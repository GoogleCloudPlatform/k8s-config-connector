# HTTP Lifecycle Report: `bigqueryreservationassignment-orgassignment`

- **Primary Resource:** `BigQueryReservationAssignment` (`bigqueryreservation.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `bqassignment-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 4 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:5aee2559bc8a820c`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigQueryReservationReservation` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `200` | `.../locations/us-west2/reservations/bigqueryreservationreservation${uniqueId}/`<br>`assignments` | Pre-create read |
| 4 | `CREATE` | `POST` `200` | `.../locations/us-west2/reservations/bigqueryreservationreservation${uniqueId}/`<br>`assignments`<br>`?assignmentId=bqassignment-${uniqueId}` | Created resource |
| 5 | `DELETE` | `GET` `200` | `.../locations/us-west2/reservations/bigqueryreservationreservation${uniqueId}/`<br>`assignments` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-west2/reservations/bigqueryreservationreservation${uniqueId}/assignments/`<br>`bqassignment-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_bigqueryreservationassignment-orgassignment.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
