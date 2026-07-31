### 2026-07-29 Greenfield Types, CRD, and Identity for HypercomputeClusterCluster
- **Context**: Implementing initial types and identity for HypercomputeClusterCluster under `hypercomputecluster.cnrm.cloud.google.com/v1alpha1`. [Issue #10277]
- **Problem**: 
  1. The googleapis pinned commit SHA `731d7f2ab6` stated in the issue didn't contain the `google.cloud.hypercomputecluster.v1` service definition or proto files.
  2. The type generator skips/fails on maps with message values, such as `map[string]ComputeResource`.
- **Solution**:
  1. Pinned a newer googleapis commit (`b8486a2f44f15dc578a9dc1e17b144253079d5c1`) in `generate.sh` that includes the complete hypercomputecluster protos.
  2. Represented the map of messages as custom Go map fields `map[string]*ComputeResource`, `map[string]*NetworkResource`, and `map[string]*StorageResource` inside `HypercomputeClusterClusterSpec`. The prune mechanism in `generate-types` automatically uncommented and generated all of the reachable nested structures in `types.generated.go` seamlessly.
- **Impact**: Demonstrates a clean, reproducible way to model maps of sub-structures in KCC direct resources without requiring tedious manual definition of all transient sub-types.
