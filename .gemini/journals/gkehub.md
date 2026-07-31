# GKEHub Service Journal

### [2026-07-02] GKEHubFleet Type Generation and Google APIs Pinning
- **Context**: Implementing direct types for `GKEHubFleet` (`v1alpha1`) in KCC.
- **Problem**: The previously pinned Google APIs commit (`1765b559c4`) did not include `google/cloud/gkehub/v1/fleet.proto`. Attempts to generate the types failed with `failed to find the proto message google.cloud.gkehub.v1.Fleet`.
- **Solution**: We searched the history and updated `apis/git.versions` to pin googleapis to the newer commit `0fcabfc28371e7bab8107402eb06ad58134ee383`. After pinning this commit, the protobuf compiler and the controllerbuilder successfully compiled the proto and generated all of the required types (including `DefaultClusterConfig`, `FleetLifecycleStateObservedState`, and sub-structs) automatically since they were marked as reachable in `gkehubfleet_types.go`.
- **Impact**: Ensures that future developments or regenerations of GKEHub resources have the correct proto source files fully accessible and aligned.
