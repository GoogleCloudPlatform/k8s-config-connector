# BigQueryReservationReservationGroup Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Proto Pinning in `generate.sh`**:
   - The global protobuf version pinned in `apis/git.versions` (`1765b559c42386788ff0c6412491277b4791107a`) did not yet include `message ReservationGroup`.
   - Pinned `PROTO_SHA="2193a2bfcecb92b92aad7a4d81baa428cafd7dfd"` (which introduces `ReservationGroup` in `google/cloud/bigquery/reservation/v1/reservation.proto`) in `apis/bigqueryreservation/generate.sh`.
   - Passed `--proto-source-path ${PROTO_OUT}` to controllerbuilder commands so the build is hermetic and does not disturb the global `apis/git.versions`.

2. **Types & Schema Design**:
   - Added `cnrm.cloud.google.com/stability-level: alpha` metadata label to the CRD schema.
   - Upstream `google.cloud.bigquery.reservation.v1.ReservationGroup` only defines `string name = 1;` (which is an identifier).
   - In accordance with KCC guidelines for resources with no output-only fields, commented out `ObservedState *BigQueryReservationReservationGroupObservedState` in Status and commented out `BigQueryReservationReservationGroupObservedState` type definition so an empty schema is not generated under `status.observedState`.
   - Added `Location *string` with immutability rule `self == oldSelf` and `ProjectRef *refsv1beta1.ProjectRef` under Spec.
   - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` with golden output for missing spec fields in unstructured tests.

3. **Identity & URL Templates**:
   - Implemented `BigQueryReservationReservationGroupIdentity` with template `bigqueryreservation.googleapis.com` / `projects/{project}/locations/{location}/reservationGroups/{reservationGroup}` matching Cloud Asset Inventory format in `docs/ai/metadata/cloudassetinventory_names.jsonl`.
   - Implemented canonical `_reference.go` delegating `Normalize` to `refs.Normalize`.
