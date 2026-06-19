### [2026-05-27] SecurityHealthAnalyticsCustomModule Multiple Parent References
- **Context**: Scaffolding SecurityCenterManagementSecurityHealthAnalyticsCustomModule.
- **Problem**: The proto explicitly uses a multi-parent schema where the module can exist under an organization, folder, or project. The `generate-types` scaffolding logic defaulted to generating only `ProjectRef`.
- **Solution**: Manually added `OrganizationRef`, `FolderRef`, and `ProjectRef` to the `Spec` struct using KCC standards (`refsv1beta1.OrganizationRef` etc.). Updated the identity implementation to use `refsv1beta1.ResolveProjectID`, `refsv1beta1.ResolveFolder`, and `refsv1beta1.ResolveOrganization` appropriately to capture the correct parent dynamically.
- **Impact**: Any subsequent direct controller implementation for this resource will need to correctly route API calls based on which parent reference was supplied by the user. Ensure you test all three hierarchy levels.
