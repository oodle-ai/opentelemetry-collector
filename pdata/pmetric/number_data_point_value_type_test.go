// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetric // import "github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberDataPointValueTypeString(t *testing.T) {
	assert.Equal(t, "Empty", NumberDataPointValueTypeEmpty.String())
	assert.Equal(t, "Int", NumberDataPointValueTypeInt.String())
	assert.Equal(t, "Double", NumberDataPointValueTypeDouble.String())
	assert.Equal(t, "", (NumberDataPointValueTypeDouble + 1).String())
}
