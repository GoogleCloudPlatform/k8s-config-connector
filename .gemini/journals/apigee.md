### [2026-06-21] ApigeeApiProduct Identity and Scaffolding Quirks
- **Context**: Implementing ApigeeApiProduct direct types for issue #8061.
- **Problem**: 
  1. The proto package for Apigee is `mockgcp.cloud.apigee.v1` because Apigee uses a direct implementation but its protos are maintained in `mockgcp/apis/` rather than the official `googleapis`.
  2. For the `generate.sh` script, `--service` must be `mockgcp.cloud.apigee.v1`.
  3. `ApigeeApiProduct` parent is an Apigee Organization. The reference field uses `ApigeeOrganizationRef` (from `v1beta1`) instead of standard `ProjectRef`.
  4. ApigeeOrganization identity is of the form `organizations/{organization}`. We need to extract the raw project/org ID using `refs.ParseProject` since customers may use standard project IDs without the prefix in `OrganizationRef.External`.
- **Solution**: 
  - Adjusted `generate.sh` to use `--service mockgcp.cloud.apigee.v1`.
  - Used `apigeev1beta1.ApigeeOrganizationRef` in `apiproduct_types.go`.
  - Parsed the organization ref properly in `apiproduct_identity.go`.
- **Impact**: Future agents working on Apigee resources must use `mockgcp.cloud.apigee.v1` instead of `google.cloud.apigee.v1` for generating types/mappers, and map `OrganizationRef` properly.
