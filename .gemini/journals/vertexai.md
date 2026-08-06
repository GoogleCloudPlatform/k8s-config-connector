### 2026-07-02 VertexAIDeploymentResourcePool Direct Controller Implementation
- **Context**: Greenfield implementation of VertexAIDeploymentResourcePool direct controller and its MockGCP service.
- **Problem**: Long-running operations for Create/Update and deletion generic metadata expectations had to be carefully aligned, and mock support had to be explicitly enabled in `config/tests/samples/create/harness.go` to prevent the E2E tests from skipping the GVK.
- **Solution**: Implemented the direct controller under `pkg/controller/direct/vertexai/deploymentresourcepool` using the `apiv1beta1` REST client. Registered the mock implementation for `DeploymentResourcePoolService` in MockGCP and updated `config/tests/samples/create/harness.go` to support the GVK.
- **Impact**: Enables end-to-end testing of `VertexAIDeploymentResourcePool` against both MockGCP and real GCP without hitting skips.

### 2026-08-06 VertexAITensorboard Direct Controller Implementation
- **Context**: Greenfield implementation of VertexAITensorboard direct controller, fuzzer, and round-trip verification under the `vertexai` package (apis/vertexai/v1alpha1).
- **Problem**: Tensorboard has a server-generated ID under GCP and MockGCP (unlike user-specified names). This created a mismatch in `GetIdentity` between the fallback `specIdentity` (using `obj.GetName()`) and the actual `statusIdentity` (using the server-generated numeric ID), causing `cannot change VertexAITensorboard identity` errors and 404s when fetching the populated resource after creation.
- **Solution**: 
  1. Updated `GetIdentity()` in `vertexaitensorboard_identity.go` to return `statusIdentity` if the `spec.resourceID` is empty (indicating a server-generated ID), bypassing the rigid uniqueness/immutability check against the Kubernetes resource name.
  2. Modified the direct controller `Create()` method to query the actual newly created `pb.Tensorboard` using `created.Name` (which contains the server-generated ID), and set `status.externalRef` to the canonical GCP URI.
- **Impact**: Enables stable end-to-end reconciliation and lifecycle testing of VertexAITensorboard with server-generated IDs.
