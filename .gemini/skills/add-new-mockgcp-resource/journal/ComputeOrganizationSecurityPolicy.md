# Journal: ComputeOrganizationSecurityPolicy MockGCP Implementation

This journal captures the details and findings of implementing the `ComputeOrganizationSecurityPolicy` mock in MockGCP and matching real GCP behavior.

## Observations

1. **Server-Generated ID (`policy_id`)**:
   `ComputeOrganizationSecurityPolicy` has a server-generated numeric policy ID. In `scripts/resource-autogen/generated/servicemappings/compute.yaml`, `serverGeneratedIDField: policy_id` was previously missing. Without this, KCC assumed the resource ID was user-specified, causing the controller to attempt reading a nonexistent policy by resourceID on creation instead of executing the `POST` request. Setting `serverGeneratedIDField: policy_id` and regenerating embedded service mappings resolves this.

2. **Policy Type: `FIREWALL` vs `CLOUD_ARMOR`**:
   Real GCP rejects `type: FIREWALL` for organization security policies on organizations migrated or modern organizations with `Accessing firewall policies through organizationSecurityPolicies API is not supported in this organization` because hierarchical firewall policies must now use Network Firewall Policies (`ComputeNetworkFirewallPolicy`). `type: CLOUD_ARMOR` is supported and functions correctly on modern organizations.

3. **Request & Response Field Mapping (`displayName` vs `shortName`)**:
   For `type: CLOUD_ARMOR`, the compute API requires the policy name in the request body field `shortName`. The Terraform provider previously only populated `displayName`, which caused GCP to return `Invalid value for field 'resource.shortName': ''. SecurityPolicy must have a user-assigned name.` Updating `resource_compute_organization_security_policy.go` to send `shortName` on create and read `shortName` when `displayName` is absent aligns with GCP's API contract.

4. **`TEST_FOLDER_ID` Parent Type Check in `hack/record-gcp` and `dev/tasks/run-e2e`**:
   `hack/record-gcp` and `dev/tasks/run-e2e` used `gcloud projects describe ${GCP_PROJECT_ID} --format='value(parent.id)'` unconditionally, which set `TEST_FOLDER_ID` to the organization ID when the project's parent was an organization. As a result, the real GCP log normalization replaced the organization ID with `${folderID}` rather than `${organizationID}`. Checking that `parent.type` is `folder` before setting `TEST_FOLDER_ID` ensures correct replacement.

5. **MockGCP Implementation**:
   Implemented `OrganizationSecurityPoliciesV1` in `mockgcp/mockcompute/organizationsecuritypoliciesv1.go` providing `Insert`, `Get`, `Patch`, `Delete`, and `List`. The LRO handling in `mockgcp/mockcompute/operations.go` was updated so `createOrganizationSecurityPolicy` omits `targetId` and `targetLink` while the operation is in `RUNNING` status and includes them once `DONE`, matching authentic GCP behavior.

6. **CAIS Dynamic ID Normalization**:
   Updated `pkg/cais/caistesting/testing.go` to normalize server-generated `locations/global/securityPolicies/<id>` URLs to `unknown` in `NormalizeDynamicIDs`, ensuring parity between MockGCP and real GCP golden identity files.
