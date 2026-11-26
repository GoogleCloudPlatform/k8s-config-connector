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
