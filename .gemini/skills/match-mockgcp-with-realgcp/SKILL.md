# Skill: Match Mock behaviour with real GCP API

This skill provides a structured workflow for aligning the MockGCP implementation and normalizers with authentic real GCP HTTP logs (`_http.log`).

When the golden tests for K8s Config Connector mock output diverge from real GCP output, we inspect the discrepancies and fix either the mock service handler or the normalizers. This ensures that `hack/compare-mock` runs cleanly, golden objects are updated, and CI presubmits pass hermetically.

## Critical Rules

*   **Real GCP Baseline Required**: You must always have authentic `_http.log` files recorded from real GCP before aligning mock behavior. Please always execute the [`record-real-gcp`](.gemini/skills/record-real-gcp/SKILL.md) skill.
*   **Basic Test Case Naming**: Ensure that the basic test case directory/folder names always contain the resource kind name (e.g., use `computesubnetwork-reservedinternalrange` instead of `reservedinternalrange`).
*   **No New Missing Fields**: We should not add new missing fields to the exceptions file (`tests/apichecks/testdata/exceptions/missingfields.txt`). If fixing or renaming an existing test case exposes/causes more missing fields, you must extend the test case configuration or add a new test case (or more than one if needed) to cover those missing fields.
 You MUST first establish a baseline against `real` GCP with `E2E_GCP_TARGET=real E2E_KUBE_TARGET=envtest` with `WRITE_GOLDEN_OUTPUT=1`.
*   **WHENEVER A TEST CASE IS UPDATED, WE MUST RECORD REAL GCP LOGS AGAIN**: If you make any modifications to a test case configuration, manifest files (such as `create.yaml`, `update.yaml`, or `dependencies.yaml`), or the controller's runtime mapping configuration, you **MUST** execute the [`record-real-gcp`](.gemini/skills/record-real-gcp/SKILL.md) skill.
*   **Do NOT Update Legacy Controller**: Do not update or modify the legacy/Terraform-based controller (such as code under `pkg/krmtotf/references.go`) unless absolutely necessary. Keep the legacy controller logic unchanged and avoid modifications to common or shared legacy framework code.
*   **Focus on Target CRUD Requests**: When aligning the HTTP logs, we only care about the requests made to the CRUD endpoint of the target resource. Non-target/dependent resource requests in the log do not need to be matched as strictly.

## Workflow

### Step 0: Remove from Ratcheting Exclusions (MANDATORY)

Before running the test cases against real or mock GCP, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which is a fundamental use case KCC resources must support.

1. Open `tests/e2e/ratcheting.go`.
2. Locate the function `ShouldTestRereconiliation`.
3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

---

### Step 1: Establish Baseline with Real GCP

Before aligning mock behavior, ensure authentic `_http.log` files are recorded and verified by following the [`record-real-gcp`](.gemini/skills/record-real-gcp/SKILL.md) skill.

---

### Step 2: Match MockGCP Behaviour with Real GCP

1.  **Run compare-mock**:
    - **If `all: true` (Whole-Kind change)**: Run all tests for the Kind in a single command:
      ```bash
      hack/compare-mock "fixtures/<kind_lowercase>"
      ```
    - **If `all: false` (Selective manifest / reference change)**: Run only the affected test case(s):
      ```bash
      hack/compare-mock "fixtures/^<testname>$"
      ```

    This executes tests against MockGCP and automatically generates/updates:
    - Mock HTTP logs: `_http_mock.log`
    - Golden object snapshots: `_generated_object_*.golden.yaml`

2.  **Verify Alignment & Inspect Discrepancies**:
    Run `TestGoldenLogAlignment` (or inspect diff between `_http.log` and `_http_mock.log`) to identify discrepancies:
    ```bash
    go test -v ./pkg/test/goldenalignment -run "TestGoldenLogAlignment/<fixture_name>"
    ```

3.  **Iterative Resolution (Exit Criteria)**:
    Apply the appropriate resolution recipe below to MockGCP handlers or scoped normalizers. Re-run `hack/compare-mock` and `TestGoldenLogAlignment` until the alignment test **passes with 0 errors and zero unintended diffs**.

#### Tips for Fixing Discrepancies:

1. **Missing Default Values (Field in Real but NOT in Mock)**:
   - For missing default values (e.g. `<missing> => REGIONAL`), add or update `populateDefaultsFor<Resource>` in `mockgcp/mock<service>/<kind>.go`. Ensure it is called on both mutation (`Create`/`Insert`) and retrieval (`Get`/`List`) methods.
   - If the field is **missing from the proto**, leave a `TODO` comment noting the missing proto field instead.
2. **Unexpected Values (Field in Mock but NOT in Real)**:
   - Check if the field was added to align with another test case. If real GCP includes it in other scenarios but omits it here, it is possibly **conditionally defaulted** (or was accidentally hardcoded).
   - If verified to be conditional, implement dynamic/conditional defaulting in MockGCP and add a comment explaining the condition. Otherwise, remove or fix the hardcoded default.
   - If the field is never returned by real GCP in any scenario, remove/omit it from MockGCP.
3. **Empty / Zero-Value Fields in Real**: If real GCP outputs an empty value (e.g. `""`, `0`, `false`, `[]`) that MockGCP drops due to proto3 serialization (unset vs zero-value equivalence):
   - Try explicitly setting the field in MockGCP.
   - If proto3 JSON marshaling still omits it, leave a `TODO` comment clarifying that it is an expected proto3 zero-value limitation.
4. **Input vs Output Canonicalization**:
   - When real GCP accepts short names or unformatted references but returns canonicalized formats (e.g. converting `/zones/` to `/locations/`, expanding relative names to full URIs, or normalizing CIDRs), implement the transformation in MockGCP during `Create`/`Insert` so the stored state matches GCP's canonical output.
5. **Volatile / Dynamic Values in Self-Links & Long Names**:
   - For dynamic segments embedded in resource names, self-links, or generated resources (e.g., auto-generated hashes in `/instanceGroupManagers/gke-.*-grp`, `/forwardingRules/${forwardingRuleID}`, `${uniqueId}`, `${projectNumber}`), normalize them to standard placeholders in `mockgcp/mock<service>/normalize.go`.
   - **Critical Rule**: Always scope `Previsit` normalization in `normalize.go` to the service URL (e.g. `strings.Contains(event.URL(), "<service>.googleapis.com")`).
6. **Exempt Specific URL Prefixes**:
   - If real GCP returns full URL prefixes (e.g. `//certificatemanager.googleapis.com/` or `https://www.googleapis.com/...`) while MockGCP uses relative paths, strip the prefix in the service's normalizer, keeping the rule strictly scoped to that specific resource/field.
7. **Exempt API Version Differences (Specific / Hardcoded Only)**:
   - If real GCP and MockGCP use different API versions for a specific service (e.g., `aiplatform.v1beta1` vs `aiplatform.v1`), normalize it using explicit, hardcoded string replacements for that specific service only. Never use broad wildcards that could mask unintended API version mismatches.
8. **Operation Metadata**: Minor operation polling metadata (e.g. intermediate LRO statuses like `done: false`) can be simplified in MockGCP as long as the final resource state matches.
9. **Rule of Incremental Fixes**: Fix discrepancies **for one field** at a time and verify immediately with `hack/compare-mock` and `TestGoldenLogAlignment`. Revert (`git reset --hard`) if you get stuck in a loop.

### Step 3: Multi-Fixture Verification & Golden Alignment Test (Mandatory)

Verify that all test directories associated with your target resource are verified:
1. Ensure all discovered test subdirectories have their golden files (`_generated_object_*.golden.yaml`) and mock HTTP logs (`_http_mock.log`) cleanly generated.
   *(Note: A 0-diff in git status is valid if the resource contains no non-deterministic server-generated values or normalizers fully masked them).*
2. Confirm that `hack/compare-mock` executes with zero test failures across all resolved fixtures before concluding mock alignment.
3. **Run Golden Log Alignment Test**: Confirm that all resolved fixtures pass the golden log alignment test with zero errors:
   ```bash
   go test -v ./pkg/test/goldenalignment -run "TestGoldenLogAlignment/<kind_lowercase>"
   ```
