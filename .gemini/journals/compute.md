# Compute Service Journal

### 2026-06-13 Multi-Version Storage Version Annotation & Schema Reconstruction
- **Context**: Implementing direct KRM types for `ComputeResourcePolicy` in `v1alpha1` (previously only existed under `v1beta1`).
- **Problem**: 
  1. Adding a new version of `ComputeResourcePolicy` caused `controller-gen` to fail with `CRD for ComputeResourcePolicy.compute.cnrm.cloud.google.com has no storage version`.
  2. The generated type `apis/compute/v1alpha1/resourcepolicy_types.go` initially only contained `ProjectRef`, `Location`, and `ResourceID` fields, while all sub-types of `ResourcePolicy` generated in `types.generated.go` were commented out as unreachable.
- **Solution**: 
  1. Added `// +kubebuilder:storageversion` to the stable version (`v1beta1`) of `ComputeResourcePolicy` in `apis/compute/v1beta1/resourcepolicy_types.go`.
  2. Structured the `ComputeResourcePolicySpec` in `apis/compute/v1alpha1/resourcepolicy_types.go` to reference the generated sub-types (like `ResourcePolicyDiskConsistencyGroupPolicy`, `ResourcePolicyGroupPlacementPolicy`, etc.), which instructed the generator during the next run of `./generate.sh` to automatically uncomment those types in `types.generated.go` and successfully produce the correct CRD.
- **Impact**: When migrating existing Compute resources from Terraform (`v1beta1`) to direct KRM types under a new version, remember to add `// +kubebuilder:storageversion` to `v1beta1`, and reference generated sub-types in the new spec file so the generator automatically unprunes them.
