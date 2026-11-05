# HTTP Lifecycle Report: `aiplatformonlineevaluator-maximal`

- **Primary Resource:** `AIPlatformOnlineEvaluator` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aiplatformonlineevaluator-max-${uniqueId}`
- **Source Log:** `_http.log` (17 total HTTP entries, 9 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:160dbfba960129f6`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.config.`<br>`.maxEvaluatedSamplesPerRun` | **from:** `100`<br>**to:** `200` | Yes (#8 `GRPC`) | Yes |
| `spec.config.randomSampling.`<br>`.percentage` | **from:** `50`<br>**to:** `80` | Yes (#8 `GRPC`) | Yes |
| `spec.displayName` | **from:** `"onlineevaluator-max-${uniqueId}"`<br>**to:** `"onlineevaluator-max-updated-${uniqueId}"` | Yes (#8 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `AIPlatformReasoningEngine` | 9 | `CREATE`, `UPDATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`CreateOnlineEvaluator` | Created resource |
| 6 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Post-create read (matches `create.yaml`) |
| 7 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Pre-update read |
| 8 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`UpdateOnlineEvaluator` | updateMask: config, displayName |
| 10 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Post-update read (matches `update.yaml`) |
| 11 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Pre-delete read |
| 12 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Pre-delete read |
| 13 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`DeleteOnlineEvaluator` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformonlineevaluator-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
