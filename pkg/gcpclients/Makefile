GOOGLEAPI_VERSION?=$(shell grep https://github.com/googleapis/googleapis ../../apis/git.versions | awk '{print $$2}' )
GOLANG_X_TOOLS_VERSION?=$(shell grep -E '^\s*golang.org/x/tools\s+' ../../go.mod | awk '{print $$2}')

.PHONY: generate-grpc-for-google-protos
generate-grpc-for-google-protos: tools
	mkdir -p .build/
	git clone https://github.com/googleapis/googleapis.git .build/googleapis || (cd .build/googleapis && git fetch --all && git reset --hard $(GOOGLEAPI_VERSION))
	mkdir -p ./generated

	PATH=.build/bin/:${PATH} .build/bin/protoc/bin/protoc \
		-I ./.build/googleapis \
		--experimental_allow_proto3_optional \
		--go_out ./generated \
		--go_opt paths=source_relative \
		--go-grpc_out ./generated \
		--go-grpc_opt paths=source_relative \
		./.build/googleapis/google/cloud/tpu/v2/*.proto

	PATH=.build/bin/:${PATH} .build/bin/protoc/bin/protoc \
		-I ./.build/googleapis \
		--experimental_allow_proto3_optional \
		--go_out ./generated \
		--go_opt paths=source_relative \
		--go-grpc_out ./generated \
		--go-grpc_opt paths=source_relative \
		./.build/googleapis/google/cloud/sql/v1beta4/*.proto

	# Fetch newer googleapis for networksecurity
	mkdir -p .build/googleapis-newer
	git clone https://github.com/googleapis/googleapis.git .build/googleapis-newer || (cd .build/googleapis-newer && git fetch --all && git reset --hard cdc919ff596e263f2cc55a9780d2f74633da1ced)

	find .build/googleapis-newer/google/cloud/networksecurity/v1 -name "*.proto" -exec sed -i "s|option go_package = \".*\";|option go_package = \"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/networksecurity/v1;networksecuritypb\";|g" {} +

	PATH=.build/bin/:$${PATH} .build/bin/protoc/bin/protoc \
		-I ./.build/googleapis-newer \
		--experimental_allow_proto3_optional \
		--go_out ./generated \
		--go_opt paths=source_relative \
		--go-grpc_out ./generated \
		--go-grpc_opt paths=source_relative \
		./.build/googleapis-newer/google/cloud/networksecurity/v1/*.proto

	go run -mod=readonly golang.org/x/tools/cmd/goimports@$(GOLANG_X_TOOLS_VERSION) -w generated

.PHONY: tools
tools:
	mkdir -p .build/bin/
	GOBIN=`pwd`/.build/bin/ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=`pwd`/.build/bin/ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	wget -N -O .build/bin/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-linux-x86_64.zip
	cd .build/bin; rm -rf protoc; mkdir protoc; cd protoc; python3 -m zipfile -e ../protoc.zip .; chmod +x bin/protoc
