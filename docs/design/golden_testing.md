<!--
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Golden Testing in Config Connector

Config Connector (KCC) uses "golden testing" to ensure that the controllers behave correctly and that the interaction with Google Cloud APIs remains consistent.

## Overview

Golden testing involves running tests against real or mocked GCP APIs and capturing the resulting HTTP/gRPC traffic. This traffic is stored in "golden" files (typically named `_http.log`). Subsequent test runs compare the actual traffic against these golden files to detect regressions or changes in behavior.

## Test Fixtures

Test fixtures are rooted in `pkg/test/resourcefixture/testdata/basic`. Each fixture is a directory that typically contains:

*   `create.yaml`: The primary resource to be tested.
*   `update.yaml` (optional): An update to the primary resource.
*   `dependencies.yaml` (optional): Dependency resources (e.g., a Network needed for a Compute Instance).
*   `_http.log`: The golden traffic log.

## Testing Workflow

### 1. Recording against real GCP

When adding a new resource or changing behavior, we first record the interaction with the real GCP API.

```bash
export E2E_GCP_TARGET=real
export WRITE_GOLDEN_OUTPUT=1
go test -v ./tests/e2e -run TestAllInSeries/...
```

Helper script: `hack/record-gcp`

### 2. Normalization

Since GCP APIs return volatile values (timestamps, unique IDs, hashes), we normalize the traffic before saving it to `_http.log`. This ensures that the golden files are stable and only change when there's a meaningful change in the API interaction.

Normalization logic is often found in `mockgcp/mock<service>/normalize.go`.

### 3. Verifying against MockGCP

We then run the same tests against our mock implementation of GCP (`mockgcp`).

```bash
export E2E_GCP_TARGET=mock
go test -v ./tests/e2e -run TestAllInSeries/...
```

Helper script: `hack/compare-mock`

If the traffic matches the golden files, the test passes. If there's a mismatch, we either need to update the mock implementation to better reflect real GCP behavior, or (if the change is expected) update the golden files.

## Presubmit Checks

In CI, we run these tests to ensure no regressions are introduced. The scripts are located under `dev/ci/presubmits/tests-e2e-fixtures-*`.

If a presubmit fails, it often means the golden files need to be regenerated or the mock implementation needs fixing.

## Regenerating Golden Files

To regenerate golden files for a specific service, you can run the corresponding presubmit script locally. For example:

```bash
dev/ci/presubmits/tests-e2e-fixtures-storage
```

This will run the tests and update the golden files if `WRITE_GOLDEN_OUTPUT=1` is set (which is typically the default or handled by the script/test runner).

## Hint for Agents

When a golden test fails in CI, the output should include a "hint for agent" to guide automated tools on how to fix the problem:

`<hint_for_agent>Run "dev/ci/presubmits/tests-e2e-fixtures-{name}" to regenerate the golden test output</hint_for_agent>`
