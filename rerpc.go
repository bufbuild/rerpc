package rerpc

import (
	"fmt"
	"runtime"
)

// Version is the semantic version of the reRPC module.
const Version = "0.0.1"

// MaxHeaderBytes is 8KiB, gRPC's recommended maximum header size. To enforce
// this limit, set MaxHeaderBytes on your http.Server.
const MaxHeaderBytes = 1024 * 8

// ReRPC's supported HTTP content types. The gRPC variants follow gRPC's
// HTTP/2 protocol, while the Twirp and JSON variants follow Twirp's protocol.
// See reRPC's PROTOCOL.md for details.
const (
	TypeDefaultGRPC = "application/grpc"
	TypeProtoGRPC   = "application/grpc+proto"
	TypeProtoTwirp  = "application/protobuf"
	TypeJSON        = "application/json"
)

// ReRPC's supported compression methods.
//
// Snappy isn't supported because (a) the first-party gRPC libraries don't
// support it, and (b) github.com/golang/snappy doesn't have a stable release.
const (
	CompressionIdentity = "identity"
	CompressionGzip     = "gzip"
)

// For compile-time handshake with generated code.
const (
	SupportsCodeGenV0 = iota
)

var userAgent = fmt.Sprintf("grpc-go-rerpc/%s (%s)", Version, runtime.Version())

// UserAgent describes reRPC to servers, following the convention outlined in
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#user-agents.
func UserAgent() string {
	return userAgent
}
