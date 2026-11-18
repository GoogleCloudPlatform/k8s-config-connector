# HTTP Lifecycle Report: `aiplatformnasjob-maximal`

- **Primary Resource:** `AIPlatformNasJob` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `nasjob-maximal-${uniqueId}`
- **Source Log:** `_http.log` (10 total HTTP entries, 5 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:b1367e3b842deac1`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `IAMServiceAccount` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`CreateNasJob` | Created resource |
| 5 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetNasJob` | Pre-delete read |
| 6 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetNasJob` | Pre-delete read |
| 7 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetNasJob` | Pre-delete read |
| 8 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`DeleteNasJob` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformnasjob-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
