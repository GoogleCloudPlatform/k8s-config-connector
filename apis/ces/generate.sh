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

PROTO_SHA="1526e545e9d26f23b9c5d0f04af17297def8d045"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"
./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.ces.v1beta \
  --api-version ces.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource CESApp:App \
  --resource CESDeployment:Deployment \
  --resource CESEvaluationDataset:EvaluationDataset \
  --resource CESEvaluationExpectation:EvaluationExpectation \
  --resource CESExample:Example \
  --resource CESGuardrail:Guardrail \
  --resource CESScheduledEvaluationRun:ScheduledEvaluationRun \
  --resource CESTool:Tool \
  --prepopulate-spec \
  --emit-required-from-proto \
  --emit-plural-acronyms \
  --emit-message-maps \
  --place-server-set-fields \
  --detect-output-only-in-comments \
  --emit-parent-refs \
  --emit-sibling-refs \
  --emit-reference-hints \
  --emit-source-links

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.ces.v1beta \
  --api-version ces.cnrm.cloud.google.com/v1alpha1 \
  --resource CESDeployment:Deployment \
  --resource CESEvaluationDataset:EvaluationDataset \
  --resource CESEvaluationExpectation:EvaluationExpectation \
  --resource CESExample:Example \
  --resource CESGuardrail:Guardrail \
  --resource CESScheduledEvaluationRun:ScheduledEvaluationRun \
  --resource CESTool:Tool


cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/ces" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/ces/
fi
