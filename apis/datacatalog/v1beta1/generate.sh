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

# Find the directory of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_ROOT="$(git rev-parse --show-toplevel)"

cd "${REPO_ROOT}/dev/tools/controllerbuilder"

# We skip scaffold files because we've manually created/renamed datacatalogpolicytag_types.go
# to follow the instructions in the issue, while the generator would have created policytag_types.go.
go run . generate-types \
    --service google.cloud.datacatalog.v1beta1 \
    --api-version datacatalog.cnrm.cloud.google.com/v1beta1 \
    --resource DataCatalogPolicyTag:PolicyTag \
    --output-api "${REPO_ROOT}/apis/" \
    --skip-scaffold-files
