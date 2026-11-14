#!/usr/bin/env bash
# Copyright 2026 Google LLC
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

if [ $# -lt 1 ]; then
  echo "Usage: $0 <pr-number>"
  exit 1
fi

PR_NUMBER=$1

echo "Approving PR #${PR_NUMBER}..."
gh pr review "${PR_NUMBER}" --approve

echo "Posting /approve and /lgtm comments for Prow bot..."
gh pr comment "${PR_NUMBER}" --body $'/approve\n/lgtm'

echo "Enabling auto-merge (merge when ready) for PR #${PR_NUMBER}..."
gh pr merge "${PR_NUMBER}" --auto --merge

# Cleanup worktree and temporary branch if they exist
TEMP_BRANCH="review-pr-${PR_NUMBER}"
WORKTREE_PATH="worktrees/${TEMP_BRANCH}"
if [ -d "${WORKTREE_PATH}" ]; then
  echo "Cleaning up worktree at '${WORKTREE_PATH}'..."
  git worktree remove "${WORKTREE_PATH}" || true
fi
if git show-ref --verify --quiet "refs/heads/${TEMP_BRANCH}"; then
  echo "Cleaning up branch '${TEMP_BRANCH}'..."
  git branch -D "${TEMP_BRANCH}" || true
fi

