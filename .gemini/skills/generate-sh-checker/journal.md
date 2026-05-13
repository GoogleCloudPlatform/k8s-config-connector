# Journal for generate-sh-checker

## 2026-05-13: compute/v1alpha1 backfill

- Encountered an issue where `FutureReservation` message is missing from `google.cloud.compute.v1` in `googleapis`, even though it exists in the Go SDK `apiv1/computepb`.
- Handled this by moving `FutureReservation` and its non-shared dependent types from `types.generated.go` to `futurereservation_types.go` (since they are hand-written/preserved anyway) and only running `generate-types` for the other 3 resources that are present in `v1`.
- This ensures `types.generated.go` only contains types reachable from the resources we explicitly manage with `generate-types`.
- Noted that `compute` is special because of its size and the way protos are managed.
- Verified that `generate-mapper` from `v1beta1/generate.sh` correctly includes `v1alpha1` resources when using `--multiversion`.
