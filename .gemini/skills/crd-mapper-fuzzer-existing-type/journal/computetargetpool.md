# ComputeTargetPool Journal

During the transition of `ComputeTargetPool` to a direct KRM type, the following observations and learnings were made:

## 1. Schema Compatibility for Legacy Fields
- `ComputeTargetPool` did not originally contain `spec.projectRef` or `spec.location` in its baseline CRD, but rather `spec.region`.
- To maintain strict schema-compatibility, we avoided adding `spec.projectRef` and defined `spec.region` as required, matching the baseline CRD schema exactly.
- All reference fields like `backupTargetPoolRef`, `httpHealthCheckRef`, `instances`, and `securityPolicyRef` do not have a `kind` field in the baseline CRD, so we hand-coded a custom reference type `TargetPoolResourceRef` without a `kind` field.

## 2. Fuzz Testing without Status Mapping
- Since `ComputeTargetPool` did not have custom status mapping functions (only standard `Conditions`, `ObservedGeneration` and `SelfLink`), passing `nil` status mapping arguments to `fuzztesting.NewKRMTypedFuzzer` can cause:
  1. **Go type inference failure** at compilation time. This can be resolved by supplying explicit type parameters to the generic function: `fuzztesting.NewKRMTypedFuzzer[*pb.TargetPool, krm.ComputeTargetPoolSpec, krm.ComputeTargetPoolStatus]`.
  2. **Nil pointer dereference panics** during the fuzzing of status fields. This can be completely bypassed by registering the fuzzer using `fuzztesting.RegisterKRMSpecFuzzer` instead of `fuzztesting.RegisterKRMFuzzer`, which runs only spec fuzz tests and skips status fuzzing.
