### [2026-06-21] SecureSourceManagerInstance Identity & Refs
- **Context**: Moving SecureSourceManagerInstance to standard `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**:
    1. SecureSourceManagerInstance had an existing identity and reference structure, but `getIdentityFromSecureSourceManagerInstanceSpec` accepted `client.Object` instead of a typed `*SecureSourceManagerInstance`.
    2. There were no unit tests written for `SecureSourceManagerInstance` identity or reference.
    3. `ParentString` method was not implemented on `SecureSourceManagerInstanceIdentity`.
- **Solution**:
    1. Added `ParentString` to `SecureSourceManagerInstanceIdentity`.
    2. Updated `getIdentityFromSecureSourceManagerInstanceSpec` signature to accept a typed `*SecureSourceManagerInstance` directly.
    3. Updated the `Normalize` fallback method in `instance_reference.go` to convert the unstructured object using `common.ToStructuredType` before invoking the getter.
    4. Added comprehensive unit tests in `instance_identity_test.go` using `github.com/google/go-cmp/cmp` to verify formatting, parsing, and parent paths.
- **Impact**: SecureSourceManagerInstance is fully compliant with the modern standards for direct controller identities and references, with solid test coverage.
