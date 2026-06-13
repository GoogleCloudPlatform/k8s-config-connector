# ComputeOrganizationSecurityPolicyRule journal

We migrated/implemented the direct KRM types for `ComputeOrganizationSecurityPolicyRule` (which maps to `google.cloud.compute.v1.SecurityPolicyRule`).

## Key Observations & Learnings

- **Proto vs. REST Mismatch**: The public `google.cloud.compute.v1.SecurityPolicyRule` proto type is used for both global security policies (Cloud Armor) and organization/hierarchical firewall policies in GCP. However, the public proto definition only includes fields suitable for global security policies (e.g. omitting `direction`, `enableLogging`, `targetResources`, and `targetServiceAccounts`). These fields are accepted/needed by the REST API for organization-level security policies.
- **Manual Override Strategy**: To resolve the mismatch, we manually declared these missing fields directly in `securitypolicyrule_types.go`, while maintaining compatibility with both GCE's REST payload representation and the established direct resource design patterns.
- **Reference Type Creation**: We implemented `ComputeOrganizationSecurityPolicyRef` from scratch in `organizationsecuritypolicy_reference.go` to provide typed reference capabilities (`policyRef`) referencing the parent `ComputeOrganizationSecurityPolicy` GVK, ensuring full type safety and adherence to KCC's references standards.
