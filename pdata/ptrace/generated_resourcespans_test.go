// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlptrace "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

func TestResourceSpans_MoveTo(t *testing.T) {
	ms := generateTestResourceSpans()
	dest := NewResourceSpans()
	ms.MoveTo(dest)
	assert.Equal(t, NewResourceSpans(), ms)
	assert.Equal(t, generateTestResourceSpans(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newResourceSpans(&otlptrace.ResourceSpans{}, &sharedState)) })
	assert.Panics(t, func() { newResourceSpans(&otlptrace.ResourceSpans{}, &sharedState).MoveTo(dest) })
}

func TestResourceSpans_CopyTo(t *testing.T) {
	ms := NewResourceSpans()
	orig := NewResourceSpans()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestResourceSpans()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newResourceSpans(&otlptrace.ResourceSpans{}, &sharedState)) })
}

func TestResourceSpans_Resource(t *testing.T) {
	ms := NewResourceSpans()
	internal.FillTestResource(internal.Resource(ms.Resource()))
	assert.Equal(t, pcommon.Resource(internal.GenerateTestResource()), ms.Resource())
}

func TestResourceSpans_SchemaUrl(t *testing.T) {
	ms := NewResourceSpans()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		newResourceSpans(&otlptrace.ResourceSpans{}, &sharedState).SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	})
}

func TestResourceSpans_ScopeSpans(t *testing.T) {
	ms := NewResourceSpans()
	assert.Equal(t, NewScopeSpansSlice(), ms.ScopeSpans())
	fillTestScopeSpansSlice(ms.ScopeSpans())
	assert.Equal(t, generateTestScopeSpansSlice(), ms.ScopeSpans())
}

func generateTestResourceSpans() ResourceSpans {
	tv := NewResourceSpans()
	fillTestResourceSpans(tv)
	return tv
}

func fillTestResourceSpans(tv ResourceSpans) {
	internal.FillTestResource(internal.NewResource(tv.orig.Resource, tv.state))
	tv.orig.SchemaUrl = "https://opentelemetry.io/schemas/1.5.0"
	fillTestScopeSpansSlice(newScopeSpansSlice(&tv.orig.ScopeSpans, tv.state))
}
