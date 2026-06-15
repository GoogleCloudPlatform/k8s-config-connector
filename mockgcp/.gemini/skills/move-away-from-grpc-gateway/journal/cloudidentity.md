# Blocked Migration: Cloud Identity

## Status
Blocked

## Details
We are unable to migrate `cloudidentity` away from `grpc-gateway` to `httptogrpc`.

The `httptogrpc` multiplexer requires the service to have official protobuf/gRPC client libraries published under `cloud.google.com/go` (e.g., under a package path like `cloud.google.com/go/cloudidentity/apiv1/cloudidentitypb`).

Currently, the only official Go client library published by Google for Cloud Identity is `google.golang.org/api/cloudidentity/v1` (and `v1beta1`), which is a REST-only library generated from the Discovery Document. It does not provide the standard gRPC service/client interfaces or proto message definitions.

Without an official gRPC Go client module:
1. We cannot import standard Go proto types for `cloudidentity`.
2. `httptogrpc` cannot bind the incoming JSON/REST requests to gRPC types via reflection.

Consequently, `cloudidentity` must continue to use `grpc-gateway` with our locally-generated `mockgcp.cloud.cloudidentity.v1beta1` proto definitions until an official gRPC client module is available.
