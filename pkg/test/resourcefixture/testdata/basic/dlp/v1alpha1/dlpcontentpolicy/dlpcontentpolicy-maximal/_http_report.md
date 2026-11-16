# HTTP Lifecycle Report: `dlpcontentpolicy-maximal`

- **Primary Resource:** `DLPContentPolicy` (`dlp.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dlpcontentpolicy-maximal-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:42ccaa5789e9dede`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.defaultAction.returnVerdict` | **from:** `"ALLOW"`<br>**to:** `"BLOCK"` | Yes (#4 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal DLP Content Policy"`<br>**to:** `"Maximal DLP Content Policy Updated"` | Yes (#4 `PATCH`) | Yes |
| `spec.`<br>`.failedToScanSupportedFileType.`<br>`.returnVerdict` | **from:** `"ALLOW"`<br>**to:** `"BLOCK"` | Yes (#4 `PATCH`) | Yes |
| `spec.inputTooLarge.returnVerdict` | **from:** `"ALLOW"`<br>**to:** `"BLOCK"` | Yes (#4 `PATCH`) | Yes |
| `spec.inspectConfig.includeQuote` | **from:** `true`<br>**to:** `false` | Yes (#4 `PATCH`) | Yes |
| `spec.inspectConfig.limits.`<br>`.maxFindingsPerItem` | **from:** `10`<br>**to:** `20` | Yes (#4 `PATCH`) | Yes |
| `spec.inspectConfig.minLikelihood` | **from:** `"POSSIBLE"`<br>**to:** `"LIKELY"` | Yes (#4 `PATCH`) | Yes |
| `spec.unsupportedFileType.`<br>`.returnVerdict` | **from:** `"ALLOW"`<br>**to:** `"BLOCK"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/contentPolicies/`<br>`dlpcontentpolicy-maximal-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/contentPolicies` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/global/contentPolicies/`<br>`dlpcontentpolicy-maximal-${uniqueId}` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/global/contentPolicies/`<br>`dlpcontentpolicy-maximal-${uniqueId}` | updateMask: defaultAction, displayName, failedToScanSupportedFileType, inputTooLarge, inspectConfig, unsupportedFileType |
| 5 | `DELETE` | `GET` `200` | `.../locations/global/contentPolicies/`<br>`dlpcontentpolicy-maximal-${uniqueId}` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/global/contentPolicies/`<br>`dlpcontentpolicy-maximal-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 8/8 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dlpcontentpolicy-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
