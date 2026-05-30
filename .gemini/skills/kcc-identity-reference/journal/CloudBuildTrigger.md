# CloudBuildTrigger Identity and Reference

CloudBuildTrigger has a server-generated ID. This means:
1. `getIdentityFromCloudBuildTriggerSpec` returns `(nil, nil)` if `Status.TriggerId` is not yet populated.
2. `GetIdentity` cross-checks `Project` and `Location` but cannot cross-check the `Trigger` ID until it is known.
3. `Normalize` in `CloudBuildTriggerRef` uses the same `getIdentityFrom...Spec` logic.

The resource does not have `spec.resourceID` because the ID is server-generated and cannot be specified on creation.

The files were renamed from `trigger_*` to `cloudbuildtrigger_*` to follow the Kind-based naming convention more strictly as suggested in the skill.
