// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package connector // import "github.com/oodle-ai/opentelemetry-collector/connector"

import (
	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/consumer"
	"github.com/oodle-ai/opentelemetry-collector/internal/fanoutconsumer"
)

// MetricsRouterAndConsumer feeds the first consumer.Metrics in each of the specified pipelines.
type MetricsRouterAndConsumer interface {
	consumer.Metrics
	Consumer(...component.ID) (consumer.Metrics, error)
	PipelineIDs() []component.ID
	privateFunc()
}

type metricsRouter struct {
	consumer.Metrics
	baseRouter[consumer.Metrics]
}

func NewMetricsRouter(cm map[component.ID]consumer.Metrics) MetricsRouterAndConsumer {
	consumers := make([]consumer.Metrics, 0, len(cm))
	for _, cons := range cm {
		consumers = append(consumers, cons)
	}
	return &metricsRouter{
		Metrics:    fanoutconsumer.NewMetrics(consumers),
		baseRouter: newBaseRouter(fanoutconsumer.NewMetrics, cm),
	}
}

func (r *metricsRouter) privateFunc() {}
