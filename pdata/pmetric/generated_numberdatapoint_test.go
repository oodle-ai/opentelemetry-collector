// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlpmetrics "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/metrics/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

func TestNumberDataPoint_MoveTo(t *testing.T) {
	ms := generateTestNumberDataPoint()
	dest := NewNumberDataPoint()
	ms.MoveTo(dest)
	assert.Equal(t, NewNumberDataPoint(), ms)
	assert.Equal(t, generateTestNumberDataPoint(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newNumberDataPoint(&otlpmetrics.NumberDataPoint{}, &sharedState)) })
	assert.Panics(t, func() { newNumberDataPoint(&otlpmetrics.NumberDataPoint{}, &sharedState).MoveTo(dest) })
}

func TestNumberDataPoint_CopyTo(t *testing.T) {
	ms := NewNumberDataPoint()
	orig := NewNumberDataPoint()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestNumberDataPoint()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newNumberDataPoint(&otlpmetrics.NumberDataPoint{}, &sharedState)) })
}

func TestNumberDataPoint_Attributes(t *testing.T) {
	ms := NewNumberDataPoint()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestNumberDataPoint_StartTimestamp(t *testing.T) {
	ms := NewNumberDataPoint()
	assert.Equal(t, pcommon.Timestamp(0), ms.StartTimestamp())
	testValStartTimestamp := pcommon.Timestamp(1234567890)
	ms.SetStartTimestamp(testValStartTimestamp)
	assert.Equal(t, testValStartTimestamp, ms.StartTimestamp())
}

func TestNumberDataPoint_Timestamp(t *testing.T) {
	ms := NewNumberDataPoint()
	assert.Equal(t, pcommon.Timestamp(0), ms.Timestamp())
	testValTimestamp := pcommon.Timestamp(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.Equal(t, testValTimestamp, ms.Timestamp())
}

func TestNumberDataPoint_ValueType(t *testing.T) {
	tv := NewNumberDataPoint()
	assert.Equal(t, NumberDataPointValueTypeEmpty, tv.ValueType())
}

func TestNumberDataPoint_DoubleValue(t *testing.T) {
	ms := NewNumberDataPoint()
	assert.Equal(t, float64(0.0), ms.DoubleValue())
	ms.SetDoubleValue(float64(17.13))
	assert.Equal(t, float64(17.13), ms.DoubleValue())
	assert.Equal(t, NumberDataPointValueTypeDouble, ms.ValueType())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		newNumberDataPoint(&otlpmetrics.NumberDataPoint{}, &sharedState).SetDoubleValue(float64(17.13))
	})
}

func TestNumberDataPoint_IntValue(t *testing.T) {
	ms := NewNumberDataPoint()
	assert.Equal(t, int64(0), ms.IntValue())
	ms.SetIntValue(int64(17))
	assert.Equal(t, int64(17), ms.IntValue())
	assert.Equal(t, NumberDataPointValueTypeInt, ms.ValueType())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newNumberDataPoint(&otlpmetrics.NumberDataPoint{}, &sharedState).SetIntValue(int64(17)) })
}

func TestNumberDataPoint_Exemplars(t *testing.T) {
	ms := NewNumberDataPoint()
	assert.Equal(t, NewExemplarSlice(), ms.Exemplars())
	fillTestExemplarSlice(ms.Exemplars())
	assert.Equal(t, generateTestExemplarSlice(), ms.Exemplars())
}

func TestNumberDataPoint_Flags(t *testing.T) {
	ms := NewNumberDataPoint()
	assert.Equal(t, DataPointFlags(0), ms.Flags())
	testValFlags := DataPointFlags(1)
	ms.SetFlags(testValFlags)
	assert.Equal(t, testValFlags, ms.Flags())
}

func generateTestNumberDataPoint() NumberDataPoint {
	tv := NewNumberDataPoint()
	fillTestNumberDataPoint(tv)
	return tv
}

func fillTestNumberDataPoint(tv NumberDataPoint) {
	internal.FillTestMap(internal.NewMap(&tv.orig.Attributes, tv.state))
	tv.orig.StartTimeUnixNano = 1234567890
	tv.orig.TimeUnixNano = 1234567890
	tv.orig.Value = &otlpmetrics.NumberDataPoint_AsDouble{AsDouble: float64(17.13)}
	fillTestExemplarSlice(newExemplarSlice(&tv.orig.Exemplars, tv.state))
	tv.orig.Flags = 1
}
