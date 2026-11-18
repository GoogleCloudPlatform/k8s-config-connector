# HTTP Lifecycle Report: `dialogflowcxentitytype-minimal`

- **Primary Resource:** `DialogflowCXEntityType` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxentitytype-minimal-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:db1ffdecd188b13d`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.entities` | **from:** `[{"synonyms": ["synonym1"], "value": "value1"}]`<br>**to:** `[{"synonyms": ["synonym1", "synonym2"], "value": "value1"}]` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/`<br>`entityTypes` | Created resource |
| 2 | `CREATE` | `GET` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/entityTypes/`<br>`13c4c465-0c78-4834-9f62-ff93d1a73f8f` | Post-create read (matches `create.yaml`) |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/entityTypes/`<br>`13c4c465-0c78-4834-9f62-ff93d1a73f8f` | updateMask: entities |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/entityTypes/`<br>`13c4c465-0c78-4834-9f62-ff93d1a73f8f` | Post-update read (matches `update.yaml`) |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/4888fac5-5f1c-401e-a5e2-fe19d1f0bebf/entityTypes/`<br>`13c4c465-0c78-4834-9f62-ff93d1a73f8f` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxentitytype-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
