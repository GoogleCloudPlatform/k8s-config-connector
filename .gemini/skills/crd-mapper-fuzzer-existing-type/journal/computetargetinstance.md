# Journal: ComputeTargetInstance

## Learnings & Observations

1. **Mapping of Custom References**:
   - `ComputeTargetInstance` references `ComputeInstance`, `ComputeNetwork`, and `ComputeSecurityPolicy`.
   - By naming the Go struct fields exactly (e.g. `InstanceRef *InstanceRef`, `NetworkRef *ComputeNetworkRef`, `SecurityPolicyRef *ComputeSecurityPolicyRef`), and using structures containing `external`, `name`, and `namespace` fields, KCC's OpenAPI schema builder (`add-validation-to-crds`) automatically appends standard validation blocks (such as `oneOf` blocks for reference validation) on these custom fields.

2. **Strict Schema Compatibility Constraints**:
   - When transitioning an existing CRD type to a direct Go type definition, we must retain labels such as `cnrm.cloud.google.com/tf2crd=true` to preserve register metadata correctly.
   - We must also ensure that fields like `spec.projectRef` or `status.observedState` are NOT introduced if they were not part of the baseline CRD, to preserve 100% identical CRD schemas.
   - In Go struct definitions, omitting the `// +required` comment block on the `Spec` field can cause the generated CRD to drop `required: [spec]`. We must explicitly add `// +required` on the `Spec` field to match the original schema perfectly.

3. **GCP Proto Casing Mismatches (`NATPolicy`)**:
   - The GCP protobuf compiler maps snake-cased `nat_policy` to `NatPolicy` in proto structures, but KCC's generator expected `NATPolicy` or encountered a near-miss mismatch against KRM Spec's Go struct field.
   - Renaming the KRM Spec Go struct field from `NatPolicy` to `NATPolicy` while preserving the JSON tag `json:"natPolicy,omitempty"` completely resolved the mapper generator near-miss, allowing clean automatic mappings for that field.

4. **Handcoded Mapper for Type Mismatch**:
   - The GCP proto `Zone` field is of type `*string` (optional string pointer), while the baseline KRM Spec defines `Zone` as a required `string`.
   - Because of this `string` vs `*string` static typing mismatch in Go, direct assignment in the generated mapper fails compilation.
   - Placing our custom conversion mapping functions `ComputeTargetInstanceSpec_v1beta1_FromProto` and `ComputeTargetInstanceSpec_v1beta1_ToProto` in a hand-crafted mapping file (`computetargetinstance_mapper.go`) was extremely successful. The generator automatically detected the handcoded functions and skipped generating duplicate conflicting versions, resolving all static compilation errors.
