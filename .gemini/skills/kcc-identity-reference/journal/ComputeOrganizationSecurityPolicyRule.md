# ComputeOrganizationSecurityPolicyRule Identity & Reference Migration Journal

## Observations & Learnings

- **Resource Scope & Parent Hierarchy**: `ComputeOrganizationSecurityPolicyRule` represents a rule within an organization security policy in GCP Compute Engine (`locations/global/securityPolicies/{securityPolicy}/rules/{priority}`). Its parent is the `ComputeOrganizationSecurityPolicy` (`locations/global/securityPolicies/{securityPolicy}`).
- **Resource ID & Priority Mapping**: In GCP REST and proto (`google.cloud.compute.v1.SecurityPolicyRule`), the identifier for a rule is its integer priority (e.g. `1000`). In KCC's `ComputeOrganizationSecurityPolicyRuleSpec`, `priority` is configured via `spec.resourceID` (or defaults to `metadata.name` if unset). Therefore, the identity template uses `{priority}` as the rule identifier placeholder.
- **Spec Parent ID Parsing**: In KCC, `spec.policyId` is an immutable string field referencing the parent security policy. Similar to `ComputeOrganizationSecurityPolicyAssociation`, `policyId` can be specified in several formats:
  - Plain numeric or policy ID string (e.g., `"1234567890"`).
  - Prefix with `securityPolicies/` (e.g., `"securityPolicies/1234567890"`).
  - Relative URI (e.g., `"locations/global/securityPolicies/1234567890"`).
  - Full Compute REST URL (e.g., `"https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/1234567890"`).
  A custom helper `parseSecurityPolicyRuleParentID` normalizes all these forms to extract the plain security policy ID.
- **Canonical GCP URL & CAIS Format**: The canonical GCP URL template is `locations/global/securityPolicies/{securityPolicy}/rules/{priority}` under host `compute.googleapis.com`.
- **CAIS Exemption**: Similar to `ComputeOrganizationSecurityPolicyAssociation`, `//compute.googleapis.com/locations/global/securityPolicies/{}/rules/{}` is not in `docs/ai/metadata/cloudassetinventory_names.jsonl` because organization security policy rules are subresources managed via `compute.organizationSecurityPolicies.addRule/patchRule/removeRule/getRule`. An exemption entry was registered in `pkg/gcpurls/registry_test.go` to keep `TestRegisteredTemplatesMatchCAI` passing.
- **Legacy Fallback & Normalization Pattern**: `ComputeOrganizationSecurityPolicyRule` is currently a legacy Terraform-based resource without `status.externalRef`. Following skill instructions, `ComputeOrganizationSecurityPolicyRuleRef.Normalize` uses `refs.NormalizeWithFallback`, checking that the target resource's `Ready` condition is `True` before building the external identity string from spec. If the resource is not ready, it returns an empty string to raise a `ReferenceNotReadyError`.
- **Status Cross-Check**: `ComputeOrganizationSecurityPolicyRuleStatus` does not expose an external reference or unique ID field. In accordance with the skill mandate, no schema changes were introduced, and `GetIdentity` returns the parsed spec identity cleanly without performing a status cross-check.
- **DeepCopy Generation**: Generated deepcopy code for `ComputeOrganizationSecurityPolicyRuleRef` in `apis/compute/v1alpha1/zz_generated.deepcopy.go` via `dev/tasks/generate-crds`.
