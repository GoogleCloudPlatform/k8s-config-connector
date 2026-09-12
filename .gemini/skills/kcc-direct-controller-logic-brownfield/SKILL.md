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
    - **Labels in Brownfield Direct Controllers**: If the brownfield resource supports GCP labels, see the `kcc-direct-brownfield-labels` skill to properly map and reconcile labels using `metadata.labels` instead of `spec.labels`, preventing empty payloads, update mask issues, and untested fixtures.

1.5. **Remove from Ratcheting Exclusions (MANDATORY)**:
    This controller logic phase is the **primary stage** where the target resource must be removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. Removing the exclusion activates Server-Side Apply (SSA) creation and re-reconciliation verification within the test runner. This is a fundamental KCC capability that ensures direct controllers support 0-write re-reconciliation.

    To remove the resource from the exclusions:
    1. Open `tests/e2e/ratcheting.go`.
    2. Locate the function `ShouldTestRereconiliation`.
    3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
    4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

    After removing the exclusion, you **MUST** validate and re-record the standard fixtures (`TestAllInSeries`) against live GCP to verify 0-write re-reconciliation:
    
    - **Discover All Test Fixtures for Resource**:
      - Locate all test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` (e.g. `containernodepool`, `containernodepool-update`, `containernodepoolremovetaint`, etc.).

    - **Fast vs. Slow Resource Execution Strategy**:
      Determine whether your target resource is fast or slow to reconcile. Resources like `ContainerNodePool`, `ContainerCluster`, `SQLInstance`, and `ComposerEnvironment` are slow resources (>10 minutes per test), whereas most others are fast (<10 min per test).
      - **Fast resources (<10 min per test)**:
        - Run and record all fixtures for the resource in a single command:
          ```bash
          hack/record-gcp "fixtures/<kind_lowercase>"
          ```
      - **Slow resources (>10 min per test)**:
        - Do NOT run all test fixtures sequentially in a single command.
        - Execute each individual test fixture in parallel (e.g., triggering parallel background processes or subagents with `hack/record-gcp "fixtures/^<testname>$"`).
        - **Parallel Directory Collisions**: `hack/record-gcp` cleans and writes to `artifactz/realgcp` by default. When running multiple invocations in parallel, you **MUST** provide a unique `ARTIFACTS` path for each invocation (e.g., `ARTIFACTS=artifactz/realgcp-<testname> hack/record-gcp "fixtures/^<testname>$"`) to prevent folder collision/cleanup issues.

    - **Verification Criteria for Real GCP Recording (MANDATORY)**:
      Checking `git status` or top-level `PASS`/`FAIL` exit codes can be misleading (e.g. 0-diff is valid for unaffected fixtures, unmatched regex gives a false `PASS`, and a successful recording can mark the subtest as `FAIL` because it wrote a new golden log).
      Instead, verify that each discovered fixture completed its full lifecycle by checking stdout for:
      1. `=== RUN   TestAllInSeries/fixtures/<testname>` (confirms the subtest matched and started).
      2. Resource reached `Ready` (`status.condition.status: True`) for create (and update, if present).
      3. Deletion completed (`Done waiting for resource to delete`).
      4. `wrote updated golden output to .../_http.log` was logged (or `_http.log` was updated/written on disk).
      5. No fatal errors (`t.Fatalf` / timeout / permission errors) aborted the test before cleanup.
        
    If there are unexpected updates or writes in `_http.log`, diagnose and resolve these re-reconciliation diff bugs.

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
        Run the mock comparison tests for all fixtures of the resource in a single command (MockGCP runs in-memory and completes in seconds, so no parallel execution split is needed):
        ```bash
        hack/compare-mock "fixtures/<kind_lowercase>"
        ```
        This command will execute the tests against MockGCP and automatically generate/update the golden files (such as `_generated_object_<testname>-direct.golden.yaml`) and `_http_mock.log` under the test case directories.

    *   **Step 2.b: Multi-Fixture Verification Requirement (MANDATORY)**
        You **MUST** verify and ensure that all test directories associated with your target resource are completely verified.
        1. Ensure that all discovered test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` have their golden files (`_generated_object_*.golden.yaml`) and mock HTTP logs (`_http_mock.log`) successfully generated. Note that a 0-diff in `git status` is perfectly valid if the resource doesn't contain any service-generated, randomized values in the responses or our normalizers have overwritten those values.
        2. Commit and push all regenerated and matched log/golden files for every discovered fixture.
        
    *   **Step 2.c: Fix discrepancies in the Adapter**
        If there are any errors or discrepancies between the legacy reconciler and the direct reconciler behavior (indicated by the test failing or having incorrect updates in `_http.log` or golden objects), iteratively update your direct controller.
        
        Focus on the **Adapter** implementation:
        - `Find`: Check if fields are correctly read, populated, and mapped from the GCP SDK response.
        - `Create` / `Update`: Ensure fields are correctly mapped to the GCP SDK request and correct APIs are called. Pay attention to differences in Patch/Put calls due to incorrect diffing logic in the Update method.
        - `Delete`: Check that deletion logic works and returns correct status/error code.
        
        If a fix is needed in any place other than the Adapter in the controller code, please carefully asses it and make a well scoped change.

        Run comparison/recording commands again after making updates to verify if the changes resolve the issues. Repeat until the tests pass and the updated golden files and HTTP logs accurately match expectations.

    
## Journaling
Append any reconciliation alignment issues to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
