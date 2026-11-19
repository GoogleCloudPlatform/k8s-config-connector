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

# We compile BlueGreenDeployment using the googleapis pin containing google/cloud/sql/v1/cloud_sql_blue_green_deployments.proto.
PROTO_SHA="e7513c828b8db536fa12fc5493d541f323aedf90"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

# --- v1alpha1 ---
${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.sql.v1 \
  --api-version sql.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource SQLBlueGreenDeployment:BlueGreenDeployment \
  --prepopulate-resource

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.sql.v1 \
  --api-version sql.cnrm.cloud.google.com/v1alpha1 \
  --resource SQLBlueGreenDeployment:BlueGreenDeployment

# --- v1beta1 ---
${CONTROLLERBUILDER} generate-types \
  --service google.cloud.sql.v1beta4 \
  --api-version sql.cnrm.cloud.google.com/v1beta1  \
  --resource SQLInstance:DatabaseInstance \
  --resource SQLSSLCert:SslCert \
  --skip-scaffold-files \
  --include-skipped-output \
  --prune-unused-types=false

${CONTROLLERBUILDER} generate-mapper \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.sql.v1beta4,google.cloud.sql.v1 \
  --api-version sql.cnrm.cloud.google.com/v1beta1 \
  --multiversion \
  --include-skipped-output \
  --prepopulate-resource

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/sql" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/sql/
fi
