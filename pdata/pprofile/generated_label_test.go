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

func TestLabel_MoveTo(t *testing.T) {
	ms := generateTestLabel()
	dest := NewLabel()
	ms.MoveTo(dest)
	assert.Equal(t, NewLabel(), ms)
	assert.Equal(t, generateTestLabel(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newLabel(&otlpprofiles.Label{}, &sharedState)) })
	assert.Panics(t, func() { newLabel(&otlpprofiles.Label{}, &sharedState).MoveTo(dest) })
}

func TestLabel_CopyTo(t *testing.T) {
	ms := NewLabel()
	orig := NewLabel()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestLabel()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newLabel(&otlpprofiles.Label{}, &sharedState)) })
}

func TestLabel_Key(t *testing.T) {
	ms := NewLabel()
	assert.Equal(t, int64(0), ms.Key())
	ms.SetKey(int64(1))
	assert.Equal(t, int64(1), ms.Key())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newLabel(&otlpprofiles.Label{}, &sharedState).SetKey(int64(1)) })
}

func TestLabel_Str(t *testing.T) {
	ms := NewLabel()
	assert.Equal(t, int64(0), ms.Str())
	ms.SetStr(int64(1))
	assert.Equal(t, int64(1), ms.Str())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newLabel(&otlpprofiles.Label{}, &sharedState).SetStr(int64(1)) })
}

func TestLabel_Num(t *testing.T) {
	ms := NewLabel()
	assert.Equal(t, int64(0), ms.Num())
	ms.SetNum(int64(1))
	assert.Equal(t, int64(1), ms.Num())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newLabel(&otlpprofiles.Label{}, &sharedState).SetNum(int64(1)) })
}

func TestLabel_NumUnit(t *testing.T) {
	ms := NewLabel()
	assert.Equal(t, int64(0), ms.NumUnit())
	ms.SetNumUnit(int64(1))
	assert.Equal(t, int64(1), ms.NumUnit())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newLabel(&otlpprofiles.Label{}, &sharedState).SetNumUnit(int64(1)) })
}

func generateTestLabel() Label {
	tv := NewLabel()
	fillTestLabel(tv)
	return tv
}

func fillTestLabel(tv Label) {
	tv.orig.Key = int64(1)
	tv.orig.Str = int64(1)
	tv.orig.Num = int64(1)
	tv.orig.NumUnit = int64(1)
}
