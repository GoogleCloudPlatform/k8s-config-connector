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

### [2026-10-29] Triage judgement queue and finalize KRM types for SecureSourceManagerBranchRule
- **Context**: Triage judgement queue and finalize direct KRM types, identity, and references for `SecureSourceManagerBranchRule` (Step 1b).
- **Problem**: `SecureSourceManagerRepositoryRef` in `v1beta1` only implemented deprecated `refsv1beta1.ExternalNormalizer` and was missing `refsv1beta1.Ref` methods, which are required for `generate-identity` to normalize `spec.parentRef`.
- **Solution**: Implemented `refsv1beta1.Ref` on `SecureSourceManagerRepositoryRef` and `FromExternal` on `RepositoryIdentity`. Modelled `spec.parentRef` on `SecureSourceManagerBranchRuleSpec` pointing to `SecureSourceManagerRepositoryRef`, moved `etag` from `spec` to `status.observedState`, and added `generate-controller` and `generate-fuzzer` to `apis/securesourcemanager/generate.sh`.
- **Impact**: Enables clean single-parent identity resolution for `SecureSourceManagerBranchRule` and ensures controller and fuzzer stay in lockstep.

### [2026-11-01] Greenfield Step 2: Implement direct controller, E2E fixtures, and fuzzer for SecureSourceManagerHook
- **Context**: Implementing direct controller, golden fixtures, and fuzzer for `SecureSourceManagerHook` (`securesourcemanager.cnrm.cloud.google.com/v1alpha1`).
- **Problem**:
  1. `SecureSourceManagerInstance` dependencies take 35-45 minutes to provision and may hit regional capacity issues ("Instance creation is temporarily unavailable. Try a different region, or try again later.") in regions like `us-central1`.
  2. Real GCP returns `scanConfig` on `Repository` responses which was not in standard googleapis protos, causing V7 validation failures.
  3. Real GCP always redacts `sensitive_query_string` to `"[REDACTED]"` on `GET`, causing `CompareGreenfieldSpec` to detect a spurious diff during re-reconciliation unless normalized.
- **Solution**:
  1. Tested and used available regions (`europe-west4` and `asia-east1`) for the test fixtures.
  2. Added scoped normalization in `tests/e2e/normalize.go` guarded by `strings.Contains(t.Name(), "securesourcemanager")` to remove `.response.scanConfig` and `.scanConfig` from HTTP responses.
  3. Added a `normalize` callback in `HookAdapter.Update` to clear `SensitiveQueryString` when `observedGeneration == generation` and actual is `"[REDACTED]"`.
  4. Opted out of controller overwriting by removing `// +generated:controller` marker from `pkg/controller/direct/securesourcemanager/securesourcemanagerhook_controller.go` to preserve reference normalization before identity resolution and export parent/resourceID settings.
- **Impact**: Ensures clean re-reconciliation and passing V1-V7 HTTP lifecycle validation for `SecureSourceManagerHook`.

