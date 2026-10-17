# Skill: Generate GCP Client from Proto

This skill guides an automated agent through generating a Google Cloud Platform (GCP) Go client (protobuf and gRPC) directly from Google APIs protobuf files (`googleapis`).

## Prerequisites
- Config Connector codebase with `pkg/gcpclients/Makefile`.
- `protoc` compiler and plugins setup defined in the `Makefile`.
- Pinned `googleapis` version in `apis/git.versions`.

## Steps

1. **Verify Google APIs Version**
   - Check the `apis/git.versions` file to ensure the `https://github.com/googleapis/googleapis` repository is pinned to the correct SHA.

2. **Modify `pkg/gcpclients/Makefile`**
   - Locate the target `generate-grpc-for-google-protos` in `pkg/gcpclients/Makefile`.
   - If the resource's proto is not in the main branch of `googleapis` (like Cloud DNS `v1`), add a step to fetch and check out the folder from its branch:
     ```makefile
     	cd .build/googleapis && git fetch origin initial-test-for-dns && git checkout origin/initial-test-for-dns -- google/cloud/dns
     ```
   - Add a new block for the resource you want to generate. For example, for DNS v1, add:
     ```makefile
     	PATH=.build/bin/:$${PATH} .build/bin/protoc/bin/protoc \
     		-I ./.build/googleapis \
     		--experimental_allow_proto3_optional \
     		--go_out ./generated \
     		--go_opt paths=source_relative \
     		--go-grpc_out ./generated \
     		--go-grpc_opt paths=source_relative \
     		./.build/googleapis/google/cloud/dns/v1/*.proto
     ```
   - Note the `$${PATH}` in makefile syntax to avoid escaping issues if needed, or follow the existing pattern in the Makefile (which uses `${PATH}`).

3. **Generate the Client**
   - Navigate to `pkg/gcpclients/` and run the generation rule:
     ```bash
     make generate-grpc-for-google-protos
     ```
   - This command will:
     - Clone/fetch the `googleapis` repository at the pinned SHA.
     - Build and install the required versions of `protoc-gen-go` and `protoc-gen-go-grpc`.
     - Generate Go protobuf and gRPC code under `pkg/gcpclients/generated/`.
     - Run `goimports` to format the generated Go files and resolve imports.

4. **Verify the Generated Files**
   - Check the `pkg/gcpclients/generated/` folder.
   - For a resource under `google/cloud/dns/v1`, files should be generated in `pkg/gcpclients/generated/google/cloud/dns/v1/`.
   - Verify that package declarations and import paths in the generated files are correct.

5. **Format and Lint**
   - Run `make fmt` at the root of the project to ensure all code (including new generated files) is properly formatted.
   - Run `go vet ./...` to ensure there are no simple compilation or import issues.
