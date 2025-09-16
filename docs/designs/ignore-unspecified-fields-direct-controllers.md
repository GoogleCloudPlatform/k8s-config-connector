## Technical Design: Defaulting to GCP values for Unspecified Fields in Direct Controllers

### 1. Overview

This document proposes a new feature for KCC's "direct" controllers that provides a flexible way to manage fields that have server-side defaults or are managed out-of-band. The goal is to allow users to instruct KCC to preserve the value of a field as it exists in GCP, but only when that field is not specified in the KCC resource's spec. This allows for a smoother transition to managing existing resources with KCC and provides a safe way to handle emergency changes made directly in GCP.

This is enabled by a new annotation, `cnrm.cloud.google.com/default-to-gcp-fields`, which takes a comma-separated list of field paths.

This design is:
*   **Focused:** It applies exclusively to the "direct" controller architecture, aligning with the project's strategic goals.
*   **Safe:** It respects the Kubernetes API model by **never** mutating the user's `.spec`.
*   **Flexible:** The implementation is designed to work with all direct controllers, regardless of whether they use `CompareProtoMessage` with GCP protos or employ custom comparison logic (like the `SQLInstance` controller).

### 2. Goals

*   To allow users to provide a list of fields in the `spec` that KCC should default to the GCP value during reconciliation, if the field is not specified in the spec.
*   To implement this feature for **all** "direct" controllers in a consistent manner.
*   To ensure the user's `spec` is treated as immutable by the controller.
*   To provide a simple, declarative, and user-initiated way to enable this behavior.

### 3. Non-Goals

*   This feature will not be implemented for legacy Terraform or DCL-based controllers.

### 4. Proposed Solution

To provide this fine-grained, conditional control, we will introduce a new annotation:

`cnrm.cloud.google.com/default-to-gcp-fields: "spec.field.one,spec.field.two"`

This annotation instructs the controller to handle the listed fields with a "preserve-if-unspecified" logic. The behavior is as follows:

*   **If a field is listed in the annotation AND is NOT specified in the resource's `.spec`:** The controller will preserve the value of the field as it exists in GCP. It will not detect any drift for this field, effectively defaulting to the GCP value.
*   **If a field is listed in the annotation AND IS specified in the resource's `.spec`:** The `.spec` takes precedence. The controller will manage the field and ensure its value matches what is defined in the spec.

This is achieved by constructing a temporary, in-memory "effective desired state" object for comparison, as follows:

1.  The controller starts with the user's desired spec as a baseline.
2.  It fetches the actual, live state of the resource from GCP.
3.  It parses the comma-separated list of JSON field paths from the `default-to-gcp-fields` annotation.
4.  For each field path in the annotation, it checks if that field is present in the user's original spec.
5.  If the field is **NOT** in the spec, it retrieves the value for that field from the **actual state** and overlays it onto the baseline object.
6.  The resulting merged object is the "effective desired state", which is then used for comparison with the actual state.

### 5. Implementation Details

1.  **Design Principle: Pointers for Optional Fields:**
    *   As a design principle for all direct KRM APIs, all optional fields in the `spec` **must** be pointers (e.g., `*string`, `*int64`, `*bool`). This allows the controller to reliably distinguish between a field that is not specified by the user (`nil`) and a field that is explicitly set to its zero value. This principle is a prerequisite for the "preserve-if-unspecified" logic to work correctly.

2.  **Controller-Specific Merge Logic:**
    *   Each direct controller that supports the `default-to-gcp-fields` annotation will implement its own, private function to construct the "effective desired state".
    *   **Example for `SQLInstance`:** In `pkg/controller/direct/sql/sqlinstance_controller.go`, a new function `buildEffectiveDesiredSQLInstance` will be created.
    *   **Implementation:** This function will take the user's KRM spec and the live GCP state as input. It will create a deep copy of the live state, and then overlay the fields from the user's spec *only if the field in the spec is not `nil`*.

3.  **Controller Integration:**
    *   The `Update` function in each controller will parse the `default-to-gcp-fields` annotation.
    *   It will then call its specific merge function to get the effective desired state for comparison.

### 6. Testing Strategy

*   **Unit Tests:**
    *   Extensive unit tests for the new **Selective Merging Library**, covering:
        *   Correctly preserving a GCP value when the field is unspecified in the spec and listed in the annotation.
        *   Correctly using the spec value when the field is specified in the spec, even if it is also listed in the annotation.
        *   Handling of nested fields, lists, and different data types.
        *   Robust field presence detection logic.

*   **End-to-End Tests:**
    *   Each test fixture will execute the following scenarios:
        1.  **Preserve GCP Value:** Create a resource with a field managed out-of-band and listed in the `default-to-gcp-fields` annotation. Verify that a KCC reconciliation does not overwrite the out-of-band value.
        2.  **Take Ownership of Field:** Update the KCC resource to specify a value for the previously preserved field. Verify that KCC now enforces the value from the spec.
        3.  **Release Ownership of Field:** Update the KCC resource to remove the field from the spec (while keeping it in the `default-to-gcp-fields` annotation). Verify that KCC no longer manages the field and preserves the value from the previous step.

*   **Migration Testing for Beta Resources:**
    *   As a mandatory part of migrating any existing TF or DCL-based Beta resource to a direct controller, the migration plan **must** include the addition of an end-to-end test fixture. This test must validate the complete lifecycle of the `default-to-gcp-fields` annotation for that specific resource, ensuring that we maintain behavioral consistency and prevent regressions for our stable, Beta-level resources.

### 7. Alternatives Considered

#### 7.1. Reusing `state-into-spec: merge`

Instead of introducing a new annotation, we considered reusing the existing `cnrm.cloud.google.com/state-into-spec: "merge"` annotation.

*   **Rationale for Rejection**:
    *   **Legacy Coupling**: The `state-into-spec` annotation is tied to the legacy Terraform-based controller architecture. A key goal of the direct controller effort is to move away from these legacy concepts and adopt cleaner, more Kubernetes-native patterns. Introducing a new, purpose-built annotation provides a clean break.
    *   **Confusing Semantics**: The `state-into-spec` annotation has a history of being complex and sometimes confusing for users. A new annotation with a clear name like `default-to-gcp-fields` makes the user's intent much more explicit and self-documenting.
    *   **Future Flexibility**: A new annotation gives us the flexibility to evolve the behavior of direct controllers independently of the legacy controllers.

#### 7.2. Ignoring All Unspecified Fields

We considered a simpler approach with a single annotation, `cnrm.cloud.google.com/ignore-unspecified-fields: "true"`, which would instruct the controller to ignore *all* fields not present in the user's spec.

*   **Rationale for Rejection**:
    *   **Lack of Granularity**: This "all or nothing" approach is not flexible enough. Users often want to manage some server-defaulted fields while ignoring others. For example, they might want to manage the `activationPolicy` but ignore the `diskSize`.
    *   **Unexpected Behavior**: If a user adds a new field to their spec, they might be surprised to find that KCC suddenly starts managing it and potentially overwriting a server-side value. The proposed field-level annotation makes the user's intent more explicit.

#### 7.3. Modifying the `.spec` in-place

Another alternative was to have the controller read the live state from GCP and directly merge the unspecified fields into the `.spec` of the resource in the Kubernetes API server.

*   **Rationale for Rejection**:
    *   **Violation of Kubernetes Principles**: This is a strong anti-pattern in Kubernetes. The `.spec` should be owned by the user, and controllers should not mutate it. Mutating the spec would lead to unexpected behavior, conflicts with other controllers or webhooks, and a confusing user experience. The proposed solution respects the immutability of the user's desired state by only creating a temporary, in-memory object for comparison.