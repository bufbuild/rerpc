// Copyright 2021-2023 Buf Technologies, Inc.
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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: connect/import/v1/import.proto

package importv1connect

import (
	connect_go "github.com/bufbuild/connect-go"
	_ "github.com/bufbuild/connect-go/internal/gen/connect/import/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ImportServiceName is the fully-qualified name of the ImportService service.
	ImportServiceName = "connect.import.v1.ImportService"
)

type ()

// ImportServiceClient is a client for the connect.import.v1.ImportService service.
type ImportServiceClient interface {
}

// NewImportServiceClient constructs a client for the connect.import.v1.ImportService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewImportServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ImportServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &importServiceClient{}
}

// importServiceClient implements ImportServiceClient.
type importServiceClient struct {
}

// ImportServiceHandler is an implementation of the connect.import.v1.ImportService service.
type ImportServiceHandler interface {
}

// NewImportServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewImportServiceHandler(svc ImportServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	return "/connect.import.v1.ImportService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedImportServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedImportServiceHandler struct{}
