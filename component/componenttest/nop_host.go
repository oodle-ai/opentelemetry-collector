// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package componenttest // import "github.com/oodle-ai/opentelemetry-collector/component/componenttest"

import (
	"github.com/oodle-ai/opentelemetry-collector/component"
)

// nopHost mocks a receiver.ReceiverHost for test purposes.
type nopHost struct{}

// NewNopHost returns a new instance of nopHost with proper defaults for most tests.
func NewNopHost() component.Host {
	return &nopHost{}
}

func (nh *nopHost) GetFactory(component.Kind, component.Type) component.Factory {
	return nil
}

func (nh *nopHost) GetExtensions() map[component.ID]component.Component {
	return nil
}

func (nh *nopHost) GetExporters() map[component.DataType]map[component.ID]component.Component {
	return nil
}
