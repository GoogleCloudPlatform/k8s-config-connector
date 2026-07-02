---
name: generate-sh-checker
description: Checks apis/ subdirectories for missing generate.sh and helps create PRs to add them, following the pattern in #7293. Use this when you want to ensure all KCC resources have a standardized generation script.
---

# Generate.sh Checker

This skill helps maintain the `generate.sh` pattern across all `apis/` subdirectories in the Config Connector codebase.

## Workflow

1.  **Scan for missing scripts**: Find subdirectories in `apis/` (usually `v1beta1` or `v1alpha1`) that do not have a `generate.sh` file.
    ```bash
    find apis -maxdepth 2 -type d \( -name "v1beta1" -o -name "v1alpha1" \) | while read dir; do if [ ! -f "$dir/generate.sh" ] && ls "$dir"/*_type*.go >/dev/null 2>&1; then echo "$dir"; fi; done
    ```
    *(Note: `apis/refs` is a special folder and does not correspond to a GCP service. Since it lacks `*_types.go` files, the above command naturally skips it, which is correct.)*

2.  **Gather Resource Information**: For each identified directory, read `api_types.go` and `groupversion_info.go` to extract:
    -   `PROTO_SERVICE`: Look for `// +kcc:spec:proto=` or `// +kcc:proto=` markers in `api_types.go`. If `api_types.go` does not exist (TF-based controllers), look for the proto in `.build/third_party/googleapis/google/cloud/` or similar.
    -   `GROUP`: Look for `// +groupName=` in `groupversion_info.go`, or `var <Kind>GVK = schema.GroupVersionKind{Group: ...}` in `<service>_types.go`.
    -   `VERSION`: The directory name (e.g., `v1beta1`).
    -   `RESOURCE_MAPPINGS`: Mapping of `Kind:ProtoMessage` from `// +kcc:spec:proto=` markers, or determine the `ProtoMessage` by reading the proto files.
    -   `SERVICE_NAME`: The parent directory name in `apis/` (e.g., `apigateway`).

    *Note: If the directory does not contain any `*_type*.go` file (e.g., it only contains reference types like `service_reference.go`), there are no types or mappers to generate. In this case, `generate.sh` is not required.*

3.  **Create generate.sh**: Create a `generate.sh` file in the directory. Ensure the year in the copyright header is current (2026).
    
    Template for Direct Controllers:
    ```bash
    #!/bin/bash
    # Copyright 2026 Google LLC
    #
    # Licensed under the Apache License, Version 2.0 (the "License");
    # you may not use this file except in compliance with the License.
    # You may obtain a copy of the License at
    #
    #      http://www.apache.org/licenses/LICENSE-2.0
    #
    # Unless required by applicable law or agreed to in writing, software
    # distributed under the License is distributed on an "AS IS" BASIS,
    # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    # See the License for the specific language governing permissions and
    # limitations under the License.

    set -o errexit
    set -o nounset
    set -o pipefail

    REPO_ROOT="$(git rev-parse --show-toplevel)"
    source "${REPO_ROOT}/dev/tools/goimports.sh"
    cd ${REPO_ROOT}/dev/tools/controllerbuilder

    ./generate-proto.sh

    go run . generate-types \
      --service <PROTO_SERVICE> \
      --api-version <GROUP>/<VERSION> \
      --include-skipped-output \
      --resource <KIND1>:<PROTO_MESSAGE1> \
      [--resource <KIND2>:<PROTO_MESSAGE2> ...]

    go run . generate-mapper \
      --service <PROTO_SERVICE> \
      --api-version <GROUP>/<VERSION> \
      --include-skipped-output

    cd ${REPO_ROOT}
    dev/tasks/generate-crds

    go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/<SERVICE_NAME>/
    ```

    **Template for TF-based Controllers (where `api_types.go` does not exist)**:
    When a direct controller does not yet exist for a resource, the generated types could conflict with the existing TF types in `pkg/clients/generated/apis`. Use this pattern:
    ```bash
    #!/bin/bash
    # ... copyright ...
    set -o errexit
    set -o nounset
    set -o pipefail

    REPO_ROOT="$(git rev-parse --show-toplevel)"
    source "${REPO_ROOT}/dev/tools/goimports.sh"
    cd ${REPO_ROOT}/dev/tools/controllerbuilder

    ./generate-proto.sh

    go run . generate-types \
      --service <PROTO_SERVICE> \
      --api-version <GROUP>/<VERSION> \
      --resource <KIND1>:<PROTO_MESSAGE1> \
      --skip-scaffold-files

    # NOTYET - not yet using proto
    # go run . generate-mapper \
    #   --service <PROTO_SERVICE> \
    #   --api-version <GROUP>/<VERSION>

    # NOTYET - not yet following full pattern
    rm -f ${REPO_ROOT}/apis/<SERVICE_NAME>/<VERSION>/groupversion_info.go

    cd ${REPO_ROOT}
    dev/tasks/generate-crds

    # NOTYET - not yet following full pattern
    # go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/<SERVICE_NAME>/
    ```
    *Note: Do not use `--include-skipped-output` for TF-based controllers to avoid block-comment syntax errors if the proto contains `/*` inside strings.*

4.  **Special Handling (Multi-version & Promotion/Consolidation)**:
    -   **File Naming**: `generate-types` expects the main types file to be named `<lowercase_proto_message_name>_types.go`. If the existing file has a different name (e.g., `cluster_types.go` instead of `attachedcluster_types.go`), rename it before running the generator.
    -   **Hand-written `types.generated.go`**: If a `types.generated.go` already exists but lacks the `// Code generated by ... DO NOT EDIT.` header, it was hand-written. Rename it to `types.go` to prevent it from being overwritten.
    -   **Pointer Types**: When preserving hand-written structs that correspond to proto messages, ensure their fields use pointers (e.g., `*string` with `,omitempty` instead of `string`) where the proto fields are optional. Otherwise, `mapper.generated.go` will fail to compile with type assignment errors (e.g., `cannot use direct.LazyPtr(in.GetName()) ... as string value in assignment`).
    -   **Multi-version resources**: When generating the same resource in `v1alpha1` and `v1beta1`, we should use `// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"` on the `v1beta1` resource to generate `v1alpha1` from `v1beta1`. Often this will mean we don't need a `v1alpha1` folder at all.
    -   **Multiple generate.sh scripts**: If both `v1alpha1` and `v1beta1` directories exist and both need a `generate.sh` script, only call `generate-mapper` in one of the version's `generate.sh` (typically `v1beta1/generate.sh` and not in `v1alpha1/generate.sh`). We still need `generate-types` in `v1alpha1/generate.sh`.
        -   If multiple API versions exist (e.g. `v1alpha1` and `v1beta1`) and contain `*_types.go` files, use the `--multiversion` flag when calling `generate-mapper` to prevent naming collisions (e.g., `NotebookInstanceSpec_FromProto` becoming `NotebookInstanceSpec_v1beta1_FromProto`). This is because `generate-mapper` scans the entire `apis/<SERVICE>` directory and will generate mappers for all versions it finds, causing naming collisions if the suffix is omitted. When using `--multiversion`, you will also need to update the direct controller (`pkg/controller/direct/<SERVICE>/*_controller.go` and `*_fuzzer.go`) to use the new version-suffixed mapper functions. If there are custom manual mapper functions in `*_mappings.go` or `*_mappers.go`, rename them to match the new version suffix (e.g., `_v1beta1_FromProto` or `_v1alpha1_FromProto`) so `generate-mapper` recognizes them and skips generating duplicates.
    -   **Promotion/Consolidation**: If a `v1beta1` directory is being updated and a `v1alpha1` directory exists for the same service, check if `v1alpha1` should be consolidated.
    -   Following #7293, this involves:
        -   Removing the `v1alpha1` directory entirely if all resources are in `v1beta1`. Or, if we generate the same resource in `v1alpha1` and `v1beta1`, we should use `// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"` in the `v1beta1` type to generate `v1alpha1` from `v1beta1`. Often this will mean we don't need a `v1alpha1` folder at all.
        -   Ensuring `v1beta1` has `// +kubebuilder:storageversion`.
    -   If the existing types files are not named `<lowercasekind>_types.go` (e.g., `cluster_types.go` instead of `workstationcluster_types.go`), rename them using `git mv` to match the expected pattern *before* running `generate.sh`. Otherwise, `controllerbuilder generate-types` will fail to find the existing types and create duplicate files.
    -   **Multiple API Versions**: If a service has multiple API versions (e.g. `v1alpha1` and `v1beta1`) that map to the *same* proto service and cannot be consolidated (e.g. due to cross-references), `generate-mapper` must be called with `--multiversion`. Otherwise, identical mapper functions will be generated twice and cause compile errors. When using `--multiversion`, check existing handwritten controller code to ensure they call the updated multi-versioned mapper functions (e.g. `KMSImportJobSpec_v1beta1_FromProto`).
    -   **Different Proto Packages/Versions**: If a resource (e.g. an alpha/beta resource like `ComputeFutureReservation`) belongs to a different proto package/version (e.g. `google.cloud.compute.v1beta`) than the default `--service` package (e.g. `google.cloud.compute.v1`), you can specify its fully-qualified name in the `--resource` flag (e.g. `--resource ComputeFutureReservation:google.cloud.compute.v1beta.FutureReservation`). This allows `generate-types` to locate the correct message descriptors across different packages when compiling multiple API levels.

5.  **Execute and Verify**:
    -   Make `generate.sh` executable: `chmod +x apis/<SERVICE>/<VERSION>/generate.sh`.
    -   Run it: `./apis/<SERVICE>/<VERSION>/generate.sh`.
    -   Verify that `types.generated.go` is created in the API directory.
    -   Verify that `pkg/controller/direct/<SERVICE>/mapper.generated.go` is updated (if applicable).
    -   Verify that CRDs in `config/crds/resources/` are updated.
    -   **Ref fields with acronyms**: If you encounter `// MISSING: [Acronym]...` (like `MISSING: KMSKey` or `MISSING: CAPool`) in the generated `mapper.generated.go`, it might be because the `generate-mapper` tool expects the field in the KRM struct to use the fully capitalized acronym (e.g. `KMSKeyRef` instead of `KmsKeyRef`, `CAPoolRef` instead of `CaPoolRef`). Rename the Go struct field to match the acronym (this won't break the yaml if the `json` tag is unchanged), and update any references in `mapper.go` or `[service]_controller.go`. The generator should then automatically map the `Ref` field properly.

6.  **Commit and PR**: Create a branch, commit the changes, and propose a PR with a descriptive title like `chore: apis/<SERVICE> should follow generate.sh pattern`.

## Troubleshooting

See `notes.md` for troubleshooting uncommon edge cases.

-   **Deepcopy-gen errors (`invalid slice element type: invalid type`)**: This typically happens if `generate-types` outputs a struct name with a different capitalization than what is currently manually written in the `*_types.go` file (e.g. `PSCConfig` vs `PscConfig`). To fix, rename the type and all its usages in the `*_types.go` and `pkg/controller/direct/<SERVICE>/mapper.go` files to match the generated capitalization, then run `./generate.sh` again.
