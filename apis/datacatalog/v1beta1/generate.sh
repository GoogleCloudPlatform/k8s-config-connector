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

go run . generate-types \
  --service google.cloud.datacatalog.v1beta1 \
  --api-version datacatalog.cnrm.cloud.google.com/v1beta1 \
  --resource DataCatalogPolicyTag:PolicyTag \
  --include-skipped-output \
  --skip-scaffold-files

# Temporarily move the existing mapper.generated.go if it exists
EXISTING_MAPPER="${REPO_ROOT}/pkg/controller/direct/datacatalog/mapper.generated.go"
TEMP_MAPPER=$(mktemp)
if [ -f "${EXISTING_MAPPER}" ]; then
    cp "${EXISTING_MAPPER}" "${TEMP_MAPPER}"
fi

go run . generate-mapper \
  --service google.cloud.datacatalog.v1beta1 \
  --api-version datacatalog.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output

# Rename the newly generated mapper.generated.go
TARGET_MAPPER="${REPO_ROOT}/pkg/controller/direct/datacatalog/policytag_mapper.generated.go"
mv "${EXISTING_MAPPER}" "${TARGET_MAPPER}"

# Fix the copyright year
sed -i 's/Copyright 2025 Google LLC/Copyright 2026 Google LLC/' "${TARGET_MAPPER}"

# Restore the original mapper.generated.go
if [ -s "${TEMP_MAPPER}" ]; then
    mv "${TEMP_MAPPER}" "${EXISTING_MAPPER}"
fi
rm -f "${TEMP_MAPPER}"

# Clean up types.generated.go as we have all types in datacatalogpolicytag_types.go
rm -f "${REPO_ROOT}/apis/datacatalog/v1beta1/types.generated.go"
