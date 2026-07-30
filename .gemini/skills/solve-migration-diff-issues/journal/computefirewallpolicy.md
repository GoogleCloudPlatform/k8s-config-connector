# Migration Journal: ComputeFirewallPolicy

## Takeover Diff Issues Identified and Solved

During the direct migration validation for `ComputeFirewallPolicy`, the Direct controller adopted the resource from the legacy DCL controller with a perfect **0-write no-op takeover** on both mock and real GCP targets. No schema or behavior discrepancies were found between the two controller reconciliation engines.

However, we identified and solved an E2E test-suite normalization issue:

1. **`org-` prefix Compute Operation ID normalization discrepancy:**
   - **Symptom:** The E2E migration test reported a difference in `_http_migration_phase1_legacy_create.log` on the operation `"name"` and `"selfLink"` fields (e.g., `org-${folderID}-<id>` vs `${operationID}`).
   - **Root Cause:** In `tests/e2e/normalize_legacy.go`, the original operation ID `org-<folder_id>-<id>` was registered for string replacement. However, because `NormalizeHTTPLog` ran first and replaced the numeric `<folder_id>` with `${folderID}`, the string-replace normalizer could no longer find the original operation ID string on the formatted logs, leaving it partially un-normalized.
   - **Solution:** Updated `LegacyNormalize` inside `tests/e2e/normalize_legacy.go` to extract `org-` prefix operation IDs and also register their normalized variants (replacing folder or organization IDs with `${folderID}` or `${organizationID}`) as target operation IDs. This ensures the operation IDs are successfully substituted with `${operationID}` regardless of whether the folder/org ID is already replaced.

With this normalization improvement, the `TestMigrationToDirect` E2E test passes perfectly.
