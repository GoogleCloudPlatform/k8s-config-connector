#!/bin/bash
# Copyright 2025 Google LLC
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

pwd

export SERVICE=notebooks
export HTTP_HOST=notebooks.googleapis.com
export PROTO_SERVICE=google.cloud.notebooks.v1.NotebookService
export PROTO_MESSAGE=google.cloud.notebooks.v1.Instance
export RESOURCE=instance
export PROMPT_DIR=experiments/conductor/cmd/runner/scripts/mockgcp
export LOG_DIR=logs

mkdir -p ${LOG_DIR}

if [[ ! -e mock${SERVICE}/service.go ]]; then
controllerbuilder prompt --src-dir . --proto-dir ./.build/third_party/googleapis/ <<EOF > mockgcp/mock${SERVICE}/service.go
// +tool:mockgcp-service
// http.host: ${HTTP_HOST}
// proto.service: ${PROTO_SERVICE}
EOF
fi


if [[ ! -e mock${SERVICE}/${RESOURCE}.go ]]; then
controllerbuilder prompt --src-dir . --proto-dir ./.build/third_party/googleapis/ <<EOF > mockgcp/mock${SERVICE}/${RESOURCE}.go
// +tool:mockgcp-support
// proto.service: ${PROTO_SERVICE}
// proto.message: ${PROTO_MESSAGE}
EOF
fi

cat ${PROMPT_DIR}/03-add-service.prompt | \
    envsubst '$SERVICE' > ${LOG_DIR}/03-add-service.prompt

codebot --prompt=${LOG_DIR}/03-add-service.prompt


git status
# We add some files by name to force an error if not generated
git add mockgcp/mock${SERVICE}/service.go
git add mockgcp/mock${SERVICE}/${RESOURCE}.go
git add mock_http_roundtrip.go