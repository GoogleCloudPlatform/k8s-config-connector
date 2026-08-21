# ComputeResourcePolicy Direct KRM Transition Journal

## Learnings

### 1. Pointer Field Mappings for Optional/OneOf Boolean Fields
For Compute API resources (where fields are generated as optional/oneof in protobuf-to-go Go structs), using `direct.LazyPtr(in.GetX())` instead of `in.X` during `FromProto` conversion causes fuzzer failures.
- **Problem**: `in.GetX()` returns the zero value (`false` for bool, `""` for string) if the field is `nil` in proto. `direct.LazyPtr(...)` then wraps this zero-value into a non-nil pointer (e.g. pointer-to-false). When mapped back in `ToProto`, it serializes to `false`, causing a round-trip difference from the original `nil` value.
- **Solution**: Avoid `direct.LazyPtr(in.GetX())` for optional fields that can be directly mapped as pointers (e.g., `out.GuestFlush = in.GuestFlush`).
