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

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
REPO_ROOT="${SCRIPT_DIR}/../../.."
source "${REPO_ROOT}/dev/tools/goimports.sh"

cd ${REPO_ROOT}/dev/tools/controllerbuilder

go run . generate-types \
  --service google.cloud.alloydb.v1beta \
  --api-version alloydb.cnrm.cloud.google.com/v1beta1 \
  --resource AlloyDBCluster:Cluster \
  --resource AlloyDBInstance:Instance \
  --resource AlloyDBBackup:Backup \
  --resource AlloyDBUser:User

go run . generate-mapper --service google.cloud.alloydb.v1beta --api-version alloydb.cnrm.cloud.google.com/v1beta1

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/alloydb/
