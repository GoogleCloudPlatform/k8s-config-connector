#!/bin/bash
# Copyright 2022 Google LLC
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

# Extract API names and versions from the directory structure
# This assumes the structure is pkg/clients/generated/apis/<api>/<version>
# and we want to generate clients for all of them.
# We'll use a list of unique API/version pairs.
cd "${REPO_ROOT}/pkg/clients/generated/"
API_DIRS=(apis/*/*/)
API_VERSIONS=()
for DIR in "${API_DIRS[@]}"; do
  # DIR is e.g. apis/compute/v1beta1/
  API_VERSIONS+=($(echo ${DIR} | cut -d'/' -f 2,3))
done

# Setting GOPATH changes behaviour of k8s codegen tools
ORIGINAL_GOPATH="${GOPATH:-}"
unset GOPATH

# Generate the types
echo "Generating go types"
cd "${REPO_ROOT}"
go run ./scripts/generate-go-crd-clients

# Generate deepcopy etc
echo "Generating deepcopy for go types"
# We need to run go generate on the directory containing generators.go
cd "${REPO_ROOT}"
go generate ./pkg/clients/...

# Clear out old clients to prevent issues with deepcopy-gen
rm -rf "${REPO_ROOT}/pkg/clients/generated/client"

# Generate the clients
echo "Generating clients"

# Join API/version names into a comma-separated list
printf -v JOINED '%s,' "${API_VERSIONS[@]:1}"
JOINED="${JOINED}${API_VERSIONS[0]}"

cd ${REPO_ROOT}
go run k8s.io/code-generator/cmd/client-gen@v0.29.0 \
  --clientset-name versioned \
  --input-base github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis \
  --input ${JOINED} \
  --output-base ./ \
  --output-package github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset \
  --trim-path-prefix github.com/GoogleCloudPlatform/k8s-config-connector \
  -h ${REPO_ROOT}/hack/boilerplate_client_alpha.go.txt

# Fix up codegen using the wrong alias for k8sv1alpha1 in v1beta1 packages
echo "Fixing up imports"
# 1. We apply the import alias fix to all files in the generated directory.
# We use a robust regex that handles leading whitespace, optional aliases (including those with underscores),
# and ensures a canonical tab-indented aliased import.
find "${REPO_ROOT}/pkg/clients/generated" -name "*.go" -exec sed -E -i 's|^[[:space:]]*([a-zA-Z0-9_]*[[:space:]]+)?"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"|\tk8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"|g' {} +

# 2. In files that now use the k8sv1alpha1 alias, we replace all v1alpha1. usages with k8sv1alpha1.
# This handles all codegen tools and avoids accidental replacements in other v1alpha1 packages.
find "${REPO_ROOT}/pkg/clients/generated" -name "*.go" | while read -r file; do
  if grep -q "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1" "$file"; then
    # Aggressive replacement to catch any missed cases, then fix any accidental double-prefixing.
    # We use a loop to handle multiple occurrences on the same line if necessary (though sed /g handles it).
    sed -i 's/v1alpha1\./k8sv1alpha1\./g' "$file"
    sed -i 's/k8sk8sv1alpha1\./k8sv1alpha1\./g' "$file"
  fi
done

# Cleanup any misplaced directories that might have been created by generators
rm -rf "${REPO_ROOT}/generated"
rm -rf "${REPO_ROOT}/github.com"
rm -rf "${REPO_ROOT}/pkg/clients/github.com"

echo "Applying gofmt"
if [[ -n "${ORIGINAL_GOPATH}" ]]; then
  export GOPATH="${ORIGINAL_GOPATH}"
fi
cd ${REPO_ROOT}
make fmt # Fix up the formatting and headers

# 3. Final cleanup: Remove any non-aliased k8s/v1alpha1 imports added by goimports
# that are redundant because the aliased version is already there.
# This is a safety measure against goimports disagreement between environments.
find "${REPO_ROOT}/pkg/clients/generated" -name "*.go" -exec sed -i '/"github.com\/GoogleCloudPlatform\/k8s-config-connector\/pkg\/clients\/generated\/apis\/k8s\/v1alpha1"/ { /k8sv1alpha1/ !d }' {} +

# Run fmt again just in case the previous cleanup removed a line from an import block
make fmt
