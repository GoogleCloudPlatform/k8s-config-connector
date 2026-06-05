#!/bin/bash
# Copyright 2025 Google LLC
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

go run . generate-types --service google.cloud.discoveryengine.v1 --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1 \
  --resource DiscoveryEngineControl:Control \
  --resource DiscoveryEngineDataStore:DataStore \
  --resource DiscoveryEngineEngine:Engine \
  --resource DiscoveryEngineIdentityMappingStore:IdentityMappingStore \
  --resource DiscoveryEngineTargetSite:TargetSite \
  --resource DiscoveryEngineConversation:Conversation \
  --resource DiscoveryEngineSession:Session
mv ../../../apis/discoveryengine/v1alpha1/types.generated.go ../../../apis/discoveryengine/v1alpha1/v1_types.generated.go

go run . generate-types --service google.cloud.discoveryengine.v1beta --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1 \
  --resource DiscoveryEngineSampleQuerySet:SampleQuerySet
mv ../../../apis/discoveryengine/v1alpha1/types.generated.go ../../../apis/discoveryengine/v1alpha1/v1beta_types.generated.go

go run . generate-mapper --service google.cloud.discoveryengine.v1 --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1
mv ../../../pkg/controller/direct/discoveryengine/mapper.generated.go ../../../pkg/controller/direct/discoveryengine/v1_mapper.generated.go

go run . generate-mapper --service google.cloud.discoveryengine.v1beta --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1
mv ../../../pkg/controller/direct/discoveryengine/mapper.generated.go ../../../pkg/controller/direct/discoveryengine/v1beta_mapper.generated.go


cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/discoveryengine/
