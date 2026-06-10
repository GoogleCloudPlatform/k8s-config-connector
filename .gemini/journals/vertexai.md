### [2026-06-04] VertexAI Resource Initialization and Scaffolding
- **Context**: Implementing the initial direct types and IdentityV2 for `VertexAIModelMonitor` (GCP `google.cloud.aiplatform.v1beta1.ModelMonitor`).
- **Problem**: 
  1. The `apis/vertexai/v1alpha1` directory mixes resources from two different proto services (`google.cloud.aiplatform.v1beta1` and `google.cloud.aiplatform.v1`). Sequentially running `generate-types` for both services completely overwrote `types.generated.go`, wiping out first-pass types and leaving package types uncompilable.
  2. Both services reference `google.protobuf.Value` and `google.protobuf.ListValue`. Since the pruning tool detects references to these types across different files in the same `v1alpha1` package, it preserved them in both `types.generated.go` and `types_v1beta1.generated.go`, causing duplicate type redeclaration compile errors.
- **Solution**:
  1. Updated `apis/vertexai/v1alpha1/generate.sh` to rename the first-pass generated `types.generated.go` to `types_v1beta1.generated.go` before executing the second generation step.
  2. Implemented a post-processing Python step in `generate.sh` to automatically comment out duplicate `Value` and `ListValue` declarations in `types.generated.go`.
  3. Updated `datalabelingjob_types.go` imports and references to map `google.rpc.Status` to `common.Status` as expected by modern controllerbuilder versions.
- **Impact**: Resolves package-wide generation conflicts and duplicate type compilation issues, enabling seamless execution of `generate.sh` and validation pipelines for mixed-version API directories.
