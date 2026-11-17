# HTTP Lifecycle Report: `cesguardrail-maximal`

- **Primary Resource:** `CESGuardrail` (`ces.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cesguardrail-maximal-${uniqueId} (resourceID: cesguardrail-max-${uniqueId})`
- **Source Log:** `_http.log` (12 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2ff5811e5ba690f1`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.action.respondImmediately.`<br>`.responses` | **from:** `[{"disabled": false, "text": "This content is blocked by guar...`<br>**to:** `[{"disabled": true, "text": "Updated blocked content response...` | Yes (#7 `PATCH`) | Yes |
| `spec.contentFilter.`<br>`.bannedContents` | **from:** `["banned_general"]`<br>**to:** `["banned_general_updated"]` | Yes (#7 `PATCH`) | Yes |
| `spec.contentFilter.`<br>`.bannedContentsInAgentResponse` | **from:** `["banned_agent"]`<br>**to:** `["banned_agent_updated"]` | Yes (#7 `PATCH`) | Yes |
| `spec.contentFilter.`<br>`.bannedContentsInUserInput` | **from:** `["banned_user"]`<br>**to:** `["banned_user_updated"]` | Yes (#7 `PATCH`) | Yes |
| `spec.contentFilter.`<br>`.disregardDiacritics` | **from:** `true`<br>**to:** `false` | Yes (#7 `PATCH`) | Yes |
| `spec.contentFilter.matchType` | **from:** `"WORD_BOUNDARY_STRING_MATCH"`<br>**to:** `"SIMPLE_STRING_MATCH"` | Yes (#7 `PATCH`) | Yes |
| `spec.description` | **from:** `"Maximal CESGuardrail for testing"`<br>**to:** `"Updated maximal CESGuardrail for testing"` | Yes (#7 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"cesguardrail-maximal-${uniqueId}"`<br>**to:** `"cesguardrail-max-updated-${uniqueId}"` | Yes (#7 `PATCH`) | Yes |
| `spec.enabled` | **from:** `true`<br>**to:** `false` | Yes (#7 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `CESApp` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us/apps/cesapp-${uniqueId}/guardrails/`<br>`cesguardrail-max-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us/apps/cesapp-${uniqueId}/`<br>`guardrails`<br>`?guardrailId=cesguardrail-max-${uniqueId}` | Created resource |
| 6 | `UPDATE` | `GET` `200` | `.../locations/us/apps/cesapp-${uniqueId}/guardrails/`<br>`cesguardrail-max-${uniqueId}` | Pre-update read |
| 7 | `UPDATE` | `PATCH` `200` | `.../locations/us/apps/cesapp-${uniqueId}/guardrails/`<br>`cesguardrail-max-${uniqueId}` | updateMask: action, contentFilter, description, displayName, enabled |
| 8 | `DELETE` | `GET` `200` | `.../locations/us/apps/cesapp-${uniqueId}/guardrails/`<br>`cesguardrail-max-${uniqueId}` | Pre-delete read |
| 9 | `DELETE` | `DELETE` `200` | `.../locations/us/apps/cesapp-${uniqueId}/guardrails/`<br>`cesguardrail-max-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 9/9 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_cesguardrail-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
