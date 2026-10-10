# SaaSServiceMgmtTenant Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema Design**:
   - The GCP API for `Tenant` (`google.cloud.saasplatform.saasservicemgmt.v1beta1.Tenant`) defines:
     - `name`: Resource identifier formatted as `projects/{project}/locations/{location}/tenants/{tenant}`.
     - `consumer_resource`: Optional immutable string reference to the consumer resource.
     - `saas`: Required immutable resource reference to `saasservicemgmt.googleapis.com/Saas`.
     - `labels` / `annotations`: Handled via standard Kubernetes object metadata.
     - `uid`, `etag`, `create_time`, `update_time`: Output-only fields placed under `ObservedState`.
   - In KRM `SaaSServiceMgmtTenantSpec`:
     - `ProjectRef` (*refsv1beta1.ProjectRef, required).
     - `Location` (*string, required, immutable).
     - `ResourceID` (*string, optional).
     - `ConsumerResource` (*string, optional, immutable).
     - `SaasRef` (*SaaSServiceMgmtSaaSRef, required, immutable).
   - In KRM `SaaSServiceMgmtTenantObservedState`:
     - `Uid` (*string, optional).
     - `Etag` (*string, optional).
     - `CreateTime` (*string, optional).
     - `UpdateTime` (*string, optional).

2. **Identity & References**:
   - Identity format: `projects/{project}/locations/{location}/tenants/{tenant}` matching CAIS specification `saasservicemgmt.googleapis.com/Tenant`.
   - `SaaSServiceMgmtTenantIdentity` implements `identity.IdentityV2` and `identity.Resource`.
   - `SaaSServiceMgmtTenantRef` implements `refs.Ref` and delegates `Normalize` to `refs.Normalize`.
   - For `SaasRef`, implemented external-only reference `SaaSServiceMgmtSaaSRef` and `SaaSServiceMgmtSaaSIdentity` targeting `projects/{project}/locations/{location}/saas/{saas}`.

3. **Validation & Exceptions**:
   - Added naming violation entry for `saasservicemgmtsaas` in `tests/apichecks/testdata/exceptions/naming_violations.txt` for the external-only `SaaSServiceMgmtSaaS` reference.
   - Added alpha missing fields in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` for `saasservicemgmttenants`.
