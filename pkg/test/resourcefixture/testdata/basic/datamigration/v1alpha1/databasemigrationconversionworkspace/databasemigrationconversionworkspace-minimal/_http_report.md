# HTTP Lifecycle Report: `databasemigrationconversionworkspace-minimal`

- **Primary Resource:** `DatabaseMigrationConversionWorkspace` (`datamigration.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `db-mw-min-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:b2816bf93a472825`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `404` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConversionWorkspace` | Pre-create existence check |
| 2 | `CREATE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`CreateConversionWorkspace` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConversionWorkspace` | Post-create read (matches `create.yaml`) |
| 5 | `DELETE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConversionWorkspace` | Pre-delete read |
| 6 | `DELETE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`GetConversionWorkspace` | Pre-delete read |
| 7 | `DELETE` | `GRPC` `200` | `/google.cloud.clouddms.v1.DataMigrationService/`<br>`DeleteConversionWorkspace` | LRO: delete -> done (#8)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_databasemigrationconversionworkspace-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
