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

# We share the version with mockgcp, which is maybe a boundary violation, but is convenient.
# (It would be confusing if these were out of sync!)
DEFAULT_GOOGLE_API_VERSION=$(grep https://github.com/googleapis/googleapis ${REPO_ROOT}/apis/git.versions | awk '{print $2}')

# Take googleapi version as parameter, default to version from git.versions.
# Use "HEAD" to get the latest from remote.
GOOGLEAPI_VERSION=${1:-$DEFAULT_GOOGLE_API_VERSION}

# Take output path as parameter, default to .build/googleapis.pb
OUTPUT_PATH=${2:-"${REPO_ROOT}/.build/googleapis.pb"}


THIRD_PARTY=${REPO_ROOT}/.build/third_party
mkdir -p ${THIRD_PARTY}/
cd ${THIRD_PARTY}

if [ ! -d "googleapis" ]; then
    git clone https://github.com/googleapis/googleapis.git
fi

if [ "${GOOGLEAPI_VERSION}" == "HEAD" ]; then
    echo "Fetching latest googleapis for HEAD version"
    # Get latest version from https://github.com/googleapis/googleapis.git
    GOOGLEAPI_VERSION=$(git ls-remote https://github.com/googleapis/googleapis.git refs/heads/master | awk '{print $1}')
fi

VERSIONED_OUTPUT_PATH="${OUTPUT_PATH%.pb}-${GOOGLEAPI_VERSION}.pb"

cd googleapis

# Fetch only if we don't have the SHA locally
if ! git cat-file -e ${GOOGLEAPI_VERSION}^{commit} 2> /dev/null; then
    echo "Fetching googleapis git objects to find version ${GOOGLEAPI_VERSION}"
    git fetch origin ${GOOGLEAPI_VERSION}
fi

# Reset to the desired version
git reset --hard ${GOOGLEAPI_VERSION}


if (which protoc); then
    echo "Found protoc version $(protoc --version)"
else
    echo "Installing protoc"
    if [ "$(uname)" == "Darwin" ]; then
      brew install protobuf
    else
      echo "apt update..."
      sudo apt update
      echo "apt install..."
      sudo apt install -y protobuf-compiler
    fi
fi


if [ -f "${VERSIONED_OUTPUT_PATH}" ]; then
    echo "Using cached googleapis pb file at ${VERSIONED_OUTPUT_PATH}"
    cp "${VERSIONED_OUTPUT_PATH}" "${OUTPUT_PATH}"
    exit 0
fi

protoc --include_imports --include_source_info \
    --experimental_allow_proto3_optional \
    -I ${THIRD_PARTY}/googleapis/ \
    -I ${REPO_ROOT}/mockgcp/apis \
    ${REPO_ROOT}/mockgcp/apis/google/apps/cloudidentity/*/*.proto \
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/apigee/*/*.proto \
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/networkconnectivity/*/*.proto \
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/servicenetworking/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/analytics/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/api/*.proto \
    ${THIRD_PARTY}/googleapis/google/api/*.proto \
    ${THIRD_PARTY}/googleapis/google/api/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/bigtable/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/bigquery/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/*/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/*/*/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/dataflow/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/firestore/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/firestore/*/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/iam/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/logging/v2/*.proto \
    ${THIRD_PARTY}/googleapis/google/monitoring/v3/*.proto \
    ${THIRD_PARTY}/googleapis/google/monitoring/dashboard/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/devtools/cloudbuild/*/*.proto \
    ${THIRD_PARTY}/googleapis/google/spanner/admin/instance/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/spanner/admin/database/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/storage/control/v2/*.proto \
    ${THIRD_PARTY}/googleapis/google/pubsub/v1/*.proto \
    ${THIRD_PARTY}/googleapis/google/cloud/memorystore/v1/*.proto \
    -o ${VERSIONED_OUTPUT_PATH} 2> >(grep -v "Import .* is unused" >&2)

cp "${VERSIONED_OUTPUT_PATH}" "${OUTPUT_PATH}"