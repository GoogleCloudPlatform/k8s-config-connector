# Migration Blocked: mockapigee

## Status
**Blocked**

## Details
The migration of the `mockapigee` service away from `grpc-gateway` to `httptogrpc` is currently blocked because there is no official gRPC/protobuf-based Go client library available under `cloud.google.com/go/apigee`.

The standard library for Apigee is REST-based and located in the `google-api-go-client` repository under `google.golang.org/api/apigee/v1`. Since `httptogrpc` requires reflection over official gRPC protobuf types and client definitions, we cannot migrate this mock service until Google provides an official gRPC Go client package (typically under `cloud.google.com/go/apigee/apiv1/apigeepb`).

As per **Step 0** of the migration skill instructions, this service will continue to use the locally generated `grpc-gateway` protobuf bindings and should not be migrated at this time.
