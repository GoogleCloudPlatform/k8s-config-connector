### [2026-10-21] REST-Only Resource Lacking Public Protobuf Definitions (IAMWorkloadIdentityPoolManagedIdentity)
- **Context**: Greenfield Step 1 type generation for `IAMWorkloadIdentityPoolManagedIdentity` (`iam.cnrm.cloud.google.com/v1alpha1`, Relates to #14199, PR #14203).
- **Problem**:
  - Mechanical scaffolding via `controllerbuilder generate-types --prepopulate-resource` and `generate-identity` fails with `proto: not found`.
  - Resource exists in Google Cloud IAM v1 REST API:
    - Resource URL format in Cloud Asset Inventory: `//iam.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/workloadIdentityPools/{{WORKLOAD_POOL_ID}}/namespaces/{{NAMESPACE_ID}}/managedIdentities/{{MANAGED_IDENTITY_ID}}`
    - Discovery API type: `iam.googleapis.com/WorkloadIdentityPoolManagedIdentity`
    - Go client library representation exists in `google.golang.org/api/iam/v1` (`WorkloadIdentityPoolManagedIdentity` struct).
    - REST endpoint: `projects.locations.workloadIdentityPools.namespaces.managedIdentities` (`https://cloud.google.com/iam/docs/reference/rest/v1/projects.locations.workloadIdentityPools.namespaces.managedIdentities`).
  - While `google/iam/v1beta/workload_identity_pool.proto` exists in `googleapis/googleapis`, it only defines `WorkloadIdentityPool` and `WorkloadIdentityPoolProvider`; it lacks messages for `WorkloadIdentityPoolNamespace` and `WorkloadIdentityPoolManagedIdentity`. Google only publishes gRPC protobufs in `googleapis/googleapis`; REST-only endpoints exposed through Google API Discovery Service lack public protobuf definitions.
- **Solution**:
  - Documented findings in journal and updated `prepopulate/SKILL.md` to guide handling when resources lack public protobuf definitions.
  - Highlighted architectural options for maintainers:
    1. Providing a synthetic proto definition under `mockgcp/apis/` to enable `controllerbuilder` code generation.
    2. Hand-crafting KRM types, identity, references, and CRD schemas based on `google.golang.org/api/iam/v1` REST schema.
    3. Enhancing `controllerbuilder` to generate KRM types from Google Discovery JSON schemas.
- **Impact**: Prevents repeated discovery efforts and clarifies why `IAMWorkloadIdentityPoolManagedIdentity` (and parent `IAMWorkloadIdentityPoolNamespace`) cannot be generated through `controllerbuilder` from upstream `googleapis/googleapis`.
