// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlpprofiles "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/profiles/v1experimental"
)

// ValueType describes the type and units of a value, with an optional aggregation temporality.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewValueType function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ValueType struct {
	orig  *otlpprofiles.ValueType
	state *internal.State
}

func newValueType(orig *otlpprofiles.ValueType, state *internal.State) ValueType {
	return ValueType{orig: orig, state: state}
}

// NewValueType creates a new empty ValueType.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewValueType() ValueType {
	state := internal.StateMutable
	return newValueType(&otlpprofiles.ValueType{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ValueType) MoveTo(dest ValueType) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlpprofiles.ValueType{}
}

// Type returns the type associated with this ValueType.
func (ms ValueType) Type() int64 {
	return ms.orig.Type
}

// SetType replaces the type associated with this ValueType.
func (ms ValueType) SetType(v int64) {
	ms.state.AssertMutable()
	ms.orig.Type = v
}

// Unit returns the unit associated with this ValueType.
func (ms ValueType) Unit() int64 {
	return ms.orig.Unit
}

// SetUnit replaces the unit associated with this ValueType.
func (ms ValueType) SetUnit(v int64) {
	ms.state.AssertMutable()
	ms.orig.Unit = v
}

// AggregationTemporality returns the aggregationtemporality associated with this ValueType.
func (ms ValueType) AggregationTemporality() otlpprofiles.AggregationTemporality {
	return ms.orig.AggregationTemporality
}

// SetAggregationTemporality replaces the aggregationtemporality associated with this ValueType.
func (ms ValueType) SetAggregationTemporality(v otlpprofiles.AggregationTemporality) {
	ms.state.AssertMutable()
	ms.orig.AggregationTemporality = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ValueType) CopyTo(dest ValueType) {
	dest.state.AssertMutable()
	dest.SetType(ms.Type())
	dest.SetUnit(ms.Unit())
	dest.SetAggregationTemporality(ms.AggregationTemporality())
}
