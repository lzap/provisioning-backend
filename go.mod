module github.com/RHEnVision/provisioning-backend

// Also change: mk/tools.mk .github/workflows build/Dockerfile
go 1.19

require (
	cloud.google.com/go/compute v1.19.3
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.6.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.3.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork v1.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.1.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions v1.1.1
	github.com/IBM/pgxpoolprometheus v1.1.1
	github.com/Unleash/unleash-client-go/v3 v3.7.4
	github.com/archdx/zerolog-sentry v1.2.0
	github.com/aws/aws-sdk-go-v2 v1.18.0
	github.com/aws/aws-sdk-go-v2/config v1.18.25
	github.com/aws/aws-sdk-go-v2/credentials v1.13.24
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.11
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.97.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.19.12
	github.com/aws/aws-sdk-go-v2/service/sts v1.19.0
	github.com/aws/smithy-go v1.13.5
	github.com/deepmap/oapi-codegen v1.12.4
	github.com/exaring/otelpgx v0.4.1
	github.com/georgysavva/scany/v2 v2.0.0
	github.com/getkin/kin-openapi v0.116.0
	github.com/getsentry/sentry-go v0.21.0
	github.com/go-chi/chi/v5 v5.0.8
	github.com/go-chi/render v1.0.2
	github.com/go-openapi/runtime v0.26.0
	github.com/go-playground/mold/v4 v4.3.0
	github.com/go-playground/validator/v10 v10.13.0
	github.com/google/uuid v1.3.0
	github.com/ilyakaznacheev/cleanenv v1.4.2
	github.com/jackc/pgx-zerolog v0.0.0-20230315001418-f978528409eb
	github.com/jackc/pgx/v5 v5.3.1
	github.com/jackc/tern/v2 v2.1.0
	github.com/lzap/cloudwatchwriter2 v1.1.0
	github.com/prometheus/client_golang v1.15.1
	github.com/redhatinsights/app-common-go v1.6.6
	github.com/redhatinsights/platform-go-middlewares v0.20.0
	github.com/redis/go-redis/v9 v9.0.4
	github.com/riandyrn/otelchi v0.5.1
	github.com/rs/zerolog v1.29.1
	github.com/segmentio/kafka-go v0.4.40
	github.com/stretchr/testify v1.8.2
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.41.1
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/exporters/jaeger v1.15.1
	go.opentelemetry.io/otel/sdk v1.15.1
	go.opentelemetry.io/otel/trace v1.15.1
	go.uber.org/automaxprocs v1.5.2
	golang.org/x/crypto v0.9.0
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc
	google.golang.org/api v0.122.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.0.0 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.27 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.10 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/errors v0.20.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/spec v0.20.9 // indirect
	github.com/go-openapi/strfmt v0.21.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-openapi/validate v0.22.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/s2a-go v0.1.3 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.8.0 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/invopop/yaml v0.2.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/oleiade/lane/v2 v2.0.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.26.0 // indirect
	github.com/perimeterx/marshmallow v1.1.4 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.43.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/twmb/murmur3 v1.1.7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	go.mongodb.org/mongo-driver v1.11.6 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib v1.16.1 // indirect
	go.opentelemetry.io/otel/metric v0.38.1 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
