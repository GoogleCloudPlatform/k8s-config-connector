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

# Compile googleapis specifically for workloadmanager using the version that has it
# This avoids breaking other APIs by changing the global apis/git.versions
./generate-proto.sh df13ea51961ff225a825beab9ca112188de367a3

# Generate the KCC type structs from the GCP proto definitions
go run . generate-types \
  --service google.cloud.workloadmanager.v1 \
  --api-version workloadmanager.cnrm.cloud.google.com/v1alpha1  \
  --proto-source-path ${REPO_ROOT}/.build/googleapis-df13ea51961ff225a825beab9ca112188de367a3.pb \
  --resource WorkloadManagerEvaluation:Evaluation

# Change back to the repo root
cd ${REPO_ROOT}

# Generate the CRD YAML from the type structs
dev/tasks/generate-crds
