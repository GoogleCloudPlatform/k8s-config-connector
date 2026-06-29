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

# We need a newer googleapis to get google.cloud.dialogflow.v2.Tool
# Cache buster to force fresh HEAD commit SHA for GHA
PROTO_SHA="e9195b36c9896e43fc9fa38793b66b5cf1d59b50"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

# Unset SKIP_GENERATE_PROTOS so this specific script fetches the newer proto
OLD_SKIP_GENERATE_PROTOS="${SKIP_GENERATE_PROTOS:-}"
unset SKIP_GENERATE_PROTOS

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

# Restore SKIP_GENERATE_PROTOS
if [[ -n "${OLD_SKIP_GENERATE_PROTOS}" ]]; then
  export SKIP_GENERATE_PROTOS="${OLD_SKIP_GENERATE_PROTOS}"
fi

# Generate types for Dialogflow CX v3 service (SecuritySettings)
go run . generate-types \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource DialogflowSecuritySettings:SecuritySettings

mv ${REPO_ROOT}/apis/dialogflow/v1alpha1/types.generated.go ${REPO_ROOT}/apis/dialogflow/v1alpha1/securitysettings_types.generated.go

# Generate types for all Dialogflow V2 resources (KnowledgeBase, Generator, and Tool)
go run . generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.dialogflow.v2 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --resource DialogflowKnowledgeBase:KnowledgeBase \
  --resource DialogflowGenerator:Generator \
  --resource DialogflowConversationDataset:ConversationDataset \
  --resource DialogflowTool:Tool

# Inject missing import for apiextensionsv1 into generated types
sed -i 's/^package v1alpha1/package v1alpha1\n\nimport (\n\tapiextensionsv1 "k8s.io\/apiextensions-apiserver\/pkg\/apis\/apiextensions\/v1"\n)/' "${REPO_ROOT}/apis/dialogflow/v1alpha1/types.generated.go"

# Generate mapper for Dialogflow CX v3 service
go run . generate-mapper \
  --service google.cloud.dialogflow.cx.v3 \
  --api-version dialogflow.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/dialogflow/
