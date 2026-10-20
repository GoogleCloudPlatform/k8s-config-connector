# Config Connector Core Principles

Config Connector (KCC) is a Kubernetes add-on that enables declarative, GitOps-driven management of Google Cloud Platform (GCP) resources through Kubernetes Resource Model (KRM) custom resources.

This guide provides a comprehensive and centralized reference for the four foundational pillars of Config Connector:
1. **[The Reconciliation Loop](#1-the-reconciliation-loop)**: How KCC observes, acquires, synchronizes, and cleans up cloud infrastructure.
2. **[Managing Annotations](#2-managing-annotations)**: A centralized catalog and practical usage guide for KCC annotations controlling identity, lifecycle, safety, and GitOps workflows.
3. **[Handling Drift](#3-handling-drift)**: How KCC detects out-of-band changes, manages dynamic defaults, and balances continuous reconciliation with selective drift tolerance.
4. **[Dependency Management](#4-dependency-management)**: How KCC resolves asynchronous dependency graphs, cross-resource references, and circular requirements.

---

## 1. The Reconciliation Loop

At its core, Config Connector functions as an ensemble of Kubernetes controllers that continuously reconcile desired state declared in KRM manifests (`spec`) against actual live infrastructure in Google Cloud APIs.

```
       ┌────────────────────────────────────────────────────────┐
       │                 Kubernetes API Server                  │
       │  (KRM Custom Resources: spec, metadata, status)       │
       └───────────────────────────┬────────────────────────────┘
                                   │ Watch / Informer Events
                                   ▼
       ┌────────────────────────────────────────────────────────┐
       │                Config Connector Reconciler             │
       │                                                        │
       │   1. Read KRM Spec & Metadata                          │
       │   2. Resolve Dependencies (Refs, Secrets, Hierarchy)   │
       │   3. Query Live GCP State (GET / Search)               │
       │   4. Compute Semantic Diff (Stateless or Stateful)     │
       │   5. Apply Mutations (Create / Update / Delete RPC)    │
       │   6. Update KRM Status & Ready Conditions              │
       └───────────────────────────┬────────────────────────────┘
                                   │
                                   ▼
       ┌────────────────────────────────────────────────────────┐
       │                   Google Cloud APIs                    │
       │       (compute, storage, container, sql, etc.)         │
       └────────────────────────────────────────────────────────┘
```

### 1.1 Architecture & Controller Routing

Reconciliation is performed by controller pods running either per-namespace (Namespaced Mode) or cluster-wide (Cluster Mode):

* **Namespaced Mode (Recommended for Multi-Tenant Clusters)**: Dedicated `cnrm-controller-manager` pods are deployed for each active Kubernetes namespace governed by a `ConfigConnectorContext` resource. Each manager authenticates via a dedicated Google Service Account (GSA) using Workload Identity.
* **Cluster Mode**: A single, central `cnrm-controller-manager` pod watches all namespaces and acts through a single cluster-level identity.

#### Controller Implementations: Direct vs. Legacy

Under the hood, KCC supports three internal controller implementations:
1. **Direct Controllers (Modern Standard)**: Written in native Go using official Google Cloud Go Client Libraries (`cloud.google.com/go/...`) and gRPC/REST APIs. They offer minimal memory footprint, rapid reconciliation cycles, detailed structured logging, and robust lifecycle management.
2. **Terraform-Based Controllers (Legacy)**: Wrap the Terraform Google Beta provider internally.
3. **DCL-Based Controllers (Legacy)**: Wrap Google's Declarative Client Library.

The reconciler is chosen using the following precedence:
1. *Resource Annotation (Deprecated)*: `cnrm.cloud.google.com/reconciler: direct` (or `alpha.cnrm.cloud.google.com/reconciler: direct`).
2. *`ConfigConnectorContext` Override*: Specifying controller overrides in `spec.experiments.controllerOverrides`.
3. *Static Configuration*: Built-in defaults in KCC that route resources to `direct`, `tf`, or `dcl`.

---

### 1.2 Reconciliation Lifecycle Phases

Every KCC resource progresses through distinct operational phases:

#### Phase A: Acquisition & Adoption
When a new KRM resource is applied to Kubernetes, the reconciler checks if the underlying GCP resource already exists before calling a creation API:
* **Creation**: If the GCP resource does not exist, KCC executes the GCP `Create` RPC.
* **Adoption (Acquisition)**: If a GCP resource with matching identifiers already exists (either identified by matching `metadata.name` or explicitly specified via `spec.resourceID`), KCC **adopts** the pre-existing GCP resource. It reconciles any discrepancies between the existing GCP configuration and the KRM `spec` without destroying or recreating the resource.

#### Phase B: Steady-State Polling & Periodic Re-reconciliation
Once a resource is reconciled, it enters steady state. Re-reconciliation is triggered by:
1. **Spec Updates**: Immediate reconciliation triggered by Kubernetes watch events when a user or GitOps agent modifies the resource.
2. **Periodic Polling**: Regular background polling to detect out-of-band changes (drift) in GCP.
   * By default, KCC schedules periodic re-reconciliation every **10 minutes** (with randomized jitter to prevent API thundering herds).
   * This interval can be adjusted per-resource using `cnrm.cloud.google.com/reconcile-interval-in-seconds`.
3. **Dependency Changes**: Updates to referenced resources (e.g., changes to an underlying `ComputeNetwork` or `Secret`).

#### Phase C: Updates & FieldMask Generation
When differences are detected:
* Direct controllers construct a granular `fieldmaskpb.FieldMask` containing only the specific top-level or nested fields that have changed.
* Update operations invoke the corresponding GCP `Update` or `Patch` API call.
* Asynchronous long-running operations (LROs) are polled until completion.

#### Phase D: Deletion & Finalizer Management
When a user deletes a KRM resource (`kubectl delete`):
1. Kubernetes sets `metadata.deletionTimestamp`.
2. KCC intercepts the deletion via its attached finalizers:
   * `cnrm.cloud.google.com/finalizer`: Coordinates cloud resource deletion.
   * `cnrm.cloud.google.com/deletion-defender`: Prevents deletion if other managed resources depend on this resource.
3. KCC inspects the `cnrm.cloud.google.com/deletion-policy` annotation:
   * **`delete` (Default)**: KCC calls the GCP `Delete` RPC to destroy the physical resource in Google Cloud, verifies deletion, and removes the finalizer.
   * **`abandon`**: KCC skips the GCP `Delete` RPC, preserving the physical resource in Google Cloud untouched, and simply removes the finalizer to unmanage it from Kubernetes.

---

### 1.3 Resource Status and Conditions

Every KCC resource exposes its lifecycle status under `.status.conditions`. The primary condition is `type: Ready`:

```yaml
status:
  conditions:
  - type: Ready
    status: "True"
    reason: UpToDate
    message: The resource is up to date
    lastTransitionTime: "2026-10-20T10:00:00Z"
  observedGeneration: 2
```

#### Condition Reason Reference Table

| Condition Reason | Status | Description | Action Required |
|---|---|---|---|
| `UpToDate` | `True` | The GCP resource exists and its configuration matches the declared KRM `spec`. | None. Resource is healthy. |
| `Creating` | `False` | Creation RPC or LRO is in progress in GCP. | Wait for operation to complete. |
| `Updating` | `False` | An update RPC or LRO is being applied in GCP. | Wait for operation to complete. |
| `Deleting` | `False` | Deletion RPC is running in GCP before finalizer removal. | Wait for teardown to finish. |
| `DependencyNotReady` | `False` | A referenced KCC resource exists but has not yet reached `UpToDate`. | Inspect the referenced resource's status. |
| `DependencyNotFound` | `False` | A referenced KCC resource or Kubernetes Secret cannot be located. | Verify the target resource name/namespace exists. |
| `DependencyInvalid` | `False` | A reference format or value is malformed. | Correct the reference field syntax. |
| `ManagementConflict`| `False` | Another KCC object or external system is managing the same GCP resource. | Resolve resource ownership collision. |
| `Unmanaged` | `False` | No controller is assigned to reconcile this resource (e.g. missing `ConfigConnectorContext`). | Ensure namespace is enabled with a valid CCC. |
| `CreateFailed` | `False` | The GCP creation API returned an error (e.g. quota, permissions, invalid spec). | Check `message` and GCP API logs; fix `spec`. |
| `UpdateFailed` | `False` | An update attempt failed (e.g. immutable field modified, invalid argument). | Revert or fix invalid spec changes. |
| `DeleteFailed` | `False` | The GCP deletion API failed (e.g. bucket not empty, dependent cloud bindings). | Clear blockers in GCP or set `deletion-policy: abandon`. |

---

### 1.4 Rate Limiting, Exponential Backoff & Retry Policies

When GCP API calls fail with transient errors (such as `429 Too Many Requests`, `503 Service Unavailable`, or transient `404` replication delays), KCC enters an exponential backoff loop:
* **Base Delay**: Starts at 1 second.
* **Ceiling**: Backoff doubles up to the configured maximum delay (default: 120 seconds).
* **Jitter**: Jitter is applied to prevent synchronized retry spikes.
* **Custom Backoff**: You can override the backoff ceiling on Direct controllers using `cnrm.cloud.google.com/backoff-max-delay-in-seconds`. Setting this to `"0"` disables automated retries for scenarios requiring manual intervention.

---

## 2. Managing Annotations

Annotations in Config Connector provide vital operational levers to steer controller behavior, lifecycle policies, identity mapping, and GitOps synchronization.

### 2.1 Complete Annotations Catalog

| Annotation Key | Allowed / Example Values | Default | Scope | Description |
|---|---|---|---|---|
| `cnrm.cloud.google.com/project-id` | `"my-gcp-project-id"` | Inherited from Namespace | Resource, Namespace | Specifies the GCP Project ID where the resource resides. Overrides the Kubernetes namespace name fallback. |
| `cnrm.cloud.google.com/folder-id` | `"123456789012"` | None | Resource, Namespace | Specifies the GCP Folder ID for folder-scoped resources. |
| `cnrm.cloud.google.com/organization-id` | `"987654321098"` | None | Resource, Namespace | Specifies the GCP Organization ID for organization-scoped resources. |
| `cnrm.cloud.google.com/state-into-spec` | `absent`, `merge` | `absent` (in modern releases) | Resource, Namespace, CCC | Controls whether GCP server-defaulted values are written back into the KRM `spec`. Set to `absent` in GitOps environments to prevent diff wars. |
| `cnrm.cloud.google.com/deletion-policy` | `delete`, `abandon` | `delete` | Resource, Namespace | Determines whether deleting the KRM object deletes the underlying GCP resource (`delete`) or merely unlinks KCC (`abandon`). |
| `cnrm.cloud.google.com/actuation-mode` | `paused` | Continuous | Resource, CCC | Temporarily suspends all reconciliation and GCP mutations for the resource without deleting it. |
| `cnrm.cloud.google.com/reconcile-interval-in-seconds` | `"300"`, `"3600"` | `"600"` (~10m) | Resource, Namespace | Overrides the periodic drift detection polling interval in seconds. Useful for reducing GCP API quota consumption. |
| `cnrm.cloud.google.com/backoff-max-delay-in-seconds` | `"30"`, `"0"` | `"120"` | Resource | Sets the maximum exponential retry backoff ceiling for Direct controllers. Set to `"0"` to halt automatic retry loops. |
| `cnrm.cloud.google.com/management-conflict-prevention-policy` | `resource`, `none` | `resource` | Resource, CCC | Controls conflict detection when multiple KCC resources point to the same GCP entity. |
| `cnrm.cloud.google.com/unmanaged` | `"field1,field2.subfield"` | None | Resource | Informs Direct controllers to skip drift detection and mutations on specified field paths, allowing external tools to manage them. |
| `cnrm.cloud.google.com/default-to-gcp-fields` | `"spec.storageClass,spec.retention"` | None | Resource | Directs controllers to preserve live GCP values for omitted fields without attempting to reset them to defaults. |
| `cnrm.cloud.google.com/force-destroy` | `"true"`, `"false"` | `"false"` | Resource | Safety gate for resources containing sub-items (e.g. `StorageBucket`). Must be `"true"` to permit deleting non-empty buckets. |
| `cnrm.cloud.google.com/delete-contents-on-destroy` | `"true"`, `"false"` | `"false"` | Resource | Similar to force-destroy for datasets and tables (e.g. `BigQueryDataset`). |
| `cnrm.cloud.google.com/reconciler` | `direct`, `tf`, `dcl` | Auto-selected | Resource (Deprecated) | Explicitly selects the reconciliation engine. Deprecated in favor of `ConfigConnectorContext` controller overrides. |
| `cnrm.cloud.google.com/sqlinstance-advanced-dr` | `"enabled"` | None | Resource (`SQLInstance`) | Enables Advanced Disaster Recovery (DR) pair support for Cloud SQL, suppressing false drift during out-of-band failover. |
| `cnrm.cloud.google.com/remove-default-node-pool` | `"true"`, `"false"` | `"false"` | Resource (`ContainerCluster`) | Automatically removes the default GKE compute node pool upon cluster creation so independent `ContainerNodePool` resources can be used. |
| `cnrm.cloud.google.com/create-gsa-key-secret` | `"true"`, `"false"` | `"true"` | Resource (`IAMServiceAccountKey`) | Enables or disables automatic creation of a Kubernetes Secret containing the private key. |

---

### 2.2 Deep Dive: Core Annotations & Common Workflows

#### 1. Safe GitOps Deployments: `state-into-spec: absent`
* **Problem**: In legacy KCC versions, KCC fetched defaulted fields from GCP (e.g. disk sizes, network tiers, protocol defaults) and wrote them directly back into the KRM `.spec`. If your Git repository did not contain those defaulted fields, GitOps tools (ArgoCD, Flux) immediately saw an out-of-sync diff and attempted to strip them, causing an endless GitOps reconciliation battle.
* **Solution**: Setting `cnrm.cloud.google.com/state-into-spec: absent` ensures KCC leaves the resource `spec` untouched, maintaining full compatibility with GitOps engines.

```yaml
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: my-team-assets
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  uniformBucketLevelAccess: true
```

#### 2. Infrastructure Protection: `deletion-policy: abandon`
* **Use Case**: Safeguard production databases, high-availability storage buckets, or shared VPC networks against accidental deletion during namespace cleanup or cluster migrations.
* **Behavior**: When `deletion-policy: abandon` is set, executing `kubectl delete` removes the KRM object from the cluster while leaving the physical cloud resource intact in Google Cloud.

```yaml
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: production-db-primary
  annotations:
    cnrm.cloud.google.com/deletion-policy: abandon
spec:
  databaseVersion: POSTGRES_15
  region: us-central1
  settings:
    tier: db-custom-8-32768
```

#### 3. Pausing Reconciliation: `actuation-mode: paused`
* **Use Case**: Emergency troubleshooting, conducting out-of-band maintenance via `gcloud`, or temporarily suspending API calls during an ongoing cloud incident.
* **Behavior**: The controller completely suspends calls to GCP for this resource. You can safely modify GCP configurations without KCC immediately reverting your changes. Once maintenance is done, remove the annotation to resume automated synchronization.

```yaml
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeInstance
metadata:
  name: edge-workstation-01
  annotations:
    cnrm.cloud.google.com/actuation-mode: paused
```

#### 4. Reducing API Quota Exhaustion: `reconcile-interval-in-seconds`
* **Use Case**: For large clusters managing thousands of static resources (e.g., DNS record sets, IAM roles, Service Accounts), default 10-minute polling can consume significant GCP API quota.
* **Behavior**: Extending the interval to several hours lowers API call volume dramatically while still preserving self-healing capabilities.

```yaml
apiVersion: dns.cnrm.cloud.google.com/v1beta1
kind: DNSRecordSet
metadata:
  name: api-endpoint-record
  annotations:
    cnrm.cloud.google.com/reconcile-interval-in-seconds: "7200" # Check drift every 2 hours
spec:
  ttl: 300
  type: A
  rrdatas: ["10.0.0.50"]
```

---

## 3. Handling Drift

Drift occurs when the real-world state of a cloud resource in Google Cloud diverges from the desired state declared in the KRM specification.

### 3.1 Sources of Drift
* **Out-of-Band Modifications**: An engineer executes `gcloud` commands or edits settings in the Google Cloud Console.
* **External Systems & CI/CD**: Third-party automation (Terraform, scripts, Ansible) mutates cloud infrastructure.
* **Autoscaling & Dynamic Cloud Behaviors**: Horizontal Pod Autoscalers, Cloud SQL automatic storage increases, or GKE node auto-provisioning mutate instance counts or disk allocations.
* **Server-Side Defaulting & Normalization**: GCP APIs automatically populating unconfigured default fields or reordering lists.

---

### 3.2 Detection Architecture: Stateless vs. Stateful

KCC employs two sophisticated drift detection strategies:

#### A. The 5-Step Comparison Pipeline (Direct Controllers)
For modern Direct controllers, KCC uses an automated 5-step comparison pipeline to prevent false positives and accurately compute update masks:

```
1. Actual Proto (Live GCP API GET)
   │
   ▼ Spec_FromProto
2. Actual KRM (Strips status, output-only fields, etags)
   │
   ▼ MergeUnsetFields (copies values for fields omitted from desired spec)
3. Desired KRM (Adopts unmanaged actual values)
   │
   ▼ Spec_ToProto
4. Normalized & Ordered Protos (Sorts repeated fields, normalizes project number -> project ID)
   │
   ▼ DiffForTopLevelFields
5. Structured Diff & Minimal FieldMask (Emitted for GCP Update RPC)
```

#### B. Stateful Reconciliation via Change Cookies
For highly complex resources with dynamic server defaults (such as `SQLInstance`), KCC supports stateful reconciliation using **Change Cookies** stored in `cnrm.cloud.google.com/last-changed-cookie`:
* **Fast Path**: The controller stores a cryptographic hash of the last successfully applied `spec` alongside a hash of the observed GCP state. On periodic re-reconciliation, if both the desired spec and the observed GCP hash match the cookie, the controller exits immediately without executing costly semantic diffs.
* **Slow Path**: If the observed GCP hash changes while the spec remains constant, KCC knows out-of-band drift has occurred and enters detailed reconciliation to restore the declared state.

---

### 3.3 Strategies for Managing Drift

Depending on your operational requirements, choose between **Active Correction**, **Selective Unmanagement**, and **Disaster Recovery Suppression**:

#### Strategy 1: Active Drift Correction (GitOps Self-Healing)
* **Goal**: Git repository is the absolute source of truth. Any manual alteration in Cloud Console must be undone.
* **Configuration**: Default KCC behavior. Leave fields declared in `spec`. KCC will detect out-of-band changes during its periodic cycle and automatically overwrite them to match Git.

#### Strategy 2: Intentional Out-of-Band Management (`unmanaged`)
* **Goal**: Manage foundational infrastructure with KCC, but delegate specific fields (such as Cloud SQL storage capacity or replica counts) to external autoscalers or external teams.
* **Configuration**: Use `cnrm.cloud.google.com/unmanaged` with a comma-separated list of field paths. KCC will ignore differences on those paths:

```yaml
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: scalable-db
  annotations:
    # Allow Cloud SQL automatic storage increase to manage disk size without KCC reverting it
    cnrm.cloud.google.com/unmanaged: "spec.settings.diskSize"
spec:
  databaseVersion: MYSQL_8_0
  region: us-central1
  settings:
    tier: db-custom-4-16384
    diskAutoresize: true
```

#### Strategy 3: Disaster Recovery Pair Drift Suppression
* **Problem**: Cross-region database replication (such as Cloud SQL Primary and DR Replica) requires mutual references. When an out-of-band failover occurs in GCP, the replica is promoted to primary, which normally causes KCC to detect a conflicting role and attempt destructive reversions.
* **Solution**: Set `cnrm.cloud.google.com/sqlinstance-advanced-dr: "enabled"` on both instances. This instructs the direct controller to allow circular-dependency-free initial setup and to suppress role reversion during and after live failover events.

---

## 4. Dependency Management

Real-world cloud architectures consist of interdependent resources: Compute instances require Subnetworks, Subnetworks require VPC Networks, Pub/Sub Subscriptions require Topics, and Services require IAM Service Accounts.

Config Connector provides a declarative dependency management engine that automatically resolves references and guarantees eventual consistency.

### 4.1 Reference Syntax and Patterns

Config Connector standardizes references using a `*Ref` suffix (e.g., `networkRef`, `topicRef`, `kmsKeyRef`, `serviceAccountRef`).

#### Pattern A: Intra-Namespace Reference (`name`)
When the referenced resource is managed by KCC within the **same Kubernetes namespace**:

```yaml
spec:
  topicRef:
    name: telemetry-events # Matches metadata.name of PubSubTopic in current namespace
```

#### Pattern B: Cross-Namespace Reference (`name` + `namespace`)
When the dependency resides in a **different Kubernetes namespace** within the same cluster:

```yaml
spec:
  networkRef:
    name: shared-vpc
    namespace: networking-infra # Explicit namespace of the target ComputeNetwork
```

#### Pattern C: External Reference (`external`)
When referencing a pre-existing GCP resource that is **not managed by Config Connector** in the current cluster:

```yaml
spec:
  # Accepts standard GCP URI, relative resource name, selfLink, or resource ID
  subnetworkRef:
    external: projects/corp-prod-net/regions/us-central1/subnetworks/app-tier-subnet
```

---

### 4.2 Asynchronous Resolution & Eventual Consistency

In traditional imperative scripts (or early infrastructure-as-code tools), dependencies required rigid ordering scripts (`wait-for-resource.sh`).

In Config Connector:
1. **Apply All at Once**: You can apply your entire infrastructure graph simultaneously (`kubectl apply -f infra/` or via GitOps sync).
2. **Non-Blocking Requeue**: If resource B references resource A, and resource A is still in `Creating` or has not yet been applied:
   * Controller B marks its condition as:
     ```yaml
     status:
       conditions:
       - type: Ready
         status: "False"
         reason: DependencyNotReady
         message: "Resource 'networking-infra/shared-vpc' is not ready"
     ```
   * Controller B places itself in an exponential backoff requeue loop.
3. **Automatic Convergence**: As soon as resource A reaches `status.conditions[Ready].status: "True"`, Controller B resolves the generated URI/ID from resource A's status and automatically proceeds with creation.

---

### 4.3 Container Hierarchies (Project, Folder, Organization)

Every GCP resource belongs to a resource hierarchy: `Organization -> Folder -> Project -> Resource`.

KCC resolves parent containers through three flexible mechanisms:
1. **Explicit Spec Reference**:
   ```yaml
   spec:
     projectRef:
       external: "my-gcp-project-id" # or name: "my-kcc-project"
   ```
2. **Hierarchical Annotations**:
   ```yaml
   metadata:
     annotations:
       cnrm.cloud.google.com/project-id: "my-gcp-project-id"
   ```
3. **Namespace Defaulting**: If neither spec reference nor annotation is specified, KCC automatically assumes the GCP Project ID matches the name of the Kubernetes namespace where the resource is deployed.

---

### 4.4 Secret and Credential References

When resources require sensitive data (e.g., database root passwords, TLS private keys, API tokens), Config Connector forbids hardcoding plain text in the manifest. Instead, use `valueFrom.secretKeyRef`:

```yaml
spec:
  rootPassword:
    valueFrom:
      secretKeyRef:
        name: db-credentials
        key: password
```

* **Secret Version Tracking**: Direct controllers compute a SHA-256 hash of referenced secret values and store it in `cnrm.cloud.google.com/observed-secret-versions`. If the Kubernetes Secret is rotated, KCC detects the hash divergence and triggers an update in GCP without exposing sensitive credentials in logs or status fields.

---

### 4.5 Deletion Safety & Cascading Protection

To prevent accidental outages caused by out-of-order deletions, Config Connector attaches the `cnrm.cloud.google.com/deletion-defender` finalizer to critical infrastructure components:
* If a user or automation attempts to delete a `ComputeNetwork` while a `ComputeSubnetwork` or `ComputeFirewall` is still actively referencing it, the deletion defender blocks removal of the network until all dependent child resources have been safely decommissioned.

---

## 5. Summary & Next Steps

By understanding how the **reconciliation loop** converges state, mastering **annotations** for operational control, architecting for **drift** detection and suppression, and structuring **declarative dependencies**, platform teams can operate multi-tenant Google Cloud infrastructure with high resilience and predictability.

To explore battle-tested, production-grade manifests that demonstrate these principles across diverse GCP services, continue to:
* **[Practical Complex Examples Guide](./practical-examples.md)**: Production configurations for Cloud SQL, GKE, Cloud Storage, Pub/Sub, Compute Engine, and IAM.
* **[Managing Config Connector with GitOps (ArgoCD)](./gitops-argocd.md)**: Deploying and upgrading Config Connector via GitOps pipelines.
