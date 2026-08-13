# VertexAIDataset Journal

## Observations & Quirks

1. **Server-Generated Dataset ID**:
   - The Vertex AI Dataset service creates datasets with a server-assigned numerical ID rather than a user-specified ID (e.g., `1781291486539559728`).
   - Consequently, during identity check on updates, `statusIdentity` contains the numerical dataset ID from `status.observedState.name` whereas `specIdentity` holds the K8s name/spec resource ID (`metadata.name`).
   - We relaxed the dataset identity validation to allow matching the status ID against the spec KRM resource name ID if the status dataset ID is numeric and the spec dataset ID is non-numerical.

2. **Project ID vs. Project Number Resolution**:
   - Similarly, Vertex AI resource names typically use numerical project numbers rather than standard project ID strings.
   - We relaxed the project identity validation in `apis/vertexai/v1beta1/vertexaidataset_identity.go` to permit matching when one or both project identifiers are numerical.
