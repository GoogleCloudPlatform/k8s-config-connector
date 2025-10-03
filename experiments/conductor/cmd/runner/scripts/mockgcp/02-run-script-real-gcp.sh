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

set -x

cd $(dirname "$0")
SCRIPT_DIR=`pwd`

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

if [[ -z "${RUN_TEST}" ]]; then
  echo "RUN_TEST is required"
  exit 1
fi

mkdir -p ${LOG_DIR}

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

git co master
git co ${BRANCH_NAME}

cd ${REPO_ROOT}/mockgcp

export WRITE_GOLDEN_OUTPUT=1
export E2E_GCP_TARGET=real

echo "Running test"
# HACK: refresh token
gcloud auth print-access-token > /dev/null
# We ignore test failures, because we expect the golden output to be volatile at this stage
(go test ./mockgcptests -v -run TestScripts/${RUN_TEST} || true) | tee ${LOG_DIR}/test-realgcp.log

git status
# We add some files by name to force an error if not generated
git add ${RUN_TEST}/_http.log
git add .
git commit -m "mockgcp: Capture golden output for ${RUN_TEST}"

echo "Done"