# ComputeDisk Direct Controller Journal

## Implementation Details

- **Zonal and Regional Disks**: A ComputeDisk can be either zonal or regional based on its `location` parameter. The direct controller detects whether `location` corresponds to a zone or region (using the count of hyphens) and dynamically instantiates/uses either the `compute.DisksClient` or `compute.RegionDisksClient`.
- **Default Fields Comparison & Server-populated Defaults**: In `compareComputeDisk`, server-side or mockgcp-defaulted fields such as `type` (pd-standard), `physical_block_size_bytes` (4096), and default disk encryption keys are populated onto the cloned desired object if omitted. This avoids unnecessary trigger of GCP updates during round-trip comparisons.
- **E2E Testing against MockGCP**: All 3 test fixtures (`zonalcomputedisk`, `regionalcomputedisk`, and `computediskfromsourcedisk`) were fully verified against mockgcp and passed successfully.
