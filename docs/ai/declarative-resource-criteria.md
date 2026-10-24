# Criteria for Declarative vs. Non-Declarative GCP Resources in Config Connector

Config Connector (KCC) manages Google Cloud Platform (GCP) resources using the Kubernetes Resource Model (KRM): a user declares the desired state of a resource in `.spec`, and a level-triggered controller continuously reconciles GCP to match `.spec`.

Because GCP's [AIP-121](https://google.aip.dev/121) and [AIP-131–135](https://google.aip.dev/131) ("Resource-oriented design") model nearly *every* API concept as a REST/gRPC resource with `Get`, `Create`, `Update`, and `Delete` methods, **having CRUD methods in a protobuf service is necessary but not sufficient** to make a GCP resource suitable for Config Connector. For example, one-shot job executions, live chat sessions, pull request comments, and individual database rows often expose standard AIP CRUD methods, even though managing them as Kubernetes CRDs in `etcd` is an anti-pattern.

To avoid subjective debates ("is this resource declarative?"), Config Connector evaluates candidate GCP resources across **three ordered criteria**:

1. **Criterion 1: Mechanical Reconcilability (Objective — API Contract)**
2. **Criterion 2: Parent & Lifecycle Longevity (Structural — URL Pattern & State Machine)**
3. **Criterion 3: Control Plane vs. Data Plane / Telemetry (Policy — Scope & Cardinality)**

---

## Criterion 1: Mechanical Reconcilability (Objective)

A KRM level-triggered reconciliation loop requires four mechanical operations:
1. **Observe (`FindByName` / `Get`):** Read the current state of the resource from GCP.
2. **Actuate Creation (`Create`):** Provision the resource from `.spec` when missing and bind it to a stable identity (`metadata.name` / `.spec.resourceID` or `.status.externalRef`).
3. **Steady-State Convergence (`Update` or Immutable No-Op):** When re-reconciled periodically (e.g., every 10 minutes) with an unchanged `.spec`, remain in a steady `Ready` / `UpToDate` state without mutating or re-creating the resource.
4. **Actuate Deletion (`Delete` or Abandon):** Clean up the resource when the KRM object is deleted (or cleanly abandon it for destroy-protected resources such as `KMSKeyRing`).

A GCP resource fails **Criterion 1** if any of the following hold in its protobuf service definition (`pkg/protoapi`):

### 1.1 Read-Only / Discovered / Server-Populated (`!Create && !Update`)
- **Signal:** The protobuf service defines `Get` and/or `List` (and occasionally `Delete`), but **neither `Create` nor `Update`**.
- **Why it fails:** The resource is populated out-of-band by GCP discovery or server-side processes. A KRM controller has no RPC to provision or mutate `.spec`.
- **Examples:** `AppHubDiscoveredService`, `AppHubDiscoveredWorkload`, `BatchTask`, `EdgeContainerMachine`, `APIHubPlugin`, `ColabRuntime`, `DocumentAIProcessorVersion`.
- **Disposition:** **Do not support** as a KRM CRD.

### 1.2 One-Shot Executions & Imperative Jobs (Violates Steady-State Convergence)
- **Signal:**
  - The resource has `Create` and `Get`, lacks `Update` (`!Update`), and often lacks `Delete` (`Create`-only) or exposes a custom **`Cancel<Resource>`** RPC instead of `Update<Resource>`.
  - Its `State` enum transitions automatically to a terminal completion state (`SUCCEEDED`, `FAILED`, `CANCELLED`, `COMPLETED`, `ABORTED`, `EXPIRED`) rather than holding a steady `ACTIVE` / `READY` state, and the message carries output-only `start_time` + (`end_time` or `completion_time`) timestamps.
- **Why it fails:** Re-reconciling a finished or failed one-shot execution cannot converge `.spec` via `Update`.
- **Distinction — Job Executions vs. Declarative Job Templates:**
  - **One-shot executions (Do NOT support):** `WorkflowsExecution`, `VertexAICustomJob`, `VertexAIPipelineJob`, `VertexAITuningJob`, `NotebooksExecution`, `DataplexMetadataJob`, `BatchJob`, `TranscoderJob`, `MetastoreMetadataImport`, `KMSImportJob`.
  - **Reusable / scheduled job definitions (Support):** Resources that configure a reusable template or schedule (`RunJob` [Cloud Run Job], `DataplexTask`, `DataformWorkflowConfig`, `CloudSchedulerJob`). In protobuf descriptors, these expose `Create` + `Update` + `Delete` for the template itself, plus a separate custom RPC (`Run<Resource>`, `Trigger<Resource>`, `Start<Resource>`) to spawn executions.

### 1.3 Append-Only / Destroy-Protected Infrastructure (`Create` + `Update`, No `Delete`)
- **Signal:** The service defines `Get`, `Create`, and `Update`, but no `Delete`.
- **Disposition:** **Requires human judgement.**
  - Some core GCP infrastructure resources intentionally omit `Delete` for security/audit reasons (e.g., `KMSKeyRing`, `KMSCryptoKey`, `KMSEkmConnection`, `IAPBrand`, `ChronicleReferenceList`, `StorageRapidCache`) and **are** supported in KCC (with deletion defaulting to abandon or a custom destroy action).
  - Others omit `Delete` because they are ephemeral sessions or activity records (`DeviceStreamingSession`, `TestingDeviceSession`, `SecuresourcemanagerPullRequest`, `DialogflowParticipant`) and **should not** be supported.

---

## Criterion 2: Parent & Lifecycle Longevity (Structural)

Even when a resource exposes full AIP CRUD (`Get`, `Create`, `Update`, `Delete`), a KRM object stored in Kubernetes `etcd` must represent durable state attached to a long-lived parent (a Project, Folder, Organization, BillingAccount, or another declarative KCC CRD).

A GCP resource fails **Criterion 2** if its `google.api.resource.pattern` is scoped under (or is itself) an **ephemeral runtime parent**:

| Ephemeral URL Pattern Segment | Meaning | Examples |
|---|---|---|
| `/sessions/{session}` | Live end-user chat or reasoning session | `AIPlatformSession`, `DialogflowContext`, `DialogflowSessionEntityType`, `DiscoveryEngineSession` |
| `/conversations/{conversation}` | Active user support or chat conversation | `DialogflowParticipant`, `CCInsightsConversation`, `DiscoveryEngineConversation`, `GeminiDataAnalyticsConversation` |
| `/runs/{run}` or `/executions/{execution}` | Individual training, pipeline, or lineage execution run | `DatalineageRun`, `AIPlatformTensorboardTimeSeries`, `AIPlatformTensorboardRun` |
| `/pullRequests/{pull_request}` or `/issues/{issue}` | Developer code-review or issue-tracker activity | `SecuresourcemanagerPullRequest`, `SecuresourcemanagerPullRequestComment`, `SecuresourcemanagerIssueComment` |

- **Disposition:** **Do not support** by default. If a parent segment in `google.api.resource.pattern` is an ephemeral runtime entity that KCC does not declaratively manage, the child resource cannot be meaningfully managed via GitOps.

---

## Criterion 3: Control Plane vs. Data Plane / Telemetry (Policy)

Some GCP resources pass both Criterion 1 (full AIP CRUD) and Criterion 2 (long-lived parent such as `Instance`, `MetadataStore`, `DataStore`, `Catalog`, or `Repository`), meaning a generated KCC controller **will mechanically work**. However, they represent **high-cardinality application data, telemetry/lineage records, or point-in-time version snapshots** rather than cloud infrastructure configuration.

Because this boundary is a product/scope policy rather than a mechanical API limitation, these resources are grouped into three sub-buckets for explicit review:

### 3.1 High-Cardinality Data-Plane Rows, Documents, SKUs & Telemetry Records
- **Characteristics:** The parent resource (`DataTable`, `DataStore`, `Catalog`, `KnowledgeBase`, `MetadataStore`) is control-plane infrastructure, but the child resource represents individual records, documents, catalog items, or scan occurrences inside that store.
- **Examples:**
  - **Table Rows / Search Documents / Catalog Items:** `ChronicleDataTableRow` (rows inside a `ChronicleDataTable`), `DiscoveryengineDocument` (documents inside a `DataStore`), `DialogflowDocument` (documents inside a `KnowledgeBase`), `RetailProduct` (SKUs inside a `Catalog`).
  - **ML & Vulnerability Lineage / Telemetry Records:** `AIPlatformArtifact`, `AIPlatformContext`, `AIPlatformExecution` (ML Metadata lineage records inside a `MetadataStore`), `GrafeasOccurrence` (per-image scan/build occurrences).
- **Disposition:** **Exclude from bulk infrastructure PRs.** Manage data-plane contents via data pipelines or SDKs rather than Kubernetes `etcd` CRDs, unless a specific KCC user story requires declarative seeding.

### 3.2 Immutable Version Snapshots, Image Tags, Experiments & Test Cases
- **Characteristics:** Resources that mint immutable version snapshots of another resource (`Version`, `DatasetVersion`), mutable pointers to build artifacts (`Tag`), or evaluation/testing fixtures (`Experiment`, `TestCase`, `Example`).
- **Examples:** `AIPlatformDatasetVersion`, `ArtifactRegistryTag`, `DialogflowVersion`, `DialogflowExperiment`, `DialogflowTestCase`, `DialogflowExample`.
- **Disposition:** **Isolate for product review.** Usually managed by CI/CD or MLOps pipelines rather than infrastructure GitOps.

### 3.3 Point-in-Time Backups & Volume Snapshots
- **Characteristics:** Point-in-time data captures of stateful infrastructure (`Backup`, `Snapshot`). Contrast with declarative backup *policies/vaults/schedules* (`NetAppBackupPolicy`, `NetAppBackupVault`), which are pure control-plane configuration.
- **Examples:** `SpannerBackup`, `NetappBackup`, `NetappSnapshot` (alongside existing KCC resources such as `ComputeSnapshot`, `AlloyDBBackup`, and `FilestoreBackup`).
- **Disposition:** **Supported when needed for DR/GitOps workflows, but review separately** from pure control-plane configuration CRDs because backup/snapshot creation often requires a live source volume/database and long-running data copy operations.

---

## Automated Enforcement in `controllerbuilder`

`controllerbuilder generate-types` (`dev/tools/controllerbuilder/pkg/commands/generatetypes/nondeclarative.go`) enforces Criteria 1–3 automatically when scaffolding a new resource (where `<kind>_types.go` does not yet exist) using a hybrid of dynamic proto RPC checks and explicit collection maps:

1. **Dynamic Proto RPC Checks (`nonDeclarativeRPCReason`):**
   - Rejects read-only / discovered resources whose service defines `Get` or `Delete` but neither `Create` nor `Update` (Criterion 1.1).
   - Rejects one-shot executions whose service defines `Create` without `Update` and either lacks `Delete` or exposes a `Cancel<Resource>` RPC (Criterion 1.2).
   - Rejects high-cardinality data-plane collections whose service exposes AIP-165 / AIP-231–235 bulk mutation or purge RPCs (`Purge*`, `BulkCreate*`, `BatchCreate*`, `BatchDelete*`, `DeleteAll*`) on the collection (Criterion 3.1).
2. **Explicit Collection Maps (`nonDeclarativePatternReason`):**
   - Because ~35% of non-declarative or policy-excluded resources (`ArtifactRegistryTag`, `AIPlatformDatasetVersion`, `DialogflowExample`, `SecuresourcemanagerIssueComment`, `NetappBackup`, `NetappSnapshot`, `SpannerBackup`, etc.) expose standard `Create`/`Get`/`Update`/`Delete` RPCs in proto, `nondeclarative.go` also checks `ephemeralCollectionSegments` (Criterion 2, matched anywhere in `google.api.resource.pattern`) and `nonDeclarativeLeafCollections` (Criterion 3, matched only on the leaf collection segment so ancestor `versions` segments like `ApihubSpec` are not false-flagged), with a per-entry comment documenting the concrete GCP resource and rationale.
   - Pass `--allow-non-declarative` to `controllerbuilder generate-types` to explicitly opt in when scaffolding a resource in one of these categories.
   - Existing resources (whose `<kind>_types.go` file already exists on disk) are never blocked during regeneration (`./dev/tasks/generate-crds`).

## Summary Checklist for Bulk Generation & Archetype Triage

Before advancing a resource through `kcc-bulk-resource-generator` or `controllerbuilder`:

1. **Check RPC Shape (`report-archetypes`):**
   - If `!Create && !Update` $\rightarrow$ **Reject** (Criterion 1.1: Read-Only / Discovered).
   - If `!Update` and (`!Delete` or `Cancel<Resource>` RPC or terminal `SUCCEEDED`/`FAILED` state) $\rightarrow$ **Reject** (Criterion 1.2: One-Shot Execution).
2. **Check Resource URL Pattern (`google.api.resource.pattern`):**
   - If pattern contains `/sessions/`, `/conversations/`, `/runs/`, `/pullRequests/`, or `/issues/` $\rightarrow$ **Reject / Split** (Criterion 2: Ephemeral Parent Scope).
3. **Check Control-Plane vs. Data-Plane Role:**
   - If the resource represents table rows, indexed documents, catalog SKUs, lineage/scan occurrences, version snapshots, or point-in-time backups $\rightarrow$ **Split out of core control-plane batches** for explicit scope sign-off (Criterion 3).

