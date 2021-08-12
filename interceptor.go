package rerpc

import (
	"context"

	"google.golang.org/protobuf/proto"
)

// Func is the generic signature of a unary RPC, from both the server and the
// client's perspective. Interceptors wrap Funcs.
type Func func(context.Context, proto.Message) (proto.Message, error)

// HandlerStreamFunc is the generic signature of a streaming RPC from the
// server's perspective. Interceptors wrap HandlerStreamFuncs.
type HandlerStreamFunc func(context.Context, Stream)

// CallStreamFunc is the generic signature of a streaming RPC from the client's
// perspective. Interceptors wrap CallStreamFuncs.
type CallStreamFunc func(context.Context) Stream

// An Interceptor adds logic to a generated handler or client, like the
// decorators or middleware you may have seen in other libraries. Interceptors
// may replace the context, mutate the request, mutate the response, handle the
// returned error, retry, recover from panics, emit logs and metrics, or do
// nearly anything else.
//
// The returned functions must be safe to call concurrently. If chained
// carelessly, the interceptor's logic may run more than once - where possible,
// interceptors should be idempotent.
//
// See Chain for an example of interceptor use.
type Interceptor interface {
	Wrap(Func) Func
	WrapHandlerStream(HandlerStreamFunc) HandlerStreamFunc
	WrapCallStream(CallStreamFunc) CallStreamFunc
}

// ConfiguredCallInterceptor returns the Interceptor configured by a collection
// of call options (if any). It's used in generated code.
func ConfiguredCallInterceptor(opts ...CallOption) Interceptor {
	var cfg callCfg
	for _, o := range opts {
		o.applyToCall(&cfg)
	}
	return cfg.Interceptor
}

// ConfiguredHandlerInterceptor returns the Interceptor configured by a collection
// of handler options (if any). It's used in generated code.
func ConfiguredHandlerInterceptor(opts ...HandlerOption) Interceptor {
	var cfg handlerCfg
	for _, o := range opts {
		o.applyToHandler(&cfg)
	}
	return cfg.Interceptor
}

// A UnaryInterceptorFunc is a simple Interceptor implementation that only
// wraps unary RPCs. See CallMetadata for an example.
type UnaryInterceptorFunc func(Func) Func

// Wrap implements Interceptor by applying the interceptor function.
func (f UnaryInterceptorFunc) Wrap(next Func) Func { return f(next) }

// WrapHandlerStream implements Interceptor with a no-op.
func (f UnaryInterceptorFunc) WrapHandlerStream(next HandlerStreamFunc) HandlerStreamFunc {
	return next
}

// WrapCallStream implements Interceptor with a no-op.
func (f UnaryInterceptorFunc) WrapCallStream(next CallStreamFunc) CallStreamFunc {
	return next
}

// A Chain composes multiple interceptors into one.
type Chain struct {
	interceptors []Interceptor
}

var _ Interceptor = (*Chain)(nil)

// NewChain composes multiple interceptors into one. The first interceptor
// provided is the outermost layer of the onion: it acts first on the context
// and request, and last on the response and error.
func NewChain(interceptors ...Interceptor) *Chain {
	return &Chain{interceptors}
}

// Wrap implements Interceptor.
func (c *Chain) Wrap(next Func) Func {
	// We need to wrap in reverse order to have the first interceptor from
	// the slice act first.
	for i := len(c.interceptors) - 1; i >= 0; i-- {
		if interceptor := c.interceptors[i]; interceptor != nil {
			next = interceptor.Wrap(next)
		}
	}
	return next
}

// WrapHandlerStream implements Interceptor.
func (c *Chain) WrapHandlerStream(next HandlerStreamFunc) HandlerStreamFunc {
	for i := len(c.interceptors) - 1; i >= 0; i-- {
		if interceptor := c.interceptors[i]; interceptor != nil {
			next = interceptor.WrapHandlerStream(next)
		}
	}
	return next
}

// WrapCallStream implements Interceptor.
func (c *Chain) WrapCallStream(next CallStreamFunc) CallStreamFunc {
	// We need to wrap in reverse order to have the first interceptor from
	// the slice act first.
	for i := len(c.interceptors) - 1; i >= 0; i-- {
		if interceptor := c.interceptors[i]; interceptor != nil {
			next = interceptor.WrapCallStream(next)
		}
	}
	return next
}
