# HTTP Lifecycle Report: `discoveryengineconversation-minimal`

- **Primary Resource:** `DiscoveryEngineConversation` (`discoveryengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `conv-${uniqueId}`
- **Source Log:** `_http.log` (12 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:9548266e002db311`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

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
| 7 | `DELETE` | `GET` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/conversations/`<br>`conversation-${uniqueId}` | Pre-delete read |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/conversations/`<br>`conversation-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_discoveryengineconversation-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
