# Jobs Service Journal

### 2026-07-02 CloudTalentSolutionCompany Initial Implementation
- **Context**: Implementing the direct types, CRD, and IdentityV2 for `CloudTalentSolutionCompany` (Kind: `CloudTalentSolutionCompany`, Group: `jobs.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: 
  1. The issue instructions suggested using `--service google.cloud.jobs.v1`, but the compiled protobuf definition in Google APIs holds the `Company` proto message inside the package `google.cloud.talent.v4` (the API itself is known as Cloud Talent Solution). Running `--service google.cloud.jobs.v1` failed with "proto: not found".
  2. `Company` is a hierarchical child of `Tenant` (`projects/{project}/tenants/{tenant}/companies/{company}`). However, there is currently no KRM resource for `Tenant` in KCC.
- **Solution**: 
  1. Ran the types and mapper generator with `--service google.cloud.talent.v4` as the service name, which successfully located and scaffolded the `Company` proto message and fields.
  2. Designed and implemented a local `TenantRef` reference struct and `TenantIdentity` in the `apis/jobs/v1alpha1/` package. Since `Tenant` has no KRM resource, we registered it using `refs.Register(&TenantRef{}, nil)`.
  3. Structured the `CloudTalentSolutionCompany` identity resolver to resolve and normalize the parent `TenantRef`, parsing it into a `TenantIdentity` to retrieve the GCP `project` and `tenant` IDs needed to construct the company URL.
- **Impact**: Unblocks the implementation of CloudTalentSolutionCompany by correctly resolving the service and parent hierarchy when a parent KRM resource does not yet exist.
