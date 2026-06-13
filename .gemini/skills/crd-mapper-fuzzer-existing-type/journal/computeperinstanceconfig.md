# ComputePerInstanceConfig Journal - June 13, 2026

## Learnings & Observations

### 1. Handling Pointer-to-String vs Non-Pointer KRM string
- **Problem**: GCE's proto message maps resource references or optional strings as `*string` (e.g. `PreservedStatePreservedDisk.source`), while in the KRM types, we map them as non-pointer `string` because they are marked `required` in the original baseline CRD schema.
- **Mismatch**: This creates a round-trip failure during fuzzing. When a field is `nil` in proto, it gets translated to `""` in KRM, which gets translated back to `&""` (pointer to empty string) in proto. Since `nil != &""`, the round-trip check fails.
- **Solution**: We addressed this by:
  1. Adding a guard in the mapper's `ToProto` converter: `if in.Source != "" { out.Source = &in.Source }`, so that empty strings map back to `nil`.
  2. Specifying a custom `FilterSpec` in our fuzzer that clears out empty/nil source values in `p1` (the original generated proto) beforehand, so both `p1` and `p2` align.

### 2. Fuzzing Spec-only Resources
- **Problem**: When a resource doesn't map any fields from its Status to the underlying GCE proto, trying to fuzz the Status results in empty-proto-to-nil round-trip mismatches.
- **Solution**: We used `fuzztesting.NewKRMTypedSpecFuzzer` which has `NoStatus` as its StatusType parameter and sets Status mapping functions to `nil`. We then registered it using `fuzztesting.RegisterKRMSpecFuzzer()` which automatically skips registering the status fuzzing logic, allowing us to only fuzz the spec.
