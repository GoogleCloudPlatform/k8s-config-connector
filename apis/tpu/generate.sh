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


${CONTROLLERBUILDER} generate-types     --service google.cloud.tpu.v2     --api-version tpu.cnrm.cloud.google.com/v1alpha1     --resource TPUVirtualMachine:Node

${CONTROLLERBUILDER} generate-types     --service google.cloud.tpu.v2     --api-version tpu.cnrm.cloud.google.com/v1alpha1     --resource TPUQueuedResource:QueuedResource

${CONTROLLERBUILDER} generate-mapper     --service google.cloud.tpu.v2     --api-version tpu.cnrm.cloud.google.com/v1alpha1

# Add +kubebuilder:validation:XPreserveUnknownFields to empty structs in types.generated.go to avoid openapi schema errors
sed -i 's/type QueuedResource_Spot struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResource_Spot struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"
sed -i 's/type QueuedResourceState_AcceptedData struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResourceState_AcceptedData struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"
sed -i 's/type QueuedResourceState_ActiveData struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResourceState_ActiveData struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"
sed -i 's/type QueuedResourceState_CreatingData struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResourceState_CreatingData struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"
sed -i 's/type QueuedResourceState_DeletingData struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResourceState_DeletingData struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"
sed -i 's/type QueuedResourceState_ProvisioningData struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResourceState_ProvisioningData struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"
sed -i 's/type QueuedResourceState_SuspendedData struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResourceState_SuspendedData struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"
sed -i 's/type QueuedResourceState_SuspendingData struct {/\/\/ +kubebuilder:validation:XPreserveUnknownFields\ntype QueuedResourceState_SuspendingData struct {/g' "${REPO_ROOT}/apis/tpu/v1alpha1/types.generated.go"

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/tpu" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/tpu/
fi
