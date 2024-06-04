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

func TestMetric_MoveTo(t *testing.T) {
	ms := generateTestMetric()
	dest := NewMetric()
	ms.MoveTo(dest)
	assert.Equal(t, NewMetric(), ms)
	assert.Equal(t, generateTestMetric(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newMetric(&otlpmetrics.Metric{}, &sharedState)) })
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).MoveTo(dest) })
}

func TestMetric_CopyTo(t *testing.T) {
	ms := NewMetric()
	orig := NewMetric()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestMetric()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newMetric(&otlpmetrics.Metric{}, &sharedState)) })
}

func TestMetric_Name(t *testing.T) {
	ms := NewMetric()
	assert.Equal(t, "", ms.Name())
	ms.SetName("test_name")
	assert.Equal(t, "test_name", ms.Name())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetName("test_name") })
}

func TestMetric_Description(t *testing.T) {
	ms := NewMetric()
	assert.Equal(t, "", ms.Description())
	ms.SetDescription("test_description")
	assert.Equal(t, "test_description", ms.Description())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetDescription("test_description") })
}

func TestMetric_Unit(t *testing.T) {
	ms := NewMetric()
	assert.Equal(t, "", ms.Unit())
	ms.SetUnit("1")
	assert.Equal(t, "1", ms.Unit())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetUnit("1") })
}

func TestMetric_Metadata(t *testing.T) {
	ms := NewMetric()
	assert.Equal(t, pcommon.NewMap(), ms.Metadata())
	internal.FillTestMap(internal.Map(ms.Metadata()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Metadata())
}

func TestMetric_Type(t *testing.T) {
	tv := NewMetric()
	assert.Equal(t, MetricTypeEmpty, tv.Type())
}

func TestMetric_Gauge(t *testing.T) {
	ms := NewMetric()
	fillTestGauge(ms.SetEmptyGauge())
	assert.Equal(t, MetricTypeGauge, ms.Type())
	assert.Equal(t, generateTestGauge(), ms.Gauge())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetEmptyGauge() })
}

func TestMetric_CopyTo_Gauge(t *testing.T) {
	ms := NewMetric()
	fillTestGauge(ms.SetEmptyGauge())
	dest := NewMetric()
	ms.CopyTo(dest)
	assert.Equal(t, ms, dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newMetric(&otlpmetrics.Metric{}, &sharedState)) })
}

func TestMetric_Sum(t *testing.T) {
	ms := NewMetric()
	fillTestSum(ms.SetEmptySum())
	assert.Equal(t, MetricTypeSum, ms.Type())
	assert.Equal(t, generateTestSum(), ms.Sum())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetEmptySum() })
}

func TestMetric_CopyTo_Sum(t *testing.T) {
	ms := NewMetric()
	fillTestSum(ms.SetEmptySum())
	dest := NewMetric()
	ms.CopyTo(dest)
	assert.Equal(t, ms, dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newMetric(&otlpmetrics.Metric{}, &sharedState)) })
}

func TestMetric_Histogram(t *testing.T) {
	ms := NewMetric()
	fillTestHistogram(ms.SetEmptyHistogram())
	assert.Equal(t, MetricTypeHistogram, ms.Type())
	assert.Equal(t, generateTestHistogram(), ms.Histogram())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetEmptyHistogram() })
}

func TestMetric_CopyTo_Histogram(t *testing.T) {
	ms := NewMetric()
	fillTestHistogram(ms.SetEmptyHistogram())
	dest := NewMetric()
	ms.CopyTo(dest)
	assert.Equal(t, ms, dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newMetric(&otlpmetrics.Metric{}, &sharedState)) })
}

func TestMetric_ExponentialHistogram(t *testing.T) {
	ms := NewMetric()
	fillTestExponentialHistogram(ms.SetEmptyExponentialHistogram())
	assert.Equal(t, MetricTypeExponentialHistogram, ms.Type())
	assert.Equal(t, generateTestExponentialHistogram(), ms.ExponentialHistogram())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetEmptyExponentialHistogram() })
}

func TestMetric_CopyTo_ExponentialHistogram(t *testing.T) {
	ms := NewMetric()
	fillTestExponentialHistogram(ms.SetEmptyExponentialHistogram())
	dest := NewMetric()
	ms.CopyTo(dest)
	assert.Equal(t, ms, dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newMetric(&otlpmetrics.Metric{}, &sharedState)) })
}

func TestMetric_Summary(t *testing.T) {
	ms := NewMetric()
	fillTestSummary(ms.SetEmptySummary())
	assert.Equal(t, MetricTypeSummary, ms.Type())
	assert.Equal(t, generateTestSummary(), ms.Summary())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newMetric(&otlpmetrics.Metric{}, &sharedState).SetEmptySummary() })
}

func TestMetric_CopyTo_Summary(t *testing.T) {
	ms := NewMetric()
	fillTestSummary(ms.SetEmptySummary())
	dest := NewMetric()
	ms.CopyTo(dest)
	assert.Equal(t, ms, dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newMetric(&otlpmetrics.Metric{}, &sharedState)) })
}

func generateTestMetric() Metric {
	tv := NewMetric()
	fillTestMetric(tv)
	return tv
}

func fillTestMetric(tv Metric) {
	tv.orig.Name = "test_name"
	tv.orig.Description = "test_description"
	tv.orig.Unit = "1"
	internal.FillTestMap(internal.NewMap(&tv.orig.Metadata, tv.state))
	tv.orig.Data = &otlpmetrics.Metric_Sum{Sum: &otlpmetrics.Sum{}}
	fillTestSum(newSum(tv.orig.GetSum(), tv.state))
}
