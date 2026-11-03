# Migration Journal: AppEngineStandardAppVersion

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/appengine/v1alpha1/appenginestandardappversion/appenginestandardappversionbasic`
- Resource: `AppEngineStandardAppVersion` (`appengine.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp "TestMigrationToDirect/fixtures/appenginestandardappversionbasic$"
  ```
- **Recorded Golden Files:**
  - `_http_migration_phase1_legacy_create.log`
  - `_http_migration_phase2_legacy_re-reconciliation.log`
  - `_http_migration_phase3_direct_takeover.log`
  - `_http_migration_phase4_direct_re-reconciliation.log`
  - `_migration_diffs.json`

### Harness Shortcoming & Resolution
- **Issue:** During Phase 2 legacy re-reconciliation, the legacy Terraform provider detects diffs on `deployment.0.files` and `env_variables.PORT` because `Read()` does not persist these fields into state. Terraform consequently triggers an asynchronous version update via `google.appengine.v1.Versions.CreateVersion` (a long-running operation taking ~40s).
- **Harness Bug:** In `tests/e2e/script_test.go`, `waitForReconciliationAfterPatch` previously used a fixed 5-second timeout and treated any change in `resourceVersion` as reconciliation completion. When the legacy controller began updating, it set condition `Ready=False, reason=Updating`, incrementing `resourceVersion`. This caused `waitForReconciliationAfterPatch` to exit immediately while the 40-second GCP operation was still in flight, resulting in Terraform's ongoing POST and polling calls bleeding into Phase 3 (Direct Takeover) and causing unexpected write failures.
- **Fix:** Enhanced `waitForReconciliationAfterPatch` in `tests/e2e/script_test.go` to inspect status conditions. If the resource is in an `Updating` state (`Ready == False && reason == "Updating"`), it continues waiting up to `create.DefaultWaitForReadyTimeout` until the operation finishes and the resource is no longer updating.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully provisioned `AppEngineStandardAppVersion` via the legacy Terraform controller.
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller; Terraform's expected re-reconciliation diffs were captured in `_migration_diffs.json` and traffic recorded in `_http_migration_phase2_legacy_re-reconciliation.log`.
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single GET request and **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero direct takeover diffs were detected.
- **Phase 4 (Direct Re-reconcile):** The Direct controller re-reconciled cleanly with a single GET request and **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs detected.
- **Diffs Status:** `_migration_diffs.json` contains no `"isNewObject": false` blocks for `controller: direct`. Full state compatibility is maintained.

### Mock Alignment
- Verified `TestMigrationToDirect` passes hermetically against MockGCP (`E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest RUN_E2E=1`).
- Verified `TestAllInSeries` passes against MockGCP via `./hack/compare-mock`.

### Ratcheting Exclusions
- `AppEngineStandardAppVersion` is retained in `tests/e2e/ratcheting.go` exclusion list because its default controller is Terraform (which fails re-reconciliation due to the missing deployment state in the legacy provider). It should be removed from exclusions when the resource is promoted to use the Direct controller by default.
