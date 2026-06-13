# Compute Service Journal

### 2026-06-13 Implementation of ComputeManagedSSLCertificate Direct Types
- **Context**: Implementing direct KRM types and `generate.sh` configuration for `ComputeManagedSSLCertificate`.
- **Problem**: The GCP `SslCertificate` is naturally a global resource. The KRM generator automatically generated `Location` as a required `string` property, which caused existing test fixtures to fail schema validation because they lacked `spec.location`. Moreover, the API check `TestCRDFieldPresenceInTests` failed because the `location` field defined in the CRD was not covered/tested in the test fixture's manifest.
- **Solution**: 
  1. Updated `Location` in `ComputeManagedSSLCertificateSpec` (inside `sslcertificate_types.go`) to be an optional pointer `*string` with `// +optional`.
  2. Added `location: global` explicitly to the basic test fixture's `create.yaml` under `pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computemanagedsslcertificate/`.
- **Impact**: Ensures that global Compute resources migrated to the direct controller keep backwards compatibility with existing YAMLs (where `location` is omitted/unspecified) while successfully passing OpenAPI schema and API coverage tests.
