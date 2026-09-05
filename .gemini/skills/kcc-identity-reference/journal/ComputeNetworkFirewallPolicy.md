# ComputeNetworkFirewallPolicy Identity & Reference Migration Journal

## Observations & Learnings

- **Resource Scope:** `ComputeNetworkFirewallPolicy` is a project-scoped global resource. Consequently, its canonical identity URL is of the format `projects/{project}/global/firewallPolicies/{firewallPolicy}`.
- **User-Specified ID:** Unlike the organization-scoped `ComputeFirewallPolicy` which has server-generated IDs, `ComputeNetworkFirewallPolicy` uses a user-specified name/ID (using `spec.resourceID` or `metadata.name`). We implement `identity.IdentityV2` to capture this.
- **Reference Refactoring:** Defined `ComputeNetworkFirewallPolicyRef` struct and standard methods in `apis/compute/v1beta1/computenetworkfirewallpolicy_reference.go`.
