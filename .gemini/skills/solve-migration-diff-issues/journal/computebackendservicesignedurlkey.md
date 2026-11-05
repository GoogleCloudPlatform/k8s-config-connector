# Migration Journal: ComputeBackendServiceSignedURLKey

## Takeover Diff Issues Identified and Solved

During the direct migration validation for `ComputeBackendServiceSignedURLKey`, we verified the transition from the legacy (Terraform) controller to the Direct controller.

1. **Clean 0-Write Takeover:**
   - **Verification:** The Direct controller successfully adopted the resource during Phase 3 (direct takeover) and Phase 4 (direct re-reconciliation) with a clean 0-write no-op reconciliation on real GCP.
   - **Diffs:** No unexpected mutation diffs (`isNewObject: false`) were produced during legacy re-reconciliation, Direct takeover, or Direct re-reconciliation.

2. **Terraform Set Hash Normalization:**
   - **Symptom:** In dependency resources such as `ComputeBackendService`, Terraform represents set items with volatile numerical hashes in diff IDs (e.g., `health_checks.774372592`), which can vary across runs and environments in `_migration_diffs.json`.
   - **Solution:** Added a normalizer in `tests/e2e/migration_test.go` to normalize Terraform list/set hashes (`"id": "<field>.<hash>"`) to `"id": "<field>.${hash}"`, and normalized `health_checks.${hash}` in `_migration_diffs.json`, ensuring golden diff reproducibility across runs.

3. **Re-reconciliation Ratcheting:**
   - Removed `ComputeBackendServiceSignedURLKey` from the `ShouldTestRereconiliation` exclusion list in `tests/e2e/ratcheting.go`, confirming that re-reconciliation passes cleanly.
