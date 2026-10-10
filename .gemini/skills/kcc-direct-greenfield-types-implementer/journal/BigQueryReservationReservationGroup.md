# BigQueryReservationReservationGroup Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema & Proto Mapping**:
   - `ReservationGroup` does not currently exist as a message in the compiled protobuf descriptor (`google.cloud.bigquery.reservation.v1`), but exists in the GCP REST API under `projects.locations.reservationGroups` and in CAIS metadata (`//bigqueryreservation.googleapis.com/projects/{{PROJECT_NUMBER}}/locations/{{LOCATION}}/reservationGroups/{{RESERVATION_GROUP}}`).
   - Defined types manually under `apis/bigqueryreservation/v1alpha1/bigqueryreservationreservationgroup_types.go` adhering to KCC standards:
     - `ProjectRef` (`*refsv1beta1.ProjectRef`) and `Location` (`*string`) marked as required.
     - `ResourceID` (`*string`) marked as optional.
     - `Status` contains `Conditions`, `ObservedGeneration` (`*int64`), `ExternalRef` (`*string`), and `ObservedState` (`CreationTime` and `UpdateTime`).
     - Stability level set to `alpha`.

2. **Identity & Reference**:
   - Implemented identity under `apis/bigqueryreservation/v1alpha1/bigqueryreservationreservationgroup_identity.go`.
   - Used `gcpurls.Template[BigQueryReservationReservationGroupIdentity]("bigqueryreservation.googleapis.com", "projects/{project}/locations/{location}/reservationGroups/{reservationGroup}")`.
   - Implemented standard `ParentString()`, `GetIdentity(ctx, reader)` with authoritative drift detection cross-checking against `status.externalRef`.
   - Implemented `BigQueryReservationReservationGroupRef` under `apis/bigqueryreservation/v1alpha1/bigqueryreservationreservationgroup_reference.go` delegating `Normalize` to `refs.Normalize`.

3. **Validation & Tests**:
   - Implemented comprehensive unit tests under `apis/bigqueryreservation/v1alpha1/bigqueryreservationreservationgroup_identity_test.go`.
   - Validated `scripts/validate-prereqs.sh` and CAIS tests pass cleanly.
