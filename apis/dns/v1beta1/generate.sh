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

# Generate KRM Go types from the OpenAPI / Discovery API JSON specification, ignoring all 'kind' fields
go run "${REPO_ROOT}/dev/tools/openapi-to-krm/main.go" \
  --schema-file "${REPO_ROOT}/apis/dns/v1beta1/dns-api.json" \
  --api-version "dns.cnrm.cloud.google.com/v1beta1" \
  --resource "DNSManagedZone:ManagedZone" \
  --ignore-field "*:kind" \
  --output-file "${REPO_ROOT}/apis/dns/v1beta1/types.generated.go"

# Mapper is commented out for now; we do not need it at this stage.
# go run . generate-mapper \
#   --service google.cloud.dns.v1 \
#   --api-version dns.cnrm.cloud.google.com/v1beta1 \
#   --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds
