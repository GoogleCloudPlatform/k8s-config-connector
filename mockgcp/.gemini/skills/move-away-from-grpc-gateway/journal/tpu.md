# Migration Blocked: TPU v2

## Service Information
- **Service Name:** tpu
- **Mock Service Path:** `mockgcp/mocktpu`
- **Required API Version:** v2 (as used by `TPUVirtualMachine`'s schema and mock)

## Reason for Block
The migration of the TPU mock service from `grpc-gateway` to `httptogrpc` is **blocked** because there is no official gRPC Go client library or protobuf module supporting TPU v2. 

- `cloud.google.com/go/tpu` only provides `apiv1` (as of v1.13.0) and completely lacks any `apiv2` or `apiv2alpha1` packages.
- `google.golang.org/genproto/googleapis/cloud/tpu/v2` is outdated and missing required fields (such as `AcceleratorConfig`).
- The canonical protobuf files under `google/cloud/tpu/v2` list the option `option go_package = "cloud.google.com/go/tpu/apiv2/tpupb;tpupb";`, but this package does not actually exist or cannot be imported.

According to the updated guidance in the `move-away-from-grpc-gateway` skill (updated in #8344), if no official `cloud.google.com/go` client/protobuf package exists, the migration is blocked and must not be performed.
