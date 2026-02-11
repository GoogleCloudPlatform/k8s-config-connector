# Force Acquire (Safe Import & Verification)

> `Feature State`: `alpha` (Proposed)

Config Connector (KCC) allows users to acquire existing Google Cloud Platform (GCP) resources using the `spec.resourceID` field. Standard acquisition occurs during the first reconciliation: KCC adopts the resource and immediately enforces the desired state defined in the KRM, potentially modifying the GCP resource to match the spec.

This document proposes a "Force Acquire" feature, enabled via the `cnrm.cloud.google.com/force-acquire: "true"` annotation. This feature allows users to safely acquire or verify a resource **without risk of modification**.

## Goals

1.  **Safe Adoption:** Allow users to import existing GCP resources into KCC management without accidental modification. If the KRM spec does not strictly match the live GCP resource, the operation should fail rather than updating the resource.
2.  **Drift Verification:** Enable a "read-only" mode where KCC continuously verifies that the GCP resource matches the KRM spec but strictly prevents any actuation (correction) of drift.
3.  **Bypass KCC Immutability (Temporary):** Provide a mechanism for users to temporarily manage fields that KCC considers immutable but the underlying GCP API has made mutable. This allows users to leverage new GCP API features immediately without waiting for KCC to update its schema and controller logic. When `force-acquire` is active, the strict verification ensures that the GCP resource matches the KRM's desired state, effectively allowing a "fresh declaration" to be matched against even mutable fields, without KCC making any changes itself.
4.  **Ignore Mutability:** Allow acquisition even if the KRM spec differs from the GCP resource in immutable fields (treating the KRM as a "fresh" declaration for matching purposes).

## Proposed Solution: `force-acquire` Annotation

We introduce a new annotation: `cnrm.cloud.google.com/force-acquire`.

When this annotation is set to `"true"` on a KCC resource, the reconciliation logic changes as follows:

1.  **Acquisition Mode (Ignore Mutability):** The controller treats the reconciliation as an acquisition attempt, regardless of whether it's the first or subsequent reconciliation. This means standard mutability checks (which typically prevent changing immutable fields after creation) are skipped during the comparison. The KRM spec is treated as the ground truth definition to match against.
2.  **Strict Verification:** The controller compares the KRM `spec` against the live GCP resource.
    *   **If the GCP resource does not exist:** The reconciliation fails with a `ResourceNotFound` error. KCC **will not** attempt to create the resource.
    *   **If the GCP resource exists but differs from the KRM spec:** The reconciliation fails with an `UpdateFailed` (or similar) error, detailing the mismatch. KCC **will not** attempt to update the resource to fix the drift.
    *   **If the GCP resource exists and matches the KRM spec:** The reconciliation succeeds, and the resource status is updated to `Ready`.
3.  **No Actuation:** Under no circumstances will the controller send a `Create`, `Update`, or `Delete` request to the GCP API while this annotation is active. It effectively puts the controller in a "Verify Only" mode.

## Workflow

### Scenario: Safe Import of a Production Resource

A user wants to bring an existing Cloud Storage Bucket under KCC management but wants to ensure their KRM YAML is perfect before KCC takes over, to avoid accidental configuration changes.

1.  **Author KRM:** The user writes the `StorageBucket` YAML, setting `spec.resourceID` to the existing bucket's name.
2.  **Apply with Annotation:** The user adds `metadata.annotations["cnrm.cloud.google.com/force-acquire"]: "true"` and applies the YAML.
3.  **Observe Status:**
    *   **Mismatch:** If the KRM differs from the live bucket (e.g., incorrect `location` or `storageClass`), KCC reports an error in `status.conditions` describing the mismatch. The live bucket remains untouched.
    *   **Correction:** The user updates the KRM YAML to match the error report and re-applies.
4.  **Success:** Once the KRM matches, KCC reports `Ready: True`.
5.  **Start Management:** The user removes the `cnrm.cloud.google.com/force-acquire` annotation. KCC now manages the resource normally, and will enforce future updates.

### Scenario: Drift Monitoring (Read-Only)

A user wants to use KCC to define the "Golden State" of their infrastructure and be alerted if the live infrastructure drifts, but they do not want KCC to automatically revert changes (perhaps they use another tool for emergency fixes).

1.  **Apply with Annotation:** The user applies their KCC resources with `force-acquire: "true"`.
2.  **Monitor:** The user sets up alerts on KCC resource status.
3.  **Drift Event:** If someone modifies a resource manually in the Cloud Console, KCC detects the mismatch during the next reconciliation.
4.  **Alerting:** Because `force-acquire` prevents actuation, KCC reports an error (Drift detected) instead of fixing it. The resource status goes to `NotReady`.
5.  **Resolution:** The user receives the alert and decides whether to update the KCC YAML (accepting the change) or manually revert the GCP change.

## Comparison with Existing Features

| Feature | `force-acquire: "true"` | `actuationMode: Paused` | Standard Acquisition |
| :--- | :--- | :--- | :--- |
| **GCP Interaction** | **Reads** (Verify) | **None** (Stop) | **Reads & Writes** (Manage) |
| **Drift Detection** | **Yes** (Reports Error) | **No** | **Yes** (Auto-corrects) |
| **Actuation (Writes)** | **Blocked** | **Blocked** | **Allowed** |
| **Mutability Checks** | **Ignored** (Like new) | N/A | **Enforced** (after creation) |
| **Primary Use Case** | Safe Import, Audit | Maintenance, Emergency Stop | Normal Management |

## Implementation Details

*   **Controller Support:** This feature will be supported across all controller types: Direct, Terraform-based, and DCL-based.
*   **Error Reporting:** Mismatches should be reported clearly in the `status.conditions` (typically strictly creating a `Ready` condition with `False` status and a descriptive `Message`).

## Alternatives Considered (for Bypassing Immutability)

When addressing the specific goal of updating fields that are immutable in KCC but mutable in GCP, several alternatives were considered:

### 1. Abandon and Re-acquire
*   **Description:** The user sets the `cnrm.cloud.google.com/deletion-policy: "abandon"` annotation, deletes the KCC resource (orphaning the GCP resource), updates the YAML manifest with the new desired state (which is allowed during creation), and re-applies it to "re-acquire" the resource.
*   **Analysis:** This uses existing primitives but is disruptive to GitOps workflows. It requires manual orchestration, breaks the resource's history in Kubernetes, and carries a risk of accidental deletion if the deletion policy is not set correctly. It also causes a temporary loss of visibility and management.
*   **Auditability:** Poor. While GitOps can track the delete and create commits, the continuous history of the resource in Kubernetes is broken, making it difficult to correlate the new resource with the old one and trace the full evolution of the configuration through cluster events and logs.

### 2. Force Mutation Annotation (e.g., `force-mutate: "fieldA,fieldB"`)
*   **Description:** Introduce an annotation that lists specific fields where immutability checks should be bypassed, allowing KCC to attempt an update.
*   **Analysis:** While explicit, this is risky. If the underlying API does *not* actually support the update (despite the user's belief), it could lead to undefined behavior or complex errors. It also duplicates the logic of `force-acquire` but with dangerous write permissions. Validating which fields are truly mutable vs. destructive is difficult to do generically.
*   **Auditability:** Good. The intent to mutate specific fields is recorded directly on the object's metadata via the annotation and within GitOps history, providing a clear audit trail in both source control and Kubernetes audit logs.

### 3. Power Tool / CLI
*   **Description:** Create a separate CLI tool that can patch the KCC resource status or directly call GCP APIs to force the change outside of the controller's reconciliation loop.
*   **Analysis:** This decouples dangerous operations from the controller but offers a poor developer experience. It fragments the tooling (requiring a separate binary and credentials), is not GitOps-friendly, and is hard to discover. It moves the complexity to the user rather than solving it in the platform.
*   **Auditability:** Poor. Changes made via an external tool typically bypass both the GitOps workflow (no commit reflecting the change) and the standard Kubernetes audit logs, making it very difficult to track who made what change and why, unless the tool implements its own robust and persistent logging.

**Why `force-acquire` is preferred:** It allows the KRM spec to be updated to the true desired state in a safe, declarative manner. The user can perform the mutation via other means (Console/gcloud) if needed, and then simply "verify and accept" the new state in KCC without risking accidental destruction or fighting with the controller. The use of the annotation and KCC's status updates provides a clear and auditable trail of the intent to verify and the outcome of that verification.
