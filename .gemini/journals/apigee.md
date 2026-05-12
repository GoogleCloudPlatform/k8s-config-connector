### [2026-05-12] Apigee ApiProduct synchronous operations and ApigeeOrganizationRef
- **Context**: Implementing the direct controller for ApigeeApiProduct.
- **Problem**: Most Apigee resources like Instance use LROs, but ApigeeApiProduct operations return the actual object immediately. Additionally, Apigee resources typically use `ApigeeOrganizationRef` instead of `ProjectRef` because an Apigee Organization is a 1:1 mapping with a GCP Project.
- **Solution**: Avoided using `WaitForApigeeOp` in the ApiProduct adapter's Create/Update/Delete operations and directly processed the returned object. Ensured the types use `ApigeeOrganizationRef`. Also, since the identity is specific to Apigee, the normal ProjectRef mappings needed to be replaced by custom logic resolving `ApigeeOrganizationRef.External`.
- **Impact**: Future agents working on Apigee resources shouldn't blindly assume all endpoints are LROs or use `ProjectRef`.
