# BigtableMemoryLayer Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema Design**:
   - The GCP API for `MemoryLayer` (`google.bigtable.admin.v2.MemoryLayer`) is a singleton child of `Cluster` with resource name pattern: `projects/{project}/instances/{instance}/clusters/{cluster}/memoryLayer`.
   - The GCP API provides no user-configurable spec fields; `memoryConfig` is sent empty `{}` to enable and cleared to disable, and its child `storageSizeGib` is output-only. `state` is output-only.
   - Therefore, in KRM, creating `BigtableMemoryLayer` declaratively enables the memory layer on the referenced `spec.clusterRef`.
   - `ResourceID` is present on the spec for metadata naming consistency.
   - Added `cnrm.cloud.google.com/stability-level=alpha` kubebuilder annotation.

2. **Identity & Reference Pattern**:
   - Implemented `BigtableMemoryLayerIdentity` using `gcpurls.Template[BigtableMemoryLayerIdentity]("bigtable.googleapis.com", "projects/{project}/instances/{instance}/clusters/{cluster}/memoryLayer")`.
   - Created `BigtableMemoryLayerRef` in `apis/bigtable/v1alpha1/bigtablememorylayer_reference.go` implementing `refs.Ref` and registering with `refs.Register`.
   - `Normalize` delegates directly to `refs.Normalize` as required for direct/greenfield resources with status-based identity.
   - Added template exception in `pkg/gcpurls/registry_test.go` as `MemoryLayer` is not yet part of CAI metadata definitions.
   - Added comprehensive identity and external reference parsing unit tests in `apis/bigtable/v1alpha1/bigtablememorylayer_identity_test.go`.
