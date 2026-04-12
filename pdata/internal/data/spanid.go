// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package data // import "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data"

import (
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

const spanIDSize = 8

var (
	errMarshalSpanID = errors.New(
		"marshal: invalid buffer length for SpanID",
	)
	errUnmarshalSpanID = errors.New(
		"unmarshal: invalid SpanID length",
	)
)

// SpanID is a custom data type that is used for all
// span_id fields in OTLP Protobuf messages.
type SpanID []byte

// Size returns the size of the data to serialize.
func (sid SpanID) Size() int {
	if sid.IsEmpty() {
		return 0
	}
	return spanIDSize
}

// IsEmpty returns true if id is empty or all zero bytes.
func (sid SpanID) IsEmpty() bool {
	if len(sid) == 0 {
		return true
	}
	for _, b := range sid {
		if b != 0 {
			return false
		}
	}
	return true
}

// MarshalTo converts span ID into a binary
// representation. Called by Protobuf serialization.
func (sid SpanID) MarshalTo(data []byte) (int, error) {
	if sid.IsEmpty() {
		return 0, nil
	}
	if len(data) < spanIDSize {
		return 0, errMarshalSpanID
	}
	return copy(data, sid), nil
}

// Unmarshal inflates this span ID from binary
// representation. Called by Protobuf serialization.
func (sid *SpanID) Unmarshal(data []byte) error {
	if len(data) == 0 {
		*sid = make([]byte, spanIDSize)
		return nil
	}
	if len(data) != spanIDSize {
		return errUnmarshalSpanID
	}
	*sid = make([]byte, spanIDSize)
	copy(*sid, data)
	return nil
}

// MarshalJSON converts SpanID into a hex string
// enclosed in quotes.
func (sid SpanID) MarshalJSON() ([]byte, error) {
	if sid.IsEmpty() {
		return []byte(`""`), nil
	}
	return marshalJSON(sid)
}

// UnmarshalJSON decodes SpanID from hex string,
// possibly enclosed in quotes.
func (sid *SpanID) UnmarshalJSON(data []byte) error {
	return unmarshalOTLPJSONBytesID(
		(*[]byte)(sid), spanIDSize, data,
	)
}

// UnmarshalJsoniter reads a hex-encoded span ID from
// iter and reports any error using the given operation
// and format string.
func (sid *SpanID) UnmarshalJsoniter(
	iter *jsoniter.Iterator,
	op string,
	errFmt string,
) {
	if err := sid.UnmarshalJSON([]byte(iter.ReadString())); err != nil {
		iter.ReportError(op, fmt.Sprintf(errFmt, err))
	}
}
