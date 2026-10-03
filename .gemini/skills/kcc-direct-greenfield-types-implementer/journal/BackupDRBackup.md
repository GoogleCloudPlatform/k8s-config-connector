# BackupDRBackup Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema Design**:
   - The GCP API for `Backup` (`google.cloud.backupdr.v1.Backup`) defines fields for GCP Backups inside Backup and DR vaults.
   - Hierarchy: `projects/{project}/locations/{location}/backupVaults/{backupVault}/dataSources/{dataSource}/backups/{backup}`.
   - `BackupDRBackupSpec` maps:
     - `projectRef`: `*refsv1beta1.ProjectRef`
     - `location`: `*string`
     - `backupVault`: `*string`
     - `dataSource`: `*string`
     - `resourceID`: `*string`
     - `labels`: `map[string]string`
     - `enforcedRetentionEndTime`: `*string`
     - `expireTime`: `*string`
     - `backupApplianceLocks`: `[]BackupLock`
   - Output-only fields (e.g. `etag`, `state`, `serviceLocks`, `computeInstanceBackupProperties`, `cloudSQLInstanceBackupProperties`, `backupApplianceBackupProperties`, `diskBackupProperties`, `gcpBackupPlanInfo`, `resourceSizeBytes`, `satisfiesPzs`, `satisfiesPzi`, `createTime`, `updateTime`, `consistencyTime`) are mapped to `BackupDRBackupObservedState`.
   - Acronym capitalization fixed in `Scheduling` for `minNodeCPUs` and `localSSDRecoveryTimeout`.

2. **Identity & Reference Implementation**:
   - Canonical `gcpurls.Template` URL pattern: `"projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}/backups/{backup}"`.
   - `BackupIdentity` implements `identity.IdentityV2` and `BackupDRBackup` implements `identity.Resource`.
   - `BackupRef` implements `refs.Ref` and delegates normalization strictly to `refs.Normalize`.
   - Unit tests added in `backupdrbackup_identity_test.go` using `cmp.Diff`.

3. **Multi-version Generator Compatibility**:
   - `apis/backupdr/generate.sh` runs `generate-mapper --multiversion`, so manual mapping functions and `ToProto` stubs for nested output-only observed state types were placed in `pkg/controller/direct/backupdr/backup_mappers.go`.
