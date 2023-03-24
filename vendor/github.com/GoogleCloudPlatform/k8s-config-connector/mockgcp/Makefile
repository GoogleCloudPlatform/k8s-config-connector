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
	# May need protoc installed: apt install protobuf-compiler
	mkdir -p ./generated
	PATH=bin/:${PATH} protoc \
		-I ./third_party/googleapis \
		--grpc-gateway_out ./generated \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt standalone=true \
		--experimental_allow_proto3_optional \
		./third_party/googleapis/google/cloud/secretmanager/v1/service.proto \
		./third_party/googleapis/google/cloud/security/privateca/v1/*.proto \
