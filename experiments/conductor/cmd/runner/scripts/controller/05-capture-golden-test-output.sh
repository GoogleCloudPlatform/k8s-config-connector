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

if [[ -z "${WORKDIR}" ]]; then
  echo "WORKDIR is required"
  exit 1
fi

if [[ -z "${BRANCH_NAME}" ]]; then
  echo "BRANCH_NAME is required"
  exit 1
fi

cd ${WORKDIR}

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

hack/compare-mock fixtures/${CRD_KIND,,}-minimal || true


git status
git add .
git commit -m "${CRD_KIND}: Capture golden output"


echo "Done"