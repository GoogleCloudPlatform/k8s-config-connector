# HTTP Lifecycle Report: `backupdrbackupvault-full`

- **Primary Resource:** `BackupDRBackupVault` (`backupdr.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `backupdrbackupvault-full-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:251c703a0bc804e0`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.`<br>`.backupMinimumEnforcedRetentionDuration` | **from:** `"86400s"`<br>**to:** `"86500s"` | Yes (#5 `PATCH`) | Yes |
| `spec.description` | **from:** `"This is a full backup vault"`<br>**to:** `"This is an updated full backup vault"` | Yes (#5 `PATCH`) | Yes |
| `spec.ignoreInactiveDatasources` | **from:** `true`<br>**to:** `false` | Yes (#5 `PATCH`) | Yes |
| `metadata.labels.another-label` | **from:** *(unset)*<br>**to:** `"another-value"` | Yes (#5 `PATCH`) | Yes |
| `metadata.labels.custom-label` | **from:** `"foo"`<br>**to:** `"bar"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/backupVaults/`<br>`backupdrbackupvault-full-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/backupVaults`<br>`?backupVaultId=backupdrbackupvault-full-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `UPDATE` | `GET` `200` | `.../locations/us-central1/backupVaults/`<br>`backupdrbackupvault-full-${uniqueId}` | Pre-update read |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/backupVaults/`<br>`backupdrbackupvault-full-${uniqueId}` | LRO: update -> done (#6)<br>updateMask: backupMinimumEnforcedRetentionDuration, description, labels |
| 7 | `UPDATE` | `GET` `200` | `.../locations/us-central1/backupVaults/`<br>`backupdrbackupvault-full-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/backupVaults/`<br>`backupdrbackupvault-full-${uniqueId}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_backupdrbackupvault-full.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [ ] **V7 (Authenticity & Hand-Edit Integrity):** Protobuf schema violation: stanza #3 (`GET`) response failed protobuf schema for `@type` "type.googleapis.com/google.cloud.backupdr.v1.BackupVault": proto: (line 1:94): unknown field "backupRetentionInheritance".
  - `<hint_for_agent>`: HAND-EDITED OR HALLUCINATED PROTO PAYLOAD DETECTED: One or more JSON payloads or @type URLs in _http.log violate the protobuf schema in .build/googleapis.pb. Do not hand-edit _http.log; re-record against RealGCP.
