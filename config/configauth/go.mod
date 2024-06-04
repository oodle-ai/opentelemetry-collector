module github.com/oodle-ai/opentelemetry-collector/config/configauth

go 1.21.0

require (
	github.com/stretchr/testify v1.9.0
	github.com/oodle-ai/opentelemetry-collector/component v0.101.0
	github.com/oodle-ai/opentelemetry-collector/extension v0.101.0
	github.com/oodle-ai/opentelemetry-collector/extension/auth v0.101.0
	go.uber.org/goleak v1.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/config/configtelemetry v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/confmap v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/pdata v1.8.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/oodle-ai/opentelemetry-collector/pdata => ../../pdata

replace github.com/oodle-ai/opentelemetry-collector/confmap => ../../confmap

replace github.com/oodle-ai/opentelemetry-collector/component => ../../component

replace github.com/oodle-ai/opentelemetry-collector/config/configtelemetry => ../configtelemetry

replace github.com/oodle-ai/opentelemetry-collector/extension => ../../extension

replace github.com/oodle-ai/opentelemetry-collector/extension/auth => ../../extension/auth
