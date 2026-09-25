# Identity and Reference Implementation: BackupDRBackup

## Identity Template
- Template: `backupdr.googleapis.com/projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}/backups/{backup}`
- Matches `docs/ai/metadata/cloudassetinventory_names.jsonl` entry for `backupdr.googleapis.com/Backup`.

## Components
- Identity: `BackupIdentity` in `apis/backupdr/v1alpha1/backupdrbackup_identity.go`
- Reference: `BackupRef` in `apis/backupdr/v1alpha1/backupdrbackup_reference.go`
- Unit tests: `apis/backupdr/v1alpha1/backupdrbackup_identity_test.go`
- Cross-check: `GetIdentity` cross-checks parsed spec identity against `status.externalRef` when present.
