// Copyright 2021-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connect

import (
	"errors"
	"net/http"
)

// Thankfully, the handler stream is much simpler than the client. net/http
// gives us the request body and response writer at the same time, so we don't
// need to worry about concurrency.
func newHandlerStream(
	spec Specification,
	web bool,
	w http.ResponseWriter,
	r *http.Request,
	maxReadBytes int64,
	compressMinBytes int,
	codec Codec,
	protobuf Codec, // for errors
	requestCompressionPools compressionPool,
	responseCompressionPools compressionPool,
) (*handlerSender, *handlerReceiver) {
	sender := &handlerSender{
		spec: spec,
		web:  web,
		marshaler: marshaler{
			writer:           w,
			compressionPool:  responseCompressionPools,
			codec:            codec,
			compressMinBytes: compressMinBytes,
		},
		protobuf: protobuf,
		writer:   w,
		trailer:  make(http.Header, 3), // grpc-{status,message,status-details-bin}
	}
	receiver := &handlerReceiver{
		spec: spec,
		unmarshaler: unmarshaler{
			web:             web,
			reader:          r.Body,
			max:             maxReadBytes,
			compressionPool: requestCompressionPools,
			codec:           codec,
		},
		request: r,
	}
	return sender, receiver
}

type handlerSender struct {
	spec        Specification
	web         bool
	marshaler   marshaler
	protobuf    Codec // for errors
	writer      http.ResponseWriter
	trailer     http.Header
	wroteToBody bool
}

var _ Sender = (*handlerSender)(nil)

func (hs *handlerSender) Send(message any) error {
	defer hs.flush()
	hs.wroteToBody = true
	if !hs.web {
		// We're going to write body data, so we'll have to send gRPC's status
		// information in HTTP trailers. Since we know the trailer keys ahead of
		// time, we maximize the chance that any intervening proxies will support
		// our trailers by advertising them in the "Trailer" header.
		//
		// This doesn't apply to gRPC-Web, where we don't use HTTP trailers.
		hs.Header()["Trailer"] = []string{
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		}
	}
	if err := hs.marshaler.Marshal(message); err != nil {
		return err // already coded
	}
	// don't return typed nils
	return nil
}

func (hs *handlerSender) Close(err error) error {
	defer hs.flush()
	if connectErr, ok := asError(err); ok {
		mergeHeaders(hs.Header(), connectErr.header)
		mergeHeaders(hs.Trailer(), connectErr.trailer)
	}
	if !hs.web || !hs.wroteToBody {
		// We're using standard gRPC and/or we haven't written any data to the
		// response body. In the latter case, we should send what gRPC calls a
		// "trailers-only" response. Confusingly, gRPC's "trailers-only" response
		// puts all the data in HTTP _headers_ (even for gRPC-Web).
		for key, values := range hs.trailer {
			if hs.wroteToBody {
				// We're using standard gRPC, so we'll have to send this metadata as
				// HTTP trailers. In net/http's ResponseWriter API, we do that by
				// writing to the headers map with a special prefix.
				key = http.TrailerPrefix + key
			}
			for _, value := range values {
				hs.writer.Header().Add(key, value)
			}
		}
		return grpcErrorToTrailer(hs.writer.Header(), hs.protobuf, err)
	}
	if trailerErr := grpcErrorToTrailer(hs.trailer, hs.protobuf, err); trailerErr != nil {
		return trailerErr
	}
	if marshalErr := hs.marshaler.MarshalWebTrailers(hs.trailer); marshalErr != nil {
		return marshalErr
	}
	return nil
}

func (hs *handlerSender) Spec() Specification {
	return hs.spec
}

func (hs *handlerSender) Header() http.Header {
	return hs.writer.Header()
}

func (hs *handlerSender) Trailer() http.Header {
	return hs.trailer
}

func (hs *handlerSender) flush() {
	if f, ok := hs.writer.(http.Flusher); ok {
		f.Flush()
	}
}

type handlerReceiver struct {
	spec        Specification
	unmarshaler unmarshaler
	request     *http.Request
}

var _ Receiver = (*handlerReceiver)(nil)

func (hr *handlerReceiver) Receive(message any) error {
	if err := hr.unmarshaler.Unmarshal(message); err != nil {
		if errors.Is(err, errGotWebTrailers) {
			if hr.request.Trailer == nil {
				hr.request.Trailer = hr.unmarshaler.WebTrailer()
			} else {
				mergeHeaders(hr.request.Trailer, hr.unmarshaler.WebTrailer())
			}
		}
		return err // already coded
	}
	// don't return typed nils
	return nil
}

func (hr *handlerReceiver) Close() error {
	// We don't want to copy unread portions of the body to /dev/null here: if
	// the client hasn't closed the request body, we'll block until the server
	// timeout kicks in. This could happen because the client is malicious, but
	// a well-intentioned client may just not expect the server to be returning
	// an error for a streaming RPC. Better to accept that we can't always reuse
	// TCP connections.
	if err := hr.request.Body.Close(); err != nil {
		if connectErr, ok := asError(err); ok {
			return connectErr
		}
		return NewError(CodeUnknown, err)
	}
	return nil
}

func (hr *handlerReceiver) Spec() Specification {
	return hr.spec
}

func (hr *handlerReceiver) Header() http.Header {
	return hr.request.Header
}

func (hr *handlerReceiver) Trailer() http.Header {
	return hr.request.Trailer
}
