---
name: reviewgen-brownfield-new-types
description: Provides provides clear review criteria for reviewing PRs that add new KCC types for Brownfield resources.
---

# Review guide for KCC Brownfield new types PRs
Please respect the following review criteria and invariants when reviewing.

## 1. API Versioning
*   All Brownfield resources should be implemented in KRM as `v1beta1` (or matching their existing version in the master branch).
*   Verify that the KRM version name (`spec.versions[].name`) matches the existing CRD version (typically `v1beta1`).
*   Verify files are placed in: `apis/${resource_group}/${version}/` (typically `v1beta1`).
*   *Note:* Do not confuse the GCP API version with the KRM version.

## 2. Copyright Year
*   All new `.go` and `.sh` files must contain a copyright header. For all new files the copyright year **must be 2026**.
*   For existing files the copyright year **must not be modified**.

## 3. Pointers (Go Types)
*   For `${resource_name}_types.go`, review the field comments (e.g., Kubebuilder tags indicating `required` or `optional`).
*   **Strict Rule:** If a field is a Go scalar primitive type and it is optional (e.g., `string`, `bool`, `int`, `int32`, `int64`, `float64`), it **must be a pointer** (e.g., `*string`, `*bool`). 
*  **Strict Rule:** If a field is a Go scalar primitive type and it is required it should be implemented as a value (e.g., `string`, `bool`, `int`, `int32`, `int64`, `float64`).
*   **Collection Exception:** Do **not** make slice fields (e.g., `[]string`) or map fields (e.g., `map[string]string`) pointers (i.e., do not write `*[]string` or `*map[string]string`).

## 4. References & Identity
*   All string fields referencing a GCP resource identifier must map to a reference struct in Go to validate format/semantics.
*   If a resource is a child of another GCP resource, this relationship must be explicitly denoted in the code.
*   Compare the implementation against the provided reference files in https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7894 and https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7907/

## 5. Schema Compatibility (No Breaking Changes)
*   **Strict Rule:** No changes to the schema (fields added, removed, or type changes) are allowed compared to the existing CRD in the `master` branch.
    * **Exceptions:**
      1. Adding `status.externalRef` (as `ExternalRef *string `json:"externalRef,omitempty"``) under the Status struct is allowed and expected.
      2. Removing the following top level `status` block located at the end of the CRD yaml file is fine and should not be flagged in the review:
      ```yaml
      status:
        acceptedNames:
          kind: ""
          plural: ""
        conditions: []
        storedVersions: []
      ```  
*   The generated CRD must match the existing one exactly, except for minor differences like descriptions and the allowed addition of `status.externalRef`.
*   **Action:** Run `./dev/tasks/diff-crds --base master --allow-status-external-ref-addition` from the repository root. If the output shows any added or removed fields under `spec` or `status` (excluding comments, description text updates, or allowed `status.externalRef`), fail the review.

## 6. Completeness & Heuristics (Proto-to-CRD mapping)
*   **Parity over Completeness:** Unlike Greenfield resources (where we aim for 100% coverage of the proto), Brownfield migration PRs must maintain strict schema parity with the existing CRD. Do **NOT** add fields from the proto if they were not already present in the master branch CRD. The mapping heuristics below apply only to the fields already present in the existing CRD.
*   Find the proto definition in `.build/third_party/googleapis/google/...` matching the service named in the resource's `generate.sh` (e.g., `google.cloud.apihub.v1`).
*   Verify that fields are mapped using these rules:
    1. **`status` Mapping:** Fields containing `(google.api.field_behavior) = OUTPUT_ONLY` in the proto must map only to Go's `Status` struct (represented as `status` in the CRD).
    2. **`spec` Mapping:** Fields without `OUTPUT_ONLY` behavior in the proto must map to Go's `Spec` struct (represented as `spec` in the CRD).
*   CRD fields must align with the Kubernetes Resource Model (KRM) conventions. Useful references include:
    *   [Kubernetes Resource Management Design Proposal](https://github.com/kubernetes/design-proposals-archive/blob/main/architecture/resource-management.md)
    *   [Kubernetes API Conventions](https://github.com/kubernetes/community/blob/main/contributors/devel/sig-architecture/api-conventions.md)

## 7. Scope Boundaries (Step 1 vs Step 2)
* **Step 1 (`step/types`) PRs**:
  - Must ONLY contain KRM Go types (`_types.go`), `types.generated.go`, `zz_generated.deepcopy.go`, `groupversion_info.go`, `doc.go`, `apis/<service>/generate.sh` (`generate-types`), and regenerated CRD YAML.
  - **FAIL** if the PR contains mappers (`mapper.generated.go`, `mappers.go`) or controller code under `pkg/controller/direct/`. Mappers belong in Step 3 (Fuzzer) / Step 5 (Controller).
  - **FAIL** if the PR contains identity or reference files (`_identity.go`, `_reference.go`). Identity belongs in Step 2.
  - **FAIL** if the PR contains committed internal journal files or changes under `.gemini/skills/`.
* **Step 2 (`step/identity`) PRs**:
  - Must ONLY contain identity and reference files (`_identity.go`, `_reference.go`), `generate.sh` (`generate-identity`), and CAIS test output (`_identities.yaml`).
  - **FAIL** if the PR contains mappers or controller reconciliation code.
  - **FAIL** if the PR contains committed internal journal files or changes under `.gemini/skills/`.

## 8. Code Quality, Style & Design Practices
Pay attention to race conditions, edge cases, and input/output correctness. Ensure correctness first, then do a second pass to ensure the code follows these style and design practices:

1. **1 File = 1 Class / Struct (Single Responsibility per File)**:
   * Follow a 1-file, 1-class (1 primary struct/concept + its methods) design with a single responsibility per file; avoid dumping multiple responsibilities or bag-of-fields helper structs into a types or identity file.
   * If standalone unit tests are written, each `<file>.go` should have a matching `<file>_test.go` with tests structured into 3 visually separated blocks (Arrange, Act, Assert) separated by blank lines.

2. **Treat Files as Private by Default (No Cross-File Package Globals)**:
   * Even though Go shares lowercase symbols across a package, treat each `.go` file as a strict encapsulation boundary.
   * Across files in the same package, only interact via a struct's constructor (`newX(...)`) and its methods (`x.doWork(...)`).
   * Do NOT reach into another struct's unexported fields from a different file.
   * Do NOT call free helper functions or reference const declarations defined in another file unless that file's explicit sole purpose is a shared domain model.

3. **No Feature Envy on Value/Data Structs**:
   * Value structs (e.g., parsed identifiers, reference types, options) must only contain data and validation for that data.
   * Do not attach downstream domain logic (e.g., formatting IAM emails, constructing API client payloads, or computing GCS regions) onto a shared data struct; keep those as private methods/helpers in the file that actually performs that domain work.

4. **Own Collaborators via Constructor, Not Method Threading**:
   * If struct A uses collaborator B, wire B inside `newA(...)` (and store B as a field on A) rather than having the caller construct B and pass it into multiple methods of A (`a.foo(b)`, `a.bar(b)`).
   * Initialize optional collaborators in the constructor rather than repeating `if x.field != nil` guards across methods.

5. **Depth-First Call-Tree Ordering Within a File (Stepdown Rule)**:
   * A reader must be able to read any file from top to bottom without jumping to other files—or bouncing to the bottom of the same file—to hunt down helpers or constants.
   * Do NOT group all receiver methods at the top of a file and dump private helper functions at the bottom.
   * Place every private helper function or sub-method immediately below the function that calls it, in the exact order they are called: struct -> constructor -> constructor's helpers -> primary method -> its helpers in call order -> trivial 1-line getters/delegators at the bottom.

6. **Colocate Constants with Their First Caller**:
   * Do not dump all const declarations at the top of the file if they are only used by a specific sub-tree of functions further down.
   * Declare const blocks immediately above the first function or method that uses them.

7. **Orchestrator vs. Leaf Method Granularity**:
   * Reference: [Function Length](https://martinfowler.com/bliki/FunctionLength.html).
   * Orchestrator methods must stay at a single high level of abstraction and should NOT contain inline for loops or anonymous struct tables for sub-steps. Extract each loop/sub-step into a well-named helper placed immediately below in call order.
   * Leaf methods (methods that perform one concrete task) should treat the 5–15 line length as a guideline, not a hard quota. Do NOT shred a cohesive 20-line leaf method into 3–4 single-use 5-line wrappers.

8. **No `(nil, nil)` Return Contracts**:
   * Never return `(nil, nil)` from a fetch/lookup/parser function. Return either a valid value or a descriptive error (or `(T, bool)` for pure in-memory lookups) so callers never need defensive `if val == nil` checks after `if err != nil`.

## 9. PR Description & Metadata
* **Release Note Block**: Verify that the PR description includes a release note block set to `NONE`:
  ```release-note
  NONE
  ```
  Fail or flag if the release note block is missing or not set to `NONE`.

# Review Comment Template
When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **Scope Check**: [Pass/Fail] - (Confirm no mappers or out-of-step files included)
* **Release Note Check**: [Pass/Fail] - (Confirm release-note block is set to NONE)
* **API Version Check**: [Pass/Fail] - (Specify paths/versions checked)
* **Go Type Pointers**: [Pass/Fail] - (List any non-pointer primitives found)
* **Schema Compatibility**: [Pass/Fail] - (List any diff-crds issues)
* **References/Identity**: [Pass/Fail] - (List any missing resource references or out-of-step identity files)
* **Code Quality & Design**: [Pass/Fail] - (Check single responsibility, encapsulation, Stepdown ordering, no feature envy, no (nil, nil))

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```

