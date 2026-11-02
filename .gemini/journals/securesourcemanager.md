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

### [2026-11-02] Greenfield Step 2: Direct controller, fuzzer, and E2E fixtures for SecureSourceManagerBranchRule
- **Context**: Recorded and verified E2E fixtures against real GCP and confirmed direct controller and round-trip fuzzer for `SecureSourceManagerBranchRule`.
- **Learnings & Observations**:
  1. Real GCP `Repository` responses return a `scanConfig` field which is not yet defined in upstream `googleapis.pb`; added `removeKeysFromMap(obj, []string{"scanConfig"})` in `tests/e2e/normalize.go` for `/repositories` to pass V7 protobuf schema validation.
  2. The SecureSourceManager API does not support `annotations` in `updateMask` for BranchRules and drops them on creation; excluded `annotations` from test fixtures (tracked in `alpha-missingfields.txt`).
  3. `includePattern` cannot be changed from a specific branch name (e.g. `main`) to a wildcard on an existing rule, and SecureSourceManager requires `minimumReviewsCount >= 2` (greater than `minimumApprovalsCount`). Valid mutable fields tested in update include `disabled`, `requireCommentsResolved`, `allowStaleReviews`, and `requireLinearHistory`.
  4. Dependent `SecureSourceManagerInstance` creation and deletion take ~30m and ~25m respectively on real GCP; running with `SUBTEST_TIMEOUT_E2E=120m` is required to avoid premature test timeouts.
- **Impact**: E2E fixtures for minimal and maximal pass all 7/7 automated validation checks (V1-V7) against real GCP.

