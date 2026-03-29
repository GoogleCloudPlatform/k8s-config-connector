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

set -exu; set -o pipefail

# Step 1: Determine Release Versions
if [[ -z "${STALE_VERSION:-}" ]]; then
  if [[ -f version/VERSION ]]; then
    STALE_VERSION=$(cat version/VERSION | cut -d. -f1,2)
    echo "Auto-detected STALE_VERSION: ${STALE_VERSION}"
  else
    echo "Error: STALE_VERSION cannot be empty and version/VERSION does not exist."
    exit 1
  fi
fi

NEW_VERSION=${1:-${VERSION:-}}
if [[ -z "${NEW_VERSION}" ]]; then
  echo "Error: NEW_VERSION not provided as first argument and VERSION env var is not set."
  exit 1
fi

if [ -z "$NEW_VERSION" ]; then
  echo "Error: New version cannot be empty."
  exit 1
fi

# Step 2: Create Release Branch
echo "Creating release branch release-${NEW_VERSION}..."
git checkout -B "release-${NEW_VERSION}"

# Step 3: Propose Tag and Update Manifests
echo "Proposing tag and updating manifests..."
VERSION=${NEW_VERSION} STALE_VERSION=${STALE_VERSION} ./dev/tasks/propose-tag

# Step 4: Run Unit Tests and Update Goldens if needed
echo "Running unit tests. Golden files will be updated if there are failures..."
cd operator
# We use an if statement to handle the failure case without exiting due to set -e
if ! (go test ./pkg/controllers/...); then
  echo "Unit tests failed. Updating golden files..."
  WRITE_GOLDEN_OUTPUT="true" go test ./pkg/controllers/...
fi
cd ..

# Step 5: Format Code
echo "Formatting code..."
make fmt

# Step 6: Consolidate everything into a single commit
echo "Staging all changes and creating the release commit..."
git add .
git commit -m "Release ${NEW_VERSION}"

echo "--- Release branch release-${NEW_VERSION} prepared with a single commit ---"
