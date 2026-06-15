### [2026-05-18] Migrating mockgcp securesourcemanager to httptogrpc
- **Context**: Migrating `securesourcemanager` away from grpc-gateway.
- **Problem**: Need to ensure all components such as error rewriting and operations paths are migrated correctly.
- **Solution**: `RewriteError` for 404 errors can be safely removed. Changed `s.operations.RegisterOperationsPath` to `mux.AddOperationsPath` when registering the grpcMux.
- **Impact**: Other agents migrating similar services to `httptogrpc` should know to remove `RewriteError` and properly register operations paths with `mux.AddOperationsPath`.
