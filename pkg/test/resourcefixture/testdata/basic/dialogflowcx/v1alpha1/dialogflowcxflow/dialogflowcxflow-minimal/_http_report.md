# HTTP Lifecycle Report: `dialogflowcxflow-minimal`

- **Primary Resource:** `DialogflowCXFlow` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxflow-minimal-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:f044d61b9c19660a`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Minimal flow test"`<br>**to:** `"Updated minimal flow test"` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/`<br>`flows` | Created resource |
| 2 | `CREATE` | `GET` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`23a423c3-e060-469a-92ad-db344e601c4c` | Post-create read (matches `create.yaml`) |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`23a423c3-e060-469a-92ad-db344e601c4c` | updateMask: description |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`23a423c3-e060-469a-92ad-db344e601c4c` | Post-update read (matches `update.yaml`) |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/flows/`<br>`23a423c3-e060-469a-92ad-db344e601c4c` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxflow-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
