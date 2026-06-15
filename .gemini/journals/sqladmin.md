### [2026-06-15] SQLAdminBackup Scaffolding and Types Implementation
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `SQLAdminBackup` under the `sqladmin.cnrm.cloud.google.com` group.
- **Problem**: The GCP SQL Admin API's backup resource is named `BackupRun` in the protobuf definitions (`google/cloud/sql/v1/cloud_sql_backup_runs.proto`) and belongs to package `google.cloud.sql.v1`, rather than `google.cloud.sqladmin.v1` with message `Backup`.
- **Solution**: 
  1. Configured the types generator inside `generate.sh` to use:
     ```bash
     go run . generate-types \
       --service google.cloud.sql.v1 \
       --api-version sqladmin.cnrm.cloud.google.com/v1alpha1 \
       --resource SQLAdminBackup:BackupRun
     ```
  2. Overrode the `DiskEncryptionConfiguration` struct inside `backuprun_types.go` to reference the canonical `kmsv1beta1.KMSCryptoKeyRef` rather than raw string fields to satisfy reference-like field guidelines.
  3. Ran `dev/tasks/generate-ci-cd-jobs` to dynamically generate the required presubmit and periodic scripts for the new `sqladmin.cnrm.cloud.google.com` API group.
- **Impact**: Unblocks next-step development (controller and mapper implementation) for `SQLAdminBackup` by establishing the base KRM schemas, CRDs, and proper IdentityV2.
