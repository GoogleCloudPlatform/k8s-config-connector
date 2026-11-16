# HTTP Lifecycle Report: `ccinsightsphrasematcher-maximal`

- **Primary Resource:** `CCInsightsPhraseMatcher` (`contactcenterinsights.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `pmmax${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:6e80f3faf61f0bdc`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.active` | **from:** `true`<br>**to:** `false` | Yes (#4 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal Phrase Matcher ${uniqueId}"`<br>**to:** `"Updated Maximal Phrase Matcher ${uniqueId}"` | Yes (#4 `PATCH`) | Yes |
| `spec.phraseMatchRuleGroups` | **from:** `[{"phraseMatchRules": [{"config": {"exactMatchConfig": {"case...`<br>**to:** `[{"phraseMatchRules": [{"config": {"exactMatchConfig": {"case...` | Yes (#4 `PATCH`) | Yes |
| `spec.roleMatch` | **from:** `"END_USER"`<br>**to:** `"HUMAN_AGENT"` | Yes (#4 `PATCH`) | Yes |
| `spec.type` | **from:** `"ALL_OF"`<br>**to:** `"ANY_OF"` | Yes (#4 `PATCH`) | Yes |
| `spec.versionTag` | **from:** `"v1"`<br>**to:** `"v2"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `400` | `.../locations/us-central1/phraseMatchers/`<br>`pmmax${uniqueId}` | Pre-create read |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`phraseMatchers` | Created resource |
| 3 | `CREATE` | `GET` `200` | `.../locations/us-central1/phraseMatchers/`<br>`${phraseMatcherID}` | Post-create read (matches `create.yaml`) |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/phraseMatchers/`<br>`${phraseMatcherID}` | updateMask: active, displayName, phraseMatchRuleGroups, roleMatch, type, versionTag |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-central1/phraseMatchers/`<br>`${phraseMatcherID}` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/phraseMatchers/`<br>`${phraseMatcherID}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 6/6 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_ccinsightsphrasematcher-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
