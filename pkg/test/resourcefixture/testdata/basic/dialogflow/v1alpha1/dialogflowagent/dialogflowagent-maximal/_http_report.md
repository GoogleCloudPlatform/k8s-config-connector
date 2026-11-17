# HTTP Lifecycle Report: `dialogflowagent-maximal`

- **Primary Resource:** `DialogflowAgent` (`dialogflow.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowagent-${uniqueId} (resourceID: ${projectId})`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:cf1e530c5b9a5f8e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.apiVersion` | **from:** `"API_VERSION_V2_BETA_1"`<br>**to:** `"API_VERSION_V2"` | Yes (#4 `POST`) | Yes |
| `spec.classificationThreshold` | **from:** `0.5`<br>**to:** `0.7` | Yes (#4 `POST`) | Yes |
| `spec.description` | **from:** `"Initial description"`<br>**to:** `"Updated description"` | Yes (#4 `POST`) | Yes |
| `spec.displayName` | **from:** `"max-${uniqueId}"`<br>**to:** `"max-up-${uniqueId}"` | Yes (#4 `POST`) | Yes |
| `spec.enableLogging` | **from:** `true`<br>**to:** `false` | Yes (#4 `POST`) | Yes |
| `spec.matchMode` | **from:** `"MATCH_MODE_HYBRID"`<br>**to:** `"MATCH_MODE_ML_ONLY"` | Yes (#4 `POST`) | Yes |
| `spec.supportedLanguageCodes` | **from:** `["fr", "es"]`<br>**to:** `["de"]` | Yes (#4 `POST`) | Yes |
| `spec.timeZone` | **from:** `"America/New_York"`<br>**to:** `"Europe/London"` | Yes (#4 `POST`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../agent` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../agent` | Created resource |
| 3 | `CREATE` | `GET` `200` | `.../agent` | Post-create read (matches `create.yaml`) |
| 4 | `UPDATE` | `POST` `200` | `.../agent` | Updated resource (`POST`) |
| 5 | `UPDATE` | `GET` `200` | `.../agent` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../agent` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 8/8 mutated fields from `update.yaml` sent in `UPDATE` phase (`POST`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowagent-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
