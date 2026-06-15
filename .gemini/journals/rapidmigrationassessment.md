### [2026-06-15] Initial Greenfield Scaffolding of RapidMigrationAssessmentCollector
- **Context**: Implementing KRM types, CRD, and IdentityV2 for `RapidMigrationAssessmentCollector` under `rapidmigrationassessment.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: When generating types using the generator tool, referenced structs like `GuestOSScan` and `VSphereScan` under `CollectorObservedState` are marked as `unreachable` and commented out because they are not initially declared/used in `collector_types.go`.
- **Solution**: Explicitly defined the custom types and added appropriate fields/references in `collector_types.go`. Re-running the generator script correctly kept the Go structures active in `types.generated.go`. We also registered the GVK directly using `GroupVersion.WithKind` to avoid redeclaration in the reference file.
- **Impact**: Ensures standard Greenfield scaffolding and flawless compilation of the first direct `rapidmigrationassessment` resource.
