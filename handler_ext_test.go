// Copyright 2021-2023 The Connect Authors
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

package connect_test

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	connect "connectrpc.com/connect"
	"connectrpc.com/connect/internal/assert"
	"connectrpc.com/connect/internal/connecttest"
	pingv1 "connectrpc.com/connect/internal/gen/connect/ping/v1"
	"connectrpc.com/connect/internal/gen/connect/ping/v1/pingv1connect"
)

func TestHandler_ServeHTTP(t *testing.T) {
	t.Parallel()
	path, handler := pingv1connect.NewPingServiceHandler(successPingServer{})
	prefixed := http.NewServeMux()
	prefixed.Handle(path, handler)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	mux.Handle("/prefixed/", http.StripPrefix("/prefixed", prefixed))
	const pingProcedure = pingv1connect.PingServicePingProcedure
	const sumProcedure = pingv1connect.PingServiceSumProcedure
	server := connecttest.StartHTTPTestServer(t, mux)
	client := server.Client()

	t.Run("get_method_no_encoding", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			server.URL+pingProcedure,
			strings.NewReader(""),
		)
		assert.Nil(t, err)
		resp, err := client.Do(request)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusUnsupportedMediaType)
	})

	t.Run("get_method_bad_encoding", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			server.URL+pingProcedure+`?encoding=unk&message={}`,
			strings.NewReader(""),
		)
		assert.Nil(t, err)
		resp, err := client.Do(request)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusUnsupportedMediaType)
	})

	t.Run("idempotent_get_method", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			server.URL+pingProcedure+`?encoding=json&message={}`,
			strings.NewReader(""),
		)
		assert.Nil(t, err)
		resp, err := client.Do(request)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("prefixed_get_method", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			server.URL+"/prefixed"+pingProcedure+`?encoding=json&message={}`,
			strings.NewReader(""),
		)
		assert.Nil(t, err)
		resp, err := client.Do(request)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("method_not_allowed", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			server.URL+sumProcedure,
			strings.NewReader(""),
		)
		assert.Nil(t, err)
		resp, err := client.Do(request)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusMethodNotAllowed)
		assert.Equal(t, resp.Header.Get("Allow"), http.MethodPost)
	})

	t.Run("unsupported_content_type", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			server.URL+pingProcedure,
			strings.NewReader("{}"),
		)
		assert.Nil(t, err)
		request.Header.Set("Content-Type", "application/x-custom-json")
		resp, err := client.Do(request)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusUnsupportedMediaType)
		assert.Equal(t, resp.Header.Get("Accept-Post"), strings.Join([]string{
			"application/grpc",
			"application/grpc+json",
			"application/grpc+json; charset=utf-8",
			"application/grpc+proto",
			"application/grpc-web",
			"application/grpc-web+json",
			"application/grpc-web+json; charset=utf-8",
			"application/grpc-web+proto",
			"application/json",
			"application/json; charset=utf-8",
			"application/proto",
		}, ", "))
	})

	t.Run("charset_in_content_type_header", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			server.URL+pingProcedure,
			strings.NewReader("{}"),
		)
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json;Charset=Utf-8")
		resp, err := client.Do(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("unsupported_charset", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			server.URL+pingProcedure,
			strings.NewReader("{}"),
		)
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json; charset=shift-jis")
		resp, err := client.Do(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusUnsupportedMediaType)
	})

	t.Run("unsupported_content_encoding", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			server.URL+pingProcedure,
			strings.NewReader("{}"),
		)
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Encoding", "invalid")
		resp, err := client.Do(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusNotFound)

		type errorMessage struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		}
		var message errorMessage
		err = json.NewDecoder(resp.Body).Decode(&message)
		assert.Nil(t, err)
		assert.Equal(t, message.Message, `unknown compression "invalid": supported encodings are gzip`)
		assert.Equal(t, message.Code, connect.CodeUnimplemented.String())
	})
}

func TestHandlerMaliciousPrefix(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.Handle(pingv1connect.NewPingServiceHandler(successPingServer{}))
	server := connecttest.StartHTTPTestServer(t, mux)

	const (
		concurrency  = 256
		spuriousSize = 1024 * 1024 * 512 // 512 MB
	)
	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < concurrency; i++ {
		body := make([]byte, 16)
		// Envelope prefix indicates a large payload which we're not actually
		// sending.
		binary.BigEndian.PutUint32(body[1:5], spuriousSize)
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			server.URL+pingv1connect.PingServicePingProcedure,
			bytes.NewReader(body),
		)
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/grpc")
		wg.Add(1)
		go func(req *http.Request) {
			defer wg.Done()
			<-start
			response, err := server.Client().Do(req)
			if err == nil {
				_, _ = io.Copy(io.Discard, response.Body)
				response.Body.Close()
			}
		}(req)
	}
	close(start)
	wg.Wait()
}

type successPingServer struct {
	pingv1connect.UnimplementedPingServiceHandler
}

func (successPingServer) Ping(context.Context, *connect.Request[pingv1.PingRequest]) (*connect.Response[pingv1.PingResponse], error) {
	return &connect.Response[pingv1.PingResponse]{}, nil
}
