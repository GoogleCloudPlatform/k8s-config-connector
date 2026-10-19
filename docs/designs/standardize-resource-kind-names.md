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

Product marketing names change over time and are open to interpretation. To make the **Service Portion** deterministic and reproducible, it is derived from the GCP service's **REST API Discovery Document**, which is the machine-readable source of truth that Google API client libraries are generated from.

### Source: the Discovery Document

The Discovery Document for a service can be fetched from either:
- `https://<service>.googleapis.com/$discovery/rest?version=<version>` (e.g. `https://spanner.googleapis.com/$discovery/rest?version=v1`), or
- the Discovery directory: `https://www.googleapis.com/discovery/v1/apis/<service>/<version>/rest`.

The following top-level fields are used:

- `canonicalName`: **Primary input.** The canonical, human-readable service name. Client library generators use it for service class names (e.g. Java `SQLAdmin`, `CloudRedis`). Optional; not every API sets it.
- `name`: **Fallback input.** The API name, which equals the REST API hostname prefix. Always present.
- `title`: **Casing hint.** Used to recover word boundaries and casing when `canonicalName` is all lowercase or concatenated.

### Algorithm

1. **Pick the base name:** Use `canonicalName` if present. Otherwise, fall back to `name`.
2. **Strip affixes:** Tokenize on whitespace. Remove a leading `Cloud` or `Google` token (see rule 2 above), and a trailing `Admin` or `API` token.
3. **Recover casing:** If the result is a single token whose casing is not meaningful (e.g. `Bigquery`, `Pubsub`, `Iam`, `Dns`, `networkconnectivity`, `Contactcenterinsights`), tokenize `title`. Strip non-alphanumeric characters (e.g. `Pub/Sub` -> `PubSub`), and drop the `Cloud`, `Google`, `API`, and `Admin` tokens. Then find the contiguous or ordered sequence of title tokens (including a parenthesized acronym such as `(IAM)`) whose case-insensitive concatenation equals the base name, and use the title's casing.
4. **PascalCase:** Upper-case the first letter of each remaining token and concatenate. Keep existing all-caps acronyms (`SQL`, `KMS`, `DNS`, `IAM`) as they are.
5. **Apply acronym rule (human decision):** If the result is excessively long, a reviewer may approve a commonly recognized acronym (see rule 4 above). This is the only non-deterministic step, and the decision must be recorded in the PR or design.
6. **Unresolvable cases:** If step 3 does not find a match, use the PascalCased base name and flag it for reviewer discretion.

### Examples

| API (`name/version`) | `canonicalName` | `title` | Derived Service Portion | Notes |
|---|---|---|---|---|
| `spanner/v1` | `Spanner` | Cloud Spanner API | `Spanner` | Used as-is |
| `sqladmin/v1` | `SQL Admin` | Cloud SQL Admin API | `SQL` | Trailing `Admin` stripped |
| `alloydb/v1` | `Cloud AlloyDB Admin` | AlloyDB API | `AlloyDB` | `Cloud` and `Admin` stripped |
| `redis/v1` | `Cloud Redis` | Google Cloud Memorystore for Redis API | `Redis` | `Cloud` stripped |
| `secretmanager/v1` | `Secret Manager` | Secret Manager API | `SecretManager` | PascalCased |
| `container/v1` | `Container` | Kubernetes Engine API | `Container` | `canonicalName` takes priority over marketing name |
| `bigquery/v2` | `Bigquery` | BigQuery API | `BigQuery` | Casing from `title` |
| `pubsub/v1` | `Pubsub` | Cloud Pub/Sub API | `PubSub` | Casing from `title` (`Pub/Sub` -> `PubSub`) |
| `iam/v1` | `Iam` | Identity and Access Management (IAM) API | `IAM` | Casing from parenthesized acronym in `title` |
| `networkconnectivity/v1` | `networkconnectivity` | Network Connectivity API | `NetworkConnectivity` | Word boundaries from `title` |
| `compute/v1` | *(absent)* | Compute Engine API | `Compute` | Fallback to `name` |
| `contactcenterinsights/v1` | `Contactcenterinsights` | Contact Center AI Insights API | `ContactCenterInsights` -> `CCInsights` | Word boundaries from `title` (skipping `AI`), then acronym rule (human decision) |

## Standardization and Backward Compatibility

To prevent breaking existing users, KCC preserves these naming rules across both legacy and direct controllers.
When introducing a new KRM API group, verify that there is no existing KRM API group mapping to the same underlying GCP REST API host. 
If one exists, strongly prefer adding the new resource kind to the existing group. Otherwise developers **must** strictly follow these established rules.

## Checklist for Determining Resource Kind Names

When designing a new KCC resource, use the following checklist to determine its `Kind`:

1. [ ] **Derive the Service Portion from the Discovery Document:**
   - Fetch the service's Discovery Document and read `canonicalName` (fall back to `name` if it is absent).
   - Apply the **Service Portion Selection** algorithm above: strip `Cloud`/`Google`/`Admin`/`API`, recover casing from `title`, and PascalCase.
   - Determine whether an acronym is needed (human decision), and record the rationale.
2. [ ] **Determine the Resource Portion:** Determine the specific name of the resource within that service and use the singular form. (e.g., "Instance", "Bucket", "Topic", "HealthCheck")
3. [ ] **Combine and PascalCase:** Join the Service portion and Resource portion using PascalCase. (e.g., `SpannerInstance`, `StorageBucket`, `ComputeHealthCheck`)
4. [ ] **Check for Collisions/Prior Art:** Check `config/crds/resources` or existing resources in `apis/` to ensure the prefix and service portion are consistent with existing family members.

## Evaluation of Existing Violations

For a detailed list and evaluation of these existing violations, please refer to the tracking issue: [Resource Kind Naming Violations](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/13347).

To ensure backward compatibility and prevent breaking existing users, existing v1beta1 resource kinds that violate these rules are preserved.
