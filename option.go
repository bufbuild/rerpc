package connect

import (
	"strings"

	"github.com/bufconnect/connect/codec"
	"github.com/bufconnect/connect/compress"
)

// Option implements both ClientOption and HandlerOption, so it can be applied
// both client-side and server-side.
type Option interface {
	ClientOption
	HandlerOption
}

type replaceProcedurePrefixOption struct {
	prefix      string
	replacement string
}

// ReplaceProcedurePrefix changes the URL used to call a procedure. Typically,
// generated code sets the procedure name: for example, a protobuf procedure's
// name and URL is composed from the fully-qualified protobuf package name, the
// service name, and the method name. This option replaces a prefix of the
// procedure name with another static string. Using this option is usually a
// bad idea, but it's occasionally necessary to prevent protobuf package
// collisions. (For example, connect uses this option to serve the health and
// reflection APIs without generating runtime conflicts with grpc-go.)
//
// ReplaceProcedurePrefix doesn't change the data exposed by the reflection
// API. To prevent inconsistencies between the reflection data and the actual
// service URL, using this option disables reflection for the modified service
// (though other services can still be introspected).
func ReplaceProcedurePrefix(prefix, replacement string) Option {
	return &replaceProcedurePrefixOption{
		prefix:      prefix,
		replacement: replacement,
	}
}

func (o *replaceProcedurePrefixOption) applyToClient(cfg *clientCfg) {
	cfg.Procedure = o.transform(cfg.Procedure)
}

func (o *replaceProcedurePrefixOption) applyToHandler(cfg *handlerCfg) {
	cfg.Procedure = o.transform(cfg.Procedure)
	cfg.RegistrationName = "" // disable reflection
}

func (o *replaceProcedurePrefixOption) transform(name string) string {
	if !strings.HasPrefix(name, o.prefix) {
		return name
	}
	return o.replacement + strings.TrimPrefix(name, o.prefix)
}

type readMaxBytes struct {
	Max int64
}

// ReadMaxBytes limits the performance impact of pathologically large messages
// sent by the other party. For handlers, ReadMaxBytes limits the size of
// message that the client can send. For clients, ReadMaxBytes limits the size
// of message that the server can respond with. Limits are applied before
// decompression and apply to each protobuf message, not to the stream as a
// whole.
//
// Setting ReadMaxBytes to zero allows any message size. Both clients and
// handlers default to allowing any request size.
func ReadMaxBytes(n int64) Option {
	return &readMaxBytes{n}
}

func (o *readMaxBytes) applyToClient(cfg *clientCfg) {
	cfg.MaxResponseBytes = o.Max
}

func (o *readMaxBytes) applyToHandler(cfg *handlerCfg) {
	cfg.MaxRequestBytes = o.Max
}

type codecOption struct {
	Name  string
	Codec codec.Codec
}

// Codec registers a serialization method with a client or handler.
//
// Typically, generated code automatically supplies this option with the
// appropriate codec(s). For example, handlers generated from protobuf schemas
// using protoc-gen-go-connect automatically register binary and JSON codecs.
// Users with more specialized needs may override the default codecs by
// registering a new codec under the same name.
//
// Handlers may have multiple codecs registered, and use whichever the client
// chooses. Clients may only have a single codec.
//
// When registering protocol buffer codecs, take care to use connect's
// protobuf.NameBinary ("protobuf") rather than "proto".
func Codec(name string, c codec.Codec) Option {
	return &codecOption{
		Name:  name,
		Codec: c,
	}
}

func (o *codecOption) applyToClient(cfg *clientCfg) {
	cfg.Codec = o.Codec
	cfg.CodecName = o.Name
}

func (o *codecOption) applyToHandler(cfg *handlerCfg) {
	if o.Codec == nil {
		delete(cfg.Codecs, o.Name)
		return
	}
	cfg.Codecs[o.Name] = o.Codec
}

type compressorOption struct {
	Name       string
	Compressor compress.Compressor
}

// Compressor configures client and server compression strategies.
//
// For handlers, it registers a compression algorithm. Clients may send
// messages compressed with that algorithm and/or request compressed responses.
// By default, handlers support gzip (using the standard library), compressing
// response messages if the client supports it and the uncompressed message is
// >1KiB.
//
// For clients, registering compressors serves two purposes. First, the client
// asks servers to compress responses using one of the registered algorithms.
// (Note that gRPC's compression negotiation is complex, but most of Google's
// gRPC server implementations won't compress responses unless the request is
// compressed.) Second, it makes all the registered algorithms available for
// use with UseCompressor. Note that actually compressing requests requires
// using both Compressor and UseCompressor.
//
// To remove a previously-registered compressor, re-register the same name with
// a nil compressor.
func Compressor(name string, c compress.Compressor) Option {
	return &compressorOption{
		Name:       name,
		Compressor: c,
	}
}

func (o *compressorOption) applyToClient(cfg *clientCfg) {
	o.apply(cfg.Compressors)
}

func (o *compressorOption) applyToHandler(cfg *handlerCfg) {
	o.apply(cfg.Compressors)
}

func (o *compressorOption) apply(m map[string]compress.Compressor) {
	if o.Compressor == nil {
		delete(m, o.Name)
		return
	}
	m[o.Name] = o.Compressor
}

type interceptOption struct {
	interceptors []Interceptor
}

// Interceptors configures a client or handler's interceptor stack. Repeated
// Interceptors options are applied in order, so
//
//   Interceptors(A) + Interceptors(B, C) == Interceptors(A, B, C)
//
// Unary interceptors compose like an onion. The first interceptor provided is
// the outermost layer of the onion: it acts first on the context and request,
// and last on the response and error.
//
// Stream interceptors also behave like an onion: the first interceptor
// provided is the first to wrap the context and is the outermost wrapper for
// the (Sender, Receiver) pair. It's the first to see sent messages and the
// last to see received messages.
//
// Applied to client and handler, Interceptors(A, B, ..., Y, Z) produces:
//
//        client.Send()     client.Receive()
//              |                 ^
//              v                 |
//           A ---               --- A
//           B ---               --- B
//             ...               ...
//           Y ---               --- Y
//           Z ---               --- Z
//              |                 ^
//              v                 |
//           network            network
//              |                 ^
//              v                 |
//           A ---               --- A
//           B ---               --- B
//             ...               ...
//           Y ---               --- Y
//           Z ---               --- Z
//              |                 ^
//              v                 |
//       handler.Receive() handler.Send()
//              |                 ^
//              |                 |
//              -> handler logic --
//
// Note that in clients, the Sender handles the request message(s) and the
// Receiver handles the response message(s). For handlers, it's the reverse.
// Depending on your interceptor's logic, you may need to wrap one side of the
// stream on the clients and the other side on handlers. See the implementation
// of HeaderInterceptor for an example.
func Interceptors(interceptors ...Interceptor) Option {
	return &interceptOption{interceptors}
}

func (o *interceptOption) applyToClient(cfg *clientCfg) {
	cfg.Interceptor = o.chainWith(cfg.Interceptor)
}

func (o *interceptOption) applyToHandler(cfg *handlerCfg) {
	cfg.Interceptor = o.chainWith(cfg.Interceptor)
}

func (o *interceptOption) chainWith(current Interceptor) Interceptor {
	if len(o.interceptors) == 0 {
		return current
	}
	if current == nil && len(o.interceptors) == 1 {
		return o.interceptors[0]
	}
	if current == nil && len(o.interceptors) > 1 {
		return newChain(o.interceptors)
	}
	return newChain(append([]Interceptor{current}, o.interceptors...))
}
