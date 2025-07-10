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


# Exit immediately if a command exits with a non-zero status.
set -e

# --- Environment Variable Check ---
# This script requires GIT_COMMIT and VERSION to be set as environment variables.

if [ -z "$GIT_COMMIT" ]; then
  echo "ERROR: The GIT_COMMIT environment variable is not set."
  echo "Usage: export GIT_COMMIT=<commit-hash>"
  exit 1
fi

if [ -z "$VERSION" ]; then
  echo "ERROR: The VERSION environment variable is not set."
  echo "Usage: export VERSION=<version-string>"
  exit 1
fi

if [ -z "$REPO_PATH" ]; then
  echo "ERROR: The REPO_PATH environment variable is not set."
  echo "Usage: export REPO_PATH=<repo-path-string>"
  exit 1
fi


# --- Configuration ---
# Please edit the variables below if they differ from your setup.

# The local path to your GoB source code checkout.
# This should be the path where you have cloned the Github repository.
# e.g., "/usr/local/google/home/${username}/k8s-config-connector/""
SOURCE_CHECKOUT_PATH="${REPO_PATH}"

# --- Script ---

BRANCH_NAME="push_tag_${VERSION}"
REMOTE="sso://cnrm/cnrm"
VERSION_FILE="version/VERSION"
PUSH_OPTIONS="push-justification=b/${BUGANIZER_ID}"

echo "--- Preparing local repository ---"
echo "Using GIT_COMMIT=${GIT_COMMIT}"
echo "Using VERSION=${VERSION}"

echo "Checking out commit: ${GIT_COMMIT}"
git checkout "${GIT_COMMIT}"

echo "Creating new branch: ${BRANCH_NAME}"
git checkout -b "${BRANCH_NAME}"

# The command to be run, constructed from the main.go flags and the release document.
GO_COMMAND=(go run .
  --remote upstream
  --branch upstream/master
  --version-file "${VERSION_FILE}"
  --source "${SOURCE_CHECKOUT_PATH}"
  -v=2
)

echo ""
echo "--- Step 1: Performing Dry Run ---"
echo "The following command will be executed for a dry run:"
echo "${GO_COMMAND[@]}"
echo ""

"${GO_COMMAND[@]}"

echo ""
echo "--- Step 2: Push Tag ---"
read -p "Dry run complete. Do you want to proceed with pushing the tag? (y/n) " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
  echo "Proceeding to push the tag..."
  # The release document mentions adding '--yes=1' and the main.go file defines this flag as '--yes'.
  "${GO_COMMAND[@]}" --yes=true
  echo "--- Tag pushed successfully. ---"
else
  echo "Aborted. No tag was pushed."
  exit 1
fi
