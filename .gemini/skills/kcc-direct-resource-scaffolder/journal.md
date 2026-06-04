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
