# HTTP Lifecycle Report: `bigqueryreservationschedulingpolicy`

- **Primary Resource:** `BigQueryReservationReservation` (`bigqueryreservation.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `bigqueryreservationreservation-${uniqueId} (resourceID: bigqueryreservationreservation${uniqueId})`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:f718b9b99d3fc5e8`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.schedulingPolicy.`<br>`.concurrency` | **from:** `5`<br>**to:** `10` | Yes (#4 `PATCH`) | Yes |
| `spec.schedulingPolicy.maxSlots` | **from:** `100`<br>**to:** `200` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-west2/reservations/`<br>`bigqueryreservationreservation${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-west2/reservations`<br>`?reservationId=bigqueryreservationreservation${uniqueId}` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/us-west2/reservations/`<br>`bigqueryreservationreservation${uniqueId}` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-west2/reservations/`<br>`bigqueryreservationreservation${uniqueId}` | updateMask: schedulingPolicy |
| 5 | `DELETE` | `GET` `200` | `.../locations/us-west2/reservations/`<br>`bigqueryreservationreservation${uniqueId}` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-west2/reservations/`<br>`bigqueryreservationreservation${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_bigqueryreservationschedulingpolicy.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [ ] **V7 (Authenticity & Hand-Edit Integrity):** Protobuf schema violation: stanza #2 (`POST`) request body failed protobuf schema `google.cloud.bigquery.reservation.v1.Reservation`: proto: (line 1:68): unknown field "schedulingPolicy".
  - `<hint_for_agent>`: HAND-EDITED OR HALLUCINATED PROTO PAYLOAD DETECTED: One or more JSON payloads or @type URLs in _http.log violate the protobuf schema in .build/googleapis.pb. Do not hand-edit _http.log; re-record against RealGCP.
