# Journal - Update GCP Dependencies

## 2026-05-13
- Ran `dev/tasks/update-gcp-dependencies`.
- Found a bug where the script wrote `git.versions` to the root instead of `apis/git.versions`. Fixed it.
- Encountered a compilation error in `networkmanagement` v1alpha1. The API update added `edge_responses` and `probed_all_devices` to `ProbingDetails`.
- `ProbingDetails` was manually managed in `connectivitytest_types.go`. The generator produced `ProbingDetails_SingleEdgeResponse` in `types.generated.go` without the `ObservedState` suffix because it didn't recognize it as part of an observed state structure (since the parent was manually managed and missing the field).
- Fixed by updating `connectivitytest_types.go` to include the new fields and the `ProbingDetails_SingleEdgeResponseObservedState` type.
- Re-running generation for `networkmanagement` resolved the conflict and removed the problematic type from `types.generated.go`.
- `make ready-pr` failed due to lack of Docker in the environment, but `go work sync` and manual CRD generation succeeded.
