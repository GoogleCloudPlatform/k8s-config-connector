#!/bin/bash
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

function usage() {
  echo "Usage: $0 --title <pr_title> --body <pr_body_file> [--base <base_branch>]"
  echo "Example: $0 --title 'Fix issue 123' --body pr-body.txt"
  exit 1
}

TITLE=""
BODY_FILE=""
BASE="master"

while [[ $# -gt 0 ]]; do
  case $1 in
    --title)
      TITLE="$2"
      shift 2
      ;;
    --body)
      BODY_FILE="$2"
      shift 2
      ;;
    --base)
      BASE="$2"
      shift 2
      ;;
    *)
      echo "Unknown option $1"
      usage
      ;;
  esac
done

if [[ -z "$TITLE" || -z "$BODY_FILE" ]]; then
  echo "Error: --title and --body are required."
  usage
fi

if [[ ! -f "$BODY_FILE" ]]; then
  echo "Error: Body file $BODY_FILE not found."
  exit 1
fi

if git ls-files --error-unmatch pr-body.txt >/dev/null 2>&1; then
  echo "Error: pr-body.txt is checked into git."
  echo "Please remove it (e.g., git rm pr-body.txt) and commit the removal before sending a PR."
  exit 1
fi

if ! git diff --quiet; then
  echo "Error: There are uncommitted changes in your working directory."
  echo "Please commit your work before running this script."
  git status --short
  exit 1
fi

echo "Running make fmt..."
make fmt

if ! git diff --quiet; then
  echo "Error: 'make fmt' introduced formatting changes."
  echo "Please commit these changes before sending a PR."
  git status --short
  exit 1
fi

BRANCH=$(git branch --show-current)
if [[ -z "$BRANCH" || "$BRANCH" == "master" || "$BRANCH" == "main" ]]; then
  echo "Error: Please create a new branch for your PR. Currently on $BRANCH."
  exit 1
fi

echo "Pushing branch $BRANCH to origin..."
git push --set-upstream origin "$BRANCH" || git push --force-with-lease --set-upstream origin "$BRANCH"

echo "Checking for existing PR..."
if gh pr view "$BRANCH" >/dev/null 2>&1; then
  echo "PR already exists for branch $BRANCH."
  # Edit the existing PR title and body
  gh pr edit "$BRANCH" --title "$TITLE" --body-file "$BODY_FILE"
  echo "PR updated."
else
  echo "Creating new PR..."
  gh pr create --title "$TITLE" --body-file "$BODY_FILE" --base "$BASE"
  echo "PR created successfully."
fi
