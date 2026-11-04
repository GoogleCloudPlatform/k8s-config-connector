# HTTP Lifecycle Report: `configdeliveryrelease-maximal`

- **Primary Resource:** `ConfigDeliveryRelease` (`configdelivery.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `v1-0-0`
- **Source Log:** `_http.log` (12 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:09edb9a94e70ad52`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `.spec` or `.metadata.labels` differences between `create.yaml` and `update.yaml`._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ConfigDeliveryResourceBundle` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/releases/`<br>`v1-0-0` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/`<br>`releases`<br>`?releaseId=v1-0-0` | LRO: create -> done (#6)<br>Created resource |
| 7 | `DELETE` | `GET` `200` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/releases/`<br>`v1-0-0` | Pre-delete read |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/releases/`<br>`v1-0-0` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `.spec` or `.metadata.labels` changes between `create.yaml` and `update.yaml`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_configdeliveryrelease-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
