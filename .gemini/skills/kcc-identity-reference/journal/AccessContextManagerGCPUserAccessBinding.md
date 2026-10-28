# AccessContextManagerGCPUserAccessBinding Identity & Reference Migration Journal

## Resource Overview
- Kind: `AccessContextManagerGCPUserAccessBinding`
- Group: `accesscontextmanager.cnrm.cloud.google.com`
- Version: `v1alpha1`
- Proto Message: `google.identity.accesscontextmanager.v1.GcpUserAccessBinding`
- Canonical Template: `organizations/{organization}/gcpUserAccessBindings/{gcpUserAccessBinding}`

## Observations & Implementation Decisions

1. **Deterministic Identity Generation**:
   - Added `${CONTROLLERBUILDER} generate-identity` to `apis/accesscontextmanager/generate.sh`.
   - Generated `accesscontextmanagergcpuseraccessbinding_identity.generated.go`, `accesscontextmanagergcpuseraccessbinding_reference.generated.go`, and `accesscontextmanagergcpuseraccessbinding_identity_generated_test.go`.
   - The generator marked `identity-server-generated-id` in `judgement_queue.yaml` because the Create RPC has no client-specified ID parameter. This was resolved as `accepted` because GCP server-assigns the last path segment during creation (as documented on `status.name`).

2. **Handwritten Override for Parent Resolution**:
   - `OrganizationRef` is defined locally in `apis/accesscontextmanager/v1alpha1/accesscontextmanagergcpuseraccessbinding_types.go` with `external`, `name`, and `namespace` fields to preserve CRD schema compatibility and validation constraints without changing the existing CRD.
   - The generator attempted to pass `obj.Spec.OrganizationRef` directly to `refs.ResolveOrganization` (which expects `*refsv1beta1.OrganizationRef`).
   - Used the declaration-by-declaration override pattern: provided `getIdentityFromAccessContextManagerGCPUserAccessBindingSpec` in `apis/accesscontextmanager/v1alpha1/accesscontextmanagergcpuseraccessbinding_identity.go`.
   - Normalized `spec.organizationRef.external` by ensuring an `"organizations/"` prefix if the user passed a plain numeric organization ID (common in Terraform configurations and test fixtures).
   - Also supported fallback to the `cnrm.cloud.google.com/organization-id` annotation.

3. **Server-Generated ID & Status Handling**:
   - `AccessContextManagerGCPUserAccessBindingIdentity` implements `identity.ServerGeneratedIdentity`.
   - When `spec.resourceID` is omitted, `GetIdentity()` reads `status.name`, parses the server-generated ID, and defaults the identity's `GcpUserAccessBinding` field. If both are specified, it cross-checks them to detect drift.

4. **Verification**:
   - `go test -v ./apis/accesscontextmanager/v1alpha1/...` passed all conformance and identity unit tests.
   - `dev/tasks/diff-crds` produced zero diffs, guaranteeing schema compatibility.
   - `apis/accesscontextmanager/generate.sh` executed cleanly end-to-end.
