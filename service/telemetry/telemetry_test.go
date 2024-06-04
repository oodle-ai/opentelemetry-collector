// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/oodle-ai/opentelemetry-collector/config/configtelemetry"
)

func TestTelemetryConfiguration(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		success bool
	}{
		{
			name: "Valid config",
			cfg: &Config{
				Logs: LogsConfig{
					Level:    zapcore.DebugLevel,
					Encoding: "console",
				},
				Metrics: MetricsConfig{
					Level:   configtelemetry.LevelBasic,
					Address: "127.0.0.1:3333",
				},
			},
			success: true,
		},
		{
			name: "Invalid config",
			cfg: &Config{
				Logs: LogsConfig{
					Level: zapcore.DebugLevel,
				},
				Metrics: MetricsConfig{
					Level:   configtelemetry.LevelBasic,
					Address: "127.0.0.1:3333",
				},
			},
			success: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFactory()
			set := Settings{ZapOptions: []zap.Option{}}
			logger, err := f.CreateLogger(context.Background(), set, tt.cfg)
			if tt.success {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
			} else {
				assert.Error(t, err)
				assert.Nil(t, logger)
			}
		})
	}
}

func TestSampledLogger(t *testing.T) {
	tests := []struct {
		name string
		cfg  *Config
	}{
		{
			name: "Default sampling",
			cfg: &Config{
				Logs: LogsConfig{
					Encoding: "console",
				},
			},
		},
		{
			name: "Custom sampling",
			cfg: &Config{
				Logs: LogsConfig{
					Level:    zapcore.DebugLevel,
					Encoding: "console",
					Sampling: &LogsSamplingConfig{
						Enabled:    true,
						Tick:       1 * time.Second,
						Initial:    100,
						Thereafter: 100,
					},
				},
			},
		},
		{
			name: "Disable sampling",
			cfg: &Config{
				Logs: LogsConfig{
					Level:    zapcore.DebugLevel,
					Encoding: "console",
					Sampling: &LogsSamplingConfig{
						Enabled: false,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFactory()
			ctx := context.Background()
			set := Settings{ZapOptions: []zap.Option{}}
			logger, err := f.CreateLogger(ctx, set, tt.cfg)
			assert.NoError(t, err)
			assert.NotNil(t, logger)
		})
	}
}

func TestTelemetryShutdown(t *testing.T) {
	tests := []struct {
		name     string
		provider trace.TracerProvider
		wantErr  error
	}{
		{
			name: "No provider",
		},
		{
			name:     "Non-SDK provider",
			provider: noop.NewTracerProvider(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			telemetry := Telemetry{tracerProvider: tt.provider}
			err := telemetry.Shutdown(context.Background())
			require.Equal(t, tt.wantErr, err)
		})
	}
}
