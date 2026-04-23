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

# Setting GOPATH changes behaviour of k8s codegen tools
unset GOPATH

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "${REPO_ROOT}"

# Clear out old clients to prevent issues with deepcopy-gen
rm -rf "${REPO_ROOT}/pkg/clients/generated/client"

# Generate strong-typed definitions for existing CRDs
echo "Generating go types"
go run ./scripts/generate-go-crd-clients

# Generate deepcopy etc
echo "Generating deepcopy for go types"
go generate ./pkg/clients/...

# Generate the clients
echo "Generating clients"
cd "${REPO_ROOT}/pkg/clients/generated/"

# Extract API & version names
API_DIRS=(apis/*/*/)
API_VERSIONS=()
for DIR in "${API_DIRS[@]}";
do
  API_VERSIONS+=($(echo ${DIR} | cut -d'/' -f 2,3))
done

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
# deepcopy-gen and client-gen use 'v1alpha1' by default when there is no shadowing, but our types use 'k8sv1alpha1'
# We need to handle both the case where it used the 'v1alpha1' alias and the case where it used no alias.

echo "Fixing up imports"
# 1. We apply the import alias fix to all files in the generated directory.
# We use a more robust regex that handles leading whitespace and optional aliases correctly.
find "${REPO_ROOT}/pkg/clients/generated" -name "*.go" -exec sed -E -i 's|^([[:space:]]*)([a-zA-Z0-9]*[[:space:]]+)?"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"|\1k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"|g' {} +

# 2. In files that now use the k8sv1alpha1 alias, we replace all v1alpha1. usages with k8sv1alpha1.
# This handles all codegen tools and avoids accidental replacements in other v1alpha1 packages.
# We use a while loop to be more robust than xargs grep | xargs sed.
find "${REPO_ROOT}/pkg/clients/generated" -name "*.go" | while read -r file; do
  if grep -q 'k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"' "$file"; then
    sed -E -i 's/\bv1alpha1\./k8sv1alpha1\./g' "$file"
  fi
done

# Cleanup any misplaced directories that might have been created by generators
rm -rf "${REPO_ROOT}/generated"
rm -rf "${REPO_ROOT}/github.com"
rm -rf "${REPO_ROOT}/pkg/clients/github.com"

echo "Applying gofmt"
cd ${REPO_ROOT}
make fmt # Fix up the formatting and headers
