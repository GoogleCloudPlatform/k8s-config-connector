# HypercomputeCluster Journal

### [2026-06-15] HypercomputeCluster Bootstrapping and Google APIs SHA Update
- **Context**: Implementing direct types, CRD, and IdentityV2 for `HypercomputeClusterCluster` (Issue #10277).
- **Problem**: 
  1. The proto definition for `hypercomputecluster` was not present in the default pinned Google APIs SHA (`1765b559c42386788ff0c6412491277b4791107a`).
  2. Bootstrapping `ClusterObservedState` had a circular dependency where `generate-types` pruned `ClusterObservedState` as unreachable because it wasn't in `cluster_types.go`, but adding it to `cluster_types.go` caused `controller-gen` compile errors because it was commented out in `types.generated.go`.
- **Solution**: 
  1. Updated `apis/git.versions` to the newer Google APIs master commit `1526e545e9d26f23b9c5d0f04af17297def8d045` where `hypercomputecluster` was added.
  2. Hand-wrote the `ClusterObservedState` and bootstrapped the unreachable sub-types using a temporary empty `type OrchestratorObservedState struct{}` in `cluster_types.go`, ran the generator once to uncomment sub-types, and then removed the temporary type so that the generator could output the full types.
- **Impact**: Helps next agents understand how to resolve missing proto definitions and handle bootstrapping of output-only observed state types that depend on deep nested generated structures.
