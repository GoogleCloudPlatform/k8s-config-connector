# Migration Journal: FirebaseHostingSite

## Observations
- **Takeover Diff Result**: Ran `TestMigrationToDirect` against real GCP for the `FirebaseHostingSite` test fixture:
  - `firebasehostingsiteautogen`
- The fixture successfully completed all four migration phases:
  - Phase 1 (Legacy Create): Provisioned the initial `FirebaseHostingSite` resource using the Terraform controller.
  - Phase 2 (Legacy Re-reconcile): Verified re-reconciliation behavior with the Terraform controller.
  - Phase 3 (Direct Takeover): Successfully adopted the resource using the Direct controller with a clean **0-write no-op reconciliation** (no takeover diffs detected).
  - Phase 4 (Direct Re-reconcile): Confirmed clean **0-write no-op re-reconciliation** using the Direct controller.
- No takeover diffs were detected (`"isNewObject": false` diffs count is 0 in `_migration_diffs.json`), confirming full state compatibility between the legacy Terraform controller and the direct controller.
- **Ratcheting Exclusion**: Verified `FirebaseHostingSite` is not present in the ratcheting exclusions in `tests/e2e/ratcheting.go`.
