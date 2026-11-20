# HTTP Lifecycle Report: `dialogflowconversationprofile-minimal`

- **Primary Resource:** `DialogflowConversationProfile` (`dialogflow.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowconversationprofile-min-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:010862b764ca5a22`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"conv-profile-min-${uniqueId}"`<br>**to:** `"conv-profile-min-${uniqueId}-updated"` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/`<br>`conversationProfiles` | Created resource |
| 2 | `UPDATE` | `GET` `200` | `.../locations/global/conversationProfiles/`<br>`0GNsyoseSV-sTsFkBspSKA` | Pre-update read |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/conversationProfiles/`<br>`0GNsyoseSV-sTsFkBspSKA` | updateMask: displayName |
| 4 | `DELETE` | `GET` `200` | `.../locations/global/conversationProfiles/`<br>`0GNsyoseSV-sTsFkBspSKA` | Pre-delete read |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/conversationProfiles/`<br>`0GNsyoseSV-sTsFkBspSKA` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowconversationprofile-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
