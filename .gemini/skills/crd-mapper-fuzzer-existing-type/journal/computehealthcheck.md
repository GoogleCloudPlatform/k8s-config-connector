# Journal: ComputeHealthCheck transition to direct KRM types

## Observations & Key Learnings

1. **Schema Compatibility Validation**:
   - The Go types for `ComputeHealthCheck` already existed under `apis/compute/v1beta1/healthcheck_types.go`.
   - Running `dev/tasks/diff-crds` returned empty output, verifying that the KRM schema defined by these Go types is 100% strictly compatible with the baseline CRD schema.

2. **Fuzzer Implementation**:
   - Implemented a complete fuzzer under `pkg/controller/direct/compute/healthcheck_fuzzer.go`.
   - Mapped all KRM Spec and Status fields to enable extensive round-trip fuzz validation.

3. **Ignoring Unmapped Proto Fields**:
   - The underlying GCP proto `google.cloud.compute.v1.HealthCheck` contains the field `grpc_tls_health_check`. This field is not exposed in the baseline `ComputeHealthCheck` CRD schema.
   - Handled this by marking `.grpc_tls_health_check` as unimplemented (`f.Unimplemented_NotYetTriaged(".grpc_tls_health_check")`) in the fuzzer, allowing the round-trip validation to pass completely with 100,000 fuzzing iterations.
