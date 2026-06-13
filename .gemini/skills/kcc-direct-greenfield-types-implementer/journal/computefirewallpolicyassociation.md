### 2026-06-13 ComputeFirewallPolicyAssociation Direct Types Scaffolding
- **Context**: Implementing direct KRM types and generate.sh configuration for `ComputeFirewallPolicyAssociation`.
- **Problem**: 
  - `ComputeFirewallPolicyAssociation` is an organization/folder-scoped resource, meaning it does not belong to a specific GCP Project or regional Location. However, the `generate-types` tool scaffolds these fields (`ProjectRef`, `Location`) by default.
  - The resource already existed in `v1beta1` under KCC (as a DCL resource). Scaffolding it under `v1alpha1` would have downgraded the resource version in the generated CRD schema, potentially breaking existing deployments.
- **Solution**:
  - Generated and structured the direct types directly under `apis/compute/v1beta1/firewallpolicyassociation_types.go` instead of `v1alpha1` to maintain version parity.
  - Manually stripped out `ProjectRef` and `Location` from the Spec.
  - Maintained backward compatibility by keeping `AttachmentTargetRef` as a generic `k8sv1alpha1.ResourceRef` (as it can reference either a Folder or Organization) while using a typed `refsv1beta1.ComputeFirewallPolicyRef` for `FirewallPolicyRef`.
- **Impact**: Ensures schema compatibility and preserves existing `v1beta1` deployments when migrating the underlying controller from DCL to a direct KRM model.
