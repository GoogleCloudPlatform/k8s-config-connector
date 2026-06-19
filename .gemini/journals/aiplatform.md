# AIPlatform (Vertex AI) Journal

### 2026-06-05 aiplatform: ListValue recursion & missing imports in types generation
- **Context**: Implementing Greenfield v1alpha1 `VertexAIRagCorpus` direct resource types.
- **Problem**: 
  1. The code generator fails during the `prunetypes` phase because `types.generated.go` uses `apiextensionsv1.JSON` but does not import `apiextensionsv1` ("use of unimported package 'apiextensionsv1'"). Since pruning compiles the package before formatting, `goimports` cannot run to fix it.
  2. The `google.protobuf.ListValue` type was manually commented out in `model_types.go` because of historical `controller-gen` recursion limits. This caused the type generator to skip generating `ListValue` entirely, resulting in "unknown type ListValue" errors because fields (like `ExplanationParameters.output_indices`) referenced it.
- **Solution**:
  1. Enhanced `dev/tools/controllerbuilder/pkg/codegen/typegenerator.go` to explicitly register and inject `apiextensionsv1` import when `google.protobuf.Struct` is used by a field.
  2. Fully uncommented `ListValue` and `ListTypeValue` in `apis/aiplatform/v1alpha1/model_types.go`. Modern `controller-gen` handles the recursive structure perfectly and outputs robust OpenAPI schemas using preservation of unknown fields.
- **Impact**: Code generation compiles and builds cleanly without manual post-generation corrections, while restoring correct schemas for all recursive list types.
