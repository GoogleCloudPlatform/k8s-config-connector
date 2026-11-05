# GKEBackup Journal

### 2026-07-02 GKEBackupBackupChannel Direct Types Scaffolding
- **Context**: Implementing the initial KRM types, CRD, and IdentityV2 for the `GKEBackupBackupChannel` resource (GCP service `google.cloud.gkebackup.v1.BackupChannel`).
- **Problem**: The proto model definition for `BackupChannel` includes `destination_project` which represents where backups are allowed to be stored.
- **Solution**: Implemented `DestinationProjectRef` in the Spec as a standard KCC `ProjectRef` pointing to a `Project` resource. Output-only fields like `uid`, `create_time`, `update_time`, `etag`, and `destination_project_id` were successfully mapped under `ObservedState` as `*string` types.
- **Impact**: Provides standard structured fields and schema validations for GKEBackupBackupChannel to pave the way for implementing its direct reconciler/controller in subsequent phases.

### 2026-10-17 GKEBackupRestoreChannel Direct Controller Implementation
- **Context**: Implementing direct controller, mappers, fuzzer, and E2E fixtures for `GKEBackupRestoreChannel` (GCP service `google.cloud.gkebackup.v1.BackupForGKE`).
- **Problem**: In GCP, `destination_project` on `RestoreChannel` is immutable. When creating the restore channel, the API accepts project IDs (e.g. `projects/{projectId}`) or project numbers (e.g. `projects/{projectNumber}`), but GET responses canonicalize this field to `projects/{projectNumber}` and populate `destination_project_id` with `{projectId}`. This creates an apparent diff between desired and actual state.
- **Solution**: In `compareRestoreChannel`, normalized `maskedActual.DestinationProject` if `destination_project_id` matches desired or refers to the same project. In addition, ensured `AdapterForURL` is implemented using `id.FromExternal(url)` and fixed `Export()` to set `u.Object = uObj` before `SetName` and `SetGroupVersionKind` so metadata is preserved.
- **Impact**: Enables clean drift-free reconciliation, URL-based adapter lookup, accurate export, and full test coverage for `GKEBackupRestoreChannel`.

