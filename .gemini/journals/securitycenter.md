### [2026-05-27] SecurityCenterMuteConfig Greenfield Types
- **Context**: Implementing Greenfield types for `SecurityCenterMuteConfig`.
- **Problem**: Identity required parent-level references to `OrganizationRef` and mapping to the exact Identity template `organizations/{organization}/muteConfigs/{mute_config}`.
- **Solution**: We implemented `IdentityV2` matching this template, using `refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)` to acquire the `OrganizationID`. We explicitly removed `ProjectRef` and `Location` which are scaffolded by default because the required identity template strictly expects an organization parent. 
- **Impact**: Demonstrates handling organizational hierarchies in IdentityV2 over the more common project/location template.

### [2026-05-27] generate-types quirk for SecurityCenterBigQueryExport
- **Context**: Implementing Greenfield Step 1 types for `SecurityCenterBigQueryExport`.
- **Problem**: The KCC `generate-types` tool created an empty `types.generated.go` and omitted the primitive scalar fields (like `Description`, `Filter`, `Dataset`) from `SecurityCenterBigQueryExportSpec`. Additionally, it defaulted to scaffolding `ProjectRef` instead of `OrganizationRef`.
- **Solution**: Manually populated `SecurityCenterBigQueryExportSpec` and `SecurityCenterBigQueryExportObservedState` in `bigqueryexport_types.go` from the proto definition, replacing `ProjectRef` with `OrganizationRef` and mapping `Dataset` to `DatasetRef`.
- **Impact**: When bootstrapping types for resources with primitive top-level fields, agents must be prepared to manually transcribe fields from the protobuf into the generated `_types.go` file.

### [2026-10-12] SecurityCenterBigQueryExport Golden Log Alignment
- **Context**: Aligning MockGCP logs with RealGCP for `SecurityCenterBigQueryExport`.
- **Problem**: `SecurityCenterBigQueryExport` requires organization-level permissions and quota projects (`organizations/${TEST_ORG_ID}`) on real GCP. Similar to `SecurityCenterMuteConfig`, authentic `_http.log` cannot be recorded in standard single-project test environments.
- **Solution**: Added `bigqueryexport-maximal` and `bigqueryexport-minimal` to `realGCPSkipFixtures` in `golden_alignment_test.go`, following established project precedent. Added unit tests for MockGCP service/lifecycle and a round-trip KRM fuzzer test for the direct controller.
- **Impact**: Ensures MockGCP tests run hermetically with full test coverage while skipping real GCP log alignment checks that require organization access.
