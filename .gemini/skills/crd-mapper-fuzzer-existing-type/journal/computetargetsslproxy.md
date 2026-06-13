# ComputeTargetSSLProxy Direct Type Transition Journal

- **Issue**: Transitioning the ComputeTargetSSLProxy KCC resource to use direct KRM types and generate.sh.
- **Learnings**:
  - **Schema Compatibility**: To achieve 100% strict schema compatibility, we added `// +required` to the `Spec` field to match the baseline CRD schema's required fields.
  - **Protobuf/KRM Field Naming Deviations**: Hand-coded mappers were implemented in `pkg/controller/direct/compute/targetsslproxy_mapper.go` because the KRM struct field names deviate from default protobuf field naming (e.g., `BackendServiceRef` maps to `Service`, `CertificateMapRef` maps to `CertificateMap`, and `SslPolicyRef` maps to `SslPolicy`).
  - **Status Mapping**: Added the `// +kcc:status:proto=google.cloud.compute.v1.TargetSslProxy` annotation above the `ComputeTargetSSLProxyStatus` struct so that the status fields (`creation_timestamp`, `id`, `self_link`) are correctly recognized and mapped by the fuzzer.
  - **E2E Fuzz Testing**: Verified that the handwritten mappers and fuzzers pass `pkg/fuzztesting/fuzztests` cleanly with zero diffs over 100,000 randomized iterations.
