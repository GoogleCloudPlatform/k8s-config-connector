# BackupDR Journal

### 2026-11-06 BackupDRBackup Direct Types Prepopulation
- **Context**: Greenfield Step 1a prepopulation of direct KRM types, IdentityV2, references, and judgement queue for `BackupDRBackup` (GCP proto `google.cloud.backupdr.v1.Backup`).
- **Problem**:
  1. `generate-types` rejected `BackupDRBackup` as non-declarative because the resource pattern `projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}/backups/{backup}` ends with a point-in-time collection (`backups`).
  2. `BackupDRBackup` transitively references `AccessConfig`, `NetworkInterface`, and `Metadata`. The protobuf definitions for `AccessConfig` and `NetworkInterface` define fields with custom `json_name` tags (`external_ip` with `json_name="natIP"` and `ip_address` with `json_name="networkIP"`). `generatemapper` derived field accessors from JSON names (`in.NatIP`, `in.NetworkIP`), whereas the generated Go client types in `backupdrpb` name the fields from proto identifiers (`ExternalIp`, `IpAddress`). Furthermore, `Metadata` is skipped by default by `generatemapper` due to suffix filtering.
- **Solution**:
  1. Added `--allow-non-declarative` to `generate-types` in `apis/backupdr/generate.sh`.
  2. Implemented handwritten mapping functions for `AccessConfig`, `NetworkInterfaceObservedState`, and `Metadata` in `pkg/controller/direct/backupdr/backupdrbackup_mapper.go`.
- **Impact**: Enables clean mechanical compilation and conformance test passing for `BackupDRBackup` direct KRM types and IdentityV2 while preserving the open `judgement_queue.yaml` for Step 1b triage.
