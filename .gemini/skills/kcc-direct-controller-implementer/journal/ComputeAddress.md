# ComputeAddress Direct Controller Journal

## Implementation Details

- **Diff Comparison**: Kept `compareAddress` using `tags.DiffForTopLevelFields` per reviewer/design requirements, avoiding recursive comparison paths.
- **Immutability Behavior**: ComputeAddress is completely immutable in GCP. Reconciling changed spec fields triggers a structured reporting update status explaining that `ComputeAddress is immutable and cannot be updated`.
- **E2E Testing against MockGCP**: All direct E2E test fixtures (`computeaddressipcollection-direct`, `globalcomputeaddress-direct`, `regionalcomputeaddress-direct`) were fully verified against MockGCP and passed successfully.
