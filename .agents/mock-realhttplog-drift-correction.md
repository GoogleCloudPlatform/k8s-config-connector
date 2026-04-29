---
name: Mock RealHTTPLog Drift Correction
description: Detects drift in real http logs from e2e tests and creates issues per test to align mockgcp behavior.
schedule: "@daily"
skipPR: true
---

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

# Role
You are a software development assistant for the Kubernetes Config Connector project.
Your goal is to detect drift in real http logs from e2e test runs and create GitHub issues to track the required fixes for aligning the mockgcp behavior.
You have access to the following tools:
- GitHub CLI (`gh`)
- bash / git
Use `gh` and `bash`/`git` to perform your duties.

# Task
1. Run `./dev/tasks/download-e2e-logs` without arguments. This will fail but output a list of tests. Parse the output and filter for tests that begin with `e2e-`.
2. For each test in the list, run `TEST=<test_name> ./dev/tasks/download-e2e-logs`. This will download the logs from the storage bucket of the last run into the testdata directory.
3. Check for modified files using `git status --porcelain`. Look specifically for modifications in `_http.log` files under `pkg/test/resourcefixture/testdata/basic/`.
4. Identify the list of drifted services. For each modified `_http.log` file, determine the corresponding `<service>` (it is usually the folder name immediately under `pkg/test/resourcefixture/testdata/basic/`, e.g., if `pkg/test/resourcefixture/testdata/basic/pubsub/...` is modified, the service is `pubsub`).
5. If the number of drifted services is more than 4, limit your work to the first 4 to avoid overwhelming the team.
6. For each of the drifted services (up to 4):
    - Check if an issue already exists (open or closed) for fixing mock real http log drift for this service using: `gh issue list --state all --search "Fix mock realhttp log drift for <service>"`
    - If an issue already exists, skip creating a new one.
    - If no issue exists, create a new GitHub issue using the `gh` tool.
    - The issue title should be: `Fix mock realhttp log drift for <service>`
    - The issue should be labeled with: `overseer`, `priority/medium`, `step/mockgcp`.
    - The issue body MUST contain the exact text from the **ISSUE BODY TEMPLATE** below, replacing `<service>` with the appropriate service name.
    - Append a link to this chore file (`.agents/mock-realhttplog-drift-correction.md`) at the end of the issue body for traceability.
7. Discard all local changes to the repo: `git checkout main && git reset --hard HEAD && git clean -fd` (Use the correct default branch, it might be master or main).
8. IMPORTANT: If there are more than 15 open PRs for this task, stop and dont create any more.

## Issue Title

`Fix mock realhttp log drift for <service>`

## Issue Labels
The issue should be labeled with the following labels:
* `overseer`
* `priority/medium`
* `step/mockgcp`

## ISSUE BODY TEMPLATE

The issue body should contain this text template with the appropriate service filled in.
The body template is treated as markdown. Retain the formatting as is when filling in the service.

------------ BEGIN ISSUE BODY TEMPLATE ------------
This issue tracks the required updates to the mock definitions for `<service>` to align `mockgcp` behavior with recent real HTTP logs.

### Instructions

- [ ] 1. **Identify the drift**:
    To see the drift, you can download the real HTTP logs for `<service>` tests by running the download script for the relevant test (e.g. `TEST=e2e-service-<service>` or similar):
    ```bash
    TEST=e2e-service-<service> ./dev/tasks/download-e2e-logs
    ```
    Then, check for modified files using `git status`. Look for modifications in `_http.log` files under `pkg/test/resourcefixture/testdata/basic/<service>/`.

- [ ] 2. **Build the mock**:
    Navigate to the mock service directory and ensure it builds.
    ```bash
    cd mockgcp/mock<service>
    go build .
    ```

- [ ] 3. **Run the tests**:
    Run the tests against the mock implementation using the corresponding dev script to observe the failures.
    Example:
    ```bash
    cd dev/ci/presubmits/
    ./tests-e2e-fixtures-<service>
    ```
    *(Note: The `<service>`'s test fixtures script might be named differently or bundled with other scripts. You may need to identify the correct test script first. You can also run `dev/tools/compare-mock <testpath>` to test a specific resource fixture.)*

- [ ] 4. **Fix the mock implementation**:
    If the tests fail, update the `mockgcp` implementation for `<service>` to make the tests pass. 
    This may involve updating the mock itself or adding normalization rules to `normalize.go`. 
    **Critical:** If you update `normalize.go`, ensure that `Previsit` rules are explicitly scoped to the `<service>` URL to avoid side effects on other services.
    Repeat steps 3 and 4 until the tests pass. After a reasonable number of attempts, you can give up.

- [ ] 5. **Create a PR**:
    Create a PR including the changes in `mockgcp/mock<service>` and the updated `_http.log` files.
    Create the PR using `gh` and don't forget to link it to this issue (e.g., `Fixes #<issue number>`) in the PR description.
    Include a link to the chore file (`.agents/mock-realhttplog-drift-correction.md`) in the PR description.
    Apply the same labels to the PR as are on this issue.
    
------------ END ISSUE BODY TEMPLATE ------------
