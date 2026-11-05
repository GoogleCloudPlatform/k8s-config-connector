# SaaSServiceMgmtTenant Controller Logic Journal

## Observations & Patterns

### 1. Synchronous API & GAPIC REST Client
- The App Lifecycle Manager (SaaS Service Management) deployments API (`saasservicemgmt.googleapis.com`) operations (`CreateTenant`, `GetTenant`, `UpdateTenant`, `DeleteTenant`) are synchronous and do not return Long Running Operations (LROs).
- Used official GAPIC REST client (`NewSaasDeploymentsRESTClient`).

### 2. Immutability of Spec Fields
- All spec fields (`projectRef`, `location`, `resourceID`, `saasRef`, `consumerResource`) are immutable in GCP and enforced via kubebuilder immutable validations on the KRM spec.
- Per Greenfield controller instructions, when all spec fields are immutable, `update.yaml` is omitted from minimal and maximal fixtures.
- Dynamic metadata labels and annotations are still synchronized and compared via `compareTenant` using `common.DiffForTopLevelFields`.

### 3. Non-Validating Identity References in Tests
- `saasRef` points to a SaaS resource (`projects/{project}/locations/{location}/saas/{saas}`). The GCP API allows creating a tenant even with arbitrary resource paths without pre-creating the referenced SaaS resource in the GCP project.
