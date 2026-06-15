# Migration Center Journal

### 2026-06-15 Implement direct KRM types and IdentityV2 for MigrationCenterGroup
- **Context**: Implementing Greenfield types and identity logic for `MigrationCenterGroup` (GCP service `google.cloud.migrationcenter.v1` / KCC kind `MigrationCenterGroup` in `v1alpha1`). PR: https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/10345
- **Problem**: The proto-to-struct generator commented out `Group` and `GroupObservedState` fields as unreachable because they belong to the top-level resource message.
- **Solution**: Hand-mapped Spec fields (`Labels`, `DisplayName`, `Description`) and Status fields (`CreateTime`, `UpdateTime` under `MigrationCenterGroupObservedState`) from `types.generated.go` into `group_types.go`. Also mapped IdentityV2 using `projects/{project}/locations/{location}/groups/{group}` URL structure.
- **Impact**: Provides a blueprint for future Migration Center resources and demonstrates the recovery of commented-out/unreachable top-level fields for KRM.
