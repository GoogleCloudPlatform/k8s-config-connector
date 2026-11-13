# HTTP Lifecycle Report: `dialogflowcxflow-maximal`

- **Primary Resource:** `DialogflowCXFlow` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxflow-maximal-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:8e981033e23495e9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal flow test"`<br>**to:** `"Updated maximal flow test"` | Yes (#3 `PATCH`) | Yes |
| `spec.nluSettings.`<br>`.classificationThreshold` | **from:** `0.5`<br>**to:** `0.6` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/`<br>`flows` | Created resource |
| 2 | `CREATE` | `GET` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`5914a4bc-506a-42fa-9122-5301e29cb157` | Post-create read (matches `create.yaml`) |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`5914a4bc-506a-42fa-9122-5301e29cb157` | updateMask: description, nluSettings |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`5914a4bc-506a-42fa-9122-5301e29cb157` | Post-update read (matches `update.yaml`) |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`5914a4bc-506a-42fa-9122-5301e29cb157` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxflow-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
