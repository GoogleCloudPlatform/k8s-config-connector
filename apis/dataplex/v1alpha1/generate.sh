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

PROTO_SHA="ee4a3e1ce4e8d16661fcc624322000ad792ffb8a"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

# Force proto generation with correct SHA even if SKIP_GENERATE_PROTOS is set
OLD_SKIP_GENERATE_PROTOS=${SKIP_GENERATE_PROTOS:-}
unset SKIP_GENERATE_PROTOS
./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}
if [ -n "${OLD_SKIP_GENERATE_PROTOS}" ]; then
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
    --resource DataplexMetadataFeed:MetadataFeed \
    --proto-source-path ${PROTO_OUT}

# Post-process types.generated.go to inject kubebuilder validation annotations for recursive self-referential fields
python3 -c "
path = '${REPO_ROOT}/apis/dataplex/v1alpha1/types.generated.go'
with open(path, 'r') as f:
    content = f.read()
content = content.replace(
    'RecordFields []AspectType_MetadataTemplate \`json:\"recordFields,omitempty\"\`',
    '// +kubebuilder:validation:items:XPreserveUnknownFields\n\t// +kubebuilder:validation:items:Type=object\n\tRecordFields []AspectType_MetadataTemplate \`json:\"recordFields,omitempty\"\`'
)
content = content.replace(
    'MapItems *AspectType_MetadataTemplate \`json:\"mapItems,omitempty\"\`',
    '// +kubebuilder:validation:XPreserveUnknownFields\n\t// +kubebuilder:validation:Type=object\n\tMapItems *AspectType_MetadataTemplate \`json:\"mapItems,omitempty\"\`'
)
content = content.replace(
    'ArrayItems *AspectType_MetadataTemplate \`json:\"arrayItems,omitempty\"\`',
    '// +kubebuilder:validation:XPreserveUnknownFields\n\t// +kubebuilder:validation:Type=object\n\tArrayItems *AspectType_MetadataTemplate \`json:\"arrayItems,omitempty\"\`'
)
with open(path, 'w') as f:
    f.write(content)
"

go run . generate-mapper \
    --service google.cloud.dataplex.v1 \
    --api-version "dataplex.cnrm.cloud.google.com/v1alpha1" \
    --proto-source-path ${PROTO_OUT}

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/dataplex/
