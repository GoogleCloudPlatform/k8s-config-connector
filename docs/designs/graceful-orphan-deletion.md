# Graceful Deletion of Resources with Missing Dependencies

## 1. Overview

This document proposes a change to the direct controller reconciliation logic to handle the deletion of a Kubernetes resource whose dependencies (e.g., a parent resource) have already been deleted. The current behavior causes the resource to be stuck in a `DependencyNotFound` error state, preventing its deletion. The proposed solution is to modify the deletion logic to rely on the resource's `status.externalRef` field instead of resolving dependencies from the `spec`, allowing the controller to delete the corresponding cloud resource and remove the finalizer.

## 2. Problem Statement

In KCC, many resources have explicit dependencies defined in their `spec`. For example, a `GKEHubFeatureMembership` has a `membershipRef` field that points to its parent `GKEHubMembership` resource.

The standard reconciliation loop resolves these dependencies to construct the full name or ID of the cloud resource required for `CREATE`, `UPDATE`, and `DELETE` operations. This works well for `CREATE` and `UPDATE`.

However, this dependency resolution becomes problematic during deletion in a specific scenario:

1.  A parent resource (e.g., `GKEHubMembership`) is deleted.
2.  A child resource (e.g., `GKEHubFeatureMembership`) that depends on it is then deleted.
3.  The `GKEHubFeatureMembership` controller's reconciliation loop is triggered by the `metadata.deletionTimestamp`.
4.  The controller attempts to resolve the `membershipRef` to construct the full name of the FeatureMembership resource in GCP.
5.  This resolution fails with a `DependencyNotFound` error because the parent `GKEHubMembership` Kubernetes resource is gone.
6.  The controller returns an error and the request is requeued. The finalizer on the `GKEHubFeatureMembership` is never removed.

The resource is now "orphaned" and stuck indefinitely, requiring manual intervention to clean up both the Kubernetes object (by removing the finalizer) and the underlying GCP resource.

## 3. Goals

*   Ensure KCC resources can be deleted reliably even if their dependencies are no longer present in the Kubernetes API.
*   Prevent controllers from getting stuck in a perpetual reconciliation loop during deletion of orphaned resources.
*   The solution should be generic and applicable to all direct controllers that manage resources with dependencies.
*   The deletion process should remain idempotent.

## 4. Non-Goals

*   This design does not address dependency issues during resource creation or updates.
*   It does not propose a generic garbage collection mechanism for GCP resources. The scope is limited to fixing the deletion logic for KCC-managed resources.

## 5. Proposed Design

The proposed solution integrates into the existing `directbase` controller framework by making two key changes:

1.  **Deletion-Aware Adapter Creation:** The `AdapterForObject` method in each resource's `Model` will be modified. When it detects a resource is being deleted (via `metadata.deletionTimestamp`), it will first attempt to construct an `Adapter` using the `status.externalRef`. This adapter is self-sufficient and can perform `Find` and `Delete` operations without needing any information from the `spec`. If `status.externalRef` is not present, it falls back to the original logic of resolving dependencies from the `spec`.

2.  **Contextual Error Handling in `doReconcile`:** The central `doReconcile` function in `directbase_controller.go` will be updated. When it receives a `DependencyNotFound` error from the `AdapterForObject` call, it will now check if the resource is being deleted. If it is, this combination of events (deletion timestamp + missing dependencies + no `status.externalRef`) is the definitive signal of an orphaned resource. In this specific case, it will log a warning and remove the finalizer to unblock the deletion, rather than retrying fruitlessly.

This approach is minimally invasive and leverages the existing controller structure to achieve the desired graceful deletion behavior.

### 5.1. Deletion Logic Flow

The updated flow within `doReconcile` will be as follows:

1.  The `doReconcile` function calls `model.AdapterForObject`.
2.  Inside `AdapterForObject`:
    *   It checks for a deletion timestamp.
    *   If found, it checks for `status.externalRef`.
    *   If `externalRef` exists, it returns a specialized adapter that uses it.
    *   If not, it proceeds to resolve dependencies from the `spec` as usual.
3.  Back in `doReconcile`:
    *   It checks the error returned from `AdapterForObject`.
    *   **If the error is `DependencyNotFound` AND the resource has a deletion timestamp**, it triggers the new orphan-handling logic: log a warning and call `handleDeleted()` (which removes the finalizer).
    *   If the error is `DependencyNotFound` during a create/update, it preserves the existing behavior of retrying.
    *   If no error occurs, it proceeds with the `Adapter` it received. The subsequent `adapter.Delete()` call will succeed whether the adapter was built from `status` or `spec`.

### 5.2. Pseudocode for Implementation

This pseudocode reflects the necessary changes in the resource-specific model and the generic `directbase` controller.

**1. Conceptual Change in `(model *Model) AdapterForObject(...)`**
*(This logic would be implemented in each direct controller's model, e.g., `gkehub/featuremembership_model.go`)*

```go
// In the model implementation for a resource...
func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
    // If the resource is being deleted AND has a status.externalRef,
    // construct a specialized adapter that can perform Find/Delete using only that identifier.
    // This bypasses spec dependency resolution entirely.
    if !u.GetDeletionTimestamp().IsZero() {
        externalRef, found, _ := unstructured.NestedString(u.Object, "status", "externalRef")
        if found && externalRef != "" {
            return newAdapterFromExternalRef(u, externalRef) // This adapter does not need resolved dependencies.
        }
    }

    // For create, update, or deletion where status.externalRef is missing,
    // proceed with the original logic of resolving dependencies from the spec.
    // This will correctly return an UnresolvableDeps error if parents are missing.
    if err := m.resolveDependencies(u); err != nil {
        return nil, err
    }
    return newAdapterFromSpec(u, m.resolvedRefs)
}
```

**2. Change in `(r *reconcileContext) doReconcile(...)`**
*(Located in `pkg/controller/direct/directbase/directbase_controller.go`)*

```go
func (r *reconcileContext) doReconcile(ctx context.Context, u *unstructured.Unstructured) (requeue bool, err error) {
    // ... (actuation mode check, defaulters) ...

    // Get the adapter. The model's AdapterForObject now contains the deletion-specific logic.
	adapter, adapterErr := r.Reconciler.model.AdapterForObject(ctx, r.Reconciler.Client, u)
	if adapterErr != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(adapterErr); ok {
			// **MODIFIED LOGIC**: If this is a deletion, an unresolvable dependency
            // means we have an orphan because the model already tried and failed to use status.externalRef.
			if !u.GetDeletionTimestamp().IsZero() {
				logger.Warning("could not resolve dependencies for deletion and status.externalRef is missing; resource may be orphaned in GCP", "resource", k8s.GetNamespacedName(u))
				// Remove the finalizer to unblock Kubernetes object deletion.
				return false, r.handleDeleted(ctx, u)
			}

			// This is the original behavior for create/update with missing dependencies.
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
			return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
		}
		// Handle other adapter creation errors.
		return false, r.handleUpdateFailed(ctx, u, adapterErr)
	}

    // ... (the rest of the function, including adapter.Find(), remains largely the same) ...

	// Deletion logic block (now guaranteed to have a working adapter)
	if !u.GetDeletionTimestamp().IsZero() {
		// ... (finalizer and annotation checks) ...

        existsAlready, err := adapter.Find(ctx)
        // ... (handle find errors) ...

		if !existsAlready {
			logger.Info("underlying resource does not exist; no API call necessary", "resource", k8s.GetNamespacedName(u))
			return false, r.handleDeleted(ctx, u)
		}

		logger.Info("deleting underlying resource", "resource", k8s.GetNamespacedName(u))
		deleteOp := NewDeleteOperation(r.Reconciler.Client, u)
		if _, err := adapter.Delete(ctx, deleteOp); err != nil {
			// A real GCP API error, not a dependency issue.
			return false, r.handleDeleteFailed(ctx, u, err)
		}
		return false, r.handleDeleted(ctx, u)
	}

	// ... (create/update logic) ...
}
```

### 5.3. Contract for Direct Controllers

This design imposes a requirement on all direct controllers: they **must** reliably populate a `status.externalRef` field upon the successful creation of a GCP resource. This field must contain enough information to uniquely identify and delete the resource later.

## 6. Implementation Plan

1.  **Phase 1: Proof of Concept:** Implement the proposed deletion logic for a single, known-problematic resource, such as `GKEHubFeatureMembership`, to validate the approach.
2.  **Phase 2: Generalization:** Refactor the logic into a shared package or utility function within `pkg/controller/direct/` that can be easily consumed by all direct controllers.
3.  **Phase 3: Adoption:** Update all existing direct controllers to use this new, generalized deletion logic.
4.  **Phase 4: Documentation:** Update the developer documentation for creating new direct controllers to mandate the use of this pattern and the population of `status.externalRef`.

## 7. Alternatives Considered

### 7.1. Annotation-Based Deletion Policy

An annotation like `cnrm.cloud.google.com/deletion-policy: "abandon"` could be used to instruct the controller to simply remove the finalizer without attempting to delete the underlying resource.

*   **Downsides:** This places the burden on the user to know when to apply this annotation. The proposed design is superior because it represents a more correct default behavior that requires no user intervention.

## 8. Testing Strategy

To ensure this solution is correctly implemented and to prevent future regressions, we will adopt a two-pronged testing strategy.

### 8.1. E2E Resource Fixture Test for Orphan Deletion

We will introduce a new test pattern for our resource fixture-based E2E tests. This pattern will specifically simulate the orphan scenario and verify the controller behaves correctly.

#### 8.1.1. New Test Fixture File

A new, optional file named `delete_dependencies.yaml` will be introduced into the test fixture directory structure. The presence of this file will signal the test runner to execute a special orphan deletion test case for that resource.

#### 8.1.2. Test Execution Flow

When a test fixture contains `delete_dependencies.yaml`, the test runner (`TestAllInSeries`) will be modified to perform the following sequence after the standard `create` and `update` steps have passed:

1.  **Identify Dependencies:** The test runner will parse `delete_dependencies.yaml`, which will contain a list of `GroupVersionKind` and names for the dependency resources to be deleted (typically the resources from `dependencies.yaml`).
2.  **Delete Dependencies:** The runner will delete the specified dependency resources from the Kubernetes API server and wait for them to be fully removed.
3.  **Delete Primary Resource:** The runner will then delete the primary resource being tested (the one defined in `create.yaml`).
4.  **Assert Finalization:** The core of the test is to assert that the primary resource is also successfully and completely removed from the Kubernetes API server within a reasonable timeout.

If the primary resource is not removed, it means its finalizer was not removed, indicating the controller is stuck in the `DependencyNotFound` loop and the test will fail.

### 8.2. Static Code Analysis

To proactively enforce this new pattern and catch errors at compile time, we will develop a custom static analysis check (linter).

*   **Linter Logic:** The linter will scan all Go files implementing the `directbase.Model` interface. It will enforce the following rule:
    > In any implementation of `AdapterForObject`, if the function can return an error of type `lifecyclehandler.DependencyNotFoundError`, then the function **must** also contain a check for the resource's deletion timestamp (`u.GetDeletionTimestamp()`).

*   **Integration:** This linter will be integrated into our CI/presubmit checks. This ensures that no new direct controller can be merged if it resolves dependencies without correctly handling the orphan deletion case, making the pattern self-enforcing for all future development.
