# HTTP Lifecycle Report: `aiplatformonlineevaluator-minimal`

- **Primary Resource:** `AIPlatformOnlineEvaluator` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aiplatformonlineevaluator-${uniqueId}`
- **Source Log:** `_http.log` (13 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:37d5c2dc7cd4204f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `AIPlatformReasoningEngine` | 8 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`CreateOnlineEvaluator` | Created resource |
| 6 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Post-create read (matches `create.yaml`) |
| 7 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Pre-delete read |
| 8 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`GetOnlineEvaluator` | Pre-delete read |
| 9 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/`<br>`DeleteOnlineEvaluator` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformonlineevaluator-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
