# Journal: BackupDRBackupVault Transition

## Observations
- `BackupDRBackupVault` types are already fully handcoded/defined and strictly schema-compatible with the existing CRD.
- Added `--include-skipped-output` to both `generate-types` and `generate-mapper` commands in `apis/backupdr/v1beta1/generate.sh`.
- During regeneration of `types.generated.go`, unreachable local `Status` and `Any` structs were pruned/removed and replaced with `common.Status` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/common`.
- This change affected `BackupDRBackupPlanAssociation`, which had a custom mapper (`backupplanassociation_mappers.go`) referencing the deleted `krm.Status` and `krm.Any` types.
- Updated `backupplanassociation_mappers.go` to import `apis/common` and cleanly map to/from `common.Status` instead, eliminating compile errors and keeping the status model clean.
- Successfully verified that the regenerated CRD for `BackupDRBackupVault` had absolutely zero diff against the baseline CRD.
