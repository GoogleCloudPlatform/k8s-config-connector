# HTTP Lifecycle Report: `discoveryengineconversation-maximal`

- **Primary Resource:** `DiscoveryEngineConversation` (`discoveryengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `conversation-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:deed229898bb3909`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.state` | **from:** `"IN_PROGRESS"`<br>**to:** `"COMPLETED"` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DiscoveryEngineDataStore` | 4 | `CREATE`, `DELETE` |
| `DiscoveryEngineEngine` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `500` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/conversations/`<br>`conversation-${uniqueId}` | Pre-create read |
| 6 | `CREATE` | `POST` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/`<br>`conversations` | Created resource |
| 7 | `UPDATE` | `GET` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/conversations/`<br>`conversation-${uniqueId}` | Pre-update read |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/conversations/`<br>`conversation-${uniqueId}` | updateMask: state |
| 9 | `DELETE` | `GET` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/conversations/`<br>`conversation-${uniqueId}` | Pre-delete read |
| 10 | `DELETE` | `DELETE` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/conversations/`<br>`conversation-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_discoveryengineconversation-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
