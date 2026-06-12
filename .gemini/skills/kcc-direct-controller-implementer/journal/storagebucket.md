# StorageBucket Direct Controller Journal

## Architectural Choices & Design

### GCS REST Client Adaptation
Google Cloud Storage (GCS) is primarily a REST-based API. While `google.golang.org/genproto/googleapis/storage/v1` defines protobuf schemas for GCS, using a raw gRPC connection in the direct controller leads to compatibility issues with mockgcp and standard environments due to protobuf field number mismatches in Google's split `service.proto`.

To deliver a robust direct controller that is completely compatible with both real GCP and mockgcp's REST intercept layer, we:
1. Kept `storagebucket_mapper.go` fully generated and typed with `pb.Bucket` (`google.golang.org/genproto/googleapis/storage/v1`).
2. Configured the controller `storagebucket_controller.go` to instantiate GCS's official HTTP REST client library (`google.golang.org/api/storage/v1`).
3. Implemented clean `pbToRESTBucket` and `restToPBBucket` serialization mappers using `protojson` to bridge the protobuf schema and the REST GCS client structures.

This design completely bypasses gRPC service mismatch limits, uses GCS's standard REST interface, and allows our controller to cleanly interoperate with mockgcp's excellent REST mocking.

### mockgcp Service Descriptor Mapping
In `mockgcp/mockstorage/service.go`, GCS's gRPC services are split into `BucketsServer`, `ObjectsServer`, etc. under the `mockgcp.storage.v1` package namespace. To ensure any gRPC GCS clients are fully compatible with mockgcp's gRPC server, we registered mockgcp's `buckets` handler under the canonical GCS gRPC service name `google.storage.v1.Storage`.
