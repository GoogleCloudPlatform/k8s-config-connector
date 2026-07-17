# ComputeImage Direct Controller Journal

## Implementation Details

- **Direct Controller and Mapper**: The direct controller for `ComputeImage` is implemented under `pkg/controller/direct/compute/computeimage_controller.go` and its mapper functions are defined in `pkg/controller/direct/compute/computeimage_mapper.go`.
- **Fuzzer Integration**: The KRM fuzzer is registered in `pkg/controller/direct/compute/image_fuzzer.go` and was verified using the focused roundtrip tests via `FOCUS=ComputeImage go test ./pkg/fuzztesting/fuzztests`.
- **E2E Testing and MockGCP Validation**: We verified the direct reconciliation paths using both the `computeimage` and `computeimage-direct` test fixtures. They execute successfully and completely under mockgcp, validating all CRUD operations, updates, status mapping, and deletion.
- **Gradual Opt-in (Brownfield Migration)**: In `pkg/controller/resourceconfig/static_config.go`, `ComputeImage` is configured with `k8s.ReconcilerTypeDirect` in its `SupportedControllers` list, while retaining `k8s.ReconcilerTypeTerraform` as the default reconciler. This ensures that existing configurations continue to reconcile seamlessly using the legacy controller while new or overridden resources can leverage the direct controller.
