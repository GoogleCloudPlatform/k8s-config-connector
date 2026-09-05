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
    - **Service-Generated Resource IDs**: If the brownfield resource has a service-generated resource ID, see the `kcc-direct-service-generated-id` skill for mandatory identity and reconciler rules (including `Find()` pre-check guards, `Create()` request handling, and `GetIdentity()` comparison rules).

1.5. **Remove from Ratcheting Exclusions (MANDATORY)**:
    This controller logic phase is the **primary stage** where the target resource must be removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. Removing the exclusion activates Server-Side Apply (SSA) creation and re-reconciliation verification within the test runner. This is a fundamental KCC capability that ensures direct controllers support 0-write re-reconciliation.

    To remove the resource from the exclusions:
    1. Open `tests/e2e/ratcheting.go`.
    2. Locate the function `ShouldTestRereconiliation`.
    3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
    4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

    After removing the exclusion, you **MUST** validate and re-record the standard fixtures (`TestAllInSeries`) against live GCP to verify 0-write re-reconciliation:
    1. Run the recorder: `./hack/record-gcp TestAllInSeries/fixtures/<fixture-name>` (or all fixtures for this kind).
    2. If there are unexpected updates or writes in `_http.log`, diagnose and resolve these re-reconciliation diff bugs.

    **Resolution Steps for Re-reconciliation Diff Bugs**:
    * **Comparing Diffs with `common.CompareBrownfieldSpec` for Server Defaults**: If the server returns default values for fields not specified in KRM, replace the diffing logic with `common.CompareBrownfieldSpec` to adopt the actual server values into the desired state before comparison, preventing spurious update attempts on unspecified fields.
    * **Link Normalization**: If the server returns formatted links (e.g., using relative paths or project numbers instead of project IDs) that differ from what the user specifies in KRM, write normalizers to convert both user-specified and server-returned URIs into a canonical format before comparison.
    * **Slice/List Sorting**: If the server returns elements in a repeated field in a non-deterministic or different order than specified in KRM (for unordered lists), sort the slice elements deterministically in both desired and actual states before comparison.

2.  **Verify and Record against MockGCP / Fix Discrepancies**:
    Run the fixtures tests against mock GCP to check behavior, update golden files, and verify correctness.

    > [!WARNING]
    > **WHENEVER A TEST CASE IS UPDATED, WE MUST RECORD REAL GCP LOGS AGAIN.**
    > If you make any modifications to a test case configuration, manifest files (such as `create.yaml`, `update.yaml`, or `dependencies.yaml`), or the controller's runtime mapping configuration, you **MUST** run the test case against real GCP (`hack/record-gcp` or with `E2E_GCP_TARGET=real`) to regenerate the authentic `_http.log` baseline before comparing or committing any mock log changes. Do not attempt to manually edit the logs or bypass recording live traffic.

    When the direct controller becomes available for a resource, the test framework automatically injects the 'direct' 
    annotation into the test YAML and generates golden logs using the direct controller. If both the legacy and direct 
    controllers are available, the framework also records the legacy controller logs (_http_old_controller.log, 
    _final_object_old_controller.golden.yaml, and _exported_old_controller.golden.yaml) along with the diff files between 
    the two controllers (_final_object.diff and _http.diff).
    
    *   **Step 2.a: Run compare-mock**
        Run the test using `hack/compare-mock`:
        ```bash
        hack/compare-mock "fixtures/^<testname>$"
        ```
        This command will execute the tests against MockGCP and automatically generate/update the golden files (such as `_generated_object_<testname>-direct.golden.yaml`) and `_http.log` under the test case directory.
        
    *   **Step 2.b: Fix discrepancies in the Adapter**
        If there are any errors or discrepancies between the legacy reconciler and the direct reconciler behavior (indicated by the test failing or having incorrect updates in `_http.log` or golden objects), iteratively update your direct controller.
        
        Focus on the **Adapter** implementation:
        - `Find`: Check if fields are correctly read, populated, and mapped from the GCP SDK response.
        - `Create` / `Update`: Ensure fields are correctly mapped to the GCP SDK request and correct APIs are called. Pay attention to differences in Patch/Put calls due to incorrect diffing logic in the Update method.
        - `Delete`: Check that deletion logic works and returns correct status/error code.
        
        If a fix is needed in any place other than the Adapter in the controller code, please carefully asses it and make a well scoped change.

        Run the `compare-mock` command again after making updates to verify if the changes resolve the issues. Repeat until the tests pass and the updated golden files and HTTP logs accurately match expectations.

    
## Journaling
Append any reconciliation alignment issues to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.

## GCP Labels Handling for Migrated Resources (Brownfield)

To maintain backward compatibility during brownfield migrations, standard GCP labels must be mapped to/from Kubernetes `metadata.labels`. The `labels` field must not be present in `spec`.

Follow this template:

1. **Remove labels from Spec**: In `apis/<service>/<version>/<resource>_types.go`, ensure that `Labels` is not in `MyResourceSpec`. Comment it out if present and run `dev/tasks/generate-crds`.
2. **Configure Fuzzer**: In `pkg/controller/direct/<service>/<resource>_fuzzer.go`, call `f.Unimplemented_LabelsAnnotations(".labels")`.
3. **Map Labels in Controller**: In `pkg/controller/direct/<service>/<resource>_controller.go`, use the `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label` package to set the labels:
   ```go
   desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())
   ```
4. **Create Labels Test Fixture**: In `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<resource_lowercase>-labels/`, define a dedicated test case that modifies `metadata.labels` between `create.yaml` and `update.yaml`.

