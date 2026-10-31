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

${CONTROLLERBUILDER} generate-types \
  --service google.cloud.dataplex.v1 \
  --api-version dataplex.cnrm.cloud.google.com/v1alpha1 \
  --resource DataplexLake:Lake \
  --resource DataplexZone:Zone \
  --resource DataplexTask:Task \
  --resource DataplexEntryGroup:EntryGroup \
  --resource DataplexEntryType:EntryType \
  --resource DataplexDataTaxonomy:DataTaxonomy \
  --resource DataplexAspectType:AspectType \
  --resource DataplexDataScan:DataScan \
  --resource DataplexMetadataJob:MetadataJob \
  --resource DataplexMetadataFeed:MetadataFeed

PROTO_SHA="28ba5d15234eff44241492ca72b135e1171fc537"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"
./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.dataplex.v1 \
  --api-version dataplex.cnrm.cloud.google.com/v1alpha1 \
  --resource DataplexGlossary:Glossary

${CONTROLLERBUILDER} generate-types \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.dataplex.v1 \
  --api-version dataplex.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource DataplexAsset:Asset \
  --resource DataplexDataAsset:DataAsset \
  --resource DataplexDataAttribute:DataAttribute \
  --resource DataplexEncryptionConfig:EncryptionConfig \
  --resource DataplexEntity:Entity \
  --resource DataplexEntry:Entry \
  --resource DataplexEntryLink:EntryLink \
  --resource DataplexGlossaryCategory:GlossaryCategory \
  --resource DataplexGlossaryTerm:GlossaryTerm \
  --resource DataplexContent:Content \
  --resource DataplexEnvironment:Environment

${CONTROLLERBUILDER} generate-identity \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.dataplex.v1 \
  --api-version dataplex.cnrm.cloud.google.com/v1alpha1 \
  --resource DataplexAsset:Asset \
  --resource DataplexDataAsset:DataAsset \
  --resource DataplexDataAttribute:DataAttribute \
  --resource DataplexEncryptionConfig:EncryptionConfig \
  --resource DataplexEntity:Entity \
  --resource DataplexEntry:Entry \
  --resource DataplexEntryLink:EntryLink \
  --resource DataplexGlossaryCategory:GlossaryCategory \
  --resource DataplexGlossaryTerm:GlossaryTerm \
  --resource DataplexContent:Content \
  --resource DataplexEnvironment:Environment

${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.dataplex.v1 \
  --api-version dataplex.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --proto-source-path ${PROTO_OUT}

${CONTROLLERBUILDER} generate-controller \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.dataplex.v1 \
  --api-version dataplex.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource DataplexEntry:Entry

${CONTROLLERBUILDER} generate-fuzzer \
  --proto-source-path ${PROTO_OUT} \
  --service google.cloud.dataplex.v1 \
  --api-version dataplex.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource DataplexEntry:Entry

cd ${REPO_ROOT}
if [[ -z "${SKIP_GENERATE_CRDS:-}" ]]; then
  dev/tasks/generate-crds
fi

if [ -d "${REPO_ROOT}/pkg/controller/direct/dataplex" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/dataplex/
fi
