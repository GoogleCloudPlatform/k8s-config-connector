### [2026-06-17] PrivateCACertificateTemplate Identity & Refs
- **Context**: Implementing the modern `identity.IdentityV2` and `refs.Ref` patterns for the `PrivateCACertificateTemplate` resource under `apis/privateca/v1beta1`.
- **Problem**: 
    1. `PrivateCACertificateTemplate` lacked structured `IdentityV2` and `refs.Ref` implementations.
    2. Golden tests in `TestGoldenIdentitiesYamlFiles` expected a valid CAIS URL pattern rather than "unknown".
- **Solution**:
    1. Implemented `apis/privateca/v1beta1/privatecacertificatetemplate_identity.go` to use `identity.IdentityV2` backed by `gcpurls.Template` matching the URL path `projects/{project}/locations/{location}/certificateTemplates/{certificatetemplate}`.
    2. Implemented `apis/privateca/v1beta1/privatecacertificatetemplate_reference.go` defining `PrivateCACertificateTemplateRef` conforming to `refs.Ref`.
    3. Added comprehensive unit testing under `apis/privateca/v1beta1/privatecacertificatetemplate_identity_test.go` to verify standard template-based URL parsing, reference validation, and interface adherence using `cmp.Diff`.
    4. Generated the updated golden test files using `WRITE_GOLDEN_OUTPUT=1 go test` under the CAIS package.
- **Impact**: `PrivateCACertificateTemplate` now fully conforms to the modern Config Connector patterns for declarative GCP identity and reference resolution, with updated golden test assertions.
