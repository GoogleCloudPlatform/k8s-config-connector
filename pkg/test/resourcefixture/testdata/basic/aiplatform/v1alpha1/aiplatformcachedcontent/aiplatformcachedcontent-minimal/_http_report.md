# HTTP Lifecycle Report: `aiplatformcachedcontent-minimal`

- **Primary Resource:** `AIPlatformCachedContent` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aiplatformcachedcontent-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:00e78ce69209a1b7`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.ttl` | **from:** `"300s"`<br>**to:** `"600s"` | Yes (#3 `GRPC`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.GenAiCacheService/`<br>`CreateCachedContent` | Created resource |
| 2 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.GenAiCacheService/`<br>`GetCachedContent` | Pre-update read |
| 3 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.GenAiCacheService/`<br>`UpdateCachedContent` | updateMask: expireTime |
| 4 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.GenAiCacheService/`<br>`GetCachedContent` | Pre-delete read |
| 5 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.GenAiCacheService/`<br>`GetCachedContent` | Pre-delete read |
| 6 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.GenAiCacheService/`<br>`DeleteCachedContent` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformcachedcontent-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
