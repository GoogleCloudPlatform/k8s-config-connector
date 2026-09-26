# Greenfield Direct KRM Types Implementation: BackupDRBackup

## Overview
Implemented Greenfield direct KRM types for GCP BackupDR Backup resource under `apis/backupdr/v1alpha1/`.

## Proto and Resource Mapping
- Proto service: `google.cloud.backupdr.v1`
- Proto message: `google.cloud.backupdr.v1.Backup` (in `google/cloud/backupdr/v1/backupvault.proto`)
- Resource mapping: `BackupDRBackup:Backup`
- API group: `backupdr.cnrm.cloud.google.com/v1alpha1`
- GCP Resource Pattern: `projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}/backups/{backup}`

## Design Decisions
1. **Hierarchical Parent**: The resource parent contains `projectRef`, `location`, `backupVault`, and `dataSource`. `Location`, `BackupVault`, and `DataSource` are pointer scalar strings per Greenfield type requirements.
2. **Field Behavior & Validation**:
   - `spec`: Contains mutable/user-provided fields (`projectRef`, `location`, `backupVault`, `dataSource`, `resourceID`, `labels`, `enforcedRetentionEndTime`, `expireTime`, `backupApplianceLocks`).
   - `observedState`: Contains GCP-assigned output-only fields (`description`, `createTime`, `updateTime`, `consistencyTime`, `state`, `serviceLocks`, `backupApplianceLocks`, `backupType`, `gcpBackupPlanInfo`, `resourceSizeBytes`, `satisfiesPzs`, `satisfiesPzi`, `etag`).
   - `etag` is mapped to `observedState` in accordance with KCC conventions for concurrency tokens.
3. **Mappers**: Added `pkg/controller/direct/backupdr/backup_mappers.go` to cleanly map `BackupDRBackupSpec` and `BackupDRBackupObservedState` to/from protobuf, handling oneof union fields such as `ClientLockInfo`, `LockSource`, and `PlanInfo`.
