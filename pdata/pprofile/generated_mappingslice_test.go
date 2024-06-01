// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlpprofiles "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/profiles/v1experimental"
)

func TestMappingSlice(t *testing.T) {
	es := NewMappingSlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newMappingSlice(&[]otlpprofiles.Mapping{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewMapping()
	testVal := generateTestMapping()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestMapping(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestMappingSliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newMappingSlice(&[]otlpprofiles.Mapping{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewMappingSlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestMappingSlice_CopyTo(t *testing.T) {
	dest := NewMappingSlice()
	// Test CopyTo to empty
	NewMappingSlice().CopyTo(dest)
	assert.Equal(t, NewMappingSlice(), dest)

	// Test CopyTo larger slice
	generateTestMappingSlice().CopyTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)

	// Test CopyTo same size slice
	generateTestMappingSlice().CopyTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)
}

func TestMappingSlice_EnsureCapacity(t *testing.T) {
	es := generateTestMappingSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestMappingSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestMappingSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestMappingSlice(), es)
}

func TestMappingSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestMappingSlice()
	dest := NewMappingSlice()
	src := generateTestMappingSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestMappingSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestMappingSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewMappingSlice()
	emptySlice.RemoveIf(func(el Mapping) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestMappingSlice()
	pos := 0
	filtered.RemoveIf(func(el Mapping) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func generateTestMappingSlice() MappingSlice {
	es := NewMappingSlice()
	fillTestMappingSlice(es)
	return es
}

func fillTestMappingSlice(es MappingSlice) {
	*es.orig = make([]otlpprofiles.Mapping, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = otlpprofiles.Mapping{}
		fillTestMapping(newMapping(&(*es.orig)[i], es.state))
	}
}
