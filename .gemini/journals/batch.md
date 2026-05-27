### [2026-05-26] Multi-version Generation in Same Package
- **Context**: Implementing CloudBatchResourceAllowance alongside BatchJob and BatchTask.
- **Problem**: BatchJob and BatchTask use v1, but ResourceAllowance only exists in v1alpha. generate-types overwrites types.generated.go.
- **Solution**: Run generate-types for v1alpha, then rename types.generated.go to resourceallowance_types.generated.go. Then run generate-types for v1 for the rest. Ensure all fields referencing output-only structs map to their ObservedState variant explicitly.
- **Impact**: Enables mixing GCP API versions within a single KRM api_version group without overwriting generated files.

### [2026-05-26] Output-Only Wrapped Messages
- **Context**: Implementing CloudBatchResourceAllowance spec and status.
- **Problem**: A UsageResourceAllowance wrapper containing Spec and output-only Status was pruned incorrectly.
- **Solution**: Explicitly define the wrapper types and their ObservedState variants in resourceallowance_types.go so prunetypes treats them as reachable.
- **Impact**: Prevents prunetypes from commenting out essential types.
