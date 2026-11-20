# HTTP Lifecycle Report: `migrationcenterimportdatafile-minimal`

- **Primary Resource:** `MigrationCenterImportDataFile` (`migrationcenter.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `migrationcenterimportdatafile-minimal-${uniqueId}`
- **Source Log:** `_http.log` (19 total HTTP entries, 13 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ade20d41423de98f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `MigrationCenterImportJob` | 6 | `CREATE`, `DELETE` |
| `MigrationCenterSource` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 8 | `CREATE` | `GET` `404` | `.../locations/us-central1/importJobs/migrationcenterimportjob-idf-min-${uniqueId}/importDataFiles/`<br>`migrationcenterimportdatafile-minimal-${uniqueId}` | Pre-create existence check |
| 9 | `CREATE` | `POST` `200` | `.../locations/us-central1/importJobs/migrationcenterimportjob-idf-min-${uniqueId}/`<br>`importDataFiles`<br>`?importDataFileId=migrationcenterimportdatafile-minimal-${uniqueId}` | LRO: create -> done (#10)<br>Created resource |
| 11 | `DELETE` | `GET` `200` | `.../locations/us-central1/importJobs/migrationcenterimportjob-idf-min-${uniqueId}/importDataFiles/`<br>`migrationcenterimportdatafile-minimal-${uniqueId}` | Pre-delete read |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/importJobs/migrationcenterimportjob-idf-min-${uniqueId}/importDataFiles/`<br>`migrationcenterimportdatafile-minimal-${uniqueId}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_migrationcenterimportdatafile-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
