# ComputeTargetTCPProxy KRM Mapper & Fuzzer Journal

- **Acronyms and capitalization**: The KRM kind uses `ComputeTargetTCPProxy` while the Go protobuf types and generation script use `ComputeTargetTcpProxy:TargetTcpProxy` which maps back to the proto message `TargetTcpProxy`.
- **Fuzzer Refactoring**: Refactored the pre-existing fuzzer `pkg/controller/direct/compute/targettcpproxy_fuzzer.go` to adhere to the type-safe helpers defined on `KRMTypedFuzzer`. Changed `.Insert` calls on raw fields to `.SpecField()`, `.StatusField()`, `.Unimplemented_Identity()`, and `.Unimplemented_Internal()`. This improves code safety and ensures alignment with direct controller standards.
- **Top-level mapping skipped**: The top-level `+kcc:proto` was not needed on the main `ComputeTargetTCPProxy` struct itself, as KRM spec and status are mapped individually from the protobuf `TargetTcpProxy` message (meaning adding it would result in generating empty/unreachable mappings for the main KRM struct).
