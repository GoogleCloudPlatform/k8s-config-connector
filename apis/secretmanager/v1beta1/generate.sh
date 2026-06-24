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

go run . generate-types \
    --service google.cloud.secretmanager.v1 \
    --api-version secretmanager.cnrm.cloud.google.com/v1beta1 \
    --resource SecretManagerSecret:Secret \
    --resource SecretManagerSecretVersion:SecretVersion

cd ${REPO_ROOT}

# HACK: Fix the slice bug in controllerbuilder for ReplicationStatus_UserManagedStatus_ReplicaStatusObservedState
sed -i 's/Replicas \[\]ReplicationStatus_UserManagedStatus_ReplicaStatus `json:"replicas,omitempty"`/Replicas \[\]ReplicationStatus_UserManagedStatus_ReplicaStatusObservedState `json:"replicas,omitempty"`/g' apis/secretmanager/v1beta1/types.generated.go

dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/secretmanager/
