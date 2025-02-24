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

if [[ -z "${GCLOUD_COMMAND}" ]]; then
  echo "GCLOUD_COMMAND is required"
  exit 1
fi

export WORKDIR=~/kccai/work1/
export BRANCH_NAME=gcloud_${SERVICE}_${RESOURCE}
export LOG_DIR=/tmp/conductor/${BRANCH_NAME}

export EXPECTED_PATH=mock${SERVICE}/testdata/${RESOURCE}/crud
./01-generate-script.sh

export RUN_TEST=${EXPECTED_PATH}
./02-run-script-real-gcp.sh 

./03a-add-to-makefile.sh

./03-implement-mocks.sh

./04-run-script-mockgcp.sh
