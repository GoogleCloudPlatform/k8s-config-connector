# HTTP Lifecycle Report: `dataplexpartition-minimal`

- **Primary Resource:** `DataplexPartition` (`dataplex.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dataplexpartition-minimal-${uniqueId}`
- **Source Log:** `_http.log` (32 total HTTP entries, 29 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:67dfaed12556f93e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataplexAsset` | 6 | `CREATE`, `DELETE` |
| `DataplexEntity` | 3 | `CREATE`, `DELETE` |
| `DataplexLake` | 7 | `CREATE`, `DELETE` |
| `DataplexZone` | 7 | `CREATE`, `DELETE` |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 14 | `CREATE` | `POST` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/entities/entity_${uniqueId}/`<br>`partitions` | Created resource |
| 15 | `DELETE` | `GET` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/entities/entity_${uniqueId}/partitions/`<br>`2026-11-16` | Pre-delete read |
| 16 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/entities/entity_${uniqueId}/partitions/`<br>`2026-11-16` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataplexpartition-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
