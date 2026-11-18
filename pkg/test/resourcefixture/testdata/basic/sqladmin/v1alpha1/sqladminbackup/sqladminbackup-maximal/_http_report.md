# HTTP Lifecycle Report: `sqladminbackup-maximal`

- **Primary Resource:** `SQLAdminBackup` (`sqladmin.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `sqladminbackup-maximal-${uniqueId}`
- **Source Log:** `_http.log` (11 total HTTP entries, 9 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:0091a855d6f4704a`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `SQLInstance` | 9 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `200` | `.../instances/sqlinstance-max-${uniqueId}/backupRuns/`<br>`1794960364195` | Pre-create read |
| 8 | `DELETE` | `DELETE` `200` | `.../instances/sqlinstance-max-${uniqueId}/backupRuns/`<br>`1794960364195` | LRO: delete -> pending (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource acquired via existing state (`GET 200 OK` in `CREATE` phase).
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_sqladminbackup-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
