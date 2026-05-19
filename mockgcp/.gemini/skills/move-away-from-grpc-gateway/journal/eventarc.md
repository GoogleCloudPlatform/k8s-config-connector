# Eventarc Migration to httptogrpc

- Moved Eventarc away from `grpc-gateway` to `httptogrpc`.
- Switched to official `cloud.google.com/go/eventarc/apiv1/eventarcpb` protos.
- Updated `NewHTTPMux` to use `httptogrpc.NewGRPCMux`.
- No major type mismatches encountered.
- `proto.CloneOf` was already in use and compatible with the new proto types.
