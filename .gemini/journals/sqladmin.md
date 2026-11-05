### [2026-06-29] SQLAdminBackup Greenfield Types Generation
- **Context**: Implementing direct types, CRD, and IdentityV2 for `SQLAdminBackup` under Issue #10298.
- **Problem**: The issue description pinned the GCP service to `google.cloud.sqladmin.v1` and the resource to `SQLAdminBackup:Backup`. However, the actual proto package and messages are structured as `google.cloud.sql.v1` and `BackupRun` in the official googleapis. Attempting to use the suggested fields resulted in "proto: not found" errors.
- **Solution**: Set the `--service` flag to `google.cloud.sql.v1` and the `--resource` flag to `SQLAdminBackup:BackupRun`. This successfully mapped to the proto definitions and allowed types, CRD, and IdentityV2 generation. KMS reference fields (`KMSCryptoKeyRef`) and instance reference fields (`SQLInstanceRef`) were then added manually to `sqladminbackup_types.go` to conform to KCC standards.
- **Impact**: The next agent implementing the SQLAdminBackup controller or mappers can build directly on these types and use `BackupRun` as the underlying GCP proto.

### [2026-10-15] SQLSSLCert Direct Controller Migration
- **Context**: Implementing direct controller and recording/verifying test fixtures for brownfield `SQLSSLCert` under Issue #13993.
- **Problem**: `SQLSSLCert` has a server-generated resource ID (`sha1Fingerprint`). Furthermore, `privateKey` and `serverCaCert` are returned only on initial cert creation in `SslCertsInsertResponse`, not during subsequent `Get` calls.
- **Solution**: Followed the 4 rules in `kcc-direct-service-generated-id` to guard `Find()` when ID is unset, leave ID out of create payload, use `createOp.SetSpecResourceID()` upon creation, and validate in `GetIdentity()`. Preserved `privateKey` and `serverCaCert` in adapter status across reconciliations. Added `k8s.ReconcilerTypeDirect` to `SupportedControllers` in `static_config.go` and verified dual-testing with MockGCP.

