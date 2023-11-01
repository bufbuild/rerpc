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

package connect

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// flagEnvelopeCompressed indicates that the data is compressed. It has the
// same meaning in the gRPC-Web, gRPC-HTTP2, and Connect protocols.
const flagEnvelopeCompressed = 0b00000001

var errSpecialEnvelope = errorf(
	CodeUnknown,
	"final message has protocol-specific flags: %w",
	// User code checks for end of stream with errors.Is(err, io.EOF).
	io.EOF,
)

// envelope is a block of arbitrary bytes wrapped in gRPC and Connect's framing
// protocol.
//
// Each message is preceded by a 5-byte prefix. The first byte is a uint8 used
// as a set of bitwise flags, and the remainder is a uint32 indicating the
// message length. gRPC and Connect interpret the bitwise flags differently, so
// envelope leaves their interpretation up to the caller.
type envelope struct {
	Data  *bytes.Buffer
	Flags uint8

	offset int
}

func (e *envelope) IsSet(flag uint8) bool {
	return e.Flags&flag == flag
}

// Read implements io.Reader.
func (e *envelope) Read(data []byte) (readN int, err error) {
	if e.offset < 5 {
		prefix := makeEnvelopePrefix(e.Flags, e.Data.Len())
		readN = copy(data, prefix[e.offset:])
		e.offset += readN
		if e.offset < 5 {
			return readN, nil
		}
		data = data[readN:]
	}
	n := copy(data, e.Data.Bytes()[e.offset-5:])
	e.offset += n
	readN += n
	if readN == 0 && e.offset == e.Data.Len()+5 {
		err = io.EOF
	}
	return readN, err
}

// WriteTo implements io.WriterTo.
func (e *envelope) WriteTo(dst io.Writer) (wroteN int64, err error) {
	if e.offset < 5 {
		prefix := makeEnvelopePrefix(e.Flags, e.Data.Len())
		prefixN, err := dst.Write(prefix[e.offset:])
		e.offset += prefixN
		wroteN += int64(prefixN)
		if e.offset < 5 {
			return wroteN, err
		}
	}
	n, err := dst.Write(e.Data.Bytes()[e.offset-5:])
	e.offset += n
	wroteN += int64(n)
	return wroteN, err
}

type envelopeWriter struct {
	codec            Codec
	compressMinBytes int
	compressionPool  *compressionPool
	bufferPool       *bufferPool
	sendMaxBytes     int
}

func (w *envelopeWriter) Marshal(dst io.Writer, message any) *Error {
	if message == nil {
		if _, err := dst.Write(nil); err != nil {
			if connectErr, ok := asError(err); ok {
				return connectErr
			}
			return NewError(CodeUnknown, err)
		}
		return nil
	}
	buffer, err := w.marshal(message)
	if err != nil {
		return err
	}
	defer w.bufferPool.Put(buffer)
	return w.Write(dst, &envelope{Data: buffer, Flags: 0})
}

// Write writes the enveloped message, compressing as necessary. It doesn't
// retain any references to the supplied envelope or its underlying data.
func (w *envelopeWriter) Write(dst io.Writer, env *envelope) *Error {
	if !env.IsSet(flagEnvelopeCompressed) &&
		w.compressionPool != nil &&
		env.Data.Len() > w.compressMinBytes {
		if err := w.compress(env.Data); err != nil {
			return err
		}
		env.Flags |= flagEnvelopeCompressed
	}
	return w.write(dst, env)
}

func (w *envelopeWriter) marshal(message any) (*bytes.Buffer, *Error) {
	if appender, ok := w.codec.(marshalAppender); ok {
		return w.marshalAppend(message, appender)
	}
	return w.marshalBase(message)
}

func (w *envelopeWriter) marshalAppend(message any, codec marshalAppender) (*bytes.Buffer, *Error) {
	// Codec supports MarshalAppend; try to re-use a []byte from the pool.
	buffer := w.bufferPool.Get()
	raw, err := codec.MarshalAppend(buffer.Bytes(), message)
	if err != nil {
		w.bufferPool.Put(buffer)
		return nil, errorf(CodeInternal, "marshal message: %w", err)
	}
	if cap(raw) > buffer.Cap() {
		// The buffer from the pool was too small, so MarshalAppend grew the slice.
		// Pessimistically assume that the too-small buffer is insufficient for the
		// application workload, so there's no point in keeping it in the pool.
		// Instead, replace it with the larger, newly-allocated slice. This
		// allocates, but it's a small, constant-size allocation.
		*buffer = *bytes.NewBuffer(raw)
	} else {
		// MarshalAppend didn't allocate, but we need to fix the internal state of
		// the buffer. Compared to replacing the buffer (as above), buffer.Write
		// copies but avoids allocating.
		buffer.Write(raw)
	}
	return buffer, nil
}

func (w *envelopeWriter) marshalBase(message any) (*bytes.Buffer, *Error) {
	// Codec doesn't support MarshalAppend; let Marshal allocate a []byte.
	raw, err := w.codec.Marshal(message)
	if err != nil {
		return nil, errorf(CodeInternal, "marshal message: %w", err)
	}
	return bytes.NewBuffer(raw), nil
}

func (w *envelopeWriter) compress(buffer *bytes.Buffer) *Error {
	compressed := w.bufferPool.Get()
	defer w.bufferPool.Put(compressed)
	if err := w.compressionPool.Compress(compressed, buffer); err != nil {
		return err
	}
	*buffer, *compressed = *compressed, *buffer // Swap buffer contents.
	return nil
}

func (w *envelopeWriter) checkSize(env *envelope) *Error {
	if w.sendMaxBytes > 0 && env.Data.Len() > w.sendMaxBytes {
		str := "message"
		if env.IsSet(flagEnvelopeCompressed) {
			str = "compressed message"
		}
		return errorf(CodeResourceExhausted,
			"%s size %d exceeds sendMaxBytes %d",
			str, env.Data.Len(), w.sendMaxBytes)
	}
	return nil
}

func (w *envelopeWriter) write(dst io.Writer, env *envelope) *Error {
	if err := w.checkSize(env); err != nil {
		return err
	}
	if _, err := env.WriteTo(dst); err != nil {
		if connectErr, ok := asError(err); ok {
			return connectErr
		}
		return errorf(CodeUnknown, "write message: %w", err)
	}
	return nil
}

type envelopeReader struct {
	codec           Codec
	last            envelope
	compressionPool *compressionPool
	bufferPool      *bufferPool
	readMaxBytes    int
}

func (r *envelopeReader) Unmarshal(message any, src io.Reader) *Error {
	buffer := r.bufferPool.Get()
	defer r.bufferPool.Put(buffer)

	env := &envelope{Data: buffer}
	err := r.Read(env, src)
	switch {
	case err == nil &&
		(env.Flags == 0 || env.Flags == flagEnvelopeCompressed) &&
		env.Data.Len() == 0:
		// This is a standard message (because none of the top 7 bits are set) and
		// there's no data, so the zero value of the message is correct.
		return nil
	case err != nil && errors.Is(err, io.EOF):
		// The stream has ended. Propagate the EOF to the caller.
		return err
	case err != nil:
		// Something's wrong.
		return err
	}

	data := env.Data
	if data.Len() > 0 && env.IsSet(flagEnvelopeCompressed) {
		if r.compressionPool == nil {
			return errorf(
				CodeInvalidArgument,
				"protocol error: sent compressed message without compression support",
			)
		}
		decompressed := r.bufferPool.Get()
		defer r.bufferPool.Put(decompressed)
		if err := r.compressionPool.Decompress(decompressed, data, int64(r.readMaxBytes)); err != nil {
			return err
		}
		data = decompressed
	}

	if env.Flags != 0 && env.Flags != flagEnvelopeCompressed {
		// Drain the rest of the stream to ensure there is no extra data.
		if n, err := discard(src); err != nil {
			return errorf(CodeInternal, "corrupt response: I/O error after end-stream message: %w", err)
		} else if n > 0 {
			return errorf(CodeInternal, "corrupt response: %d extra bytes after end of stream", n)
		}
		// One of the protocol-specific flags are set, so this is the end of the
		// stream. Save the message for protocol-specific code to process and
		// return a sentinel error. Since we've deferred functions to return env's
		// underlying buffer to a pool, we need to keep a copy.
		copiedData := make([]byte, data.Len())
		copy(copiedData, data.Bytes())
		r.last = envelope{
			Data:  bytes.NewBuffer(copiedData),
			Flags: env.Flags,
		}
		return errSpecialEnvelope
	}

	if err := r.codec.Unmarshal(data.Bytes(), message); err != nil {
		return errorf(CodeInvalidArgument, "unmarshal message: %w", err)
	}
	return nil
}

func (r *envelopeReader) Read(env *envelope, src io.Reader) *Error {
	prefixes := [5]byte{}
	// io.ReadFull reads the number of bytes requested, or returns an error.
	// io.EOF will only be returned if no bytes were read.
	if _, err := io.ReadFull(src, prefixes[:]); err != nil {
		if errors.Is(err, io.EOF) {
			// The stream ended cleanly. That's expected, but we need to propagate an EOF
			// to the user so that they know that the stream has ended. We shouldn't
			// add any alarming text about protocol errors, though.
			return NewError(CodeUnknown, err)
		}
		// Something else has gone wrong - the stream didn't end cleanly.
		if connectErr, ok := asError(err); ok {
			return connectErr
		}
		if maxBytesErr := asMaxBytesError(err, "read 5 byte message prefix"); maxBytesErr != nil {
			// We're reading from an http.MaxBytesHandler, and we've exceeded the read limit.
			return maxBytesErr
		}
		return errorf(
			CodeInvalidArgument,
			"protocol error: incomplete envelope: %w", err,
		)
	}
	size := int64(binary.BigEndian.Uint32(prefixes[1:5]))
	if r.readMaxBytes > 0 && size > int64(r.readMaxBytes) {
		_, err := io.CopyN(io.Discard, src, size)
		if err != nil && !errors.Is(err, io.EOF) {
			return errorf(CodeUnknown, "read enveloped message: %w", err)
		}
		return errorf(CodeResourceExhausted, "message size %d is larger than configured max %d", size, r.readMaxBytes)
	}
	// We've read the prefix, so we know how many bytes to expect.
	// CopyN will return an error if it doesn't read the requested
	// number of bytes.
	if readN, err := io.CopyN(env.Data, src, size); err != nil {
		if maxBytesErr := asMaxBytesError(err, "read %d byte message", size); maxBytesErr != nil {
			// We're reading from an http.MaxBytesHandler, and we've exceeded the read limit.
			return maxBytesErr
		}
		if errors.Is(err, io.EOF) {
			// We've gotten fewer bytes than we expected, so the stream has ended
			// unexpectedly.
			return errorf(
				CodeInvalidArgument,
				"protocol error: promised %d bytes in enveloped message, got %d bytes",
				size,
				readN,
			)
		}
		return errorf(CodeUnknown, "read enveloped message: %w", err)
	}
	env.Flags = prefixes[0]
	return nil
}

func makeEnvelopePrefix(flags uint8, size int) [5]byte {
	prefix := [5]byte{}
	prefix[0] = flags
	binary.BigEndian.PutUint32(prefix[1:5], uint32(size))
	return prefix
}
