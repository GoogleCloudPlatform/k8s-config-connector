## Technical Design: Stateful Reconciliation with Spec Hashing

### 1. Overview

This document proposes a fundamental improvement to the KCC reconciliation model to reliably detect changes in user intent. The current stateless model, which compares the user's `.spec` against the live GCP state, cannot distinguish between a server-defaulted field and a field a user has intentionally stopped managing. This ambiguity can lead to incorrect controller behavior, such as fighting to unset a field the user wants to be managed by GCP defaults.

To solve this, we will introduce a stateful mechanism. The controller will store a hash of the last successfully applied `.spec` in the resource's `.status`. This provides a lightweight, GitOps-safe historical record, allowing the controller to know with certainty when a user has modified their desired state.

### 2. Goals

*   To reliably and efficiently detect any change to a resource's `.spec`.
*   To provide a definitive trigger for the controller to re-evaluate the resource and calculate a new update plan.
*   To be lightweight, avoiding the storage of large objects in the resource's `.status`.
*   To be fully compatible with GitOps tools (e.g., ArgoCD, Flux) by storing state exclusively in the `.status` subresource.

### 3. Proposed Solution

We will introduce a new field, `lastAppliedSpecHash`, to the `.status` struct of all KCC resources.

#### New Status Field

```go
// Example Status struct
type ResourceStatus struct {
    // ... existing fields like Conditions, ObservedGeneration, etc.

    // LastAppliedSpecHash is a hash of the spec that was last successfully applied.
    // This is used by the controller to detect if the user's desired state has changed.
    // +optional
    LastAppliedSpecHash string `json:"lastAppliedSpecHash,omitempty"`
}
```

#### Reconciliation Logic

The reconciliation loop will be modified to use this hash as the primary signal for user intent changes.

1.  **On Reconciliation:** The controller starts its loop.
2.  **Calculate Current Hash:** It calculates a hash of the `resource.spec` from the object currently in the cluster. The hashing process must be deterministic (e.g., canonical JSON serialization followed by SHA256).
3.  **Compare Hashes:** It compares this newly calculated hash with the value stored in `resource.status.lastAppliedSpecHash`.
4.  **Decision:**
    *   **If the hashes match:** The user's desired state has not changed. The controller can proceed to check for drift from the GCP state (see `fast-path-reconciliation-with-gcp-hash.md`).
    *   **If the hashes do NOT match:** This is a clear and unambiguous signal that the user has changed their intent. The controller must now enter the "slow path" to calculate a precise update. This involves a detailed comparison of the `currentSpec` against the `liveGCPObject` to determine the exact set of fields to create, update, or unset.

#### Status Update

After every successful `CREATE` or `UPDATE` operation sent to GCP, the controller **must** update the resource's status, setting the `lastAppliedSpecHash` field to the hash of the `.spec` that was just applied.

### 4. Interaction with `default-to-gcp-fields`

This design is a prerequisite for the correct implementation of the "unmanaging" a field feature. The `lastAppliedSpecHash` provides the crucial *trigger*.

When the hash check fails because a user has removed a field from their spec, the controller knows it must act. It then consults the `cnrm.cloud.google.com/default-to-gcp-fields` annotation to resolve the ambiguity of *how* to act:
*   If the removed field is in the annotation, the controller ignores it.
*   If the removed field is NOT in the annotation, the controller unsets it in GCP.

### 5. Alternatives Considered

*   **Storing Full `lastAppliedSpec` in `.status`:** This would work but would significantly increase the size of every KCC resource in etcd, which is inefficient at scale.
*   **Using `metadata.annotations`:** Storing state in annotations is not GitOps-safe, as GitOps controllers will revert any changes made by the KCC controller to the live object's metadata.
