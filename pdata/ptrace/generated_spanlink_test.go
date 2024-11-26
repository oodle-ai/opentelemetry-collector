// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal/data"
	otlptrace "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

func TestSpanLink_MoveTo(t *testing.T) {
	ms := generateTestSpanLink()
	dest := NewSpanLink()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpanLink(), ms)
	assert.Equal(t, generateTestSpanLink(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newSpanLink(&otlptrace.Span_Link{}, &sharedState)) })
	assert.Panics(t, func() { newSpanLink(&otlptrace.Span_Link{}, &sharedState).MoveTo(dest) })
}

func TestSpanLink_CopyTo(t *testing.T) {
	ms := NewSpanLink()
	orig := NewSpanLink()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSpanLink()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newSpanLink(&otlptrace.Span_Link{}, &sharedState)) })
}

func TestSpanLink_TraceID(t *testing.T) {
	ms := NewSpanLink()
	tid, _ := ms.TraceID()
	assert.Equal(t, pcommon.TraceID(data.TraceID([16]byte{})), tid)
	testValTraceID := pcommon.TraceID(data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, tid)
}

func TestSpanLink_SpanID(t *testing.T) {
	ms := NewSpanLink()
	sp, _ := ms.SpanID()
	assert.Equal(t, pcommon.SpanID(data.SpanID([8]byte{})), sp)
	testValSpanID := pcommon.SpanID(data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetSpanID(testValSpanID)
	sp, _ = ms.SpanID()
	assert.Equal(t, testValSpanID, sp)
}

func TestSpanLink_TraceState(t *testing.T) {
	ms := NewSpanLink()
	internal.FillTestTraceState(internal.TraceState(ms.TraceState()))
	assert.Equal(t, pcommon.TraceState(internal.GenerateTestTraceState()), ms.TraceState())
}

func TestSpanLink_Flags(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, uint32(0), ms.Flags())
	ms.SetFlags(uint32(0xf))
	assert.Equal(t, uint32(0xf), ms.Flags())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSpanLink(&otlptrace.Span_Link{}, &sharedState).SetFlags(uint32(0xf)) })
}

func TestSpanLink_Attributes(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestSpanLink_DroppedAttributesCount(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSpanLink(&otlptrace.Span_Link{}, &sharedState).SetDroppedAttributesCount(uint32(17)) })
}

func generateTestSpanLink() SpanLink {
	tv := NewSpanLink()
	fillTestSpanLink(tv)
	return tv
}

func fillTestSpanLink(tv SpanLink) {
	tv.orig.TraceId = []byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}
	tv.orig.SpanId = []byte{8, 7, 6, 5, 4, 3, 2, 1}
	internal.FillTestTraceState(internal.NewTraceState(&tv.orig.TraceState, tv.state))
	tv.orig.Flags = uint32(0xf)
	internal.FillTestMap(internal.NewMap(&tv.orig.Attributes, tv.state))
	tv.orig.DroppedAttributesCount = uint32(17)
}
