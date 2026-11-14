### [2026-11-14] Prepopulating ResourceManagerTagHold
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for ResourceManagerTagHold (resourcemanager.cnrm.cloud.google.com/v1alpha1).
- **Problem**: 
  1. `google.cloud.resourcemanager.v3` has `servicePrefix: CloudResourceManager` in `service-prefixes.yaml`. Without an entry in `knownKindNames`, `TestKindNamesFollowServicePrefixes` expected `CloudResourceManagerTagHold`.
  2. Enabling `--multiversion` on `apis/resourcemanager/generate.sh` caused `generate-mapper` to search for `*_v1beta1_*` mapping functions for `Project` and `Folder`.
- **Solution**: 
  1. Added `{ServicePrefix: "CloudResourceManager", Message: "TagHold", Want: "ResourceManagerTagHold"}` to `knownKindNames` in `tests/apichecks/kind_name_violations_test.go`.
  2. Added `_v1beta1_` forwarding functions in `pkg/controller/direct/resourcemanager/mapper.go` delegating to `FolderSpec_FromProto`, `FolderSpec_ToProto`, `ProjectSpec_FromProto`, etc.
- **Impact**: Allows `ResourceManagerTagHold` to use the canonical Kind name and cleanly integrates `v1alpha1` alongside `v1beta1` in `apis/resourcemanager/generate.sh`.
