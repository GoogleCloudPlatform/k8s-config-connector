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
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# Generate the KCC type structs from the GCP proto definitions
go run . generate-types \
  --service google.cloud.parametermanager.v1 \
  --api-version parametermanager.cnrm.cloud.google.com/v1alpha1  \
  --resource ParameterManagerParameter:Parameter

# Generate the mapper functions that convert between the KCC structs and the GCP proto structs
go run . generate-mapper \
  --service google.cloud.parametermanager.v1 \
  --api-version parametermanager.cnrm.cloud.google.com/v1alpha1

# Change back to the repo root
cd ${REPO_ROOT}

# Generate the CRD YAML from the type structs
dev/tasks/generate-crds

# Format the generated Go code
go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg/controller/direct/parametermanager/