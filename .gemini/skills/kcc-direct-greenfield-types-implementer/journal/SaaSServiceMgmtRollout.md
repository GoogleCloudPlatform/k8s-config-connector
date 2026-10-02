# SaaSServiceMgmtRollout Journal

## Context
Implementing greenfield direct KRM types, identity, reference, and generate.sh configuration for `SaaSServiceMgmtRollout` under the `saasservicemgmt.cnrm.cloud.google.com` group in version `v1alpha1`.

## Actions taken

1. **Scaffold & generate.sh**:
   - Added `--resource SaaSServiceMgmtRollout:Rollout` to `apis/saasservicemgmt/generate.sh`.
   - Generated protobuf types, mapper stubs, deepcopy methods, and CRD manifests.

2. **Types Definition**:
   - Defined `SaaSServiceMgmtRolloutSpec` with required `projectRef` and immutable `location`, as well as optional fields `resourceID`, `releaseRef`, `rolloutOrchestrationStrategy`, `unitFilter`, `rolloutKindRef`, and `control`.
   - Marked `releaseRef` and `rolloutKindRef` as immutable via CEL validation rules.
   - Defined `SaaSServiceMgmtRolloutObservedState` capturing output-only proto fields: `startTime`, `endTime`, `state`, `stateMessage`, `stateTransitionTime`, `rootRollout`, `parentRollout`, `stats`, `uid`, `etag`, `createTime`, and `updateTime`.
   - Applied stability label `cnrm.cloud.google.com/stability-level=alpha`.

3. **Identity & Reference**:
   - Created `saasservicemgmtrollout_identity.go` with `SaaSServiceMgmtRolloutIdentity` using URL template `projects/{project}/locations/{location}/rollouts/{rollout}`.
   - Implemented status cross-check in `GetIdentity` to detect configuration drift.
   - Created `saasservicemgmtrollout_reference.go` implementing `refs.Ref` and delegating normalization to `refs.Normalize`.
   - Added unit tests in `saasservicemgmtrollout_identity_test.go` using `cmp.Diff`.
   - Also implemented `SaasServiceMgmtReleaseRef` in `saasservicemgmtrelease_reference.go` and `SaaSServiceMgmtRolloutKindRef` in `saasservicemgmtrolloutkind_reference.go` with `saasservicemgmtrolloutkind_identity.go` for cross-resource references.
