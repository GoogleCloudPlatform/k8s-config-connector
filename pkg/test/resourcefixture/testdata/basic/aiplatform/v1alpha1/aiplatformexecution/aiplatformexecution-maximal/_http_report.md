# HTTP Lifecycle Report: `aiplatformexecution-maximal`

- **Primary Resource:** `AIPlatformExecution` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aiplatformexecution-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:c3a82caa9ef84a02`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A test execution for AIPlatformExecution"`<br>**to:** `"An updated test execution for AIPlatformExecution"` | Yes (#9 `GRPC`) | Yes |
| `spec.displayName` | **from:** `"Maximal Execution"`<br>**to:** `"Maximal Execution Updated"` | Yes (#9 `GRPC`) | Yes |
| `spec.metadata.input_param` | **from:** `"value1"`<br>**to:** `"value2"` | Yes (#9 `GRPC`) | Yes |
| `spec.metadata.key2` | **from:** `123`<br>**to:** `456` | Yes (#9 `GRPC`) | Yes |
| `spec.state` | **from:** `"RUNNING"`<br>**to:** `"COMPLETE"` | Yes (#9 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `Project` | 1 | `CREATE` |
| `VertexAIMetadataStore` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 6 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetExecution` | Pre-create existence check |
| 7 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`CreateExecution` | Created resource |
| 8 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetExecution` | Pre-update read |
| 9 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`UpdateExecution` | updateMask: description, displayName, metadata, state |
| 10 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetExecution` | Pre-delete read |
| 11 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetExecution` | Pre-delete read |
| 12 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`DeleteExecution` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformexecution-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
