// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumer // import "github.com/oodle-ai/opentelemetry-collector/consumer"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
)

// Traces is an interface that receives ptrace.Traces, processes it
// as needed, and sends it to the next processing node if any or to the destination.
type Traces interface {
	baseConsumer
	// ConsumeTraces receives ptrace.Traces for consumption.
	ConsumeTraces(ctx context.Context, td ptrace.Traces) error
}

// ConsumeTracesFunc is a helper function that is similar to ConsumeTraces.
type ConsumeTracesFunc func(ctx context.Context, td ptrace.Traces) error

// ConsumeTraces calls f(ctx, td).
func (f ConsumeTracesFunc) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	return f(ctx, td)
}

type baseTraces struct {
	*baseImpl
	ConsumeTracesFunc
}

// NewTraces returns a Traces configured with the provided options.
func NewTraces(consume ConsumeTracesFunc, options ...Option) (Traces, error) {
	if consume == nil {
		return nil, errNilFunc
	}
	return &baseTraces{
		baseImpl:          newBaseImpl(options...),
		ConsumeTracesFunc: consume,
	}, nil
}
