### [2026-05-27] SecurityCenterMuteConfig Greenfield Types
- **Context**: Implementing Greenfield types for `SecurityCenterMuteConfig`.
- **Problem**: Identity required parent-level references to `OrganizationRef` and mapping to the exact Identity template `organizations/{organization}/muteConfigs/{mute_config}`.
- **Solution**: We implemented `IdentityV2` matching this template, using `refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)` to acquire the `OrganizationID`. We explicitly removed `ProjectRef` and `Location` which are scaffolded by default because the required identity template strictly expects an organization parent. 
- **Impact**: Demonstrates handling organizational hierarchies in IdentityV2 over the more common project/location template.
