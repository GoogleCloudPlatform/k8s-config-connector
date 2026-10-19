# Migration Journal: ComputeNetworkFirewallPolicy

## Observations
- **Takeover Diff Result**: Ran `TestMigrationToDirect` against real GCP (`cnrm-barni-4`) for the fixture:
  - `pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computenetworkfirewallpolicy`
- **Phase Breakdown**:
  - **Phase 1 (Legacy Create)**: Successfully created the network firewall policy resource using the legacy Terraform/DCL controller.
  - **Phase 2 (Legacy Re-reconciliation)**: Successfully re-reconciled with a no-op update under the legacy controller (GET only, 0 writes).
  - **Phase 3 (Direct Takeover)**: Transitioned reconciler to `direct`. The Direct controller adopted the resource cleanly with a **0-write no-op reconciliation** (GET only, 0 writes).
  - **Phase 4 (Direct Re-reconciliation)**: Tested Direct re-reconciliation with a no-op annotation. The Direct controller re-reconciled cleanly with a **0-write no-op reconciliation** (GET only, 0 writes).
- **Recorded Traffic**: The HTTP golden logs (`_http_migration_phase2_legacy_re-reconciliation.log`, `_http_migration_phase3_direct_takeover.log`, and `_http_migration_phase4_direct_re-reconciliation.log`) and `_migration_diffs.json` were re-recorded against live GCP.
- **Diff Analysis**: `_migration_diffs.json` contains only the initial creation diff (`"isNewObject": true`) from Terraform. No diffs were reported during direct takeover (Phase 3) or direct re-reconciliation (Phase 4).
- **Mock Alignment**: `TestMigrationToDirect` and `TestAllInSeries` both pass cleanly against MockGCP.
