---
name: reviewgen-greenfield-new-types
description: Provides provides clear review criteria for reviewing PRs that add new KCC types for Greenfield resources.
---

# Review guide for KCC Greenfield new types PRs
Please respect the following review criteria and invariants when reviewing.

## 1. API Versioning
*   All Greenfield resources must be implemented as `v1alpha1` in KRM.
*   Verify that `spec.versions.name = v1alpha1` in the generated CRD YAML.
*   Verify files are placed in: `apis/${resource_group}/v1alpha1/`.
*   *Note:* Do not confuse the GCP API version (which can be higher, e.g., v1) with the KRM version (which must be `v1alpha1`).

## 2. Copyright Year
*   All new `.go` and `.sh` files must contain a copyright header.
<!-- TODO: Dynamically determine year -->
*   The copyright year **must be 2026**. 

## 3. Pointers (Go Types)
*   For `${resource_name}_types.go`, review the field comments (e.g., Kubebuilder tags indicating `required` or `optional`).
*   **Strict Rule:** If a field is a Go scalar primitive type (e.g., `string`, `bool`, `int`, `int32`, `int64`, `float64`), it **must be a pointer** (e.g., `*string`, `*bool`), regardless of whether it is optional or required.
*   **Collection Exception:** Do **not** make slice fields (e.g., `[]string`) or map fields (e.g., `map[string]string`) pointers (i.e., do not write `*[]string` or `*map[string]string`).

## 4. References & Identity
*   All string fields referencing a GCP resource identifier must map to a reference struct in Go to validate format/semantics.
*   If a resource is a child of another GCP resource, this relationship must be explicitly denoted in the code.
*   Compare the implementation against the provided reference files in https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7894 and https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7907/

## 5. Completeness & Heuristics (Proto-to-CRD mapping)
*   **Completeness Goal:** Greenfield resources must aim for 100% coverage of the fields defined in the Google API proto. Compare the generated CRD YAML against the generated proto files. The CRD must map **all** fields declared in the Proto.
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
* **Completeness & Heuristics**: [Pass/Fail] - (List any missing or incorrectly mapped fields)
* **References/Identity**: [Pass/Fail] - (List any missing resource references)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
