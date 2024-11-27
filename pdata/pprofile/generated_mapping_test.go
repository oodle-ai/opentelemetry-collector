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
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

func TestMapping_MoveTo(t *testing.T) {
	ms := generateTestMapping()
	dest := NewMapping()
	ms.MoveTo(dest)
	assert.Equal(t, NewMapping(), ms)
	assert.Equal(t, generateTestMapping(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newMapping(&otlpprofiles.Mapping{}, &sharedState)) })
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).MoveTo(dest) })
}

func TestMapping_CopyTo(t *testing.T) {
	ms := NewMapping()
	orig := NewMapping()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestMapping()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newMapping(&otlpprofiles.Mapping{}, &sharedState)) })
}

func TestMapping_ID(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, uint64(0), ms.ID())
	ms.SetID(uint64(1))
	assert.Equal(t, uint64(1), ms.ID())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetID(uint64(1)) })
}

func TestMapping_MemoryStart(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, uint64(0), ms.MemoryStart())
	ms.SetMemoryStart(uint64(1))
	assert.Equal(t, uint64(1), ms.MemoryStart())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetMemoryStart(uint64(1)) })
}

func TestMapping_MemoryLimit(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, uint64(0), ms.MemoryLimit())
	ms.SetMemoryLimit(uint64(1))
	assert.Equal(t, uint64(1), ms.MemoryLimit())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetMemoryLimit(uint64(1)) })
}

func TestMapping_FileOffset(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, uint64(0), ms.FileOffset())
	ms.SetFileOffset(uint64(1))
	assert.Equal(t, uint64(1), ms.FileOffset())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetFileOffset(uint64(1)) })
}

func TestMapping_Filename(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, int64(0), ms.Filename())
	ms.SetFilename(int64(1))
	assert.Equal(t, int64(1), ms.Filename())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetFilename(int64(1)) })
}

func TestMapping_BuildID(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, int64(0), ms.BuildID())
	ms.SetBuildID(int64(1))
	assert.Equal(t, int64(1), ms.BuildID())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetBuildID(int64(1)) })
}

func TestMapping_BuildIDKind(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, otlpprofiles.BuildIdKind(0), ms.BuildIDKind())
	ms.SetBuildIDKind(otlpprofiles.BuildIdKind(1))
	assert.Equal(t, otlpprofiles.BuildIdKind(1), ms.BuildIDKind())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetBuildIDKind(otlpprofiles.BuildIdKind(1)) })
}

func TestMapping_Attributes(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, pcommon.NewUInt64Slice(), ms.Attributes())
	internal.FillTestUInt64Slice(internal.UInt64Slice(ms.Attributes()))
	assert.Equal(t, pcommon.UInt64Slice(internal.GenerateTestUInt64Slice()), ms.Attributes())
}

func TestMapping_HasFunctions(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, false, ms.HasFunctions())
	ms.SetHasFunctions(true)
	assert.Equal(t, true, ms.HasFunctions())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetHasFunctions(true) })
}

func TestMapping_HasFilenames(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, false, ms.HasFilenames())
	ms.SetHasFilenames(true)
	assert.Equal(t, true, ms.HasFilenames())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetHasFilenames(true) })
}

func TestMapping_HasLineNumbers(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, false, ms.HasLineNumbers())
	ms.SetHasLineNumbers(true)
	assert.Equal(t, true, ms.HasLineNumbers())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetHasLineNumbers(true) })
}

func TestMapping_HasInlineFrames(t *testing.T) {
	ms := NewMapping()
	assert.Equal(t, false, ms.HasInlineFrames())
	ms.SetHasInlineFrames(true)
	assert.Equal(t, true, ms.HasInlineFrames())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMapping(&otlpprofiles.Mapping{}, &sharedState).SetHasInlineFrames(true) })
}

func generateTestMapping() Mapping {
	tv := NewMapping()
	fillTestMapping(tv)
	return tv
}

func fillTestMapping(tv Mapping) {
	tv.orig.Id = uint64(1)
	tv.orig.MemoryStart = uint64(1)
	tv.orig.MemoryLimit = uint64(1)
	tv.orig.FileOffset = uint64(1)
	tv.orig.Filename = int64(1)
	tv.orig.BuildId = int64(1)
	tv.orig.BuildIdKind = otlpprofiles.BuildIdKind(1)
	internal.FillTestUInt64Slice(internal.NewUInt64Slice(&tv.orig.Attributes, tv.state))
	tv.orig.HasFunctions = true
	tv.orig.HasFilenames = true
	tv.orig.HasLineNumbers = true
	tv.orig.HasInlineFrames = true
}