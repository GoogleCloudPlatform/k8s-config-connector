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

# --- v1alpha1 ---

${CONTROLLERBUILDER} generate-types \
    --service google.cloud.bigquery.biglake.v1 \
    --api-version "bigquerybiglake.cnrm.cloud.google.com/v1alpha1" \
    --include-skipped-output \
    --resource BigLakeCatalog:Catalog \
    --resource BigLakeDatabase:Database

${CONTROLLERBUILDER} generate-types \
    --service google.cloud.biglake.v1 \
    --api-version "bigquerybiglake.cnrm.cloud.google.com/v1alpha1" \
    --include-skipped-output \
    --resource BigLakeIcebergCatalog:IcebergCatalog \
    --resource BigLakeIcebergNamespace:IcebergNamespace \
    --prepopulate-resource

${CONTROLLERBUILDER} generate-identity \
    --service google.cloud.biglake.v1 \
    --api-version "bigquerybiglake.cnrm.cloud.google.com/v1alpha1" \
    --resource BigLakeIcebergCatalog:IcebergCatalog \
    --resource BigLakeIcebergNamespace:IcebergNamespace \
    --pattern BigLakeIcebergCatalog=projects/{project}/catalogs/{catalog} \
    --pattern BigLakeIcebergNamespace=projects/{project}/catalogs/{catalog}/namespaces/{namespace}

${CONTROLLERBUILDER} generate-types \
    --service google.cloud.biglake.hive.v1 \
    --api-version "bigquerybiglake.cnrm.cloud.google.com/v1alpha1" \
    --include-skipped-output \
    --resource BigLakeHiveCatalog:HiveCatalog \
    --resource BigLakeHiveDatabase:HiveDatabase \
    --prepopulate-resource

${CONTROLLERBUILDER} generate-identity \
    --service google.cloud.biglake.hive.v1 \
    --api-version "bigquerybiglake.cnrm.cloud.google.com/v1alpha1" \
    --resource BigLakeHiveCatalog:HiveCatalog \
    --resource BigLakeHiveDatabase:HiveDatabase \
    --pattern BigLakeHiveCatalog=projects/{project}/catalogs/{catalog} \
    --pattern BigLakeHiveDatabase=projects/{project}/catalogs/{catalog}/databases/{database}

# --- v1beta1 ---

${CONTROLLERBUILDER} generate-types \
    --service google.cloud.bigquery.biglake.v1 \
    --api-version "bigquerybiglake.cnrm.cloud.google.com/v1beta1" \
    --resource BigLakeTable:Table

${CONTROLLERBUILDER} generate-mapper \
    --multiversion \
    --service google.cloud.bigquery.biglake.v1,google.cloud.biglake.v1,google.cloud.biglake.hive.v1 \
    --api-version "bigquerybiglake.cnrm.cloud.google.com/v1beta1" \
    --prepopulate-resource

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/bigquerybiglake" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/bigquerybiglake/
fi
