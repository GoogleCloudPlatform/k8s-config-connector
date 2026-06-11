# Blocked Migration: ServiceNetworking

## Status
Blocked

## Details
We are unable to migrate `servicenetworking` away from `grpc-gateway` to `httptogrpc`. 

The `httptogrpc` multiplexer requires the service to have official protobuf/gRPC client libraries published under `cloud.google.com/go` (e.g., under a package path like `cloud.google.com/go/servicenetworking/apiv1/servicenetworkingpb`).

Currently, the only official Go client library published by Google for Service Networking is `google.golang.org/api/servicenetworking/v1`, which is a REST-only library generated from the Discovery Document. It does not provide the standard gRPC service/client interfaces or proto message definitions.

Without an official gRPC Go client module:
1. We cannot import standard Go proto types for `servicenetworking`.
2. `httptogrpc` cannot bind the incoming JSON/REST requests to gRPC types via reflection.

Consequently, `servicenetworking` must continue to use `grpc-gateway` with our locally-generated `mockgcp.cloud.servicenetworking.v1` proto definitions until an official gRPC client module is available.
