## Technical Design: Stateful Reconciliation with a Coalesced Hash Cookie

### 1. Overview

This document proposes a fundamental improvement to the KCC reconciliation model to reliably and efficiently detect both changes in user intent and drift in the live GCP resource state. The current stateless model, which compares the user's `.spec` against the live GCP state, cannot distinguish between a server-defaulted field and a field a user has intentionally stopped managing. This ambiguity can lead to incorrect controller behavior. Furthermore, without a lightweight drift detection mechanism, the controller must perform expensive semantic diffs on every reconciliation cycle.

To solve this, we will introduce a stateful mechanism using a single, coalesced "cookie". The controller will store two hashes in a single `.status` field: one for the last successfully applied `.spec` and one for the last observed state of the live GCP resource. This provides a lightweight, GitOps-safe historical record, allowing the controller to quickly determine if the user's desired state has changed, if the GCP resource has drifted, or if both are in sync.

### 2. Goals

*   To reliably and efficiently detect any change to a resource's `.spec`.
*   To reliably and efficiently detect drift between the KCC resource and the live GCP resource.
*   To provide a definitive trigger for the controller to re-evaluate the resource and calculate a new update plan, avoiding expensive computations when no changes have occurred.
*   To be lightweight, avoiding the storage of large objects in the resource's `.status`.
*   To be fully compatible with GitOps tools (e.g., ArgoCD, Flux) by storing state exclusively in the `.status` subresource.

### 3. Proposed Solution

We will introduce a new field, `lastModifiedCookie`, to the `.status` struct of all KCC resources. This field will store two hashes, concatenated with a `/`.

#### New Status Field

```go
// Example Status struct
type ResourceStatus struct {
    // ... existing fields like Conditions, ObservedGeneration, etc.

    // LastModifiedCookie contains hashes of the last applied spec and the last observed GCP state.
    // The format is "<spec-hash>/<gcp-hash>".
    // This is used by the controller to detect if the user's desired state has changed or if the GCP resource has drifted.
    // +optional
    LastModifiedCookie string `json:"lastModifiedCookie,omitempty"`
}
```

#### Reconciliation Logic

The reconciliation loop will be modified to use this cookie as the primary signal for changes.

1.  **On Reconciliation:** The controller starts its loop.
2.  **Calculate Current Hashes:**
    *   It calculates a hash of the `resource.spec` from the object currently in the cluster. The hashing process must be deterministic (e.g., canonical JSON serialization followed by SHA256).
    *   It fetches the live GCP object and calculates a normalized hash of its state. Normalization involves removing or zeroing-out volatile or server-generated fields (e.g., `etag`, `lastModifiedTimestamp`, `selfLink`) to ensure the hash only reflects user-configurable fields.
3.  **Compare Hashes:** It parses the `resource.status.lastModifiedCookie` to extract the previously stored spec and GCP hashes.
4.  **Decision:**
    *   **If the new spec hash does NOT match the stored spec hash:** This is a clear signal that the user has changed their intent. The controller must enter the "slow path" to calculate a precise update.
    *   **If the spec hashes match, but the new GCP hash does NOT match the stored GCP hash:** This indicates that the GCP resource has drifted from its last known state. The controller must enter the "slow path" to reconcile the drift.
    *   **If both hashes match:** The user's desired state and the live GCP state are in sync with the last applied configuration. The controller can terminate the reconciliation early (the "fast path").

#### Status Update

After every successful `CREATE` or `UPDATE` operation sent to GCP, the controller **must** update the resource's status. It will construct a new cookie by concatenating the hash of the just-applied `.spec` and the hash of the full GCP object returned by the API call, and set the `lastModifiedCookie` field to this new value.

### 4. Interaction with `default-to-gcp-fields`

This design is a prerequisite for the correct implementation of the "unmanaging" a field feature. The spec hash part of the cookie provides the crucial *trigger*.

When the spec hash check fails because a user has removed a field from their spec, the controller knows it must act. It then consults the `cnrm.cloud.google.com/default-to-gcp-fields` annotation to resolve the ambiguity of *how* to act:
*   If the removed field is in the annotation, the controller ignores it.
*   If the removed field is NOT in the annotation, the controller unsets it in GCP.

### 5. Alternatives Considered

*   **Storing Full `lastAppliedSpec` in `.status`:** This would work but would significantly increase the size of every KCC resource in etcd, which is inefficient at scale.
*   **Using two separate hash fields in `.status`:** Instead of a single cookie, we could have used `lastAppliedSpecHash` and `lastAppliedGCPHash`. A single field was chosen for conciseness, though two fields would be slightly more explicit and readable. The core logic remains the same.
*   **Using `metadata.annotations`:** Storing state in annotations is not GitOps-safe, as GitOps controllers will revert any changes made by the KCC controller to the live object's metadata.
