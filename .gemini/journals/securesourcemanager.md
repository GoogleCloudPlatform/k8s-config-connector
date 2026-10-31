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

### [2026-10-29] Greenfield Step 1b: Triage judgement queue and finalize types for SecureSourceManagerHook
- **Context**: Triaging judgement queue and finalizing direct KRM types and identity for `SecureSourceManagerHook`.
- **Problem**: `SecureSourceManagerHook` is a child resource under `SecureSourceManagerRepository`. The parent reference `SecureSourceManagerRepositoryRef` in `v1beta1` was a legacy reference implementing only `ExternalNormalizer`, lacking `refs.Ref` methods (`GetGVK`, `GetNamespacedName`, `GetExternal`, `SetExternal`, `ValidateExternal`, `Normalize`) and `RepositoryIdentity.FromExternal` required by `generate-identity`.
- **Solution**: Implemented `refs.Ref` and `refs.ExternalRef` on `SecureSourceManagerRepositoryRef`, added `FromExternal` to `RepositoryIdentity`, replaced `projectRef` and `location` with `parentRef *securesourcemanagerv1beta1.SecureSourceManagerRepositoryRef` per the single-parent rule, and regenerated identity, controller, fuzzer, and CRD manifests.
- **Impact**: Child resources under `SecureSourceManagerRepository` can now resolve parent references canonically using `parentRef.Normalize`.

