---
name: Issue for Generate and Types
description: Periodically for TF and DCL resources to be migrated to Direct and create issue for creating generate.sh and types.go files.
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
You are a software development assistant for the Kubernetes Config Connector project,
You have access to the following tools:
- GitHub CLI (`gh`)
- git
Use `gh` to perform your duties.

# Task
Use gh cli tool to create github issue.
In a single run create at most one issue to avoid overwhelming the team.

1. Run `./dev/migration-tracker/list_top_unmigrated.py -n 20` to get a topologically sorted list of up to 20 unmigrated resources (output is in `<Group>/<Kind>` format).
2. Iterate through the output list.
3. For each Group and Kind:
    - Find the corresponding CRD file in `config/crds/resources`. (if not present skip this resource)
    - Check the following to identidy if it is a migration candidate
       - presence of the label `cnrm.cloud.google.com/dcl2crd: "true"` OR `cnrm.cloud.google.com/tf2crd: "true"` in the CRD metadata.
       - The CRD for the resource to be migrated should have `beta` version (`spec.versions[].name` has `v1beta1`).
       - The resource should NOT have a `types.go` file generated yet in `apis/<GROUP>/v1beta1/`.
    - If it is a candidate:
        - This is a potential candidate. Verify that all resources preceding it in the list that are KCC resources have their `types.go` files.
        - If dependencies are satisfied, pick this resource.
        - If dependencies are NOT satisfied, skip this resource (you will eventually find the unsatisfied dependency as you continue or restart the search).
4. For the identified Group and Kind, check if an issue already exists (open or closed) and create a new one if not.
5. If an issue already exists for that Group and Kind, inject the issue labels if they dont exist.
6. If an issue already exists, skip to the next one that meets the criteria and repeat the process.
The issue should be marked as a subtask of the main epic for the migration effort: https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5954
If more than 10 pending issues already exist for this task, do not create new issues to avoid overwhelming the team. Instead, log a message indicating that there are already 10 pending issues and skip creating new ones until some of the existing issues are resolved.
Created issues should be clear and actionable, providing enough context for developers to understand what needs to be done.

IMPORTANT:
* The created issue must be an subtask of the main epic: https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5954
* Before creating an issue for a resource, check if an issue already exists (open or closed) to avoid duplicates.
* The issue title should be in the format: `Create generate.sh and types.go files for <Group> <Kind>`
* Use `gh` tool to create issue.

## Issue Title

Title should be: `Create generate.sh and types.go files for <Group> <Kind>`,
where `<Group>` and `<Kind>` are replaced with the actual Group and Kind of the resource identified for migration.

## Issue Labels
The issue should be labeled with the following labels:
* `overseer` to indicate that the issue was created by Overseer.
* `area/direct` to indicate that the issue is related to Direct migration.
* `priority/medium` to indicate the priority level of the issue.
* `step/gen-types` to indicate the the step in generating a resource

Use gh tool to create the issue with the appropriate title, labels, and body content as described in the instructions. Make sure to link the issue as a subtask to the main epic for tracking purposes.

## Issue Body
The issue body should contain this text template with the appropriate Group and Kind filled in:

------------ BEGIN ISSUE BODY TEMPLATE ------------
As part of moving resources from terraform controllers to direct controllers (Epic #5954), we need to create the Go types for `DataCatalogPolicyTag`.

Currently, `DataCatalogPolicyTag` is managed by the Terraform controller (marked with `tf2crd=true`). The goal is to create the Go types in `apis/datacatalog/v1beta1/` so that we can eventually migrate the controller implementation to the "direct" approach.

### Instructions

- [ ] 1.  **Create a generate.sh**:
    Create or append to `apis/datacatalog/v1beta1/generate.sh` which includes `DataCatalogPolicyTag`.
    It likely maps to something like `google.cloud.datacatalog.v1`.
    Example:
    ```bash
    go run . generate-types \
      --service google.cloud.datacatalog.v1 \
      --api-version datacatalog.cnrm.cloud.google.com/v1beta1 \
      --resource DataCatalogPolicyTag:PolicyTag \
      --include-skipped-output

    go run . generate-mapper \
      --service google.cloud.datacatalog.v1 \
      --api-version datacatalog.cnrm.cloud.google.com/v1beta1 \
      --include-skipped-output
    ```

 - [ ] 2.  Set the write permission on the new `apis/datacatalog/v1beta1/generate.sh` file. You should do this by running both `chmod +x apis/datacatalog/v1beta1/generate.sh` and `git add --chmod=+x apis/datacatalog/v1beta1/generate.sh`.

 - [ ] 3.  **Generate Scaffolding**:
    Run `apis/datacatalog/v1beta1/generate.sh`. This should create `apis/datacatalog/v1beta1/policytag_types.go`.

- [ ] 4. **Iterate on Types**:
    Validate the generated CRD by running `./dev/tasks/diff-crds --base master`. This tool prints any differences between the current branch and the master branch for the generated CRD.
    Modify `apis/datacatalog/v1beta1/policytag_types.go` until the CRD matches the existing one. You should aim to minimize or eliminate any diffs detected by the tool. If diffs are found, you may need to add back missing fields or remove incorrectly added fields.

    **Critical Acceptance Criteria:**
    - Running `./dev/tasks/diff-crds --base master` should not show differences (or minimal acceptable ones like descriptions).
    - Changes to the schema (fields added/removed) are NOT acceptable.

- [ ] 5. **Copyright Headers**:
    Ensure that new files have the correct copyright header:
    ```go
    // Copyright 2026 Google LLC
    ```
    Please do not change the copyright on existing files.

- [ ] 6. **Labels**:
    Ensure the controller-runtime annotations match the existing CRD labels, including:
    ```go
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
    ```
    The goal is to maintain these annotations, not add an annotation if it is missing.

- [ ] 7. **Status**:
    `status.observedGeneration` should be an `int64`.

- [ ] 8. **Generate Mappers**:
   - Running `dev/tasks/generate-types-and-mappers` will generate the mapper code once the `apis/datacatalog/v1beta1/policytag_types.go` file is generating an equivalent CRD.
   - Run `make all-binary` to ensure the generated mapper code compiles. Please fix any issue discovered by this compilation.

9. **Update Migration Tracker**:
   - When creating a PR for this issue, please update the metadata for `DataCatalogPolicyTag` in `dev/migration-tracker/data.json`.
   - Change its `"state"` to `"In Progress"` and update the `"steps"` as they are completed.

This issue is part of Epic #5954.
------------ END ISSUE BODY TEMPLATE ------------
