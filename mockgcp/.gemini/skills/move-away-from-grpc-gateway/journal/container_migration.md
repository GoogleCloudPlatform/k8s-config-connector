## container migration
- Moved mockgcp container (GKE) to httptogrpc.
- Deleted generated container proto files in `mockgcp/generated/mockgcp/container/v1beta1/`.
- Modified `mockgcp/Makefile` to remove container protos from code generation.
- Updated `mockgcp/mockcontainer/service.go`, `cluster.go`, `lro.go`, and `nodepool.go` imports to use the official package `cloud.google.com/go/container/apiv1/containerpb`.
- Refactored `NewHTTPMux` to use `httptogrpc.NewGRPCMux` and `pb.NewClusterManagerClient`.
- Implemented `/v1beta1/` to `/v1/` path rewriting for incoming HTTP requests to map GKE v1beta1 endpoints used by Terraform to the official v1-based client annotations.
- Removed legacy, deprecated assignments of `PrivateCluster`, `ClusterTelemetry`, and `ProtectConfig` from the mock GKE cluster resource representation as they are not defined in GKE's v1 protobuf schema.
- Verified that all mockcontainer components successfully compile.
