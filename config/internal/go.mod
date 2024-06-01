module github.com/oodle-ai/opentelemetry-collector/config/internal

go 1.21.0

require (
	github.com/stretchr/testify v1.9.0
	github.com/oodle-ai/opentelemetry-collector v0.101.0
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/featuregate v1.8.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/oodle-ai/opentelemetry-collector => ../../

replace github.com/oodle-ai/opentelemetry-collector/featuregate => ../../featuregate

replace github.com/oodle-ai/opentelemetry-collector/confmap => ../../confmap

replace github.com/oodle-ai/opentelemetry-collector/config/configtelemetry => ../configtelemetry

replace github.com/oodle-ai/opentelemetry-collector/pdata => ../../pdata

replace github.com/oodle-ai/opentelemetry-collector/consumer => ../../consumer

replace github.com/oodle-ai/opentelemetry-collector/component => ../../component

replace github.com/oodle-ai/opentelemetry-collector/pdata/testdata => ../../pdata/testdata
