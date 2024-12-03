#!/bin/bash
# Copyright 2024 Google LLC
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
cd ${REPO_ROOT}/dev/tools/controllerbuilder

THIRD_PARTY=${REPO_ROOT}/.build/third_party
mkdir -p ${THIRD_PARTY}/

# We share the version with mockgcp, which is maybe a boundary violation, but is convenient.
# (It would be confusing if these were out of sync!)
GOOGLEAPI_VERSION=$(grep https://github.com/googleapis/googleapis ${REPO_ROOT}/mockgcp/git.versions | awk '{print $2}' )

cd ${REPO_ROOT}/.build/third_party
git clone https://github.com/googleapis/googleapis.git ${THIRD_PARTY}/googleapis || (cd ${THIRD_PARTY}/googleapis && git reset --hard ${GOOGLEAPI_VERSION})

if (which protoc); then
    echo "Found protoc version $(protoc --version)"
else
    echo "Installing protoc"
    if [ "$(uname)" == "Darwin" ]; then
      brew install protobuf
    else
      sudo apt install -y protobuf-compiler
    fi
fi


protoc --include_imports --include_source_info \
    -I ${THIRD_PARTY}/googleapis/ \
    -I ${REPO_ROOT}/mockgcp/apis \
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/networkconnectivity/*/*.proto \
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/apigee/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/api/*.proto \
    ${THIRD_PARTY}/googleapis/google/api/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/bigtable/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/bigquery/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/*/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/dataflow/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/firestore/admin/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/iam/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/logging/v2/*.proto \
    ${THIRD_PARTY}/googleapis/google/monitoring/v3/*.proto \
    ${THIRD_PARTY}/googleapis/google/monitoring/dashboard/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/devtools/cloudbuild/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/spanner/admin/instance/v1/*.proto \
    -o ${REPO_ROOT}/.build/googleapis.pb
