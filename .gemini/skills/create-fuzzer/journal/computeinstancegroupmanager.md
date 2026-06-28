# Journal: ComputeInstanceGroupManager Fuzzer Implementation

## Observations & Key Learnings

1. **Return Value Mismatch**:
   - The existing hand-coded FromProto mapper function `ComputeInstanceGroupManagerStatus_v1beta1_FromProto` returns a value of `krm.ComputeInstanceGroupManagerStatus` instead of a pointer (`*krm.ComputeInstanceGroupManagerStatus`).
   - To avoid modifying the existing mapper's signature and risking breakage of other modules/controllers, we implemented a lightweight wrapper function `computeInstanceGroupManagerStatus_FromProto` directly in the fuzzer file to return a pointer, satisfying the requirements of `NewKRMTypedFuzzer`.

2. **Required Non-Pointer Field vs. Optional Pointer Mismatch**:
   - In KRM Spec, `targetSize` is defined as a non-pointer `int64` (so it's required and defaults to `0` if empty), whereas in the proto, `target_size` is an optional `*int32` pointer.
   - When the fuzzer generated a random proto with a `nil` target size, it translated to `0` in KRM, which then mapped back to a non-nil `&0` in the round-tripped proto. This mismatch (`nil` vs `&0`) caused fuzzer failure.
   - We resolved this cleanly using `f.FilterSpec` to pre-initialize any `nil` `TargetSize` in the source proto to `0`, ensuring consistent round-trip equivalence.

3. **Coverage and Unimplemented Registrations**:
   - We marked several fields that are currently unsupported/unmapped by the direct controller as `Unimplemented_NotYetTriaged` to cleanly pass the lossless verification loop.
