# HTTP Lifecycle Report: `biglakehivedatabase-minimal`

- **Primary Resource:** `BigLakeHiveDatabase` (`bigquerybiglake.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `biglakehivedatabase-min-${uniqueId} (resourceID: biglakehivedatabase_min_${uniqueId})`
- **Source Log:** `_http.log` (17 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:85b7806b2a671595`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** *(unset)*<br>**to:** `"Updated description for minimal Hive database"` | Yes (#9 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GRPC` `404` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-create existence check |
| 5 | `CREATE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`CreateHiveCatalog` | Created resource |
| 6 | `CREATE` | `GET` `404` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/`<br>`biglakehivedatabase_min_${uniqueId}` | Post-create eventual-consistency poll |
| 7 | `CREATE` | `POST` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/`<br>`databases`<br>`?hiveDatabaseId=biglakehivedatabase_min_${uniqueId}` | Composite create follow-up (`POST`) |
| 8 | `UPDATE` | `GET` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/`<br>`biglakehivedatabase_min_${uniqueId}` | Pre-update read |
| 9 | `UPDATE` | `PATCH` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/`<br>`biglakehivedatabase_min_${uniqueId}` | updateMask: description |
| 10 | `DELETE` | `GET` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/`<br>`biglakehivedatabase_min_${uniqueId}` | Pre-delete read |
| 11 | `DELETE` | `DELETE` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/`<br>`biglakehivedatabase_min_${uniqueId}` | Deleted resource |
| 12 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 13 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 14 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`DeleteHiveCatalog` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_biglakehivedatabase-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
