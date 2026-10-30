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
PROTO_SHA="e7513c828b8db536fa12fc5493d541f323aedf90"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"
./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}
${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.bigquery.analyticshub.v1 \
  --api-version bigqueryanalyticshub.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource BigQueryAnalyticsHubQueryTemplate:QueryTemplate

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.bigquery.analyticshub.v1 \
  --api-version bigqueryanalyticshub.cnrm.cloud.google.com/v1alpha1 \
  --resource BigQueryAnalyticsHubQueryTemplate:QueryTemplate

${CONTROLLERBUILDER} generate-controller \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.bigquery.analyticshub.v1 \
  --api-version bigqueryanalyticshub.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource BigQueryAnalyticsHubQueryTemplate:QueryTemplate

${CONTROLLERBUILDER} generate-fuzzer \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.bigquery.analyticshub.v1 \
  --api-version bigqueryanalyticshub.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource BigQueryAnalyticsHubQueryTemplate:QueryTemplate

cd ${REPO_ROOT}
if [[ -z "${SKIP_GENERATE_CRDS:-}" ]]; then
  dev/tasks/generate-crds
fi

if [ -d "${REPO_ROOT}/pkg/controller/direct/bigqueryanalyticshub" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/bigqueryanalyticshub/
fi
