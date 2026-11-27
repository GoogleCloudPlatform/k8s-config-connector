# HTTP Lifecycle Report: `dialogflowcxtool-minimal`

- **Primary Resource:** `DialogflowCXTool` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxtool-minimal-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 3 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:f146a9f00cf7a815`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Minimal tool test"`<br>**to:** `"Minimal tool test updated"` | Yes (#4 `PATCH`) | Yes |
| `spec.functionSpec.inputSchema.`<br>`.properties.param2.type` | **from:** *(unset)*<br>**to:** `"string"` | Yes (#4 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DialogflowCXAgent` | 3 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/agents/aa771d89-00c7-4fe8-b1e6-7947d4b91660/`<br>`tools` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/global/agents/aa771d89-00c7-4fe8-b1e6-7947d4b91660/tools/`<br>`790fa71f-8288-4f5a-9f91-06fe977ab22c` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/aa771d89-00c7-4fe8-b1e6-7947d4b91660/tools/`<br>`790fa71f-8288-4f5a-9f91-06fe977ab22c` | updateMask: description, functionSpec |
| 5 | `DELETE` | `GET` `200` | `.../locations/global/agents/aa771d89-00c7-4fe8-b1e6-7947d4b91660/tools/`<br>`790fa71f-8288-4f5a-9f91-06fe977ab22c` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/aa771d89-00c7-4fe8-b1e6-7947d4b91660/tools/`<br>`790fa71f-8288-4f5a-9f91-06fe977ab22c` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxtool-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
