---
name: Issue for mockgcp
description: Periodically find TF and DCL resources to be migrated to Direct and create an issue to create the mockgcp implementation.
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
You have access to the following tools:
- GitHub CLI (`gh`)
- git
Use `gh` to perform your duties.

# Filter Criteria
5. Find the e2e test fixtures for this resource. Tests are usually in `pkg/test/resourcefixture/testdata/basic/<group>/<version>/<kind>/` or similar locations. If an e2e test DOES NOT exist for this resource, record the response (log that the test doesn't exist) and SKIP the resource. If a test DOES exist and the mock does NOT, pick this resource.


# Task
Use gh cli tool to create a github issue.
In a single run create at most one issue to avoid overwhelming the team.

1. Run `./dev/migration-tracker/list_top_unmigrated.py -n 20 --step mocks` to get the prioritized list of resource Groups and Kinds in <Group>/<Kind> format
2. Iterate through the output list.
3. For each Group and Kind:
    - Find the corresponding CRD file in `config/crds/resources`. If not found skip it.
    -  Candidate resources can be identified by the presence of the label `cnrm.cloud.google.com/dcl2crd: "true"` OR `cnrm.cloud.google.com/tf2crd: "true"` in their CRD metadata.
    - The resource to be migrated should be in `beta` (`spec.versions[].name` is `v1beta1`).
    - Check if it is a migration candidate (DCL or TF, Beta version).
    - Check if the mock already exists in `mockgcp/`.
    - Find the e2e test fixtures for this resource. If an e2e test DOES NOT exist, SKIP the resource.
    - If a candidate is valid and a test DOES exist but the mock does NOT:
        - Check if an issue already exists (open or closed) for this Group and Kind and create a new one if not.
        - If an issue already exists for that Group and Kind, inject the issue labels if they don't exist using `gh issue edit`.
        - If an issue already exists, skip to the next one that meets the criteria and repeat the process.
4. The issue should be marked as a subtask of the main epic for the migration effort: https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5954
5. If more than 10 pending issues already exist for this task (labeled with `step/mocks`), do not create new issues to avoid overwhelming the team. Instead, log a message indicating that there are already pending issues and skip creating new ones.
6. Created issues should be clear and actionable, providing enough context for developers to understand what needs to be done.

IMPORTANT:
* The created issue must be a subtask of the main epic: https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5954
* Before creating an issue for a resource, check if an issue already exists (open or closed) to avoid duplicates.
* The issue title should be in the format: `Create mockgcp implementation for <Group> <Kind>`
* Use `gh` tool to create issue.

## Issue Title

Title should be: `Create mockgcp implementation for <Group> <Kind>`,
where `<Group>` and `<Kind>` are replaced with the actual Group and Kind of the resource identified for mock creation.

## Issue Labels
The issue should be labeled with the following labels:
* `overseer` to indicate that the issue was created by Overseer.
* `area/direct` to indicate that the issue is related to Direct migration.
* `priority/medium` to indicate the priority level of the issue.
* `step/mocks` to indicate the step in the migration process.

Use the gh tool to create the issue with the appropriate title, labels, and body content as described in the instructions. Make sure to link the issue as a subtask to the main epic for tracking purposes.

## Issue Body
The issue body should contain this text template with the appropriate Group, Kind, and lowercase Kind filled in.
The body template is treated as markdown. Retain the formatting as is when filling in the values.

------------ BEGIN ISSUE BODY TEMPLATE ------------
As part of moving resources from terraform and DCL controllers to direct controllers (Epic #5954), we need to create the `mockgcp` implementation for `<Group> <Kind>`.

Currently, `<Group> <Kind>` is managed by the Terraform or DCL controller. The goal is to create its mock implementation in `mockgcp/` to enable hermetic testing, which is a prerequisite for migrating the controller implementation to the "direct" approach.

Instructions:

- [ ] 1. **Capture Golden Traffic**:
    Run the existing terraform/DCL controller to grab the httplogs against real GCP.
    Use the following command to run the tests against real GCP and write the golden output (replace `<lowercase_kind>` with the resource's kind in lowercase):
    ```bash
    E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=real GOLDEN_OBJECT_CHECKS=1 GOLDEN_REQUEST_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/<lowercase_kind> 2>&1 | tee log
    ```

- [ ] 2. **Add the Proto Definition**:
    Locate the relevant `.proto` file for the GCP service in the `googleapis` repository and add it to `mockgcp/Makefile`.

- [ ] 3. **Generate Go Code**:
    Run `make all` in `mockgcp/` to generate the necessary Go files from the proto definition.

- [ ] 4. **Create a Mock Service Directory**:
    Create a new directory named `mockgcp/mock<servicename>`.

- [ ] 5. **Implement the basic service**:
    In your service directory, create `service.go` and `normalize.go`. Register the service in `mockgcp/register.go` for both GRPC and HTTP.

- [ ] 6. **Implement the Service**:
    Implement the core CRUD methods in a `.go` file within the new directory based on the `_http.log` output.

- [ ] 7. **Normalization**:
    Ensure you explicitly scope the `Previsit` logic in `normalize.go` to the relevant service URL to avoid corrupting logs in other services.

- [ ] 8. **Verify your mock implementation**:
    Verify the mock implementation by running the tests against it:
    ```bash
    E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock GOLDEN_OBJECT_CHECKS=1 GOLDEN_REQUEST_CHECKS=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/<lowercase_kind> 2>&1 | tee log
    ```
    The tests should pass against your mock implementation without modifications to the `_http.log` files if possible.

- [ ] 9. **Create a PR**:
    Create a branch, commit the new mockgcp implementation, and open a PR.
    The PR title should be: `mockgcp: feat: For <Group> <Kind>`
    In the PR body, link back to this issue.

- [ ] 10. **Update Migration Tracker**:
    When creating a PR for this issue, please update the metadata for `<Group> <Kind>` in `dev/migration-tracker/data.json`.
    Change its `"state"` to `"In Progress"` for this step (or set `"mocks": true` in the `"steps"` object as completed).

This issue is part of Epic #5954.
------------ END ISSUE BODY TEMPLATE ------------