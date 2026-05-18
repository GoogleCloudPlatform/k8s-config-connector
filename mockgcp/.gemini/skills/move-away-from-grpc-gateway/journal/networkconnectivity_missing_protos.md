# Network Connectivity Missing Protos

When moving `networkconnectivity` away from `grpc-gateway`, it was discovered that `RegionalEndpoint` is present in the Discovery JSON but missing from the official `googleapis` gRPC protos and the Go client library.

As a result, the `RegionalEndpoint` mock implementation and its tests were removed, as they are not currently used by any KCC resource and cannot be easily moved to `httptogrpc` without gRPC type definitions.

Other resources (`InternalRange` and `ServiceConnectionPolicy`) were successfully moved to the official client.
