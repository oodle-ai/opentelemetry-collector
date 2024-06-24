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

// ScopeLogs is a collection of logs from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewScopeLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeLogs struct {
	orig  *otlplogs.ScopeLogs
	state *internal.State
}

func newScopeLogs(orig *otlplogs.ScopeLogs, state *internal.State) ScopeLogs {
	return ScopeLogs{orig: orig, state: state}
}

// NewScopeLogs creates a new empty ScopeLogs.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewScopeLogs() ScopeLogs {
	state := internal.StateMutable
	return newScopeLogs(&otlplogs.ScopeLogs{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ScopeLogs) MoveTo(dest ScopeLogs) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlplogs.ScopeLogs{}
}

// IsNil returns whether the struct is nil value.
func (ms ScopeLogs) IsNil() bool {
	return ms.orig == nil
}

// Scope returns the scope associated with this ScopeLogs.
func (ms ScopeLogs) Scope() pcommon.InstrumentationScope {
	return pcommon.InstrumentationScope(internal.NewInstrumentationScope(&ms.orig.Scope, ms.state))
}

// SchemaUrl returns the schemaurl associated with this ScopeLogs.
func (ms ScopeLogs) SchemaUrl() string {
	return ms.orig.SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ScopeLogs.
func (ms ScopeLogs) SetSchemaUrl(v string) {
	ms.state.AssertMutable()
	ms.orig.SchemaUrl = v
}

// LogRecords returns the LogRecords associated with this ScopeLogs.
func (ms ScopeLogs) LogRecords() LogRecordSlice {
	return newLogRecordSlice(&ms.orig.LogRecords, ms.state)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ScopeLogs) CopyTo(dest ScopeLogs) {
	dest.state.AssertMutable()
	ms.Scope().CopyTo(dest.Scope())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.LogRecords().CopyTo(dest.LogRecords())
}
