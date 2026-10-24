# ComputeMachineImage Fuzzer Journal

## Observations & Learnings

- **Fuzzer Implementation Details:** The KRM fuzzer for `ComputeMachineImage` is located in `pkg/controller/direct/compute/computemachineimage_fuzzer.go` under package `compute`.
- **Central Fuzz Testing Integration:** Confirmed that `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute"` is imported and registered in `pkg/controller/direct/register/register.go` so the fuzzer executes under the central fuzz test suite.
- **Subfield Mapping Details:**
  - `machine_image_encryption_key`: Mapped to `pb.CustomerEncryptionKey`. The subfields `kms_key_name`, `kms_key_service_account`, `raw_key`, and `sha256` are round-tripped as `SpecField`s. The proto field `rsa_encrypted_key` is marked as `Unimplemented_NotYetTriaged`.
- **Detailed Comparisons Documentation:** Added structured comparison documentation inside `computemachineimage_fuzzer.go` explicitly mapping each KRM Spec and Status field to its Protobuf counterpart and fuzzer status.
- **Testing Verification:**
  - Focused test: `FOCUS=ComputeMachineImage go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` (PASS)
  - Full suite: `go test -count=1 -v ./pkg/fuzztesting/fuzztests/` (PASS)
