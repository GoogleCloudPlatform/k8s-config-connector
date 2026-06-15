# Filestore (file) Journal

### 2026-06-15 Greenfield Direct Types & Identity for FilestoreBackup
- **Context**: Implementing direct types, CRD, and IdentityV2 for `FilestoreBackup` under the `file.cnrm.cloud.google.com/v1alpha1` group.
- **Problem**: The proto-to-KRM generator failed when using the `google.cloud.file.v1` package specified in the issue because the actual protobuf package name in the cloned `googleapis` repo is `google.cloud.filestore.v1`.
- **Solution**: Updated the `generate.sh` script to reference `google.cloud.filestore.v1` as the protobuf package, while maintaining `file.cnrm.cloud.google.com/v1alpha1` as the API group/version.
- **Impact**: Enables flawless code generation of direct types and CRDs for FilestoreBackup and other filestore resources.
