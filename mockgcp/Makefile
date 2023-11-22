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
		./apis/mockgcp/cloud/resourcemanager/v1/*.proto \
		./third_party/googleapis/mockgcp/iam/admin/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/billing/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/certificatemanager/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/compute/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/gkemulticloud/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/edgecontainer/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/edgenetwork/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/functions/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/networkservices/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/resourcemanager/v3/*.proto \
		./third_party/googleapis/mockgcp/cloud/security/privateca/v1/*.proto \
		./third_party/googleapis/mockgcp/cloud/secretmanager/v1/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1/*.proto \
		./third_party/googleapis/mockgcp/api/serviceusage/v1beta1/*.proto
