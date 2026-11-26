# HTTP Lifecycle Report: `dialogflowcxgenerator-maximal`

- **Primary Resource:** `DialogflowCXGenerator` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxgenerator-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 4 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:26e43e257d1c626a`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"generator-max-${uniqueId}"`<br>**to:** `"generator-max-updated-${uniqueId}"` | Yes (#5 `PATCH`) | Yes |
| `spec.modelParameter.`<br>`.maxDecodeSteps` | **from:** `100`<br>**to:** `150` | Yes (#5 `PATCH`) | Yes |
| `spec.modelParameter.temperature` | **from:** `0.5`<br>**to:** `0.7` | Yes (#5 `PATCH`) | Yes |
| `spec.modelParameter.topK` | **from:** `20`<br>**to:** `30` | Yes (#5 `PATCH`) | Yes |
| `spec.modelParameter.topP` | **from:** `0.9`<br>**to:** `0.95` | Yes (#5 `PATCH`) | Yes |
| `spec.placeholders` | **from:** `[{"id": "topic_id", "name": "topic"}]`<br>**to:** `[{"id": "topic_id", "name": "topic"}, {"id": "topic_id_2", "n...` | Yes (#5 `PATCH`) | Yes |
| `spec.promptText.text` | **from:** `"Summarize $conversation for topic $topic"`<br>**to:** `"Summarize $conversation for topic $topic and subtopic $topic2"` | Yes (#5 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DialogflowCXAgent` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `POST` `200` | `.../locations/global/agents/71b71683-20a4-475f-b949-80a5e45a561f/`<br>`generators` | Created resource |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/71b71683-20a4-475f-b949-80a5e45a561f/generators/`<br>`617434c9-6885-4e52-bf3f-4e085ce4de44` | Pre-update read |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/71b71683-20a4-475f-b949-80a5e45a561f/generators/`<br>`617434c9-6885-4e52-bf3f-4e085ce4de44` | updateMask: displayName, modelParameter, placeholders, promptText |
| 6 | `DELETE` | `GET` `200` | `.../locations/global/agents/71b71683-20a4-475f-b949-80a5e45a561f/generators/`<br>`617434c9-6885-4e52-bf3f-4e085ce4de44` | Pre-delete read |
| 7 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/71b71683-20a4-475f-b949-80a5e45a561f/generators/`<br>`617434c9-6885-4e52-bf3f-4e085ce4de44` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 7/7 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxgenerator-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
