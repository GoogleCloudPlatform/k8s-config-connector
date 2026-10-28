# Migration Journal: DocumentAIProcessorDefaultVersion

## Observations
- **Target Resource**: `DocumentAIProcessorDefaultVersion` (`documentai.cnrm.cloud.google.com/v1alpha1`).
- **Test Fixture**: `pkg/test/resourcefixture/testdata/basic/documentai/v1alpha1/documentaiprocessordefaultversion/documentaiprocessordefaultversion-basic` (the only test fixture under `basic`, and does not have the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation).
- **Takeover Diff Result**: Executed `TestMigrationToDirect` against real GCP (`cnrm-barni-4`) using `./hack/record-gcp "TestMigrationToDirect/fixtures/documentaiprocessordefaultversion-basic$"`.
- **Phase Breakdown**:
  - **Phase 1 (Legacy Create)**: Successfully provisioned/managed the resource using the legacy Terraform controller (GET request to verify the parent processor).
  - **Phase 2 (Legacy Re-reconcile)**: Successfully re-reconciled with the legacy Terraform controller (GET request, 0 mutations).
  - **Phase 3 (Direct Takeover)**: Transitioned reconciler to `direct`. The Direct controller adopted the resource cleanly with a **0-write no-op reconciliation** (single read-only GET request, 0 mutations, no takeover diffs).
  - **Phase 4 (Direct Re-reconcile)**: Tested Direct re-reconciliation with a no-op annotation. The Direct controller re-reconciled cleanly with a **0-write no-op reconciliation** (single read-only GET request, 0 mutations, no diffs).
- **Recorded Traffic**: HTTP migration cassettes (`_http_migration_phase1_legacy_create.log`, `_http_migration_phase2_legacy_re-reconciliation.log`, `_http_migration_phase3_direct_takeover.log`, and `_http_migration_phase4_direct_re-reconciliation.log`) and `_migration_diffs.json` were recorded against live GCP project `cnrm-barni-4`.
- **Diff Analysis**: `_migration_diffs.json` is `null` (zero diffs across all phases). No takeover diffs (`"isNewObject": false` is absent) were reported during direct takeover (Phase 3) or direct re-reconciliation (Phase 4), confirming full state compatibility between the legacy Terraform controller and the Direct controller.
- **Mock Alignment**: `TestMigrationToDirect` and `TestAllInSeries` both pass cleanly against MockGCP.
- **Ratcheting Exclusion**: Verified `DocumentAIProcessorDefaultVersion` is not present in the ratcheting exclusions in `tests/e2e/ratcheting.go`.
