GOOGLEAPI_VERSION?=$(shell grep https://github.com/googleapis/googleapis ../../apis/git.versions | awk '{print $$2}' )
GOLANG_X_TOOLS_VERSION?=$(shell grep -E '^\s*golang.org/x/tools\s+' ../../go.mod | awk '{print $$2}')

.PHONY: generate-grpc-for-google-protos
generate-grpc-for-google-protos: tools
	mkdir -p .build/
	@if [ ! -d .build/googleapis ]; then \
		git clone https://github.com/googleapis/googleapis.git .build/googleapis; \
	fi
	cd .build/googleapis && git fetch --all && git reset --hard $(GOOGLEAPI_VERSION) && git clean -fd
	mkdir -p ./generated

	PATH=.build/bin/:${PATH} .build/bin/protoc/bin/protoc \
		-I ./.build/googleapis \
		--experimental_allow_proto3_optional \
		--go_out ./generated \
		--go_opt paths=source_relative \
		--go-grpc_out ./generated \
		--go-grpc_opt paths=source_relative \
		./.build/googleapis/google/cloud/tpu/v2/*.proto

	printf "// The server CA mode.\noptional string server_ca_mode = 11;\n\n// The server CA pool.\noptional string server_ca_pool = 12;\n\n// Custom Subject Alternative Names.\nrepeated string custom_subject_alternative_names = 13;\n" | go -C ../../mockgcp/tools/patch-proto run . --file $$(pwd)/.build/googleapis/google/cloud/sql/v1beta4/cloud_sql_resources.proto --message IpConfiguration --mode append

	sed -i '/POSTGRES_16 = 272;/a \  // The database version is PostgreSQL 17.\n  POSTGRES_17 = 408;\n\n  // The database version is PostgreSQL 18.\n  POSTGRES_18 = 557;' ./.build/googleapis/google/cloud/sql/v1beta4/cloud_sql_resources.proto

	PATH=.build/bin/:${PATH} .build/bin/protoc/bin/protoc \
		-I ./.build/googleapis \
		--experimental_allow_proto3_optional \
		--go_out ./generated \
		--go_opt paths=source_relative \
		--go-grpc_out ./generated \
		--go-grpc_opt paths=source_relative \
		./.build/googleapis/google/cloud/sql/v1beta4/*.proto

	go run -mod=readonly golang.org/x/tools/cmd/goimports@$(GOLANG_X_TOOLS_VERSION) -w generated

.PHONY: tools
tools:
	mkdir -p .build/bin/
	GOBIN=`pwd`/.build/bin/ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=`pwd`/.build/bin/ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	wget -N -O .build/bin/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-linux-x86_64.zip
	cd .build/bin; rm -rf protoc; mkdir protoc; cd protoc; python3 -m zipfile -e ../protoc.zip .; chmod +x bin/protoc
