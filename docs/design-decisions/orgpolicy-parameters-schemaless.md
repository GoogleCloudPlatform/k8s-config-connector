# Design Decision: Schema Representation for `parameters` in `OrgPolicyPolicy`

## 1. Context and Problem Statement

In the `OrgPolicyPolicy` CRD, the `parameters` field allows users to pass specific configuration values when policy enforcement is enabled for a managed constraint.

The underlying Google Cloud API definition for these parameters encompasses a massive surface area. There are currently over **9,218 available fields** across all possible constraint types. Furthermore, these fields are highly conditional: a specific constraint (e.g., "Constraint A") only accepts a specific, small subset of these 9,218 fields.

The standard Config Connector (KCC) best practice is to define a strict OpenAPI structural schema for every field to enable synchronous `kubectl` validation. However, we must decide how to represent the `parameters` field in the Kubernetes CustomResourceDefinition (CRD) given its extreme size and highly polymorphic, conditional nature.

## 2. Decision

We will represent the `parameters` field as an arbitrary JSON object using `apiextensionsv1.JSON` in the Go API types, which translates to `x-kubernetes-preserve-unknown-fields: true` in the generated CRD manifest.

We will **not** attempt to generate a strict, declarative OpenAPI structural schema for the 9,218 conditional fields within the CRD.

## 3. Rationale

This decision is driven by hard physical constraints within the Kubernetes architecture (specifically `etcd` and the `kube-apiserver`).

### 3.1. Etcd Object Size Limits (The 1.5MB Limit)
Kubernetes stores CRDs in `etcd`, which enforces a strict 1.5 MB limit per object (including the gRPC payload overhead). 

To put this into perspective, Config Connector's largest CRD today is `MonitoringDashboard`. It defines approximately ~1,300 unique fields and its CRD YAML size is **~555 KB**. 

If we were to generate a strict schema for the `parameters` field containing 9,218 fields, the resulting CRD would be roughly 7 times larger, resulting in an estimated size of **~3.8 MB**. This would immediately cause the API Server to reject the installation of the `OrgPolicyPolicy` CRD entirely due to the `etcd` request size limit being exceeded.

### 3.2. Operational Delay for Missing Fields
An alternative approach would be to artificially truncate the schema—only defining the top 500 most commonly used parameter fields. However, if a user needs a parameter that is not in the truncated schema, they would be blocked. Adding support for a new field requires a code change and a new KCC release, which typically takes at least a couple of weeks. By keeping the field schemaless, users can immediately use any parameter supported by the underlying GCP API without waiting for a KCC release to explicitly whitelist it.

## 4. Pros and Cons

### Pros
*   **Installable CRD:** The `OrgPolicyPolicy` CRD remains well within the `etcd` size limits and installs smoothly on all standard Kubernetes clusters.
*   **Zero-Day Support for New Constraints:** As GCP adds new OrgPolicy constraints and parameters to the backend API, KCC users can immediately use them in their YAML without requiring KCC to update the CRD schema.
*   **Performance:** Controller and API Server memory profiles remain stable as they do not have to parse a massive schema tree.

### Cons
*   **No Synchronous Validation (Asynchronous Error Surfacing):** Because the Kubernetes API server does not validate the internal structure, typos, incorrect data types (e.g., passing a string instead of an integer), or missing required fields will not be caught during `kubectl apply`.
    *   *Mitigation:* The KCC controller will surface these errors asynchronously on the resource's `status.conditions` when the GCP API rejects the payload (`400 Bad Request`).
*   **Loss of IDE Support and Tooling:** Users lose native YAML autocomplete, intellisense, and inline documentation in editors (like VSCode or IntelliJ), as there is no OpenAPI schema for the IDE tooling to read.
*   **No Server-Side Pruning:** The Kubernetes API server normally drops unknown fields to prevent garbage data from consuming `etcd` space. With `preserve-unknown-fields`, any arbitrary data the user accidentally includes in the `parameters` block will be accepted and stored in `etcd`.

## 5. Alternatives Considered

### 5.1. Comprehensive OpenAPI Schema
We considered generating a full, strict OpenAPI schema for all 9,218 possible parameter fields across all OrgPolicy constraints. This would follow the standard KCC pattern for other resources.

*   **Result:** Rejected. As detailed in the Rationale (Section 3.1), the resulting CRD size (~3.8 MB) far exceeds the 1.5 MB `etcd` limit, making the CRD impossible to install on standard Kubernetes clusters.

### 5.2. Hybrid Side-by-Side Schema
We considered a hybrid model where the `parameters` object contains two sub-fields: `parameters.schema` (a strict OpenAPI object containing the ~500 most common fields) and `parameters.schemaless` (an arbitrary JSON object for everything else).

*   **Static Snapshot and Configuration Fragmentation:** To maintain KCC's own backward compatibility and avoid a maintenance treadmill, this schema would likely never be updated. This results in logical configuration being split across two different fields based solely on when those fields were added to the underlying GCP API.

### 5.3. Comparison of Approaches

| Feature | Pure Schemaless (Chosen) | Hybrid Side-by-Side | Comprehensive Schema |
| :--- | :--- | :--- | :--- |
| **User Complexity** | Simple ("Put JSON here") | Complex (Split configuration) | Simple (Fully typed) |
| **Typo Detection** | Asynchronous (GCP Error) | Sync for top 500 only | Synchronous (API Server) |
| **IDE Autocomplete** | None (requires tools) | Native for top 500 only | Full native support |
| **CRD Maintenance** | Zero | Mid (Initial curation) | Extremely High |
| **Breaking Changes** | None | Low (GCP stability) | Low (GCP stability) |
| **Installability** | **Success** (~5KB) | **Success** (~200KB) | **Failed** (~3.8MB) |

## 6. Conclusion: When to Use Schemaless Fields

While Config Connector's default stance remains to use strict OpenAPI schemas for all resources to maximize native Kubernetes UX, this design decision establishes a clear precedent for when to use `apiextensionsv1.JSON` (`x-kubernetes-preserve-unknown-fields: true`).

The decision to use a schemaless field should be based on two sets of criteria:

### 6.1. Drivers for Bypassing Strict Schema (The "Why Not" Drivers)
A strict schema should be avoided if:
1. **Physical Size Constraints:** A strict schema representation would cause the generated CRD to approach or exceed Kubernetes physical limits (e.g., the 1.5MB `etcd` limit).
2. **Massive or Polymorphic Surface Area:** The field acts as a container for thousands of potential sub-fields, or represents an inherently polymorphic payload where the structure is highly conditional.
3. **Zero-Day Support Requirement:** The backend API frequently adds new sub-fields, and users require immediate access to these features without being bottlenecked by KCC release cycles.

### 6.2. Factors Enabling a Schemaless Approach (The "Why Possible" Drivers)
A schemaless approach is viable if:
1. **Opaque Pass-Through:** The KCC controller does not need to internally inspect, default, or mutate the contents of the field to perform its reconciliation loop.
2. **Robust Backend Validation:** The underlying GCP API performs strict validation, ensuring that asynchronous error surfacing (via `status.conditions`) provides clear feedback to the user.
