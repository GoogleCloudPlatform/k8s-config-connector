### [2026-11-15] BigtableCluster Greenfield Prepopulation
- **Context**: Prepopulating direct KRM types, identity, and `judgement_queue.yaml` for `BigtableCluster` under `apis/bigtable/v1alpha1`.
- **Observations**:
  - `BigtableCluster` had existing KRM types and CRD schema from earlier development, referencing `bigtablev1beta1.InstanceRef` and `kmsv1beta1.KMSCryptoKeyRef`.
  - Adding `// +kcc:required-from-proto` to `BigtableClusterSpec` allowed `controllerbuilder generate-types --prepopulate-resource` to generate `apis/bigtable/judgement_queue.yaml` with `reason: untriaged-bulk-generation` while preserving the existing types and CRD backwards compatibility.
  - Because `bigtablev1beta1.InstanceRef` uses `NormalizedExternal` rather than standard `refs.Ref`, `getIdentityFromBigtableClusterSpec` is implemented in `apis/bigtable/v1alpha1/bigtablecluster_identity.go` as a hand-written declaration override. `controllerbuilder generate-identity` then emits `BigtableClusterIdentityFormat`, `BigtableClusterRef`, and conformance tests deterministically.
  - Legacy `ClusterIdentity` and `ClusterRef` types are retained in `bigtablecluster_identity.go` and `bigtablecluster_reference.go` because they are referenced by `BigtableBackup` and `BigtableMemoryLayer`.
  - The `v1beta1` `generate-types` block for `BigtableAppProfile`, `BigtableTable`, and `BigtableGCPolicy` is maintained in `apis/bigtable/generate.sh`. To prevent the newly added parameterless proto message `AppProfile.StandardIsolation.MemoryConfig` from introducing an empty object into the released `v1beta1` `BigtableAppProfile` CRD, `AppProfile_StandardIsolation` is explicitly defined in `apis/bigtable/v1beta1/bigtableappprofile_types.go` following the established manual override pattern.

