# Migration Journal: AppEngineServiceSplitTraffic

## Observations
- **Real GCP Execution**: Ran `TestMigrationToDirect` and `TestAllInSeries` against real GCP project `cnrm-barni-4` using `./hack/record-gcp`:
  - Fixture: `pkg/test/resourcefixture/testdata/basic/appengine/v1alpha1/appengineservicesplittraffic/appengineservicesplittrafficbasic`
- **Root Cause & Fixes**:
  1. **Legacy Terraform Provider Missing Flattener**:
     - `resourceAppEngineServiceSplitTrafficRead` in the vendored Terraform provider (`third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/appengine/resource_app_engine_service_split_traffic.go`) failed to read and persist `split` into Terraform state.
     - Consequently, Terraform detected a diff on every reconciliation attempt and sent an unexpected `PATCH` request during Phase 2 (legacy re-reconciliation).
     - Fixed by adding `flattenAppEngineServiceSplitTrafficSplit` and setting `split` and `migrate_traffic` in `resourceAppEngineServiceSplitTrafficRead`.
  2. **Direct Controller Takeover 0-Write Verification**:
     - In `pkg/controller/direct/appengine/appengineservicesplittraffic_controller.go`, `Update` checked `if diffRes.Empty() && mburMatch`. Because objects adopted from Terraform do not have the `cnrm.cloud.google.com/mutable-but-unreadable-fields-hash` annotation, `mburMatch` evaluated to `false`, causing an unnecessary `PATCH` to be sent to GCP on takeover.
     - Fixed by updating `Update` so that when `diffRes.Empty()` is true (the live split allocations and shardBy in GCP match the desired spec), it records the MBUR hash annotation and updates status with zero GCP writes.
  3. **Non-Deletable Fixture Classification**:
     - Added `AppEngineServiceSplitTraffic` to `IsNonDeletable()` in `pkg/test/httplogreport/fixture/loader.go` to match GCP AppEngine semantics where split traffic cannot be deleted.
  4. **Ratcheting Exclusion Removal**:
     - Removed `AppEngineServiceSplitTraffic` from the ratcheting exclusions list in `tests/e2e/ratcheting.go`, as both legacy Terraform and Direct controllers now achieve clean 0-write re-reconciliation.
- **Verification Results**:
  - `TestMigrationToDirect` passes cleanly against real GCP (`cnrm-barni-4`) with 0 writes in Phases 2, 3, and 4, and `null` diffs in `_migration_diffs.json`.
  - `TestAllInSeries` passes cleanly against real GCP (`cnrm-barni-4`).
  - Hermetic tests (`TestMigrationToDirect`) pass cleanly against `mockgcp`.
