# Technical Design: Standardizing Resource Kind Names

## Problem Statement
Config Connector (KCC) maps each Google Cloud Platform (GCP) resource to a different Custom Resource Definition (CRD) and controller. Historically, KCC legacy controllers (such as the Terraform-based controller and the DCL-based controller) followed a set of implicit naming conventions when determining the `Kind` of a KRM resource. 

As we transition to direct controllers and continue to introduce greenfield resources, it is critical to document, formalize, and standardize these naming rules. This ensures long-term consistency, maintains backward compatibility, and guides future development when designing new resource kinds.

## Naming Conventions

1. **Two-part Structure (Service Name + Specific Name):**
   The resource kind name is divided into two distinct parts:
   - **Service Portion:** Indicates the name of the GCP service that the resource belongs to.
   - **Resource Portion:** A specific name for that resource within the context of that service.
   
   *Example:* For GCP Compute Engine's health check, the kind is `ComputeHealthCheck`. The service portion is `Compute` and the resource portion is `HealthCheck`.

2. **Removal of the "Cloud" Prefix:**
   To keep kind names concise and avoid redundancy, any `"Cloud"` prefix is stripped from the service portion of the resource name.
   
   *Example:* For the service "Cloud SQL", the service portion is normalized to `"SQL"`, resulting in resource kinds like `SQLInstance` or `SQLDatabase` (rather than `CloudSQLInstance` or `CloudSQLDatabase`). Similarly, "Cloud Spanner" becomes `Spanner` (e.g., `SpannerInstance`).

3. **PascalCase and Singular Form:**
   - **KRM Kind:** Always written in **PascalCase** and is **singular** (e.g., `ComputeHealthCheck`, `StorageBucket`, `PubSubTopic`).
   - **GCP API Representation:** In contrast, the underlying GCP API typically uses camelCase, is dot-separated, and uses pluralized resource collections (e.g., `compute.healthChecks`, `pubsub.topics`).

4. **Using Acronym** :
   - If the full name is excessively long, use a clean, commonly recognized acronym (this requires human decision and reviewer discretion).
   - *Example:* "Identity and Access Management" -> `IAM`, "Google Kubernetes Engine" -> `GKE`, "Contact Center Insights" -> `CCInsights`.

## Service Portion Selection

When determining the **Service Portion** of a new resource kind name, there are two options from highest to lowest priority:

1. **Latest GCP official product name**:
   - Use the full, modern official GCP product name whenever available. 
   - *Example:* "Cloud Spanner" -> `Spanner`, "Cloud Build" -> `Build`, "Artifact Registry" -> `ArtifactRegistry`.
2. **Hostname of the REST API**:
   - In most cases, the official GCP product name maps to the lowercase of the REST API hostname (e.g., `spanner.googleapis.com` for Spanner, `redis.googleapis.com` for Redis).
   - Prioritize using GCP official product name whenever available if the brand/product name and API hostname differ.
   - Use the REST API host name as a backup.
   - *Example:* `container.googleapis.com` -> `Container`, `bigquery.googleapis.com` -> `BigQuery`.

## Standardization and Backward Compatibility

To prevent breaking existing users, KCC preserves these naming rules across both legacy and direct controllers. When introducing new KRM resources or migrating existing ones to direct controllers, developers **must** strictly adhere to these established rules.

## Checklist for Determining Resource Kind Names

When designing a new KCC resource, use the following checklist to determine its `Kind`:

1. [ ] **Identify the GCP Service and Apply Selection Hierarchy:**
   - Use the **Service Portion Selection** above to determine the service prefix.
   - Strip any `"Cloud"` prefix if using the official product name.
   - Determine the necessity to use acronym.
2. [ ] **Determine the Resource Portion:** What is the specific name of the resource within that service? Use the singular form. (e.g., "Instance", "Bucket", "Topic", "HealthCheck")
3. [ ] **Combine and PascalCase:** Join the Service portion and Resource portion using PascalCase. (e.g., `SpannerInstance`, `StorageBucket`, `ComputeHealthCheck`)
4. [ ] **Check for Collisions/Prior Art:** Check `config/crds/resources` or existing resources in `apis/` to ensure the prefix and service portion are consistent with existing family members.

## Evaluation of Existing Violations

To ensure backward compatibility and prevent breaking existing users, existing resource kinds that violate these rules are preserved. However, they are treated as legacy naming exceptions and should not be used as patterns for new resources.

### 1. Service Portion Retaining the "Cloud" Prefix
The following 27 legacy resource kinds incorrectly retain the `"Cloud"` prefix in their `Kind` name. Newly introduced resources must strip this prefix.

| Current Resource Kind | API Group | Version(s) | Expected Standardized Kind |
| :--- | :--- | :--- | :--- |
| **CloudAssetFolderFeed** | `cloudasset.cnrm.cloud.google.com` | `v1alpha1` | `AssetFolderFeed` |
| **CloudAssetOrganizationFeed** | `cloudasset.cnrm.cloud.google.com` | `v1alpha1` | `AssetOrganizationFeed` |
| **CloudAssetProjectFeed** | `cloudasset.cnrm.cloud.google.com` | `v1alpha1` | `AssetProjectFeed` |
| **CloudBatchResourceAllowance** | `batch.cnrm.cloud.google.com` | `v1alpha1` | `BatchResourceAllowance` |
| **CloudBuildConnection** | `cloudbuild.cnrm.cloud.google.com` | `v1alpha1` | `BuildConnection` |
| **CloudBuildTrigger** | `cloudbuild.cnrm.cloud.google.com` | `v1beta1` | `BuildTrigger` |
| **CloudBuildWorkerPool** | `cloudbuild.cnrm.cloud.google.com` | `v1beta1`, `v1alpha1` | `BuildWorkerPool` |
| **CloudDMSConversionWorkspace** | `clouddms.cnrm.cloud.google.com` | `v1alpha1` | `DMSConversionWorkspace` |
| **CloudDMSMigrationJob** | `clouddms.cnrm.cloud.google.com` | `v1alpha1` | `DMSMigrationJob` |
| **CloudDMSPrivateConnection** | `clouddms.cnrm.cloud.google.com` | `v1alpha1` | `DMSPrivateConnection` |
| **CloudDeployAutomation** | `clouddeploy.cnrm.cloud.google.com` | `v1alpha1` | `DeployAutomation` |
| **CloudDeployCustomTargetType** | `clouddeploy.cnrm.cloud.google.com` | `v1alpha1` | `DeployCustomTargetType` |
| **CloudDeployDeliveryPipeline** | `clouddeploy.cnrm.cloud.google.com` | `v1beta1`, `v1alpha1` | `DeployDeliveryPipeline` |
| **CloudDeployDeployPolicy** | `clouddeploy.cnrm.cloud.google.com` | `v1alpha1` | `DeployDeployPolicy` |
| **CloudDeployTarget** | `clouddeploy.cnrm.cloud.google.com` | `v1alpha1` | `DeployTarget` |
| **CloudFunctions2Function** | `cloudfunctions2.cnrm.cloud.google.com` | `v1alpha1` | `Functions2Function` |
| **CloudFunctionsFunction** | `cloudfunctions.cnrm.cloud.google.com` | `v1beta1` | `FunctionsFunction` |
| **CloudIDSEndpoint** | `cloudids.cnrm.cloud.google.com` | `v1beta1`, `v1alpha1` | `IDSEndpoint` |
| **CloudIOTDevice** | `cloudiot.cnrm.cloud.google.com` | `v1alpha1` | `IOTDevice` |
| **CloudIOTDeviceRegistry** | `cloudiot.cnrm.cloud.google.com` | `v1alpha1` | `IOTDeviceRegistry` |
| **CloudIdentityGroup** | `cloudidentity.cnrm.cloud.google.com` | `v1beta1` | `IdentityGroup` |
| **CloudIdentityMembership** | `cloudidentity.cnrm.cloud.google.com` | `v1beta1` | `IdentityMembership` |
| **CloudNumberRegistryRegistryBook**| `cloudnumberregistry.cnrm.cloud.google.com`| `v1alpha1` | `NumberRegistryRegistryBook` |
| **CloudSchedulerJob** | `cloudscheduler.cnrm.cloud.google.com` | `v1beta1` | `SchedulerJob` |
| **CloudSecurityComplianceCloudControl**| `cloudsecuritycompliance.cnrm.cloud.google.com`| `v1alpha1` | `SecurityComplianceCloudControl` |
| **CloudSecurityComplianceFramework**| `cloudsecuritycompliance.cnrm.cloud.google.com`| `v1alpha1` | `SecurityComplianceFramework` |
| **CloudTalentSolutionCompany** | `cloudtalentsolution.cnrm.cloud.google.com` | `v1alpha1` | `TalentSolutionCompany` |

### 2. Service Portion Names Different From API Hostnames

#### A. Brand vs. Underlying API Hostname (Different Names)
These services use a different market-facing brand name compared to the underlying REST API host:
- **Vertex AI** (`vertexai`) & **Vector Search** (`vectorsearch`) $\rightarrow$ use host **`aiplatform.googleapis.com`** (the unified machine learning platform).
- **Cloud Talent Solution** (`cloudtalentsolution`) $\rightarrow$ uses host **`jobs.googleapis.com`** (the jobs search and matchmaking engine).
- **Grafeas** (`grafeas`) $\rightarrow$ uses host **`containeranalysis.googleapis.com`** (the container scanning metadata storage API).
- **Database Migration Service** $\rightarrow$ uses host **`datamigration.googleapis.com`** (the Database Migration Service API).

#### B. Shorter or Variant Hostnames (Slight Variations)
These hostnames are derived from the product name but are shortened or use administration-focused suffixes:
- **Filestore** (`filestore`) $\rightarrow$ uses host **`file.googleapis.com`** (drops the "store" suffix).
- **BigQuery Analytics Hub** (`bigqueryanalyticshub`) $\rightarrow$ uses host **`analyticshub.googleapis.com`** (drops "BigQuery").
- **BigQuery BigLake** (`bigquerybiglake`) $\rightarrow$ uses host **`biglake.googleapis.com`** (drops "BigQuery").
- **Cloud Translation** (`translate`) $\rightarrow$ uses host **`translation.googleapis.com`** ("translation" instead of "translate").

#### C. Active Mismatches in the Codebase
- **NotebookInstanceV2** (`notebooks.cnrm.cloud.google.com` | `v1alpha1`): The API hostname is `notebooks`, the GCP official product name is "Workbench Instance" (using notebooks v2 API). We also have a "NotebookInstance" resource (v1beta1, using notebooks v1 API, deprecated).
- **CloudDMSWorkspace**, **CloudDMSMigrationJob**, and **CloudDMSPrivateConnection** (`clouddms.cnrm.cloud.google.com` | `v1alpha1`): The API hostname is `datamigration`, but the GCP official product name is "Database Migration Service".

### 3. Pluralized Settings Kinds
A few legacy resource kinds end in `"Settings"`, which is technically a plural noun. These are treated as acceptable domain-specific exceptions where no clear singular "Setting" noun maps cleanly to GCP API concepts.

- **IAPSettings** (`iap.cnrm.cloud.google.com` | `v1beta1`)
- **DialogflowSecuritySettings** (`dialogflow.cnrm.cloud.google.com` | `v1alpha1`)
- **APIQuotaAdjusterSettings** (`cloudquota.cnrm.cloud.google.com` | `v1beta1`)

### 4. Mismapped/Incomplete Resource Portion Names
Some legacy resource kinds drop words from their official GCP resource name, leading to incomplete or ambiguous KRM kinds:

- **DeviceStreamingSession** (`devicestreaming.cnrm.cloud.google.com` | `v1alpha1`): The underlying GCP resource is "Device Session". Under the Two-Part Structure rule (Service Portion `DeviceStreaming` + Resource Portion `DeviceSession`), this resource kind should be named **`DeviceStreamingDeviceSession`**.
