# HTTP Lifecycle Report: `dialogflowcxagent-maximal`

- **Primary Resource:** `DialogflowCXAgent` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxagent-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:6adce6bac654def5`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A test agent for Dialogflow CX"`<br>**to:** `"Updated test agent for Dialogflow CX"` | Yes (#3 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"agent-max-${uniqueId}"`<br>**to:** `"agent-max-${uniqueId}-updated"` | Yes (#3 `PATCH`) | Yes |
| `spec.enableSpellCorrection` | **from:** `false`<br>**to:** `true` | Yes (#3 `PATCH`) | Yes |
| `spec.speechToTextSettings.`<br>`.enableSpeechAdaptation` | **from:** *(unset)*<br>**to:** `true` | Yes (#3 `PATCH`) | Yes |
| `spec.supportedLanguageCodes` | **from:** `["fr", "de"]`<br>**to:** `["es"]` | Yes (#3 `PATCH`) | Yes |
| `spec.timeZone` | **from:** `"America/New_York"`<br>**to:** `"America/Los_Angeles"` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/agents` | Created resource |
| 2 | `CREATE` | `GET` `200` | `.../locations/global/agents/`<br>`962ddb85-5a13-429a-9ddd-bb115af6759c` | Post-create read (matches `create.yaml`) |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/`<br>`962ddb85-5a13-429a-9ddd-bb115af6759c` | updateMask: displayName, supportedLanguageCodes, timeZone, description, speechToTextSettings, enableSpellCorrection |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/`<br>`962ddb85-5a13-429a-9ddd-bb115af6759c` | Post-update read (matches `update.yaml`) |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/`<br>`962ddb85-5a13-429a-9ddd-bb115af6759c` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 6/6 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxagent-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
