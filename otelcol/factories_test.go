// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol

import (
	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/connector"
	"github.com/oodle-ai/opentelemetry-collector/connector/connectortest"
	"github.com/oodle-ai/opentelemetry-collector/exporter"
	"github.com/oodle-ai/opentelemetry-collector/exporter/exportertest"
	"github.com/oodle-ai/opentelemetry-collector/extension"
	"github.com/oodle-ai/opentelemetry-collector/extension/extensiontest"
	"github.com/oodle-ai/opentelemetry-collector/processor"
	"github.com/oodle-ai/opentelemetry-collector/processor/processortest"
	"github.com/oodle-ai/opentelemetry-collector/receiver"
	"github.com/oodle-ai/opentelemetry-collector/receiver/receivertest"
)

func nopFactories() (Factories, error) {
	var factories Factories
	var err error

	if factories.Connectors, err = connector.MakeFactoryMap(connectortest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	if factories.Extensions, err = extension.MakeFactoryMap(extensiontest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	if factories.Receivers, err = receiver.MakeFactoryMap(receivertest.NewNopFactory(), receivertest.NewNopFactoryForType(component.DataTypeLogs)); err != nil {
		return Factories{}, err
	}

	if factories.Exporters, err = exporter.MakeFactoryMap(exportertest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	if factories.Processors, err = processor.MakeFactoryMap(processortest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	return factories, err
}
