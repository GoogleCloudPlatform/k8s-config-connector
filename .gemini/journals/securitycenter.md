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
