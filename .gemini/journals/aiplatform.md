### [2026-06-05] Implementing VertexAISchedule Direct Types and Handling Protobuf Value Collisions
- **Context**: Implementing the initial KRM types, CRD, and IdentityV2 for `VertexAISchedule` (Issue #9248).
- **Problem**: When generating types for `VertexAISchedule:Schedule`, two name collisions occurred:
  1. `google.cloud.aiplatform.v1.Value` vs `google.protobuf.Value`. Both mapped to the Go type `Value`, causing a redeclared in this block error.
  2. `google.protobuf.ListValue` was unhandled and generated as an undefined type.
- **Solution**: 
  1. Mapped both `google.protobuf.Value` and `google.protobuf.ListValue` to `apiextensionsv1.JSON` inside the `protoMessagesNotMappedToGoStruct` map in `dev/tools/controllerbuilder/pkg/codegen/common.go`.
  2. Added logic inside `typegenerator.go` to automatically add the `apiextensionsv1` import ("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1") if either `google.protobuf.Struct`, `google.protobuf.Value`, or `google.protobuf.ListValue` are encountered.
  3. Ensured that `apis/aiplatform/v1alpha1/generate.sh` passes both `AIPlatformModel:Model` and `VertexAISchedule:Schedule` to `generate-types` so that shared and specific types for both resources are correctly generated and preserved in `types.generated.go`.
- **Impact**: Clean, type-safe generation of complex protobuf-based resources utilizing Google's well-known Protobuf types (Struct, Value, ListValue) mapped directly to Kubernetes' native `apiextensionsv1.JSON`.
