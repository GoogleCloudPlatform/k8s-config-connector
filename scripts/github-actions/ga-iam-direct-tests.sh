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

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/

APIGROUPS=(
    "iam.cnrm.cloud.google.com"
)

export KCC_USE_DIRECT_RECONCILERS="IAMPartialPolicy"

# Running direct tests
export ONLY_TEST_APIGROUPS=$(IFS=,; echo "${APIGROUPS[*]}")
${REPO_ROOT}/dev/tasks/run-e2e

# Running direct and tracker tests
export KCC_RECONCILE_FLAG_GATE=USE_DEPENDENCY_TRACKER
unset ONLY_TEST_APIGROUPS

echo "Running fixture and sample iam tests with the gcptracker..."
export RUN_TESTS='TestAllInSeries/samples/iampartialpolicy'
${REPO_ROOT}/dev/tasks/run-e2e