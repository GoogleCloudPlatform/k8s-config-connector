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

# We need a newer googleapis to get DataplexDataProduct (data_products.proto)
PROTO_SHA="69ca7ae2e66cd5623cafecc00971ef5397a4b258"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

# Unset SKIP_GENERATE_PROTOS so this specific script fetches the newer proto
OLD_SKIP_GENERATE_PROTOS="${SKIP_GENERATE_PROTOS:-}"
unset SKIP_GENERATE_PROTOS

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

# Restore SKIP_GENERATE_PROTOS
if [[ -n "${OLD_SKIP_GENERATE_PROTOS}" ]]; then
  export SKIP_GENERATE_PROTOS="${OLD_SKIP_GENERATE_PROTOS}"
fi

go run . generate-types \
    --service google.cloud.dataplex.v1 \
    --api-version "dataplex.cnrm.cloud.google.com/v1alpha1" \
    --resource DataplexLake:Lake \
    --resource DataplexZone:Zone \
    --resource DataplexTask:Task \
    --resource DataplexEntryGroup:EntryGroup \
    --resource DataplexEntryType:EntryType \
    --resource DataplexDataTaxonomy:DataTaxonomy \
    --resource DataplexAspectType:AspectType \
    --resource DataplexDataScan:DataScan \
    --resource DataplexMetadataJob:MetadataJob \
    --resource DataplexDataProduct:DataProduct \
    --proto-source-path ${PROTO_OUT}

# Handled recursive self-referential fields by defining AspectType_MetadataTemplate manually in dataplexaspecttype_types.go

go run . generate-mapper \
    --service google.cloud.dataplex.v1 \
    --api-version "dataplex.cnrm.cloud.google.com/v1alpha1" \
    --proto-source-path ${PROTO_OUT}

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/dataplex/
