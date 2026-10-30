# HTTP Lifecycle Report: `datamigrationconnectionprofile-minimal`

- **Primary Resource:** `DatamigrationConnectionProfile` (`datamigration.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dmsconnprofile-${uniqueId}`
- **Source Log:** `_http.log` (10 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:efc8787e61e7128b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"dms-profile-${uniqueId}"`<br>**to:** `"dms-profile-updated-${uniqueId}"` | Yes (#5 `GRPC`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `404` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConnectionProfile` | Pre-create existence check |
| 2 | `CREATE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`CreateConnectionProfile` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConnectionProfile` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`UpdateConnectionProfile` | LRO: update -> done (#6)<br>updateMask: displayName |
| 7 | `UPDATE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConnectionProfile` | Post-update read (matches `update.yaml`) |
| 8 | `UPDATE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConnectionProfile` | Post-update read (matches `update.yaml`) |
| 9 | `DELETE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`DeleteConnectionProfile` | LRO: delete -> done (#10)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_datamigrationconnectionprofile-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
