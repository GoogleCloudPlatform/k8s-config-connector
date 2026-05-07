---
name: Mock Drift Correction
description: Detects drift in mockgcp protos and creates issues per service to address them.
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
Your goal is to detect upstream proto changes in `mockgcp` and create GitHub issues to track the required fixes for the drift.
You have access to the following tools:
- GitHub CLI (`gh`)
- git
Use `gh` and `git` to perform your duties.

# Task
1. Navigate to the `mockgcp` directory: `cd mockgcp`
2. Run `make all` to regenerate the generated code and detect any proto drift from upstream.
3. Check for modified files: `git status --porcelain`. Look specifically for modifications in `.proto` files under `apis/mockgcp/cloud/`.
4. If no `.proto` files are modified under `apis/mockgcp/cloud/`, try updating the `https://github.com/googleapis/googleapis` entry in `mockgcp/git.versions` to point to `master`, and run `make all` again to detect drift against master.
5. Check again for modified files: `git status --porcelain`. Look specifically for modifications in `.proto` files under `apis/mockgcp/cloud/`.
6. Identify the list of drifted services. For each modified `.proto` file, determine the corresponding `<service>` (e.g., if `apis/mockgcp/cloud/apigee/v1/service.proto` is modified, the service is `apigee`).
7. If the number of modified services is more than 10, limit your work to 10 to avoid overwhelming the team.
8. For each of the up to 10 drifted services:
    - Check if an issue already exists (open or closed) for fixing mock drift for this service using: `gh issue list --state all --search "Fix mock drift for <service>"`.
    - If an issue already exists, skip creating a new one.
    - If no issue exists, create a new GitHub issue using the `gh` tool.
    - The issue title should be: `Fix mock drift for <service>`
    - The issue should be labeled with: `overseer`, `priority/medium`, `step/mockgcp`.
    - The issue body MUST contain the exact text from the **ISSUE BODY TEMPLATE** below, replacing `<service>` with the appropriate service name.
    - Append a link to this chore file (`.agents/mock-drift-correction.md`) at the end of the issue body for traceability.
9. Discard all local changes to the repo: `git checkout master && git reset --hard HEAD`

## Issue Title

`Fix mock drift for <service>`

## Issue Labels
The issue should be labeled with the following labels:
* `overseer`
* `priority/medium`
* `step/mockgcp`

## ISSUE BODY TEMPLATE

The issue body should contain this text template with the appropriate service filled in.
The body template is treated as markdown. Retain the formatting as is when filling in the service.

------------ BEGIN ISSUE BODY TEMPLATE ------------
This issue tracks the required updates to the mock definitions for `<service>` to address drift from upstream proto changes.

### Context & Guidelines

We should **always prefer to use the proto where there is a proto**. The service team is responsible for publishing an updated proto and client libraries.

There are two ways we consume protos in `mockgcp`:
1. **via the `googleapis/googleapis` repo**: This means we compile our own protos to `.pb.go`. This approach makes it easier to "patch the protos" and compile in a change (e.g., for service teams testing new fields).
2. **using the precompiled versions in the google clients**: (e.g., as seen in `mockbillingbudgets`). This is much easier because there is only one proto version in play and it avoids `grpc-gateway`. The disadvantage is that sometimes the google clients are not up to date.

Keep these approaches in mind when updating the mock definitions.

### Instructions

- [ ] 1. **Identify the drift**:
    Navigate to the `mockgcp` directory and run `make all` to regenerate the generated code and detect proto drift from upstream.
    ```bash
    cd mockgcp
    make all
    ```
    If no `.proto` files are modified under `apis/mockgcp/cloud/`, update the `https://github.com/googleapis/googleapis` entry in `git.versions` to point to `master`, and run `make all` again to detect drift against master.
    ```bash
    sed -i 's|https://github.com/googleapis/googleapis .*|https://github.com/googleapis/googleapis master|' git.versions
    make all
    ```
    Check for modified files using `git status`. Look for modifications in `.proto` files and their corresponding generated `.go` files for `<service>`.

- [ ] 2. **Build the mock**:
    Navigate to the mock service directory and ensure it builds.
    ```bash
    cd mockgcp/mock<service>
    go build .
    ```
    If it doesn't build, you need to fix the compile errors in the mock implementation.

- [ ] 3. **Run the tests**:
    If it builds, run the tests using the corresponding dev script.
    example:
    ```bash
    cd dev/ci/presubmits/
    ./tests-e2e-fixtures-<service>
    ```
    The <service>'s test fixtures script could be names differently or bundled with other scripts.
    You may have to identify the corresponding test script first.

- [ ] 4. **Fix the mock implementation**:
    If the tests fail, update the mockgcp implementation for `<service>` to make the tests pass.
    Repeat step 4 till the tests pass. After a reasonable attemps you can give up.

- [ ] 5. **Create a PR**:
    Add the changes in mock<service> that you worked on. Its ok if there are no changes in that folder.
    Add the modified proto files for <service> in `apis/mockgcp/cloud/` and their corresponding generated `.pb.go`, `.pb.gw.go`, and `_grpc.pb.go` files.
    Create a PR using gh and don't forget to link the Fixes #<issue number> in the PR.
    Include a link to the chore file (`.agents/mock-drift-correction.md`) in the PR description.
    Apply the same labels to the PR as are on this issue.
    
    ------------ END ISSUE BODY TEMPLATE ------------