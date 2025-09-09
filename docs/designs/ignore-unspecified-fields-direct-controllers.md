## Technical Design: Ignoring Unspecified Fields in Direct Controllers

### 1. Overview

This document proposes a new feature for KCC's "direct" controllers to give users more flexible control over how resource fields are managed. The goal is to allow users to have KCC manage only the fields they explicitly define in their resource manifest, while ignoring all other fields and preserving any values set directly on the GCP resource.

This design is:
*   **Focused:** It applies exclusively to the "direct" controller architecture, aligning with the project's strategic goals.
*   **Safe:** It respects the Kubernetes API model by **never** mutating the user's `.spec`.
*   **Flexible:** The implementation is designed to work with all direct controllers, regardless of whether they use `CompareProtoMessage` with GCP protos or employ custom comparison logic (like the `SQLInstance` controller).

### 2. Goals

*   To allow users to instruct KCC to ignore any field that is not explicitly set in the resource `spec`.
*   To implement this feature for **all** "direct" controllers in a consistent manner.
*   To ensure the user's `spec` is treated as immutable by the controller.
*   To provide a simple, declarative, and user-initiated way to enable this behavior.

### 3. Non-Goals

*   This feature will not be implemented for legacy Terraform or DCL-based controllers.

### 4. Proposed Solution

The user-facing contract remains a simple, resource-level annotation:

`cnrm.cloud.google.com/ignore-unspecified-fields: "true"`

When a direct controller reconciles a resource with this annotation, it will construct a new, **temporary, in-memory "effective desired state" object** before running its diff logic. This object is *only used for the comparison step* and is *never written back to the resource's `.spec`*.

The process for constructing this "effective desired state" is as follows:

1.  The controller fetches the actual, live state of the resource from GCP.
2.  It creates a deep copy of this **actual state** to use as a baseline for the comparison.
3.  It then overlays the fields from the user's `spec` (the original, immutable desired state) on top of this baseline. Any field that is explicitly defined in the `spec` will overwrite the corresponding value from the actual state in this in-memory copy. The logic for this overlay will be implemented on a **per-controller basis**, ensuring it is tailored to the specific resource structure.
4.  The resulting merged object is the "effective desired state". It effectively represents "what the user wants, plus what GCP has for everything else."
5.  This "effective desired state" is then passed to the controller's existing comparison function (e.g., `CompareProtoMessage`, `InstancesMatch`, or other custom diff logic) in place of the original `spec`.

Because the effective desired state already contains the server-side values for any unspecified fields, the comparison function will find no difference for those fields, and KCC will not attempt to modify them. This approach avoids the complexity of a generic, shared library for merging, and instead empowers each controller to handle its own resource type correctly.

### 5. Implementation Details

The logic for constructing the "effective desired state" will be implemented at the controller level, providing type safety and allowing for resource-specific handling.

1.  **Controller-Specific Merge Function:**
    *   Each direct controller that supports this feature will implement a new, private function to handle the merge logic for its specific resource type.
    *   **Example for `SQLInstance`:** In `pkg/controller/direct/sql/sqlinstance_controller.go`, a new function could be introduced:
        *   **Signature:** `func buildEffectiveDesiredSQLInstance(userKRM *krm.SQLInstance, liveState *api.DatabaseInstance) (*api.DatabaseInstance, error)`
        *   **Purpose:** This function will take the user's desired KRM resource and the live state from GCP, and return a new `DatabaseInstance` object representing the effective desired state for comparison.
        *   **Implementation:** The function will create a deep copy of the `liveState` and then carefully overlay the fields from the `userKRM.Spec`. This ensures that any field not present in the user's spec will retain its value from the `liveState`.

2.  **Controller Integration Strategy:**
    The reconciliation logic in each direct controller will be updated to use this new function.

    *   **Common Pattern:** In the `Update` function of each direct controller, after fetching the user's resource and the live state, the controller will check for the `cnrm.cloud.google.com/ignore-unspecified-fields: "true"` annotation.
    *   **Conditional Execution:**
        *   If the annotation is **present**, the controller will call its specific merge function (e.g., `buildEffectiveDesiredSQLInstance`) to construct the effective desired state. This effective state will then be used as the "desired" input for the controller's comparison function.
        *   If the annotation is **not present**, the controller will convert the original user spec to the GCP API format (as it does today) and use that for the comparison, maintaining the default "prune" behavior.

    *   **Example for `SQLInstance`:**
        *   In `pkg/controller/direct/sql/sqlinstance_controller.go`, the `Update` function will be modified:
        ```go
        func (a *sqlInstanceAdapter) Update(...) {
            ...
            var desiredForComparison *api.DatabaseInstance
            var err error

            if u.GetAnnotations()["cnrm.cloud.google.com/ignore-unspecified-fields"] == "true" {
                desiredForComparison, err = buildEffectiveDesiredSQLInstance(a.desired, a.actual)
            } else {
                desiredForComparison, err = SQLInstanceKRMToGCP(a.desired, a.actual, false)
            }
            if err != nil {
                return err
            }

            if !InstancesMatch(desiredForComparison, a.actual, ...) {
                // ... perform update
            }
            ...
        }
        ```

This per-controller approach ensures that the merge logic is tailored to the specific resource, is type-safe, and is easier to maintain and debug than a single, generic, and untyped library function.

### 6. Testing Strategy

The testing strategy will focus on both the correctness of the per-controller merge logic and the end-to-end user experience.

*   **Unit Tests:**
    *   For each direct controller that implements this feature, extensive unit tests will be added for its **new, resource-specific merge function** (e.g., `buildEffectiveDesiredSQLInstance`).
    *   These tests will cover various scenarios, including:
        *   Correctly overlaying fields from the user's spec onto the live state.
        *   Handling of nested fields, lists, and other complex structures specific to that resource.
        *   Ensuring that fields *not* in the user's spec retain their values from the live state.
        *   Asserting that the original user spec and live state objects are not mutated by the merge function.

*   **End-to-End Tests:**
    *   We will add test fixtures for at least two resources to ensure broad coverage of the pattern:
        1.  The **`SQLInstance` resource**, to validate the approach for controllers with custom comparison logic.
        2.  Another key direct controller (e.g., `StorageBucket` if it becomes a direct controller, or another suitable candidate) to validate the pattern for more standard controllers.
    *   Each test fixture will execute the following scenarios to verify the complete lifecycle:
        1.  **Create with Annotation:** Create the resource with the `cnrm.cloud.google.com/ignore-unspecified-fields: "true"` annotation and a minimal `spec`. Verify that server-side default values are preserved and KCC does not try to overwrite them.
        2.  **Update to Manage Field:** Update the resource to explicitly specify a field that was previously ignored. Verify that KCC now correctly manages this field and updates it to the value in the spec.
        3.  **Remove Annotation:** Remove the annotation from the resource. Verify that the controller reverts to its default "prune" behavior, and that any fields present in the live state but not in the spec are now removed (or reset to their zero value) by KCC.
        3.  **Remove Annotation:** Remove the annotation from the resource. Verify that the controller reverts to its default "prune" behavior, and that any fields present in the live state but not in the spec are now removed (or reset to their zero value) by KCC.
        3.  **Remove Annotation:** Remove the annotation from the resource. Verify that the controller reverts to its default "prune" behavior, and that any fields present in the live state but not in the spec are now removed (or reset to their zero value) by KCC.

*   **Migration Testing for Beta Resources:**
    *   As a mandatory part of migrating any existing TF or DCL-based Beta resource to a direct controller, the migration plan **must** include the addition of an end-to-end test fixture. This test must validate the complete lifecycle of the `ignore-unspecified-fields` annotation for that specific resource, ensuring that we maintain behavioral consistency and prevent regressions for our stable, Beta-level resources.

### 7. Alternatives Considered

#### 7.1. Reusing `state-into-spec: merge`

Instead of introducing a new annotation, we considered reusing the existing `cnrm.cloud.google.com/state-into-spec: "merge"` annotation.

*   **Rationale for Rejection**:
    *   **Legacy Coupling**: The `state-into-spec` annotation is tied to the legacy Terraform-based controller architecture. A key goal of the direct controller effort is to move away from these legacy concepts and adopt cleaner, more Kubernetes-native patterns. Introducing a new, purpose-built annotation provides a clean break.
    *   **Confusing Semantics**: The `state-into-spec` annotation has a history of being complex and sometimes confusing for users. A new annotation with a clear name like `ignore-unspecified-fields` makes the user's intent much more explicit and self-documenting.
    *   **Future Flexibility**: A new annotation gives us the flexibility to evolve the behavior of direct controllers independently of the legacy controllers.

#### 7.2. Field-level Annotation

We considered an alternative where the user would specify which fields to ignore at a more granular level, for example, by listing them in an annotation: `cnrm.cloud.google.com/ignore-fields: "fieldA,fieldB.subFieldC"`.

*   **Rationale for Rejection**:
    *   **High User Burden**: This approach would be very cumbersome for users. They would need to identify and list every single server-defaulted field they want to preserve. This is error-prone and creates a high maintenance burden.
    *   **Brittleness**: If the GCP API introduces new server-defaulted fields, users' manifests would become outdated and would need to be updated to preserve the new fields.
    *   **Complexity**: The implementation would be more complex, as it would require parsing the field paths and selectively merging them. The proposed "all or nothing" approach is much simpler to implement and for users to understand.

#### 7.3. Modifying the `.spec` in-place

Another alternative was to have the controller read the live state from GCP and directly merge the unspecified fields into the `.spec` of the resource in the Kubernetes API server.

*   **Rationale for Rejection**:
    *   **Violation of Kubernetes Principles**: This is a strong anti-pattern in Kubernetes. The `.spec` should be owned by the user, and controllers should not mutate it. Mutating the spec would lead to unexpected behavior, conflicts with other controllers or webhooks, and a confusing user experience. The proposed solution respects the immutability of the user's desired state by only creating a temporary, in-memory object for comparison.

#### 7.4. Centralized, Generic Merge Library

We initially considered implementing the merge logic in a single, shared library function, `BuildEffectiveDesiredStateForComparison`, that would operate on untyped `unstructured.Unstructured` objects. This function would have been called by all direct controllers to ensure consistency.

*   **Rationale for Rejection**:
    *   **High Complexity and Brittleness**: A generic function that can correctly merge any two arbitrary resource structures is extremely difficult to write and maintain. GCP resources have a wide variety of structures, including nested objects, lists of objects with different merge keys (e.g., by `name` or by `port`), and other idiosyncrasies. A "one size fits all" approach is likely to be complex, buggy, and result in a "magic" function that is hard to debug.
    *   **Loss of Type Safety**: Operating on `unstructured.Unstructured` (i.e., `map[string]interface{}`) means giving up Go's type safety. This makes the code harder to reason about, more prone to runtime errors, and moves error detection from compile-time to runtime.
    *   **Controller-Specific Knowledge**: The logic for merging two resource states often requires knowledge of the specific resource's semantics. For example, merging a list of disks requires knowing that the `deviceName` is the key. This domain-specific knowledge is best encapsulated within the resource's own controller, rather than in a generic, central function. The proposed per-controller approach allows for this, leading to more robust and maintainable code.
