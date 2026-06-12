.PHONY: all
all: tools generate-protos-from-openapi generate-grpc-for-google-protos gen-proto

.PHONY: tools
tools:
	mkdir -p bin/
	GOBIN=`pwd`/bin/ go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.3 \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.3
	GOBIN=`pwd`/bin/ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=`pwd`/bin/ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	wget -N -O bin/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-linux-x86_64.zip
	cd bin; rm -rf protoc; mkdir protoc; cd protoc; python3 -m zipfile -e ../protoc.zip .
	chmod +x bin/protoc/bin/protoc

GOOGLEAPI_VERSION?=$(shell grep https://github.com/googleapis/googleapis git.versions | awk '{print $$2}' )
.PHONY: sync-repo
sync-repo:
	mkdir -p third_party
	@if [ ! -d "third_party/googleapis" ]; then \
			echo "Cloning googleapis repository..."; \
			git clone --depth 1 https://github.com/googleapis/googleapis.git third_party/googleapis; \
	fi
	@echo "Syncing to commit $(GOOGLEAPI_VERSION)...";
	cd third_party/googleapis && (git cat-file -e $(GOOGLEAPI_VERSION)^{commit} 2> /dev/null || git fetch origin $(GOOGLEAPI_VERSION)) && git reset --hard $(GOOGLEAPI_VERSION) && git clean -df . && git status

.PHONY: gen-proto
gen-proto: tools sync-repo
	mkdir -p ./generated

	./apply-proto-patches.sh
	./fixup-third-party.sh

	make gen-proto-no-fixup

.PHONY: gen-proto-no-fixup
gen-proto-no-fixup:
	PATH=bin/:${PATH} bin/protoc/bin/protoc \
		-I ./apis \
		-I ./third_party/googleapis \
		--go_out ./generated \
		--go_opt paths=source_relative \
		--go-grpc_out ./generated \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./generated \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--experimental_allow_proto3_optional \
		./apis/google/apps/cloudidentity/v1beta1/*.proto \
		./apis/google/cloud/dns/v1/*.proto \
		./apis/google/cloud/networkservices/v1/*.proto \
		./apis/mockgcp/cloud/apigee/v1/*.proto \
		./apis/mockgcp/cloud/bigquery/v2/*.proto \
		./apis/mockgcp/cloud/ids/v1/*.proto \
		./apis/mockgcp/cloud/networkconnectivity/v1/*.proto \
		./apis/mockgcp/cloud/servicenetworking/v1/*.proto \
		./apis/mockgcp/cloud/resourcemanager/v1/*.proto \
		./apis/mockgcp/cloud/privilegedaccessmanager/v1/*.proto \
		./apis/mockgcp/storage/v1/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/apigateway/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/alloydb/v1beta/*.proto \
		./third_party/googleapis/mockgcp/cloud/tasks/v2/*.proto \
		./third_party/googleapis/mockgcp/cloud/common/*.proto \
		./third_party/googleapis/mockgcp/cloud/compute/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/documentai/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/filestore/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/functions/v1/*.proto \
		./apis/mockgcp/cloud/gkehub/v1beta/*.proto \
		./apis/mockgcp/cloud/gkehub/policycontroller/v1beta/*.proto \
		./apis/mockgcp/cloud/gkehub/v1beta/configmanagement/*.proto \
		./apis/mockgcp/cloud/gkehub/v1beta/metering/*.proto \
		./apis/mockgcp/cloud/gkehub/v1beta/multiclusteringress/*.proto \
		./apis/mockgcp/cloud/gkehub/v1beta1/membership.proto \
		./third_party/googleapis/mockgcp/cloud/privilegedaccessmanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/pubsublite/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/redis/cluster/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/redis/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/resourcemanager/v3/*.proto \
		./third_party/googleapis/mockgcp/cloud/security/privateca/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/servicedirectory/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/sql/v1beta4/*.proto \
		./third_party/googleapis/mockgcp/cloud/tpu/v2/*.proto \
		./third_party/googleapis/mockgcp/container/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/dataflow/v1beta3/*.proto \
		./third_party/googleapis/mockgcp/devtools/cloudbuild/v1/*.proto	 \
		./third_party/googleapis/mockgcp/iam/admin/v1/*.proto \
		./third_party/googleapis/mockgcp/iam/v1/resource_policy_member.proto \
		./third_party/googleapis/mockgcp/spanner/admin/database/v1/*.proto \
		./third_party/googleapis/mockgcp/spanner/admin/instance/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/bigquery/biglake/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/bigquery/reservation/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/orchestration/airflow/service/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/datacatalog/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/metastore/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkebackup/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/run/v2/*.proto \
		./third_party/googleapis/mockgcp/devtools/sourcerepo/v1/*.proto


.PHONY: generate-grpc-for-google-protos
generate-grpc-for-google-protos:
	mkdir -p third_party
	if [ ! -d "third_party/googleapis" ]; then \
			echo "Cloning googleapis repository..."; \
			git clone --depth 1 https://github.com/googleapis/googleapis.git third_party/googleapis; \
	fi
	cd third_party/googleapis && (git cat-file -e $(GOOGLEAPI_VERSION)^{commit} 2> /dev/null || git fetch origin $(GOOGLEAPI_VERSION)) && git reset --hard $(GOOGLEAPI_VERSION) && git clean -df . && git status
	mkdir -p ./generated

	PATH=bin/:${PATH} bin/protoc/bin/protoc \
		-I ./apis \
		-I ./third_party/googleapis \
		--experimental_allow_proto3_optional \
		--grpc-gateway_out ./generated \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt standalone=true \
		./third_party/googleapis/google/bigtable/admin/v2/*.proto \
		./third_party/googleapis/google/storage/control/v2/*.proto \
		./third_party/googleapis/google/cloud/apigateway/v1/apigateway_service.proto \
		./third_party/googleapis/google/cloud/networkservices/v1/*.proto \
		./third_party/googleapis/google/cloud/recaptchaenterprise/v1/*.proto \
		./third_party/googleapis/google/iam/v2/*.proto

.PHONY: generate-protos-from-openapi
generate-protos-from-openapi:
	mkdir -p temp/

	wget -O temp/cloudresourcemanager-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/refs/heads/main/cloudresourcemanager/v1/cloudresourcemanager-api.json
	mkdir -p apis/mockgcp/cloud/resourcemanager/v1
	cd tools/gapic; go run . --proto-version=2 --proto-package mockgcp.cloudresourcemanager.v1 ../../temp/cloudresourcemanager-v1.json > ../../apis/mockgcp/cloud/resourcemanager/v1/api.proto

	wget -O temp/bigquery-v2.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/bigquery/v2/bigquery-api.json
	mkdir -p apis/mockgcp/cloud/bigquery/v2/
	cd tools/gapic; go run . --proto-version=2 --proto-package mockgcp.cloud.bigquery.v2 ../../temp/bigquery-v2.json > ../../apis/mockgcp/cloud/bigquery/v2/api.proto

	wget -O temp/networkconnectivity-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/networkconnectivity/v1/networkconnectivity-api.json
	mkdir -p apis/mockgcp/cloud/networkconnectivity/v1/
	cd tools/gapic; go run . --proto-package mockgcp.cloud.networkconnectivity.v1 ../../temp/networkconnectivity-v1.json > ../../apis/mockgcp/cloud/networkconnectivity/v1/networkconnectivity.proto

	wget -O temp/servicenetworking-api.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/servicenetworking/v1/servicenetworking-api.json
	mkdir -p apis/mockgcp/cloud/servicenetworking/v1/
	cd tools/gapic; go run . --proto-package mockgcp.cloud.servicenetworking.v1 ../../temp/servicenetworking-api.json > ../../apis/mockgcp/cloud/servicenetworking/v1/servicenetworking.proto

	wget -O temp/storage-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/storage/v1/storage-api.json
	mkdir -p apis/mockgcp/storage/v1/
	cd tools/gapic; go run . --proto-version=2 --proto-package mockgcp.storage.v1 ../../temp/storage-v1.json > ../../apis/mockgcp/storage/v1/service.proto

	wget -O temp/ids-api-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/b49e3b908a8ed562e068736f1c42e992538ba6e0/ids/v1/ids-api.json
	mkdir -p apis/mockgcp/cloud/ids/v1/
	cd tools/gapic; go run . --proto-version=2 --proto-package mockgcp.cloud.ids.v1 ../../temp/ids-api-v1.json > ../../apis/mockgcp/cloud/ids/v1/service.proto


	wget -O temp/apigee-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/refs/heads/main/apigee/v1/apigee-api.json
	mkdir -p apis/mockgcp/cloud/apigee/v1/
	cd tools/gapic; go run . --proto-package mockgcp.cloud.apigee.v1 ../../temp/apigee-v1.json > ../../apis/mockgcp/cloud/apigee/v1/service.proto

	wget -O temp/cloudidentity-api.json https://raw.githubusercontent.com/googleapis/google-api-go-client/refs/heads/main/cloudidentity/v1beta1/cloudidentity-api.json
	mkdir -p apis/google/apps/cloudidentity/v1beta1/
	cd tools/gapic; go run . --proto-version=2 --proto-package google.apps.cloudidentity.v1beta1 ../../temp/cloudidentity-api.json > ../../apis/google/apps/cloudidentity/v1beta1//service.proto

	wget -O temp/dns-api.json https://raw.githubusercontent.com/googleapis/google-api-go-client/refs/heads/main/dns/v1/dns-api.json
	mkdir -p apis/google/cloud/dns/v1/
	# DNS sadly uses a bespoke operation type
	cd tools/gapic; go run . --proto-version=2 --proto-package google.cloud.dns.v1 --normalize-lro=false ../../temp/dns-api.json > ../../apis/google/cloud/dns/v1/service.proto

	rm -r temp
