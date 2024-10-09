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
source ${REPO_ROOT}/scripts/shared-vars-public.sh
cd ${REPO_ROOT}
source ${REPO_ROOT}/scripts/fetch_ext_bins.sh && \
	fetch_tools && \
	setup_envs

cd ${REPO_ROOT}/
echo "Running e2e fixtures test with vcr replay mode..."
E2E_KUBE_TARGET=envtest \
RUN_E2E=1 \
E2E_GCP_TARGET=vcr \
VCR_MODE=replay \
  go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures 2>&1
