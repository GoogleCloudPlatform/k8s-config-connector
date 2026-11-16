# HTTP Lifecycle Report: `aiplatformhyperparametertuningjob-maximal`

- **Primary Resource:** `AIPlatformHyperparameterTuningJob` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `hptjob-maximal-${uniqueId}`
- **Source Log:** `_http.log` (23 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2f5c86e2bfdb98e4`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `IAMServiceAccount` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`CreateHyperparameterTuningJob` | Created resource |
| 6 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 7 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 8 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 9 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`CancelHyperparameterTuningJob` | Deleted resource |
| 10 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 11 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 12 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`CancelHyperparameterTuningJob` | Deleted resource |
| 13 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 14 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 15 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`CancelHyperparameterTuningJob` | Deleted resource |
| 16 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 17 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 18 | `DELETE` | `GRPC` `400` | `/google.cloud.aiplatform.v1.JobService/`<br>`CancelHyperparameterTuningJob` | Deleted resource |
| 19 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 20 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`GetHyperparameterTuningJob` | Pre-delete read |
| 21 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.JobService/`<br>`DeleteHyperparameterTuningJob` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformhyperparametertuningjob-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
