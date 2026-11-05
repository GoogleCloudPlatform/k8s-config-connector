# Migration Journal: ApigeeNATAddress

## Observations
- **Takeover Diff Result**: Ran `TestMigrationToDirect` against real GCP (`cnrm-barni-4`) for the `ApigeeNATAddress` test fixture:
  - `apigeenataddress-basic`
- The fixture successfully completed all four migration phases:
  - Phase 1 (Legacy Create): Successfully provisioned `ApigeeInstance` dependency and `ApigeeNATAddress` resource using the legacy Terraform controller.
  - Phase 2 (Legacy Re-reconcile): Verified re-reconciliation behavior with the Terraform controller (read-only `GET` request, 0 mutations).
  - Phase 3 (Direct Takeover): Successfully adopted the resource using the Direct controller with a clean **0-write no-op reconciliation** (single read-only `GET` request, no takeover diffs).
  - Phase 4 (Direct Re-reconcile): Confirmed clean **0-write no-op re-reconciliation** using the Direct controller (single read-only `GET` request, 0 mutations).
- No takeover diffs were detected (`"isNewObject": false` diffs count is 0 in `_migration_diffs.json`), confirming full state compatibility between the legacy Terraform controller and the Direct controller.
- **Mock Alignment**: Verified that both `TestMigrationToDirect` and `TestAllInSeries` pass cleanly against MockGCP. Generated `_http_mock.diff` between `_http_old_controller_mock.log` and `_http_mock.log`.
- **Ratcheting Exclusion**: Verified `ApigeeNATAddress` is not present in the ratcheting exclusions in `tests/e2e/ratcheting.go`.
