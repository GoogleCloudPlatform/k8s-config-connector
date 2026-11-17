# HTTP Lifecycle Report: `dataplexasset-maximal`

- **Primary Resource:** `DataplexAsset` (`dataplex.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dataplexasset-maximal-${uniqueId}`
- **Source Log:** `_http.log` (29 total HTTP entries, 20 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:72f8836b4634738f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Initial description for maximal asset"`<br>**to:** `"Updated description for maximal asset"` | Yes (#14 `PATCH`) | Yes |
| `spec.discoverySpec.csvOptions.`<br>`.delimiter` | **from:** `","`<br>**to:** `"\|"` | Yes (#14 `PATCH`) | Yes |
| `spec.discoverySpec.csvOptions.`<br>`.disableTypeInference` | **from:** `false`<br>**to:** `true` | Yes (#14 `PATCH`) | Yes |
| `spec.discoverySpec.csvOptions.`<br>`.headerRows` | **from:** `1`<br>**to:** `2` | Yes (#14 `PATCH`) | Yes |
| `spec.discoverySpec.`<br>`.excludePatterns` | **from:** `["tmp/*"]`<br>**to:** `["tmp2/*"]` | Yes (#14 `PATCH`) | Yes |
| `spec.discoverySpec.`<br>`.includePatterns` | **from:** `["data/*"]`<br>**to:** `["data2/*"]` | Yes (#14 `PATCH`) | Yes |
| `spec.discoverySpec.jsonOptions.`<br>`.disableTypeInference` | **from:** `false`<br>**to:** `true` | Yes (#14 `PATCH`) | Yes |
| `spec.discoverySpec.schedule` | **from:** `"0 1 * * *"`<br>**to:** `"0 2 * * *"` | Yes (#14 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal Asset ${uniqueId}"`<br>**to:** `"Updated Maximal Asset ${uniqueId}"` | Yes (#14 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataplexLake` | 7 | `CREATE`, `DELETE` |
| `DataplexZone` | 7 | `CREATE`, `DELETE` |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 10 | `CREATE` | `GET` `404` | `.../locations/us-central1/lakes/dplake-max-${uniqueId}/zones/dpzone-max-${uniqueId}/assets/`<br>`dataplexasset-maximal-${uniqueId}` | Pre-create existence check |
| 11 | `CREATE` | `POST` `200` | `.../locations/us-central1/lakes/dplake-max-${uniqueId}/zones/dpzone-max-${uniqueId}/`<br>`assets`<br>`?assetId=dataplexasset-maximal-${uniqueId}` | LRO: create -> done (#12)<br>Created resource |
| 13 | `UPDATE` | `GET` `200` | `.../locations/us-central1/lakes/dplake-max-${uniqueId}/zones/dpzone-max-${uniqueId}/assets/`<br>`dataplexasset-maximal-${uniqueId}` | Pre-update read |
| 14 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/lakes/dplake-max-${uniqueId}/zones/dpzone-max-${uniqueId}/assets/`<br>`dataplexasset-maximal-${uniqueId}` | LRO: update -> done (#15)<br>updateMask: description, discoverySpec, displayName, resourceSpec |
| 16 | `DELETE` | `GET` `200` | `.../locations/us-central1/lakes/dplake-max-${uniqueId}/zones/dpzone-max-${uniqueId}/assets/`<br>`dataplexasset-maximal-${uniqueId}` | Pre-delete read |
| 17 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/lakes/dplake-max-${uniqueId}/zones/dpzone-max-${uniqueId}/assets/`<br>`dataplexasset-maximal-${uniqueId}` | LRO: delete -> done (#18)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 9/9 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataplexasset-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
