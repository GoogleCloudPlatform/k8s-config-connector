# HTTP Lifecycle Report: `migrationcenterreport-maximal`

- **Primary Resource:** `MigrationCenterReport` (`migrationcenter.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `migrationcenterreport-maximal-${uniqueId}`
- **Source Log:** `_http.log` (25 total HTTP entries, 19 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:70093e759c4ecdd6`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `MigrationCenterGroup` | 6 | `CREATE`, `DELETE` |
| `MigrationCenterPreferenceSet` | 7 | `CREATE`, `DELETE` |
| `MigrationCenterReportConfig` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 11 | `CREATE` | `GET` `404` | `.../locations/us-central1/reportConfigs/migrationcenterreportconfig-maximal-${uniqueId}/reports/`<br>`migrationcenterreport-maximal-${uniqueId}` | Pre-create existence check |
| 12 | `CREATE` | `POST` `200` | `.../locations/us-central1/reportConfigs/migrationcenterreportconfig-maximal-${uniqueId}/`<br>`reports`<br>`?reportId=migrationcenterreport-maximal-${uniqueId}` | LRO: create -> done (#13)<br>Created resource |
| 14 | `DELETE` | `GET` `200` | `.../locations/us-central1/reportConfigs/migrationcenterreportconfig-maximal-${uniqueId}/reports/`<br>`migrationcenterreport-maximal-${uniqueId}` | Pre-delete read |
| 15 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/reportConfigs/migrationcenterreportconfig-maximal-${uniqueId}/reports/`<br>`migrationcenterreport-maximal-${uniqueId}` | LRO: delete -> done (#16)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_migrationcenterreport-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
