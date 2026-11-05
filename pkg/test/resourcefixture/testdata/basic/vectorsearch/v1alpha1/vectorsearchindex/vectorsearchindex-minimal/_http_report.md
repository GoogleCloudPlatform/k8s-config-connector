# HTTP Lifecycle Report: `vectorsearchindex-minimal`

- **Primary Resource:** `VectorsearchIndex` (`vectorsearch.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `vsi-min-${uniqueId}`
- **Source Log:** `_http.log` (16 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e7692162a350f70e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** *(unset)*<br>**to:** `"Updated Minimal Vectorsearch Index"` | Yes (#9 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `VectorSearchCollection` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/collections/vsc-min-${uniqueId}/indexes/`<br>`vsi-min-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/collections/vsc-min-${uniqueId}/`<br>`indexes`<br>`?indexId=vsi-min-${uniqueId}` | LRO: create -> done (#7)<br>Created resource |
| 8 | `UPDATE` | `GET` `200` | `.../locations/us-central1/collections/vsc-min-${uniqueId}/indexes/`<br>`vsi-min-${uniqueId}` | Pre-update read |
| 9 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/collections/vsc-min-${uniqueId}/indexes/`<br>`vsi-min-${uniqueId}` | LRO: update -> done (#10)<br>updateMask: displayName |
| 11 | `DELETE` | `GET` `200` | `.../locations/us-central1/collections/vsc-min-${uniqueId}/indexes/`<br>`vsi-min-${uniqueId}` | Pre-delete read |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/collections/vsc-min-${uniqueId}/indexes/`<br>`vsi-min-${uniqueId}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_vectorsearchindex-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
