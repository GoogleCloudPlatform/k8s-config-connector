### [2026-10-19] BigQueryReservationReservationGroup Direct Controller
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for greenfield resource BigQueryReservationReservationGroup.
- **Observations**:
  - `ReservationGroup` in the GCP API (`google.cloud.bigquery.reservation.v1`) is an immutable container resource; the API only provides `CreateReservationGroup`, `GetReservationGroup`, `DeleteReservationGroup`, and `ListReservationGroups` methods (no `UpdateReservationGroup`).
  - The protobuf `ReservationGroup` message only contains the `Name` identifier field.
  - The controller handles `Update` by performing spec diff comparison and returning a clear error if changes are attempted, while successfully reconciling if unchanged.
  - Since there are no observed status fields in proto, `fuzztesting.NewKRMTypedSpecFuzzer` and `fuzztesting.RegisterKRMSpecFuzzer` are used for round-trip fuzz testing.
