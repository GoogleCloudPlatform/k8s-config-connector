### [2026-10-06] NetAppActiveDirectory Greenfield Direct Controller
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for Greenfield resource `NetAppActiveDirectory`.
- **Problem**:
  1. `netBiosPrefix` in GCP NetApp Volumes must be at most 10 characters (otherwise GCP returns `400: Invalid netbios prefix, length must not be longer than 10 characters`).
  2. GCP defaults `organizational_unit` to `"CN=Computers"` on the server side when omitted.
- **Solution**:
  1. Used concise static prefixes (`adnetbios` / `maxnetbios`) for fixture test configurations.
  2. Handled server default for `organizational_unit` in `compareActiveDirectory` to prevent unwanted diffs and extra updates during re-reconciliation.
- **Impact**: All minimal and maximal fixtures successfully verified, recorded against real GCP, and passed re-reconciliation testing.

### [2026-11-26] NetAppStoragePool Types Finalization
- **Context**: Greenfield Step 1b types finalization and judgement queue triage for `NetAppStoragePool`.
- **Hierarchy & References**:
  - `StoragePool` is regional project-parented (`projects/{project}/locations/{location}/storagePools/{storagePool}`), so `ProjectRef` and `Location` were replaced by inlined `*parent.ProjectAndLocationRef`.
  - Converted `network` to `networkRef` (`*computerefs.ComputeNetworkRef`).
  - Converted `active_directory` to `activeDirectoryRef` (`*NetAppActiveDirectoryRef`).
  - Converted `kms_config` to `kmsConfigRef` (`*NetAppKMSConfigRef`).
- **Fields & Conventions**:
  - `spec.labels` was removed in accordance with KCC conventions (`metadata.labels` managed by KCC).
  - Deprecated field `global_access_allowed` was verified omitted from Spec.
  - `spec.zone` was kept as a plain string parameter specifying the active zone for regional storage pools.

### [2026-11-26] NetAppStoragePool Greenfield Direct Controller
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for Greenfield resource `NetAppStoragePool`.
- **GCP NetApp Volumes API Behaviors**:
  1. `network`: GCP converts project ID in `projects/<project-id>/global/networks/<name>` to project number `projects/<project-number>/global/networks/<name>` upon creation, and `network` cannot be updated (returns `400: invalid update mask path "network"`). Added normalization logic in `StoragePoolAdapter.Update` to normalize network string when pointing to the same network.
  2. `zone` & `replicaZone`: These fields are write-only / create-only parameters for regional storage pools and are not returned in GET/LIST responses by the GCP API. Furthermore, including them in `updateMask` causes GCP to fail with `400: invalid update mask path "replica_zone"`. Added normalization in `StoragePoolAdapter.Update` adopting actual state when empty to prevent false update diffs.
  3. `labels`: Handled via `label.GCPLabels` on creation and `common.CompareGreenfieldSpecAndLabels` on update.
- **Fixtures & Tests**:
  - `netappstoragepool-minimal` and `netappstoragepool-maximal` fixtures recorded and verified against real GCP with clean CREATE -> UPDATE -> RE-RECONCILE -> DELETE lifecycles.

### [2026-11-27] NetAppVolume Types Finalization
- **Context**: Greenfield Step 1b types finalization and judgement queue triage for `NetAppVolume`.
- **Hierarchy & References**:
  - `Volume` is regional project-parented (`projects/{project}/locations/{location}/volumes/{volume}`), so `ProjectRef` and `Location` were replaced by inlined `*parent.ProjectAndLocationRef`.
  - Converted `storagePool` to `storagePoolRef` (`*NetAppStoragePoolRef`).
  - Converted `backupConfig.backupPolicies` to `backupPolicyRefs` (`[]BackupPolicyRef`) and `backupConfig.backupVault` to `backupVaultRef` (`*BackupVaultRef`).
  - Converted `blockDevices[].hostGroups` to `hostGroupRefs` (`[]NetAppHostGroupRef`).
  - Converted `hybridReplicationParameters.replication` to `replicationRef` (`*NetappReplicationRef`).
  - Created external-only references `NetAppBackupRef` and `NetAppSnapshotRef` for `restoreParameters.sourceBackupRef` and `restoreParameters.sourceSnapshotRef`.
- **Fields & Conventions**:
  - `spec.labels` is preserved on `NetAppVolumeSpec` per Greenfield direct conventions.
  - Moved child structs `BackupConfig` (plus `BackupConfigObservedState`), `BlockDevice` (plus `BlockDeviceObservedState`), `HybridReplicationParameters`, and `RestoreParameters` into `netappvolume_types.go`.
  - Added custom mapper functions in `pkg/controller/direct/netapp/netappvolume_mapper.go` for `BackupConfig` and `RestoreParameters` to handle irregular pluralization and oneof proto source fields.

### [2026-11-27] NetAppVolume Greenfield Direct Controller
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for Greenfield resource `NetAppVolume`.
- **GCP NetApp Volumes API Behaviors**:
  1. `storagePool`: GCP NetApp Volumes requires `storagePool` in `CreateVolumeRequest` and `UpdateVolumeRequest` payloads to be the short storage pool ID (`<storagePoolId>`) rather than the fully-qualified resource name (`projects/<project>/locations/<location>/storagePools/<name>`). Supplying the full resource name causes the backend to prepend the parent resource prefix, producing `GetInternalStoragePool failed for "projects/.../locations/.../storagePools/projects/...": Malformed name`. Normalized `StoragePool` to `tokens[len(tokens)-1]` in `VolumeAdapter.Create`, `VolumeAdapter.Update`, and the `normalize` callback in `CompareGreenfieldSpec`. In `VolumeAdapter.Export`, formatted the exported `StoragePoolRef.External` back to the canonical fully-qualified resource URL.
  2. `Private Service Access (PSA)`: GCP Cloud NetApp Volumes requires VPC peering with service `netapp.servicenetworking.goog` on the consumer network (e.g. `gcloud services vpc-peerings connect --service=netapp.servicenetworking.goog`).
  3. `Fuzzer FilterStatus for Volume_CloneDetails`: The Go client protobuf struct `pb.Volume_CloneDetails` includes `SplitState` (added in newer SDK releases), which is not part of the pinned KRM `Volume_CloneDetailsObservedState`. Configured `f.FilterStatus` in `netappvolume_fuzzer.generated.go` to zero out `SplitState` during fuzz round-trips.
- **Fixtures & Tests**:
  - `netappvolume-minimal` and `netappvolume-maximal` recorded against real GCP. Both fixtures passed all 7 automated checks (V1 through V7) in `_http_report.md` with complete and clean CREATE -> UPDATE -> RE-RECONCILE -> DELETE lifecycles.


