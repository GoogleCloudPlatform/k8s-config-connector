# HTTP Lifecycle Report: `aiplatformcontext-maximal`

- **Primary Resource:** `AIPlatformContext` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aiplatformcontext-${uniqueId} (resourceID: context-${uniqueId})`
- **Source Log:** `_http.log` (14 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e9bb43b85a8d93a3`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Initial description for maximal context"`<br>**to:** `"Updated description for maximal context"` | Yes (#9 `GRPC`) | Yes |
| `spec.displayName` | **from:** `"Maximal Context ${uniqueId}"`<br>**to:** `"Updated Maximal Context ${uniqueId}"` | Yes (#9 `GRPC`) | Yes |
| `spec.metadata.framework` | **from:** `"tensorflow"`<br>**to:** `"pytorch"` | Yes (#9 `GRPC`) | Yes |
| `metadata.labels.test-label` | **from:** `"initial-value"`<br>**to:** `"updated-value"` | Yes (#9 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `Project` | 1 | `CREATE` |
| `VertexAIMetadataStore` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 6 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetContext` | Pre-create existence check |
| 7 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`CreateContext` | Created resource |
| 8 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetContext` | Pre-update read |
| 9 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`UpdateContext` | updateMask: description, displayName, labels, metadata |
| 10 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetContext` | Pre-delete read |
| 11 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`GetContext` | Pre-delete read |
| 12 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.MetadataService/`<br>`DeleteContext` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformcontext-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
