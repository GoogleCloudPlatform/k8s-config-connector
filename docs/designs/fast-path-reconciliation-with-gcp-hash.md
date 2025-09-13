## Technical Design: Fast-Path Reconciliation with GCP Object Hashing

### 1. Overview

This document proposes a performance optimization for the KCC reconciliation loop, building upon the stateful model introduced in `stateful-reconciliation-with-spec-hash.md`.

In a stable system, the vast majority of reconciliations are no-ops: the user's spec has not changed, and the resource in GCP has not drifted. The controller's primary job is to verify this state as efficiently as possible. Performing a full, field-by-field semantic comparison between the spec and the live GCP state on every cycle is computationally expensive and unnecessary.

This design introduces a "fast path" using a hash of the live GCP object. By comparing the hash of the current GCP object with a stored hash from the previous successful reconciliation, the controller can almost instantly determine if any drift has occurred, avoiding the expensive comparison logic in the common case.

### 2. Goals

*   To create a highly efficient "fast path" for the reconciliation loop.
*   To avoid expensive, field-by-field semantic diffs when the resource has not changed.
*   To improve the overall performance and reduce the CPU overhead of KCC controllers at scale.

### 3. Prerequisite

This design assumes the implementation of the `lastAppliedSpecHash` feature. The fast-path check for GCP drift should only occur *after* the controller has verified that the user's intent has not changed.

### 4. Proposed Solution

We will introduce a second field to the `.status` struct of all KCC resources: `lastAppliedGCPHash`.

#### New Status Field

```go
// Example Status struct
type ResourceStatus struct {
    // ... existing fields
    LastAppliedSpecHash string `json:"lastAppliedSpecHash,omitempty"`

    // LastAppliedGCPHash is a hash of the GCP resource state that was observed
    // after the last successful reconciliation. This is used for fast-path drift detection.
    // +optional
    LastAppliedGCPHash string `json:"lastAppliedGCPHash,omitempty"`
}
```

#### Reconciliation Logic

The fast-path check is the second step in the stateful reconciliation process.

1.  **Check Intent:** The controller first compares `hash(currentSpec)` with `status.lastAppliedSpecHash`. If they differ, the fast path is aborted, and the controller proceeds directly to calculating a full update.
2.  **Fast-Path Drift Detection:** If the spec hashes match, the controller proceeds to check for GCP drift:
    a.  It makes a `GET` call to fetch the `liveGCPObject`.
    b.  It calculates a normalized hash of this `liveGCPObject`.
    c.  It compares this new hash with the value in `status.lastAppliedGCPHash`.
3.  **Decision:**
    *   **If the hashes match:** The controller can immediately conclude that the resource is in sync. The reconciliation loop terminates successfully. This is the "fast path".
    *   **If the hashes do NOT match:** This signals that *something* has changed in GCP. The controller must now proceed to the "slow path": a full, semantic, field-by-field comparison of the `currentSpec` against the `liveGCPObject` to determine if the drift is meaningful (i.e., it affected a managed field) or benign (e.g., a change to an output-only field).

#### Status Update

After every successful `CREATE` or `UPDATE` operation, the controller must update the resource's status with both hashes:
1.  `lastAppliedSpecHash` is set to the hash of the spec that was just applied.
2.  `lastAppliedGCPHash` is set to the normalized hash of the full GCP object that was returned by the successful API call.

### 5. Normalization

A simple hash of the raw GCP response is insufficient, as many responses contain volatile data (e.g., `etag`, `lastModifiedTimestamp`, `selfLink`) that changes on every `GET` call without representing a change in the resource's configuration.

Before hashing the `liveGCPObject`, the controller must first run it through a **normalization** function. This function will create a temporary copy of the object and remove or zero-out any known volatile or purely informational fields. This ensures that the hash only represents the actual configuration of the resource. This normalization logic may need to be tailored for each resource type.

### 6. Summary

The `lastAppliedGCPHash` is a pure performance optimization. It allows the controller to handle the most common case (no changes) with maximum efficiency, reserving the expensive and complex semantic diffing logic for the rare cases where a change has actually occurred.
