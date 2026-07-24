#!/bin/bash
# Copyright 2026 Google LLC
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
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# We must compile GKEHub using a newer googleapis pin than the global default in apis/git.versions.
# Therefore, we override SKIP_GENERATE_PROTOS and specify a custom output path to compile the newer version.
if [[ -n ${SKIP_GENERATE_PROTOS:-} ]]; then
  unset SKIP_GENERATE_PROTOS
fi
./generate-proto.sh "0fcabfc28371e7bab8107402eb06ad58134ee383" "${REPO_ROOT}/.build/googleapis-gkehub.pb"

go run . generate-types \
  --proto-source-path "${REPO_ROOT}/.build/googleapis-gkehub.pb" \
  --service google.cloud.gkehub.v1 \
  --api-version gkehub.cnrm.cloud.google.com/v1alpha1 \
  --resource GKEHubFleet:Fleet

cd ${REPO_ROOT}
dev/tasks/generate-crds
