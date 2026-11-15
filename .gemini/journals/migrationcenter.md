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

### 2026-11-15 Greenfield Step 1b judgement queue triage for MigrationCenterReportConfig
- **Context**: Triaging judgement queue and finalizing direct KRM types, identity, and references for `MigrationCenterReportConfig` under `v1alpha1`.
- **Solution**:
  - Replaced separate `projectRef` and `location` fields in `MigrationCenterReportConfigSpec` with inlined `*parent.ProjectAndLocationRef`.
  - Moved `ReportConfig_GroupPreferenceSetAssignment` from `types.generated.go` to `migrationcenterreportconfig_types.go`, replacing raw `group` and `preferenceSet` strings with strongly-typed `GroupRef *MigrationCenterGroupRef` and `PreferenceSetRef *MigrationCenterPreferenceSetRef`.
  - Confirmed client-specified resource ID support (`CreateReportConfigRequest` has `report_config_id`), retaining `IdentityV2`.
  - Resolved all 5 entries in `apis/migrationcenter/judgement_queue.yaml`.
  - Regenerated CRD, deepcopy code, mappers, and updated pre-generated controller and fuzzer in lockstep.
- **Impact**: Clean, type-safe API for `MigrationCenterReportConfig` with validated reference relationships and complete test passes.

### 2026-11-15 Greenfield Step 2 direct controller, E2E fixtures, and fuzzer for MigrationCenterReportConfig
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `MigrationCenterReportConfig` under `v1alpha1`.
- **Details**:
  - Direct controller in `pkg/controller/direct/migrationcenter/migrationcenterreportconfig_controller.go` handles immutable `MigrationCenterReportConfig` (GCP Migration Center API only provides `Create`, `Get`, `Delete`, `List` operations and no `Update` RPC).
  - Round-trip fuzzer in `migrationcenterreportconfig_fuzzer.generated.go` covers all proto fields (`description`, `display_name`, `group_preferenceset_assignments`, `create_time`, `update_time`, and identity `.name`).
  - Created minimal and maximal E2E test fixtures with dependencies on `MigrationCenterGroup` and `MigrationCenterPreferenceSet`. Since the resource is immutable, `update.yaml` is omitted.
  - Recorded golden HTTP logs (`_http.log`), golden objects, and audit probes against real GCP using `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`.
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` as all spec fields are now tested.
- **Impact**: Full lifecycle reconciliation for `MigrationCenterReportConfig` verified with authentic GCP audit logs and golden traffic.

### 2026-11-15 Greenfield Step 1a prepopulate types, identity, and judgement queue for MigrationCenterReport
- **Context**: Prepopulating direct KRM types, identity, references, and `judgement_queue.yaml` for `MigrationCenterReport` under `v1alpha1`.
- **Details**:
  - Added `MigrationCenterReport:Report` to `generate-types` and `generate-identity` in `apis/migrationcenter/generate.sh`.
  - Added `google.type` to `generate-mapper`'s `--service` flag in `apis/migrationcenter/generate.sh` to resolve missing external mappers for `Money_FromProto` / `Money_ToProto` referenced in `ReportSummary`.
  - Added `MigrationCenterReport` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Ran `apis/migrationcenter/generate.sh` to scaffold `migrationcenterreport_types.go`, identity, reference, CRD, and entries in `apis/migrationcenter/judgement_queue.yaml`.
  - Recorded untriaged fields in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Mechanical prepopulation complete with clean compilation, passing unit tests, and passing `apichecks`.

### 2026-11-15 Greenfield Step 1b judgement queue triage for MigrationCenterReport
- **Context**: Triaging judgement queue and finalizing direct KRM types, identity, and references for `MigrationCenterReport` under `v1alpha1`.
- **Solution**:
  - Replaced separate `projectRef`, `location`, and guessed `reportConfigRef` fields in `MigrationCenterReportSpec` with `spec.parentRef *MigrationCenterReportConfigRef` following `PARENT_AND_IDENTITY.md` Case B.
  - Moved server-populated `state` field from `MigrationCenterReportSpec` to `MigrationCenterReportObservedState.state`.
  - Confirmed client-specified resource ID support (`CreateReportRequest` has `report_id`), retaining `IdentityV2`.
  - Resolved all 6 entries in `apis/migrationcenter/judgement_queue.yaml`.
  - Regenerated CRD, deepcopy code, identity, and mappers via `apis/migrationcenter/generate.sh`.
  - Re-recorded `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` to reflect the updated spec fields.
- **Impact**: Clean, type-safe API for `MigrationCenterReport` ready for Phase 2 direct controller implementation.




