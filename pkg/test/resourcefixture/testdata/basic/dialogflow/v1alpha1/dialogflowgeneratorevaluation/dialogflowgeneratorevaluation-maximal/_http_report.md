# HTTP Lifecycle Report: `dialogflowgeneratorevaluation-maximal`

- **Primary Resource:** `DialogflowGeneratorEvaluation` (`dialogflow.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dfgeneratoreval-max-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:c9f28a4d34bd0667`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DialogflowGenerator` | 4 | `CREATE`, `DELETE` |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 6 | `CREATE` | `POST` `200` | `.../locations/global/generators/dialogflowgenerator-${uniqueId}/`<br>`evaluations` | LRO: create -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GET` `200` | `.../locations/global/generators/dialogflowgenerator-${uniqueId}/evaluations/`<br>`MTI2MTg0NzU5MjY5Mzg3MTQxMTM` | Post-create read (matches `create.yaml`) |
| 9 | `DELETE` | `DELETE` `200` | `.../locations/global/generators/dialogflowgenerator-${uniqueId}/evaluations/`<br>`MTI2MTg0NzU5MjY5Mzg3MTQxMTM` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowgeneratorevaluation-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
