### [2026-06-05] Resolved recursive type imports and struct builder imports for VertexAITuningJob and AIPlatformModel
- **Context**: Implementing VertexAITuningJob and regenerating types for AIPlatformModel (under Service `google.cloud.aiplatform.v1`).
- **Problem**: 
  1. The controllerbuilder's code generator mapping for `google.protobuf.Struct` results in `apiextensionsv1.JSON` fields but fails to inject the corresponding package import `"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"`, leading to compilation errors.
  2. The `ListValue` struct was commented out in `model_types.go` because of previous recursive definition issues, causing `findTypeDeclarationWithProtoTag` to skip generating `ListValue` automatically. This led to `unknown type ListValue` errors.
- **Solution**: 
  1. Patched `dev/tools/controllerbuilder/pkg/codegen/typegenerator.go` to automatically add the `apiextensionsv1` import whenever `google.protobuf.Struct` is mapped.
  2. Deleted the commented-out `ListValue` and `ListTypeValue` declarations in `model_types.go` to let the code generator natively and cleanly output them in `types.generated.go`.
- **Impact**: Cleaned up code-gen for all other resources in the same group and prevented manual import/stubbing workarounds in type files.
