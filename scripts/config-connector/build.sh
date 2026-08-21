#!/usr/bin/env bash
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

# builds the config-connector binary for a single system and architecture
# to build for a system / arch other than the default define the ${GOOS} and ${GOARCH} variables
REPO_ROOT=$(git rev-parse --show-toplevel)
cd ${REPO_ROOT}

VERSION=${VERSION:-dev}
BASE_OUTPUT_DIR=bin/config-connector
# if goarch OR goos is not set, grab it from the go env
export GOOS=${GOOS:-$(go env GOOS)}
export GOARCH=${GOARCH:-$(go env GOARCH)}
# create a target directory for the given system & architecture
OUTPUT_DIR=${BASE_OUTPUT_DIR}/${GOOS}/${GOARCH}
mkdir -p ${OUTPUT_DIR}
# run the build
echo "Building config-connector for ${GOOS}/${GOARCH}"
LDFLAGS="-X \"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd.version=${VERSION}\""
OUTPUT_PATH=${OUTPUT_DIR}/config-connector
if [[ ${GOOS} == "windows" ]]; then
  OUTPUT_PATH="${OUTPUT_PATH}.exe"
fi
go build -ldflags "${LDFLAGS}" -o ${OUTPUT_PATH} github.com/GoogleCloudPlatform/k8s-config-connector/cmd/config-connector
