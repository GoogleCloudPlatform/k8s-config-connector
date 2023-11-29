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

# Check if any commit has changes both within and outside the TF Git Subtree
subtree_dir="third_party/github.com/hashicorp/terraform-provider-google-beta/"


for commit in $(git rev-list --topo-order --max-count=$COMMIT_CNT $COMMIT_HEAD); do
  parent_commit=$(git rev-parse $commit^)

  if git diff --name-only $parent_commit..$commit | grep "^$subtree_dir" > /dev/null && git diff --name-only $parent_commit..$commit | grep -v "^$subtree_dir" > /dev/null
  then
    echo -e "Error: Your commit \"$commit\" includes changes both within and outside the\n\n\"$subtree_dir\" directory.\n\nPlease ensure that changes made within this directory are grouped and\n\nsubmitted as a separate, dedicated commit."
    exit 1
fi
done

# Regular validations on fmt, generated doc, generated code, etc.
#${REPO_ROOT}/scripts/validate-prereqs.sh
