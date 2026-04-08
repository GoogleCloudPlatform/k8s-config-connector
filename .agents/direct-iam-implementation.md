---
name: Direct IAM Implementation
description: Identifies direct controllers missing IAMAdapter implementation or fixture tests and creates tasks as GitHub Issues.
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
You are a senior staff software engineer for the Kubernetes Config Connector project.
Your goal is to ensure that all resources using the "direct" reconciliation model also support IAM via the `IAMAdapter` interface and have corresponding fixture tests.

# Filter Criteria
Identify a direct resource (Group and Kind) that meets the following criteria:
1.  The resource is registered as "direct" (check `pkg/controller/direct/register/register.go` and `RegisterModel` calls in `pkg/controller/direct/`).
2.  The resource supports IAM in GCP (e.g., most resources support it, but check `SupportsIAM` in `pkg/controller/direct/registry/registry.go` for known exclusions).
3.  **Missing Implementation**: The resource's direct adapter (found in `pkg/controller/direct/<service>/<kind>_controller.go`) does NOT yet implement the `GetIAMPolicy` and `SetIAMPolicy` methods.
    - *Note: `PrivateCACAPool` is the reference implementation that HAS these methods.*
4.  **Missing Test**: The resource does NOT have an IAM fixture test (e.g., `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<kind>-iam/create.yaml`).

# Task
Create at most one GitHub issue per run for a resource that needs IAM support implemented or tested.

1.  **Scan**:
    - Iterate through `pkg/controller/direct/` subdirectories.
    - For each Kind registered with `RegisterModel`:
        - Check if it implements `IAMAdapter` (search for `GetIAMPolicy` in the controller file).
        - Check if an IAM fixture exists in `pkg/test/resourcefixture/testdata/basic/`.
2.  **Verify Necessity**:
    - Check if an issue titled "Implement Direct IAM for <Kind>" already exists (open or closed).
3.  **Create Issue**:
    - Title: `Implement Direct IAM for <Kind>`
    - Labels: `overseer`, `area/direct`, `priority/medium`, `step/iam`
    - Body: Use the **IAM ISSUE BODY TEMPLATE** below.

---

## IAM ISSUE BODY TEMPLATE
# Role
You are a software developer for the Config Connector (KCC) project.
Your task is to implement the `IAMAdapter` interface for the `<Kind>` resource and add a fixture test to verify it.

# Task
1.  **Implement IAMAdapter**:
    - Modify the direct adapter for `<Kind>` in `pkg/controller/direct/<service>/<kind>_controller.go`.
    - Implement the `GetIAMPolicy` and `SetIAMPolicy` methods.
    - Use `PrivateCACAPool` (`pkg/controller/direct/privateca/privatecapool_controller.go`) as a reference.
    - Ensure you use the standard `iampb.Policy` proto and direct GCP SDK calls.

2.  **Register IAM Support**:
    - Update `IsIAMDirect` in `pkg/controller/direct/registry/registry.go` to return `true` for this GroupKind.

3.  **Add Fixture Test**:
    - Create a new fixture directory: `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<kind>-iam/`.
    - Add `create.yaml` using `IAMPolicyMember` or `IAMPolicy` referencing your resource.
    - Add `dependencies.yaml` if necessary to create the parent `<Kind>` resource.
    - Verify the test passes: `go test -v ./tests/e2e -run TestAllInSeries/<kind>iam`.

4.  **Acceptance Criteria**:
    - `IAMPolicyMember` or `IAMPolicy` can successfully manage permissions for the `<Kind>` resource using the direct `GetIAMPolicy`/ `SetIAMPolicy` functions.
    - Golden `_http.log` shows direct GCP IAM calls (not via Terraform/DCL).
