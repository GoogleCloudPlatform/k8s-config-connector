# Migration Journal: ComputeNetworkEndpoint

## Observations
- **Target Resource**: `ComputeNetworkEndpoint` (`compute.cnrm.cloud.google.com/v1alpha1`).
- **Test Fixture**: `pkg/test/resourcefixture/testdata/basic/compute/v1alpha1/computenetworkendpoint/computenetworkendpointbasic` (the only test fixture under `basic`, and does not have the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation in `create.yaml`).
- **Takeover Diff Result**: Executed `TestMigrationToDirect` against real GCP (`cnrm-barni-4`) using `./hack/record-gcp TestMigrationToDirect/fixtures/computenetworkendpointbasic$`.
- **Phase Breakdown**:
  - **Phase 1 (Legacy Create)**: Successfully provisioned the dependencies (`ComputeNetwork`, `ComputeSubnetwork`, `ComputeInstance`, `ComputeNetworkEndpointGroup`) and primary `ComputeNetworkEndpoint` resource using the legacy Terraform controller.
  - **Phase 2 (Legacy Re-reconcile)**: Successfully verified re-reconciliation behavior under the legacy Terraform controller (0 mutations, read-only `POST .../listNetworkEndpoints` call).
  - **Phase 3 (Direct Takeover)**: Transitioned reconciler to `direct`. The Direct controller adopted the resource cleanly with a **0-write no-op reconciliation** (read-only `POST .../listNetworkEndpoints` call, 0 mutations, zero takeover diffs).
  - **Phase 4 (Direct Re-reconcile)**: Tested Direct re-reconciliation with a no-op cookie annotation (`test.cnrm.cloud.google.com/reconcile-cookie: re-reconcile-direct-v1`). The Direct controller re-reconciled cleanly with a **0-write no-op reconciliation** (read-only `POST .../listNetworkEndpoints` call, 0 mutations, no diffs).
- **Recorded Traffic**: HTTP migration cassettes (`_http_migration_phase1_legacy_create.log`, `_http_migration_phase2_legacy_re-reconciliation.log`, `_http_migration_phase3_direct_takeover.log`, and `_http_migration_phase4_direct_re-reconciliation.log`) and `_migration_diffs.json` are verified against live GCP project `cnrm-barni-4`.
- **Diff Analysis**: `_migration_diffs.json` contains only initial creation diffs (`"isNewObject": true`) from Terraform. No diffs (`"isNewObject": false`) were reported during direct takeover (Phase 3) or direct re-reconciliation (Phase 4), confirming full state compatibility between the legacy Terraform controller and the Direct controller.
- **Controller Immutability & Comparison**: In `computenetworkendpoint_controller.go`, `ComputeNetworkEndpoint` is immutable. `compareComputeNetworkEndpoint` normalizes instance references by extracting the last component and runs `DiffForTopLevelFields`. Because the instance name and port match the desired state, zero diffs are produced.
- **Mock Alignment**: `TestMigrationToDirect` and `TestAllInSeries` both pass cleanly against MockGCP.
- **Ratcheting Exclusion**: Verified `ComputeNetworkEndpoint` is not present in the ratcheting exclusions in `tests/e2e/ratcheting.go`.
