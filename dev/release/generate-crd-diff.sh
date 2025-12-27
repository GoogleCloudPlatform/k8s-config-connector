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


set -e

# Get the last 4 commit subjects
COMMIT_MSG_1=$(git log -1 --pretty=%s)
COMMIT_MSG_2=$(git log -2 --pretty=%s | tail -n 1)
COMMIT_MSG_3=$(git log -3 --pretty=%s | tail -n 1)
COMMIT_MSG_4=$(git log -4 --pretty=%s | tail -n 1)

# The expected commit messages.
EXPECTED_MSG_FMT="[make fmt] Apply formatting"
EXPECTED_MSG_GOLDEN="Update golden files for operator controllers"
EXPECTED_MSG_CRD="Update alpha CRDs for Release"
RELEASE_COMMIT_REGEX="^Release [0-9]+\.[0-9]+\.[0-9]+"

CRD_COMMIT_HASH=""

# The script should handle the case where there are optional format and/or golden file update commits.

if [[ "$COMMIT_MSG_1" == "$EXPECTED_MSG_FMT" ]] && \
   [[ "$COMMIT_MSG_2" == "$EXPECTED_MSG_GOLDEN" ]] && \
   [[ "$COMMIT_MSG_3" == *"$EXPECTED_MSG_CRD"* ]] && \
   [[ "$COMMIT_MSG_4" =~ $RELEASE_COMMIT_REGEX ]]; then
    echo "Found the expected sequence of commits with format and golden file update."
    CRD_COMMIT_HASH=$(git log -3 --pretty=%H | tail -n 1)
# Case 2: Format update present, Golden file update ABSENT
elif [[ "$COMMIT_MSG_1" == "$EXPECTED_MSG_FMT" ]] && \
     [[ "$COMMIT_MSG_2" == *"$EXPECTED_MSG_CRD"* ]] && \
     [[ "$COMMIT_MSG_3" =~ $RELEASE_COMMIT_REGEX ]]; then
    echo "Found the expected sequence of commits with format update."
    CRD_COMMIT_HASH=$(git log -2 --pretty=%H | tail -n 1)
# Case 3: Format update absent, Golden file update PRESENT
elif [[ "$COMMIT_MSG_1" == "$EXPECTED_MSG_GOLDEN" ]] && \
   [[ "$COMMIT_MSG_2" == *"$EXPECTED_MSG_CRD"* ]] && \
   [[ "$COMMIT_MSG_3" =~ $RELEASE_COMMIT_REGEX ]]; then
    echo "Found the expected sequence of commits with golden file update."
    CRD_COMMIT_HASH=$(git log -2 --pretty=%H | tail -n 1)
# Case 4: Format update absent, Golden file update ABSENT
elif [[ "$COMMIT_MSG_1" == *"$EXPECTED_MSG_CRD"* ]] && \
     [[ "$COMMIT_MSG_2" =~ $RELEASE_COMMIT_REGEX ]]; then
    echo "Found the expected sequence of commits without optional updates."
    CRD_COMMIT_HASH=$(git log -1 --pretty=%H)
else
    echo "Error: The last commits do not match the expected release commit sequence."
    echo "Commit 1: $COMMIT_MSG_1"
    echo "Commit 2: $COMMIT_MSG_2"
    echo "Commit 3: $COMMIT_MSG_3"
    echo "Commit 4: $COMMIT_MSG_4"
    exit 1
fi

if [ -z "$CRD_COMMIT_HASH" ]; then
    echo "Error: CRD_COMMIT_HASH could not be determined."
    exit 1
fi

echo "Generating diff for commit: $CRD_COMMIT_HASH"
git diff $CRD_COMMIT_HASH^! > crd_diff.tmp

echo "Diff saved to crd_diff.tmp"
