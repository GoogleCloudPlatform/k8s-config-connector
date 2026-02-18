# Technical Design: Read-Only Resources via Actuation Mode

## 1. Overview

This document proposes a mechanism to support "read-only" resources in Config Connector (KCC). This feature allows users to designate specific KCC resources as "observed-only," meaning KCC will reconcile their status from the Google Cloud Platform (GCP) but will **not** attempt to create, update, or delete the underlying GCP resource. This effectively puts the resource into a "read-only" mode from the perspective of the controller's actuation logic.

## 2. Motivation

There are several scenarios where users need KCC to be aware of a GCP resource without managing its lifecycle:

*   **Split Management / Shared Resources:** A central team manages a resource (e.g., a shared VPC Network or a Folder) using one KCC instance (or another tool like Terraform), while tenant teams need to reference that resource in their own namespaces. Tenant controllers need to read the resource's status (e.g., `status.selfLink`, `status.id`) to resolve dependencies but should not be able to modify it.
*   **Migration/Adoption:** Users may want to import existing resources into KCC for visibility or to use as references for other resources, without risking accidental modifications or deletions during the transition period.
*   **Security/Compliance:** Certain critical resources must be modified only through specific pipelines, but their state needs to be visible in the Kubernetes cluster for auditing or policy enforcement.
*   **Multi-Master Scenarios:** In a namespaced mode setup, multiple KCC managers might exist. To prevent conflicts, only one manager should be the "writer" (actuator) for a specific GCP resource, while others can be "readers."

Currently, users can use `cnrm.cloud.google.com/deletion-policy: abandon` to prevent deletion, or `cnrm.cloud.google.com/reconcile-interval-in-seconds: "0"` to stop periodic reconciliation. However, neither fully prevents actuation:
*   `deletion-policy: abandon` still allows updates and creation.
*   `reconcile-interval-in-seconds: "0"` still triggers actuation if the spec changes or upon initial creation.

There is a need for a resource-level control to explicitly disable all actuation (Create, Update, Delete) while maintaining status reconciliation (Read).

## 3. Goals

*   **Explicit Control:** Provide a clear, resource-level annotation to control the actuation mode.
*   **Read-Only Semantics:** Ensure that when a resource is in "read-only" mode, KCC performs **no** write operations (Create, Update, Delete) against the GCP API.
*   **Status Fidelity:** Ensure that KCC continues to update the resource's `status` (including `observedGeneration` and `conditions`) to reflect the actual state of the GCP resource.
*   **Precedence:** Define a clear precedence order for actuation configuration (Resource > Namespace/Context > Cluster/System).

## 4. Non-Goals

*   **Fine-grained Field Control:** This design controls actuation at the whole-resource level, not for individual fields.
*   **Access Control:** This feature does not replace RBAC or IAM. It is a controller behavior configuration, not a security boundary.

## 5. Proposed Solution

We propose introducing a new annotation, `cnrm.cloud.google.com/actuation-mode`, which controls the controller's actuation behavior for the annotated resource.

### 5.1. Annotation

**Key:** `cnrm.cloud.google.com/actuation-mode`

**Values:**

*   `Reconciling` (Default): The controller manages the resource's full lifecycle (Create, Read, Update, Delete).
*   `Paused`: The controller **reads** the resource state from GCP and updates the K8s status, but **skips** any Create, Update, or Delete operations.

### 5.2. Behavior in `Paused` Mode

When a resource is in `Paused` mode:

1.  **Creation:** If the resource exists in K8s but not in GCP, the controller will **not** create it in GCP. It may report a `Ready=False` condition indicating the resource is not found or is unmanaged.
2.  **Updates:** If the resource spec in K8s differs from the state in GCP, the controller will **not** update GCP. It will, however, update the K8s resource's `status` to reflect the GCP state.
3.  **Deletion:** If the resource is deleted from K8s, the controller will **not** delete the resource in GCP. It will simply remove the finalizer (effectively behaving like `deletion-policy: abandon`).
4.  **Status Updates:** The controller will continue to periodically poll GCP (subject to `reconcile-interval-in-seconds`) and update the K8s resource's `status`, including `status.conditions`, `status.observedGeneration`, and resource-specific status fields (e.g., `status.id`, `status.selfLink`).

### 5.3. Precedence

The actuation mode will be determined by checking the following sources in order of precedence:

1.  **Resource Annotation:** `cnrm.cloud.google.com/actuation-mode` on the resource itself.
2.  **ConfigConnectorContext (Namespaced Mode):** `spec.actuationMode` in the `ConfigConnectorContext` object for the resource's namespace.
3.  **ConfigConnector (Cluster Mode):** `spec.actuationMode` in the `ConfigConnector` object.
4.  **Default:** `Reconciling`.

This allows cluster admins to set a default policy (e.g., everything is reconciling), namespace admins to override it (e.g., a "read-only" namespace), and users to override it for specific resources.

## 6. User Experience

### 6.1. Defining a Read-Only Resource

To define a resource that KCC should only observe:

```yaml
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: shared-vpc
  namespace: tenant-a
  annotations:
    cnrm.cloud.google.com/actuation-mode: "Paused"
spec:
  # The spec should match the existing resource, but even if it drifts,
  # KCC will NOT update the actual GCP resource.
  resourceID: existing-shared-vpc
```

### 6.2. Status Reporting

Users can inspect the status to see the observed state:

```yaml
status:
  conditions:
  - type: Ready
    status: "True"
    reason: "UpToDate" # or potentially a new reason "Paused" or "ReadOnly"
    message: "Resource is in Paused actuation mode"
  selfLink: "..."
  observedGeneration: 1
```

## 7. Implementation Details

### 7.1. Modify `DecideActuationMode`

The function `DecideActuationMode` in `pkg/controller/resourceactuation/resourceactuation.go` needs to be updated or wrapped to accept the resource's annotations.

```go
// Current signature
// func DecideActuationMode(cc opv1beta1.ConfigConnector, ccc opv1beta1.ConfigConnectorContext) opv1beta1.ActuationMode

// New signature or logic
func DecideActuationMode(u *unstructured.Unstructured, cc opv1beta1.ConfigConnector, ccc opv1beta1.ConfigConnectorContext) opv1beta1.ActuationMode {
    // 1. Check annotation on u
    if val, ok := k8s.GetAnnotation("cnrm.cloud.google.com/actuation-mode", u); ok {
        return opv1beta1.ActuationMode(val)
    }
    // 2. Fallback to existing logic (CCC, CC, Default)
    // ...
}
```

### 7.2. Controller Logic

In the reconciler (e.g., `pkg/controller/direct/directbase/directbase_controller.go` and DCL/TF controllers), the `DecideActuationMode` result is already checked. We need to ensure that the `Paused` mode logic:

1.  **Skip Actuation:** Does not call `Create`, `Update`, or `Delete` on the underlying adapter/client.
2.  **Continue Status Sync:** Still calls `Get` or `Find` to retrieve the current GCP state and updates the K8s status.

*Note: The current `directbase` implementation for `Paused` mode (seen in `doReconcile`) seems to skip everything, including status updates, or at least it logs "Skipping actuation" and returns. We need to verify and possibly adjust this to ensure *read* operations still happen if the goal is "Read-Only" rather than "Completely Paused".*

**Refinement:**
If `Paused` currently means "do nothing at all" (no read, no write), we might need to distinguish between `Paused` (no-op) and `ReadOnly` (read-only).
However, for this design, we define `Paused` as "No Actuation", which implies we can still read.
If the existing `Paused` implementation completely stops the reconciliation loop, we should change it to allow the `Get` (Read) and `Status Update` paths, but gate the `Create`, `Update`, and `Delete` paths.

**Revised `doReconcile` flow (Conceptual):**

```go
// 1. Fetch State (Read)
gcpResource, err := adapter.Find(ctx)
// ... handle err ...

// 2. Determine Actuation Mode
mode := DecideActuationMode(resource, cc, ccc)

// 3. Update Status (Always, unless completely suspended)
if gcpResource != nil {
    // Update status based on gcpResource
}

// 4. Actuation (Gated by Mode)
if mode == "Paused" {
    // Skip Create/Update/Delete
    // Ensure Finalizers are handled (if deleting, remove finalizer without GCP delete)
    return result, nil
}

// 5. Proceed with Create/Update/Delete if Reconciling...
```

## 8. Testing Strategy

### 8.1. Unit Tests

*   **`DecideActuationMode`:** Verify that the annotation takes precedence over CCC and CC settings. Verify default behavior.
*   **Controller Logic:** Mock the adapter and verify that in `Paused` mode, `Create`/`Update`/`Delete` methods are **not** called, but `Find`/`Get` methods **are** called.

### 8.2. Integration/E2E Tests

*   **Scenario 1: Paused Resource Creation:** Create a resource with `actuation-mode: Paused`. Verify it is not created in GCP.
*   **Scenario 2: Paused Resource Drift:** Create a resource (reconciling), then switch to `Paused`. Modify the GCP resource externally (or via another KCC resource). Verify KCC updates the status to match GCP but does not revert the change.
*   **Scenario 3: Paused Resource Deletion:** Delete a `Paused` resource in K8s. Verify the GCP resource remains.

## 9. Alternatives Considered

### 9.1. Using `management-conflict-prevention-policy`

We considered reusing `cnrm.cloud.google.com/management-conflict-prevention-policy`. However, its values (`resource`, `none`) are semantically tied to conflict resolution (leasing) rather than explicit lifecycle management. Overloading it with `paused` or `read-only` might be confusing.

### 9.2. Using `deletion-policy: abandon` + `reconcile-interval: 0`

As noted in Motivation, this combination is insufficient because `reconcile-interval: 0` does not prevent actuation on spec changes, and `deletion-policy` only controls deletion.

### 9.3. New CRD `ReadOnlyResource`

Creating a separate CRD (e.g., `ComputeNetworkReadOnly`) would be too heavy-handed and require duplicating the entire schema. An annotation is lightweight and works with all existing resources.
