// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/oodle-ai/opentelemetry-collector/pdata/internal/cmd/pdatagen/internal"

import (
	"bytes"
	"strings"
	"text/template"
)

const messageValueTemplate = `{{ .description }}
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use New{{ .structName }} function to create new instances.
// Important: zero-initialized instance is not valid for use.
{{- if .isCommon }}
type {{ .structName }} internal.{{ .structName }}
{{- else }}
type {{ .structName }} struct {
	orig *{{ .originName }}
	state *internal.State
}
{{- end }}

func new{{ .structName }}(orig *{{ .originName }}, state *internal.State) {{ .structName }} {
	{{- if .isCommon }}
	return {{ .structName }}(internal.New{{ .structName }}(orig, state))
	{{- else }}
	return {{ .structName }}{orig: orig, state: state}
	{{- end }}
}

// New{{ .structName }} creates a new empty {{ .structName }}.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func New{{ .structName }}() {{ .structName }} {
	state := internal.StateMutable
	return new{{ .structName }}(&{{ .originName }}{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms {{ .structName }}) MoveTo(dest {{ .structName }}) {
	ms.{{- if .isCommon }}getState(){{ else }}state{{ end }}.AssertMutable()
	dest.{{- if .isCommon }}getState(){{ else }}state{{ end }}.AssertMutable()
	*dest.{{ .origAccessor }} = *ms.{{ .origAccessor }}
	*ms.{{ .origAccessor }} = {{ .originName }}{}
}

{{ if .isCommon -}}
func (ms {{ .structName }}) getOrig() *{{ .originName }} {
	return internal.GetOrig{{ .structName }}(internal.{{ .structName }}(ms))
}

func (ms {{ .structName }}) getState() *internal.State {
	return internal.Get{{ .structName }}State(internal.{{ .structName }}(ms))
}
{{- end }}

{{ range .fields -}}
{{ .GenerateAccessors $.messageStruct }}
{{ end }}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms {{ .structName }}) CopyTo(dest {{ .structName }}) {
	dest.{{- if .isCommon }}getState(){{ else }}state{{ end }}.AssertMutable()
	{{- range .fields }}
	{{ .GenerateCopyToValue $.messageStruct }}
	{{- end }}
}`

const messageValueTestTemplate = `
func Test{{ .structName }}_MoveTo(t *testing.T) {
	ms := {{ .generateTestData }}
	dest := New{{ .structName }}()
	ms.MoveTo(dest)
	assert.Equal(t, New{{ .structName }}(), ms)
	assert.Equal(t, {{ .generateTestData }}, dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(new{{ .structName }}(&{{ .originName }}{}, &sharedState)) })
	assert.Panics(t, func() { new{{ .structName }}(&{{ .originName }}{}, &sharedState).MoveTo(dest) })
}

func Test{{ .structName }}_CopyTo(t *testing.T) {
	ms := New{{ .structName }}()
	orig := New{{ .structName }}()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = {{ .generateTestData }}
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(new{{ .structName }}(&{{ .originName }}{}, &sharedState)) })
}

{{ range .fields }}
{{ .GenerateAccessorsTest $.messageStruct }}
{{ end }}`

const messageValueGenerateTestTemplate = `func {{ upperIfInternal "g" }}enerateTest{{ .structName }}() {{ .structName }} {
	{{- if .isCommon }}
	orig := {{ .originName }}{}
	state := StateMutable
	{{- end }}
	tv := New{{ .structName }}({{ if .isCommon }}&orig, &state{{ end }})
	{{ upperIfInternal "f" }}illTest{{ .structName }}(tv)
	return tv
}

func {{ upperIfInternal "f" }}illTest{{ .structName }}(tv {{ .structName }}) {
	{{- range .fields }}
	{{ .GenerateSetWithTestValue $.messageStruct }}
	{{- end }}
}`

const messageValueAliasTemplate = `
type {{ .structName }} struct {
	orig *{{ .originName }}
	state *State
}

func GetOrig{{ .structName }}(ms {{ .structName }}) *{{ .originName }} {
	return ms.orig
}

func Get{{ .structName }}State(ms {{ .structName }}) *State {
	return ms.state
}

func New{{ .structName }}(orig *{{ .originName }}, state *State) {{ .structName }} {
	return {{ .structName }}{orig: orig, state: state}
}`

type baseStruct interface {
	getName() string
	generateStruct(sb *bytes.Buffer)
	generateTests(sb *bytes.Buffer)
	generateTestValueHelpers(sb *bytes.Buffer)
	generateInternal(sb *bytes.Buffer)
}

// messageValueStruct generates a struct for a proto message. The struct can be generated both as a common struct
// that can be used as a field in struct from other packages and as an isolated struct with depending on a package name.
type messageValueStruct struct {
	structName     string
	packageName    string
	description    string
	originFullName string
	fields         []baseField
}

func (ms *messageValueStruct) getName() string {
	return ms.structName
}

func (ms *messageValueStruct) generateStruct(sb *bytes.Buffer) {
	t := template.Must(template.New("messageValueTemplate").Parse(messageValueTemplate))
	if err := t.Execute(sb, ms.templateFields()); err != nil {
		panic(err)
	}
}

func (ms *messageValueStruct) generateTests(sb *bytes.Buffer) {
	t := template.Must(template.New("messageValueTestTemplate").Parse(messageValueTestTemplate))
	if err := t.Execute(sb, ms.templateFields()); err != nil {
		panic(err)
	}
}

func (ms *messageValueStruct) generateTestValueHelpers(sb *bytes.Buffer) {
	funcs := template.FuncMap{
		"upperIfInternal": func(in string) string {
			if usedByOtherDataTypes(ms.packageName) {
				return strings.ToUpper(in)
			}
			return in
		},
	}
	t := template.Must(template.New("messageValueGenerateTestTemplate").Funcs(funcs).Parse(messageValueGenerateTestTemplate))
	if err := t.Execute(sb, ms.templateFields()); err != nil {
		panic(err)
	}
}

func (ms *messageValueStruct) generateInternal(sb *bytes.Buffer) {
	t := template.Must(template.New("messageValueAliasTemplate").Parse(messageValueAliasTemplate))
	if err := t.Execute(sb, ms.templateFields()); err != nil {
		panic(err)
	}
}

func (ms *messageValueStruct) templateFields() map[string]any {
	return map[string]any{
		"messageStruct": ms,
		"fields":        ms.fields,
		"structName":    ms.structName,
		"originName":    ms.originFullName,
		"generateTestData": func() string {
			if usedByOtherDataTypes(ms.packageName) {
				return ms.structName + "(internal.GenerateTest" + ms.structName + "())"
			}
			return "generateTest" + ms.structName + "()"
		}(),
		"description":  ms.description,
		"isCommon":     usedByOtherDataTypes(ms.packageName),
		"origAccessor": origAccessor(ms),
	}
}

var _ baseStruct = (*messageValueStruct)(nil)
