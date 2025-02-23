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
cd "${REPO_ROOT}"

### This check ensures that the generated Go types for direct resources are not manually modified by accident.
### Ensures that the code generation tools can be safely re-run.
### This check ignores copyright year changes.
make generate-types
diff_output=$(git diff --unified=0 | grep -v -E "^(diff --git|index |--- |\+\+\+ |@@ |[+-][[:space:]]*// Copyright.*Google LLC)") || true
if echo "$diff_output" | grep -q "^[+-]"; then
    echo "Full diff (excluding copyright changes):"
    echo "$diff_output"
    echo "ERROR: The generated types are outdated. Run 'make generate-types' to update them."
    echo "If you need to modify any types, first move them out of the generated file."
    echo "Then run 'make generate-types' again to ensure the generated file remains unchanged."
    echo "Affected files:"
    git diff --name-only
    exit 1
fi