#!/bin/bash
# Copyright 2022 Google LLC
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

make fmt
changed_file_count=$(git diff --name-only | wc -l)
if [[ "${changed_file_count}" != "0" ]]; then
    echo "Full diff:"
    git diff
    echo "ERROR: Files not formatted correctly. Please run 'make fmt' and update your PR."
    echo "Affected files:"
    git diff --name-only
    exit 1
fi
make generate
changed_file_count=$(git diff --name-only | wc -l)
if [[ "${changed_file_count}" != "0" ]]; then
    echo "Full diff:"
    git diff
    echo "ERROR: Generated code out-of-date. Please run 'make generate' and update your PR."
    echo "Affected files:"
    git diff --name-only
    exit 1
fi
make manifests
changed_file_count=$(git diff --name-only | wc -l)
if [[ "${changed_file_count}" != "0" ]]; then
    echo "Full diff:"
    git diff
    echo "ERROR: Manifests must be regenerated. Please run 'make ready-pr' or 'make manifests' and update your PR."
    echo "Affected files:"
    git diff --name-only
    exit 1
fi
make generate-go-client
changed_file_count=$(git diff --name-only | wc -l)
added_go_client_file_count=$(git ls-files --others --exclude-standard pkg/clients/generated/ | wc -l)
if [[ "${changed_file_count}" != "0" ]] || [[ "${added_go_client_file_count}" != "0" ]]; then
    echo "Full diff:"
    git diff
    echo "ERROR: Resource Go Clients must be regenerated. Please run 'make ready-pr' or 'make generate-go-client' and update your PR."
    echo "Affected files:"
    git diff --name-only
    git ls-files --others --exclude-standard pkg/clients/generated/
    echo "First 100 lines of diff:"
    git diff | head -n100
    exit 1
fi
make ensure
if [[ $? -ne 0 ]]; then
  echo "'make ensure' failed. Please validate the override patch files."
fi
changed_file_count=$(git diff --name-only | wc -l)
if [[ "${changed_file_count}" != "0" ]]; then
    echo "Full diff:"
    git diff
    echo "ERROR: Vendored code does not match go.mod (+ override patches). Please run 'make ensure' and update your PR."
    echo "Affected files:"
    git diff --name-only
    exit 1
fi
make resource-docs
changed_file_count=$(git diff --name-only | wc -l)
added_reference_doc_file_count=$(git ls-files --others --exclude-standard scripts/generate-google3-docs/resource-reference/generated/ | wc -l)
if [[ "${changed_file_count}" != "0" ]] || [[ "${added_reference_doc_file_count}" != "0" ]]; then
    echo "Full diff:"
    git diff
    echo "ERROR: Resource docs must be regenerated. Please run 'make ready-pr' or 'make resource-docs' and update your PR."
    echo "Affected files:"
    git diff --name-only
    git ls-files --others --exclude-standard scripts/generate-google3-docs/resource-reference/generated/
    exit 1
fi
