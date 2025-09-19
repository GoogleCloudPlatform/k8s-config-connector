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
if git rev-parse "$VERSION" >/dev/null 2>&1; then
  echo "Tag $VERSION already exists. Nothing to do."
  exit 0
fi
echo "Tag $VERSION does not exist. Proceeding."

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

# 5. Derive branch name for the release process.
VERSION_MAJOR_MINOR=$(echo "$VERSION" | cut -d. -f1,2)
if ! [[ "$VERSION_MAJOR_MINOR" =~ ^[0-9]+\.[0-9]+$ ]]; then
  echo "ERROR: Could not derive major.minor version from $VERSION"
  exit 1
fi
BRANCH_NAME="release_${VERSION_MAJOR_MINOR}"

# Save current branch/state to return to it later.
ORIGINAL_STATE=$(git rev-parse --abbrev-ref HEAD)
if [ "$ORIGINAL_STATE" = "HEAD" ]; then
  # We are in a detached HEAD state, common in CI. Save the commit hash instead.
  ORIGINAL_STATE=$(git rev-parse HEAD)
fi


# 6. Create a new branch from the specific commit to match the release process.
echo "Creating new branch '${BRANCH_NAME}' from commit ${COMMIT_HASH}..."
# Use -B to create or reset the branch, making the script resilient to previous failures.
git checkout -B "${BRANCH_NAME}" "${COMMIT_HASH}"

# 7. Set variables and construct the Go command for pushing the tag.
if [ -z "${BUGANIZER_ID}" ]; then
  echo "ERROR: BUGANIZER_ID environment variable not set."
  exit 1
fi
REMOTE="sso://cnrm/cnrm"
PUSH_OPTIONS="push-justification=b/${BUGANIZER_ID}"
REPO_PATH=$(pwd)

GO_COMMAND=(go run .
  --remote "${REMOTE}"
  --branch "${BRANCH_NAME}"
  --version-file "${VERSION_FILE}"
  --source "${REPO_PATH}"
  --push-options "${PUSH_OPTIONS}"
)

# 8. Perform a dry run.
echo "--- Step 1: Performing Dry Run ---"
DRY_RUN_EXIT_CODE=0
DRY_RUN_OUTPUT=$("${GO_COMMAND[@]}" 2>&1) || DRY_RUN_EXIT_CODE=$?

if [ ${DRY_RUN_EXIT_CODE} -ne 0 ] || echo "${DRY_RUN_OUTPUT}" | grep -q -E '^\'EWF\''; then
  echo "ERROR: Dry run failed."
  echo "Exit code: ${DRY_RUN_EXIT_CODE}."
  echo "Full output:"
  echo "${DRY_RUN_OUTPUT}"
  git checkout "${ORIGINAL_STATE}"
  exit 1
else
  echo "Dry run successful."
fi

# 9. If dry run was successful, proceed with the actual push.
echo "--- Step 2: Pushing Tag ---"
if "${GO_COMMAND[@]}" --yes=true; then
  echo "--- Tag push command executed successfully. ---"
else
  echo "ERROR: Tag push command failed."
  git checkout "${ORIGINAL_STATE}"
  exit 1
fi

# 10. Return to the previous branch/state and clean up.
echo "Returning to previous git state and cleaning up."
git checkout "${ORIGINAL_STATE}"
git branch -D "${BRANCH_NAME}"

echo "--- Successfully created and pushed tag ${VERSION} ---