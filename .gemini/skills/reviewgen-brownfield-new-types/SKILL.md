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
    * **Exception:** Removing the following top level `status` block located at the end of the CRD yaml file is fine and should not be flagged in the review:
      ```yaml
      status:
        acceptedNames:
          kind: ""
          plural: ""
        conditions: []
        storedVersions: []
      ```  
*   The generated CRD must match the existing one exactly, except for minor differences like descriptions.
*   **Action:** Run `./dev/tasks/diff-crds --base master` from the repository root. If the output shows any added or removed fields under `spec` or `status` (excluding comments or description text updates), fail the review.

## 6. Completeness & Heuristics (Proto-to-CRD mapping)
*   **Parity over Completeness:** Unlike Greenfield resources (where we aim for 100% coverage of the proto), Brownfield migration PRs must maintain strict schema parity with the existing CRD. Do **NOT** add fields from the proto if they were not already present in the master branch CRD. The mapping heuristics below apply only to the fields already present in the existing CRD.
*   Find the proto definition in `.build/third_party/googleapis/google/...` matching the service named in the resource's `generate.sh` (e.g., `google.cloud.apihub.v1`).
*   Verify that fields are mapped using these rules:
    1. **`status` Mapping:** Fields containing `(google.api.field_behavior) = OUTPUT_ONLY` in the proto must map only to Go's `Status` struct (represented as `status` in the CRD).
    2. **`spec` Mapping:** Fields without `OUTPUT_ONLY` behavior in the proto must map to Go's `Spec` struct (represented as `spec` in the CRD).
*   CRD fields must align with the Kubernetes Resource Model (KRM) conventions. Useful references include:
    *   [Kubernetes Resource Management Design Proposal](https://github.com/kubernetes/design-proposals-archive/blob/main/architecture/resource-management.md)
    *   [Kubernetes API Conventions](https://github.com/kubernetes/community/blob/main/contributors/devel/sig-architecture/api-conventions.md)

# Review Comment Template
When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **API Version Check**: [Pass/Fail] - (Specify paths/versions checked)
* **Go Type Pointers**: [Pass/Fail] - (List any non-pointer primitives found)
* **Schema Compatibility**: [Pass/Fail] - (List any diff-crds issues)
* **References/Identity**: [Pass/Fail] - (List any missing resource references)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
