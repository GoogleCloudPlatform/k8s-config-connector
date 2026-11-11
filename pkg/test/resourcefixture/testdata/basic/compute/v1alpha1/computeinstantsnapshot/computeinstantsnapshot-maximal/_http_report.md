# HTTP Lifecycle Report: `computeinstantsnapshot-maximal`

- **Primary Resource:** `ComputeInstantSnapshot` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computeinstantsnapshot-${uniqueId}`
- **Source Log:** `_http.log` (13 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:4a3867f85ee37c12`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeDisk` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../zones/us-central1-a/instantSnapshots/`<br>`${instantSnapshotID}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../zones/us-central1-a/`<br>`instantSnapshots` | LRO: insert -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GET` `200` | `.../zones/us-central1-a/instantSnapshots/`<br>`${instantSnapshotID}` | Post-create read (matches `create.yaml`) |
| 9 | `DELETE` | `DELETE` `200` | `.../zones/us-central1-a/instantSnapshots/`<br>`${instantSnapshotID}` | LRO: delete -> done (#10)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computeinstantsnapshot-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
