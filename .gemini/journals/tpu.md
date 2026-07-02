# TPU Journal

### [2026-07-02] TPUQueuedResource Greenfield Scaffolding
- **Context**: Scaffolding direct types and IdentityV2 for `TPUQueuedResource` (tpu.cnrm.cloud.google.com/v1alpha1).
- **Problem**: 
  1. The instructions requested using `google.cloud.tpu.v1` as the proto service, but the `QueuedResource` proto message is only defined in TPU `v2` (and `v2alpha1`). TPU `v1` does not support Queued Resources.
  2. The `QueuedResource` proto message embeds `Node` and `NodeObservedState`. Since `google.cloud.tpu.v2.Node` is already registered as a top-level KRM resource (`TPUVirtualMachine`), the type generator skipped generating the standalone `Node` and `NodeObservedState` structures. This resulted in undefined-type compilation errors when compiling `types.generated.go`.
- **Solution**: 
  1. Updated the service flag in `apis/tpu/v1alpha1/generate.sh` to `--service google.cloud.tpu.v2` for `TPUQueuedResource`.
  2. Defined `Node` and `NodeObservedState` manually in `apis/tpu/v1alpha1/tpuqueuedresource_types.go` mapping to the correct speculative and observed fields of the TPU `v2` `Node` proto.
- **Impact**: Ensures that when other services or nested structures reference top-level registered KRM Kinds, we can safely define and map them locally without breaking the automatic type generation/pruning pipeline.
