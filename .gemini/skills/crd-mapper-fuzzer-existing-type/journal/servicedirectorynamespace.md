# Journal: ServiceDirectoryNamespace Transition to Direct KRM types

## Learnings & Observations

### 1. Comment format issue in Go structs
Using block comments `/* ... */` to describe fields that contain GCP paths like `projects/*/locations/*/namespaces/*` is highly dangerous because the `*/` in the path prematurely terminates the block comment. The remaining text is then parsed as Go code, causing compilation and `controller-gen` failures.
**Recommendation:** Always use line comments `//` instead of block comments `/* ... */`.

### 2. Handling of missing ObservedState in compiler/mapper vs CRD schema
When a baseline CRD's `status` doesn't have an `observedState` field, we must not add it to the `ServiceDirectoryNamespaceStatus` Go struct to preserve strict schema compatibility.
However, because the GCP proto contains output-only fields (e.g., `create_time`, `update_time`, `uid`), the `generate-mapper` tool expects a `ServiceDirectoryNamespaceObservedState` struct to map these fields.
**Solution:** Define an empty `ServiceDirectoryNamespaceObservedState` struct in `namespace_types.go`, but do NOT reference it as a field inside `ServiceDirectoryNamespaceStatus`. This allows the generated mapper to compile successfully while ensuring `controller-gen` does not generate an `observedState` property in the final CRD.
