# Fuzzer Journal: NotebookInstance

## Observations
- `NotebookInstance` maps directly to the protobuf message `google.cloud.notebooks.v1.Instance`.
- All 36 protobuf fields are registered under either `SpecField`, `StatusField`, or `Unimplemented_Identity` (specifically `.name`), ensuring a 100% complete mapping coverage.
- The KRM Spec fields `zone` and `projectRef` under the `Parent` inline struct map to the GCP parent path segments rather than being direct fields of `pb.Instance` itself. The GCP resource name (fully qualified path) is marked as `Unimplemented_Identity` in the fuzzer.
- The `oneof` fields in `pb.Instance` (specifically `vm_image` and `container_image` which form the `environment` oneof) map cleanly to KRM Spec types. The Go protobuf library's randomized `oneof` field generation ensures that only one of these fields is filled during fuzzing, satisfying round-trip correctness.
- The NotebookInstance fuzzer is registered under the central fuzz test suite by its presence in `pkg/controller/direct/notebooks/` and the package's inclusion in `pkg/controller/direct/register/register.go`.
