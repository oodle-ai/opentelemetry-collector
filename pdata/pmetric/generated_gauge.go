// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlpmetrics "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/metrics/v1"
)

// Gauge represents the type of a numeric metric that always exports the "current value" for every data point.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewGauge function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Gauge struct {
	orig  *otlpmetrics.Gauge
	state *internal.State
}

func newGauge(orig *otlpmetrics.Gauge, state *internal.State) Gauge {
	return Gauge{orig: orig, state: state}
}

// NewGauge creates a new empty Gauge.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewGauge() Gauge {
	state := internal.StateMutable
	return newGauge(&otlpmetrics.Gauge{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms Gauge) MoveTo(dest Gauge) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.Gauge{}
}

// IsNil returns whether the struct is nil value.
func (ms Gauge) IsNil() bool {
	return ms.orig == nil
}

// DataPoints returns the DataPoints associated with this Gauge.
func (ms Gauge) DataPoints() NumberDataPointSlice {
	return newNumberDataPointSlice(&ms.orig.DataPoints, ms.state)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms Gauge) CopyTo(dest Gauge) {
	dest.state.AssertMutable()
	ms.DataPoints().CopyTo(dest.DataPoints())
}
