# HTTP Lifecycle Report: `bigqueryrowaccesspolicy-minimal`

- **Primary Resource:** `BigQueryRowAccessPolicy` (`bigquery.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `bigqueryrowaccesspolicy${uniqueId} (resourceID: bqrap_${uniqueId})`
- **Source Log:** `_http.log` (20 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:396a02140733695a`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.filterPredicate` | **from:** `"region = \"US\""`<br>**to:** `"region = \"EU\""` | Yes (#11 `PUT`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigQueryDataset` | 5 | `CREATE`, `DELETE` |
| `BigQueryTable` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `404` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}` | Pre-create existence check |
| 8 | `CREATE` | `POST` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/`<br>`rowAccessPolicies` | Created resource |
| 9 | `UPDATE` | `GET` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}` | Pre-update read |
| 10 | `UPDATE` | `POST` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}:getIamPolicy` | Pre-update read |
| 11 | `UPDATE` | `PUT` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}` | Updated resource (`PUT`) |
| 12 | `DELETE` | `GET` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}` | Pre-delete read |
| 13 | `DELETE` | `POST` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}:getIamPolicy` | Pre-delete read |
| 14 | `DELETE` | `GET` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}` | Pre-delete read |
| 15 | `DELETE` | `POST` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}:getIamPolicy` | Pre-delete read |
| 16 | `DELETE` | `DELETE` `200` | `.../datasets/${datasetID}/tables/bigquerytable_${uniqueId}/rowAccessPolicies/`<br>`bqrap_${uniqueId}`<br>`?force=true` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PUT`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_bigqueryrowaccesspolicy-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
