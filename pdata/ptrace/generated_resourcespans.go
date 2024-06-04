// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlptrace "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

// ResourceSpans is a collection of spans from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpans struct {
	orig  *otlptrace.ResourceSpans
	state *internal.State
}

func newResourceSpans(orig *otlptrace.ResourceSpans, state *internal.State) ResourceSpans {
	return ResourceSpans{orig: orig, state: state}
}

// NewResourceSpans creates a new empty ResourceSpans.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewResourceSpans() ResourceSpans {
	state := internal.StateMutable
	return newResourceSpans(&otlptrace.ResourceSpans{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ResourceSpans) MoveTo(dest ResourceSpans) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlptrace.ResourceSpans{}
}

// Resource returns the resource associated with this ResourceSpans.
func (ms ResourceSpans) Resource() pcommon.Resource {
	return pcommon.Resource(internal.NewResource(&ms.orig.Resource, ms.state))
}

// SchemaUrl returns the schemaurl associated with this ResourceSpans.
func (ms ResourceSpans) SchemaUrl() string {
	return ms.orig.SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ResourceSpans.
func (ms ResourceSpans) SetSchemaUrl(v string) {
	ms.state.AssertMutable()
	ms.orig.SchemaUrl = v
}

// ScopeSpans returns the ScopeSpans associated with this ResourceSpans.
func (ms ResourceSpans) ScopeSpans() ScopeSpansSlice {
	return newScopeSpansSlice(&ms.orig.ScopeSpans, ms.state)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ResourceSpans) CopyTo(dest ResourceSpans) {
	dest.state.AssertMutable()
	ms.Resource().CopyTo(dest.Resource())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.ScopeSpans().CopyTo(dest.ScopeSpans())
}
