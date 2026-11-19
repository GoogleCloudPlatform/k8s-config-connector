### [2026-11-15] BigtableCluster Greenfield Prepopulation
- **Context**: Prepopulating direct KRM types, identity, and `judgement_queue.yaml` for `BigtableCluster` under `apis/bigtable/v1alpha1`.
- **Observations**:
  - `BigtableCluster` had existing KRM types and CRD schema from earlier development, referencing `bigtablev1beta1.InstanceRef` and `kmsv1beta1.KMSCryptoKeyRef`.
  - Adding `// +kcc:required-from-proto` to `BigtableClusterSpec` allowed `controllerbuilder generate-types --prepopulate-resource` to generate `apis/bigtable/judgement_queue.yaml` with `reason: untriaged-bulk-generation` while preserving the existing types and CRD backwards compatibility.
  - Because `bigtablev1beta1.InstanceRef` uses `NormalizedExternal` rather than standard `refs.Ref`, `getIdentityFromBigtableClusterSpec` is implemented in `apis/bigtable/v1alpha1/bigtablecluster_identity.go` as a hand-written declaration override. `controllerbuilder generate-identity` then emits `BigtableClusterIdentityFormat`, `BigtableClusterRef`, and conformance tests deterministically.
  - Legacy `ClusterIdentity` and `ClusterRef` types are retained in `bigtablecluster_identity.go` and `bigtablecluster_reference.go` because they are referenced by `BigtableBackup` and `BigtableMemoryLayer`.
  - The `v1beta1` `generate-types` block for `BigtableAppProfile`, `BigtableTable`, and `BigtableGCPolicy` is maintained in `apis/bigtable/generate.sh`. To prevent the newly added parameterless proto message `AppProfile.StandardIsolation.MemoryConfig` from introducing an empty object into the released `v1beta1` `BigtableAppProfile` CRD, `AppProfile_StandardIsolation` is explicitly defined in `apis/bigtable/v1beta1/bigtableappprofile_types.go` following the established manual override pattern.

### [2026-11-15] BigtableCluster Greenfield Types Judgement Triage
- **Context**: Triaging `apis/bigtable/judgement_queue.yaml` and finalizing direct KRM types, identity, and references for `BigtableCluster`.
- **Observations**:
  - Implemented `refs.Ref` methods (`GetGVK`, `GetNamespacedName`, `GetExternal`, `SetExternal`, `ValidateExternal`, `Normalize`) on `bigtablev1beta1.InstanceRef`, enabling `controllerbuilder generate-identity` to generate `getIdentityFromBigtableClusterSpec` automatically in `bigtablecluster_identity.generated.go` without handwritten overrides.
  - Preserved existing `BigtableClusterParent` with `spec.instanceRef` (`bigtablev1beta1.InstanceRef`) to maintain CRD backward compatibility for the existing released CRD.
  - Retained legacy `ClusterIdentity` and `ClusterRef` in `bigtablecluster_identity.go` and `bigtablecluster_reference.go` to preserve compatibility with dependent resources (`BigtableBackup` and `BigtableMemoryLayer`).
  - Resolved `location-or-parent-ref` and `authoritative-proto-reference` for `.spec.location` by keeping it as a plain string: Bigtable clusters are zonal while their parent instance is not, and `locations.googleapis.com/Location` is not a KCC-managed resource.
  - Resolved `required-not-enforced` on autoscaling fields by keeping them optional in KRM to support clusters configured with static `serveNodes`.

### [2026-11-16] BigtableCluster Greenfield Direct Controller Implementation
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `BigtableCluster` under Greenfield Step 2.
- **Observations**:
  - `BigtableInstanceAdminClient` from `cloud.google.com/go/bigtable/admin/apiv2` is a gRPC-only client (`m.config.GRPCClientOptions()`); GAPIC REST client does not exist for Bigtable Admin API.
  - Updating a cluster in Bigtable Admin API must use `PartialUpdateCluster` with `*adminpb.PartialUpdateClusterRequest` (`Cluster` and `UpdateMask`) rather than `UpdateCluster`. `UpdateCluster` does not accept a FieldMask and does not support updating `cluster_config.cluster_autoscaling_config`.
  - `CreateCluster` and `PartialUpdateCluster` return Long-Running Operations (`*CreateClusterOperation` and `*PartialUpdateClusterOperation`) whose `Wait(ctx)` returns `*adminpb.Cluster`. `DeleteCluster` is synchronous and returns an immediate error.
  - Bigtable requires each cluster within an instance to be located in a distinct zone (e.g., primary cluster in `us-central1-a`, secondary in `us-central1-b`).
  - Implemented `ID() string` on `*BigtableClusterIdentity` in `apis/bigtable/v1alpha1/bigtablecluster_identity.go` to cleanly expose the cluster ID string for `CreateClusterRequest`.
  - Re-reconciliation testing was enabled by removing `BigtableCluster` from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. Both minimal (updating `serveNodes`) and maximal (updating `clusterConfig` autoscaling limits and targets) fixtures passed all 7/7 automated validation checks against real GCP.


