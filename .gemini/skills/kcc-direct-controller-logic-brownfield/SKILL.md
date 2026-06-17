---
name: kcc-direct-controller-logic-brownfield
description: Implement the core reconciliation logic (Adapter) and E2E fixtures for a direct controller migrating from legacy
---

# KCC Direct Controller Logic Implementer for Brownfield resource

This skill guides the implementation of the `Adapter` interface

## Inputs
- `resource_kind`: The KCC Kind.
- `service_name`: The GCP service name (short, e.g., `apigee`).
- `api_version`: The KCC API version.

## Workflow

1.  **Implement Adapter Logic**:
    Update `pkg/controller/direct/<service>/<resource_lower>_controller.go`.
    - Implement `Find`, `Create`, `Update`, and `Delete`.
    - Use the generated mappers and manual mappers as needed.
    - Ensure correct error handling (e.g., handling 404s in `Find`).

2.  **Duplicate and Annotate Test Cases for Direct Reconciler**:
    We need to duplicate the existing test fixtures and configure the duplicates to use the direct reconciler.
    
    *   **Step 2.a: Duplicate test cases**
        Locate the existing test cases under `pkg/test/resourcefixture/testdata/basic/<service_name>/<version>/<kind>`.
        For each test case directory (e.g., `<test_name>`), create a duplicate copy of the entire directory in the same folder, appending `-direct` to its name (e.g., `<test_name>-direct`).
        
        *Example*: For `bigtableappprofile` under `pkg/test/resourcefixture/testdata/basic/bigtable/v1beta1/bigtableappprofile`, the directory is duplicated as `bigtableappprofile-direct`.
        
        Commit this duplication change to git before modifying the files.
        
    *   **Step 2.b: Inject direct reconciler annotation**
        In the new `-direct` test directories, edit the resource YAML files (typically `create.yaml` and `update.yaml`, if present).
        Add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation under `metadata.annotations`.
        
        *Example configuration change in `create.yaml` / `update.yaml`*:
        ```yaml
        metadata:
          annotations:
            alpha.cnrm.cloud.google.com/reconciler: direct
        ```
        
        Do not add this annotation to dependency resources defined in `dependencies.yaml` unless they are also specifically part of the migration.


3.  **Verify and Record against MockGCP / Fix Discrepancies**:
    Run the newly added duplicate tests against mock GCP to check behavior, update golden files, and verify correctness.
    
    *   **Step 3.a: Run compare-mock**
        Run the test using `hack/compare-mock`:
        ```bash
        hack/compare-mock "fixtures/^<testname>-direct$"
        ```
        This command will execute the tests against MockGCP and automatically generate/update the golden files (such as `_generated_object_<testname>-direct.golden.yaml`) and `_http.log` under the test case directory.
        
    *   **Step 3.b: Fix discrepancies in the Adapter**
        If there are any errors or discrepancies between the legacy reconciler and the direct reconciler behavior (indicated by the test failing or having incorrect updates in `_http.log` or golden objects), iteratively update your direct controller.
        
        Focus on the **Adapter** implementation:
        - `Find`: Check if fields are correctly read, populated, and mapped from the GCP SDK response.
        - `Create` / `Update`: Ensure fields are correctly mapped to the GCP SDK request and correct APIs are called. Pay attention to differences in Patch/Put calls due to incorrect diffing logic in the Update method.
        - `Delete`: Check that deletion logic works and returns correct status/error code.
        
        If a fix is needed in any place other than the Adapter in the controller code, please carefully asses it and make a well scoped change.

        Run the `compare-mock` command again after making updates to verify if the changes resolve the issues. Repeat until the tests pass and the updated golden files and HTTP logs accurately match expectations.

    
## Journaling
Append any reconciliation alignment issues to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
