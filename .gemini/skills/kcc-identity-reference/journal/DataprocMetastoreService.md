# DataprocMetastoreService Journal

## Context
`DataprocMetastoreService` is a greenfield resource that does not have a KCC `_types.go` implementation because it is not yet fully managed by Config Connector, but other resources need to reference it.

## Findings
- **No `_types.go`**: Since it is a greenfield resource with no `_types.go`, we placed the identity and reference files directly under `apis/refs/` alongside similar resources like `memorystoreinstance`.
- **GVK & Group**: Based on DCL metadata (`pkg/dcl/metadata/metadata.go`), the service group prefix is `dataprocmetastore`, leading to GVK Group: `dataprocmetastore.cnrm.cloud.google.com` (with Version `v1beta1`).
- **Template Format**: The CAIS (`cloudassetinventory_names.jsonl`) lists this under `metastore.googleapis.com/Service` with the path format `projects/{project}/locations/{location}/services/{service}`.
- **Verification**: Tests for `pkg/gcpurls/` and `apis/refs/` successfully compiled and passed, verifying that the template correctly resolves and matches both CAIS and canonical KCC reference standards.
