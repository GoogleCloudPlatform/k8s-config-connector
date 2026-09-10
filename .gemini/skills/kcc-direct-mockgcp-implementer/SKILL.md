---
name: kcc-direct-mockgcp-implementer
description: Guides the implementation of Phase 3 (MockGCP and Alignment) for a direct KCC resource, verifying behavioral correctness against simulated GCP services. Use this when you need to implement or align mockgcp for a KCC resource.
---

# KCC Direct MockGCP Implementer

This skill guides you through implementing Phase 3 (MockGCP and Alignment) for a direct KCC resource to verify behavioral correctness against simulated GCP services.

## Inputs
- `ResourceKind`: The kind of the resource (e.g., `VertexAIDataset`).
- `service_name`: The short name of the GCP service (e.g., `aiplatform`).
- `api_version`: The KCC API version (e.g., `v1alpha1`, `v1beta1`).
- `group`: The API group (e.g., `vertexai`).
- `kind_lowercase`: The lowercase kind name (e.g., `vertexaidataset`).
- `testname`: The specific test folder name under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/`.

## Workflow

### 1. Locate E2E Fixtures
- The test fixtures are located under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/`.

### 2. Add or Enhance Mock Service
- If a mock service for `<service_name>` does not exist under `mockgcp/mock<service_name>/`, create one:
  - Follow the guide in `mockgcp/GEMINI.md` and `mockgcp/README.md`.
  - Add the relevant proto to the Makefile and run `make gen-proto` if needed.
  - Implement the mock service entrypoint in `mockgcp/mock<service_name>/service.go` and register it in `mockgcp/register.go`.
- If the mock service already exists, implement the necessary CRUD (Create, Read, Update, Delete) methods for `<ResourceKind>` in `mockgcp/mock<service_name>/<kind_lowercase>.go`.

### 2b. Remove from Ratcheting Exclusions (MANDATORY)
Before running the test cases against real or mock GCP, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which is a fundamental use case KCC resources must support.
1. Open `tests/e2e/ratcheting.go`.
2. Locate the function `ShouldTestRereconiliation`.
3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

### 3. Incremental Mock Alignment

Before executing mock comparison/alignment, you **MUST** discover all test cases/fixtures associated with the target resource.

- **Discover All Test Fixtures for Resource**:
  - Locate all test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` (e.g. `containernodepool`, `containernodepool-update`, `containernodepoolremovetaint`, etc.).

- **Fast vs. Slow Resource Execution Strategy**:
  Determine whether your target resource is fast or slow to reconcile against real GCP. Resources like `ContainerNodePool`, `ContainerCluster`, `SQLInstance`, and `ComposerEnvironment` are slow resources (>10 minutes per test), whereas most others are fast (<10 min per test).
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

  > [!WARNING]
  > **WHENEVER A TEST CASE IS UPDATED, WE MUST RECORD REAL GCP LOGS AGAIN.**
  > If you make any modifications to a test case configuration, manifest files (such as `create.yaml`, `update.yaml`, or `dependencies.yaml`), or the controller's runtime mapping configuration, you **MUST** run the test case against real GCP (`hack/record-gcp` or with `E2E_GCP_TARGET=real`) to regenerate the authentic `_http.log` baseline before comparing or committing any mock log changes. Do not attempt to manually edit the logs or bypass recording live traffic.

- Use the `fix-diffs-mockgcp` skill (`mockgcp/.gemini/skills/fix-diffs-mockgcp/SKILL.md`) to align the mock logs with the real GCP output:
  - **Output-Only Fields/IDs**: If real GCP produces dynamic values that mockgcp lacks, implement a `populate<ResourceKind>Defaults` function in `mockgcp/mock<service_name>/<kind_lowercase>.go` called on `Insert` and `Get` to match the required format.
  - **Volatile/Random Values**: For values like timestamps or etags that are functionally identical but structurally unpredictable, update `normalize.go` for the service.
  - **Critical Rule**: Always scope the `Previsit` normalization in `normalize.go` to ensure it only applies to your service URL (e.g. `strings.Contains(event.URL(), "<service_name>.googleapis.com")`) to prevent log corruption in unrelated services.
- Iterate on running mock alignment commands and making incremental code updates until the HTTP logs match real GCP perfectly with clean, minimal diffs.
- **Run compare-mock**:
  Run the mock comparison tests for all fixtures of the resource in a single command (MockGCP runs in-memory and completes in seconds, so no parallel execution split or Fast/Slow execution strategy is needed for mock tests):
  ```bash
  hack/compare-mock "fixtures/<kind_lowercase>"
  ```
  Specifically, re-run `hack/compare-mock "fixtures/<kind_lowercase>"` (or `hack/compare-mock "fixtures/^<testname>$"`) to update `_http_mock.log` and the golden object files (`_generated_object_*.golden.yaml`), and verify log alignment against `_http.log`. Repeat until all discrepancies are resolved and the tests pass.
  Make sure both `_http.log` and `_http_mock.log` are present in all the fixture directories.

### 4. Verify and Run Presubmits
- Run local validation: `scripts/validate-prereqs.sh`.
- Run the e2e fixtures presubmit: `./dev/ci/presubmits/tests-e2e-fixtures-<kind_lowercase>`.
- **Multi-Fixture Verification Requirement (MANDATORY)**:
  - You **MUST** verify and ensure that all test directories associated with your target resource are completely verified.
  1. Ensure that all discovered test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` have their golden files (`_generated_object_*.golden.yaml`) and mock HTTP logs (`_http_mock.log`) successfully generated. Note that a 0-diff in `git status` is perfectly valid if the fixture was unaffected and produced identical output.
  2. Commit and push all regenerated and matched log/golden files for every discovered fixture in your Pull Request.
