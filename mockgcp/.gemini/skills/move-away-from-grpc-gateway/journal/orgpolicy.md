### [2026-05-18] orgpolicy migration
- **Context**: Moving mockgcp orgpolicy away from grpc-gateway.
- **Problem**: Need to migrate routing logic. `mux.RewriteError` is not supported by `httptogrpc`. Protos were under `mockgcp/cloud/orgpolicy`.
- **Solution**: Removed `mux.RewriteError`. Verified that `pb.NewOrgPolicyClient(conn)` is used to register the service with `grpcMux.AddService`.
- **Impact**: Helps next agents handle mockgcp orgpolicy migrations and avoids `RewriteError` issues with `httptogrpc`.