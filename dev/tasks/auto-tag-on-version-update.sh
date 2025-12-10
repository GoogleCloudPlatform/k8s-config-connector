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

# 3. Use the current commit (HEAD) as the tag target.
# We tag the end of the release PR (which includes 3 commits), not just the first commit that bumped the version.
# We assume this script runs on the tip of the release branch/PR.
COMMIT_HASH=$(git rev-parse HEAD)
if [ -z "${COMMIT_HASH}" ]; then
  echo "ERROR: Could not determine commit hash."
  exit 1
fi
echo "Using HEAD as commit to tag: ${COMMIT_HASH}"

# 4. Verify the commit messages of the last 3 commits.
# Expected order (oldest to newest):
# 1. Release <VERSION>
# 2. Update alpha CRDs for Release <VERSION>
# 3. Update golden files for operator controllers

# HEAD~2 (First commit)
MSG_FIRST=$(git log --format=%s -n 1 HEAD~2)
EXPECTED_FIRST="Release ${VERSION}"
if [ "${MSG_FIRST}" != "${EXPECTED_FIRST}" ]; then
  echo "ERROR: HEAD~2 commit message '${MSG_FIRST}' does not match expected '${EXPECTED_FIRST}'"
  exit 1
fi

# HEAD~1 (Second commit)
MSG_SECOND=$(git log --format=%s -n 1 HEAD~1)
EXPECTED_SECOND="Update alpha CRDs for Release ${VERSION}"
if [ "${MSG_SECOND}" != "${EXPECTED_SECOND}" ]; then
  echo "ERROR: HEAD~1 commit message '${MSG_SECOND}' does not match expected '${EXPECTED_SECOND}'"
  exit 1
fi

# HEAD (Third commit)
MSG_THIRD=$(git log --format=%s -n 1 HEAD)
EXPECTED_THIRD="Update golden files for operator controllers"
if [ "${MSG_THIRD}" != "${EXPECTED_THIRD}" ]; then
  echo "ERROR: HEAD commit message '${MSG_THIRD}' does not match expected '${EXPECTED_THIRD}'"
  exit 1
fi

echo "Verified last 3 commit messages match release pattern."

# 4. Verify the version in the file at the target commit matches the version from HEAD.
# This ensures we're tagging the right commit.
VERSION_AT_COMMIT=$(git show "${COMMIT_HASH}:${VERSION_FILE}" | tr -d '[:space:]')
if [ "$VERSION_AT_COMMIT" != "$VERSION" ]; then
    echo "ERROR: Version at HEAD ('$VERSION') does not match version at commit ${COMMIT_HASH} ('${VERSION_AT_COMMIT}')."
    echo "This can happen if ${VERSION_FILE} was modified after the version bump commit."
    exit 1
fi

# 5. Create the release branch.
# Extract Major.Minor from the version (e.g., 1.139.0 -> 1.139)
MAJOR_MINOR=$(echo "$VERSION" | cut -d. -f1,2)
RELEASE_BRANCH="release-${MAJOR_MINOR}"

echo "Checking if release branch ${RELEASE_BRANCH} exists on remote..."
if git ls-remote --exit-code --heads origin "${RELEASE_BRANCH}"; then
  echo "Release branch ${RELEASE_BRANCH} already exists on remote. Skipping branch creation."
else
  echo "Creating release branch ${RELEASE_BRANCH} from commit ${COMMIT_HASH}"
  git push origin "${COMMIT_HASH}:refs/heads/${RELEASE_BRANCH}"
fi

# 6. Create an annotated tag on the release branch (or the commit if branch exists).
# We tag the commit explicitly, but pushing the branch first ensures the commit is reachable there.
echo "Creating annotated tag v${VERSION} for commit ${COMMIT_HASH}"
git tag -a "v${VERSION}" -m "Release ${VERSION}" "${COMMIT_HASH}"

# 7. Push the tag to the OSS remote.
echo "Pushing tag v${VERSION} to origin."
git push origin "v${VERSION}"

echo "--- Successfully processed release branch ${RELEASE_BRANCH} and tag v${VERSION} ---"
