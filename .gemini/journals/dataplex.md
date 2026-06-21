### [2026-05-26] Implement DataplexDataScan Types
- **Context**: Implementing KRM types for DataplexDataScan (google.cloud.dataplex.v1)
- **Problem**: DataplexDataScan requires manual scaffolding of inner struct fields inside `datascan_types.go` because the `prunetypes` generator removes nested types in `types.generated.go` if they are unreachable. Additionally, `k8s.io/apimachinery/pkg/runtime/schema` must be correctly imported in `datascan_reference.go` instead of `k8s.io/apimachinery/pkg/schema`.
- **Solution**: Scaffolded fields `DataQualitySpec`, `DataProfileSpec`, `DataDiscoverySpec` and their related `ResultObservedState` fields manually in `datascan_types.go`. Fixed import path in `datascan_reference.go`.
- **Impact**: Agents working on direct controllers for `dataplex` or any other nested specification APIs need to remember to correctly instantiate nested structure types in `datascan_types.go` and ensure proper imports for `schema.GroupVersionKind`.
