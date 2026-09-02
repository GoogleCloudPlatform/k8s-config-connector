# KCC Bulk CRD Generation: Architectural Learnings & Corrections

**Date:** 2026-09-02  
**Context:** Bulk Deterministic CRD & Go Types Generation Experiment  
**Branch:** `bulk-crd-gen-experiment`  
**Purpose:** Document key domain rules, architectural insights, and critical course corrections provided during the design and evaluation sessions to ensure alignment with KCC core principles.

---

## 1. Enum Validation & Forward-Compatibility Policy

### Initial Flawed Proposal
* Proposed generating Go string types with Kubebuilder enum validation markers (`+kubebuilder:validation:Enum=VAL1;VAL2`).

### The Correction & Principle
* **Enums in KCC must always be unvalidated `*string` types with strictly NO `+kubebuilder:validation:Enum` markers.**
* **Why:**
  1. **External Ownership & Asynchronous Evolution:** GCP APIs are developed and evolved by independent service teams across Google who frequently add new enum values.
  2. **Preventing Premature Client Rejection:** If KCC enforces enum values at the CRD OpenAPI level, Kubernetes will reject user manifests containing new enum values until KCC releases a CRD update.
  3. **Server-Side Authority:** The underlying GCP service API already validates enum values and returns standard `400 Bad Request` or `404` errors if an invalid string is passed.
  4. **Documentation Only:** Allowed enum values should be documented in Go docstrings/comments only, never as schema constraints.

---

## 2. Rejecting Overlay DSLs in Favor of KCC's Native Override Model

### Initial Flawed Proposal
* Proposed an elaborate YAML overlay DSL (`overlays.yaml`), custom AST patch engines, and abstract mathematical formulas to modify generated types.

### The Correction & Principle
* **Overlays are unnecessary over-engineering. KCC already has a native Go override mechanism:**
  1. **`types.generated.go`:** 100% deterministic output from protobufs, overwritten on every run.
  2. **`<kind>_types.go`:** Holds the top-level CRD types (`Spec`, `Status`) and **only the specific child structs that require manual/agent modifications** (such as swapping a raw string for a `*Ref`).
  3. **Native Deduplication in Tooling:** `controllerbuilder` automatically checks if a struct exists in `<kind>_types.go` and **skips** generating it in `types.generated.go`. Child and parent structs link naturally via standard Go compilation.
  4. **Isolated Files:** Identity logic lives in `<kind>_identity.go` and reference definitions live in `<kind>_reference.go`.
* **Workflow Alignment:** AI agents and developers should follow the standard `kcc-direct-greenfield-types-implementer` and `kcc-direct-base-types-implementer` skills directly in Go, rather than writing custom YAML overlays.

---

## 3. Realities of GCP Proto Annotations & Fallbacks

### Initial Flawed Proposal
* Assumed `google.api.resource_reference` was widely available and could serve as the primary mechanism for detecting reference fields.
* Assumed `google.api.resource` was consistently present on all GCP resource protos.

### The Correction & Principle
1. **`google.api.resource_reference` is Sparse (< 10%):**
   * Populated in less than 10% of GCP APIs. Relying on it as a primary signal misses >90% of references.
   * **Real Strategy:** Use multi-signal heuristics (field names like `*_link`, `*_network`, `*_kms_key`, `*_service_account`, `*_topic` + docstring URI regexes) and flag remaining ambiguous strings for AI agent / developer review.
2. **`google.api.resource` is Canonical but Incomplete:**
   * When `(google.api.resource).pattern` is present, it is the canonical source for extracting parent hierarchy (`projects/`, `folders/`, `organizations/`) and location.
   * **Mandatory Fallback:** Because many GCP protos omit this option, the generator must have a robust fallback (inspecting `parent` docstring conventions or defaulting to standard `ProjectRef` + `Location`).

---

## 4. Classification of "Credential Gaps" as Resource References

### Initial Flawed Formulation
* Treated "Credential Gaps" as a separate, distinct foundational gap category alongside Reference Gaps and Parent Gaps.

### The Correction & Principle
* **Credential fields are not an independent fundamental gap; they are simply a specific subtype of Resource Reference (`SecretRef` / `SecretKeySelector`).**
* In Kubernetes and KCC, sensitive fields (passwords, TLS private keys, certificates, API tokens) are modeled as references to Kubernetes `Secret` resources (`v1alpha1.SecretRef` / `v1alpha1.SecretKeySelector`) or GCP Secret Manager secrets.
* Recognizing sensitive field patterns (e.g., `*_password`, `*_private_key`, `*_secret`) during generation is still useful for suggesting `SecretKeySelector` references, but architecturally it falls under the unified Resource Reference framework rather than a separate abstraction layer.

---

## 5. Planning Ahead for Validation: The Ground-Truth Testing Failure

### The Planning Failure
* **Failed to plan ahead on how to test and validate results before producing/reporting accuracy numbers.**
* When developing generator enhancements, initial coverage metrics were derived from ad-hoc synthetic AST parsing scripts and direct subcommands rather than establishing an end-to-end, reproducible regeneration validation test harness against `master`.
* Reporting accuracy data without a pre-agreed, reproducible ground-truth harness led to premature assertions and lack of confidence in the metrics.

### The True Ground-Truth Regeneration & Diff Methodology
To measure deterministic generation fidelity with 100% certainty:

#### Why Avoid Per-Service Resets Against Master:
* If you switch or reset against `master` after each service, you lose the local generator modifications and documentation in your experiment branch (`bulk-crd-gen-experiment`), risk leaking intermediate state, and cannot perform a holistic side-by-side comparison across services.

#### The Canonical Multi-Service Batch Pass on `bulk-crd-gen-experiment`:
1. **Committed Experiment Baseline:** Ensure all generator enhancements and documentation are cleanly committed on `bulk-crd-gen-experiment`.
2. **Multi-Service Batch Wipe:** Temporarily remove existing local types in `apis/<service>/<version>/` and CRD YAMLs in `config/crds/resources/` for all target evaluation services.
3. **Canonical Batch Execution:** Execute the official `./apis/<service>/generate.sh` script for each target service using the enhanced `controllerbuilder`, followed by a single `dev/tasks/generate-crds`.
4. **Holistic 3-Dimensional Diff Against Master/Baseline:**
   * **Gate 1 (Build):** Verify that all `generate.sh` scripts and `dev/tasks/generate-crds` exit `0`.
   * **Gate 2 (CRD Schema Parity):** Run `git diff config/crds/resources/` to inspect the exact OpenAPI v3 schema delta across all services side-by-side.
   * **Gate 3 (Go Types Parity):** Run `git diff apis/` to measure deterministic field coverage and catalog all structs requiring manual override.
5. **Scorecard Documentation:** Record verified metrics in `docs/designs/benchmark-results.md`.
6. **Single Clean Restoration:** Restore modified APIs/CRDs via `git checkout -- apis/ config/crds/resources/`.

---

## 6. Summary Matrix of Corrections

| Topic | Initial Mistake | Correct KCC Standard | Impact |
| :--- | :--- | :--- | :--- |
| **Enums** | Add `+kubebuilder:validation:Enum` | Unvalidated Go `*string` (no enum markers) | Prevents breaking CRDs when GCP adds enum values |
| **Override Mechanism** | Custom YAML overlay DSL (`overlay.yaml`) | Native Go override (`types.generated.go` vs `<kind>_types.go`) | Eliminates unnecessary abstraction; aligns with skills |
| **Proto Annotations** | Rely on `(google.api.resource_reference)` | Multi-signal heuristics + agent triage | Works across >90% of APIs where annotation is absent |
| **Parent Hierarchy** | Assume `google.api.resource` always exists | Use `pattern` when present; robust fallback when absent | Prevents generator failure on unannotated protos |
| **Credential Gaps** | Treat as distinct category from references | Classify as a subtype of Resource Reference (`SecretRef`) | Unifies reference resolution model |
| **Validation Planning** | Report numbers from ad-hoc AST scripts before building validation harness | Define test harness first; wipe $\to$ `generate.sh` $\to$ `git diff master` | 100% objective, ground-truth verification via Git |
| **Test Execution Flow** | Reset per-service against master (risking state leaks & generator context loss) | Multi-service batch run on experiment branch; single diff & clean reset | Preserves generator state; enables holistic side-by-side comparison |
