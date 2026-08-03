# AssuredWorkloadsWorkload Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Stability Label**:
   - Added `cnrm.cloud.google.com/stability-level: alpha` as a metadata label to the schema.

2. **Identity & Reference**:
   - Implemented identity under `apis/assuredworkloads/v1alpha1/assuredworkloadsworkload_identity.go`.
   - The GCP resource URL structure for Assured Workloads is organization-scoped rather than project-scoped: `organizations/{organization}/locations/{location}/workloads/{workload}`.
   - We leveraged `refs.ResolveOrganizationID` to resolve the organization ID from either `spec.organizationRef.external` or the `cnrm.cloud.google.com/organization-id` annotation in structured or unstructured forms.
   - Replaced old status-based fallback in Normalize with the modern `refs.Normalize` which queries `status.externalRef` directly.

3. **Validation and Golden Files**:
   - Ran unit tests for identity and successfully updated CAIS golden output `_identities.yaml` files via `TestGoldenIdentitiesYamlFiles`.
   - Ran `dev/ci/presubmits/tests-e2e-fixtures-assuredworkloads` successfully to confirm the resource works seamlessly against simulated/mock GCP APIs.
