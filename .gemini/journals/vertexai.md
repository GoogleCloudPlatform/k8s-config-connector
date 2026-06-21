# VertexAI Journal

### 2026-06-05 VertexAIOnlineEvaluator Types and Identity Implementation
- **Context**: Implementing Greenfield VertexAIOnlineEvaluator under `apis/vertexai/v1alpha1`.
- **Problem**: 
  1. Pinned Google APIs SHA (`1765b559c4`) does not have the `OnlineEvaluator` proto messages or its updated `evaluation_service.proto` definitions, causing `generate-types` to fail.
  2. Sequential `generate-types` runs for `v1beta1` and `v1` services in `generate.sh` completely overwrite the shared `types.generated.go` file, deleting each other's types.
  3. `gcpurls.Template` converts placeholders like `{online_evaluator}` to lowercase (`online_evaluator`), which causes a runtime panic if the Go identity struct field is named `OnlineEvaluator` (as it looks for `onlineevaluator` after lowercase conversion).
- **Solution**:
  1. Added updated/missing proto files under `mockgcp/apis/google/cloud/aiplatform/v1beta1/`, and modified `generate-proto.sh` to copy `mockgcp/apis/google/*` overrides into `googleapis` cache directory before compiling. This avoids shadowing issues in `protoc`.
  2. Configured `generate.sh` to output to separate `types_v1beta1.generated.go` and `types_v1.generated.go` files (both matching the `generated.go` suffix, which `findTypeDeclaration` and `prunetypes` skip/support safely).
  3. Swapped `datalabelingjob_types.go` from using the local `Status` struct to `common.Status` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/common` to avoid compiler and duplicate deepcopy type generation conflicts.
  4. Implemented `VertexAIOnlineEvaluatorIdentity` with field name `Online_evaluator` (with underscore) so that lowercase conversion matches the placeholder perfectly.
- **Impact**: Enables co-existence of types from multiple services and API versions of the same service inside the same package without conflicts.
