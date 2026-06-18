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

export REPO_ROOT="$(git rev-parse --show-toplevel)"
source "${REPO_ROOT}/dev/tools/goimports.sh"

cd ${REPO_ROOT}/dev/tools/controllerbuilder

./generate-proto.sh

go run . generate-types \
  --service google.cloud.aiplatform.v1 \
  --api-version aiplatform.cnrm.cloud.google.com/v1alpha1 \
  --resource VertexAISpecialistPool:SpecialistPool

# Revert types.generated.go to avoid deleting types of other resources in the same package
git checkout HEAD -- "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"

go run . generate-mapper \
  --service google.cloud.aiplatform.v1 \
  --api-version aiplatform.cnrm.cloud.google.com/v1alpha1 \
  --api-dir "${REPO_ROOT}/apis/aiplatform/v1alpha1" \
  --api-go-package-path "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"

cd ${REPO_ROOT}

# Remove mappers for resources in other KRM groups or problematic ones.
python3 -c '
import sys
import os

path = "pkg/controller/direct/aiplatform/mapper.generated.go"
if not os.path.exists(path):
    sys.exit(0)

content = open(path).read()
lines = content.splitlines()
out = []
skip = False
# Problematic functions that refer to other KRM groups or missing types
problematic = ["ActiveLearningConfig", "SampleConfig", "TrainingConfig", "DataLabelingJob", "Money", "Value_FromProto", "Value_ToProto", "ListValue"]

for line in lines:
    if "func " in line and any(p in line for p in problematic):
        skip = True
        continue
    if skip:
        if line == "}":
            skip = False
        continue
    out.append(line)

# Also remove the import if it exists
final_out = []
for line in out:
    if "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1" in line:
        continue
    final_out.append(line)

open(path, "w").write("\n".join(final_out) + "\n")
'

dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w apis/aiplatform/v1alpha1/
go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/aiplatform/
