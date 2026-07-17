### [2026-05-18] Migrating mockgcp securesourcemanager to httptogrpc
- **Context**: Migrating `securesourcemanager` away from grpc-gateway.
- **Problem**: Need to ensure all components such as error rewriting and operations paths are migrated correctly.
- **Solution**: `RewriteError` for 404 errors can be safely removed. Changed `s.operations.RegisterOperationsPath` to `mux.AddOperationsPath` when registering the grpcMux.
- **Impact**: Other agents migrating similar services to `httptogrpc` should know to remove `RewriteError` and properly register operations paths with `mux.AddOperationsPath`.

### [2026-06-22] Implement direct controller and E2E fixtures for SecureSourceManagerInstance
- **Context**: Standardizing and fully implementing the direct controller for `SecureSourceManagerInstance`.
- **Observation**: SecureSourceManagerInstance does not have an `UpdateInstance` method in the GCP Go Client, making the resource functionally immutable.
- **Solution**: Updated `Create` and `Update` methods using standard `compareSecureSourceManagerInstance` and `updateStatus` helper patterns. Implemented explicit immutability check in the `Update` method returning a clean error containing the exact field diffs if any updates are attempted.
- **Impact**: Fully aligned the controller with direct controller requirements and ensured schema correctness in validation and fuzz-testing.
