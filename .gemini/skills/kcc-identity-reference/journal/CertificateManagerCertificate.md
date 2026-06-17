### [2026-06-06] CertificateManagerCertificate Identity & Refs
- **Context**: Migrating `CertificateManagerCertificate` to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**:
    1. `CertificateManagerCertificate` used the legacy manual normalization logic and was missing the standard `identity.IdentityV2` implementation.
    2. The URL templates were not using the centralized `gcpurls.Template` registry.
- **Solution**:
    1. Implemented `apis/certificatemanager/v1beta1/certificatemanagercertificate_identity.go` to use `identity.IdentityV2` backed by `gcpurls.Template`.
    2. Rewrote `apis/certificatemanager/v1beta1/certificatemanagercertificate_reference.go` to implement `refs.Ref` and register with `refs.Register`.
    3. Merged separate reference tests into a comprehensive `apis/certificatemanager/v1beta1/certificatemanagercertificate_identity_test.go` verifying parsing, validation, and normalization.
- **Impact**: `CertificateManagerCertificate` matches the modern direct controller identity and references conventions, enabling fully safe type-based operations across the controller layer.
