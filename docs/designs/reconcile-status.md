## Technical Design: Setting Status while updating a resource on GCP

### 1. Overview

This document proposes a standardized way for "direct" controllers in Config Connector (KCC) to provide visibility into the reconciliation process via the resource's `status.conditions`. Specifically, it addresses the need to indicate when a reconciliation is actively performing an update against the GCP API, why that update is occurring, and the outcome of the update.

Currently, many direct controllers only update the status once the reconciliation is complete. This leaves users in the dark during long-running operations, such as creating a large BigQuery dataset or a Cloud SQL instance. By updating the `Ready` condition to `False` with a reason of `Updating` at the start of the GCP interaction, we provide immediate feedback to the user.

### 2. Goals

*   **Real-time Visibility:** Provide immediate feedback when KCC starts an update against GCP.
*   **Reasoned Updates:** Indicate why an update is occurring (e.g., which field in the spec triggered the update).
*   **Standardized Conditions:** Reuse the standard `Ready` condition type and follow Kubernetes API conventions.
*   **Single Source of Truth:** Ensure there is exactly one `Ready` condition in the `status.conditions` list.
*   **Error Reporting:** Clearly indicate the outcome of the update, including detailed error messages on failure.

### 3. Non-Goals

*   This design does not apply to legacy Terraform or DCL-based controllers.
*   Comprehensive Long Running Operation (LRO) tracking (e.g., progress percentages) is not a goal for the initial implementation, though the `reason` field could potentially store an LRO ID in the future.

### 4. Proposed Solution

We will utilize the `Ready` condition in `status.conditions` to communicate the status of the reconciliation. The reconciler will transition this condition through different states as it progresses.

#### 4.1. Reconciliation States and Condition Values

| State | Status | Reason | Message |
| :--- | :--- | :--- | :--- |
| **Starting Update** | `False` | `Updating` | `Updating resource (first field path: spec.foo)` |
| **Successful Completion** | `True` | `UpToDate` | `Resource is up to date` |
| **Update Failure** | `False` | `UpdateFailure` | `Failed to update resource (first field path: spec.foo): <error message>` |

#### 4.2. Triggering Field Identification

The `Message` field will be used to provide context on why the reconciliation is performing an update. For updates triggered by a change in the spec, the message should identify the first field that was detected as different between the desired state and the actual state.

Example: `Updating resource (first field path: spec.description)`

If the update is not triggered by a field change (e.g., a periodic sync or a forced reconciliation), a generic message like `Updating resource (triggered by periodic sync)` can be used.

### 5. Implementation Details

#### 5.1. Adapter Interface Enhancement

To support reporting the first changed field, the `Adapter.Update` method (and potentially `Create`) needs a way to communicate this back to the reconciler, or it should set the status itself using the provided `UpdateOperation`.

Since the `Adapter` already has access to the `UpdateOperation` (or `CreateOperation`), it can call `UpdateStatus` with a specifically crafted `Ready` condition before it begins the actual GCP API call.

```go
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
    // 1. Perform diffing between a.desired and a.actual
    // 2. Identify the first changed field (e.g., "spec.description")
    
    // 3. If changes detected, set the Updating status
    if len(updateMask.Paths) > 0 {
        status := a.GetStatus() // Helper to get current status
        readyCondition := v1alpha1.Condition{
            Type: "Ready",
            Status: corev1.ConditionFalse,
            Reason: "Updating",
            Message: fmt.Sprintf("Updating resource (first field path: %s)", firstField),
        }
        if err := updateOp.UpdateStatus(ctx, status, &readyCondition); err != nil {
            return err
        }
        
        // 4. Call GCP API
        // ...
    }
    return nil
}
```

#### 5.2. Standardizing `UpdateStatus`

The `UpdateStatus` method in `pkg/controller/direct/directbase/operations.go` already uses `SetStatusCondition`, which ensures that only one condition of a given type exists.

```go
func SetStatusCondition(conditions *[]v1alpha1.Condition, newCondition v1alpha1.Condition) {
    // ... logic to find and update or append the condition ...
}
```

This implementation satisfies the requirement of having at most one `Ready` condition.

#### 5.3. Handling Errors

If the GCP API call fails, the adapter should return an error. The `DirectReconciler` catches this error and calls `handleUpdateFailed`, which sets the `Ready` condition to `False` with the error message.

We should ensure `handleUpdateFailed` uses a consistent `Reason` (e.g., `UpdateFailure`) and preserves the "first field path" information if available. This might require passing more context in the error or setting the status one last time before returning the error from the Adapter.

### 6. Kubernetes API Conventions Compliance

This design aligns with the [Kubernetes API conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status):

*   **Condition Type:** Reuses "Ready", which is a standard predicate.
*   **Status:** Uses "True" and "False".
*   **Reason:** Uses CamelCase strings ("Updating", "UpToDate", "UpdateFailure").
*   **ObservedGeneration:** `UpdateStatus` already sets `status.observedGeneration` to the current metadata generation.

### 7. Future Considerations: Long Running Operations (LROs)

For resources that use LROs, the `Reason` field could be updated to include the LRO ID once the operation has started. This would allow a controller that restarts to potentially "resume" or at least poll the status of an existing operation.

Example Reason: `Updating:operation-123456789`

### 8. Testing Strategy

#### 8.1. Unit Tests

*   Update `SetStatusCondition` tests to verify it correctly handles multiple updates to the same condition type.
*   Add unit tests for helper functions that extract the first changed field from a diff.

#### 8.2. End-to-End Tests

*   **Status Verification:** Enhance E2E test suites to verify that the `Ready` condition transitions through the expected states.
*   **Golden Files:** While golden HTTP logs don't capture K8s status updates, we can add assertions in our E2E tests (e.g., using `kubectl wait` or custom Go assertions) to verify the status during the reconciliation process.
*   **Negative Testing:** Verify that on update failure, the `Ready` condition accurately reflects the error and the reason.
