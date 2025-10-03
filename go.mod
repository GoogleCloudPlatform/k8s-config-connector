module github.com/GoogleCloudPlatform/k8s-config-connector

go 1.24.0

toolchain go1.24.6

replace github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp => ./mockgcp

require (
	cloud.google.com/go/accesscontextmanager v1.9.6
	cloud.google.com/go/aiplatform v1.99.0
	cloud.google.com/go/alloydb v1.16.1
	cloud.google.com/go/apigateway v1.7.6
	cloud.google.com/go/apikeys v1.1.12
	cloud.google.com/go/apphub v0.2.4
	cloud.google.com/go/asset v1.21.1
	cloud.google.com/go/backupdr v1.4.0
	cloud.google.com/go/batch v1.12.2
	cloud.google.com/go/bigquery v1.69.0
	cloud.google.com/go/billing v1.20.4
	cloud.google.com/go/certificatemanager v1.9.5
	cloud.google.com/go/cloudbuild v1.22.2
	cloud.google.com/go/clouddms v1.8.7
	cloud.google.com/go/cloudquotas v1.3.2
	cloud.google.com/go/cloudtasks v1.13.6
	cloud.google.com/go/compute v1.38.0
	cloud.google.com/go/datacatalog v1.26.0
	cloud.google.com/go/dataflow v0.11.0
	cloud.google.com/go/dataform v0.12.0
	cloud.google.com/go/dataplex v1.25.3
	cloud.google.com/go/dataproc/v2 v2.11.2
	cloud.google.com/go/datastream v1.14.1
	cloud.google.com/go/deploy v1.27.2
	cloud.google.com/go/discoveryengine v1.15.0
	cloud.google.com/go/documentai v1.37.0
	cloud.google.com/go/edgecontainer v1.4.3
	cloud.google.com/go/essentialcontacts v1.7.6
	cloud.google.com/go/eventarc v1.15.5
	cloud.google.com/go/firestore v1.18.0
	cloud.google.com/go/gkebackup v1.8.0
	cloud.google.com/go/gkemulticloud v1.5.3
	cloud.google.com/go/iam v1.5.2
	cloud.google.com/go/iap v1.11.2
	cloud.google.com/go/kms v1.22.0
	cloud.google.com/go/logging v1.13.0
	cloud.google.com/go/managedkafka v0.4.0
	cloud.google.com/go/memorystore v0.3.0
	cloud.google.com/go/metastore v1.14.7
	cloud.google.com/go/monitoring v1.24.2
	cloud.google.com/go/netapp v1.7.1
	cloud.google.com/go/networkmanagement v1.19.1
	cloud.google.com/go/networksecurity v0.10.6
	cloud.google.com/go/networkservices v0.5.0
	cloud.google.com/go/notebooks v1.12.6
	cloud.google.com/go/orchestration v1.11.9
	cloud.google.com/go/orgpolicy v1.15.0
	cloud.google.com/go/privilegedaccessmanager v0.2.5
	cloud.google.com/go/profiler v0.4.1
	cloud.google.com/go/pubsub/v2 v2.0.0
	cloud.google.com/go/recaptchaenterprise/v2 v2.20.4
	cloud.google.com/go/redis v1.18.2
	cloud.google.com/go/resourcemanager v1.10.6
	cloud.google.com/go/run v1.10.0
	cloud.google.com/go/secretmanager v1.14.7
	cloud.google.com/go/securesourcemanager v1.1.1
	cloud.google.com/go/security v1.18.5
	cloud.google.com/go/spanner v1.82.0
	cloud.google.com/go/speech v1.27.1
	cloud.google.com/go/storage v1.55.0
	cloud.google.com/go/vmwareengine v1.3.5
	cloud.google.com/go/workflows v1.14.2
	cloud.google.com/go/workstations v1.1.1
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	github.com/GoogleCloudPlatform/declarative-resource-client-library v1.62.0
	github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder v0.0.0-20250915202832-88fd3f7499ff
	github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp v0.0.0-20240614222432-4bde5b345380
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30
	github.com/blang/semver v3.5.1+incompatible
	github.com/blang/semver/v4 v4.0.0
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/fatih/color v1.18.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v1.4.3
	github.com/go-logr/zapr v1.3.0
	github.com/google/go-cmp v0.7.0
	github.com/google/uuid v1.6.0
	github.com/googleapis/gax-go/v2 v2.15.0
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/hcl/v2 v2.24.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.24.0
	github.com/hashicorp/terraform-provider-google-beta v0.0.0-00010101000000-000000000000
	github.com/nasa9084/go-openapi v0.0.0-20200604141640-2875b7376353
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/gomega v1.36.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.20.4
	github.com/prometheus/procfs v0.15.1
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.10.0
	github.com/tmccombs/hcl2json v0.6.7
	github.com/zclconf/go-cty v1.16.3
	go.opencensus.io v0.24.0
	go.uber.org/mock v0.5.2
	go.uber.org/zap v1.27.0
	golang.org/x/oauth2 v0.30.0
	golang.org/x/sync v0.17.0
	golang.org/x/time v0.12.0
	google.golang.org/api v0.246.0
	google.golang.org/genproto v0.0.0-20250603155806-513f23925822
	google.golang.org/genproto/googleapis/api v0.0.0-20250804133106-a7a43d27e69b
	google.golang.org/genproto/googleapis/api/serviceusage v0.0.0-20250519155744-55703ea1f237
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.7
	gopkg.in/dnaeon/go-vcr.v3 v3.2.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.32.1
	k8s.io/apiextensions-apiserver v0.32.1
	k8s.io/apimachinery v0.32.1
	k8s.io/client-go v0.32.1
	k8s.io/klog/v2 v2.130.1
	sigs.k8s.io/controller-runtime v0.20.4
	sigs.k8s.io/controller-tools v0.16.5
	sigs.k8s.io/kubebuilder-declarative-pattern v0.20.0-beta.1.0.20250514194322-871029137730
	sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver v0.0.0-20230303024857-d1f76c15e05b
	sigs.k8s.io/structured-merge-diff/v4 v4.4.2
	sigs.k8s.io/yaml v1.4.0
)

require go.opentelemetry.io/auto/sdk v1.1.0 // indirect

require (
	cloud.google.com/go/osconfig v1.14.6 // indirect
	github.com/envoyproxy/go-control-plane/envoy v1.32.4 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-jose/go-jose/v4 v4.0.5 // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/spiffe/go-spiffe/v2 v2.5.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/zeebo/errs v1.4.0 // indirect
	golang.org/x/telemetry v0.0.0-20250908211612-aef8a434d053 // indirect
	golang.org/x/tools/godoc v0.1.0-deprecated // indirect
	gonum.org/v1/gonum v0.16.0 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	sigs.k8s.io/kubebuilder-declarative-pattern/ktest v0.0.0-20250514194322-871029137730 // indirect
)

require (
	bitbucket.org/creachadair/stringset v0.0.8 // indirect
	cel.dev/expr v0.24.0 // indirect
	cloud.google.com/go v0.121.2 // indirect
	cloud.google.com/go/auth v0.16.3 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
	cloud.google.com/go/bigtable v1.38.0
	cloud.google.com/go/compute/metadata v0.7.0 // indirect
	cloud.google.com/go/longrunning v0.6.7
	dario.cat/mergo v1.0.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect; indsirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ProtonMail/go-crypto v1.1.5 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apache/arrow/go/v15 v15.0.2 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/cncf/xds/go v0.0.0-20250501225837-2ac532fd4443 // indirect
	github.com/cyphar/filepath-securejoin v0.3.6 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.12.1 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
	github.com/evanphx/json-patch/v5 v5.9.11 // indirect
	github.com/exponent-io/jsonpath v0.0.0-20210407135951-1de76d718b3f // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gammazero/deque v0.0.0-20190521012701-46e4ffb7a622 // indirect
	github.com/gammazero/workerpool v0.0.0-20190608213748-0ed5e40ec55e // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.6.2 // indirect
	github.com/go-git/go-git/v5 v5.13.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/gobuffalo/flect v1.0.3 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.2.5 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4
	github.com/google/btree v1.1.3 // indirect
	github.com/google/go-cpy v0.0.0-20211218193943-a9c933c06932 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20241029153458-d1b30febd7db // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.6 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.1 // indirect
	github.com/hashicorp/go-plugin v1.4.8 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hc-install v0.4.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.17.3 // indirect
	github.com/hashicorp/terraform-json v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-framework v1.1.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.14.3 // indirect
	github.com/hashicorp/terraform-plugin-log v0.7.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.1.0 // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/spdystream v0.5.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pjbgf/sha1cd v0.3.2 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.59.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/skeema/knownhosts v1.3.0 // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xlab/treeprint v1.2.0 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.61.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.61.0 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/exp v0.0.0-20250911091902-df9299821621 // indirect
	golang.org/x/mod v0.28.0 // indirect
	golang.org/x/net v0.44.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/term v0.35.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	golang.org/x/tools v0.37.0
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/cli-runtime v0.32.1 // indirect
	k8s.io/component-base v0.32.1 // indirect
	k8s.io/kube-openapi v0.0.0-20241105132330-32ad38e42d3f // indirect
	k8s.io/kubectl v0.32.1 // indirect
	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738
	sigs.k8s.io/json v0.0.0-20241010143419-9aa6b5e7a4b3 // indirect
	sigs.k8s.io/kubebuilder-declarative-pattern/applylib v0.0.0-20230420203711-4abaa68e1923 // indirect
	sigs.k8s.io/kustomize/api v0.18.0 // indirect
	sigs.k8s.io/kustomize/kstatus v0.0.2-0.20200509233124-065f70705d4d // indirect
	sigs.k8s.io/kustomize/kyaml v0.18.1 // indirect
)

replace github.com/hashicorp/terraform-provider-google-beta => ./third_party/github.com/hashicorp/terraform-provider-google-beta

replace github.com/GoogleCloudPlatform/declarative-resource-client-library => ./third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library
