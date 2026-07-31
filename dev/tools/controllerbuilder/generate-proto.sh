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
# Extract the default git commit SHA version of the googleapis repository from apis/git.versions
DEFAULT_GOOGLE_API_VERSION=$(grep https://github.com/googleapis/googleapis ${REPO_ROOT}/apis/git.versions | awk '{print $2}')

# Take googleapi version as parameter, default to version from git.versions.
# Use "HEAD" to get the latest from remote.
GOOGLEAPI_VERSION=${1:-$DEFAULT_GOOGLE_API_VERSION}

# Take output path as parameter, default to .build/googleapis.pb
OUTPUT_PATH=${2:-"${REPO_ROOT}/.build/googleapis.pb"}


THIRD_PARTY="${REPO_ROOT}/.build/third_party"
mkdir -p "${THIRD_PARTY}/"

DEFAULT_GOOGLEAPI_DIR="${THIRD_PARTY}/googleapis"
if [ ! -d "${DEFAULT_GOOGLEAPI_DIR}" ]; then
    git clone --depth 1 https://github.com/googleapis/googleapis.git "${DEFAULT_GOOGLEAPI_DIR}"
fi

if [ "${GOOGLEAPI_VERSION}" == "HEAD" ]; then
    echo "Fetching latest googleapis for HEAD version"
    cd "${DEFAULT_GOOGLEAPI_DIR}"
    GOOGLEAPI_VERSION=$(git ls-remote https://github.com/googleapis/googleapis.git refs/heads/master | awk '{print $1}')
fi

if [ -n "${2:-}" ]; then
    # Explicitly provided output path, use it directly
    VERSIONED_OUTPUT_PATH="${OUTPUT_PATH}"
else
    # Default output path, version it with the SHA.
    # ${OUTPUT_PATH%.pb} strips the suffix '.pb' from the path, then we append the SHA.
    VERSIONED_OUTPUT_PATH="${OUTPUT_PATH%.pb}-${GOOGLEAPI_VERSION}.pb"
fi

if [[ "${SKIP_GENERATE_PROTOS:-0}" == "1" ]] && [ -f "${VERSIONED_OUTPUT_PATH}" ]; then
    echo "Skipping generate-proto.sh as requested by SKIP_GENERATE_PROTOS=1 and output file exists: ${VERSIONED_OUTPUT_PATH}"
    exit 0
fi

if [ "${GOOGLEAPI_VERSION}" == "${DEFAULT_GOOGLE_API_VERSION}" ]; then
    VERSION_DIR="${DEFAULT_GOOGLEAPI_DIR}"
else
    VERSION_DIR="${THIRD_PARTY}/googleapis-${GOOGLEAPI_VERSION}"
    if [ ! -d "${VERSION_DIR}" ]; then
        git clone --depth 1 https://github.com/googleapis/googleapis.git "${VERSION_DIR}"
    fi
fi

cd "${VERSION_DIR}"

if ! git cat-file -e ${GOOGLEAPI_VERSION}^{commit} 2> /dev/null; then
    echo "Fetching googleapis git objects to find version ${GOOGLEAPI_VERSION}"
    git fetch origin ${GOOGLEAPI_VERSION}
fi

git reset --hard ${GOOGLEAPI_VERSION}

# Overwrite config.proto with the updated version that has DeploymentGroup
mkdir -p google/cloud/config/v1
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/config/v1/config.proto google/cloud/config/v1/config.proto


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
    if [ "${VERSIONED_OUTPUT_PATH}" != "${OUTPUT_PATH}" ]; then
        cp "${VERSIONED_OUTPUT_PATH}" "${OUTPUT_PATH}"
    fi
    exit 0
fi


# Enable nullglob shell option so that unmatched glob patterns (e.g. if certain subdirectories don't contain any .proto files)
# expand to an empty string instead of the literal wildcard string, avoiding protoc errors.
shopt -s nullglob
PROTO_FILES=(
    ${REPO_ROOT}/mockgcp/apis/google/apps/cloudidentity/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/apigee/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/networkconnectivity/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/servicenetworking/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/google/cloud/binaryauthorization/*/*.proto
    ${VERSION_DIR}/google/*/*.proto
    ${VERSION_DIR}/google/analytics/*/*/*.proto
    ${VERSION_DIR}/google/privacy/dlp/v2/*.proto
    ${VERSION_DIR}/google/api/*.proto
    ${VERSION_DIR}/google/api/*/*/*.proto
    ${VERSION_DIR}/google/bigtable/*/*/*.proto
    ${VERSION_DIR}/google/cloud/bigquery/*/*.proto
    ${VERSION_DIR}/google/cloud/*/*/*.proto
    ${VERSION_DIR}/google/cloud/*/*/*/*.proto
    ${VERSION_DIR}/google/cloud/*/*/*/*/*.proto
    ${VERSION_DIR}/google/dataflow/*/*.proto
    ${VERSION_DIR}/google/firestore/*/*.proto
    ${VERSION_DIR}/google/firestore/*/*/*.proto
    ${VERSION_DIR}/google/iam/v1/*.proto
    ${VERSION_DIR}/google/iam/admin/v1/*.proto
    ${VERSION_DIR}/google/logging/v2/*.proto
    ${VERSION_DIR}/google/monitoring/v3/*.proto
    ${VERSION_DIR}/google/monitoring/metricsscope/v1/*.proto
    ${VERSION_DIR}/google/monitoring/dashboard/v1/*.proto
    ${VERSION_DIR}/google/devtools/cloudbuild/*/*.proto
    ${VERSION_DIR}/google/devtools/artifactregistry/*/*.proto
    ${VERSION_DIR}/google/devtools/testing/*/*.proto
    ${VERSION_DIR}/google/spanner/admin/instance/v1/*.proto
    ${VERSION_DIR}/google/spanner/admin/database/v1/*.proto
    ${VERSION_DIR}/google/storage/control/v2/*.proto
    ${VERSION_DIR}/google/storage/v1/*.proto
    ${VERSION_DIR}/google/pubsub/v1/*.proto
    ${VERSION_DIR}/google/maps/mapmanagement/*/*.proto
    ${VERSION_DIR}/google/cloud/memorystore/v1/*.proto
    ${VERSION_DIR}/google/container/*/*.proto
    ${VERSION_DIR}/google/privacy/dlp/v2/*.proto
    ${VERSION_DIR}/grafeas/v1/*.proto
)
# Disable nullglob shell option to restore the default shell globbing behavior
shopt -u nullglob

# Compile the protocols to a single descriptor set binary file (.pb).
# 2> >(...) redirects stderr into process substitution to filter out noisy "Import ... is unused" warnings.
protoc --include_imports --include_source_info \
    --experimental_allow_proto3_optional \
    -I ${VERSION_DIR}/ \
    -I ${REPO_ROOT}/mockgcp/apis \
    "${PROTO_FILES[@]}" \
    -o ${VERSIONED_OUTPUT_PATH} 2> >(grep -v "Import .* is unused" >&2)

if [ "${VERSIONED_OUTPUT_PATH}" != "${OUTPUT_PATH}" ]; then
    cp "${VERSIONED_OUTPUT_PATH}" "${OUTPUT_PATH}"
fi
