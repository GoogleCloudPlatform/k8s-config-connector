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

./generate-proto.sh

# Generate types for Dialogflow v2 service (KnowledgeBase, Generator)
go run . generate-types \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowKnowledgeBase:KnowledgeBase \
  --resource DialogflowGenerator:Generator \
  --resource DialogflowConversationDataset:ConversationDataset

# Generate types for Dialogflow CX v3 service (SecuritySettings)
go run . generate-types \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource DialogflowSecuritySettings:SecuritySettings

mv ${REPO_ROOT}/apis/dialogflow/v1alpha1/types.generated.go ${REPO_ROOT}/apis/dialogflow/v1alpha1/securitysettings_types.generated.go

# Generate types for Dialogflow v2 service again to restore types.generated.go for Dialogflow v2
go run . generate-types \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowKnowledgeBase:KnowledgeBase \
  --resource DialogflowGenerator:Generator \
  --resource DialogflowConversationDataset:ConversationDataset

# Generate mapper for Dialogflow CX v3 service
go run . generate-mapper \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/dialogflow/
