# HTTP Lifecycle Report: `vertexaifeaturestoreentitytype-basic`

- **Primary Resource:** `VertexAIFeaturestoreEntityType` (`vertexai.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `entitytype${uniqueId}`
- **Source Log:** `_http.log` (12 total HTTP entries, 5 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e7bedb63f88a7329`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Test entity type"`<br>**to:** `"Updated test entity type"` | Yes (#8 `PATCH`) | Yes |
| `spec.offlineStorageTtlDays` | **from:** `30`<br>**to:** `60` | Yes (#8 `PATCH`) | Yes |
| `metadata.labels.label-one` | **from:** `"value-one"`<br>**to:** `"value-two"` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `VertexAIFeaturestore` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/featurestores/featurestore${uniqueId}/entityTypes/`<br>`entitytype${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/featurestores/featurestore${uniqueId}/`<br>`entityTypes`<br>`?entityTypeId=entitytype${uniqueId}` | LRO: create -> done (#6)<br>Created resource |
| 7 | `CREATE` | `GET` `200` | `.../locations/us-central1/featurestores/featurestore${uniqueId}/entityTypes/`<br>`entitytype${uniqueId}` | Post-create read (matches `create.yaml`) |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/featurestores/featurestore${uniqueId}/entityTypes/`<br>`entitytype${uniqueId}` | updateMask: description, offlineStorageTtlDays, labels |
| 9 | `DELETE` | `GET` `200` | `.../locations/us-central1/featurestores/featurestore${uniqueId}/entityTypes/`<br>`entitytype${uniqueId}` | Pre-delete read |
| 10 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/featurestores/featurestore${uniqueId}/entityTypes/`<br>`entitytype${uniqueId}` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_vertexaifeaturestoreentitytype-basic.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
