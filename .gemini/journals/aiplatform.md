### 2026-06-05 VertexAIStudy Direct Types Implementation
- **Context**: Implementing direct KRM types and IdentityV2 for `VertexAIStudy` under `apis/aiplatform/v1alpha1`.
- **Problem**: 
  1. Running the `generate-types` command for `AIPlatformModel:Model` and `VertexAIStudy:Study` caused compilation errors in `types.generated.go` because the package `apiextensionsv1` was unimported. Even though we added the import to `model_types.go`, `goimports` automatically deleted the import from `model_types.go` because it was not used by any struct within `model_types.go`.
  2. `ListValue` struct was previously commented out in `model_types.go` to avoid infinite recursion/CRD generation errors. However, with the current controller builder generator, `ListValue` was expected as a struct and caused type undefined errors.
- **Solution**: 
  1. We added a dummy variable `var _ apiextensionsv1.JSON` in `model_types.go` right below the imports block to prevent `goimports` from pruning the import of `apiextensionsv1`.
  2. We defined `ListValue` properly as a Go struct containing `[]Value` (similar to how `vertexai` API group defines it).
  3. We updated `apis/aiplatform/v1alpha1/generate.sh` to run `goimports` on `apis/aiplatform/v1alpha1` right before running `dev/tasks/generate-crds` so that the generated imports are resolved before CRD validation.
- **Impact**: Keeps the `aiplatform` API group types generating and compiling cleanly, and provides a clear template for adding future resources in the same group.
