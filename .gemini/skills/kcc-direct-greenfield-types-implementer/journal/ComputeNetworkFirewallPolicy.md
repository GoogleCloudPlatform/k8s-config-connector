# ComputeNetworkFirewallPolicy Types Implementation Journal

## Mapping Summary
- **GCP Proto Resource**: `google.cloud.compute.v1.FirewallPolicy` (mapping to GCP path `projects/{project}/global/firewallPolicies/{name}`).
- **KRM Kind**: `ComputeNetworkFirewallPolicy`
- **KRM Version**: `v1beta1` (Existing resource, updated to direct KRM types).

## Architecture & Design Decisions
1. **Rule Separation**:
   - The GCP `FirewallPolicy` proto contains a list of nested `rules`. In Config Connector, however, these rules are modeled as separate individual resources via the `ComputeNetworkFirewallPolicyRule` CRD. Thus, `rules` is omitted from `ComputeNetworkFirewallPolicySpec` to maintain user configuration standards and decouple rule management from policy management.
2. **ObservedState**:
   - Fields such as `creationTimestamp`, `fingerprint`, `id`, `ruleTupleCount`, `selfLink`, and `selfLinkWithId` are modeled inside `ComputeNetworkFirewallPolicyObservedState` under `status.observedState`. This aligns with direct controller standards while preserving the ability of other resources to reference this policy (e.g., via `networkFirewallPolicyId`).
3. **Labels**:
   - Configured `cnrm.cloud.google.com/stability-level=stable` to reflect its promotion/stability as a `v1beta1` stable resource.
