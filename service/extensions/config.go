// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extensions // import "github.com/oodle-ai/opentelemetry-collector/service/extensions"

import "github.com/oodle-ai/opentelemetry-collector/component"

// Config represents the ordered list of extensions configured for the service.
type Config []component.ID
