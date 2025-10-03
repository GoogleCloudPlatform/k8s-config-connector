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

echo "--- Checking for new version to tag from version/VERSION ---"

# 1. Read the version from the VERSION file at HEAD. 
VERSION_FILE="version/VERSION"
if [ ! -f "$VERSION_FILE" ]; then
    echo "ERROR: Version file not found at ${VERSION_FILE}"
    exit 1
fi
VERSION=$(cat "${VERSION_FILE}" | tr -d '[:space:]')
if [ -z "$VERSION" ]; then
  echo "ERROR: Could not read version from ${VERSION_FILE}."
  exit 1
fi
echo "Found version in ${VERSION_FILE}: $VERSION"

# 2. Check if the tag already exists.
if [ "$(git tag -l "v${VERSION}")" ]; then
  echo "Tag v${VERSION} already exists"
  exit 0
fi
echo "Tag v$VERSION does not exist. Proceeding."

# 3. Find the commit that last modified the VERSION file. This is the commit we will tag.
COMMIT_HASH=$(git log -1 --pretty=format:%H "${VERSION_FILE}")
if [ -z "$COMMIT_HASH" ]; then
  echo "ERROR: Could not find a commit for ${VERSION_FILE}."
  exit 1
fi
echo "Found commit to tag: ${COMMIT_HASH}"

# 4. Verify the version in the file at the target commit matches the version from HEAD.
# This ensures we're tagging the right commit.
VERSION_AT_COMMIT=$(git show "${COMMIT_HASH}:${VERSION_FILE}" | tr -d '[:space:]')
if [ "$VERSION_AT_COMMIT" != "$VERSION" ]; then
    echo "ERROR: Version at HEAD ('$VERSION') does not match version at commit ${COMMIT_HASH} ('${VERSION_AT_COMMIT}')."
    echo "This can happen if ${VERSION_FILE} was modified after the version bump commit."
    exit 1
fi

# 5. Create an annotated tag.
echo "Creating annotated tag v${VERSION} for commit ${COMMIT_HASH}"
git tag -a "v${VERSION}" -m "Release ${VERSION}" "${COMMIT_HASH}"

# 6. Push the tag to the OSS remote.
echo "Pushing tag v${VERSION} to origin."
git push origin "v${VERSION}"

echo "--- Successfully created and pushed tag v${VERSION} ---"
