# ComputeNetworkFirewallPolicyAssociation Greenfield Types Implementation Journal

## Observations & Implementation Details

1. **Unreachable Proto Types**:
   - The code generator identified `FirewallPolicyAssociation` as an "unreachable type" because it wasn't referenced directly by other types the generator walked automatically.
   - Consequently, the proto fields (`attachment_target`, `display_name`, `firewall_policy_id`, `name`, `short_name`) were put inside `types.generated.go` as commented-out structures.
   - We leveraged these commented-out structures to hand-craft `ComputeNetworkFirewallPolicyAssociationSpec` and `ComputeNetworkFirewallPolicyAssociationObservedState` fields with precise `+kcc:proto:field` annotations.

2. **First-Class Reference Fields**:
   - We avoided the generic `v1alpha1.ResourceRef` and instead implemented modern, strongly typed KCC reference types:
     - `AttachmentTargetRef` -> `*computev1beta1.ComputeNetworkRef`
     - `FirewallPolicyRef` -> `*refsv1beta1.ComputeFirewallPolicyRef`
     - `ProjectRef` -> `*refsv1beta1.ProjectRef`
   - This ensures type safety, better CRD validation schemas, and complete compatibility with modern direct controllers.

3. **Validation & Metadata**:
   - Appended the `cnrm.cloud.google.com/stability-level=alpha` label to mark this resource's new direct API status.
   - Added appropriate `// +required` tags to ensure spec fields are validated strictly at the CRD/schema level.
