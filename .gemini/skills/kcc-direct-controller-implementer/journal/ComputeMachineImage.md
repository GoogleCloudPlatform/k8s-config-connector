# ComputeMachineImage Direct Controller Journal

## Implementation Details

- **GCP Client**: Added `newMachineImagesClient` to `pkg/controller/direct/compute/client.go` using GAPIC REST client (`compute.NewMachineImagesRESTClient`).
- **Immutability**: `ComputeMachineImage` has no Update or Patch method in the GCP Compute Engine API. The `Update` method performs spec comparison using `common.CompareBrownfieldSpec` and returns an error if any mutable change is attempted, surfacing diffs via `structuredreporting.ReportDiff`.
- **Reference Trimming**: In `Update` normalizer and `Export`, `refs.TrimComputeURIPrefix` is applied to `sourceInstance` to ensure canonical `projects/.../zones/.../instances/...` representation and avoid spurious diffs against GCP full URLs.
- **Testing**: Golden files and HTTP mock logs were generated and verified against MockGCP, passing `TestAllInSeries` and `dev/ci/presubmits/tests-e2e-fixtures-suite compute`.
