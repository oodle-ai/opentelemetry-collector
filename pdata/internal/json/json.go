// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package json // import "github.com/oodle-ai/opentelemetry-collector/pdata/internal/json"

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

var marshaler = &jsonpb.Marshaler{
	// https://github.com/open-telemetry/opentelemetry-specification/pull/2758
	EnumsAsInts: true,
	// https://github.com/open-telemetry/opentelemetry-specification/pull/2829
	OrigName: false,
}

// OTLP JSON deviates from standard proto3 JSON: traceId,
// spanId, and parentSpanId must be hex-encoded, not base64.
// gogo's jsonpb encodes []byte as base64 (proto3 default),
// so we post-process to convert these fields to hex.
var idFieldKeys = [][]byte{
	[]byte(`"traceId"`),
	[]byte(`"spanId"`),
	[]byte(`"parentSpanId"`),
}

func Marshal(
	out io.Writer,
	pb proto.Message,
) error {
	var buf bytes.Buffer
	if err := marshaler.Marshal(&buf, pb); err != nil {
		return err
	}
	_, err := out.Write(
		fixIDFieldEncoding(buf.Bytes()),
	)
	return err
}

// fixIDFieldEncoding scans JSON bytes for traceId, spanId,
// and parentSpanId string values and converts them from
// base64 to lowercase hex per the OTLP spec.
func fixIDFieldEncoding(data []byte) []byte {
	for _, key := range idFieldKeys {
		data = replaceIDValues(data, key)
	}
	return data
}

func replaceIDValues(
	data []byte,
	key []byte,
) []byte {
	searchFrom := 0
	for {
		if searchFrom >= len(data) {
			break
		}
		idx := bytes.Index(
			data[searchFrom:], key,
		)
		if idx < 0 {
			break
		}
		idx += searchFrom

		// Skip past key and find the colon then the
		// opening quote of the value.
		pos := idx + len(key)
		for pos < len(data) && data[pos] != ':' {
			pos++
		}
		if pos >= len(data) {
			break
		}
		pos++ // skip ':'
		for pos < len(data) &&
			(data[pos] == ' ' || data[pos] == '\t') {
			pos++
		}
		if pos >= len(data) || data[pos] != '"' {
			searchFrom = pos
			continue
		}
		valStart := pos + 1 // after opening quote
		valEnd := valStart
		for valEnd < len(data) && data[valEnd] != '"' {
			valEnd++
		}
		if valEnd >= len(data) {
			break
		}

		val := data[valStart:valEnd]
		hexVal := tryBase64ToHex(val)
		if hexVal != nil {
			var result []byte
			result = append(
				result, data[:valStart]...,
			)
			result = append(result, hexVal...)
			result = append(
				result, data[valEnd:]...,
			)
			searchFrom = valStart + len(hexVal) + 1
			data = result
		} else {
			searchFrom = valEnd + 1
		}
	}
	return data
}

// tryBase64ToHex decodes a base64 string and re-encodes
// as lowercase hex. Returns nil if the input is not valid
// base64 or is already hex-sized (meaning it's likely
// already hex).
func tryBase64ToHex(val []byte) []byte {
	if len(val) == 0 {
		return nil
	}
	// If the value length matches hex encoding of 16 or
	// 8 bytes, it's likely already hex — leave it alone.
	if len(val) == 32 || len(val) == 16 {
		return nil
	}
	decoded, err := base64.StdEncoding.DecodeString(
		string(val),
	)
	if err != nil {
		return nil
	}
	if len(decoded) != 16 && len(decoded) != 8 {
		return nil
	}
	hexStr := make(
		[]byte, hex.EncodedLen(len(decoded)),
	)
	hex.Encode(hexStr, decoded)
	return hexStr
}
