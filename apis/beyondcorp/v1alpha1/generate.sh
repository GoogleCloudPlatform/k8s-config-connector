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
cd "${REPO_ROOT}/dev/tools/controllerbuilder"

./generate-proto.sh

go run . generate-types \
  --service google.cloud.beyondcorp.clientconnectorservices.v1 \
  --api-version beyondcorp.cnrm.cloud.google.com/v1alpha1 \
  --resource BeyondCorpClientConnectorService:ClientConnectorService

mv "${REPO_ROOT}/apis/beyondcorp/v1alpha1/types.generated.go" "${REPO_ROOT}/apis/beyondcorp/v1alpha1/clientconnectorservice_types.generated.go" || true

go run . generate-types \
    --service google.cloud.beyondcorp.clientgateways.v1 \
    --api-version beyondcorp.cnrm.cloud.google.com/v1alpha1 \
    --resource BeyondCorpClientGateway:ClientGateway

mv "${REPO_ROOT}/apis/beyondcorp/v1alpha1/types.generated.go" "${REPO_ROOT}/apis/beyondcorp/v1alpha1/clientgateway_types.generated.go" || true

go run . generate-mapper \
  --service google.cloud.beyondcorp.clientconnectorservices.v1 \
  --api-version beyondcorp.cnrm.cloud.google.com/v1alpha1 \
  --output-dir "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientconnectorservice"

mv "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientconnectorservice/beyondcorp/mapper.generated.go" "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientconnectorservice/mapper.generated.go" || true
rmdir "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientconnectorservice/beyondcorp/" || true
sed -i.bak 's/package beyondcorp/package beyondcorpclientconnectorservice/g' "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientconnectorservice/mapper.generated.go" && rm -f "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientconnectorservice/mapper.generated.go.bak"

go run . generate-mapper \
    --service google.cloud.beyondcorp.clientgateways.v1 \
    --api-version beyondcorp.cnrm.cloud.google.com/v1alpha1 \
    --output-dir "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientgateway"

mv "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientgateway/beyondcorp/mapper.generated.go" "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientgateway/mapper.generated.go" || true
rmdir "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientgateway/beyondcorp/" || true
sed -i.bak 's/package beyondcorp/package beyondcorpclientgateway/g' "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientgateway/mapper.generated.go" && rm -f "${REPO_ROOT}/pkg/controller/direct/beyondcorp/beyondcorpclientgateway/mapper.generated.go.bak"

cd "${REPO_ROOT}"
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w apis/beyondcorp/v1alpha1/
go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/beyondcorp/
