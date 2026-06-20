# Journal: ComputeSecurityPolicy MockGCP Implementation

This journal captures the details and findings of implementing the `ComputeSecurityPolicy` mock in MockGCP.

## Observations

1. **Terraform Update Behavior**:
   Unlike many GCP resources that update the entire specification in a single `Patch` or `Update` call, the `ComputeSecurityPolicy` (associated with the `google_compute_security_policy` Terraform resource) uses a collection of granular rule-based methods to update rules:
   - `AddRule`: to append a new security policy rule.
   - `PatchRule`: to update fields/actions of an existing rule matching a specific priority.
   - `RemoveRule`: to delete a security policy rule.
   - `Patch`: to modify top-level policy configurations (like `AdaptiveProtectionConfig`).

2. **Mocking Granular Methods**:
   To accurately mock `ComputeSecurityPolicy`, we had to implement all of the above methods on the `SecurityPoliciesV1` service under `mockgcp/mockcompute/securitypoliciesv1.go`.

3. **Rule Ordering**:
   In real GCP, security policy rules are evaluated and ordered by their priority. To replicate this behavior and ensure deterministic output in mock tests, we sort the rules list by the `Priority` field after any mutating action (`Insert`, `AddRule`, `PatchRule`, `RemoveRule`).

4. **Resource Verification**:
   The E2E test `computesecuritypolicy` under `pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computesecuritypolicy` was successfully run against both real GCP (recording a pristine `_http.log`) and MockGCP. MockGCP successfully intercept and mock all interactions.
