# Journal: ComputeOrganizationSecurityPolicyAssociation Transition to Direct KRM types

## Learnings & Observations

### 1. Underlying Proto Casing and File Naming Conventions
- **Problem**: In GCP Compute API, the resource body expected for the association is represented by `google.cloud.compute.v1.FirewallPolicyAssociation` proto message.
- **Solution**: We configured `generate.sh` to map the KRM `ComputeOrganizationSecurityPolicyAssociation` to `FirewallPolicyAssociation`.
- **Insight**: In order for standard generator tooling (`generate-types` and `generate-mapper`) to function seamlessly without generator panics or untracked duplicates, we must name the Go types file `firewallpolicyassociation_types.go` (lowercase proto message name) rather than matching the KRM Kind name.

### 2. Matching Required Fields and Casing with Custom Mappers
- **Problem**: In KRM types, required spec fields `attachmentId` and `policyId` are defined as non-pointers (`string`). However, in proto, they are pointer strings (`*string`). This causes a build failure due to direct assignment type mismatch in automatically generated mappers.
- **Solution**: We implemented custom spec mapper functions `ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto` and `ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto` in `computeorganizationsecuritypolicyassociation_mappings.go`. The generator automatically detected these existing functions and skipped/commented out the conflicting generated versions in `mapper.generated.go`.

### 3. Maintaining Strict Schema Compatibility
- **Problem**: Direct controllers usually define fields like `observedState` and `externalRef` under Status, and `projectRef` under Spec. However, the baseline CRD for `ComputeOrganizationSecurityPolicyAssociation` did not have these.
- **Solution**: To maintain strict schema compatibility (so that `diff-crds` is completely clean), we omitted `observedState`, `externalRef`, and `projectRef` from the Go types file. By matching the baseline CRD structure exactly, `dev/tasks/diff-crds` produced an empty output, confirming 100% schema parity.
