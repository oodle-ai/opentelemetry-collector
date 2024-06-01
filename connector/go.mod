module github.com/oodle-ai/opentelemetry-collector/connector

go 1.21.0

require (
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.9.0
	github.com/oodle-ai/opentelemetry-collector v0.101.0
	github.com/oodle-ai/opentelemetry-collector/component v0.101.0
	github.com/oodle-ai/opentelemetry-collector/consumer v0.101.0
	github.com/oodle-ai/opentelemetry-collector/pdata v1.8.0
	github.com/oodle-ai/opentelemetry-collector/pdata/testdata v0.101.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/config/configtelemetry v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/confmap v0.101.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.49.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240520151616-dc85e6b867a5 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/oodle-ai/opentelemetry-collector => ../

replace github.com/oodle-ai/opentelemetry-collector/component => ../component

replace github.com/oodle-ai/opentelemetry-collector/config/configtelemetry => ../config/configtelemetry

replace github.com/oodle-ai/opentelemetry-collector/confmap => ../confmap

replace github.com/oodle-ai/opentelemetry-collector/consumer => ../consumer

replace github.com/oodle-ai/opentelemetry-collector/featuregate => ../featuregate

replace github.com/oodle-ai/opentelemetry-collector/pdata => ../pdata

replace github.com/oodle-ai/opentelemetry-collector/pdata/testdata => ../pdata/testdata
