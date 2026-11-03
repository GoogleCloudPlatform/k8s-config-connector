# HTTP Lifecycle Report: `bigqueryanalyticshubquerytemplate-maximal`

- **Primary Resource:** `BigQueryAnalyticsHubQueryTemplate` (`bigqueryanalyticshub.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `bigqueryanalyticshubquerytemplatemax${uniqueId}`
- **Source Log:** `_http.log` (10 total HTTP entries, 4 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:177e2082506a429f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Sample description for query template ${uniqueId}"`<br>**to:** `"Updated description for query template ${uniqueId}"` | Yes (#6 `PATCH`) | Yes |
| `spec.routine.definitionBody` | **from:** `"bigqueryanalyticshubquerytemplatemax${uniqueId}() AS (SELECT...`<br>**to:** `"bigqueryanalyticshubquerytemplatemax${uniqueId}() AS (SELECT...` | Yes (#6 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigQueryAnalyticsHubDataExchange` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `404` | `.../locations/US/dataExchanges/bigqueryanalyticshubdataexchange${uniqueId}/queryTemplates/`<br>`bigqueryanalyticshubquerytemplatemax${uniqueId}` | Pre-create existence check |
| 4 | `CREATE` | `POST` `200` | `.../locations/US/dataExchanges/bigqueryanalyticshubdataexchange${uniqueId}/`<br>`queryTemplates`<br>`?queryTemplateId=bigqueryanalyticshubquerytemplatemax${uniqueId}` | Created resource |
| 5 | `CREATE` | `GET` `200` | `.../locations/US/dataExchanges/bigqueryanalyticshubdataexchange${uniqueId}/queryTemplates/`<br>`bigqueryanalyticshubquerytemplatemax${uniqueId}` | Post-create read (matches `create.yaml`) |
| 6 | `UPDATE` | `PATCH` `200` | `.../locations/US/dataExchanges/bigqueryanalyticshubdataexchange${uniqueId}/queryTemplates/`<br>`bigqueryanalyticshubquerytemplatemax${uniqueId}` | updateMask: description, routine |
| 7 | `UPDATE` | `GET` `200` | `.../locations/US/dataExchanges/bigqueryanalyticshubdataexchange${uniqueId}/queryTemplates/`<br>`bigqueryanalyticshubquerytemplatemax${uniqueId}` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/US/dataExchanges/bigqueryanalyticshubdataexchange${uniqueId}/queryTemplates/`<br>`bigqueryanalyticshubquerytemplatemax${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_bigqueryanalyticshubquerytemplate-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
