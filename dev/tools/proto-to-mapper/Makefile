.PHONY: generate-code
generate-code: generate-pb
	go run . -apis ../../../apis/ --api-packages github.com/GoogleCloudPlatform/apis

.PHONY: generate-pb
generate-pb: install-protoc-linux
	mkdir -p third_party
	mkdir -p build
	git clone https://github.com/googleapis/googleapis.git third_party/googleapis || (cd third_party/googleapis && git fetch --all && git reset --hard origin/master && git pull)
	protoc --include_imports --include_source_info \
		-I ./third_party/googleapis/ \
		-I ../../../mockgcp/apis \
		../../../mockgcp/apis/mockgcp/cloud/networkconnectivity/*/*.proto \
		../../../mockgcp/apis/google/apps/cloudidentity/groups/*/*.proto \
		./third_party/googleapis/google/api/*.proto \
		./third_party/googleapis/google/api/*/*/*.proto \
		./third_party/googleapis/google/bigtable/*/*/*.proto \
		./third_party/googleapis/google/cloud/bigquery/*/*.proto \
		./third_party/googleapis/google/cloud/*/*/*.proto \
		./third_party/googleapis/google/cloud/*/*/*/*.proto \
		./third_party/googleapis/google/dataflow/*/*.proto \
		./third_party/googleapis/google/firestore/admin/v1/*.proto \
		./third_party/googleapis/google/iam/v1/*.proto \
		./third_party/googleapis/google/logging/v2/*.proto \
		./third_party/googleapis/google/monitoring/v3/*.proto \
		./third_party/googleapis/google/monitoring/dashboard/v1/*.proto \
		./third_party/googleapis/google/devtools/cloudbuild/*/*.proto \
		./third_party/googleapis/google/spanner/admin/instance/v1/*.proto \
		-o build/googleapis.pb

install-protoc-linux:
	which protoc || sudo apt install -y protobuf-compiler

install-protoc-mac:
	sudo brew install protobuf
