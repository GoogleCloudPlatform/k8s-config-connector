# Journal: ComputeOrganizationSecurityPolicyAssociation MockGCP Implementation

This journal captures the details and findings of implementing the `ComputeOrganizationSecurityPolicyAssociation` mock in MockGCP and matching real GCP behavior.

## Observations & Learnings

1. **Security Policy Reference Resolution**:
   - `ComputeOrganizationSecurityPolicy` uses server-generated numeric policy IDs.
   - For `ComputeOrganizationSecurityPolicyAssociation`, the association must link to a valid security policy ID.
   - Added `securityPolicyRef` to `ComputeOrganizationSecurityPolicyAssociationSpec` in `apis/compute/v1alpha1/computeorganizationsecuritypolicyassociation_types.go` and mapped it in `scripts/resource-autogen/generated/servicemappings/compute.yaml` and `config/servicemappings/compute.yaml`.
   - Used `valueTemplate: "locations/global/securityPolicies/{{value}}"` to ensure the resolved reference matches the full URL path expected by the Terraform provider and live state comparison.

2. **Terraform Provider URL & Id Handling**:
   - In `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute/resource_compute_organization_security_policy_association.go`, normalized `policy_id` across `Create`, `Read`, `Delete`, and `Import` using a helper `normalizeSecurityPolicyId` to ensure the `locations/global/securityPolicies/` prefix is uniformly handled.
   - This prevents diff mismatches on immutable fields during re-reconciliation.

3. **GCP Attachment Permissions**:
   - In real GCP, creating an association directly at the organization level (`organizations/<org-id>`) requires `compute.organizations.setSecurityPolicy`.
   - Attaching to a project (`projects/<project-number>`) succeeds with standard test project credentials and correctly executes the full lifecycle (`addAssociation`, `getAssociation`, `removeAssociation`).

4. **Error Format Alignment**:
   - When querying a non-existent association via `getAssociation`, GCP Compute Engine returns HTTP 400 Bad Request with:
     ```json
     {
       "error": {
         "code": 400,
         "errors": [
           {
             "domain": "global",
             "message": "Invalid value for field 'name': '...'. An association with that name does not exist.",
             "reason": "invalid"
           }
         ],
         "message": "Invalid value for field 'name': '...'. An association with that name does not exist."
       }
     }
     ```
   - Updated `mockcompute/service.go` `RewriteError` to populate the `errors` detail slice for 400 errors, ensuring exact alignment between MockGCP and real GCP.

5. **Mock Operations Lifecycle**:
   - `addAssociation` returns an operation of type `attachOrganizationSecurityPolicy`.
   - `removeAssociation` returns an operation of type `detachOrganizationSecurityPolicy`.
   - Both retain `targetId` and `targetLink` pointing to the associated security policy during operation execution and completion.
