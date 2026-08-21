# Export Support for LoggingLogBucket Journal

## Findings and Observations

During the implementation of export support for `LoggingLogBucket`, several key observations were made:

1. **CAIS Identity Resolution Mapping:**
   `LoggingLogBucket` was previously marked as an unsupported kind in `unsupportedKinds` within `pkg/cais/cais.go`. Removing it from the `unsupportedKinds` map allowed the export framework to correctly resolve and extract its CAIS identity.

2. **Parsing Dynamic parent / location in Export:**
   Since the `LogBucket` GCP resource does not contain `location` as a distinct field within its API schema (as location is embedded inside the full API path), `LoggingLogBucketSpec_FromProto` originally mapped the `location` field as empty. In `LoggingLogBucketAdapter.Export(...)`, we resolved this by manually populating `obj.Spec.Location` from `a.id.Location` (which is parsed from the external identifier / URI).

3. **Dynamic Resource Normalization for Folder-Scoped Resources:**
   For tests where `Folder` resources are dynamically created at runtime (e.g. `folderlogbucket` test fixture), the generated folder IDs are volatile and change on every test run.
   To prevent these volatile IDs from causing non-deterministic diffs in the exported golden files (`_exported.yaml`), we added an automatic look-up/discovery step in `tests/e2e/unified_test.go`. If there is a dynamically created `Folder` resource in the namespace, the test harness discovers its resolved numeric `folderId` from its status, and maps it in the replacements table to be properly normalized to `${folderID}`.
