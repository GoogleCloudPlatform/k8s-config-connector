# HTTP Lifecycle Report: `projectsink-gsbucketref`

- **Primary Resource:** `LoggingLogSink` (`logging.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `logsink-${uniqueId}`
- **Source Log:** `_http.log` (10 total HTTP entries, 5 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:801dd6a0f43f4347`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `StorageBucket` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 2 | `CREATE` | `POST` `200` | `/storage/v1/b?project=${projectId}` | Created resource |
| 4 | `CREATE` | `GET` `404` | `.../sinks/logsink-${uniqueId}` | Post-create eventual-consistency poll |
| 5 | `CREATE` | `POST` `200` | `.../sinks?uniqueWriterIdentity=true` | Composite create follow-up (`POST`) |
| 6 | `DELETE` | `GET` `200` | `.../sinks/logsink-${uniqueId}` | Pre-delete read |
| 7 | `DELETE` | `DELETE` `200` | `.../sinks/logsink-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_projectsink-gsbucketref.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
