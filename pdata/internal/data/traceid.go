// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package data // import "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data"

import (
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

const traceIDSize = 16

var (
	errMarshalTraceID = errors.New(
		"marshal: invalid buffer length for TraceID",
	)
	errUnmarshalTraceID = errors.New(
		"unmarshal: invalid TraceID length",
	)
)

// TraceID is a custom data type that is used for all
// trace_id fields in OTLP Protobuf messages.
type TraceID []byte

// Size returns the size of the data to serialize.
func (tid TraceID) Size() int {
	if tid.IsEmpty() {
		return 0
	}
	return traceIDSize
}

// IsEmpty returns true if id is empty or all zero bytes.
func (tid TraceID) IsEmpty() bool {
	if len(tid) == 0 {
		return true
	}
	for _, b := range tid {
		if b != 0 {
			return false
		}
	}
	return true
}

// MarshalTo converts trace ID into a binary
// representation. Called by Protobuf serialization.
func (tid TraceID) MarshalTo(data []byte) (int, error) {
	if tid.IsEmpty() {
		return 0, nil
	}
	if len(data) < traceIDSize {
		return 0, errMarshalTraceID
	}
	return copy(data, tid), nil
}

// Unmarshal inflates this trace ID from binary
// representation. Called by Protobuf serialization.
func (tid *TraceID) Unmarshal(data []byte) error {
	if len(data) == 0 {
		*tid = make([]byte, traceIDSize)
		return nil
	}
	if len(data) != traceIDSize {
		return errUnmarshalTraceID
	}
	*tid = make([]byte, traceIDSize)
	copy(*tid, data)
	return nil
}

// MarshalJSON converts trace id into a hex string
// enclosed in quotes.
func (tid TraceID) MarshalJSON() ([]byte, error) {
	if tid.IsEmpty() {
		return []byte(`""`), nil
	}
	return marshalJSON(tid)
}

// UnmarshalJSON inflates trace id from hex string,
// possibly enclosed in quotes.
func (tid *TraceID) UnmarshalJSON(data []byte) error {
	return unmarshalOTLPJSONBytesID(
		(*[]byte)(tid), traceIDSize, data,
	)
}

// UnmarshalJsoniter reads a hex-encoded trace ID from
// iter and reports any error using the given operation
// and format string.
func (tid *TraceID) UnmarshalJsoniter(
	iter *jsoniter.Iterator,
	op string,
	errFmt string,
) {
	if err := tid.UnmarshalJSON([]byte(iter.ReadString())); err != nil {
		iter.ReportError(op, fmt.Sprintf(errFmt, err))
	}
}
