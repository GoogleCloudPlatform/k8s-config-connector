# Journal: Moving away from grpc-gateway

## Network Connectivity (2026-05-18)
- Moved `InternalRange` and `ServiceConnectionPolicy` to official Go client and `httptogrpc`.
- Discovered `RegionalEndpoint` is missing from official gRPC protos. Removed its mock implementation as it's not a KCC resource.
- Updated `mappergenerator.go` to remove hardcoded `networkconnectivity` exception.
- Fixed several type mismatches in `mapper.generated.go` and `internalrange_mapper.go` related to Enums and Etags.
