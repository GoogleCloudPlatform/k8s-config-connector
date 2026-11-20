# HTTP Lifecycle Report: `dialogflowconversationprofile-maximal`

- **Primary Resource:** `DialogflowConversationProfile` (`dialogflow.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowconversationprofile-max-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:df6bd385a2026d28`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"conv-profile-max-${uniqueId}"`<br>**to:** `"conv-profile-max-${uniqueId}-updated"` | Yes (#3 `PATCH`) | Yes |
| `spec.loggingConfig.`<br>`.enableStackdriverLogging` | **from:** `true`<br>**to:** `false` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/`<br>`conversationProfiles` | Created resource |
| 2 | `UPDATE` | `GET` `200` | `.../locations/global/conversationProfiles/`<br>`HtnzaOUwRTaVeS7PEaMQMw` | Pre-update read |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/conversationProfiles/`<br>`HtnzaOUwRTaVeS7PEaMQMw` | updateMask: displayName, loggingConfig |
| 4 | `DELETE` | `GET` `200` | `.../locations/global/conversationProfiles/`<br>`HtnzaOUwRTaVeS7PEaMQMw` | Pre-delete read |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/conversationProfiles/`<br>`HtnzaOUwRTaVeS7PEaMQMw` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowconversationprofile-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
