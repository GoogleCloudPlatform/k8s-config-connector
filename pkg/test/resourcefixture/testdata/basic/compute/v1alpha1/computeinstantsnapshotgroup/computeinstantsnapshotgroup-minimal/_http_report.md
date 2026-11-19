# HTTP Lifecycle Report: `computeinstantsnapshotgroup-minimal`

- **Primary Resource:** `ComputeInstantSnapshotGroup` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computeinstantsnapshotgroup-${uniqueId}`
- **Source Log:** `_http.log` (20 total HTTP entries, 14 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:f6f21712feb41231`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeDisk` | 7 | `CREATE`, `DELETE` |
| `ComputeResourcePolicy` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 9 | `CREATE` | `GET` `404` | `.../zones/us-central1-a/instantSnapshotGroups/`<br>`${instantSnapshotGroupID}` | Pre-create existence check |
| 10 | `CREATE` | `POST` `200` | `.../zones/us-central1-a/`<br>`instantSnapshotGroups` | LRO: insert -> done (#11)<br>Created resource |
| 12 | `CREATE` | `GET` `200` | `.../zones/us-central1-a/instantSnapshotGroups/`<br>`${instantSnapshotGroupID}` | Post-create read (matches `create.yaml`) |
| 13 | `DELETE` | `DELETE` `200` | `.../zones/us-central1-a/instantSnapshotGroups/`<br>`${instantSnapshotGroupID}` | LRO: delete -> done (#14)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computeinstantsnapshotgroup-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
