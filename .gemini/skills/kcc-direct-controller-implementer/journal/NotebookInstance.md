# NotebookInstance Direct Controller Journal

## Observations & Implementation Details
- **Reference Resolution**: Verified that `NotebookInstance` resolves references cleanly using the standardized `common.NormalizeReferences(ctx, reader, obj, nil)` flow in `AdapterForObject`.
- **Identity Integration**: `NotebookInstance` features a dedicated identity class `NotebookInstanceIdentity` which parses external references using standard gcpurls format: `projects/{project}/locations/{location}/instances/{instance}`. This keeps identity-building robust and re-usable.
- **Update Multi-step Orchestration**: The GCP API for `NotebookInstance` has special update requirements. Specifically:
  - Metadata is updated directly using `UpdateInstanceMetadataItems`.
  - Shielded Instance Config changes require first stopping the instance via `StopInstance`, waiting for the operation to complete, applying the `UpdateShieldedInstanceConfig` update, waiting for that operation to complete, and then starting the instance again using `StartInstance` and waiting for the final operation to complete.
  - The custom controller implementation successfully coordinates this sequential state-machine flow within the `Update` reconciliation.
- **Verification**: Verified using `dev/ci/presubmits/tests-e2e-fixtures-notebooks` which executes the basic, full, and minimal e2e tests successfully under `mockgcp`.
