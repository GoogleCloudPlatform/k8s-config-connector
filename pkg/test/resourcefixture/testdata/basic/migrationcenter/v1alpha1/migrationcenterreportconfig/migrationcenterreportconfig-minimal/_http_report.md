# HTTP Lifecycle Report: `migrationcenterreportconfig-minimal`

- **Primary Resource:** `MigrationCenterReportConfig` (`migrationcenter.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `mcrc-${uniqueId}`
- **Source Log:** `_http.log` (19 total HTTP entries, 13 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:183b99a7b75f347a`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `metadata.labels.example-label` | **from:** *(unset)*<br>**to:** `"updated"` | No (0 Update RPCs) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `MigrationCenterGroup` | 6 | `CREATE`, `DELETE` |
| `MigrationCenterPreferenceSet` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 8 | `CREATE` | `GET` `404` | `.../locations/us-central1/reportConfigs/`<br>`mcrc-${uniqueId}` | Pre-create existence check |
| 9 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`reportConfigs`<br>`?reportConfigId=mcrc-${uniqueId}` | LRO: create -> done (#10)<br>Created resource |
| 11 | `CREATE` | `GET` `200` | `.../locations/us-central1/reportConfigs/`<br>`mcrc-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/reportConfigs/`<br>`mcrc-${uniqueId}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 0 mutation RPCs occurred in `UPDATE` phase despite 1 modified field(s) in `update.yaml` (metadata.labels.example-label).
  - `<hint_for_agent>`: SILENT UPDATE NO-OP DETECTED: `update.yaml` modified metadata.labels.example-label, but the controller made zero mutation RPCs (`PATCH`/`PUT`/`POST`) during the UPDATE phase. Check the controller's Update() diff detection and updateMask construction.
- [x] **V3 (End-State):** Final GCP response matches `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
