# SQLSSLCert KRM Types Journal

## Observations & Learnings

1. **CRD Schema Parity**:
   - The baseline CRD `sqlsslcerts.sql.cnrm.cloud.google.com` is `v1beta1` with `spec.commonName` (required string), `spec.instanceRef` (required SQLInstanceRef), and `spec.resourceID` (optional string).
   - The status contains `conditions`, `cert`, `certSerialNumber`, `createTime`, `expirationTime`, `observedGeneration`, `privateKey`, `serverCaCert`, and `sha1Fingerprint`.
   - By matching the exact types, comments, and tags, `dev/tasks/diff-crds` produces an empty diff.
   - Note on godoc comments: the baseline CRD did not have descriptions on the top-level `spec` and `status` properties (only on individual fields). Omitting doc comments on `SQLSSLCertSpec` and `SQLSSLCertStatus` ensures that `crd-to-simple-schema` produces identical schema without `_no_description.spec` or `_no_description.status` diffs.

2. **Avoiding Skip Conflict with Shared SSLCert Type**:
   - `SQLInstance` in `types.generated.go` uses `ServerCACert *SSLCert` which maps to proto message `google.cloud.sql.v1beta4.SslCert`.
   - If `SQLSSLCertSpec` is tagged with `// +kcc:spec:proto=google.cloud.sql.v1beta4.SslCert`, `controllerbuilder generate-types` treats `SslCert` as overridden by a non-generated type in the package, skipping the generation of `type SSLCert struct` in `types.generated.go`. This would break `SQLInstance` compilation (`invalid field type: invalid type`).
   - Leaving the `// +kcc:spec:proto` annotation off `SQLSSLCertSpec` prevents `generate-types` from skipping `SSLCert` in `types.generated.go`. Both `SQLInstance` and `SQLSSLCert` compile and generate cleanly.

3. **Configuring generate.sh**:
   - Updated `apis/sql/generate.sh` to include `--resource SQLSSLCert:SslCert`.
   - Re-running `generate.sh` succeeds, generates CRDs and deepcopy helpers, and produces zero diff in `dev/tasks/diff-crds`.
