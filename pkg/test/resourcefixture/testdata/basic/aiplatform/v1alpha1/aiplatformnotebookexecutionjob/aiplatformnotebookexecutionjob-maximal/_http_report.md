# HTTP Lifecycle Report: `aiplatformnotebookexecutionjob-maximal`

- **Primary Resource:** `AIPlatformNotebookExecutionJob` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aipnb-max-${uniqueId}`
- **Source Log:** `_http.log` (16 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:60f89700a0a2b1cf`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `StorageBucket` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1.NotebookService/`<br>`GetNotebookExecutionJob` | Pre-create existence check |
| 5 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.NotebookService/`<br>`CreateNotebookExecutionJob` | LRO: create -> done (#6)<br>Created resource |
| 7 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.NotebookService/`<br>`GetNotebookExecutionJob` | Post-create read (matches `create.yaml`) |
| 8 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.NotebookService/`<br>`GetNotebookExecutionJob` | Pre-delete read |
| 9 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.NotebookService/`<br>`GetNotebookExecutionJob` | Pre-delete read |
| 10 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.NotebookService/`<br>`DeleteNotebookExecutionJob` | LRO: delete -> done (sync)<br>Deleted resource |
| 13 | `DELETE` | `DELETE` `204` | `/storage/v1/b/aipnb-max-dep-${uniqueId}/o/`<br>`aipnb-max-${uniqueId}%2Fcontent.html`<br>`?generation=1794626960360161` | Deleted resource |
| 14 | `DELETE` | `DELETE` `204` | `/storage/v1/b/aipnb-max-dep-${uniqueId}/o/`<br>`aipnb-max-${uniqueId}%2Fcontent.ipynb`<br>`?generation=1794626960272129` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformnotebookexecutionjob-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
