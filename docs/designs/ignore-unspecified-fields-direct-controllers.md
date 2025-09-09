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
3.  It then recursively overlays the fields from the user's `spec` (the original, immutable desired state) on top of this baseline. Any field that is explicitly defined in the `spec` will overwrite the corresponding value from the actual state in this in-memory copy.
4.  The resulting merged object is the "effective desired state". It effectively represents "what the user wants, plus what GCP has for everything else."
5.  This "effective desired state" is then passed to the controller's existing comparison function (e.g., `CompareProtoMessage`, `InstancesMatch`, or other custom diff logic) in place of the original `spec`.

Because the effective desired state already contains the server-side values for any unspecified fields, the comparison function will find no difference for those fields, and KCC will not attempt to modify them.

### 5. Implementation Details

The key to making this work universally is a generic, centralized function that constructs this "effective desired state" object, operating on `unstructured.Unstructured` objects.

1.  **New Generic Function: `BuildEffectiveDesiredStateForComparison`:**
    *   A new function, `BuildEffectiveDesiredStateForComparison`, will be created in a shared location like `pkg/controller/direct/common/comparison.go` (or a new file like `pkg/controller/direct/common/effectivedesiredstate.go`).
    *   **Signature:** `func BuildEffectiveDesiredStateForComparison(userSpec, liveState map[string]interface{}) (map[string]interface{}, error)`
    *   **Purpose:** This function's documentation will explicitly state that it creates a *temporary, in-memory object solely for comparison purposes* and *does not modify the original `userSpec` or the Kubernetes resource's `.spec`*.
    *   **Implementation:** This function will perform a deep, recursive merge. It will start with a deep copy of `liveState` and then recursively overlay the `userSpec` onto it. It must correctly handle nested maps and slices.

2.  **Controller Integration Strategy:**
    Each direct controller's reconciliation logic will be updated to include this standard pre-comparison step.

    *   **Common Pattern:** In the `Update` function of each direct controller, after fetching the `userSpec` (from the KRM object) and the `liveState` (from GCP), the controller will check for the `cnrm.cloud.google.com/ignore-unspecified-fields: "true"` annotation.
    *   **Conditional Execution:**
        *   If the annotation is **present**, the controller will call `effectiveDesiredState, err := BuildEffectiveDesiredStateForComparison(userSpec, liveState)`.
        *   This `effectiveDesiredState` will then be used as the "desired" input for the controller's specific comparison function.
        *   If the annotation is **not present**, the original `userSpec` will be used as the "desired" input for the comparison function, maintaining current behavior.

    *   **Example for controllers with custom logic (e.g., `SQLInstance`):**
        *   In `pkg/controller/direct/sql/sqlinstance_controller.go`, before calling `InstancesMatch(desiredGCP, a.actual, ...)`, the controller will check for the annotation.
        *   If present, it will construct a new `desiredGCP` object by calling `BuildEffectiveDesiredStateForComparison` on the original spec and the actual GCP state.
        *   This new, merged object will be passed to `InstancesMatch`.

    *   **Example for controllers using `CompareProtoMessage`:**
        *   In the controller's `Update` function, it will check for the annotation.
        *   If present, it will perform the merge using `BuildEffectiveDesiredStateForComparison` on the unstructured `spec` and `actual` objects.
        *   The resulting `unstructured` object (the `effectiveDesiredState`) will then be converted to the required GCP protobuf message before being passed to `CompareProtoMessage`.

This pattern ensures that all direct controllers, regardless of their internal diffing mechanism, can leverage the same core logic for ignoring unspecified fields, providing a consistent and reliable user experience.

### 6. Testing Strategy

*   **Unit Tests:**
    *   Extensive unit tests for the new `BuildEffectiveDesiredStateForComparison` function, covering nested fields, lists, replacement of existing values, and preservation of non-specified values. These tests will explicitly assert that the original `userSpec` object remains unchanged.
*   **End-to-End Tests:**
    *   Add a test fixture for a resource that **uses `CompareProtoMessage`** (e.g., `StorageBucket`). This test will verify the end-to-end flow for the most common type of direct controller.
    *   Add a test fixture for the **`SQLInstance` resource**. This is critical to prove the solution works for controllers with custom comparison logic.
    *   Both test fixtures will execute the same scenarios:
        1.  Create the resource with the annotation and a minimal `spec`; verify server-side defaults are preserved.
        2.  Update the resource to explicitly specify a previously ignored field; verify KCC now manages it.
        3.  Remove the annotation; verify the resource reverts to the default KCC behavior.
