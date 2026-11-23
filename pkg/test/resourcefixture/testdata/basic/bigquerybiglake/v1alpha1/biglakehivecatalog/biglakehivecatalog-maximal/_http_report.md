# HTTP Lifecycle Report: `biglakehivecatalog-maximal`

- **Primary Resource:** `BigLakeHiveCatalog` (`bigquerybiglake.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `biglakehivecatalog-max-${uniqueId} (resourceID: biglakehivecatalog_max_${uniqueId})`
- **Source Log:** `_http.log` (13 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:939a9e967a4bb969`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal BigLake Hive Catalog"`<br>**to:** `"Updated Maximal BigLake Hive Catalog"` | Yes (#7 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GRPC` `404` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-create existence check |
| 5 | `CREATE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`CreateHiveCatalog` | Created resource |
| 6 | `UPDATE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-update read |
| 7 | `UPDATE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`UpdateHiveCatalog` | updateMask: description |
| 8 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 9 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 10 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`DeleteHiveCatalog` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_biglakehivecatalog-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
