# HTTP Lifecycle Report: `biglakeicebergcatalog-maximal`

- **Primary Resource:** `BigLakeIcebergCatalog` (`bigquerybiglake.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `biglakeicebergcatalog-max-${uniqueId}`
- **Source Log:** `_http.log` (19 total HTTP entries, 12 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:23d58c46d2b842f7`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.credentialMode` | **from:** `"CREDENTIAL_MODE_END_USER"`<br>**to:** `"CREDENTIAL_MODE_VENDED_CREDENTIALS"` | Yes (#10 `GRPC`) | Yes |
| `spec.description` | **from:** `"Maximal BigLake Iceberg Catalog"`<br>**to:** `"Updated Maximal BigLake Iceberg Catalog"` | Yes (#10 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `StorageBucket` | 12 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GRPC` `404` | `/google.cloud.biglake.v1.IcebergCatalogService/`<br>`GetIcebergCatalog` | Pre-create existence check |
| 8 | `CREATE` | `GRPC` `200` | `/google.cloud.biglake.v1.IcebergCatalogService/`<br>`CreateIcebergCatalog` | Created resource |
| 9 | `UPDATE` | `GRPC` `200` | `/google.cloud.biglake.v1.IcebergCatalogService/`<br>`GetIcebergCatalog` | Pre-update read |
| 10 | `UPDATE` | `GRPC` `200` | `/google.cloud.biglake.v1.IcebergCatalogService/`<br>`UpdateIcebergCatalog` | updateMask: credentialMode, description |
| 11 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.v1.IcebergCatalogService/`<br>`GetIcebergCatalog` | Pre-delete read |
| 12 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.v1.IcebergCatalogService/`<br>`GetIcebergCatalog` | Pre-delete read |
| 13 | `DELETE` | `GRPC` `200` | `/google.cloud.biglake.v1.IcebergCatalogService/`<br>`DeleteIcebergCatalog` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_biglakeicebergcatalog-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
