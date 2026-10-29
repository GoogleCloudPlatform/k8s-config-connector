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
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowKnowledgeBase:KnowledgeBase \
  --resource DialogflowGenerator:Generator \
  --resource DialogflowConversationDataset:ConversationDataset

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource DialogflowSecuritySettings:SecuritySettings
mv ${REPO_ROOT}/apis/dialogflow/v1alpha1/types.generated.go ${REPO_ROOT}/apis/dialogflow/v1alpha1/securitysettings_types.generated.go

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.dialogflow.v2beta1 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowSipTrunk:SipTrunk
mv ${REPO_ROOT}/apis/dialogflow/v1alpha1/types.generated.go ${REPO_ROOT}/apis/dialogflow/v1alpha1/siptrunk_types.generated.go

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowKnowledgeBase:KnowledgeBase \
  --resource DialogflowGenerator:Generator \
  --resource DialogflowConversationDataset:ConversationDataset

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource DialogflowConversationModel:ConversationModel

${CONTROLLERBUILDER} generate-identity \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowConversationModel:ConversationModel

${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${REPO_ROOT}/.build/googleapis-e7513c828b8db536fa12fc5493d541f323aedf90.pb \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource DialogflowEnvironment:Environment \
  --resource DialogflowPlaybook:Playbook \
  --resource DialogflowTransitionRouteGroup:TransitionRouteGroup \
  --resource DialogflowFlow:Flow \
  --resource DialogflowPage:Page

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${REPO_ROOT}/.build/googleapis-e7513c828b8db536fa12fc5493d541f323aedf90.pb \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowEnvironment:Environment \
  --resource DialogflowPlaybook:Playbook \
  --resource DialogflowTransitionRouteGroup:TransitionRouteGroup \
  --resource DialogflowFlow:Flow \
  --resource DialogflowPage:Page

${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --prepopulate-resource

${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.dialogflow.v2beta1 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --output-dir "${REPO_ROOT}/pkg/controller/direct/dialogflow/siptrunk"
mv "${REPO_ROOT}/pkg/controller/direct/dialogflow/siptrunk/dialogflow/mapper.generated.go" "${REPO_ROOT}/pkg/controller/direct/dialogflow/siptrunk/mapper.generated.go" || true
rmdir "${REPO_ROOT}/pkg/controller/direct/dialogflow/siptrunk/dialogflow/" || true
sed -i 's/package dialogflow/package siptrunk/g' "${REPO_ROOT}/pkg/controller/direct/dialogflow/siptrunk/mapper.generated.go"

${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --output-dir "${REPO_ROOT}/pkg/controller/direct/dialogflow/generator"
mv "${REPO_ROOT}/pkg/controller/direct/dialogflow/generator/dialogflow/mapper.generated.go" "${REPO_ROOT}/pkg/controller/direct/dialogflow/generator/mapper.generated.go" || true
rmdir "${REPO_ROOT}/pkg/controller/direct/dialogflow/generator/dialogflow/" || true
sed -i 's/package dialogflow/package generator/g' "${REPO_ROOT}/pkg/controller/direct/dialogflow/generator/mapper.generated.go"

${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --output-dir "${REPO_ROOT}/pkg/controller/direct/dialogflow/conversationdataset"
mv "${REPO_ROOT}/pkg/controller/direct/dialogflow/conversationdataset/dialogflow/mapper.generated.go" "${REPO_ROOT}/pkg/controller/direct/dialogflow/conversationdataset/mapper.generated.go" || true
rmdir "${REPO_ROOT}/pkg/controller/direct/dialogflow/conversationdataset/dialogflow/" || true
sed -i 's/package dialogflow/package conversationdataset/g' "${REPO_ROOT}/pkg/controller/direct/dialogflow/conversationdataset/mapper.generated.go"

cd ${REPO_ROOT}
if [[ -z "${SKIP_GENERATE_CRDS:-}" ]]; then
  dev/tasks/generate-crds
fi

if [ -d "${REPO_ROOT}/pkg/controller/direct/dialogflow" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/dialogflow/
fi
