# HTTP Lifecycle Report: `dialogflowcxtool-maximal`

- **Primary Resource:** `DialogflowCXTool` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxtool-maximal-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 3 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:34e7bfd7152586dc`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal tool test"`<br>**to:** `"Maximal tool test updated"` | Yes (#4 `PATCH`) | Yes |
| `spec.openAPISpec.textSchema` | **from:** `"openapi: 3.0.0\ninfo:\n  title: Sample API\n  version: 1.0.0...`<br>**to:** `"openapi: 3.0.0\ninfo:\n  title: Sample API Updated\n  versio...` | Yes (#4 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DialogflowCXAgent` | 3 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/agents/5d1cda79-5a98-47b9-8d8e-7b21e4872dd3/`<br>`tools` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/global/agents/5d1cda79-5a98-47b9-8d8e-7b21e4872dd3/tools/`<br>`483eb7b5-3a7e-4643-9b01-33dd34ff9b77` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/5d1cda79-5a98-47b9-8d8e-7b21e4872dd3/tools/`<br>`483eb7b5-3a7e-4643-9b01-33dd34ff9b77` | updateMask: description, openApiSpec |
| 5 | `DELETE` | `GET` `200` | `.../locations/global/agents/5d1cda79-5a98-47b9-8d8e-7b21e4872dd3/tools/`<br>`483eb7b5-3a7e-4643-9b01-33dd34ff9b77` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/5d1cda79-5a98-47b9-8d8e-7b21e4872dd3/tools/`<br>`483eb7b5-3a7e-4643-9b01-33dd34ff9b77` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxtool-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
