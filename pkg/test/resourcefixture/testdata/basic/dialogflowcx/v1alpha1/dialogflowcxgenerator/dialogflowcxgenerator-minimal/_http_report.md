# HTTP Lifecycle Report: `dialogflowcxgenerator-minimal`

- **Primary Resource:** `DialogflowCXGenerator` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxgenerator-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 4 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:f905e03db016bf1b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"generator-${uniqueId}"`<br>**to:** `"generator-updated-${uniqueId}"` | Yes (#5 `PATCH`) | Yes |
| `spec.promptText.text` | **from:** `"Summarize the customer conversation: $conversation"`<br>**to:** `"Summarize the customer query and intent: $conversation"` | Yes (#5 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DialogflowCXAgent` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `POST` `200` | `.../locations/global/agents/3b72e995-aa60-428f-84a4-cd420c50d6e3/`<br>`generators` | Created resource |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/3b72e995-aa60-428f-84a4-cd420c50d6e3/generators/`<br>`93af9224-55e2-4529-9651-33fad7ce4ada` | Pre-update read |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/3b72e995-aa60-428f-84a4-cd420c50d6e3/generators/`<br>`93af9224-55e2-4529-9651-33fad7ce4ada` | updateMask: displayName, promptText |
| 6 | `DELETE` | `GET` `200` | `.../locations/global/agents/3b72e995-aa60-428f-84a4-cd420c50d6e3/generators/`<br>`93af9224-55e2-4529-9651-33fad7ce4ada` | Pre-delete read |
| 7 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/3b72e995-aa60-428f-84a4-cd420c50d6e3/generators/`<br>`93af9224-55e2-4529-9651-33fad7ce4ada` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxgenerator-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
