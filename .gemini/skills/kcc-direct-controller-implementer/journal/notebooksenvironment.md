# NotebooksEnvironment Direct Controller Journal

## Observations & Implementation Details
- **Immutability Handling**: The `NotebooksEnvironment` resource is completely immutable in GCP (there is no `UpdateEnvironment` API). Following KCC direct controller standards, we implemented the `compareEnvironment` helper using `mappers.OnlySpecFields` and `tags.DiffForTopLevelFields`.
- **Update Verification**: In the `Update` method, we call `compareEnvironment(ctx, a.actual, a.desired)`. If no diff is detected, we succeed and update the Kubernetes status. If a diff is present, we log the diff with `structuredreporting.ReportDiff(ctx, diffs)` and return a descriptive error stating that the resource is immutable and cannot be updated. This ensures correct and robust KRM reconciliation behavior.
- **Reference Resolution**: We integrated `common.NormalizeReferences(ctx, reader, obj, nil)` in `AdapterForObject` to support any future and existing resource references correctly.
- **Clean Proto Representation**: We avoided holding raw KRM spec structs in the adapter by mapping and storing the desired state as a proto `*pb.Environment` directly in `AdapterForObject`.
