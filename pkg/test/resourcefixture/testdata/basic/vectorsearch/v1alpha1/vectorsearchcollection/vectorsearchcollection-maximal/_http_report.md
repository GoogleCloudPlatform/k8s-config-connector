# HTTP Lifecycle Report: `vectorsearchcollection-maximal`

- **Primary Resource:** `VectorSearchCollection` (`vectorsearch.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `vsc-max-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:562691e69bc506f9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.dataSchema.properties.year.`<br>`.type` | **from:** *(unset)*<br>**to:** `"integer"` | Yes (#5 `PATCH`) | Yes |
| `spec.description` | **from:** `"Initial description of maximal collection"`<br>**to:** `"Updated description of maximal collection"` | Yes (#5 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal Vector Search Collection"`<br>**to:** `"Maximal Vector Search Collection - Updated"` | Yes (#5 `PATCH`) | Yes |
| `spec.labels.env` | **from:** `"test"`<br>**to:** `"prod"` | Yes (#5 `PATCH`) | Yes |
| `spec.labels.owner` | **from:** `"kcc"`<br>**to:** `"kcc-team"` | Yes (#5 `PATCH`) | Yes |
| `spec.labels.updated` | **from:** *(unset)*<br>**to:** `"true"` | Yes (#5 `PATCH`) | Yes |
| `spec.vectorSchema.`<br>`.my_vector_field.denseVector.`<br>`.vertexEmbeddingConfig.taskType` | **from:** `"RETRIEVAL_DOCUMENT"`<br>**to:** `"RETRIEVAL_QUERY"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/collections/`<br>`vsc-max-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/collections`<br>`?collectionId=vsc-max-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `UPDATE` | `GET` `200` | `.../locations/us-central1/collections/`<br>`vsc-max-${uniqueId}` | Pre-update read |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/collections/`<br>`vsc-max-${uniqueId}` | LRO: update -> done (#6)<br>updateMask: dataSchema, description, displayName, labels, vectorSchema |
| 7 | `DELETE` | `GET` `200` | `.../locations/us-central1/collections/`<br>`vsc-max-${uniqueId}` | Pre-delete read |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/collections/`<br>`vsc-max-${uniqueId}`<br>`?force=true` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 7/7 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_vectorsearchcollection-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
