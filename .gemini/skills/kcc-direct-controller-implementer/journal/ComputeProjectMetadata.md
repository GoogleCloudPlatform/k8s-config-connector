# ComputeProjectMetadata Direct Controller Journal

## Implementation Details

- **Controller & Adapter**:
  - Implemented the direct controller for `ComputeProjectMetadata` under `pkg/controller/direct/compute/computeprojectmetadata_controller.go`.
  - Used `compute.NewProjectsRESTClient` via `gcpClient.newProjectsClient(ctx)`.
  - Registered the GVK `ComputeProjectMetadataGVK` with `registry.RegisterModel`.
  - Registered dual-controller support in `pkg/controller/resourceconfig/static_config.go` with `SupportedControllers: [Direct, Terraform]` keeping default as `Terraform` for brownfield gradual migration.

- **Diff Comparison**:
  - Used `mappers.OnlySpecFields` and `common.DiffForTopLevelFields` for clean top-level field diffing on `items`.
  - Both desired and actual items are deterministically sorted by key in proto mappers, ensuring zero false diffs and clean steady-state re-reconciliation (0 writes).

- **E2E & Migration Testing**:
  - Tested both the Direct controller and Fallback-to-legacy controller scenarios using `hack/compare-mock`.
  - Verified seamless migration and takeover using `TestMigrationToDirect`.
