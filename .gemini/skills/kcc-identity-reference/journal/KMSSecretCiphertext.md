### [2026-06-22] KMSSecretCiphertext Identity & Refs
- **Context**: Migrating `KMSSecretCiphertext` under `apis/kms/v1alpha1/` to the modern `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**:
    1. `KMSSecretCiphertext` is a client-side wrapper resource without a standard GCP-side URL (or CAI representation).
    2. It was missing modern `identity.IdentityV2` and `refs.Ref` implementations.
    3. The CryptoKey parameter was referenced as a string, requiring robust extraction and parsing into parent segments.
- **Solution**:
    1. Implemented `apis/kms/v1alpha1/kmssecretciphertext_identity.go` utilizing `identity.ServerGeneratedIdentity` because ciphertext value is optional in the spec and can be server-generated (retrieved from `status.ciphertext`).
    2. Reused package-level `cryptoKeyFormat` and `cryptoKeyIdentity` declared in `cryptokeyversion_identity.go` to safely parse the required `spec.cryptoKey` field.
    3. Implemented `apis/kms/v1alpha1/kmssecretciphertext_reference.go` to support `KMSSecretCiphertextRef` and register it using `refs.Register`.
    4. Created `apis/kms/v1alpha1/kmssecretciphertext_identity_test.go` to thoroughly test parsing, formatting, validation, and normalization using `cmp.Diff`.
    5. Added a URL template exception to `TestRegisteredTemplatesMatchCAI` in `pkg/gcpurls/registry_test.go` since KMSSecretCiphertext is not tracked by the standard Cloud Asset Inventory.
- **Impact**: `KMSSecretCiphertext` fully conforms to modern Config Connector identity and references patterns, avoiding circular package dependencies and ensuring total type-safety during reconciliation.

### [2026-06-23] Address Review Feedback on KMSSecretCiphertext
- **Context**: Reviewers requested cleaning up `kmssecretciphertext_identity.go` to use the standard `KMSCryptoKeyIdentityFormat` and `KMSCryptoKeyIdentity` from `apis/kms/v1beta1` instead of duplicate parsing vars. Also requested sorting `ignoredTemplates` in `pkg/gcpurls/registry_test.go` alphabetically and grouped by service.
- **Problem**:
    1. Duplicate parsing vars and structs (`cryptoKeyFormat` and `cryptoKeyIdentity`) was sub-optimal.
    2. Map in `pkg/gcpurls/registry_test.go` was unsorted, causing merge conflicts.
    3. Regeneration of `zz_generated.deepcopy.go` was needed.
- **Solution**:
    1. Imported `kmsv1beta1` package and used `kmsv1beta1.KMSCryptoKeyIdentityFormat` directly to parse `obj.Spec.CryptoKey` in `kmssecretciphertext_identity.go`.
    2. Sorted the entire `ignoredTemplates` map alphabetically in `pkg/gcpurls/registry_test.go`, grouped by service host name, and added clean section comments.
    3. Ran `make generate` to regenerate `apis/kms/v1alpha1/zz_generated.deepcopy.go`.
