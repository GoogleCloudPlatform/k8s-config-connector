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
	cd bin; rm -rf protoc; mkdir protoc; cd protoc; unzip ../protoc.zip

GOOGLEAPI_VERSION?=$(shell grep https://github.com/googleapis/googleapis git.versions | awk '{print $$2}' )
.PHONY: sync-repo
sync-repo:
	mkdir -p third_party
	@if [ ! -d "third_party/googleapis" ]; then \
			echo "Cloning googleapis repository..."; \
			git clone https://github.com/googleapis/googleapis.git third_party/googleapis; \
	fi
	@echo "Syncing to commit $(GOOGLEAPI_VERSION)...";
	@cd third_party/googleapis && git fetch --all && git reset --hard $(GOOGLEAPI_VERSION)

.PHONY: gen-proto
gen-proto: tools sync-repo
	mkdir -p ./generated

	./apply-proto-patches.sh
	./fixup-third-party.sh

	make gen-proto-no-fixup

.PHONY: gen-proto-no-fixup
gen-proto-no-fixup:
	PATH=bin/:${PATH} bin/protoc/bin/protoc \
		-I ./third_party/googleapis \
		-I ./apis \
		--go_out ./generated \
		--go_opt paths=source_relative \
		--go-grpc_out ./generated \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./generated \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--experimental_allow_proto3_optional \
		./apis/google/apps/cloudidentity/groups/v1beta1/*.proto \
		./apis/mockgcp/cloud/apigee/v1/*.proto \
		./apis/mockgcp/cloud/bigquery/v2/*.proto \
		./apis/mockgcp/cloud/ids/v1/*.proto \
		./apis/mockgcp/cloud/networkconnectivity/v1/*.proto \
		./apis/mockgcp/cloud/servicenetworking/v1/*.proto \
		./apis/mockgcp/cloud/resourcemanager/v1/*.proto \
		./apis/mockgcp/storage/v1/*.proto \
		./third_party/googleapis/mockgcp/api/apikeys/v2/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/annotation*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/data_item.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/dataset*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/encryption_spec.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/api_auth.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/artifact.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/lineage_subgraph.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/context.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/execution.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/event.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/metadata*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/model*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/deploy*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/env_var.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/evaluated_annotation.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/job_state.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/publisher_model.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/io.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/operation.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/saved_query.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/tensorboard*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/user_action_reference.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/entity_type.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/feature.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/feature_monitor.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/feature_selector.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/featurestore*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/feature_monitoring_stats.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/types.proto \
		./third_party/googleapis/mockgcp/cloud/apigateway/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/apphub/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/alloydb/v1beta/*.proto \
		./third_party/googleapis/mockgcp/cloud/batch/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/bigquery/analyticshub/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/bigquery/connection/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/bigquery/datatransfer/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/billing/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/certificatemanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/tasks/v2/*.proto \
		./third_party/googleapis/mockgcp/cloud/clouddms/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/common/*.proto \
		./third_party/googleapis/mockgcp/cloud/compute/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/documentai/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/edgecontainer/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/edgenetwork/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/filestore/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/functions/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/policycontroller/v1beta/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/configmanagement/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/metering/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/multiclusteringress/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta1/membership.proto \
		./third_party/googleapis/mockgcp/cloud/gkemulticloud/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/kms/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/managedkafka/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/deploy/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/networkservices/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/privilegedaccessmanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/pubsublite/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/redis/cluster/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/redis/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/resourcemanager/v3/*.proto \
		./third_party/googleapis/mockgcp/cloud/secretmanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/securesourcemanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/security/privateca/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/servicedirectory/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/sql/v1beta4/*.proto \
		./third_party/googleapis/mockgcp/cloud/tpu/v2/*.proto \
		./third_party/googleapis/mockgcp/cloud/vpcaccess/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/workflows/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/workflows/executions/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/workstations/v1/*.proto \
		./third_party/googleapis/mockgcp/container/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/dataflow/v1beta3/*.proto \
		./third_party/googleapis/mockgcp/devtools/artifactregistry/v1/*.proto \
		./third_party/googleapis/mockgcp/devtools/cloudbuild/v1/*.proto	 \
		./third_party/googleapis/mockgcp/firestore/admin/v1/*.proto \
		./third_party/googleapis/mockgcp/iam/admin/v1/*.proto \
		./third_party/googleapis/mockgcp/logging/v2/*.proto \
		./third_party/googleapis/mockgcp/monitoring/dashboard/v1/*.proto \
		./third_party/googleapis/mockgcp/monitoring/metricsscope/v1/*.proto \
		./third_party/googleapis/mockgcp/monitoring/v3/*.proto \
		./third_party/googleapis/mockgcp/pubsub/v1/*.proto \
		./third_party/googleapis/mockgcp/spanner/admin/database/v1/*.proto \
		./third_party/googleapis/mockgcp/spanner/admin/instance/v1/*.proto \
		./third_party/googleapis/mockgrafeas/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/bigquery/biglake/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/bigquery/reservation/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/notebooks/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/orchestration/airflow/service/v1/*.proto \
		./third_party/googleapis/mockgcp/api/cloudquotas/v1beta/*.proto \
		./third_party/googleapis/mockgcp/cloud/asset/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/schedule.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/schedule_service.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/pipeline_service.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/notebook_service.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/notebook_runtime.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/notebook_execution_job.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/network_spec.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/notebook_idle_shutdown_config.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/notebook_euc_config.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/notebook_runtime_template_ref.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/training_pipeline.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/pipeline_job.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/pipeline_state.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/service_networking.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/ui_pipeline_spec.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/pipeline_failure_policy.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/value.proto \
		./third_party/googleapis/mockgcp/cloud/datacatalog/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/eventarc/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/backupdr/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/essentialcontacts/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/networkmanagement/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/netapp/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/dataplex/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/vmwareengine/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/kms/v1/resources.proto \
		./third_party/googleapis/mockgcp/cloud/recaptchaenterprise/v1/*.proto

.PHONY: generate-grpc-for-google-protos
generate-grpc-for-google-protos:
	mkdir -p third_party
	git clone https://github.com/googleapis/googleapis.git third_party/googleapis || (cd third_party/googleapis && git fetch --all && git reset --hard $(GOOGLEAPI_VERSION) && git pull)
	mkdir -p ./generated

	PATH=bin/:${PATH} bin/protoc/bin/protoc \
		-I ./third_party/googleapis \
		-I ./apis \
		--experimental_allow_proto3_optional \
		--grpc-gateway_out ./generated \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt standalone=true \
		./third_party/googleapis/google/cloud/discoveryengine/v1/data_store_service.proto \
		./third_party/googleapis/google/cloud/discoveryengine/v1/engine_service.proto \
		./third_party/googleapis/google/bigtable/admin/v2/*.proto \
		./third_party/googleapis/google/cloud/dataproc/v1/*.proto \
		./third_party/googleapis/google/cloud/aiplatform/v1/model_garden_service.proto \
		./third_party/googleapis/google/storage/control/v2/*.proto \
		./third_party/googleapis/google/cloud/batch/v1/*.proto \
		./third_party/googleapis/google/cloud/edgecontainer/v1/*.proto \
		./third_party/googleapis/google/cloud/apigateway/v1/apigateway_service.proto

.PHONY: generate-protos-from-openapi
generate-protos-from-openapi:
	mkdir -p temp/

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

	wget -O temp/kms-api-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/cloudkms/v1/cloudkms-api.json
	mkdir -p apis/mockgcp/cloud/kms/v1/
	cd tools/gapic; go run . --proto-version=2 --proto-package mockgcp.cloud.kms.v1 ../../temp/kms-api-v1.json > ../../apis/mockgcp/cloud/kms/v1/service.proto

	wget -O temp/apigee-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/refs/heads/main/apigee/v1/apigee-api.json
	mkdir -p apis/mockgcp/cloud/apigee/v1/
	cd tools/gapic; go run . --proto-package mockgcp.cloud.apigee.v1 ../../temp/apigee-v1.json > ../../apis/mockgcp/cloud/apigee/v1/service.proto

	wget -O temp/cloudidentity-api.json https://raw.githubusercontent.com/googleapis/google-api-go-client/refs/heads/main/cloudidentity/v1beta1/cloudidentity-api.json
	mkdir -p apis/google/apps/cloudidentity/groups/v1beta1/
	cd tools/gapic; go run . --proto-version=2 --proto-package google.apps.cloudidentity.groups.v1beta1 ../../temp/cloudidentity-api.json > ../../apis/google/apps/cloudidentity/groups/v1beta1//service.proto

	rm -r temp
