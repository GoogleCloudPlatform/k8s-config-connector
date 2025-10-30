# GitHub Actions Workflow Generation

The main CI presubmit workflow file, `.github/workflows/ci-presubmit.yaml`, is automatically generated. This document explains how the generation works and how to add new service-specific test workflows.

## Overview

The generation process is handled by the `dev/tasks/generate-github-actions` script. This script scans the `dev/ci/presubmits/` directory for executable files. For each file it finds, it creates a corresponding job in the `.github/workflows/ci-presubmit.yaml` file.

Most services are tested as part of a general "unclassified" test suite, which is defined in `dev/ci/presubmits/tests-e2e-samples-unclassified`.

To improve test isolation and speed up CI, we are moving towards having dedicated, per-service test workflows. This involves creating a new script for the service in `dev/ci/presubmits/` and excluding it from the "unclassified" suite.

## Instructions for Adding a Per-Service Test Workflow

To create a new CI job that runs E2E tests for a specific service, follow these steps. The instructions use `storage` service as an example.

### 1. Create a new presubmit script for the service

Create a new file named `dev/ci/presubmits/tests-e2e-samples-<service_name>`. For the `storage` service, the file would be `dev/ci/presubmits/tests-e2e-samples-storage`.

This script should run the E2E tests for only the target service. You can use the `RUN_TEST_APIGROUP` environment variable for this. The value should be the service's API group, e.g., `storage.cnrm.cloud.google.com`.

Here is a template for the script. Remember to make it executable (`chmod +x <file_path>`).

**Example for `dev/ci/presubmits/tests-e2e-samples-storage`:**
```bash
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

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/

. ${REPO_ROOT}/dev/tasks/setup-envtest

echo "Running fixtures in tests/e2e for storage..."

# The env var 'KCC_ENG_GROUP_EMAIL' needs to be set to unblock sample
# 'folder-level-entitlement'.
MOCK_GROUP_EMAIL="test@group.com"

RUN_E2E=1 \
E2E_KUBE_TARGET=envtest \
E2E_GCP_TARGET=mock \
KCC_USE_DIRECT_RECONCILERS="ComputeForwardingRule,SpannerInstance" \
KCC_ENG_GROUP_EMAIL="${MOCK_GROUP_EMAIL}" \
RUN_TEST_APIGROUP="storage.cnrm.cloud.google.com" \
  go test -test.count=1 -timeout 1h30m -v ./tests/e2e -run TestAllInSeries/samples
```

### 2. Exclude the service from the unclassified test suite

Modify the `dev/ci/presubmits/tests-e2e-samples-unclassified` script to add the service's API group to the `SKIP_TEST_APIGROUP` environment variable. This prevents the tests from running twice.

For the `storage` service, you would add `storage.cnrm.cloud.google.com` to the comma-separated list in `SKIP_TEST_APIGROUP`.

### 3. Regenerate the GitHub Actions workflow file

Run the `dev/tasks/generate-github-actions` script to update the `.github/workflows/ci-presubmit.yaml` file.

```bash
./dev/tasks/generate-github-actions
```

This will create a new job in the workflow file named `tests-e2e-samples-<service_name>`.

### 4. Verify and commit the changes

Review the changes to ensure the new job is present in `.github/workflows/ci-presubmit.yaml` and the service is excluded from the unclassified tests. Then commit the new script, the updated `tests-e2e-samples-unclassified` script, and the regenerated `ci-presubmit.yaml`.

---

### Reference: `dev/tasks/generate-github-actions`

For reference, here is the content of the generator script:

```bash
#!/usr/bin/env bash
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

set -o errexit
set -o nounset
set -o pipefail

set -x

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}


cat > ${REPO_ROOT}/.github/workflows/ci-presubmit.yaml <<EOF
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

# Generated by dev/tasks/generate-github-actions

name: ci-presubmit

on:
  pull_request:
    types: [opened, synchronize, reopened]
    paths-ignore:
      - "**.md"
      - "experiments/**"
  push:
    branches: ["master"]
    paths-ignore:
      - "**.md"
      - "experiments/**"
  merge_group:
    types: [checks_requested]
    branches: [ "master" ]


jobs:
EOF


for f in $(find dev/ci/presubmits -type f | sort ); do
    name=$(basename ${f})

cat >> ${REPO_ROOT}/.github/workflows/ci-presubmit.yaml <<EOF

  ${name}:
    runs-on: ubuntu-latest
    timeout-minutes: 60
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: 'go.mod'
      - name: "Run ${f}"
        run: |
          ./${f}
        env:
          ARTIFACTS: /tmp/artifacts
      - name: "Upload artifacts"
        uses: actions/upload-artifact@v4
        with:
          name: artifacts-${name}
          path: /tmp/artifacts/

EOF

done


cat >> ${REPO_ROOT}/.github/workflows/ci-presubmit.yaml <<'EOF'

concurrency:
  group: ${{ github.workflow }}-${{ github.head_ref || github.ref }}
  cancel-in-progress: true
EOF
```