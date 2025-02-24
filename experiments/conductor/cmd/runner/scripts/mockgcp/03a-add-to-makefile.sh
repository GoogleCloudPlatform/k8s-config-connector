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


if [[ -z "${PROTO_PACKAGE}" ]]; then
  echo "PROTO_PACKAGE is required"
  exit 1
fi

mkdir -p ${LOG_DIR}

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

cd ${REPO_ROOT}/mockgcp

cat ${PROMPT_DIR}/03-add-to-makefile.prompt | \
    envsubst '$PROTO_PACKAGE' > ${LOG_DIR}/03-add-to-makefile.prompt

codebot --prompt=${LOG_DIR}/03-add-to-makefile.prompt


git status
git add Makefile
git add .
git commit -m "autogen: update Makefile for ${PROTO_PACKAGE}"

make generate-grpc-for-google-protos

git add .
git commit -m "autogen: updated GRPC/proto dependencies for ${PROTO_PACKAGE}"

echo "Done"





