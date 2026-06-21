### 2026-06-06 Resolving Type Collisions and Recursion in AI Platform

- **Context**: Implementing direct types, KRM, and IdentityV2 for `VertexAIPipelineJob` (GCP `PipelineJob` in `google.cloud.aiplatform.v1`) under the `aiplatform.cnrm.cloud.google.com/v1alpha1` API group.
- **Problem**: 
  1. Running `generate-types` on `VertexAIPipelineJob` pruned existing `AIPlatformModel` types in `types.generated.go`.
  2. A naming conflict occurred because both `google.protobuf.Value` and `google.cloud.aiplatform.v1.Value` map to a Go struct named `Value` in the same package scope, causing `Value redeclared in this block` compile errors.
  3. `controller-gen` failed due to `apiextensionsv1` not being imported in `types.generated.go`.
  4. `ListValue` generation was skipped because the generator found a commented-out manual type declaration with the tag `// +kcc:proto=google.protobuf.ListValue` in `model_types.go`, leading to undefined `ListValue` compile errors.
- **Solution**:
  1. Updated `generate.sh` to include both `--resource AIPlatformModel:Model` and `--resource VertexAIPipelineJob:PipelineJob` to ensure both resources' types are generated into the shared `types.generated.go`.
  2. Mapped `"google.cloud.aiplatform.v1.Value"` to `"apiextensionsv1.JSON"` in `protoMessagesNotMappedToGoStruct` within `dev/tools/controllerbuilder/pkg/codegen/common.go`, which resolves the name conflict and represents the union of types cleanly in KRM.
  3. Enhanced `typegenerator.go` in `dev/tools/controllerbuilder` to dynamically detect when any fields map to `apiextensionsv1.*` types and automatically add the `apiextensionsv1` import to the generated file.
  4. Renamed the `// +kcc:proto=google.protobuf.ListValue` commented-out tag in `model_types.go` to `// +kcc:proto:legacy=...` to avoid conflicting with the type generator's lookup, allowing `ListValue` to generate and compile successfully.
- **Impact**: Clean, compilation-ready type and CRD generation for any resource within the `aiplatform` API group that uses recursive structs, `Value`, or `ListValue` types.
