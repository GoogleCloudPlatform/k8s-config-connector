# Migration Center Journal

### 2026-07-01 Greenfield controller for MigrationCenterGroup
- **Context**: Implementing Greenfield direct controller and E2E fixtures for `MigrationCenterGroup` under `v1alpha1`.
- **Problem**: Migration Center is a newly introduced service with no existing mappers, controllers, or MockGCP implementation. 
- **Solution**: Scaffolded a brand new direct controller in `pkg/controller/direct/migrationcenter/`, generated mappers by updating the `generate.sh` script to include `--generate-mapper` and `goimports` formatting, and registered it in the dynamic static config mapping and `register.go`. Created both minimal and maximal KRM E2E test fixtures, then enabled the `migrationcenter.googleapis.com` API on the real GCP sandbox project to record golden objects and HTTP traffic.
- **Impact**: Demonstrates standard Greenfield controller patterns with LRO creation/update/deletion, structured reporting, and full E2E validation against real GCP.
