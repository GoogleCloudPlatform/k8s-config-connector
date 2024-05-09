.PHONY: gen-proto
gen-proto:
	mkdir -p bin/
	GOBIN=`pwd`/bin/ go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.3 \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.3
	GOBIN=`pwd`/bin/ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=`pwd`/bin/ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	mkdir -p third_party
	git clone https://github.com/googleapis/googleapis.git third_party/googleapis || (cd third_party/googleapis && git reset --hard origin/master && git pull)
	wget -N -O bin/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-linux-x86_64.zip
	cd bin; rm -rf protoc; mkdir protoc; cd protoc; unzip ../protoc.zip
	mkdir -p ./generated

	./fixup-third-party.sh

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
		./apis/mockgcp/cloud/bigquery/v2/*.proto \
		./apis/mockgcp/cloud/resourcemanager/v1/*.proto \
		./apis/mockgcp/storage/v1/*.proto \
		./third_party/googleapis/mockgcp/api/apikeys/v2/*.proto \
		./third_party/googleapis/mockgcp/devtools/artifactregistry/v1/*.proto \
		./third_party/googleapis/mockgcp/iam/admin/v1/*.proto \
		./third_party/googleapis/mockgcp/logging/v2/*.proto \
		./third_party/googleapis/mockgcp/pubsub/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/annotation*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/data_item.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/dataset*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/deployed_index_ref.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/encryption_spec.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/index.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/index_service.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/io.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/operation.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/saved_query.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/tensorboard*.proto \
		./third_party/googleapis/mockgcp/cloud/alloydb/v1beta/*.proto \
		./third_party/googleapis/mockgcp/cloud/aiplatform/v1beta1/user_action_reference.proto \
		./third_party/googleapis/mockgcp/cloud/billing/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/certificatemanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/common/*.proto \
		./third_party/googleapis/mockgcp/cloud/compute/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkemulticloud/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/edgecontainer/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/edgenetwork/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/functions/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/networkservices/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/pubsublite/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/redis/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/resourcemanager/v3/*.proto \
		./third_party/googleapis/mockgcp/cloud/security/privateca/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/secretmanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/sql/v1beta4/*.proto \
		./third_party/googleapis/mockgcp/container/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/monitoring/dashboard/v1/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1beta1/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/configmanagement/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/metering/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkehub/v1beta/multiclusteringress/*.proto

.PHONY: generate-protos-from-openapi
generate-protos-from-openapi:
	wget -O bigquery-v2.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/bigquery/v2/bigquery-api.json
	mkdir -p apis/mockgcp/cloud/bigquery/v2/
	cd tools/gapic; go run . --proto-version=2 ../../bigquery-v2.json > ../../apis/mockgcp/cloud/bigquery/v2/api.proto

	wget -O servicenetworking-api.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/servicenetworking/v1/servicenetworking-api.json
	mkdir -p apis/mockgcp/cloud/servicenetworking/v1/
	cd tools/gapic; go run . ../../servicenetworking-api.json > ../../apis/mockgcp/cloud/servicenetworking/v1/servicenetworking.proto

	wget -O storage-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/main/storage/v1/storage-api.json
	mkdir -p apis/mockgcp/storage/v1/
	cd tools/gapic; go run . --proto-version=2 ../../storage-v1.json > ../../apis/mockgcp/storage/v1/service.proto
