#!/bin/bash
# Copyright 2023 Google LLC
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
REPO_ROOT="$(git rev-parse --show-toplevel)"
source ${REPO_ROOT}/scripts/shared-vars-public.sh
cd ${REPO_ROOT}
source ${REPO_ROOT}/scripts/fetch_ext_bins.sh && \
	fetch_tools && \
	setup_envs

echo "Running validations..."

# If EVENT_NAME is "pull_request", then check if any commit contains changes both within and outside the TF Git Subtree
if [[ "${EVENT_NAME:-}" == "pull_request" ]]; then
  # Check if any commit contains changes both within and outside the TF Git Subtree
  SUBTREE_DIR="third_party/github.com/hashicorp/terraform-provider-google-beta/"
  if git diff --name-only $BASE_COMMIT..$COMMIT_HEAD | grep "^$SUBTREE_DIR" > /dev/null
  then
    COMMIT_HASHES=($(git rev-list --topo-order -n $COMMIT_CNT $COMMIT_HEAD))

    for COMMIT in "${COMMIT_HASHES[@]}"; do
      PARENT_COMMIT=$(git rev-parse $COMMIT^)

      if git diff --name-only $PARENT_COMMIT..$COMMIT | grep "^$SUBTREE_DIR" > /dev/null && git diff --name-only $PARENT_COMMIT..$COMMIT | grep -v "^$SUBTREE_DIR" > /dev/null
      then
        echo -e "Error: Your commit \"$COMMIT\" includes changes both within and outside the\n\"$SUBTREE_DIR\" directory.\nPlease ensure that changes made within this directory are grouped and\nsubmitted as a separate, dedicated commit.\n"
        exit 1
      fi
    done
  fi
fi

# Regular validations on fmt, generated doc, generated code, etc.
${REPO_ROOT}/scripts/validate-prereqs.sh
