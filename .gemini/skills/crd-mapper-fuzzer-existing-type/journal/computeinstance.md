# Journal: ComputeInstance transition to direct KRM types

## Observations & Key Learnings

1. **Configured generate.sh**:
   - Configured `apis/compute/v1beta1/generate.sh` to include `ComputeInstance:Instance` as a resource and passed the `--include-skipped-output` flag.
   - This ensures that KRM types under `apis/compute/v1beta1/instance_types.go` are registered properly, and auto-generation for DeepCopy and client clientsets are executed seamlessly.

2. **Strict Schema Compatibility (diff-crds)**:
   - Verified that the newly direct-enabled schema for `ComputeInstance` aligns 100% with its baseline counterpart.
   - Running `dev/tasks/diff-crds` yields an absolutely empty/zero output, confirming there are no schema drift or unintended breaking API changes.

3. **Handcoded Spec & Status Mapping**:
   - Created handcoded mapper functions in `pkg/controller/direct/compute/instance_mapper.go` to handle custom mappings.
   - **Resource Policies**: Handled mapping between `[]krm.InstanceResourceRef` (which uses reference names/external strings) and proto's string slice `ResourcePolicies`.
   - **Metadata**: Key-value pairs inside `InstanceMetadata` are mapped to proto's structured `pb.Metadata` array of `pb.Items` (which contains pointer string items).
   - **Sub-struct placeholders**: Generated placeholders returning nil for `InstanceAdvancedMachineFeatures`, `InstanceConfidentialInstanceConfig`, `InstanceNetworkPerformanceConfig`, `InstanceParams`, `InstanceReservationAffinity`, `InstanceScheduling`, and `InstanceShieldedInstanceConfig` for future controller reconciliation steps.
