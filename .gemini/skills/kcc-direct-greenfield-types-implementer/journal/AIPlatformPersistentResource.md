# AIPlatformPersistentResource Types implementation Journal

## Observations & Findings

1. **Mapping Snake Case in gcpurls.Template**:
   - The GCP API path is `/projects/{project}/locations/{location}/persistentResources/{persistent_resource}`.
   - Per step 3 of the `kcc-identity-reference` skill, any snake_case path variables must be mapped to camelCase `persistentResource` in the `gcpurls.Template` and Go struct field `PersistentResource string`.

2. **Reaching Commented Out Types**:
   - When running `generate.sh` initially, many nested type definitions in `types.generated.go` (like `ResourcePool`, `ResourceRuntime`, etc.) were pruned and commented out as unreachable because they were not yet referenced.
   - After explicitly defining `AIPlatformPersistentResourceSpec` and `AIPlatformPersistentResourceObservedState` and referencing these nested structures, `generate.sh` successfully detected them, uncommented them, and generated the deepcopy methods and mapper methods automatically.

3. **External Git Diff Config Error**:
   - The workspace environment had an empty `diff.external` configuration value which caused `git diff` commands to fail.
   - We resolved this locally by overriding the GIT_CONFIG_COUNT (e.g. `GIT_CONFIG_COUNT=7`) or by running `git diff --no-ext-diff`.
