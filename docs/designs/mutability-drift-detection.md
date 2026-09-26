# Mutability Drift Detection: Comparing KCC Schemas with GCP Source of Truth

## Objective
Detect mutability drift by comparing the mutability of fields in Config Connector's (KCC) CRDs against the authoritative mutability defined in Google Cloud Platform (GCP) APIs.

## High-Level Plan

### Step 1: Identify Supported TF & DCL Resources
**Goal:** Create a definitive list of all KCC resources currently backed by Terraform or DCL controllers.

**Sub-steps:**
1.  **Locate Resource Configuration:** Identify the source of truth for resource configuration (e.g., `pkg/controller/resourceconfig/static_config.go`).
2.  **Extract Resource List:** Parse configuration to list all supported KCC resources.
3.  **Determine Controller Type:** Identify if each resource uses the Terraform or DCL controller.
4.  **Extract Resource Name Mapping:** Extract the mapping to the underlying Terraform resource name (e.g., `google_compute_instance`) or DCL resource name.

**Exit Criteria:**
A structured output (e.g., JSON or YAML) listing **all** supported TF-based or DCL-based KCC resources that support `v1beta1`. For each resource, the output will include:
1.  **Controller Type:** (TF or DCL).
2.  **Resource Name:** The underlying Terraform/DCL resource name.
3.  **Service HostName:** (Optional) The GCP Service API hostname (e.g., `compute.googleapis.com`).
4.  **API Endpoint:** (Optional) The REST Resource Name (e.g., `v1.projects.locations.clusters`).
5.  **Official REST API URL:** (Optional) The URL to the official REST API documentation.

### Step 2: Extract Current Mutability (from Snapshots)
**Goal:** Determine field mutability as defined in current KCC schemas (snapshots in `third_party`).

**Sub-steps:**
1.  **Locate Schema Files (TF):** Find the file in `third_party/github.com/hashicorp/terraform-provider-google-beta` defining the TF resource schema.
2.  **Locate Schema Files (DCL):** Find the file in `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library` defining the DCL resource schema.
3.  **Parse TF Schemas:** Extract the `ForceNew` property for every field.
4.  **Parse DCL Schemas:** Extract the `x-kubernetes-immutable` property for every field.
5.  **Create Unified Output:** Standardize output format (e.g., `Field: <name>, Immutable: <true/false>`).

**Exit Criteria:**
A structured output (e.g., JSON or YAML) containing the current KCC mutability status for **all** identified TF-based and DCL-based KCC resources. For each resource, the output will list fields, each with:
1.  **Field Name:** The name of the field.
2.  **Mutability Status:** `true` if immutable (`ForceNew` for TF or `x-kubernetes-immutable` for DCL), `false` otherwise.

### Step 3: Determine Actual Mutability (Source of Truth)
**Goal:** Determine the authoritative field mutability from GCP Protobuf definitions or alternative sources.

**Sub-steps:**
1.  **Map KCC Resource to GCP Resource (with Audit Logging):**
    *   **Input:** Metadata from Step 1.
    *   **Action:** Identify the specific GCP resource (e.g., `compute.instances`) and its **Protobuf Definition** in `googleapis`.
    *   **Audit Output:** Log the mapping decision (e.g., "Mapped 'ComputeInstance' to 'compute.v1.Instance' based on kind match").
2.  **Determine Primary Source (Protobuf Definition):**
    *   **Action:** Retrieve Protobuf definition from `googleapis`.
    *   **Goal:** Extract `field_behavior` annotation (specifically `IMMUTABLE`).
    *   **Initial Extraction:** Create preliminary list of immutable fields.
3.  **Evaluate Source Reliability (Heuristic Check):**
    *   **Action:** Compare count of immutable fields in KCC (Step 2) vs. Protobuf.
    *   **Heuristic:** If `Count(KCC Immutable) >> Count(Protobuf Immutable)`, flag Protobuf as likely incomplete.
    *   **Decision:** Log result (e.g., "Protobuf insufficient, proceeding to deeper analysis").
4.  **Deep Analysis (If Protobuf Insufficient):**
    *   **Check Update Methods:** Analyze `Update`/`Patch` RPC methods. Fields absent from update masks/bodies are effectively immutable.
    *   **Internal Code Analysis (Advanced):** Fallback to manual/semi-automated analysis of internal implementations if needed.
5.  **Create Normalized Output:**
    *   **Output:** Structured list of fields with *authoritative* mutability status.

**Exit Criteria:**
A structured output (e.g., JSON or YAML) containing the authoritative mutability status for **all** corresponding GCP resources of identified TF-based and DCL-based KCC resources. For each GCP resource, the output will list fields, each with:
1.  **Field Name:** The name of the field.
2.  **Mutability Status:** `true` if immutable, `false` otherwise.
3.  **Source/Confidence Information:** Details on how the mutability was determined (e.g., "protobuf_annotation", "inferred_from_update_methods").

### Step 4: Compare and Report Discrepancies
**Goal:** Map fields between KCC and GCP and report mutability differences.

**Sub-steps:**
1.  **Establish Field Mapping Strategy (Rules & Special Cases):**
    *   **Rules:** `snake_case` (API) <-> `camelCase` (CRD); `[fieldName]Ref` (CRD) <-> `field_name` (API).
    *   **Special Cases:** Configurable overrides for non-standard mappings.
2.  **Iterative Field Mapping and Mutability Comparison:**
    *   **Action:** For each resource, map fields using rules.
    *   **Audit:** Log applied mapping rules and unmapped fields.
    *   **Compare:**
        *   **Feature Gap:** KCC `Immutable` vs GCP `Mutable`.
        *   **Mutability Discrepancy:** KCC `Mutable` vs GCP `Immutable` (Note: might be intentional).
        *   **Match:** Statuses agree.
3.  **Generate Detailed Report per Resource:**
    *   **Format:** Structured report (Markdown/JSON).
    *   **Content:**
        *   Mapping stats (mapped vs unmapped counts).
        *   List of Discrepancies (Feature Gaps & Mutability Discrepancies).
        *   List of Unmapped Fields (CRD-only & API-only) to guide rule refinement.
4.  **Refinement Loop:**
    *   Refine mapping rules based on "Unmapped Fields" findings.

**Exit Criteria:**
A comprehensive report (e.g., Markdown or JSON) detailing the comparison for **all** identified TF-based and DCL-based KCC resources. For each resource, the report will include:
1.  **Field Mapping Summary:** Statistics on mapped, unmapped (CRD-only), and unmapped (API-only) fields.
2.  **Mutability Discrepancies:** A list of fields where KCC's mutability differs from GCP's authoritative mutability, categorized as "Feature Gap" (KCC immutable, GCP mutable) or "Mutability Discrepancy" (KCC mutable, GCP immutable).
3.  **Unmapped Fields:** A clear listing of fields from both KCC CRDs and GCP APIs that could not be automatically mapped.
