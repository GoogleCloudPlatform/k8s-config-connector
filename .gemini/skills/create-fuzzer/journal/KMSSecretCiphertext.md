# Journal: KMSSecretCiphertext Fuzzer Integration

This journal documents the specific details and learnings from implementing and validating the round-trip KRM fuzzer for `KMSSecretCiphertext`.

## Learnings & Observations

- **Resource:** `KMSSecretCiphertext`
- **Location of Controller & Mappers:** None (currently registered to ReconcilerTypeTerraform only).
- **Location of Fuzzer:** `pkg/controller/direct/kms/secretciphertext_fuzzer.go`
- **GCP Proto/API Type:** This resource is a client-side wrapper representing a GCP KMS ciphertext resource (originally a Terraform resource `google_kms_secret_ciphertext`). It does not correspond to a standard REST resource or a direct Proto message in `google.cloud.kms.v1` (since encrypting is an RPC call, not a standalone resource lifecycle).

### 1. Fuzzing Client-Side/Wrapper Resources (NoProto)
- Because `KMSSecretCiphertext` is a client-side wrapper with no standard REST resource or direct Proto message, we used KCC's `KRMTypedFuzzer_NoProto` framework (`RegisterKRMFuzzer_NoProto`).
- We defined a custom `KMSSecretCiphertextAPIParent` struct inside the fuzzer to represent the client-side API schema. This API struct perfectly mirrors both Spec and Status fields, including nested custom reference types like `secret.Legacy` for `Plaintext` and `AdditionalAuthenticatedData`.
- Implementing `FromAPI` and `ToAPI` mappers that copy all fields directly ensured a 100% lossless, reliable round-trip validation under extensive randomized fuzzing.

### 2. Validation
- Run focused mapper tests using:
  ```bash
  FOCUS=KMSSecretCiphertext go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers
  ```
- All fuzzing checks completed successfully without structural diffs, ensuring KRM schemas and fields map losslessly.
