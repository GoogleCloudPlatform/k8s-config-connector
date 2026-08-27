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

# This script creates sub-issues for tracking test case validation of important KCC resources.
# It checks if an issue for each resource already exists to avoid creating duplicates.

REPO="GoogleCloudPlatform/k8s-config-connector"
PARENT_ISSUE="12569"

resources=(
  "ContainerCluster"
  "FirestoreDatabase"
  "BigtableInstance"
  "BigtableTable"
  "RedisCluster"
  "RedisInstance"
  "MemcacheInstance"
  "SQLInstance"
  "SQLDatabase"
  "AlloyDBCluster"
  "SpannerInstance"
  "SpannerDatabase"
  "FilestoreInstance"
  "StorageBucket"
  "BigQueryDataset"
  "BigQueryTable"
  "ArtifactRegistryRepository"
  "IAMPolicy"
  "IAMPartialPolicy"
)

echo "Checking and creating sub-issues for resource validation..."

for resource in "${resources[@]}"; do
  title="Validate test cases for ${resource}"
  echo "----------------------------------------"
  echo "Resource: ${resource}"

  # Search for an existing issue with the exact title to prevent duplicates
  # Using fixed-string query and jq to verify exact title match
  existing=$(gh issue list --repo "${REPO}" --search "in:title \"Validate test cases for ${resource}\"" --state all --json number,title --jq ".[] | select(.title == \"${title}\") | .number")

  if [ -n "${existing}" ]; then
    echo "Sub-issue already exists: #${existing}"
    continue
  fi

  # Create the issue body based on the template in issue #12569
  body=$(cat <<EOF
Run test cases for ${resource} against real GCP APIs and record the HTTP logs. Log the error and try to fix the test case if the underlying GCP resources can't be created/updated/deleted. If the test can't be fixed, report back in the PR comments.

Part of #${PARENT_ISSUE}.
EOF
)

  echo "Creating sub-issue: '${title}'"
  new_issue_url=$(gh issue create --repo "${REPO}" --title "${title}" --body "${body}" --label "overseer")
  echo "Created: ${new_issue_url}"
done

echo "----------------------------------------"
echo "All sub-issues processed!"
