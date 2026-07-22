### [2026-07-11] PrivateCACertificate Identity & Refs
- **Context**: Implementing the modern `identity.IdentityV2` and `refs.Ref` patterns for the `PrivateCACertificate` resource under `apis/privateca/v1beta1`.
- **Problem**:
    1. `PrivateCACertificate` lacked structured `IdentityV2` and `refs.Ref` implementations.
    2. Golden tests in `TestGoldenIdentitiesYamlFiles` expected a valid identity/reference mapping instead of "unknown".
- **Solution**:
    1. Implemented `apis/privateca/v1beta1/privatecacertificate_identity.go` to use `identity.IdentityV2` backed by `gcpurls.Template` matching the URL path `projects/{project}/locations/{location}/caPools/{caPool}/certificates/{certificate}`.
    2. Resolved the dependency reference `CaPoolRef` to obtain the parent CAPool ID dynamically when fetching the certificate identity.
    3. Implemented `apis/privateca/v1beta1/privatecacertificate_reference.go` defining `PrivateCACertificateRef` conforming to `refs.Ref` and delegating normalization to `refs.NormalizeWithFallback`.
    4. Added comprehensive unit testing under `apis/privateca/v1beta1/privatecacertificate_identity_test.go` to verify template-based URL parsing, reference validation, and interface adherence using `cmp.Diff`.
    5. Generated the updated golden test files using `WRITE_GOLDEN_OUTPUT=1 go test` under the CAIS package.
- **Impact**: `PrivateCACertificate` now fully conforms to the modern Config Connector patterns for declarative GCP identity and reference resolution, with updated golden test assertions.
