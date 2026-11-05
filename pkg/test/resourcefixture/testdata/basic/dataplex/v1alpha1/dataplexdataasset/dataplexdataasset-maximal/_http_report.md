# HTTP Lifecycle Report: `dataplexdataasset-maximal`

- **Primary Resource:** `DataplexDataAsset` (`dataplex.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dataplexdataasset-maximal-${uniqueId}`
- **Source Log:** `_http.log` (18 total HTTP entries, 12 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:8316c31687c7549b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigQueryDataset` | 5 | `CREATE`, `DELETE` |
| `DataplexDataProduct` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 8 | `CREATE` | `GET` `404` | `.../locations/us-central1/dataProducts/dpmax-${uniqueId}/dataAssets/`<br>`dataplexdataasset-maximal-${uniqueId}` | Pre-create existence check |
| 9 | `CREATE` | `POST` `200` | `.../locations/us-central1/dataProducts/dpmax-${uniqueId}/`<br>`dataAssets`<br>`?dataAssetId=dataplexdataasset-maximal-${uniqueId}` | LRO: create -> done (#10)<br>Created resource |
| 11 | `DELETE` | `GET` `200` | `.../locations/us-central1/dataProducts/dpmax-${uniqueId}/dataAssets/`<br>`dataplexdataasset-maximal-${uniqueId}` | Pre-delete read |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/dataProducts/dpmax-${uniqueId}/dataAssets/`<br>`dataplexdataasset-maximal-${uniqueId}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataplexdataasset-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
