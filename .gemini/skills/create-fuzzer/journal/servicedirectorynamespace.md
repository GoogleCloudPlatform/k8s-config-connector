# Journal: ServiceDirectoryNamespace Fuzzer Implementation

## Observations & Learnings

### 1. Fuzzing a Resource with Empty KRM Spec and ObservedState
`ServiceDirectoryNamespace` contains only meta-level identifiers like `location` and `projectRef` in its KRM Spec and a single `name` in its KRM Status, while having an empty `ServiceDirectoryNamespaceObservedState`. This results in zero mapped fields between KRM types (`Spec` / `ObservedState`) and the `pb.Namespace` protobuf.
In such cases:
- Every field of the protobuf must be explicitly ignored or marked as identity/unimplemented in the fuzzer.
- The fuzzer still compiles and runs to verify structure mapping, even when all actual fields of the proto are marked as unimplemented (either because they are identity fields, like `.name`, or because they are not yet supported/mapped in KCC's schema, such as `.labels`, `.uid`, `.create_time`, and `.update_time`).
