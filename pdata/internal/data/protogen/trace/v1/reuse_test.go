// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

// A pooled Span keeps its ID slices after ResetVT. If a zero-copy
// decode left those slices pointing into the caller's buffer, a later
// copying decode must not write into that buffer.
func TestSpanCopyingDecodeAfterUnsafeDecode(t *testing.T) {
	first := &Span{
		TraceId:      bytes.Repeat([]byte{0x11}, 16),
		SpanId:       bytes.Repeat([]byte{0x12}, 8),
		ParentSpanId: bytes.Repeat([]byte{0x13}, 8),
		Links: []*Span_Link{{
			TraceId: bytes.Repeat([]byte{0x14}, 16),
			SpanId:  bytes.Repeat([]byte{0x15}, 8),
		}},
	}
	second := &Span{
		TraceId:      bytes.Repeat([]byte{0x21}, 16),
		SpanId:       bytes.Repeat([]byte{0x22}, 8),
		ParentSpanId: bytes.Repeat([]byte{0x23}, 8),
		Links: []*Span_Link{{
			TraceId: bytes.Repeat([]byte{0x24}, 16),
			SpanId:  bytes.Repeat([]byte{0x25}, 8),
		}},
	}
	firstBuf, err := first.MarshalVT()
	require.NoError(t, err)
	secondBuf, err := second.MarshalVT()
	require.NoError(t, err)

	span := &Span{}
	require.NoError(t, span.UnmarshalVTUnsafe(firstBuf))
	span.ResetVT()
	require.NoError(t, span.UnmarshalVT(secondBuf))

	// The caller owns firstBuf again and can reuse it.
	for i := range firstBuf {
		firstBuf[i] = 0xff
	}
	require.Equal(t, second.String(), span.String())
}
