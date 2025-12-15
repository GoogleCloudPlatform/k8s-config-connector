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

# 4. Verify the commit messages match the release pattern.
# We expect the sequence (oldest to newest):
# 1. Release <VERSION>
# 2. Update alpha CRDs for Release <VERSION>
# 3. (Optional) Update golden files for operator controllers

# We start checking from HEAD and work backwards.
CURRENT_REF="HEAD"

MSG=$(git log --format=%s -n 1 "${CURRENT_REF}")

# Check for Optional commit 0: Merge commit
if [[ "${MSG}" == "Merge pull request"* ]]; then
  echo "Found merge commit at ${CURRENT_REF} ('${MSG}'). Verifying content from the merged branch (HEAD^2)."
  CURRENT_REF="${CURRENT_REF}^2"
  MSG=$(git log --format=%s -n 1 "${CURRENT_REF}")
fi

# Check for Optional commit 3: Golden files
EXPECTED_GOLDEN="Update golden files for operator controllers"
if [ "${MSG}" = "${EXPECTED_GOLDEN}" ]; then
  echo "Found golden files update commit at ${CURRENT_REF}."
  CURRENT_REF="${CURRENT_REF}~1"
  MSG=$(git log --format=%s -n 1 "${CURRENT_REF}")
fi

# Check for Required commit 2: Alpha CRDs
EXPECTED_CRDS="Update alpha CRDs for Release ${VERSION}"
if [ "${MSG}" = "${EXPECTED_CRDS}" ]; then
  echo "Found alpha CRDs update commit at ${CURRENT_REF}."
  CURRENT_REF="${CURRENT_REF}~1"
  MSG=$(git log --format=%s -n 1 "${CURRENT_REF}")
else
  echo "ERROR: Expected commit message '${EXPECTED_CRDS}' at ${CURRENT_REF}, but found '${MSG}'"
  echo "The release PR must include an '${EXPECTED_CRDS}' commit."
  exit 1
fi

# Check for Required commit 1: Release version
EXPECTED_RELEASE="Release ${VERSION}"
if [ "${MSG}" != "${EXPECTED_RELEASE}" ]; then
  echo "ERROR: Expected commit message '${EXPECTED_RELEASE}' at ${CURRENT_REF} (derived from HEAD), but found '${MSG}'"
  echo "The release PR must typically start with a 'Release ${VERSION}' commit, optionally followed by CRD updates and/or golden file updates."
  exit 1
fi

echo "Verified commit messages match release pattern."

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
