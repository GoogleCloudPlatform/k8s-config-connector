# HTTP Lifecycle Report: `featureonlinestore-basic`

- **Primary Resource:** `VertexAIFeatureOnlineStore` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `featureonlinestore${uniqueId}`
- **Source Log:** `_http.log` (11 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:f59bd22581207d3e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.labels.foo` | **from:** `"bar"`<br>**to:** `"baz"` | Yes (#6 `GRPC`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Pre-create existence check |
| 2 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`CreateFeatureOnlineStore` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Pre-update read |
| 6 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`UpdateFeatureOnlineStore` | LRO: update -> done (#7)<br>updateMask: labels |
| 8 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Post-update read (matches `update.yaml`) |
| 9 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Pre-delete read |
| 10 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Pre-delete read |
| 11 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`DeleteFeatureOnlineStore` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_featureonlinestore-basic.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
