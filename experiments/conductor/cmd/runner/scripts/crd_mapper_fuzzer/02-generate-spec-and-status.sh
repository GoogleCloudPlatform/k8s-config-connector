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

# TODO: sanity check the branch
git checkout "${BRANCH_NAME}"

type_file="${REPO_ROOT}/apis/${SERVICE}/${CRD_VERSION}/$(eval "echo ${RESOURCE,,}_types.go")"

controllerbuilder prompt --src-dir "${REPO_ROOT}" --proto-dir "${REPO_ROOT}"/.build/third_party/googleapis/ <<EOF >> "${type_file}"
// +kcc:proto=${PROTO_SERVICE}.${PROTO_RESOURCE}
// crd.kind: ${CRD_KIND} 
EOF

git status
git add .
git commit -m "${CRD_KIND}: Update types from generated"

echo "Done"