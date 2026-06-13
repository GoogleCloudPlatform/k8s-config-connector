### [2026-06-05] Reachability Analysis and Recursive Structure (ListValue) in AIPlatform Direct Types
- **Context**: Implementing Greenfield types for `VertexAIPersistentResource` in `aiplatform.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: When generating types for `VertexAIPersistentResource:PersistentResource`, the generator pruned types needed by the existing `AIPlatformModel` resource because they weren't specified in the `--resource` list. Furthermore, `google.protobuf.ListValue` is used as a recursive type inside `google.protobuf.Value` which caused a compiler/generator error when omitted or commented out.
- **Solution**:
  1. Specified both `AIPlatformModel:Model` and `VertexAIPersistentResource:PersistentResource` in the same `generate-types` command under `generate.sh` to prevent pruning.
  2. Manually injected the definition of `ListValue` and the import of `apiextensionsv1` in `types.generated.go` automatically in the `generate.sh` script via a python inline script, allowing both AST validation and `controller-gen` (CRD generator) to succeed without compilation errors.
- **Impact**: Ensures that adding any new direct types under the `aiplatform` service package does not break existing `AIPlatformModel` compilation due to type pruning or missing protobuf helper types.
