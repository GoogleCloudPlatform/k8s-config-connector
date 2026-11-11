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

# Pin a googleapis SHA that contains the Template definition
PROTO_SHA="4fabd6e4b356307d319b55a86a3f200e9c52b594"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

if [[ -z "${SKIP_GENERATE_PROTOS:-}" ]]; then
  ./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}
fi

# Generate the KCC type structs from the GCP proto definitions
${CONTROLLERBUILDER} generate-types \
  --service google.cloud.parametermanager.v1 \
  --api-version parametermanager.cnrm.cloud.google.com/v1alpha1  \
  --resource ParameterManagerParameter:Parameter \
  --resource ParameterManagerParameterVersion:ParameterVersion \
  --proto-source-path ${PROTO_OUT}

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.parametermanager.v1 \
  --api-version parametermanager.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --prepopulate-resource \
  --resource ParameterManagerTemplate:Template \
  --proto-source-path ${PROTO_OUT}

${CONTROLLERBUILDER} generate-identity \
  --service google.cloud.parametermanager.v1 \
  --api-version parametermanager.cnrm.cloud.google.com/v1alpha1 \
  --resource ParameterManagerTemplate:Template \
  --proto-source-path ${PROTO_OUT}

# Generate the mapper functions that convert between the KCC structs and the GCP proto structs
${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.parametermanager.v1 \
  --api-version parametermanager.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --prepopulate-resource \
  --proto-source-path ${PROTO_OUT}

# Change back to the repo root
cd ${REPO_ROOT}

# Generate the CRD YAML from the type structs

# Format the generated Go code
if [[ -z "${SKIP_GENERATE_CRDS:-}" ]]; then
  dev/tasks/generate-crds
fi

if [ -d "${REPO_ROOT}/pkg/controller/direct/parametermanager" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/parametermanager/
fi
