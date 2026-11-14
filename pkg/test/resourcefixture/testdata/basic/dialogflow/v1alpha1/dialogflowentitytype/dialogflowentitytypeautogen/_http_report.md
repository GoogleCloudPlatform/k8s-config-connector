# HTTP Lifecycle Report: `dialogflowentitytypeautogen`

- **Primary Resource:** `DialogflowEntityType` (`dialogflow.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowentitytype-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 5 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:0619453ea27f72c6`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DialogflowAgent` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `POST` `200` | `.../agent/entityTypes/` | Created resource |
| 5 | `CREATE` | `GET` `200` | `.../agent/entityTypes/${entityTypeID}` | Post-create read (matches `create.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../agent/entityTypes/${entityTypeID}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowentitytypeautogen.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
