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

To prevent breaking existing users, KCC preserves these naming rules across both legacy and direct controllers.
When introducing a new KRM API group, verify that there is no existing KRM API group mapping to the same underlying GCP REST API host. 
If one exists, strongly prefer adding the new resource kind to the existing group. Otherwise developers **must** strictly follow these established rules.

## Checklist for Determining Resource Kind Names

When designing a new KCC resource, use the following checklist to determine its `Kind`:

1. [ ] **Identify the GCP Service and Apply Selection Hierarchy:**
   - Use the **Service Portion Selection** above to determine the service prefix.
   - Strip any `"Cloud"` prefix if using the official product name.
   - Determine the necessity to use acronym.
2. [ ] **Determine the Resource Portion:** Determine the specific name of the resource within that service and use the singular form. (e.g., "Instance", "Bucket", "Topic", "HealthCheck")
3. [ ] **Combine and PascalCase:** Join the Service portion and Resource portion using PascalCase. (e.g., `SpannerInstance`, `StorageBucket`, `ComputeHealthCheck`)
4. [ ] **Check for Collisions/Prior Art:** Check `config/crds/resources` or existing resources in `apis/` to ensure the prefix and service portion are consistent with existing family members.

## Evaluation of Existing Violations

For a detailed list and evaluation of these existing violations, please refer to the tracking issue: [Resource Kind Naming Violations](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/13347).

To ensure backward compatibility and prevent breaking existing users, existing v1beta1 resource kinds that violate these rules are preserved.
