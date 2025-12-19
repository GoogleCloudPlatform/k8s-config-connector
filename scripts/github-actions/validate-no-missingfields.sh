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
MISSING_FIELDS_FILE="${REPO_ROOT}/tests/apichecks/testdata/exceptions/missingfields.txt"
# Checks for a partial string to determine if the missing field is a reference field
MISSING_REFERENCE_STRING="neither 'external' nor 'name' are set"

# Get the lines added to missingfields.txt in the current branch compared to master
added_lines=$(git diff --unified=0 $BASE_COMMIT..$COMMIT_HEAD "${MISSING_FIELDS_FILE}" | grep '^+[^+]' | sed 's/^+//')
# Filter for missing reference fields
added_reference_lines=$(echo "${added_lines}" | grep "${MISSING_REFERENCE_STRING}")

if [[ -n "${added_reference_lines}" ]]; then
  echo "ERROR: Additions to ${MISSING_FIELDS_FILE} are not allowed."
  echo "The following lines were added:"
  echo "${added_reference_lines}."
  echo "Please ensure the reference field(s) are tested."
  exit 1
elif [[ -n "${added_lines}" ]]; then
  echo "WARNING: New lines were added to ${MISSING_FIELDS_FILE}."
  echo "The following lines were added:"
  echo "${added_lines}"
  echo "Please ensure they are intentional and covered by tests in the future."
  exit 0
else
  echo "No new lines were added to ${MISSING_FIELDS_FILE}."
  exit 0
fi
