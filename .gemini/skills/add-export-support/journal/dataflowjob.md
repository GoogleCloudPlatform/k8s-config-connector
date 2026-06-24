# Export Support Journal for DataflowJob

## Observations
- Implementing `AdapterForURL` for `DataflowJob` was straightforward since `DataflowJobIdentity` supports parsing full external URLs using `FromExternal`.
- We created the Direct Controller adapter and registered the Model for `DataflowJob` under the `dataflow` package, implementing `AdapterForURL` and `Export` mappings.
- The `DataflowJob` resource has `Region` as an optional spec field, but the identity v2 logic requires `Region` to be present in `obj.Spec.Region`. Because of this constraint and because mockgcp does not emulate classic template launch APIs (e.g., `LaunchTemplate`), the E2E tests for `batchdataflowjob` are skipped against mockgcp. However, the Direct controller implementation was thoroughly verified using targeted package-level unit tests for the `Export` method mapping.
