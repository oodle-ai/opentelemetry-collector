module github.com/oodle-ai/opentelemetry-collector/exporter/otlphttpexporter

go 1.21.0

require (
	github.com/stretchr/testify v1.9.0
	github.com/oodle-ai/opentelemetry-collector v0.101.0
	github.com/oodle-ai/opentelemetry-collector/component v0.101.0
	github.com/oodle-ai/opentelemetry-collector/config/configcompression v1.8.0
	github.com/oodle-ai/opentelemetry-collector/config/confighttp v0.101.0
	github.com/oodle-ai/opentelemetry-collector/config/configopaque v1.8.0
	github.com/oodle-ai/opentelemetry-collector/config/configretry v0.101.0
	github.com/oodle-ai/opentelemetry-collector/config/configtls v0.101.0
	github.com/oodle-ai/opentelemetry-collector/confmap v0.101.0
	github.com/oodle-ai/opentelemetry-collector/consumer v0.101.0
	github.com/oodle-ai/opentelemetry-collector/exporter v0.101.0
	github.com/oodle-ai/opentelemetry-collector/pdata v1.8.0
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240520151616-dc85e6b867a5
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
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
	github.com/rs/cors v1.10.1 // indirect
	github.com/oodle-ai/opentelemetry-collector/config/configauth v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/config/configtelemetry v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/config/internal v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/extension v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/extension/auth v0.101.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/featuregate v1.8.0 // indirect
	github.com/oodle-ai/opentelemetry-collector/receiver v0.101.0 // indirect
	go.opentelemetry.io/contrib/config v0.7.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.49.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.27.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	go.opentelemetry.io/proto/otlp v1.2.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240520151616-dc85e6b867a5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/oodle-ai/opentelemetry-collector => ../../

replace github.com/oodle-ai/opentelemetry-collector/component => ../../component

replace github.com/oodle-ai/opentelemetry-collector/config/configauth => ../../config/configauth

replace github.com/oodle-ai/opentelemetry-collector/config/configcompression => ../../config/configcompression

replace github.com/oodle-ai/opentelemetry-collector/config/confighttp => ../../config/confighttp

replace github.com/oodle-ai/opentelemetry-collector/config/configopaque => ../../config/configopaque

replace github.com/oodle-ai/opentelemetry-collector/config/configtelemetry => ../../config/configtelemetry

replace github.com/oodle-ai/opentelemetry-collector/config/configtls => ../../config/configtls

replace github.com/oodle-ai/opentelemetry-collector/config/internal => ../../config/internal

replace github.com/oodle-ai/opentelemetry-collector/confmap => ../../confmap

replace github.com/oodle-ai/opentelemetry-collector/exporter => ../

replace github.com/oodle-ai/opentelemetry-collector/extension => ../../extension

replace github.com/oodle-ai/opentelemetry-collector/extension/auth => ../../extension/auth

replace github.com/oodle-ai/opentelemetry-collector/featuregate => ../../featuregate

replace github.com/oodle-ai/opentelemetry-collector/pdata => ../../pdata

replace github.com/oodle-ai/opentelemetry-collector/pdata/testdata => ../../pdata/testdata

replace github.com/oodle-ai/opentelemetry-collector/receiver => ../../receiver

replace github.com/oodle-ai/opentelemetry-collector/consumer => ../../consumer

retract (
	v0.76.0 // Depends on retracted pdata v1.0.0-rc10 module, use v0.76.1
	v0.69.0 // Release failed, use v0.69.1
)

replace github.com/oodle-ai/opentelemetry-collector/config/configretry => ../../config/configretry
