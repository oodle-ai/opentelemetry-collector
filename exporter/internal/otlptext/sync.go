// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlptext // import "github.com/oodle-ai/opentelemetry-collector/exporter/internal/otlptext"

import (
	"context"
	"errors"
	"os"

	"go.uber.org/zap"
)

func LoggerSync(logger *zap.Logger) func(context.Context) error {
	return func(context.Context) error {
		// Currently Sync() return a different error depending on the OS.
		// Since these are not actionable ignore them.
		err := logger.Sync()
		osErr := &os.PathError{}
		if errors.As(err, &osErr) {
			wrappedErr := osErr.Unwrap()
			if knownSyncError(wrappedErr) {
				err = nil
			}
		}
		return err
	}
}
