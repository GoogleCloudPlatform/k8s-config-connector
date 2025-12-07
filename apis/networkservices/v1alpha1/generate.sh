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
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# ServiceBinding
go run . generate-types \
    --service google.cloud.networkservices.v1 \
    --api-version "networkservices.cnrm.cloud.google.com/v1alpha1" \
    --resource NetworkServicesServiceBinding:ServiceBinding

mv ${REPO_ROOT}/apis/networkservices/v1alpha1/types.generated.go ${REPO_ROOT}/apis/networkservices/v1alpha1/servicebinding_types.generated.go

go run . generate-mapper \
    --service google.cloud.networkservices.v1 \
    --api-version "networkservices.cnrm.cloud.google.com/v1alpha1"

mv ${REPO_ROOT}/pkg/controller/direct/networkservices/mapper.generated.go ${REPO_ROOT}/pkg/controller/direct/networkservices/servicebinding_mapper.generated.go

# EdgeCacheService
go run . generate-types \
    --config config/v1alpha1/edgecacheservice.yaml

mv ${REPO_ROOT}/apis/networkservices/v1alpha1/types.generated.go ${REPO_ROOT}/apis/networkservices/v1alpha1/edgecacheservice_types.generated.go

go run . generate-mapper \
    --config config/v1alpha1/edgecacheservice.yaml

mv ${REPO_ROOT}/pkg/controller/direct/networkservices/mapper.generated.go ${REPO_ROOT}/pkg/controller/direct/networkservices/edgecacheservice_mapper.generated.go

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg/controller/direct/networkservices/
