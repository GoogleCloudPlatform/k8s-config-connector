### [2026-06-13] ComputeTargetSSLProxy Direct Type Implementation
- **Context**: Implementing direct KRM types and generate.sh configuration for `ComputeTargetSSLProxy`.
- **Problem**: Greenfield types generation script creates a very minimal KRM type from the proto structure, omitting several pre-existing, critical fields (e.g., `BackendServiceRef`, `SslCertificatesRef`, `SslPolicyRef`, etc.) and using the default/incorrect short name plurals (e.g., `gcpcomputetargetsslproxys` instead of `gcpcomputetargetsslproxies`).
- **Solution**: Hand-coded the complete schema fields in `targetsslproxy_types.go`, pulling references from existing standard types such as `ComputeBackendServiceRef`, `ComputeSSLPolicyRef`, and `ComputeSSLCertificateRef`. Standardized the GVK and shortName tags to ensure 100% strict compatibility with the pre-existing CRD.
- **Impact**: Ensures that when migrating any existing stable/beta resources to direct KRM types, schema compatibility is perfectly maintained.
