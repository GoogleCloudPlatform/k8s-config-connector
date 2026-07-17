# Export Support Journal for BackupDRBackupVault

## Observations
- `BackupDRBackupVault` uses the direct controller reconciler structure, and its `BackupVaultIdentity` natively parses CAI resource URIs (`//backupdr.googleapis.com/projects/{project}/locations/{location}/backupVaults/{backupvault}`).
- In `backupvault_controller.go`, we implemented `AdapterForURL` to support parsing external URIs.
- In the `Export` method, we updated the map assignment sequence so that `u.Object = uObj` is set prior to GVK and name fields. This prevents unstructured metadata from being silently overwritten and wiped out, resolving a tricky serialization bug.
- Added a case to `tests/e2e/export.go` to handle `BackupDRBackupVault` using `resolveCAISURI(h, obj)`.
- Verified and generated the corresponding `_exported.yaml` files under both the `backupdrbackupvault-minimal` and `backupdrbackupvault-full` test fixture data folders.
- Ran tests against mockgcp and envtest successfully: all E2E export tests passed with no differences.
