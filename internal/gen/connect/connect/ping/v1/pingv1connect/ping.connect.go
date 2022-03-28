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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: connect/ping/v1/ping.proto

package pingv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/bufbuild/connect-go/internal/gen/go/connect/ping/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant isn't defined, this code was generated
// with a version of connect newer than the one compiled into your binary. You can fix the problem
// by either regenerating this code with an older version of connect or updating the connect version
// compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_0_1

const (
	// PingServiceName is the fully-qualified name of the PingService service.
	PingServiceName = "connect.ping.v1.PingService"
)

// PingServiceClient is a client for the connect.ping.v1.PingService service.
type PingServiceClient interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect_go.Request[v1.FailRequest]) (*connect_go.Response[v1.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context) *connect_go.ClientStreamForClient[v1.SumRequest, v1.SumResponse]
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect_go.Request[v1.CountUpRequest]) (*connect_go.ServerStreamForClient[v1.CountUpResponse], error)
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context) *connect_go.BidiStreamForClient[v1.CumSumRequest, v1.CumSumResponse]
}

// NewPingServiceClient constructs a client for the connect.ping.v1.PingService service. By default,
// it uses the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests.
// It doesn't have a default protocol; you must supply either the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPingServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) (PingServiceClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	pingClient, err := connect_go.NewClient[v1.PingRequest, v1.PingResponse](
		httpClient,
		baseURL+"/connect.ping.v1.PingService/Ping",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	failClient, err := connect_go.NewClient[v1.FailRequest, v1.FailResponse](
		httpClient,
		baseURL+"/connect.ping.v1.PingService/Fail",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	sumClient, err := connect_go.NewClient[v1.SumRequest, v1.SumResponse](
		httpClient,
		baseURL+"/connect.ping.v1.PingService/Sum",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	countUpClient, err := connect_go.NewClient[v1.CountUpRequest, v1.CountUpResponse](
		httpClient,
		baseURL+"/connect.ping.v1.PingService/CountUp",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	cumSumClient, err := connect_go.NewClient[v1.CumSumRequest, v1.CumSumResponse](
		httpClient,
		baseURL+"/connect.ping.v1.PingService/CumSum",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return &pingServiceClient{
		ping:    pingClient,
		fail:    failClient,
		sum:     sumClient,
		countUp: countUpClient,
		cumSum:  cumSumClient,
	}, nil
}

// pingServiceClient implements PingServiceClient.
type pingServiceClient struct {
	ping    *connect_go.Client[v1.PingRequest, v1.PingResponse]
	fail    *connect_go.Client[v1.FailRequest, v1.FailResponse]
	sum     *connect_go.Client[v1.SumRequest, v1.SumResponse]
	countUp *connect_go.Client[v1.CountUpRequest, v1.CountUpResponse]
	cumSum  *connect_go.Client[v1.CumSumRequest, v1.CumSumResponse]
}

// Ping calls connect.ping.v1.PingService.Ping.
func (c *pingServiceClient) Ping(ctx context.Context, req *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// Fail calls connect.ping.v1.PingService.Fail.
func (c *pingServiceClient) Fail(ctx context.Context, req *connect_go.Request[v1.FailRequest]) (*connect_go.Response[v1.FailResponse], error) {
	return c.fail.CallUnary(ctx, req)
}

// Sum calls connect.ping.v1.PingService.Sum.
func (c *pingServiceClient) Sum(ctx context.Context) *connect_go.ClientStreamForClient[v1.SumRequest, v1.SumResponse] {
	return c.sum.CallClientStream(ctx)
}

// CountUp calls connect.ping.v1.PingService.CountUp.
func (c *pingServiceClient) CountUp(ctx context.Context, req *connect_go.Request[v1.CountUpRequest]) (*connect_go.ServerStreamForClient[v1.CountUpResponse], error) {
	return c.countUp.CallServerStream(ctx, req)
}

// CumSum calls connect.ping.v1.PingService.CumSum.
func (c *pingServiceClient) CumSum(ctx context.Context) *connect_go.BidiStreamForClient[v1.CumSumRequest, v1.CumSumResponse] {
	return c.cumSum.CallBidiStream(ctx)
}

// PingServiceHandler is an implementation of the connect.ping.v1.PingService service.
type PingServiceHandler interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect_go.Request[v1.FailRequest]) (*connect_go.Response[v1.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context, *connect_go.ClientStream[v1.SumRequest, v1.SumResponse]) error
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect_go.Request[v1.CountUpRequest], *connect_go.ServerStream[v1.CountUpResponse]) error
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context, *connect_go.BidiStream[v1.CumSumRequest, v1.CumSumResponse]) error
}

// NewPingServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the gRPC and gRPC-Web protocols with the binary Protobuf and JSON
// codecs.
func NewPingServiceHandler(svc PingServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/connect.ping.v1.PingService/Ping", connect_go.NewUnaryHandler(
		"/connect.ping.v1.PingService/Ping",
		svc.Ping,
		opts...,
	))
	mux.Handle("/connect.ping.v1.PingService/Fail", connect_go.NewUnaryHandler(
		"/connect.ping.v1.PingService/Fail",
		svc.Fail,
		opts...,
	))
	mux.Handle("/connect.ping.v1.PingService/Sum", connect_go.NewClientStreamHandler(
		"/connect.ping.v1.PingService/Sum",
		svc.Sum,
		opts...,
	))
	mux.Handle("/connect.ping.v1.PingService/CountUp", connect_go.NewServerStreamHandler(
		"/connect.ping.v1.PingService/CountUp",
		svc.CountUp,
		opts...,
	))
	mux.Handle("/connect.ping.v1.PingService/CumSum", connect_go.NewBidiStreamHandler(
		"/connect.ping.v1.PingService/CumSum",
		svc.CumSum,
		opts...,
	))
	return "/connect.ping.v1.PingService/", mux
}

// UnimplementedPingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPingServiceHandler struct{}

func (UnimplementedPingServiceHandler) Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("connect.ping.v1.PingService.Ping isn't implemented"))
}

func (UnimplementedPingServiceHandler) Fail(context.Context, *connect_go.Request[v1.FailRequest]) (*connect_go.Response[v1.FailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("connect.ping.v1.PingService.Fail isn't implemented"))
}

func (UnimplementedPingServiceHandler) Sum(context.Context, *connect_go.ClientStream[v1.SumRequest, v1.SumResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("connect.ping.v1.PingService.Sum isn't implemented"))
}

func (UnimplementedPingServiceHandler) CountUp(context.Context, *connect_go.Request[v1.CountUpRequest], *connect_go.ServerStream[v1.CountUpResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("connect.ping.v1.PingService.CountUp isn't implemented"))
}

func (UnimplementedPingServiceHandler) CumSum(context.Context, *connect_go.BidiStream[v1.CumSumRequest, v1.CumSumResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("connect.ping.v1.PingService.CumSum isn't implemented"))
}
