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

func TestResourceProfiles_MoveTo(t *testing.T) {
	ms := generateTestResourceProfiles()
	dest := NewResourceProfiles()
	ms.MoveTo(dest)
	assert.Equal(t, NewResourceProfiles(), ms)
	assert.Equal(t, generateTestResourceProfiles(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newResourceProfiles(&otlpprofiles.ResourceProfiles{}, &sharedState)) })
	assert.Panics(t, func() { newResourceProfiles(&otlpprofiles.ResourceProfiles{}, &sharedState).MoveTo(dest) })
}

func TestResourceProfiles_CopyTo(t *testing.T) {
	ms := NewResourceProfiles()
	orig := NewResourceProfiles()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestResourceProfiles()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newResourceProfiles(&otlpprofiles.ResourceProfiles{}, &sharedState)) })
}

func TestResourceProfiles_Resource(t *testing.T) {
	ms := NewResourceProfiles()
	internal.FillTestResource(internal.Resource(ms.Resource()))
	assert.Equal(t, pcommon.Resource(internal.GenerateTestResource()), ms.Resource())
}

func TestResourceProfiles_SchemaUrl(t *testing.T) {
	ms := NewResourceProfiles()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		newResourceProfiles(&otlpprofiles.ResourceProfiles{}, &sharedState).SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	})
}

func TestResourceProfiles_ScopeProfiles(t *testing.T) {
	ms := NewResourceProfiles()
	assert.Equal(t, NewScopeProfilesSlice(), ms.ScopeProfiles())
	fillTestScopeProfilesSlice(ms.ScopeProfiles())
	assert.Equal(t, generateTestScopeProfilesSlice(), ms.ScopeProfiles())
}

func generateTestResourceProfiles() ResourceProfiles {
	tv := NewResourceProfiles()
	fillTestResourceProfiles(tv)
	return tv
}

func fillTestResourceProfiles(tv ResourceProfiles) {
	internal.FillTestResource(internal.NewResource(&tv.orig.Resource, tv.state))
	tv.orig.SchemaUrl = "https://opentelemetry.io/schemas/1.5.0"
	fillTestScopeProfilesSlice(newScopeProfilesSlice(&tv.orig.ScopeProfiles, tv.state))
}
