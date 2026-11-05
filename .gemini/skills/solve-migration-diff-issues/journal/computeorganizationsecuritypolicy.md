# Migration Journal: ComputeOrganizationSecurityPolicy

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/compute/v1alpha1/computeorganizationsecuritypolicy/organizationsecuritypolicyautogen`
- Resource: `ComputeOrganizationSecurityPolicy` (`compute.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp TestMigrationToDirect/fixtures/organizationsecuritypolicyautogen$
  ```
- **Execution Result:**
  - The E2E migration test `TestMigrationToDirect/fixtures/organizationsecuritypolicyautogen` ran against live GCP and passed cleanly in ~66s.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully provisioned the `ComputeOrganizationSecurityPolicy` resource using the Terraform controller. API traffic recorded in `_http_migration_phase1_legacy_create.log`.
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller with a no-op annotation; no unexpected writes or modifications occurred (`_http_migration_phase2_legacy_re-reconciliation.log`).
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero diffs (`"isNewObject": false`) were detected.
- **Phase 4 (Direct Re-reconcile):** Re-reconciled with a no-op annotation under the Direct controller with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs were detected.
- The recorded `_migration_diffs.json` confirmed that there are zero `"isNewObject": false` diff blocks for `ComputeOrganizationSecurityPolicy`. Full state compatibility is maintained between the legacy Terraform controller and the direct controller.

### Ratcheting Exclusions
- Verified `ComputeOrganizationSecurityPolicy` is not excluded in `tests/e2e/ratcheting.go`.
