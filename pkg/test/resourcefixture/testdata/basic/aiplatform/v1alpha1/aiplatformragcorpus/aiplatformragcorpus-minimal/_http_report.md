# HTTP Lifecycle Report: `aiplatformragcorpus-minimal`

- **Primary Resource:** `AIPlatformRAGCorpus` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ragcorpus-min-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2818875d7fc4f92c`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.VertexRagDataService/`<br>`CreateRagCorpus` | LRO: create -> done (#2)<br>Created resource |
| 3 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.VertexRagDataService/`<br>`GetRagCorpus` | Pre-delete read |
| 4 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.VertexRagDataService/`<br>`GetRagCorpus` | Pre-delete read |
| 5 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.VertexRagDataService/`<br>`DeleteRagCorpus` | LRO: delete -> done (#6)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformragcorpus-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
