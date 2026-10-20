# ComputeOrganizationSecurityPolicy Identity & Reference Migration Journal

## Observations & Learnings

- **Resource Scope & Parent Hierarchy**: `ComputeOrganizationSecurityPolicy` represents organization security policies in GCP Compute Engine (`compute.googleapis.com`). Despite the resource being associated with an organization or folder (e.g., `organizations/123456` or `folders/123456`), its canonical URI format according to Cloud Asset Inventory (CAI) and GCP compute API is:
  `//compute.googleapis.com/locations/global/securityPolicies/{securityPolicy}`
  where `{securityPolicy}` is the numerical or string policy ID.

- **Server-Generated ID Behavior**:
  - The resource supports `spec.resourceID`, which when provided acts as the desired policy ID, but can also be omitted and generated server-side.
  - The resource status reflects this via `status.policyId`.
  - Therefore, `ComputeOrganizationSecurityPolicyIdentity` implements `identity.ServerGeneratedIdentity` with `HasIdentitySpecified() bool` (checking `i.SecurityPolicy != ""`).
  - In `GetIdentity()`, if `specIdentity.SecurityPolicy` is not specified, it falls back to `status.policyId`. If both are present, they are cross-checked for drift.

- **Reference Normalization**:
  - Because `ComputeOrganizationSecurityPolicy` is a legacy/Terraform resource without `status.externalRef`, `ComputeOrganizationSecurityPolicyRef` implements `refs.NormalizeWithFallback`.
  - The fallback function inspects `status.policyId` on the unstructured object to build the external identity string.
  - `ValidateExternal` and `Normalize` also accept and trim full Compute Engine URL prefixes (e.g. `https://compute.googleapis.com/compute/v1/...`) via `apirefs.TrimComputeURIPrefix`.

- **CAI & Golden Identity Verification**:
  - The template `locations/global/securityPolicies/{securityPolicy}` is registered in CAI under `compute.googleapis.com/SecurityPolicy`.
  - Running `WRITE_GOLDEN_OUTPUT=1 go test -v ./pkg/cli/powertools/cais/...` updated the fixture golden file `pkg/test/resourcefixture/testdata/basic/compute/v1alpha1/computeorganizationsecuritypolicy/organizationsecuritypolicyautogen/_identities.yaml` from `caisURL: unknown` to `caisURL: //compute.googleapis.com/locations/global/securityPolicies/computeorganizationsecuritypolicy${uniqueId}`.
