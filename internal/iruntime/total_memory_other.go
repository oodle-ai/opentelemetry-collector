// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package iruntime // import "github.com/oodle-ai/opentelemetry-collector/internal/iruntime"

// TotalMemory returns total available memory for non-linux platforms.
func TotalMemory() (uint64, error) {
	return readMemInfo()
}
