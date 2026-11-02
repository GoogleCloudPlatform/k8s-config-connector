# Journal: ComputeOrganizationSecurityPolicyRule MockGCP Implementation

This journal captures findings and decisions while implementing the `ComputeOrganizationSecurityPolicyRule` mock in MockGCP and matching real GCP behavior.

## Observations & Implementation Decisions

1. **GCP API Scope for Organization Security Policies**:
   - Modern GCP computes organization security policies as Cloud Armor policies (`type: CLOUD_ARMOR`). The ancient alpha/beta `FIREWALL` versioned expressions and Layer 4 configurations (`layer4Config`, `direction`) are rejected by real GCP on modern organizations (`Can not specify direction on security policy rule for Cloud Armor`).
   - The supported matcher for Cloud Armor organization security policy rules uses `versionedExpr: "SRC_IPS_V1"` with `config.srcIpRanges`.
   - Updated the CRD types and Terraform provider schema to allow `SRC_IPS_V1` and make `layer4Config` optional.

2. **Server-Generated ID & SecurityPolicy Reference**:
   - Organization security policies in GCP have server-generated IDs (e.g. `123456789012`).
   - Similar to `ComputeOrganizationSecurityPolicyAssociation`, `ComputeOrganizationSecurityPolicyRule` requires `securityPolicyRef` to resolve the parent policy ID dynamically.
   - Updated `ComputeOrganizationSecurityPolicyRuleSpec` to support `securityPolicyRef` alongside `policyId`.
   - Configured `resourceReferences` in `servicemappings/compute.yaml` and normalized `policy_id` in Terraform provider's Create, Read, Update, Delete, and Import functions to standard `locations/global/securityPolicies/<id>` paths.

3. **HTTP 400 on Non-Existent Rule**:
   - Real GCP returns HTTP 400 (`InvalidArgument`) with message `The security policy does not contain a rule at priority <X>` when `getRule` is queried for a non-existent rule.
   - Because standard Terraform `HandleNotFoundError` only checks HTTP 404, reading a non-existent rule during initial create caused an immediate failure.
   - Implemented `transformSecurityPolicyRuleReadError` in the provider to translate this 400 to 404, matching the pattern in `ComputeOrganizationSecurityPolicyAssociation`.
   - Aligned MockGCP's `GetRule` to return `status.Errorf(codes.InvalidArgument, ...)` matching GCP's exact error message.

4. **PatchRule Repeated Field Handling**:
   - In MockGCP `PatchRule`, using `proto.Merge` on repeated fields like `srcIpRanges` appends entries on every update, resulting in duplicates that trigger continuous re-reconciliation loops.
   - Replaced `proto.Merge` with field replacement for `match` and other configurable fields.

5. **Test Alignment & Normalization**:
   - Added `ComputeOrganizationSecurityPolicy` placeholder `${securityPolicyID}` to `getPlaceholdersForKind` in `golden_alignment_test.go` so dependency policy CRUD requests are filtered out cleanly when testing rules.
   - Handled server-assigned `ruleNumber` normalization in `mockcompute/normalize.go` and `golden_alignment_test.go`.
