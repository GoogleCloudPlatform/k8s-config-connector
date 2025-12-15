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
STALE_VERSION=$(git tag | sort -V | tail -n 1 | sed 's/^v//')
# echo "Stale version is: ${STALE_VERSION}" # Optional: print stale version for clarity, though prompt does it too.
read -p "Enter the new version (stale is ${STALE_VERSION}): " NEW_VERSION

if [ -z "$NEW_VERSION" ]; then
  echo "Error: New version cannot be empty."
  exit 1
fi

# Step 2: Create Release Branch
echo "Creating release branch release-${NEW_VERSION}..."
git checkout -b "release-${NEW_VERSION}"

# Step 3: Propose Tag and Update Manifests
echo "Proposing tag and updating manifests..."
VERSION=${NEW_VERSION} STALE_VERSION=${STALE_VERSION} ./dev/tasks/propose-tag
git add .
git commit -m "Release ${NEW_VERSION}"

# Step 4: Synchronize CRDs
echo "Synchronizing CRDs..."
VERSION=${NEW_VERSION} ./dev/tasks/sync-crds-folder.sh
git add .
git commit -m "Update alpha CRDs for Release ${NEW_VERSION}"

# Step 5: Run Unit Tests
echo "Running unit tests..."
cd operator
# We use an if statement to handle the failure case without exiting due to set -e
if ! (go test ./pkg/controllers/...); then
  echo "Unit tests failed. Updating golden files..."
  WRITE_GOLDEN_OUTPUT="true" go test ./pkg/controllers/...
  git add .
  git commit -m "Update golden files for operator controllers"
  
  echo "Retrying unit tests..."
  go test ./pkg/controllers/...
fi

# Step 6: Format Code
echo "Formatting code..."
cd ..
make fmt
git add .
# Only commit if there are changes
if ! git diff --staged --quiet; then
  git commit -m "[make fmt] Apply formatting"
fi
