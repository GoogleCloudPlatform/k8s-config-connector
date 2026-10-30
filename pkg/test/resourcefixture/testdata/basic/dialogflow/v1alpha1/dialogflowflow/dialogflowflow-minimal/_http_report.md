# HTTP Lifecycle Report: `dialogflowflow-minimal`

- **Primary Resource:** `DialogflowFlow` (`dialogflow.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dfflow-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2fdf75b6488614fb`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"minimal dialogflow cx flow"`<br>**to:** `"updated dialogflow cx flow"` | Yes (#3 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Flow ${uniqueId}"`<br>**to:** `"Flow Updated ${uniqueId}"` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/agents/b674ae89-59b4-4587-ac8d-71b203f68dbd/`<br>`flows` | Created resource |
| 2 | `CREATE` | `GET` `200` | `.../locations/global/agents/b674ae89-59b4-4587-ac8d-71b203f68dbd/flows/`<br>`11dcd49c-0880-4b78-b75b-0c85f46c9023` | Post-create read (matches `create.yaml`) |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/b674ae89-59b4-4587-ac8d-71b203f68dbd/flows/`<br>`11dcd49c-0880-4b78-b75b-0c85f46c9023` | updateMask: description, displayName |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/b674ae89-59b4-4587-ac8d-71b203f68dbd/flows/`<br>`11dcd49c-0880-4b78-b75b-0c85f46c9023` | Post-update read (matches `update.yaml`) |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/b674ae89-59b4-4587-ac8d-71b203f68dbd/flows/`<br>`11dcd49c-0880-4b78-b75b-0c85f46c9023` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowflow-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
