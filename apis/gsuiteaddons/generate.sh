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

CONTROLLERBUILDER="${CONTROLLERBUILDER:-}"
if [[ -z "${CONTROLLERBUILDER}" ]]; then
  if [[ -x "${REPO_ROOT}/bin/controllerbuilder" ]]; then
    CONTROLLERBUILDER="${REPO_ROOT}/bin/controllerbuilder"
  else
    CONTROLLERBUILDER="go run ${REPO_ROOT}/dev/tools/controllerbuilder"
  fi
fi
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder
./generate-proto.sh

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.gsuiteaddons.v1 \
  --api-version gsuiteaddons.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource GSuiteAddonsDeployment:Deployment

${CONTROLLERBUILDER} generate-identity \
  --service google.cloud.gsuiteaddons.v1 \
  --api-version gsuiteaddons.cnrm.cloud.google.com/v1alpha1 \
  --resource GSuiteAddonsDeployment:Deployment

${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.gsuiteaddons.v1,google.apps.script.type,google.apps.script.type.calendar,google.apps.script.type.docs,google.apps.script.type.drive,google.apps.script.type.gmail,google.apps.script.type.sheets,google.apps.script.type.slides,google.protobuf \
  --api-version gsuiteaddons.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource


cd ${REPO_ROOT}
if [[ -z "${SKIP_GENERATE_CRDS:-}" ]]; then
  dev/tasks/generate-crds
fi

if [ -d "${REPO_ROOT}/pkg/controller/direct/gsuiteaddons" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/gsuiteaddons/
fi
