# Design Proposal: Bulk Deterministic CRD & Go Types Generation for KCC

**Author:** Senior SWE, Kubernetes Config Connector (KCC)  
**Status:** In Progress / Sandbox Experimentation  
**Branch:** `bulk-crd-gen-experiment`  
**Scope:** Phase 1 (CRD & Go Types Greenfield Bulk Generation)  

---

## 1. Executive Summary & Problem Statement

Kubernetes Config Connector (KCC) currently supports hundreds of GCP resources, but the surface area of Google Cloud APIs continues to expand rapidly. While our agentic greenfield pipeline can generate end-to-end direct controllers, the current workflow suffers from two primary bottlenecks:

1. **Review Fatigue on Mechanical Boilerplate:** Reviewing 1,000–3,000 lines of generated Go types per resource is slow and exhausting when 85–90% of the code consists of straightforward proto-to-Go mappings (primitive types, strings, integers, JSON tags, comments, and standard Spec/Status separation).
2. **AI Agent Hallucination & Token Inefficiency:** Prompting LLMs to generate entire `types.go` files from scratch introduces syntax errors, missing deepcopy annotations, and inconsistent conventions.

### Core Architecture: Deterministic-First with Native Go Overrides
**85–90% of KCC CRD definitions can be generated 100% deterministically from source-of-truth Protobuf schemas via `controllerbuilder generate-types`.**

Rather than inventing complex overlay DSLs, we strictly adhere to KCC's native Go file structure and greenfield skills:
- **`types.generated.go`**: 100% deterministic output generated directly from protobufs. Overwritten on every run.
- **`<kind>_types.go`**: Contains the root CRD structs (`Kind`, `KindSpec`, `KindStatus`) and **only the specific child structs that require manual/agent modifications** (such as swapping a raw string for a `*Ref`).
- **`<kind>_reference.go` / `<kind>_identity.go`**: Dedicated files for reference and identity implementations.
- **`controllerbuilder` Native Deduplication**: `controllerbuilder` automatically detects if a struct exists in `<kind>_types.go` and skips re-generating it in `types.generated.go`.

```mermaid
flowchart TD
    subgraph Step1["1. Deterministic Engine (controllerbuilder)"]
        P[Googleapis Protos / Descriptors] --> G[Deterministic Go AST Generator]
        G --> R1[Draft types.generated.go]
        G --> R2[Draft <kind>_types.go boilerplate]
    end

    subgraph Step2["2. Practical Health Checks & Gap Scanner"]
        R1 --> V1[controller-gen & CRD Builder]
        R1 --> V2[Reference Scanner / missingrefs]
        R2 --> V3[KCC Standards Validator]
        V1 --> S1[Inventory Report & Triage Status]
        V2 --> S1
        V3 --> S1
    end

    subgraph Step3["3. Targeted Agent / Developer Action"]
        S1 -->|Action: Add Ref| A1[Agent/Human moves struct to <kind>_types.go]
        S1 -->|Action: Add Ref| A2[Agent/Human implements <kind>_reference.go]
        S1 -->|Action: Fix CRD| A3[Fix schema recursion or deep nesting]
        S1 -->|Ready| PR[Immediate PR Ready]
    end
```

---

## 2. Practical Measurement: The 4 Concrete Health Checks

To objectively measure how much of a resource is correctly generated deterministically versus what requires human/AI attention, we run four practical checks:

```mermaid
flowchart LR
    C1["1. CRD Build Health<br>(controller-gen)"] --> R[Inventory Report]
    C2["2. Reference Completeness<br>(Proto String vs *Ref)"] --> R
    C3["3. Dependency Blockers<br>(Target Ref Exists?)"] --> R
    C4["4. KCC API Standards<br>(Labels, Generation, Enums)"] --> R
```

### 2.1 Check 1: CRD Build & Syntactic Health (Pass / Fail)
- Runs `dev/tasks/generate-crds` (invoking `controller-gen`).
- Verifies that `types.generated.go` compiles, deepcopy functions are generated, and OpenAPI v3 schemas are emitted without recursion errors or panics.

### 2.2 Check 2: Reference Completeness
- Scans proto field definitions for GCP resource path strings (e.g., `google.api.resource_reference`, fields ending in `_name`, `_link`, `_kms_key`, `_network`, `_service_account`).
- Checks if the corresponding Go field is mapped to a structured `*Ref` (e.g., `kmsKeyRef`) or left as raw `*string`.

### 2.3 Check 3: Dependency Graph & Identity Blockers
- When a resource requires a reference to another GCP resource $X$:
  - Checks if $X$ already has an Identity (`<x>_identity.go`) and Reference (`<x>_reference.go`) in KCC.
  - If not, flags the blocker so the agent/developer can scaffold an external-only reference for $X$.

### 2.4 Check 4: KCC Standard Validation Rules
- **Enums**: Must be unvalidated `*string` (strictly **no** `+kubebuilder:validation:Enum` markers). KCC avoids enum validation markers to prevent breaking client-side validation when GCP APIs add new enum variants.
- **Labels**: Contains required labels (`cnrm.cloud.google.com/managed-by-kcc=true`, `cnrm.cloud.google.com/system=true`, `stability-level`).
- **ObservedGeneration**: `status.observedGeneration` is present and typed as `*int64`.
- **Output-Only Separation**: Fields with `google.api.field_behavior = OUTPUT_ONLY` or server-generated fields (`etag`) land in `Status` / `ObservedState`, not `Spec`.

---

## 3. Addressing Gaps in Generation

### 3.1 Tier 1: Deterministic Engine Improvements (`controllerbuilder`)
Before involving AI agents or human intervention, we maximize what `controllerbuilder` handles:
1. **Proto Annotations Parsing:** Ingest `google.api.resource_reference` and `google.api.field_behavior` directly from descriptor sets.
2. **Boilerplate Scaffolding:** Automatically generate initial `<kind>_types.go` with standard KCC labels, `observedGeneration`, and proto annotations (`// +kcc:spec:proto=...`).
3. **Reference Scaffolding:** If `google.api.resource` URL templates exist in the proto, automatically scaffold initial `<kind>_identity.go` and external-only `<kind>_reference.go`.

### 3.2 Tier 2: Targeted Agent Interventions (Following Greenfield Skills)
When a field is flagged as needing a reference or manual adjustment:
1. Agent moves **only that specific struct** from `types.generated.go` to `<kind>_types.go`.
2. Agent edits the field to use the appropriate `*Ref` type (e.g. `pubsubv1beta1.PubSubTopicRef`, `refsv1beta1.ProjectRef`, or custom external-only `<kind>Ref`).
3. Agent implements `<kind>_reference.go` and `<kind>_identity.go` if the target reference is new.
4. `controllerbuilder` automatically skips generating the moved struct on subsequent runs, keeping the generated file clean.

---

## 4. Practical Tracking & Work Queue Process

We maintain a straightforward, actionable inventory matrix generated by a lightweight CLI script or Go benchmark command:

### Sample Inventory Matrix

| Service | Kind | Proto Message | CRD Builds? | Missing Refs | Dependency Blockers | Status / Action |
| :--- | :--- | :--- | :---: | :--- | :--- | :--- |
| `alloydb` | `AlloyDBCluster` | `...alloydb.v1beta.Cluster` | ✅ Yes | None | None | **Ready** |
| `alloydb` | `AlloyDBInstance` | `...alloydb.v1beta.Instance` | ✅ Yes | `kmsKeyName` | KMSCryptoKey (exists) | **Action: Add Ref** |
| `bigquery` | `BigQueryDataset` | `...bigquery.v2.Dataset` | ✅ Yes | None | None | **Ready** |
| `ces` | `CESAgent` | `...ces.v1alpha.Agent` | ❌ No | `network` | ComputeNetwork | **Action: Fix OpenAPI Recursion** |
| `spanner` | `SpannerDatabase` | `...spanner.v1.Database` | ✅ Yes | `encryptionConfig.kmsKeyName` | None | **Action: Add Ref** |

### Work Queue Lifecycle:
1. **Batch Generation**: Run `generate-types` across target services.
2. **Automated Health Check**: Run `controller-gen` and reference gap scanner.
3. **Triage**:
   - **Green (Ready)**: 100% clean deterministic generation, ready for PR.
   - **Yellow (Action: Add Ref)**: Assigned to AI Agent to implement external-only ref and move struct to `<kind>_types.go`.
   - **Red (Action: Fix Schema)**: Assigned to human/agent for recursive schema or deeply nested struct fixes.

---

## 5. Full Greenfield Scope & Programmatic Measurement Pipeline

To scale deterministic generation validation across the entire KCC codebase without relying on manual diff counting, we define the complete Greenfield universe and execute a **100% code-based measurement pipeline**.

### 5.1 The Complete Greenfield Resource Taxonomy (~310 Resources)

Greenfield resources in KCC are identified via two canonical sources:

1. **Class A: Pure Greenfield Direct Resources (Direct-Only in `static_config.go` and scaffolded in `apis/`):**
   - Resources where `SupportedControllers` contains strictly `[k8s.ReconcilerTypeDirect]` and their KRM Go types are canonically generated in `apis/<service>/`.
   - Examples: `VertexAIPipelineJob`, `APIGatewayAPI`, `ApigeeEndpointAttachment`, `BackupDRBackupVault`, `BigLakeTable`, `CloudBatchResourceAllowance`, `EdgeNetworkNetwork`, `GKEHubMembership`, `SecuredLandingZone`.
2. **Class B: In-Flight / Unregistered Greenfield in `apis/` (90 Resources):**
   - Greenfield resources implemented in `apis/<service>/` that do not yet have an entry in `pkg/controller/resourceconfig/static_config.go`.
   - Examples: `ApigeeRegistryAPI`, `DiscoveryEngineSearchEngine`, `DiscoveryEngineControl`, `TestingDeviceSession`, `DataformRepository`.

*Critical Taxonomy Boundary Discovered During Calibration:*
* **Brownfield Direct Migrations (Exclude):** Resources like `CloudBuildWorkerPool` and `ArtifactRegistryVPCSCConfig` have `DefaultController: Direct` in `static_config.go`, but their KRM types reside in `pkg/clients/generated/apis/` or use `--skip-scaffold-files` in `generate.sh`. These are migrated legacy resources and must not be treated as proto-scaffolded Greenfield types.
* **Dual-Mode & Legacy (Exclude):** 78 Brownfield resources (dual Direct + TF/DCL) and 226 Legacy resources (TF/DCL only) are preserved and excluded from Greenfield wipes.

### 5.2 The Automated Programmatic Measurement Pipeline (`dev/tasks/evaluate-greenfield-parity.py`)

```mermaid
flowchart TD
    A["1. Schema AST Baseline Snapshot<br/>Load existing OpenAPI v3 schemas from config/crds/resources/"] --> B["2. Greenfield Target Discovery<br/>Filter kinds scaffolded by apis/<service>/generate.sh"]
    B --> C["3. Greenfield Batch Wipe<br/>rm target <kind>_types.go and types.generated.go"]
    C --> D["4. Batch Canonical Generation<br/>Execute ./apis/<service>/generate.sh for all targets + generate-crds"]
    D --> E["5. Automated Schema AST Traversal<br/>Recursively compare property trees (Spec & Status)"]
    E --> F["6. Quantitative Metrics & Report<br/>Export .build/greenfield-parity-report.json & Markdown summary"]
    F --> G["7. Clean Restoration<br/>git checkout HEAD -- apis/ config/crds/resources/"]
```

The measurement engine performs automated schema AST comparisons:
- **Build Success Rate:** Evaluates whether `generate.sh` and `dev/tasks/generate-crds` exit `0`.
- **Property Tree Traversal:** Recursively navigates OpenAPI v3 schemas (`spec.versions[...].schema.openAPIV3Schema`) before and after regeneration.
- **Classification Categories:**
  - **Exact Deterministic Matches:** Properties where name, type, format, and description match 1:1.
  - **Resource Reference Overrides:** Properties transitioning from raw `type: string` to structured `*Ref` objects (`external`, `name`, `namespace`).
  - **Credential Overrides:** Properties transitioning from raw strings to `SecretRef` / `SecretKeySelector` (`password.valueFrom`).
  - **Parent Pattern Parity:** Validates whether root `projectRef`, `folderRef`, `organizationRef`, and `location` match.
  - **Upstream Proto Additions:** Identifies new fields present in latest GCP proto definitions not yet present in static baseline CRDs.
- **Reporting:** Emits structured JSON (`.build/greenfield-parity-report.json`) and populates `docs/designs/benchmark-results.md`.

