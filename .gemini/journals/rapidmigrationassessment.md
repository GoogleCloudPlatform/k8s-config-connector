# Rapid Migration Assessment Journal

### [2026-07-02] Scaffolding Types and Identity V2 for RapidMigrationAssessmentCollector
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for the `RapidMigrationAssessmentCollector` direct resource. Issue #10295.
- **Problem**: The proto-to-struct generator initially marked types like `Collector`, `GuestOSScan`, and `VSphereScan` as unreachable. We had to manually incorporate these types into the `rapidmigrationassessmentcollector_types.go` file.
- **Solution**: We declared `RapidMigrationAssessmentCollectorSpec` to include the standard spec fields (such as `DisplayName`, `Description`, `ExpectedAssetCount`, `CollectionDays`, `EulaURI`, `Labels`, `ServiceAccount`). We also declared the output-only fields under `RapidMigrationAssessmentCollectorObservedState` and uncommented the `GuestOSScan` and `VSphereScan` structs. Finally, we implemented IdentityV2 in `rapidmigrationassessmentcollector_identity.go` and the `Ref` reference type in `rapidmigrationassessmentcollector_reference.go`.
- **Impact**: Enables a clean schema generation for the `RapidMigrationAssessmentCollector` CRD, aligning perfectly with GCP API schemas and paving the way for the direct controller implementation in subsequent steps.
