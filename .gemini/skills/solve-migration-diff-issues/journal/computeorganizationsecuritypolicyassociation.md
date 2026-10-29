# Migration Journal: ComputeOrganizationSecurityPolicyAssociation

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/compute/v1alpha1/computeorganizationsecuritypolicyassociation/computeorganizationsecuritypolicyassociationbasic`
- Resource: `ComputeOrganizationSecurityPolicyAssociation` (`compute.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp TestMigrationToDirect/fixtures/computeorganizationsecuritypolicyassociationbasic$
  ```
- **Execution Result:**
  - The E2E migration test `TestMigrationToDirect/fixtures/computeorganizationsecuritypolicyassociationbasic` ran against live GCP and passed cleanly in ~73s.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully provisioned the `ComputeOrganizationSecurityPolicy` dependency and `ComputeOrganizationSecurityPolicyAssociation` resource using the Terraform controller. API traffic recorded in `_http_migration_phase1_legacy_create.log`.
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller with a no-op annotation; only a single read-only GET request occurred (`_http_migration_phase2_legacy_re-reconciliation.log`), with 0 writes or modifications.
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero diffs (`"isNewObject": false`) were detected.
- **Phase 4 (Direct Re-reconcile):** Re-reconciled with a no-op annotation under the Direct controller with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs were detected.
- The recorded `_migration_diffs.json` confirmed that there are zero `"isNewObject": false` diff blocks for `ComputeOrganizationSecurityPolicyAssociation`. Full state compatibility is maintained between the legacy Terraform controller and the direct controller.

### Hermetic Verification
- Also verified hermetically against MockGCP via:
  ```bash
  RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test ./tests/e2e -v -run TestMigrationToDirect/fixtures/computeorganizationsecuritypolicyassociationbasic$
  ```
  and `hack/compare-mock`, which passed with zero discrepancies.

### Ratcheting Exclusions
- Verified `ComputeOrganizationSecurityPolicyAssociation` is not in `tests/e2e/ratcheting.go` (re-reconciliation is already enabled).
