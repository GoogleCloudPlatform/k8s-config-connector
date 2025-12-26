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
read -p "Enter the stale version: " STALE_VERSION
read -p "Enter the new version (stale is ${STALE_VERSION}): " NEW_VERSION

if [ -z "$NEW_VERSION" ]; then
  echo "Error: New version cannot be empty."
  exit 1
fi

if [ "$NEW_VERSION" == "$STALE_VERSION" ]; then
    echo "Error: New version ($NEW_VERSION) is the same as the stale version ($STALE_VERSION)."
    exit 1
fi

# sort -V sorts version numbers. If the first line is the new VERSION, then it is not strictly greater than STALE_VERSION.
# We expect STALE_VERSION to be smaller, so it should be first in sort -V output.
SORTED_VERSIONS=$(printf "%s\n%s" "$NEW_VERSION" "$STALE_VERSION" | sort -V)
SMALLEST_VERSION=$(echo "$SORTED_VERSIONS" | head -n1)

if [ "$SMALLEST_VERSION" == "$NEW_VERSION" ]; then
     echo "Error: New version ($NEW_VERSION) is not greater than stale version ($STALE_VERSION)."
     exit 1
fi

# Step 2: Create Release Branch
RELEASE_BRANCH="release-${NEW_VERSION}"
if git rev-parse --verify "${RELEASE_BRANCH}" >/dev/null 2>&1; then
    echo "Release branch ${RELEASE_BRANCH} already exists."
    read -p "Do you want to delete the existing branch and recreate it? (y/N): " response
    if [[ "$response" =~ ^[Yy]$ ]]; then
        git branch -D "${RELEASE_BRANCH}"
        echo "Deleted existing branch ${RELEASE_BRANCH}."
        echo "Creating release branch ${RELEASE_BRANCH}..."
        git checkout -b "${RELEASE_BRANCH}"
    else
        echo "Exiting without creating release branch."
        exit 1
    fi
else
    echo "Creating release branch ${RELEASE_BRANCH}..."
    git checkout -b "${RELEASE_BRANCH}"
fi

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
if ! (cd operator && go test ./pkg/controllers/...); then
  echo "Unit tests failed. Updating golden files..."
  WRITE_GOLDEN_OUTPUT="true" go test ./pkg/controllers/...
  git add .
  git commit -m "Update golden files for operator controllers"

  echo "Retrying unit tests..."
  if ! (go test ./pkg/controllers/...); then
    echo "Unit tests failed even after updating golden files."
    echo "Please fix the test failures manually and try again."
    exit 1
  fi
fi

echo "Validating resource reference docs..."
cd ..
# With VALIDATE_URLS=="true", the doc validation test also validates
# whether the embedded URLs in the template files are accessible.
# If failed, fix the inaccessible URLs in the template files and rerun `make resource-docs`.
VALIDATE_URLS="true" go test ./scripts/generate-google3-docs/...

# Step 6: Format Code
echo "Formatting code..."
<<<<<<< HEAD
=======
cd "$(git rev-parse --show-toplevel)"
>>>>>>> 3f00e08c2f (Improve release scripts)
make fmt
git add .
# Only commit if there are changes
if ! git diff --staged --quiet; then
  git commit -m "[make fmt] Apply formatting"
fi
