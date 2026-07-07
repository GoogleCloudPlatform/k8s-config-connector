# Skill: Solve Migration Diff Issues for Direct Takeover

This skill guides an automated agent through diagnosing, explaining, and fixing "takeover diff" issues when migrating a Kubernetes controller in KCC from Terraform/DCL to the Direct approach. 

The migration test (`TestMigrationToDirect`) executes in four distinct phases:
* **Phase 1 (Legacy Create):** Provisions the initial resource using the legacy (TF or DCL) controller, recording API traffic to `_http_migration_phase1.log`.
* **Phase 2 (Legacy Re-reconcile):** Applies a no-op annotation (`re-reconcile-legacy-v1`) while the legacy controller is still active to test re-reconciliation behavior and record any unexpected legacy traffic or diffs to `_http_migration_phase2.log`. (Note: if the legacy controller makes no API requests during re-reconciliation, `_http_migration_phase2.log` will not be created or will remain empty).
* **Phase 3 (Direct Takeover):** Applies the `reconciler: direct` annotation to force the Direct controller to adopt the resource. During this takeover, the Direct controller must adopt the resource with a clean **0-write no-op reconciliation**, recording read-only verification traffic to `_http_migration_phase3.log`. Any diff detected during this phase is considered a bug.
* **Phase 4 (Direct Re-reconcile):** Applies another no-op annotation (`re-reconcile-direct-v1`) while the Direct controller is active to verify re-reconciliation behavior. The Direct controller must perform a clean **0-write no-op reconciliation**, recording read-only verification traffic to `_http_migration_phase4.log`.

---

## Steps & Best Practices

### 1. Diagnose the Takeover Diff
1. Run the E2E migration test suite for the specific fixture:
   ```bash
   ./hack/record-gcp TestMigrationToDirect/fixtures/<fixture-name>$
   ```
2. Open the recorded structured diff file:
   `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<resource>/<fixture-name>/_migration_diffs.json`
3. Look for blocks containing `"isNewObject": false`. Note that diffs may appear from legacy re-reconciliation (Phase 2), Direct takeover (Phase 3), or Direct re-reconciliation (Phase 4). Diffs occurring during **Phase 3** or **Phase 4** indicate mismatches or re-reconciliation bugs where the Direct controller attempts to modify fields on an unchanged resource.
4. Analyze the diff. A diff of the form `"old": <value>, "new": null` (or vice-versa) indicates a mismatch where a field is populated in one state but not the other.

---

### 2. Identify the Root Cause
* **Derived/Computed Fields:** In GCP, some fields (e.g., a BigQuery view's `schema`, default database settings, or server-generated metadata) are automatically computed/derived by the server. These are omitted in the KRM spec (desired is `nil`) but populated by GCP (actual is non-nil).
* **Casing & Aliases:** Strings returned by the GCP API might have different casing (e.g. `INT64` vs `INTEGER`, `true` vs `TRUE`) or format (e.g. fully-qualified URIs vs relative paths).
* **Default Values:** The old controller might have applied a default value that the new Direct controller does not apply, or vice-versa.

---

### 3. Formulate the Fix in the Direct Controller
1. Locate the comparison logic for the resource:
   * For most resources, the comparison is done directly inside the `Update` method (or helper functions) in the controller file: `pkg/controller/direct/<service>/<resource>_controller.go`.
   * For extremely large or complex resources (like BigQuery Table), it may be split out into a separate file: `pkg/controller/direct/<service>/<resource>_compare.go`.
2. **Prevent Parameter Swap Bugs:** When writing or editing comparison functions, **always** explicitly name the parameters `actual` and `desired` instead of `a` and `b`. This prevents accidentally swapping them during comparison and diff reporting.
3. **Ignore Undesired Optional Fields:** If a field is optional in KRM and is omitted from the spec (desired is `nil`), the comparison logic should **ignore** the actual value returned by GCP rather than attempting to delete it.
   * *Implementation Pattern:*
     ```go
     func compareFieldEq(actual, desired *Type, prefix string, diff *structuredreporting.Diff) (bool, error) {
         if desired == nil {
             // If the desired state is not specified in the KRM spec, we do not enforce it.
             return true, nil
         }
         if actual == nil {
             // Desired is specified, but actual is nil. This is a diff.
             diff.AddField(prefix, actual, desired)
             return false, nil
         }
         // Perform deep comparison...
     }
     ```
4. **Normalize Values before Comparison:** If the diff is due to formatting or casing differences, implement normalization helpers to format both `actual` and `desired` identically before calling `reflect.DeepEqual`.

---

### 4. Validate the Fix
1. Run the E2E migration test with `WRITE_GOLDEN_OUTPUT=1` to update the recorded traffic and diffs:
   ```bash
   E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock GOLDEN_REQUEST_CHECKS=1 GOLDEN_OBJECT_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 RUN_E2E=1 go test ./tests/e2e -v -run TestMigrationToDirect/fixtures/<fixture-name>$
   ```
2. Verify that the `"isNewObject": false` block in `_migration_diffs.json` is **completely gone**, indicating a clean 0-write takeover.
3. **Iterate If Diffs Persist:** If a diff block with `"isNewObject": false` still exists in `_migration_diffs.json`, analyze the remaining/new diff fields, loop back to **Step 2 (Identify Root Cause)** and **Step 3 (Formulate Fix)**, and refine or add fixes until no takeover diffs remain.
4. Run the E2E migration test one final time without `WRITE_GOLDEN_OUTPUT=1` to ensure the test passes with a perfect green status.

---

### 5. Format and Clean Up
1. Run `make fmt` and `go vet ./pkg/controller/direct/...` to ensure formatting and compilation are perfect.
2. Clean up any accidentally generated untracked `.log` and `.json` files in other fixture directories:
   ```bash
   rm -f pkg/test/resourcefixture/testdata/basic/<service>/v1beta1/<resource>/*/_http_migration_phase*.log
   rm -f pkg/test/resourcefixture/testdata/basic/<service>/v1beta1/<resource>/*/_migration_diffs.json
   ```
