module github.com/oodle-ai/opentelemetry-collector/service

go 1.21.0

require (
	github.com/google/uuid v1.6.0
	github.com/prometheus/client_golang v1.19.1
	github.com/prometheus/client_model v0.6.1
	github.com/prometheus/common v0.53.0
	github.com/shirou/gopsutil/v3 v3.24.4
	github.com/stretchr/testify v1.9.0
	go.opencensus.io v0.24.0
	github.com/oodle-ai/opentelemetry-collector v0.101.0
	github.com/oodle-ai/opentelemetry-collector/component v0.101.0
	github.com/oodle-ai/opentelemetry-collector/config/confignet v0.101.0
	github.com/oodle-ai/opentelemetry-collector/config/configtelemetry v0.101.0
	github.com/oodle-ai/opentelemetry-collector/confmap v0.101.0
	github.com/oodle-ai/opentelemetry-collector/connector v0.101.0
	github.com/oodle-ai/opentelemetry-collector/consumer v0.101.0
	github.com/oodle-ai/opentelemetry-collector/exporter v0.101.0
	github.com/oodle-ai/opentelemetry-collector/extension v0.101.0
	github.com/oodle-ai/opentelemetry-collector/extension/zpagesextension v0.101.0
	github.com/oodle-ai/opentelemetry-collector/featuregate v1.8.0
	github.com/oodle-ai/opentelemetry-collector/pdata v1.8.0
	github.com/oodle-ai/opentelemetry-collector/pdata/testdata v0.101.0
	github.com/oodle-ai/opentelemetry-collector/processor v0.101.0
	github.com/oodle-ai/opentelemetry-collector/receiver v0.101.0
	github.com/oodle-ai/opentelemetry-collector/semconv v0.101.0
	go.opentelemetry.io/contrib/config v0.7.0
	go.opentelemetry.io/contrib/propagators/b3 v1.27.0
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/bridge/opencensus v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.27.0
	go.opentelemetry.io/otel/exporters/prometheus v0.49.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.27.0
	go.opentelemetry.io/otel/metric v1.27.0
	go.opentelemetry.io/otel/sdk v1.27.0
	go.opentelemetry.io/otel/sdk/metric v1.27.0
	go.opentelemetry.io/otel/trace v1.27.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	gonum.org/v1/gonum v0.15.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/contrib/zpages v0.52.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.27.0 // indirect
	go.opentelemetry.io/proto/otlp v1.2.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240520151616-dc85e6b867a5 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240520151616-dc85e6b867a5 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/oodle-ai/opentelemetry-collector => ../

replace github.com/oodle-ai/opentelemetry-collector/connector => ../connector

replace github.com/oodle-ai/opentelemetry-collector/component => ../component

replace github.com/oodle-ai/opentelemetry-collector/pdata => ../pdata

replace github.com/oodle-ai/opentelemetry-collector/pdata/testdata => ../pdata/testdata

replace github.com/oodle-ai/opentelemetry-collector/extension/zpagesextension => ../extension/zpagesextension

replace github.com/oodle-ai/opentelemetry-collector/extension => ../extension

replace github.com/oodle-ai/opentelemetry-collector/exporter => ../exporter

replace github.com/oodle-ai/opentelemetry-collector/confmap => ../confmap

replace github.com/oodle-ai/opentelemetry-collector/config/configtelemetry => ../config/configtelemetry

replace github.com/oodle-ai/opentelemetry-collector/processor => ../processor

replace github.com/oodle-ai/opentelemetry-collector/consumer => ../consumer

replace github.com/oodle-ai/opentelemetry-collector/semconv => ../semconv

replace github.com/oodle-ai/opentelemetry-collector/receiver => ../receiver

replace github.com/oodle-ai/opentelemetry-collector/featuregate => ../featuregate

replace github.com/oodle-ai/opentelemetry-collector/config/confignet => ../config/confignet

replace github.com/oodle-ai/opentelemetry-collector/config/configretry => ../config/configretry
