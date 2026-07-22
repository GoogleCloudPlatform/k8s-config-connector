# BackupDRBackupPlanAssociation Identity and Refs migration

Conformed `BackupDRBackupPlanAssociation` to the `identity.IdentityV2` and `refs.Ref` standard patterns.

## Learnings

- `BackupDRBackupPlanAssociation` is managed under `apis/backupdr/v1beta1`.
- Because some of the existing controllers and test fixtures used compatibility/parsing helper methods such as `NewBackupPlanAssociationIdentity` and `ParseBackupPlanAssociationExternal`, we kept these helpers in place so that we did not have to break any existing controllers/consumers.
- This pattern works well when transitioning resources without performing a massive direct controller refactor at the same time.
- Used `dev/tasks/generate-types-and-mappers` to clean up old deepcopy types from `zz_generated.deepcopy.go`.
- Generated the updated `_identities.yaml` golden file using `WRITE_GOLDEN_OUTPUT=1 go test -v ./pkg/cli/powertools/cais/...`.
