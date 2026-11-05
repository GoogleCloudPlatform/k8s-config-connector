# Migration Center Journal

### 2026-07-01 Greenfield controller for MigrationCenterGroup
- **Context**: Implementing Greenfield direct controller and E2E fixtures for `MigrationCenterGroup` under `v1alpha1`.
- **Problem**: Migration Center is a newly introduced service with no existing mappers, controllers, or MockGCP implementation. 
- **Solution**: Scaffolded a brand new direct controller in `pkg/controller/direct/migrationcenter/`, generated mappers by updating the `generate.sh` script to include `--generate-mapper` and `goimports` formatting, and registered it in the dynamic static config mapping and `register.go`. Created both minimal and maximal KRM E2E test fixtures, then enabled the `migrationcenter.googleapis.com` API on the real GCP sandbox project to record golden objects and HTTP traffic.
- **Impact**: Demonstrates standard Greenfield controller patterns with LRO creation/update/deletion, structured reporting, and full E2E validation against real GCP.

### 2026-09-28 Greenfield controller for MigrationCenterPreferenceSet
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `MigrationCenterPreferenceSet` under `v1alpha1`.
- **Solution**: Implemented direct controller and fuzzer in `pkg/controller/direct/migrationcenter/` reusing generated mappers. Configured minimal and maximal test fixtures with complete field coverage across all compute, VMware Engine, and Sole Tenancy preferences, and recorded golden HTTP traffic and objects against real GCP.
- **Impact**: Full CRUD reconciliation support for `MigrationCenterPreferenceSet` with 100% test coverage for all CRD spec and status fields.

### 2026-10-14 Greenfield controller for MigrationCenterSource
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `MigrationCenterSource` under `v1alpha1`.
- **Solution**: Implemented direct controller in `pkg/controller/direct/migrationcenter/migrationcentersource_controller.go` and fuzzer in `migrationcentersource_fuzzer.go` reusing existing generated mappers. Registered `MigrationCenterSource` in `pkg/controller/resourceconfig/static_config.go`. Created minimal and maximal test fixtures, recorded golden outputs and HTTP traffic against real GCP using `hack/record-gcp` with `RECORD_AUDIT_PROBE=1`, and updated test coverage exceptions.
- **Impact**: Full CRUD reconciliation support for `MigrationCenterSource` with 100% test coverage for all CRD spec and status fields.

