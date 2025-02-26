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

PROMPT=${SCRIPT_DIR}/01-generate-script.prompt

if [[ -z "${WORKDIR}" ]]; then
  echo "WORKDIR is required"
  exit 1
fi

if [[ -z "${BRANCH_NAME}" ]]; then
  echo "BRANCH_NAME is required"
  exit 1
fi

if [[ -z "${GCLOUD_COMMAND}" ]]; then
  echo "GCLOUD_COMMAND is required"
  exit 1
fi

if [[ -z "${LOG_DIR}" ]]; then
  echo "LOG_DIR is required"
  exit 1
fi

if [[ -z "${EXPECTED_PATH}" ]]; then
  echo "EXPECTED_PATH is required"
  exit 1
fi

mkdir -p ${LOG_DIR}
cat ${PROMPT} | \
    envsubst '$GCLOUD_COMMAND,$EXPECTED_PATH' > ${LOG_DIR}/prompt

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME} || git co -b ${BRANCH_NAME}

cd ${REPO_ROOT}/mockgcp

codebot --prompt=${LOG_DIR}/prompt | tee ${LOG_DIR}/codebot.log

git status
git add ${EXPECTED_PATH}/script.yaml
git add .
git commit -m "mockgcp: test script for ${GCLOUD_COMMAND}"

echo "Done"