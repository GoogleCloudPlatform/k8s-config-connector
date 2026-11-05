# ComputeOrganizationSecurityPolicyAssociation Identity & Reference Migration Journal

## Observations & Learnings

- **Resource Scope & Parent Hierarchy**: `ComputeOrganizationSecurityPolicyAssociation` represents the association between a `ComputeOrganizationSecurityPolicy` and an organization or folder attachment target in GCP Compute Engine.
- **URL Format**: The resource is identified by the parent security policy ID and the association name:
  `locations/global/securityPolicies/{securityPolicy}/associations/{association}`
  Full URL format: `//compute.googleapis.com/locations/global/securityPolicies/{securityPolicy}/associations/{association}`
- **CAIS Exemption**:
  `//compute.googleapis.com/locations/global/securityPolicies/{}/associations/{}` is not in `cloudassetinventory_names.jsonl` (similar to several other compute subresources and association types). An exception entry was added to `pkg/gcpurls/registry_test.go` to keep `TestRegisteredTemplatesMatchCAI` passing.
- **Spec Resolution & External Ref Trimming**:
  `obj.Spec.PolicyID` can be provided as a numeric ID, a relative path (`locations/global/securityPolicies/{id}`), or a full Compute URL (`https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/{id}`). `parseSecurityPolicyID` uses `apirefs.TrimComputeURIPrefix` and parses the security policy ID reliably.
- **Resource ID & Name**:
  The association name is resolved via `refs.GetResourceID(obj)` (defaulting to `metadata.name` if `spec.resourceID` is omitted).
- **Status Cross-Check**:
  `ComputeOrganizationSecurityPolicyAssociationStatus` does not expose an external reference or unique ID field, so no status cross-checking is performed during `GetIdentity`, returning the parsed spec identity cleanly without altering the CRD schema.
