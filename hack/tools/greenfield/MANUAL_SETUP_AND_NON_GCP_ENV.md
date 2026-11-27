# `non-gcp-env` and `manual-setup` Resource Reference

This document lists every rule and resource in **`non-gcp-env`** ([`calculate_coverage.py`](./calculate_coverage.py)) and **`manual-setup`** ([`manual_setup.json`](./manual_setup.json)), with each resource's KCC status and the exact environment or manual setup behind it.

**How `KCC Status` is labeled below:**
- **`net-new`**: No KCC CRD or controller yet.
- **`types-only`**: CRD / Go types exist in KCC, no controller yet.
- **`implemented (direct, unreleased)`**: Direct controller exists in KCC (`alpha`), still listed in `BLOCKED_KINDS` (`dev/tasks/build-release-bundle`) and not yet in the release bundle.
- **Note on upstream-only rows**: A few entries target resources present in newer upstream `googleapis` that are not yet in KCC's local `.build/third_party/googleapis` checkout; these are marked `(upstream googleapis)`.

---

## 1. `non-gcp-env` Resources

GCP is cloud software infrastructure, so resources whose target environment is either a **non-GCP product/deployment** (`NON_GCP_ENV_PACKAGES`) or **physical hardware** (`NON_GCP_ENV_RESOURCE_TYPES`) are excluded here. No released KCC kinds match this rule today, so all matching resources listed below are currently unreleased.

### A. Non-GCP Products & Deployments (`NON_GCP_ENV_PACKAGES`)

| Resource Type | KCC Kind | KCC Status | Non-GCP Environment |
|---|---|---|---|
| `gsuiteaddons.googleapis.com/Deployment` | `GSuiteAddonsDeployment` | `implemented (direct, unreleased)` | **Google Workspace Add-ons** (`google.cloud.gsuiteaddons.*`): deploys Workspace add-ons for Gmail, Docs, Sheets, Calendar, and Drive rather than GCP infrastructure. *(Read-only `Authorization` and `InstallStatus` drop earlier at `read-only`.)* |
| `resultstore.googleapis.com/Action` | `ResultStoreAction` | `net-new` | **Bazel ResultStore** (`google.devtools.resultstore.*`): stores Bazel build/test action results rather than GCP infrastructure. |
| `resultstore.googleapis.com/Configuration` | `ResultStoreConfiguration` | `net-new` | **Bazel ResultStore** (`google.devtools.resultstore.*`): stores Bazel build configuration metadata. |
| `resultstore.googleapis.com/ConfiguredTarget` | `ResultStoreConfiguredTarget` | `net-new` | **Bazel ResultStore** (`google.devtools.resultstore.*`): stores Bazel configured build target data. |
| `resultstore.googleapis.com/FileSet` | `ResultStoreFileSet` | `net-new` | **Bazel ResultStore** (`google.devtools.resultstore.*`): stores Bazel build output file sets. |
| `resultstore.googleapis.com/Invocation` | `ResultStoreInvocation` | `net-new` | **Bazel ResultStore** (`google.devtools.resultstore.*`): stores Bazel build/test invocation records. *(Read-only `DownloadMetadata` and `UploadMetadata` drop earlier at `read-only`.)* |
| `resultstore.googleapis.com/Target` | `ResultStoreTarget` | `net-new` | **Bazel ResultStore** (`google.devtools.resultstore.*`): stores Bazel build target records. |
| `spanner.googleapis.com/ExternalStorage` *(upstream googleapis)* | `SpannerExternalStorage` | `net-new` | **Spanner Omni** (`google.spanner.admin.omni.*`): external storage config for non-GCP Spanner Omni deployments (e.g., AWS/Azure). |
| `spanner.googleapis.com/LocationDistance` *(upstream googleapis)* | `SpannerLocationDistance` | `net-new` | **Spanner Omni** (`google.spanner.admin.omni.*`): topology distance config for non-GCP Spanner Omni deployments. |
| `spanner.googleapis.com/Server` *(upstream googleapis)* | `SpannerServer` | `net-new` | **Spanner Omni** (`google.spanner.admin.omni.*`): server registration for non-GCP Spanner Omni deployments. *(Other `omni` / `auth` resources like `BackupDescriptor`, `Role`, `Rollout`, `RolloutPhase`, `Location`, `Zone` drop earlier at `read-only`, `no-crud`, or `deprecated`.)* |
| `spanner.googleapis.com/User` *(upstream googleapis)* | `SpannerUser` | `net-new` | **Spanner Omni Standalone Auth** (`google.spanner.auth.*`): standalone user accounts for non-GCP Spanner Omni (instead of GCP IAM). |
| `gdchardwaremanagement.googleapis.com/*` (`8` resources, currently `v1alpha`) | `GDCHardwareManagement*` | `net-new` (drops at `non-ga-only`) | **Google Distributed Cloud Connected** (`google.cloud.gdchardwaremanagement.*`): physical rack/hardware ordering and on-prem site management. Currently drops earlier at `non-ga-only`. |
| `baremetalsolution.googleapis.com/*` (`15` resources) | `BareMetalSolution*` | `net-new` (drops earlier at `no-crud` [`5`], `read-only` [`3`], `nondeclarative:data-plane` [`1`], or here at `non-gcp-env` [`6`]) | **Bare Metal Solution** (`google.cloud.baremetalsolution.*`): dedicated physical bare-metal servers, SAN LUNs, and NFS/provisioning hardware under a specialized allowlist-only availability model. |

### B. Physical Hardware & External-Environment Resources (`NON_GCP_ENV_RESOURCE_TYPES`)

| Resource Type | KCC Kind | KCC Status | Physical Hardware / External-Environment Dependency |
|---|---|---|---|
| `compute.googleapis.com/Interconnect` | `ComputeInterconnect` | `types-only` | Requires physical carrier colocation fiber cross-connect and LOA-CFA provisioning outside GCP. |
| `compute.googleapis.com/WireGroup` *(upstream googleapis)* | `ComputeWireGroup` | `net-new` | Requires physical/dedicated interconnect wire attachments between colocation sites. |
| `cloudkms.googleapis.com/SingleTenantHsmInstance` *(upstream googleapis)* | `KMSSingleTenantHSMInstance` | `net-new` | Requires dedicated physical single-tenant HSM hardware allocation (`#16010`, `#16003`). |
| `cloudkms.googleapis.com/SingleTenantHsmInstanceProposal` *(upstream googleapis)* | `KMSSingleTenantHSMInstanceProposal` | `net-new` | Approval/proposal workflow tied to a physical single-tenant HSM hardware unit (`#16010`). |
| `vmmigration.googleapis.com/DatacenterConnector` | `VMMigrationDatacenterConnector` | `implemented (direct, unreleased)` | Connects an on-prem VMware vCenter source to GCP via a user-installed OVA appliance (`registration_id`). |
| `vmmigration.googleapis.com/MigratingVm` | `VMMigrationMigratingVM` | `implemented (direct, unreleased)` | Migrates a live `source_vm_id` hosted in an external VMware vCenter or AWS environment. |
| `vmmigration.googleapis.com/Source` | `VMMigrationSource` | `implemented (direct, unreleased)` | Points to an external VMware vCenter or AWS environment outside GCP. |
| `vmwareengine.googleapis.com/Cluster` | `VMwareEngineCluster` | `implemented (direct, unreleased)` | Requires dedicated VMware Engine PrivateCloud backed by bare-metal VE1 physical nodes (`#15133`). |
| `vmwareengine.googleapis.com/DnsForwarding` | `VMwareEngineDNSForwarding` | `net-new` | Requires a VMware Engine PrivateCloud backed by dedicated bare-metal hardware. |
| `vmwareengine.googleapis.com/HcxActivationKey` | `VMwareEngineHCXActivationKey` | `net-new` | Requires a VMware Engine PrivateCloud backed by dedicated bare-metal hardware. |
| `vmwareengine.googleapis.com/LoggingServer` | `VMwareEngineLoggingServer` | `implemented (direct, unreleased)` | Requires a VMware Engine PrivateCloud backed by dedicated bare-metal hardware. |
| `vmwareengine.googleapis.com/ManagementDnsZoneBinding` | `VMwareEngineManagementDNSZoneBinding` | `implemented (direct, unreleased)` | Requires a VMware Engine PrivateCloud backed by dedicated bare-metal hardware (`#15133`). |
| `vmwareengine.googleapis.com/Subnet` | `VMwareEngineSubnet` | `net-new` | Requires a VMware Engine PrivateCloud backed by dedicated bare-metal hardware. |

---

## 2. `manual-setup` Resources

`manual-setup` (`manual_setup.json`) applies **only to unreleased resources** (`net-new`, `types-only`, or `implemented (direct, unreleased)` in `BLOCKED_KINDS`). Released `v1alpha1` / `v1beta1` KCC kinds are never dropped by this rule.

### A. Service-Team Project Allowlisting or Cloud Console UI Opt-In

| Resource Type | KCC Kind | KCC Status | Manual Setup Needed |
|---|---|---|---|
| `compute.googleapis.com/HealthCheckService` | `ComputeHealthCheckService` | `types-only` | Service-team GCP project allowlisting for Compute HealthCheck-as-a-Service (`#15638`). |
| `compute.googleapis.com/NotificationEndpoint` | `ComputeNotificationEndpoint` | `types-only` | Service-team GCP project allowlisting for Compute HealthCheck-as-a-Service (`#15638`, `#15535`). |
| `eventarc.googleapis.com/ChannelConnection` | `EventarcChannelConnection` | `implemented (direct, unreleased)` | Service-team server-side project allowlisting / third-party Eventarc provider approval (`#15936`). |
| `financialservices.googleapis.com/BacktestResult` | `FinancialServicesBacktestResult` | `implemented (direct, unreleased)` | Early Access service-team project allowlist approval for Anti-Money Laundering AI (`#14955`). |
| `financialservices.googleapis.com/Dataset` | `FinancialServicesDataset` | `implemented (direct, unreleased)` | Early Access service-team project allowlist approval for Anti-Money Laundering AI (`#14955`). |
| `financialservices.googleapis.com/EngineConfig` | `FinancialServicesEngineConfig` | `implemented (direct, unreleased)` | Early Access service-team project allowlist approval for Anti-Money Laundering AI (`#14955`). |
| `financialservices.googleapis.com/Instance` | `FinancialServicesInstance` | `implemented (direct, unreleased)` | Early Access service-team project allowlist approval for Anti-Money Laundering AI (`#14955`). |
| `financialservices.googleapis.com/Model` | `FinancialServicesModel` | `implemented (direct, unreleased)` | Early Access service-team project allowlist approval for Anti-Money Laundering AI (`#14955`). |
| `financialservices.googleapis.com/PredictionResult` | `FinancialServicesPredictionResult` | `implemented (direct, unreleased)` | Early Access service-team project allowlist approval for Anti-Money Laundering AI (`#14955`). |
| `jobs.googleapis.com/Company` | `CloudTalentSolutionCompany` | `implemented (direct, unreleased)` | Manual Cloud Console UI data-permission opt-in click for Cloud Talent Solution (`#15827`). |
| `jobs.googleapis.com/Tenant` | `CloudTalentSolutionTenant` | `implemented (direct, unreleased)` | Manual Cloud Console UI data-permission opt-in click for Cloud Talent Solution (`#15827`). |

### B. Commercial / Partner Enrollment or Marketplace Subscription

| Resource Type | KCC Kind | KCC Status | Manual Setup Needed |
|---|---|---|---|
| `cloudchannel.googleapis.com/ChannelPartnerLink` | `CloudchannelChannelPartnerLink` | `implemented (direct, unreleased)` | Google Cloud Partner Advantage reseller/distributor enrollment and Partner Sales Console setup (`#15951`). |
| `cloudchannel.googleapis.com/ChannelPartnerRepricingConfig` | `CloudchannelChannelPartnerRepricingConfig` | `implemented (direct, unreleased)` | Google Cloud Partner Advantage reseller/distributor enrollment and Partner Sales Console setup (`#15951`). |
| `cloudchannel.googleapis.com/Customer` | `CloudchannelCustomer` | `implemented (direct, unreleased)` | Google Cloud Partner Advantage reseller/distributor enrollment and Partner Sales Console setup (`#15951`). |
| `cloudchannel.googleapis.com/CustomerRepricingConfig` | `CloudchannelCustomerRepricingConfig` | `implemented (direct, unreleased)` | Google Cloud Partner Advantage reseller/distributor enrollment and Partner Sales Console setup (`#15951`). |
| `cloudchannel.googleapis.com/Entitlement` | `CloudchannelEntitlement` | `net-new` | Google Cloud Partner Advantage reseller/distributor enrollment and Partner Sales Console setup (`#15951`). |
| `cloudcontrolspartner.googleapis.com/Customer` | `CloudcontrolspartnerCustomer` | `implemented (direct, unreleased)` | Cloud Controls Partner organization enrollment / partner onboarding. |
| `oracledatabase.googleapis.com/AutonomousDatabase` | `OracleDatabaseAutonomousDatabase` | `types-only` | Active commercial Oracle Database@Google Cloud Marketplace subscription and OCI tenancy linked to Cloud Billing (`#15655`). |
| `oracledatabase.googleapis.com/CloudExadataInfrastructure` | `OracleDatabaseCloudExadataInfrastructure` | `implemented (direct, unreleased)` | Active commercial Oracle Database@Google Cloud Marketplace subscription and OCI tenancy linked to Cloud Billing (`#15655`). |
| `oracledatabase.googleapis.com/CloudVmCluster` | `OracleDatabaseCloudVMCluster` | `implemented (direct, unreleased)` | Active commercial Oracle Database@Google Cloud Marketplace subscription and OCI tenancy linked to Cloud Billing (`#15655`). *(Matches all mutable `oracledatabase.googleapis.com/*` resources via prefix rule, including upstream-only kinds such as `ExadbVmCluster` [`types-only`] and `Database` / `DbSystem` [`net-new`].)* |
| `cloudsupport.googleapis.com/SupportEventSubscription` *(upstream googleapis)* | `CloudSupportSupportEventSubscription` | `types-only` | Active paid Enhanced or Premium Cloud Support tier on the organization plus `roles/cloudsupport.admin` (`#15735`). |

### C. External Verification or Third-Party Endpoint Setup

| Resource Type | KCC Kind | KCC Status | Manual Setup Needed |
|---|---|---|---|
| `appengine.googleapis.com/AuthorizedCertificate` | `AppEngineAuthorizedCertificate` | `types-only` | Out-of-band domain ownership verification via Google Search Console or public DNS TXT records (`#15767`). |
| `domains.googleapis.com/Registration` | `DomainsRegistration` | `net-new` | Live public DNS domain purchase and ICANN registrant contact verification. |
| `compute.googleapis.com/PublicAdvertisedPrefix` | `ComputePublicAdvertisedPrefix` | `types-only` | Public BYOIP IPv4/IPv6 ownership proof, RIR RPKI ROA / reverse-DNS setup, and isolated testing org (`#16073`, `#16058`). |
| `compute.googleapis.com/PublicDelegatedPrefix` | `ComputePublicDelegatedPrefix` | `net-new` | Requires a manually verified and provisioned BYOIP `PublicAdvertisedPrefix` (`#16073`, `#16058`). |
| `cloudkms.googleapis.com/EkmConfig` | `KMSEKMConfig` | `net-new` | External third-party KMS/EKM service endpoint and trust setup outside GCP. |
| `cloudkms.googleapis.com/EkmConnection` | `KMSEkmConnection` | `implemented (direct, unreleased)` | External third-party EKM TLS/IAM endpoint reachable by Cloud KMS. |

### D. Billing Account, Commitments, or Isolated Billed-Project Singletons

| Resource Type | KCC Kind | KCC Status | Manual Setup Needed |
|---|---|---|---|
| `accessapproval.googleapis.com/AccessApprovalSettings` | `AccessApprovalAccessApprovalSettings` | `net-new` | Organization/project with a qualifying paid Google Cloud Support SLA / Enterprise tier and active billing. |
| `apihub.googleapis.com/ApiHubInstance` | `APIHubInstance` | `types-only` | Dedicated isolated GCP project linked to an active billing account (one instance per project). |
| `apihub.googleapis.com/HostProjectRegistration` | `APIHubHostProjectRegistration` | `net-new` | Dedicated isolated GCP project linked to an active billing account for API Hub registration. |
| `apihub.googleapis.com/StyleGuide` | `APIHubStyleGuide` | `net-new` | API Hub host project singleton requiring a dedicated isolated project with an active billing account. |
| `appengine.googleapis.com/Application` | `AppEngineApplication` | `net-new` | Permanent, non-deletable one-per-project singleton requiring an isolated project with an active billing account. |
| `cloudbilling.googleapis.com/ProjectBillingInfo` | `BillingProjectBillingInfo` | `net-new` | Real billing account and billing-admin permissions (`roles/billing.user`) to attach/change a project's billing account. |
| `compute.googleapis.com/Commitment` | `ComputeCommitment` | `net-new` | Purchases an uncancellable 1-year or 3-year CUD billing commitment contract (no `Delete` RPC). |
| `dlp.googleapis.com/DiscoveryConfig` | `DLPDiscoveryConfig` | `implemented (direct, unreleased)` | Paid one-per-project/org discovery singleton requiring an isolated billed project or org. |
| `storage.googleapis.com/IntelligenceConfig` | `StorageIntelligenceConfig` | `net-new` | Paid Storage Intelligence billing enablement requiring an isolated project with an active billing account. |
| `auditmanager.googleapis.com/AuditSchedule` *(upstream googleapis)* | `AuditmanagerAuditSchedule` | `net-new` | Manual Audit Manager enrollment on a billing-enabled project plus destination GCS bucket setup. |
| `datalineage.googleapis.com/Config` *(upstream googleapis)* | `DataLineageConfig` | `net-new` | Billed Data Lineage project/org singleton requiring an isolated project with an active billing account. |

### E. Organization-Wide Security / IAM Enforcement Requiring a Separate Testing Org

| Resource Type | KCC Kind | KCC Status | Manual Setup Needed |
|---|---|---|---|
| `cloudsecuritycompliance.googleapis.com/FrameworkDeployment` | `CloudSecurityComplianceFrameworkDeployment` | `implemented (direct, unreleased)` | Dedicated isolated testing org/folder so compliance enforcement does not affect shared org environments. |
| `dataplex.googleapis.com/EncryptionConfig` | `DataplexEncryptionConfig` | `implemented (direct, unreleased)` | Dedicated testing organization for org-wide Dataplex CMEK enforcement. |
| `iam.googleapis.com/PolicyBinding` | `IAMPolicyBinding` | `net-new` | Dedicated testing organization to bind org/folder-wide Principal Access Boundary / Access Policies without lockout risk. |
| `iam.googleapis.com/PrincipalAccessBoundaryPolicy` | `IAMPrincipalAccessBoundaryPolicy` | `implemented (direct, unreleased)` | Dedicated testing organization for org-level IAM principal access boundary policies (lockout risk on shared org). |
| `iam.googleapis.com/AccessPolicy` *(upstream googleapis)* | `IAMAccessPolicy` | `types-only` | Dedicated testing organization for org/folder-wide IAM access policies (lockout risk on shared org). |
| `securitycentermanagement.googleapis.com/SecurityCenterService` | `SecurityCenterManagementSecurityCenterService` | `net-new` | Dedicated test org/folder and SCC tier setup to toggle top-level Security Command Center services. |
| `securityposture.googleapis.com/PostureDeployment` | `SecurityposturePostureDeployment` | `implemented (direct, unreleased)` | Dedicated testing organization to deploy and enforce an organization-wide security posture. |
