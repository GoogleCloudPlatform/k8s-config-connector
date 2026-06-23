# BigQuery Reservation Journal

### [2026-06-02] Handling Missing Proto Definitions for New Greenfield Resources
- **Context**: Implementing direct types, CRD, and IdentityV2 for `BigQueryReservationReservationGroup` which is a brand-new GCP resource.
- **Problem**: The pinned Google APIs SHA in `apis/git.versions` (`731d7f2ab6`) did not contain the proto definition for `ReservationGroup` in `google/cloud/bigquery/reservation/v1/reservation.proto`. Attempting to bump `git.versions` to the latest master SHA caused massive, unrelated out-of-date code regeneration across dozens of other KCC direct packages when running `make generate` and `validate-prereqs.sh`.
- **Solution**: 
  1. Vendored the updated `reservation.proto` (containing `ReservationGroup` and its RPCs) directly into `mockgcp/apis/google/cloud/bigquery/reservation/v1/reservation.proto`.
  2. Updated `dev/tools/controllerbuilder/generate-proto.sh` to include this vendored file and reordered the `protoc` include paths so that `-I ${REPO_ROOT}/mockgcp/apis` is placed before `-I ${THIRD_PARTY}/googleapis/`.
  3. Added an automatic cleanup step right before the `protoc` command to delete the shadowed file `rm -f ${THIRD_PARTY}/googleapis/google/cloud/bigquery/reservation/v1/reservation.proto` to avoid "Input is shadowed" compilation errors.
- **Impact**: Future developers working on greenfield resources whose protos are missing from the pinned SHA can easily vendor the updated `.proto` files into `mockgcp/apis/google/cloud/...` and update `generate-proto.sh` without polluting other packages by bumping the global `git.versions` SHA.
