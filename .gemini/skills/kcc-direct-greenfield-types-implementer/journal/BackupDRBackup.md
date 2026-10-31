# BackupDRBackup Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema Design & Hierarchical References**:
   - The GCP API for `Backup` (`google.cloud.backupdr.v1.Backup`) defines fields for GCP Backups inside Backup and DR vaults.
   - Resource Hierarchy: `projects/{project}/locations/{location}/backupVaults/{backupVault}/dataSources/{dataSource}/backups/{backup}`.
   - `BackupDRBackupSpec` maps:
     - `projectRef`: `*refsv1beta1.ProjectRef`
     - `location`: `*string`
     - `backupVaultRef`: `*backupdrv1beta1.BackupVaultRef`
     - `dataSourceRef`: `*DataSourceRef`
     - `resourceID`: `*string`
     - `labels`: `map[string]string`
     - `enforcedRetentionEndTime`: `*string`
     - `expireTime`: `*string`
     - `backupApplianceLocks`: `[]BackupLock`
   - Output-only fields (e.g. `etag`, `state`, `serviceLocks`, `computeInstanceBackupProperties`, `cloudSQLInstanceBackupProperties`, `backupApplianceBackupProperties`, `diskBackupProperties`, `gcpBackupPlanInfo`, `resourceSizeBytes`, `satisfiesPzs`, `satisfiesPzi`, `createTime`, `updateTime`, `consistencyTime`) are mapped to `BackupDRBackupObservedState`.
   - Output-only nested structs use the `*ObservedState` suffix (`DiskBackupPropertiesObservedState`, `Backup_GcpBackupPlanInfoObservedState`).
   - Acronym capitalization fixed in `Scheduling` for `minNodeCPUs` and `localSSDRecoveryTimeout`.

2. **Package Isolation and Import Cycle Resolution**:
   - `BackupDRBackup` in `v1alpha1` references `backupdrv1beta1.BackupVaultRef`.
   - `apis/backupdr/v1beta1` previously imported `apis/backupdr/v1alpha1` solely to reference `v1alpha1.Parent` for `BackupDRBackupPlan` and `BackupDRBackupPlanAssociation`. Since `Parent` is already declared in `v1beta1`, updating those types to use `Parent` eliminates the circular import dependency between `v1alpha1` and `v1beta1`.

3. **Inline External-Only DataSource Reference**:
   - `DataSource` does not yet have a standalone KCC CRD. To avoid naming violations (`naming_violations.txt`), `DataSourceRef` (external-only) and `DataSourceIdentity` are defined inline inside `apis/backupdr/v1alpha1/backupdrbackup_reference.go` and registered via `refs.Register(&DataSourceRef{})`.

4. **Identity & Reference Implementation**:
   - Canonical `gcpurls.Template` URL pattern: `"projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}/backups/{backup}"`.
   - `BackupIdentity` implements `identity.IdentityV2` and `BackupDRBackup` implements `identity.Resource`.
   - `BackupRef` implements `refs.Ref` and delegates normalization strictly to `refs.Normalize`.
   - Unit tests added in `backupdrbackup_identity_test.go` covering `BackupIdentity`, `DataSourceIdentity`, and `getIdentityFromBackupSpec`.

5. **Multi-version Generator Compatibility & Complete Mappers**:
   - `apis/backupdr/generate.sh` runs `generate-mapper --multiversion`.
   - Full non-empty implementations were added to `pkg/controller/direct/backupdr/backup_mappers.go` for `AccessConfig_v1alpha1_ToProto` (mapping `in.ExternalIP` to `NatIP`, `Type`, `NetworkTier`), `AttachedDiskObservedState_v1alpha1_ToProto`, `ComputeInstanceBackupPropertiesObservedState_v1alpha1_ToProto`, `DiskBackupPropertiesObservedState_v1alpha1_ToProto`, `GuestOSFeature_v1alpha1_ToProto`, `Metadata_v1alpha1_FromProto/ToProto`, `Scheduling_v1alpha1_FromProto/ToProto`, and `NetworkInterfaceObservedState_v1alpha1_FromProto/ToProto`.
