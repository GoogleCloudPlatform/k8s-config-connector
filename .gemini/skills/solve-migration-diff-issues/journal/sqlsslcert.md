# Migration Journal: SQLSSLCert

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlsslcert`
- Resource: `SQLSSLCert` (`sql.cnrm.cloud.google.com/v1beta1`)

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp "TestMigrationToDirect/fixtures/sqlsslcert$"
  ```
- **Execution Result:**
  - The E2E test `TestMigrationToDirect/fixtures/sqlsslcert` ran against live GCP and completed in 437.86s with a PASS status.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully provisioned the Cloud SQL instance and client `SQLSSLCert` using the Terraform controller.
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller with a no-op annotation; no unexpected writes or modifications occurred.
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero diffs (`"isNewObject": false`) were detected.
- **Phase 4 (Direct Re-reconcile):** Re-reconciled with a no-op annotation under the Direct controller with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs were detected.
- The recorded `_migration_diffs.json` confirmed that there are zero `"isNewObject": false` diff blocks for `SQLSSLCert`. Full state compatibility is maintained between the legacy Terraform controller and the direct controller.

### Ratcheting Exclusions
- Verified `SQLSSLCert` is not in `tests/e2e/ratcheting.go`.
