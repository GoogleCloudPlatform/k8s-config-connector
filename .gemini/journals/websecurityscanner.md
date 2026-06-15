# WebSecurityScanner Service Journal

### 2026-06-15 Initial Scaffolding and Identity for WebSecurityScannerScanConfig
- **Context**: Greenfield implementation of WebSecurityScannerScanConfig types, CRD, and IdentityV2 under `websecurityscanner.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: Custom account and Google account login forms in `ScanConfig` have password fields which represent credentials in GCP API. Storing plaintext passwords violates security best practices.
- **Solution**: Annotated the password fields in hand-coded `ScanConfig_Authentication_GoogleAccount` and `ScanConfig_Authentication_CustomAccount` structs to use KCC's standard secure secret reference type `*refsv1beta1secret.Legacy`, while mapping to the correct proto fields using `+kcc:proto:field` annotations. This is fully supported by the direct toolchain mapper.
- **Impact**: Ensures that user passwords/credentials are always stored securely inside Kubernetes Secrets rather than in plaintext in KRM Specs.
