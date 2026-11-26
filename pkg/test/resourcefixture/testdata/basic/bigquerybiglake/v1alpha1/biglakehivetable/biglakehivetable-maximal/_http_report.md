# HTTP Lifecycle Report: `biglakehivetable-maximal`

- **Primary Resource:** `BigLakeHiveTable` (`bigquerybiglake.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `biglakehivetable-max-${uniqueId} (resourceID: biglakehivetable_max_${uniqueId})`
- **Source Log:** `_http.log` (21 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:baeab43e8715dce6`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal BigLake Hive Table"`<br>**to:** `"Updated Maximal BigLake Hive Table"` | Yes (#11 `PATCH`) | Yes |
| `spec.parameters.environment` | **from:** `"staging"`<br>**to:** `"production"` | Yes (#11 `PATCH`) | Yes |
| `spec.parameters.owner` | **from:** `"kcc-test"`<br>**to:** `"kcc-test-updated"` | Yes (#11 `PATCH`) | Yes |
| `spec.partitionKeys` | **from:** `[{"comment": "Partition by date", "name": "date", "type": "st...`<br>**to:** `[{"comment": "Partition by date updated", "name": "date", "ty...` | Yes (#11 `PATCH`) | Yes |
| `spec.storageDescriptor.columns` | **from:** `[{"comment": "User identifier", "name": "user_id", "type": "i...`<br>**to:** `[{"comment": "User identifier", "name": "user_id", "type": "i...` | Yes (#11 `PATCH`) | Yes |
| `spec.storageDescriptor.`<br>`.numBuckets` | **from:** `4`<br>**to:** `8` | Yes (#11 `PATCH`) | Yes |
| `spec.storageDescriptor.`<br>`.parameters.key` | **from:** `"value"`<br>**to:** `"updated_value"` | Yes (#11 `PATCH`) | Yes |
| `spec.storageDescriptor.serdeInfo.`<br>`.description` | **from:** `"Parquet SerDe"`<br>**to:** `"Updated Parquet SerDe"` | Yes (#11 `PATCH`) | Yes |
| `spec.storageDescriptor.serdeInfo.`<br>`.parameters.serde_param` | **from:** `"serde_val"`<br>**to:** `"updated_serde_val"` | Yes (#11 `PATCH`) | Yes |
| `spec.storageDescriptor.sortCols` | **from:** `[{"col": "user_id", "order": 1}]`<br>**to:** `[{"col": "user_id", "order": 0}]` | Yes (#11 `PATCH`) | Yes |

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
| 8 | `CREATE` | `GET` `404` | `.../catalogs/biglakehivecatalog_max_${uniqueId}/databases/biglakehivedatabase_max_${uniqueId}/tables/`<br>`biglakehivetable_max_${uniqueId}` | Post-create eventual-consistency poll |
| 9 | `CREATE` | `POST` `200` | `.../catalogs/biglakehivecatalog_max_${uniqueId}/databases/biglakehivedatabase_max_${uniqueId}/`<br>`tables`<br>`?hiveTableId=biglakehivetable_max_${uniqueId}` | Composite create follow-up (`POST`) |
| 10 | `UPDATE` | `GET` `200` | `.../catalogs/biglakehivecatalog_max_${uniqueId}/databases/biglakehivedatabase_max_${uniqueId}/tables/`<br>`biglakehivetable_max_${uniqueId}` | Pre-update read |
| 11 | `UPDATE` | `PATCH` `200` | `.../catalogs/biglakehivecatalog_max_${uniqueId}/databases/biglakehivedatabase_max_${uniqueId}/tables/`<br>`biglakehivetable_max_${uniqueId}` | updateMask: description, parameters, partitionKeys, storageDescriptor |
| 12 | `DELETE` | `GET` `200` | `.../catalogs/biglakehivecatalog_max_${uniqueId}/databases/biglakehivedatabase_max_${uniqueId}/tables/`<br>`biglakehivetable_max_${uniqueId}` | Pre-delete read |
| 13 | `DELETE` | `DELETE` `200` | `.../catalogs/biglakehivecatalog_max_${uniqueId}/databases/biglakehivedatabase_max_${uniqueId}/tables/`<br>`biglakehivetable_max_${uniqueId}` | Deleted resource |
| 16 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 17 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`GetHiveCatalog` | Pre-delete read |
| 18 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.hive.v1.HiveMetastoreService/`<br>`DeleteHiveCatalog` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 10/10 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_biglakehivetable-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
