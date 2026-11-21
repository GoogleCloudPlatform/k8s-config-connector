# HTTP Lifecycle Report: `backupdrbackupplanassociation-minimal`

- **Primary Resource:** `BackupDRBackupPlanAssociation` (`backupdr.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `backupdrbackupplanassociation-minimal-${uniqueId}`
- **Source Log:** `_http.log` (33 total HTTP entries, 24 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:53e4970df855227b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BackupDRBackupPlan` | 6 | `CREATE`, `DELETE` |
| `BackupDRBackupVault` | 6 | `CREATE`, `DELETE` |
| `ComputeInstance` | 11 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 1 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 4 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 5 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 18 | `CREATE` | `GET` `404` | `.../locations/us-central1/backupPlanAssociations/`<br>`backupdrbackupplanassociation-minimal-${uniqueId}` | Pre-create existence check |
| 19 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`backupPlanAssociations`<br>`?backupPlanAssociationId=backupdrbackupplanassociation-minimal-${uniqueId}` | LRO: create -> done (#20)<br>Created resource |
| 21 | `DELETE` | `GET` `200` | `.../locations/us-central1/backupPlanAssociations/`<br>`backupdrbackupplanassociation-minimal-${uniqueId}` | Pre-delete read |
| 22 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/backupPlanAssociations/`<br>`backupdrbackupplanassociation-minimal-${uniqueId}` | LRO: delete -> done (#23)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_backupdrbackupplanassociation-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
