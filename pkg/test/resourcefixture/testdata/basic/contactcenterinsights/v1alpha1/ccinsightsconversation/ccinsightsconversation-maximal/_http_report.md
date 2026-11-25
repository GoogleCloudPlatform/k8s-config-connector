# HTTP Lifecycle Report: `ccinsightsconversation-maximal`

- **Primary Resource:** `CCInsightsConversation` (`contactcenterinsights.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ccinsightsconv-max-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:02c4c5c36ffe733d`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.agentID` | **from:** `"agent-1"`<br>**to:** `"agent-2"` | Yes (#4 `PATCH`) | Yes |
| `spec.labels.key1` | **from:** `"val1"`<br>**to:** `"val1_updated"` | Yes (#4 `PATCH`) | Yes |
| `spec.labels.key2` | **from:** `"val2"`<br>**to:** *(unset)* | Yes (#4 `PATCH`) | Yes |
| `spec.labels.key3` | **from:** *(unset)*<br>**to:** `"val3"` | Yes (#4 `PATCH`) | Yes |
| `spec.languageCode` | **from:** `"en-US"`<br>**to:** `"es-ES"` | Yes (#4 `PATCH`) | Yes |
| `spec.qualityMetadata.agentInfo` | **from:** `[{"agentID": "agent-1", "agentType": "HUMAN_AGENT", "displayN...`<br>**to:** `[{"agentID": "agent-2", "agentType": "HUMAN_AGENT", "displayN...` | Yes (#4 `PATCH`) | Yes |
| `spec.qualityMetadata.`<br>`.customerSatisfactionRating` | **from:** `5`<br>**to:** `4` | Yes (#4 `PATCH`) | Yes |
| `spec.qualityMetadata.menuPath` | **from:** `"main_menu"`<br>**to:** `"menu_updated"` | Yes (#4 `PATCH`) | Yes |
| `spec.qualityMetadata.`<br>`.waitDuration` | **from:** `"10s"`<br>**to:** `"15s"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/conversations/`<br>`ccinsightsconv-max-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`conversations`<br>`?conversationId=ccinsightsconv-max-${uniqueId}` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/us-central1/conversations/`<br>`ccinsightsconv-max-${uniqueId}` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/conversations/`<br>`ccinsightsconv-max-${uniqueId}` | updateMask: agentId, labels, languageCode, qualityMetadata |
| 5 | `DELETE` | `GET` `200` | `.../locations/us-central1/conversations/`<br>`ccinsightsconv-max-${uniqueId}` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/conversations/`<br>`ccinsightsconv-max-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 9/9 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_ccinsightsconversation-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
