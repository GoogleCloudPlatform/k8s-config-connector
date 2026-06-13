### [2026-05-19] Empty Proto Messages Mapping to Empty Structs
- **Context**: Scaffolding `AutoMLDataset` direct types.
- **Problem**: The proto `Dataset` contains fields like `textExtractionDatasetMetadata` that map to empty proto messages. The `generate-types` tooling generates empty Go structs for these. Kubebuilder CRD generation will fail for empty structs with the error `missing properties, additionalProperties, or x-kubernetes-preserve-unknown-fields`.
- **Solution**: Added `// +kubebuilder:pruning:PreserveUnknownFields` and `// +kubebuilder:validation:Schemaless` comments directly above the empty struct definitions in the Go types file (`apis/automl/v1alpha1/dataset_types.go`).
- **Impact**: Ensures that `dev/tasks/generate-crds` correctly generates the CRD with `x-kubernetes-preserve-unknown-fields: true` for these empty objects, and the `TestCRDObjectTypes` test passes without errors.

### [2026-05-19] Unreachable Types
- **Context**: Scaffolding `AutoMLDataset` direct types.
- **Problem**: `generate-types` tooling comments out many subtypes in `types.generated.go` as "unreachable" because `AutoMLDatasetSpec` initially doesn't reference them.
- **Solution**: Manually copy the needed metadata fields (like `translationDatasetMetadata`) from the commented out proto specs into `AutoMLDatasetSpec` inside `dataset_types.go`, then re-run `generate.sh`.
- **Impact**: Restores visibility of necessary subtypes and ensures the schema covers all the proto fields.

### [2026-05-19] Preserving Manual Hacks in types.generated.go
- **Context**: Scaffolding types for VertexAIFeatureOnlineStore in the existing aiplatform/v1alpha1 package.
- **Problem**: generate-types rewrites types.generated.go entirely based on the --resource flags. If you omit existing resources, it deletes their types. If you include them, it overwrites manual hacks (like commenting out recursive ListValue fields or manually fixing StructValue imports) that were necessary for CRD generation to pass.
- **Solution**: You must pass all existing --resource flags to generate-types to preserve their generation, and you MUST re-apply any manual hacks using sed or go AST manipulation within generate.sh before running generate-mapper or generate-crds.
- **Impact**: Failing to re-apply manual hacks will cause generate-crds (via controller-gen) to fail with recursive structure errors. Failing to include existing resources will cause missing type errors downstream in mapper.generated.go.

### [2026-05-19] Generate-Mapper Flag Pitfalls
- **Context**: Running generate-mapper via generate.sh for VertexAIFeatureOnlineStore.
- **Problem**: Running generate-mapper without --api-dir processes all packages under apis/ that match the --service flag (e.g., both aiplatform/v1alpha1 and vertexai/v1alpha1 for google.cloud.aiplatform.v1), causing type pollution in mapper.generated.go. Passing --api-dir but not --api-go-package-path causes the generator to miscalculate the import alias (resulting in apis instead of krm or krmv1alpha1), breaking the mapper compilation.
- **Solution**: When generating mappers for a specific package that shares a --service with other packages, you must pass BOTH --api-dir and --api-go-package-path.
- **Impact**: Prevents broken imports and polluted mapper files when multiple API packages share the same underlying proto service.
