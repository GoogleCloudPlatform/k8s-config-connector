# HTTP Lifecycle Report: `biglakehivetable-minimal`

- **Primary Resource:** `BigLakeHiveTable` (`bigquerybiglake.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `biglakehivetable-min-${uniqueId} (resourceID: biglakehivetable_min_${uniqueId})`
- **Source Log:** `_http.log` (21 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:030a559173a41cfd`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** *(unset)*<br>**to:** `"Updated description for minimal Hive table"` | Yes (#11 `PATCH`) | Yes |
| `spec.storageDescriptor.columns` | **from:** `[{"name": "id", "type": "int"}]`<br>**to:** `[{"name": "id", "type": "int"}, {"name": "name", "type": "str...` | Yes (#11 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigLakeHiveDatabase` | 4 | `CREATE`, `DELETE` |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GRPC` `404` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-create existence check |
| 5 | `CREATE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`CreateHiveCatalog` | Created resource |
| 8 | `CREATE` | `GET` `404` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/biglakehivedatabase_min_${uniqueId}/tables/`<br>`biglakehivetable_min_${uniqueId}` | Post-create eventual-consistency poll |
| 9 | `CREATE` | `POST` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/biglakehivedatabase_min_${uniqueId}/`<br>`tables`<br>`?hiveTableId=biglakehivetable_min_${uniqueId}` | Composite create follow-up (`POST`) |
| 10 | `UPDATE` | `GET` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/biglakehivedatabase_min_${uniqueId}/tables/`<br>`biglakehivetable_min_${uniqueId}` | Pre-update read |
| 11 | `UPDATE` | `PATCH` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/biglakehivedatabase_min_${uniqueId}/tables/`<br>`biglakehivetable_min_${uniqueId}` | updateMask: description, storageDescriptor |
| 12 | `DELETE` | `GET` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/biglakehivedatabase_min_${uniqueId}/tables/`<br>`biglakehivetable_min_${uniqueId}` | Pre-delete read |
| 13 | `DELETE` | `DELETE` `200` | `.../catalogs/biglakehivecatalog_min_${uniqueId}/databases/biglakehivedatabase_min_${uniqueId}/tables/`<br>`biglakehivetable_min_${uniqueId}` | Deleted resource |
| 16 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 17 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 18 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`DeleteHiveCatalog` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_biglakehivetable-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
