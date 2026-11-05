# Fuzzer Implementation Journal: OSLoginSSHPublicKey

## Observations and Learnings

- **Centralized Fuzz Testing**: The fuzzer for `OSLoginSSHPublicKey` is registered in `pkg/controller/direct/oslogin/osloginsshpublickey_fuzzer.go` and exported via `pkg/controller/direct/register/register.go`. Following the `create-fuzzer` guidelines, standalone test files like `osloginsshpublickey_fuzzer_test.go` inside the direct controller directory are removed, ensuring all fuzz test execution runs strictly under the central fuzz test suite (`pkg/fuzztesting/fuzztests`).
- **Protobuf and Schema Alignment**:
  - The proto message `google.cloud.oslogin.common.SshPublicKey` contains fields `.key`, `.expiration_time_usec`, `.fingerprint`, and `.name`.
  - In KRM Spec:
    - `.key` maps directly to `spec.key`.
    - `.expiration_time_usec` (int64 in proto) maps to `spec.expirationTimeUsec` (*string in KRM) handled via handcoded mappers in `mappers.go`.
    - Parent and acquisition fields `spec.user`, `spec.project`, and `spec.resourceID` are not part of the `SshPublicKey` payload.
  - In KRM Status:
    - `.fingerprint` maps to `status.fingerprint`.
  - Identity:
    - Proto `.name` is the full GCP resource identity (`users/{user}/sshPublicKeys/{fingerprint}`) and marked via `f.Unimplemented_Identity(".name")`.
- **Field Comparison Documentation**: Added structured field comparison comments directly above fuzzer field registrations in `osloginsshpublickey_fuzzer.go`.
- **Validation**: All round-trip mappings execute cleanly and pass without errors under the central fuzz test suite.
