# HTTP Lifecycle Report: `computebackendbucketsignedurlkeybasic`

- **Primary Resource:** `ComputeBackendBucketSignedURLKey` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computebackendbucketsignedurlkey-${uniqueId} (resourceID: key-${uniqueId})`
- **Source Log:** `_http.log` (18 total HTTP entries, 14 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:042543d450837919`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeBackendBucket` | 6 | `CREATE`, `DELETE` |
| `StorageBucket` | 8 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 8 | `CREATE` | `POST` `200` | `.../global/backendBuckets/${backendBucketID}/`<br>`addSignedUrlKey` | LRO: compute.backendBuckets.addSignedUrlKey -> done (#9)<br>Created resource |
| 11 | `DELETE` | `POST` `200` | `.../global/backendBuckets/${backendBucketID}/`<br>`deleteSignedUrlKey`<br>`?keyName=key-${uniqueId}` | LRO: compute.backendBuckets.deleteSignedUrlKey -> done (#12)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computebackendbucketsignedurlkeybasic.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
