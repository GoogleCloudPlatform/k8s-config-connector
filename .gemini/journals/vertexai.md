### 2026-07-02 VertexAIDeploymentResourcePool Direct Controller Implementation
- **Context**: Greenfield implementation of VertexAIDeploymentResourcePool direct controller and its MockGCP service.
- **Problem**: Long-running operations for Create/Update and deletion generic metadata expectations had to be carefully aligned, and mock support had to be explicitly enabled in `config/tests/samples/create/harness.go` to prevent the E2E tests from skipping the GVK.
- **Solution**: Implemented the direct controller under `pkg/controller/direct/vertexai/deploymentresourcepool` using the `apiv1beta1` REST client. Registered the mock implementation for `DeploymentResourcePoolService` in MockGCP and updated `config/tests/samples/create/harness.go` to support the GVK.
- **Impact**: Enables end-to-end testing of `VertexAIDeploymentResourcePool` against both MockGCP and real GCP without hitting skips.
