### [2026-07-02] Overriding Nested Proto Message Types in GKEMulticloud
- **Context**: Implementing the initial KRM types and Identity for `GKEMulticloudAttachedCluster` (Issue #10273).
- **Problem**: The proto definition of the `Fleet` message uses `project` as a plain string representing the Fleet host project name. In KCC, we represent this using a proper `ProjectRef` reference field for strong type safety. However, manually overriding the `Fleet` struct in `gkemulticloudattachedcluster_types.go` caused the automatic `types.generated.go` generator to skip processing the `Fleet` message, which also pruned `FleetObservedState` (needed for output-only membership fields in the ObservedState).
- **Solution**: Defined both `Fleet` (with a custom `ProjectRef` field) and `FleetObservedState` (with the output-only `Membership` field) inside `gkemulticloudattachedcluster_types.go`:
  ```go
  // +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
  type Fleet struct {
      // The host project of the fleet.
      // +required
      ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`
  }

  // +kcc:observedstate:proto=google.cloud.gkemulticloud.v1.Fleet
  type FleetObservedState struct {
      // Output only. The name of the managed Hub Membership resource associated to this cluster.
      // +optional
      Membership *string `json:"membership,omitempty"`
  }
  ```
- **Impact**: Ensures that both spec inputs and output-only state attributes for the nested `Fleet` message are fully type-safe, correctly categorized in spec/observedState, and compile without controller-gen or prunetypes errors.
