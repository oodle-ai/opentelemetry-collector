// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package plog

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlplogs "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/logs/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

// LogRecord are experimental implementation of OpenTelemetry Log Data Model.

// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogRecord function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecord struct {
	orig  *otlplogs.LogRecord
	state *internal.State
}

func newLogRecord(orig *otlplogs.LogRecord, state *internal.State) LogRecord {
	return LogRecord{orig: orig, state: state}
}

// NewLogRecord creates a new empty LogRecord.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewLogRecord() LogRecord {
	state := internal.StateMutable
	return newLogRecord(&otlplogs.LogRecord{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms LogRecord) MoveTo(dest LogRecord) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlplogs.LogRecord{}
}

// IsNil returns whether the struct is nil value.
func (ms LogRecord) IsNil() bool {
	return ms.orig == nil
}

// ObservedTimestamp returns the observedtimestamp associated with this LogRecord.
func (ms LogRecord) ObservedTimestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.orig.ObservedTimeUnixNano)
}

// SetObservedTimestamp replaces the observedtimestamp associated with this LogRecord.
func (ms LogRecord) SetObservedTimestamp(v pcommon.Timestamp) {
	ms.state.AssertMutable()
	ms.orig.ObservedTimeUnixNano = uint64(v)
}

// Timestamp returns the timestamp associated with this LogRecord.
func (ms LogRecord) Timestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.orig.TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this LogRecord.
func (ms LogRecord) SetTimestamp(v pcommon.Timestamp) {
	ms.state.AssertMutable()
	ms.orig.TimeUnixNano = uint64(v)
}

// TraceID returns the traceid associated with this LogRecord.
func (ms LogRecord) TraceID() pcommon.TraceID {
	return pcommon.TraceID(ms.orig.TraceId)
}

// SetTraceID replaces the traceid associated with this LogRecord.
func (ms LogRecord) SetTraceID(v pcommon.TraceID) {
	ms.state.AssertMutable()
	ms.orig.TraceId = v[:]
}

// SpanID returns the spanid associated with this LogRecord.
func (ms LogRecord) SpanID() pcommon.SpanID {
	return pcommon.SpanID(ms.orig.SpanId)
}

// SetSpanID replaces the spanid associated with this LogRecord.
func (ms LogRecord) SetSpanID(v pcommon.SpanID) {
	ms.state.AssertMutable()
	ms.orig.SpanId = v[:]
}

// Flags returns the flags associated with this LogRecord.
func (ms LogRecord) Flags() LogRecordFlags {
	return LogRecordFlags(ms.orig.Flags)
}

// SetFlags replaces the flags associated with this LogRecord.
func (ms LogRecord) SetFlags(v LogRecordFlags) {
	ms.state.AssertMutable()
	ms.orig.Flags = uint32(v)
}

// SeverityText returns the severitytext associated with this LogRecord.
func (ms LogRecord) SeverityText() string {
	return ms.orig.SeverityText
}

// SetSeverityText replaces the severitytext associated with this LogRecord.
func (ms LogRecord) SetSeverityText(v string) {
	ms.state.AssertMutable()
	ms.orig.SeverityText = v
}

// SeverityNumber returns the severitynumber associated with this LogRecord.
func (ms LogRecord) SeverityNumber() SeverityNumber {
	return SeverityNumber(ms.orig.SeverityNumber)
}

// SetSeverityNumber replaces the severitynumber associated with this LogRecord.
func (ms LogRecord) SetSeverityNumber(v SeverityNumber) {
	ms.state.AssertMutable()
	ms.orig.SeverityNumber = otlplogs.SeverityNumber(v)
}

// Body returns the body associated with this LogRecord.
func (ms LogRecord) Body() pcommon.Value {
	return pcommon.Value(internal.NewValue(ms.orig.Body, ms.state))
}

// Attributes returns the Attributes associated with this LogRecord.
func (ms LogRecord) Attributes() pcommon.Map {
	return pcommon.Map(internal.NewMap(&ms.orig.Attributes, ms.state))
}

// DroppedAttributesCount returns the droppedattributescount associated with this LogRecord.
func (ms LogRecord) DroppedAttributesCount() uint32 {
	return ms.orig.DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this LogRecord.
func (ms LogRecord) SetDroppedAttributesCount(v uint32) {
	ms.state.AssertMutable()
	ms.orig.DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms LogRecord) CopyTo(dest LogRecord) {
	dest.state.AssertMutable()
	dest.SetObservedTimestamp(ms.ObservedTimestamp())
	dest.SetTimestamp(ms.Timestamp())
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	dest.SetFlags(ms.Flags())
	dest.SetSeverityText(ms.SeverityText())
	dest.SetSeverityNumber(ms.SeverityNumber())
	ms.Body().CopyTo(dest.Body())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}
