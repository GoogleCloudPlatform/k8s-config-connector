# Migration Journal: ComputeRegionPerInstanceConfig

## Observations
- **Target Resource**: `ComputeRegionPerInstanceConfig` (`compute.cnrm.cloud.google.com/v1alpha1`).
- **Test Fixture**: `pkg/test/resourcefixture/testdata/basic/compute/v1alpha1/computeregionperinstanceconfig/computeregionperinstanceconfigbasic` (the only test fixture under `basic`, and does not have the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation in `create.yaml`).
- **Takeover Diff Result**: Executed `TestMigrationToDirect` against real GCP (`cnrm-barni-4`) using `./hack/record-gcp TestMigrationToDirect/fixtures/computeregionperinstanceconfigbasic$`.
- **Phase Breakdown**:
  - **Phase 1 (Legacy Create)**: Successfully provisioned dependencies (`ComputeHealthCheck`, `ComputeNetwork`, `ComputeInstanceTemplate`, `ComputeInstanceGroupManager`) and primary `ComputeRegionPerInstanceConfig` resource using the legacy Terraform controller.
  - **Phase 2 (Legacy Re-reconcile)**: Successfully verified re-reconciliation behavior under the legacy Terraform controller (0 mutations, read-only `POST .../listPerInstanceConfigs` call).
  - **Phase 3 (Direct Takeover)**: Transitioned reconciler to `direct`. The Direct controller adopted the resource cleanly with a **0-write no-op reconciliation** (read-only `POST .../listPerInstanceConfigs` call, 0 mutations, zero takeover diffs).
  - **Phase 4 (Direct Re-reconcile)**: Tested Direct re-reconciliation with a no-op cookie annotation (`test.cnrm.cloud.google.com/reconcile-cookie: re-reconcile-direct-v1`). The Direct controller re-reconciled cleanly with a **0-write no-op reconciliation** (read-only `POST .../listPerInstanceConfigs` call, 0 mutations, no diffs).
- **Test Framework Alignment**: In Compute Engine API, listing per-instance configs is performed via HTTP POST (`POST .../listPerInstanceConfigs`), which is a read-only query rather than a mutating write. Updated `tests/e2e/migration_test.go` to treat `strings.Contains(event.Request.URL, "/listPerInstanceConfigs")` as read-only for POST requests, aligning with `tests/e2e/unified_test.go`.
- **Recorded Traffic**: Real HTTP migration cassettes (`_http_migration_phase1_legacy_create.log`, `_http_migration_phase2_legacy_re-reconciliation.log`, `_http_migration_phase3_direct_takeover.log`, and `_http_migration_phase4_direct_re-reconciliation.log`) and `_migration_diffs.json` were generated and verified against live GCP project `cnrm-barni-4`.
- **Diff Analysis**: `_migration_diffs.json` contains only initial creation diffs (`"isNewObject": true`) from Terraform. No diffs (`"isNewObject": false`) were reported during direct takeover (Phase 3) or direct re-reconciliation (Phase 4), confirming full state compatibility between the legacy Terraform controller and the Direct controller.
- **Mock Alignment**: `TestMigrationToDirect` and `TestAllInSeries` both pass cleanly against MockGCP.
- **Ratcheting Exclusion**: Verified `ComputeRegionPerInstanceConfig` is not present in the ratcheting exclusions in `tests/e2e/ratcheting.go`.
