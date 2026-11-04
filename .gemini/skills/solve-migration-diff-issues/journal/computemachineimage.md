# Migration Journal: ComputeMachineImage

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/compute/v1alpha1/computemachineimage/computemachineimageautogen`
- Resource: `ComputeMachineImage` (`compute.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp TestMigrationToDirect/fixtures/computemachineimageautogen$
  ```
- **Execution Result:**
  - The E2E test `TestMigrationToDirect/fixtures/computemachineimageautogen` ran against live GCP and passed cleanly.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully created `ComputeNetwork`, `ComputeInstance`, and `ComputeMachineImage` using the legacy Terraform controller (`_http_migration_phase1_legacy_create.log`).
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller with a no-op annotation; only read requests (`GET https://compute.googleapis.com/compute/v1/projects/${projectId}/global/machineImages/...`) were made (`_http_migration_phase2_legacy_re-reconciliation.log`).
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero diffs (`"isNewObject": false`) were detected.
- **Phase 4 (Direct Re-reconcile):** Re-reconciled with a no-op annotation under the Direct controller with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs were detected.
- The recorded `_migration_diffs.json` confirms that there are zero `"isNewObject": false` diff blocks for `ComputeMachineImage`. Full state compatibility is maintained between the legacy Terraform controller and the direct controller using `common.CompareBrownfieldSpec` with `TrimComputeURIPrefix`.

### Ratcheting Exclusions
- `ComputeMachineImage` is not in the ratcheting exclusions list in `tests/e2e/ratcheting.go`.
- Hermetic MockGCP migration validation was also verified (`TestMigrationToDirect/fixtures/computemachineimageautogen`) and passed cleanly.
