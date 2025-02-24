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

cd $(dirname "$0")
SCRIPT_DIR=`pwd`

PROMPT_DIR=${SCRIPT_DIR}

if [[ -z "${WORKDIR}" ]]; then
  echo "WORKDIR is required"
  exit 1
fi

if [[ -z "${BRANCH_NAME}" ]]; then
  echo "BRANCH_NAME is required"
  exit 1
fi

if [[ -z "${LOG_DIR}" ]]; then
  echo "LOG_DIR is required"
  exit 1
fi

if [[ -z "${SERVICE}" ]]; then
  echo "SERVICE is required"
  exit 1
fi
if [[ -z "${RESOURCE}" ]]; then
  echo "RESOURCE is required"
  exit 1
fi
if [[ -z "${PROTO_SERVICE}" ]]; then
  echo "PROTO_SERVICE is required"
  exit 1
fi
if [[ -z "${PROTO_MESSAGE}" ]]; then
  echo "PROTO_SERVICE is required"
  exit 1
fi
if [[ -z "${HTTP_HOST}" ]]; then
  echo "HTTP_HOST is required"
  exit 1
fi


mkdir -p ${LOG_DIR}

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

cd ${REPO_ROOT}/mockgcp

mkdir -p mock${SERVICE}

# TODO: Or just regenerate?

if [[ ! -e mock${SERVICE}/service.go ]]; then

controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > mock${SERVICE}/service.go
// +tool:mockgcp-service
// http.host: ${HTTP_HOST}
// proto.service: ${PROTO_SERVICE}
EOF

fi


controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > mock${SERVICE}/${RESOURCE}.go
// +tool:mockgcp-support
// proto.service: ${PROTO_SERVICE}
// proto.message: ${PROTO_MESSAGE}
EOF

cat ${PROMPT_DIR}/03-add-service.prompt | \
    envsubst '$SERVICE' > ${LOG_DIR}/03-add-service.prompt

codebot --prompt=${LOG_DIR}/03-add-service.prompt


git status
# We add some files by name to force an error if not generated
git add mock${SERVICE}/service.go
git add mock${SERVICE}/${RESOURCE}.go
git add mock_http_roundtrip.go
git add .
git commit -m "autogen: generate mockgcp for ${SERVICE} ${RESOURCE}"

echo "Done"





