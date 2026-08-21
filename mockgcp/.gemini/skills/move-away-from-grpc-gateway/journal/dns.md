## 2026-06-14 - DNS migration blocked

The migration of `mockdns` away from `grpc-gateway` to `httptogrpc` is **blocked**.

### Reason
Cloud DNS does not have an official gRPC protobuf package or Go client library under `cloud.google.com/go/` or standard GCP packages.
Therefore, it cannot be registered via `httptogrpc`, which relies on official gRPC clients and reflection.
As a result, we must keep the locally generated protos in `mockgcp/generated/google/cloud/dns/v1/` and continue using `grpc-gateway` for Mock DNS.
