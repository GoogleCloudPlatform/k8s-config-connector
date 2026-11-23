# HTTP Lifecycle Report: `dataplexmetadatafeed-maximal`

- **Primary Resource:** `DataplexMetadataFeed` (`dataplex.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `metadatafeed-maximal-${uniqueId}`
- **Source Log:** `_http.log` (43 total HTTP entries, 34 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:30475fda36236e9a`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.filters.changeTypes` | **from:** `["CREATE", "UPDATE"]`<br>**to:** `["CREATE", "UPDATE", "DELETE"]` | Yes (#22 `PATCH`) | Yes |
| `spec.labels.label-a` | **from:** `"value-a"`<br>**to:** `"value-b"` | Yes (#22 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataplexAspectType` | 8 | `CREATE`, `DELETE` |
| `DataplexEntryGroup` | 7 | `CREATE`, `DELETE` |
| `DataplexEntryType` | 7 | `CREATE`, `DELETE` |
| `Project` | 1 | `UPDATE` |
| `PubSubTopic` | 11 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 17 | `CREATE` | `GET` `404` | `.../locations/us-central1/metadataFeeds/`<br>`metadatafeed-maximal-${uniqueId}` | Pre-create existence check |
| 18 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`metadataFeeds`<br>`?metadataFeedId=metadatafeed-maximal-${uniqueId}` | LRO: create -> done (#19)<br>Created resource |
| 20 | `CREATE` | `GET` `200` | `.../locations/us-central1/metadataFeeds/`<br>`metadatafeed-maximal-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 22 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/metadataFeeds/`<br>`metadatafeed-maximal-${uniqueId}` | LRO: update -> done (#23)<br>updateMask: filters, labels |
| 24 | `UPDATE` | `GET` `200` | `.../locations/us-central1/metadataFeeds/`<br>`metadatafeed-maximal-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 25 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/metadataFeeds/`<br>`metadatafeed-maximal-${uniqueId}` | LRO: delete -> done (#26)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataplexmetadatafeed-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
