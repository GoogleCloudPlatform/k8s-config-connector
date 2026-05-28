## 2026-05-17 - container migration

- Moved mockgcp container to httptogrpc and cloud.google.com/go/container/apiv1/containerpb.
- GKE (container) has significant differences between v1beta1 and v1. Since official v1beta1 Cloud Client is not available, migrated to v1.
- Required fixing type mismatches in `mockcontainer`: removed `PrivateCluster`, `ClusterTelemetry`, and `ProtectConfig` fields which are missing in v1.
- Fixed `StackType` enum usage (it moved to top-level in v1).
- Updated `httptogrpc` to support `RewriteRequest` and add `path` and `http.request.query` metadata to the context. This was necessary because `mockcontainer` (and potentially other mocks) use this metadata to determine the API version or other request details.
- Updated `service.go` rewrite logic to map `/v1beta1/` (used by Terraform/KCC) to `/v1/` (supported by the new mock).