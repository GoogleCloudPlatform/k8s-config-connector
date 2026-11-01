# HTTP Lifecycle Report: `vectorsearchindex-maximal`

- **Primary Resource:** `VectorsearchIndex` (`vectorsearch.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `vsi-max-${uniqueId}`
- **Source Log:** `_http.log` (17 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:07659c679b41764e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.dedicatedInfrastructure.`<br>`.autoscalingSpec.maxReplicaCount` | **from:** `5`<br>**to:** `6` | Yes (#9 `PATCH`) | Yes |
| `spec.dedicatedInfrastructure.`<br>`.autoscalingSpec.minReplicaCount` | **from:** `2`<br>**to:** `3` | Yes (#9 `PATCH`) | Yes |
| `spec.description` | **from:** `"Initial description of maximal index"`<br>**to:** `"Updated description of maximal index"` | Yes (#9 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal Vector Search Index"`<br>**to:** `"Maximal Vector Search Index - Updated"` | Yes (#9 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `VectorSearchCollection` | 8 | `CREATE`, `UPDATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/collections/vsc-max-dep-${uniqueId}/indexes/`<br>`vsi-max-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/collections/vsc-max-dep-${uniqueId}/`<br>`indexes`<br>`?indexId=vsi-max-${uniqueId}` | LRO: create -> done (#7)<br>Created resource |
| 8 | `UPDATE` | `GET` `200` | `.../locations/us-central1/collections/vsc-max-dep-${uniqueId}/indexes/`<br>`vsi-max-${uniqueId}` | Pre-update read |
| 9 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/collections/vsc-max-dep-${uniqueId}/indexes/`<br>`vsi-max-${uniqueId}` | LRO: update -> done (#11)<br>updateMask: dedicatedInfrastructure, description, displayName |
| 12 | `DELETE` | `GET` `200` | `.../locations/us-central1/collections/vsc-max-dep-${uniqueId}/indexes/`<br>`vsi-max-${uniqueId}` | Pre-delete read |
| 13 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/collections/vsc-max-dep-${uniqueId}/indexes/`<br>`vsi-max-${uniqueId}` | LRO: delete -> done (#14)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_vectorsearchindex-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
