# HTTP Lifecycle Report: `aiplatformfeatureview-maximal`

- **Primary Resource:** `AIPlatformFeatureView` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aiplatformfeatureview-${uniqueId} (resourceID: featureview${uniqueId})`
- **Source Log:** `_http.log` (26 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:826fc78c375f3fa9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.syncConfig.cron` | **from:** `"0 0 * * *"`<br>**to:** `"0 12 * * *"` | Yes (#15 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigQueryDataset` | 5 | `CREATE`, `DELETE` |
| `BigQueryTable` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Pre-create existence check |
| 8 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`CreateFeatureOnlineStore` | LRO: create -> done (#9)<br>Created resource |
| 10 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureView` | Post-create eventual-consistency poll |
| 11 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`CreateFeatureView` | LRO: create -> done (#12)<br>Composite create follow-up (`GRPC`) |
| 13 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureView` | Post-create read (matches `create.yaml`) |
| 14 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureView` | Pre-update read |
| 15 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`UpdateFeatureView` | LRO: update -> done (sync)<br>updateMask: syncConfig |
| 16 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureView` | Post-update read (matches `update.yaml`) |
| 17 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureView` | Pre-delete read |
| 18 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureView` | Pre-delete read |
| 19 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`DeleteFeatureView` | LRO: delete -> done (sync)<br>Deleted resource |
| 20 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Pre-delete read |
| 21 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`GetFeatureOnlineStore` | Pre-delete read |
| 22 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1.FeatureOnlineStoreAdminService/`<br>`DeleteFeatureOnlineStore` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformfeatureview-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
