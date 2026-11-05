# Migration Journal: FirebaseProject

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/firebase/v1alpha1/firebaseproject/firebaseprojectautogen`
- Resource: `FirebaseProject` (`firebase.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp "TestMigrationToDirect/fixtures/firebaseprojectautogen$"
  ```
- **Execution Result:**
  - The E2E test `TestMigrationToDirect/fixtures/firebaseprojectautogen` ran against live GCP and completed in ~28s with a PASS status.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully provisioned/imported the `FirebaseProject` using the Terraform controller with read-only GET requests (`_http_migration_phase1_legacy_create.log`).
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller with a no-op annotation; verified clean 0-write re-reconciliation with a single read-only GET request (`_http_migration_phase2_legacy_re-reconciliation.log`).
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero diffs (`"isNewObject": false`) were detected.
- **Phase 4 (Direct Re-reconcile):** Re-reconciled with a no-op annotation under the Direct controller with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs were detected.
- The recorded `_migration_diffs.json` contains `null` (zero diffs across all phases), confirming full state compatibility between the legacy Terraform controller and the direct controller.

### Ratcheting Exclusions
- Verified `FirebaseProject` is not present in the ratcheting exclusions in `tests/e2e/ratcheting.go`.

### Hermetic MockGCP Execution
- Executed `TestMigrationToDirect` against MockGCP:
  - `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run "TestMigrationToDirect/fixtures/firebaseprojectautogen$"`
  - Passed cleanly with 0 unexpected writes or diffs.
- Executed `TestAllInSeries` against MockGCP:
  - `hack/compare-mock pkg/test/resourcefixture/testdata/basic/firebase/v1alpha1/firebaseproject/firebaseprojectautogen`
  - Passed cleanly with 0 unexpected writes or diffs.
