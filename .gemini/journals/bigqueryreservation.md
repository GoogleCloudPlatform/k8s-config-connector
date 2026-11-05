### [2026-10-19] BigQueryReservationReservationGroup Direct Controller
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for greenfield resource BigQueryReservationReservationGroup.
- **Observations**:
  - `ReservationGroup` in the GCP API (`google.cloud.bigquery.reservation.v1`) is an immutable container resource; the API only provides `CreateReservationGroup`, `GetReservationGroup`, `DeleteReservationGroup`, and `ListReservationGroups` methods (no `UpdateReservationGroup`).
  - The protobuf `ReservationGroup` message only contains the `Name` identifier field.
  - The controller handles `Update` by performing spec diff comparison and returning a clear error if changes are attempted, while successfully reconciling if unchanged.
  - Since there are no observed status fields in proto, `fuzztesting.NewKRMTypedSpecFuzzer` and `fuzztesting.RegisterKRMSpecFuzzer` are used for round-trip fuzz testing.

### [2026-10-24] BigQueryReservationReservationGroup MockGCP & Alignment
- **Context**: Implementing MockGCP service and alignment for `BigQueryReservationReservationGroup`.
- **Observations**:
  - Implemented `CreateReservationGroup`, `GetReservationGroup`, `DeleteReservationGroup`, and `ListReservationGroups` on `ReservationV1` in `mockgcp/mockbigqueryreservation/reservationgroup.go`.
  - Configured `BigQueryReservationReservationGroup` in `config/tests/samples/create/harness.go` under mock GCP supported resources.
  - Aligned error messaging on 404 Not Found to match live GCP: `Reservation group <fqn> not found.`.
  - Re-reconciliation and mock tests pass for both `bigqueryreservationreservationgroup-minimal` and `bigqueryreservationreservationgroup-maximal`.

