### [2026-07-06] BigQueryReservationReservationGroup Direct Types Implementation
- **Context**: [BigQueryReservationReservationGroup KRM Types and IdentityV2](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9022)
- **Problem**: The pinned Google APIs commit `731d7f2ab6e4e2ea15030c95039e2cb66174d4fb` does not contain the `ReservationGroup` proto message definition under `google/cloud/bigquery/reservation/v1/reservation.proto`.
- **Solution**: We updated `apis/git.versions` to point to a newer Google APIs master commit `2b625c91510a2e8320a778bc88af8b65bc4a19a2` which contains `ReservationGroup`, and cleared cached `.pb` files under `.build/` to force `generate-proto.sh` to run `protoc` again and regenerate the API definitions.
- **Impact**: The type generator successfully compiled the new proto definition, scaffolded `bigqueryreservationreservationgroup_types.go`, and generated CRD schemas.
