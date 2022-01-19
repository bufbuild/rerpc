// Code generated by protoc-gen-go-rerpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-rerpc v0.0.1
// - protoc              v3.17.3
// source: rerpc/ping/v1test/ping.proto

package pingv1test

import (
	context "context"
	errors "errors"
	rerpc "github.com/rerpc/rerpc"
	clientstream "github.com/rerpc/rerpc/clientstream"
	handlerstream "github.com/rerpc/rerpc/handlerstream"
	v1test "github.com/rerpc/rerpc/internal/gen/proto/go/rerpc/ping/v1test"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the
// rerpc package are compatible. If you get a compiler error that this constant
// isn't defined, this code was generated with a version of rerpc newer than the
// one compiled into your binary. You can fix the problem by either regenerating
// this code with an older version of rerpc or updating the rerpc version
// compiled into your binary.
const _ = rerpc.SupportsCodeGenV0 // requires reRPC v0.0.1 or later

// SimplePingServiceClient is a client for the rerpc.ping.v1test.PingService
// service.
type SimplePingServiceClient interface {
	Ping(context.Context, *v1test.PingRequest) (*v1test.PingResponse, error)
	Fail(context.Context, *v1test.FailRequest) (*v1test.FailResponse, error)
	Sum(context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse]
	CountUp(context.Context, *v1test.CountUpRequest) (*clientstream.Server[v1test.CountUpResponse], error)
	CumSum(context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]
}

// FullPingServiceClient is a client for the rerpc.ping.v1test.PingService
// service. It's more complex than SimplePingServiceClient, but it gives callers
// more fine-grained control (e.g., sending and receiving headers).
type FullPingServiceClient interface {
	Ping(context.Context, *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error)
	Fail(context.Context, *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error)
	Sum(context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse]
	CountUp(context.Context, *rerpc.Request[v1test.CountUpRequest]) (*clientstream.Server[v1test.CountUpResponse], error)
	CumSum(context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]
}

// PingServiceClient is a client for the rerpc.ping.v1test.PingService service.
type PingServiceClient struct {
	client fullPingServiceClient
}

var _ SimplePingServiceClient = (*PingServiceClient)(nil)

// NewPingServiceClient constructs a client for the
// rerpc.ping.v1test.PingService service.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewPingServiceClient(baseURL string, doer rerpc.Doer, opts ...rerpc.ClientOption) (*PingServiceClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	pingFunc, err := rerpc.NewClientFunc[v1test.PingRequest, v1test.PingResponse](
		doer,
		baseURL,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"Ping",              // protobuf method
		opts...,
	)
	if err != nil {
		return nil, err
	}
	failFunc, err := rerpc.NewClientFunc[v1test.FailRequest, v1test.FailResponse](
		doer,
		baseURL,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"Fail",              // protobuf method
		opts...,
	)
	if err != nil {
		return nil, err
	}
	sumFunc, err := rerpc.NewClientStream(
		doer,
		rerpc.StreamTypeClient,
		baseURL,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"Sum",               // protobuf method
		opts...,
	)
	if err != nil {
		return nil, err
	}
	countUpFunc, err := rerpc.NewClientStream(
		doer,
		rerpc.StreamTypeServer,
		baseURL,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"CountUp",           // protobuf method
		opts...,
	)
	if err != nil {
		return nil, err
	}
	cumSumFunc, err := rerpc.NewClientStream(
		doer,
		rerpc.StreamTypeBidirectional,
		baseURL,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"CumSum",            // protobuf method
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return &PingServiceClient{client: fullPingServiceClient{
		ping:    pingFunc,
		fail:    failFunc,
		sum:     sumFunc,
		countUp: countUpFunc,
		cumSum:  cumSumFunc,
	}}, nil
}

// Ping calls rerpc.ping.v1test.PingService.Ping.
func (c *PingServiceClient) Ping(ctx context.Context, req *v1test.PingRequest) (*v1test.PingResponse, error) {
	res, err := c.client.Ping(ctx, rerpc.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// Fail calls rerpc.ping.v1test.PingService.Fail.
func (c *PingServiceClient) Fail(ctx context.Context, req *v1test.FailRequest) (*v1test.FailResponse, error) {
	res, err := c.client.Fail(ctx, rerpc.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// Sum calls rerpc.ping.v1test.PingService.Sum.
func (c *PingServiceClient) Sum(ctx context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse] {
	return c.client.Sum(ctx)
}

// CountUp calls rerpc.ping.v1test.PingService.CountUp.
func (c *PingServiceClient) CountUp(ctx context.Context, req *v1test.CountUpRequest) (*clientstream.Server[v1test.CountUpResponse], error) {
	return c.client.CountUp(ctx, rerpc.NewRequest(req))
}

// CumSum calls rerpc.ping.v1test.PingService.CumSum.
func (c *PingServiceClient) CumSum(ctx context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse] {
	return c.client.CumSum(ctx)
}

// Full exposes the underlying generic client. Use it if you need finer control
// (e.g., sending and receiving headers).
func (c *PingServiceClient) Full() FullPingServiceClient {
	return &c.client
}

type fullPingServiceClient struct {
	ping    func(context.Context, *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error)
	fail    func(context.Context, *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error)
	sum     rerpc.StreamFunc
	countUp rerpc.StreamFunc
	cumSum  rerpc.StreamFunc
}

var _ FullPingServiceClient = (*fullPingServiceClient)(nil)

// Ping calls rerpc.ping.v1test.PingService.Ping.
func (c *fullPingServiceClient) Ping(ctx context.Context, req *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error) {
	return c.ping(ctx, req)
}

// Fail calls rerpc.ping.v1test.PingService.Fail.
func (c *fullPingServiceClient) Fail(ctx context.Context, req *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error) {
	return c.fail(ctx, req)
}

// Sum calls rerpc.ping.v1test.PingService.Sum.
func (c *fullPingServiceClient) Sum(ctx context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse] {
	_, sender, receiver := c.sum(ctx)
	return clientstream.NewClient[v1test.SumRequest, v1test.SumResponse](sender, receiver)
}

// CountUp calls rerpc.ping.v1test.PingService.CountUp.
func (c *fullPingServiceClient) CountUp(ctx context.Context, req *rerpc.Request[v1test.CountUpRequest]) (*clientstream.Server[v1test.CountUpResponse], error) {
	_, sender, receiver := c.countUp(ctx)
	if err := sender.Send(req.Any()); err != nil {
		_ = sender.Close(err)
		_ = receiver.Close()
		return nil, err
	}
	if err := sender.Close(nil); err != nil {
		_ = receiver.Close()
		return nil, err
	}
	return clientstream.NewServer[v1test.CountUpResponse](receiver), nil
}

// CumSum calls rerpc.ping.v1test.PingService.CumSum.
func (c *fullPingServiceClient) CumSum(ctx context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse] {
	_, sender, receiver := c.cumSum(ctx)
	return clientstream.NewBidirectional[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
}

// FullPingServiceServer is a server for the rerpc.ping.v1test.PingService
// service.
type FullPingServiceServer interface {
	Ping(context.Context, *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error)
	Fail(context.Context, *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error)
	Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	CountUp(context.Context, *rerpc.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error
	CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

// SimplePingServiceServer is a server for the rerpc.ping.v1test.PingService
// service. It's a simpler interface than FullPingServiceServer but doesn't
// provide header access.
type SimplePingServiceServer interface {
	Ping(context.Context, *v1test.PingRequest) (*v1test.PingResponse, error)
	Fail(context.Context, *v1test.FailRequest) (*v1test.FailResponse, error)
	Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	CountUp(context.Context, *v1test.CountUpRequest, *handlerstream.Server[v1test.CountUpResponse]) error
	CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

// NewFullPingServiceHandler wraps each method on the service implementation in
// a rerpc.Handler. The returned slice can be passed to rerpc.NewServeMux.
func NewFullPingServiceHandler(svc FullPingServiceServer, opts ...rerpc.HandlerOption) []rerpc.Handler {
	handlers := make([]rerpc.Handler, 0, 5)

	ping := rerpc.NewUnaryHandler(
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"Ping",              // protobuf method
		svc.Ping,
		opts...,
	)
	handlers = append(handlers, *ping)

	fail := rerpc.NewUnaryHandler(
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"Fail",              // protobuf method
		svc.Fail,
		opts...,
	)
	handlers = append(handlers, *fail)

	sum := rerpc.NewStreamingHandler(
		rerpc.StreamTypeClient,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"Sum",               // protobuf method
		func(ctx context.Context, sender rerpc.Sender, receiver rerpc.Receiver) {
			typed := handlerstream.NewClient[v1test.SumRequest, v1test.SumResponse](sender, receiver)
			err := svc.Sum(ctx, typed)
			_ = receiver.Close()
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = sender.Close(err)
		},
		opts...,
	)
	handlers = append(handlers, *sum)

	countUp := rerpc.NewStreamingHandler(
		rerpc.StreamTypeServer,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"CountUp",           // protobuf method
		func(ctx context.Context, sender rerpc.Sender, receiver rerpc.Receiver) {
			typed := handlerstream.NewServer[v1test.CountUpResponse](sender)
			req, err := rerpc.ReceiveRequest[v1test.CountUpRequest](receiver)
			if err != nil {
				_ = receiver.Close()
				_ = sender.Close(err)
				return
			}
			if err = receiver.Close(); err != nil {
				_ = sender.Close(err)
				return
			}
			err = svc.CountUp(ctx, req, typed)
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = sender.Close(err)
		},
		opts...,
	)
	handlers = append(handlers, *countUp)

	cumSum := rerpc.NewStreamingHandler(
		rerpc.StreamTypeBidirectional,
		"rerpc.ping.v1test", // protobuf package
		"PingService",       // protobuf service
		"CumSum",            // protobuf method
		func(ctx context.Context, sender rerpc.Sender, receiver rerpc.Receiver) {
			typed := handlerstream.NewBidirectional[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
			err := svc.CumSum(ctx, typed)
			_ = receiver.Close()
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = sender.Close(err)
		},
		opts...,
	)
	handlers = append(handlers, *cumSum)

	return handlers
}

type pluggablePingServiceServer struct {
	ping    func(context.Context, *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error)
	fail    func(context.Context, *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error)
	sum     func(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	countUp func(context.Context, *rerpc.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error
	cumSum  func(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

func (s *pluggablePingServiceServer) Ping(ctx context.Context, req *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error) {
	return s.ping(ctx, req)
}

func (s *pluggablePingServiceServer) Fail(ctx context.Context, req *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error) {
	return s.fail(ctx, req)
}

func (s *pluggablePingServiceServer) Sum(ctx context.Context, stream *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error {
	return s.sum(ctx, stream)
}

func (s *pluggablePingServiceServer) CountUp(ctx context.Context, req *rerpc.Request[v1test.CountUpRequest], stream *handlerstream.Server[v1test.CountUpResponse]) error {
	return s.countUp(ctx, req, stream)
}

func (s *pluggablePingServiceServer) CumSum(ctx context.Context, stream *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error {
	return s.cumSum(ctx, stream)
}

// NewPingServiceHandler wraps each method on the service implementation in a
// rerpc.Handler. The returned slice can be passed to rerpc.NewServeMux.
//
// Unlike NewFullPingServiceHandler, it allows the service to mix and match the
// signatures of FullPingServiceServer and SimplePingServiceServer. For each
// method, it first tries to find a SimplePingServiceServer-style
// implementation. If a simple implementation isn't available, it falls back to
// the more complex FullPingServiceServer-style implementation. If neither is
// available, it returns an error.
//
// Taken together, this approach lets implementations embed
// UnimplementedPingServiceServer and implement each method using whichever
// signature is most convenient.
func NewPingServiceHandler(svc any, opts ...rerpc.HandlerOption) ([]rerpc.Handler, error) {
	var impl pluggablePingServiceServer

	// Find an implementation of Ping
	if pinger, ok := svc.(interface {
		Ping(context.Context, *v1test.PingRequest) (*v1test.PingResponse, error)
	}); ok {
		impl.ping = func(ctx context.Context, req *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error) {
			res, err := pinger.Ping(ctx, req.Msg)
			if err != nil {
				return nil, err
			}
			return rerpc.NewResponse(res), nil
		}
	} else if pinger, ok := svc.(interface {
		Ping(context.Context, *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error)
	}); ok {
		impl.ping = pinger.Ping
	} else {
		return nil, errors.New("no Ping implementation found")
	}

	// Find an implementation of Fail
	if failer, ok := svc.(interface {
		Fail(context.Context, *v1test.FailRequest) (*v1test.FailResponse, error)
	}); ok {
		impl.fail = func(ctx context.Context, req *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error) {
			res, err := failer.Fail(ctx, req.Msg)
			if err != nil {
				return nil, err
			}
			return rerpc.NewResponse(res), nil
		}
	} else if failer, ok := svc.(interface {
		Fail(context.Context, *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error)
	}); ok {
		impl.fail = failer.Fail
	} else {
		return nil, errors.New("no Fail implementation found")
	}

	// Find an implementation of Sum
	if sumer, ok := svc.(interface {
		Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	}); ok {
		impl.sum = sumer.Sum
	} else {
		return nil, errors.New("no Sum implementation found")
	}

	// Find an implementation of CountUp
	if countUper, ok := svc.(interface {
		CountUp(context.Context, *v1test.CountUpRequest, *handlerstream.Server[v1test.CountUpResponse]) error
	}); ok {
		impl.countUp = func(ctx context.Context, req *rerpc.Request[v1test.CountUpRequest], stream *handlerstream.Server[v1test.CountUpResponse]) error {
			return countUper.CountUp(ctx, req.Msg, stream)
		}
	} else if countUper, ok := svc.(interface {
		CountUp(context.Context, *rerpc.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error
	}); ok {
		impl.countUp = countUper.CountUp
	} else {
		return nil, errors.New("no CountUp implementation found")
	}

	// Find an implementation of CumSum
	if cumSumer, ok := svc.(interface {
		CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
	}); ok {
		impl.cumSum = cumSumer.CumSum
	} else {
		return nil, errors.New("no CumSum implementation found")
	}

	return NewFullPingServiceHandler(&impl, opts...), nil
}

var _ FullPingServiceServer = (*UnimplementedPingServiceServer)(nil) // verify interface implementation

// UnimplementedPingServiceServer returns CodeUnimplemented from all methods.
type UnimplementedPingServiceServer struct{}

func (UnimplementedPingServiceServer) Ping(context.Context, *rerpc.Request[v1test.PingRequest]) (*rerpc.Response[v1test.PingResponse], error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "rerpc.ping.v1test.PingService.Ping isn't implemented")
}

func (UnimplementedPingServiceServer) Fail(context.Context, *rerpc.Request[v1test.FailRequest]) (*rerpc.Response[v1test.FailResponse], error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "rerpc.ping.v1test.PingService.Fail isn't implemented")
}

func (UnimplementedPingServiceServer) Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "rerpc.ping.v1test.PingService.Sum isn't implemented")
}

func (UnimplementedPingServiceServer) CountUp(context.Context, *rerpc.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "rerpc.ping.v1test.PingService.CountUp isn't implemented")
}

func (UnimplementedPingServiceServer) CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "rerpc.ping.v1test.PingService.CumSum isn't implemented")
}
