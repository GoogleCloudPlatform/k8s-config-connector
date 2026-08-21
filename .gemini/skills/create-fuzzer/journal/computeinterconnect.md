# Fuzzer Journal: ComputeInterconnect

## Resource Details
- **Kind**: `ComputeInterconnect`
- **Group**: `compute.cnrm.cloud.google.com`
- **Version**: `v1alpha1`
- **Proto**: `google.cloud.compute.v1.Interconnect`

## Observations & Findings
1. **Wrapper-Helper Transition**: The existing fuzzer (`computeinterconnect_fuzzer.go`) was written in 2024 using legacy `.Insert` operations (such as `f.SpecFields.Insert` and `f.StatusFields.Insert`). We refactored it to use the clean wrapper-helpers `f.SpecField()`, `f.StatusField()`, and `f.Unimplemented_Identity()`.
2. **Field Mapping Documentation**: Added explicit comments comparing KRM Spec fields to their protobuf counterparts, verifying complete coverage of the spec mapping.
3. **Lossless Round-trip Validation**: Validated the updated fuzzer centrally under `pkg/fuzztesting/fuzztests/...`, confirming zero data-loss during bidirectional transformation.
