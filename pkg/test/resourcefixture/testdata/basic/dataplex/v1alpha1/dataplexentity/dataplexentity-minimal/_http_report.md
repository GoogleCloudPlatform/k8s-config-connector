# HTTP Lifecycle Report: `dataplexentity-minimal`

- **Primary Resource:** `DataplexEntity` (`dataplex.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dataplexentity-minimal-${uniqueId}`
- **Source Log:** `_http.log` (31 total HTTP entries, 26 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e65d4eb98f78825e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** *(unset)*<br>**to:** `"Updated description"` | Yes (#15 `PUT`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataplexAsset` | 6 | `CREATE`, `DELETE` |
| `DataplexLake` | 7 | `CREATE`, `DELETE` |
| `DataplexZone` | 7 | `CREATE`, `DELETE` |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 13 | `CREATE` | `POST` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/`<br>`entities` | Created resource |
| 14 | `UPDATE` | `GET` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/entities/`<br>`entity_${uniqueId}`<br>`?view=FULL` | Pre-update read |
| 15 | `UPDATE` | `PUT` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/entities/`<br>`entity_${uniqueId}` | Updated resource (`PUT`) |
| 16 | `DELETE` | `GET` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/entities/`<br>`entity_${uniqueId}`<br>`?view=FULL` | Pre-delete read |
| 17 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/lakes/dataplexlake-${uniqueId}/zones/dataplexzone-${uniqueId}/entities/`<br>`entity_${uniqueId}`<br>`?etag=abcdef0123A` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PUT`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataplexentity-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
