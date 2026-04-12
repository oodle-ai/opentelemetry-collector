// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package data // import "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data"

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// marshalJSON converts trace id into a hex string enclosed in quotes.
// Called by Protobuf JSON deserialization.
func marshalJSON(id []byte) ([]byte, error) {
	// Plus 2 quote chars at the start and end.
	hexLen := hex.EncodedLen(len(id)) + 2

	b := make([]byte, hexLen)
	hex.Encode(b[1:hexLen-1], id)
	b[0], b[hexLen-1] = '"', '"'

	return b, nil
}

// unmarshalOTLPJSONBytesID decodes a fixed-size ID from a
// hex-encoded JSON string, possibly enclosed in quotes.
// size is the expected decoded byte count (16 for TraceID,
// 8 for SpanID). dst is re-allocated or resliced as needed
// to hold exactly size bytes.
func unmarshalOTLPJSONBytesID(
	dst *[]byte, size int, src []byte,
) error {
	if cap(*dst) < size {
		*dst = make([]byte, size)
	} else {
		*dst = (*dst)[:size]
	}
	if l := len(src); l >= 2 &&
		src[0] == '"' && src[l-1] == '"' {
		src = src[1 : l-1]
	}
	if len(src) == 0 {
		return nil
	}
	if len(src) != hex.EncodedLen(size) {
		return errors.New("invalid length for ID")
	}
	_, err := hex.Decode(*dst, src)
	if err != nil {
		return fmt.Errorf(
			"cannot unmarshal ID from string '%s': %w",
			string(src), err,
		)
	}
	return nil
}
