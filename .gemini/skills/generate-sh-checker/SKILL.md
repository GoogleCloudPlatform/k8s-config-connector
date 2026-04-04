---
name: generate-sh-checker
description: Checks apis/ subdirectories for missing generate.sh and helps create PRs to add them, following the pattern in #7293. Use this when you want to ensure all KCC resources have a standardized generation script.
---

# Generate.sh Checker

This skill helps maintain the `generate.sh` pattern across all `apis/` subdirectories in the Config Connector codebase.

## Workflow

1.  **Scan for missing scripts**: Find subdirectories in `apis/` (usually `v1beta1` or `v1alpha1`) that do not have a `generate.sh` file.
    ```bash
    find apis -maxdepth 2 -type d \( -name "v1beta1" -o -name "v1alpha1" \) | while read dir; do [ ! -f "$dir/generate.sh" ] && echo "$dir"; done
    ```

2.  **Gather Resource Information**: For each identified directory, read `api_types.go` and `groupversion_info.go` to extract:
    -   `PROTO_SERVICE`: Look for `// +kcc:spec:proto=` or `// +kcc:proto=` markers in `api_types.go`.
    -   `GROUP`: Look for `// +groupName=` in `groupversion_info.go`.
    -   `VERSION`: The directory name (e.g., `v1beta1`).
    -   `RESOURCE_MAPPINGS`: Mapping of `Kind:ProtoMessage` from `// +kcc:spec:proto=` markers.
    -   `SERVICE_NAME`: The parent directory name in `apis/` (e.g., `apigateway`).

3.  **Create generate.sh**: Create a `generate.sh` file in the directory. Ensure the year in the copyright header is current (2026).
    
    Template:
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

4.  **Special Handling (Promotion/Consolidation)**:
    -   If a `v1beta1` directory is being updated and a `v1alpha1` directory exists for the same service, check if `v1alpha1` should be consolidated.
    -   Following #7293, this involves:
        -   Removing the `v1alpha1` directory.
        -   Adding `// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"` to the `v1beta1` `api_types.go` file (near the Kind struct).
        -   Ensuring `v1beta1` has `// +kubebuilder:storageversion`.

5.  **Execute and Verify**:
    -   Make `generate.sh` executable: `chmod +x apis/<SERVICE>/<VERSION>/generate.sh`.
    -   Run it: `./apis/<SERVICE>/<VERSION>/generate.sh`.
    -   Verify that `types.generated.go` is created in the API directory.
    -   Verify that `pkg/controller/direct/<SERVICE>/mapper.generated.go` is updated.
    -   Verify that CRDs in `config/crds/resources/` are updated.

6.  **Commit and PR**: Create a branch, commit the changes, and propose a PR with a descriptive title like `chore: apis/<SERVICE> should follow generate.sh pattern`.
