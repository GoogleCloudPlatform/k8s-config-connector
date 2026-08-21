# Fuzzer Journal: BigtableGCPolicy

## Observations
- `BigtableGCPolicy` is a spec-only resource in Config Connector, mapping directly to `google.bigtable.admin.v2.GcRule`. Because it has no custom status/observed-state mapping, registering it with `RegisterKRMSpecFuzzer` and initializing it with `NewKRMTypedFuzzer` using nil status mappers is ideal.
- The fuzztesting framework's `FuzzStatus` lacked a nil check on the status mapper functions, which led to nil pointer dereferences when executing focused fuzz runs (e.g. `FOCUS=BigtableGCPolicy` or `FOCUS=PubSubTopic`) on spec-only fuzzers.
- We modified `FuzzStatus` in `pkg/fuzztesting/fuzzkrm.go` to safely skip status fuzzing and return when `StatusFromProto` is nil.
- Standard fields in `GcRule` (intersection and union rules) are not yet implemented or exposed in KCC, and were marked as `f.Unimplemented_NotYetTriaged` to satisfy the round-trip fuzz test.
