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
cd "${REPO_ROOT}"

# Generate strong-typed definitions for existing CRDs
echo "Generating go types"
go run ./scripts/generate-go-crd-clients
make fmt # Fix up the formatting and headers

# HACK: Some of the kubernetes generation tools still run better in GOPATH
rm -rf ${REPO_ROOT}/.build/go/src/github.com/GoogleCloudPlatform/k8s-config-connector
mkdir -p "${REPO_ROOT}/.build/go/src/github.com/GoogleCloudPlatform"
ln -sf "${REPO_ROOT}" "${REPO_ROOT}/.build/go/src/github.com/GoogleCloudPlatform/k8s-config-connector" 
export GOPATH="${REPO_ROOT}/.build/go"

GOPATH_REPO_ROOT="${GOPATH}/src/github.com/GoogleCloudPlatform/k8s-config-connector"

# Generate deepcopy etc
echo "Generating deepcopy for go types"
cd "${GOPATH_REPO_ROOT}"
go generate ./pkg/clients/generated/...

# Generate the clients
echo "Generating clients"
cd "${GOPATH_REPO_ROOT}/pkg/clients/generated/"

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

go run ${GOPATH_REPO_ROOT}/scripts/client-gen/main.go --clientset-name versioned --input-base github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis --input ${JOINED} --output-base ../ --output-package github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset -h ${GOPATH_REPO_ROOT}/hack/boilerplate_client_alpha.go.txt

# Clients are generated in a temp github.com/ folder where we pull the
# generated files out into a cleared pkg/client/ folder
rm -rf client/
mv ../github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client .
rm -rf ../cnrm.googlesource.com
