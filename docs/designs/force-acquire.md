# Force Acquire (Safe Import & Verification)

> `Feature State`: `alpha` (Proposed)

Config Connector (KCC) allows users to acquire existing Google Cloud Platform (GCP) resources using the `spec.resourceID` field. Standard acquisition occurs during the first reconciliation: KCC adopts the resource and immediately enforces the desired state defined in the KRM, potentially modifying the GCP resource to match the spec.

This document proposes a "Force Acquire" feature, enabled via the `cnrm.cloud.google.com/force-acquire: "true"` annotation. This feature allows users to safely acquire or verify a resource **at any point in its KRM object lifecycle** without risk of modification.

## Problem Statement & Gap Analysis

Users frequently encounter scenarios where a field on a GCP resource becomes mutable in the underlying GCP API, but Config Connector (KCC) still enforces immutability on that field due to schema lag. This creates a blocking issue where users cannot manage the resource's desired state effectively.

To address this, users need a solution that bridges this gap while satisfying the following requirements:

1.  **Flexibility:** Users need a way to update their KCC manifests to match the true desired state (including changed "immutable" fields) without waiting for a new KCC release to update the schema.
2.  **Safety:** KCC must not become a tool for unsafe or undefined mutations. If KCC's controller logic isn't updated to handle a mutation, forcing it to *perform* the update is risky. Ideally, the user should perform the mutation safely outside of KCC (e.g., via Cloud Console or gcloud), and KCC should then be able to *accept* this new state without fighting it.
3.  **Auditability:** Any operation to bypass standard behavior must be fully auditable. Ad-hoc CLI tools or manual "abandon and re-acquire" workflows often break the audit trail. The intent and the acceptance of the new state should be recorded in the cluster and GitOps history.
4.  **Faster Error Detection:** In GitOps environments, visibility into reconciliation errors is often delayed. It's critical to avoid silent failures or infinite reconciliation loops where KCC might incorrectly try to revert a valid external change.

## Benefits & Use Cases

1.  **Bypass KCC Immutability (Primary Use Case):** Provide a mechanism for users to temporarily manage fields that KCC considers immutable but the underlying GCP API has made mutable. This allows users to leverage new GCP API features immediately. When `force-acquire` is active, the strict verification ensures that the GCP resource matches the KRM's desired state, effectively allowing a "fresh declaration" to be matched against even mutable fields, without KCC making any changes itself.
2.  **Safe Adoption:** Allow users to import existing GCP resources into KCC management without accidental modification. If the KRM spec does not strictly match the live GCP resource, the operation should fail rather than updating the resource.
3.  **Drift Verification:** Enable a "read-only" mode where KCC continuously verifies that the GCP resource matches the KRM spec but strictly prevents any actuation (correction) of drift.

## Proposed Solution: `force-acquire` Annotation

We introduce a new annotation: `cnrm.cloud.google.com/force-acquire`.

When this annotation is set to `"true"` on a KCC resource, the reconciliation logic changes as follows:

1.  **Acquisition Mode (Ignore Mutability):** The controller treats the reconciliation as an acquisition attempt, regardless of whether it's the first or subsequent reconciliation. This means standard mutability checks (which typically prevent changing immutable fields after creation) are skipped during the comparison. The KRM spec is treated as the ground truth definition to match against.
2.  **Strict Verification:** The controller (or webhook, depending on implementation choice) compares the KRM `spec` against the live GCP resource.
    *   **If the GCP resource does not exist:** The reconciliation fails with a `ResourceNotFound` error. KCC **will not** attempt to create the resource.
    *   **If the GCP resource exists but differs from the KRM spec:** The reconciliation fails with an `UpdateFailed` (or similar) error, detailing the mismatch. KCC **will not** attempt to update the resource to fix the drift.
    *   **If the GCP resource exists and matches the KRM spec:** The reconciliation succeeds, and the resource status is updated to `Ready`.
3.  **No Actuation:** Under no circumstances will the controller send a `Create`, `Update`, or `Delete` request to the GCP API while this annotation is active. It effectively puts the controller in a "Verify Only" mode.

## Workflow

### Scenario: Updating an 'Immutable' Field in KCC

A user wants to update a field on a GCP resource (e.g., a Pub/Sub Topic's `messageRetentionDuration`) that is currently marked as immutable in KCC's schema but has recently become mutable in the underlying GCP API. They need to do this without waiting for a KCC release.

1.  **Perform External Mutation:** The user manually updates the field on the live GCP resource using an external tool (e.g., `gcloud pubsub topics update <topic-name> --message-retention-duration=604800s`) or the Cloud Console.
2.  **Update KRM:** The user updates their KCC YAML manifest to reflect this new desired state for the field.
3.  **Apply with Annotation:** The user adds `metadata.annotations["cnrm.cloud.google.com/force-acquire"]: "true"` to their KRM manifest and applies it.
4.  **Observe Status:** KCC detects the (now matching) state of the GCP resource against the KRM's desired state.
    *   **Success:** If the manually updated GCP resource now perfectly matches the KRM `spec`, KCC reports `Ready: True`. The controller effectively "accepts" the externally performed mutation.
    *   **Mismatch (Error):** If there are still discrepancies (e.g., the user's KRM YAML doesn't exactly match the external change, or another field drifted), KCC reports an error, detailing the mismatch, and makes no further changes to GCP.
5.  **Switch to Normal Reconciliation:** Once KCC's schema and controller are updated in a future release to properly support the field's mutability, the user can remove the `cnrm.cloud.google.com/force-acquire` annotation and **apply the updated manifest**. KCC will then manage the field normally, performing updates as specified in the KRM without external intervention.

### Scenario: Safe Import of a Production Resource

A user wants to bring an existing Cloud Storage Bucket under KCC management but wants to ensure their KRM YAML is perfect before KCC takes over, to avoid accidental configuration changes.

1.  **Author KRM:** The user writes the `StorageBucket` YAML, setting `spec.resourceID` to the existing bucket's name.
2.  **Apply with Annotation:** The user adds `metadata.annotations["cnrm.cloud.google.com/force-acquire"]: "true"` and applies the YAML.
3.  **Observe Status:**
    *   **Mismatch:** If the KRM differs from the live bucket (e.g., incorrect `location` or `storageClass`), KCC reports an error describing the mismatch. The live bucket remains untouched.
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

## Implementation Details

*   **Controller Support:** This feature will be supported across all controller types: Direct, Terraform-based, and DCL-based.
*   **Error Reporting:**
    *   **Controller:** Mismatches will be reported in the resource's `status.conditions` (typically setting `Ready=False` with a descriptive `Message`).
    *   **Webhook:** Mismatches will cause the `kubectl apply` request to be rejected immediately with a descriptive error message returned to the client.

## Open Questions

### Where should the live state verification occur?

There is a trade-off between immediate feedback and system reliability when deciding where to perform the `GET` and verification against the live GCP resource.

#### Option 1: In the Controller (Standard Reconciliation)
*   **Description:** The verification logic is implemented as part of the standard reconciliation loop. The controller reads the annotation, fetches the live state, compares it, and updates the `status`.
*   **Pros:**
    *   **Robustness:** Standard Kubernetes pattern. The controller is designed for asynchronous operations and retries.
    *   **No Timeouts:** Long-running GCP API calls won't time out a webhook request.
    *   **Safety:** Decouples the control plane from the data plane (GCP).
    *   **Development Cost: Low.** Reuses the existing reconciliation machinery, client factories, and error handling patterns.
*   **Cons:**
    *   **Delayed Feedback:** The user (or GitOps tool) applies the manifest successfully, but the error only appears later in the resource `status`. This requires additional monitoring or tooling to surface.

#### Option 2: In the Webhook (Immutable Fields Validation)
*   **Description:** The verification logic is injected into the validating webhook, specifically during the immutable fields check. The webhook would fetch the live GCP resource and verify the match *before* accepting the KRM update.
*   **Pros:**
    *   **Immediate Feedback:** If the verification fails (mismatch or not found), the `kubectl apply` command fails immediately with a clear error message. This prevents invalid config from even entering the cluster.
*   **Cons:**
    *   **"Hacky" / Anti-Pattern:** Making external network calls (to GCP) from a synchronous admission webhook is generally discouraged.
    *   **Reliability:** A failure in the GCP API (500s or downtime) blocks the `kubectl apply` operation entirely for annotated resources. The user cannot persist their configuration to the cluster, halting deployment pipelines.
    *   **Latency:** Webhooks are synchronous. Even a simple GET call to GCP typically adds 100-500ms of latency per request. This accumulates with other webhooks in the chain and can significantly degrade the interactive `kubectl apply` experience, especially during cold starts, network contention, or API throttling.
    *   **Timeouts:** Kubernetes admission webhooks have strict timeout limits (default 10s, max 30s). If the GCP API call exceeds this limit due to contention or transient issues, the entire `kubectl apply` fails with a timeout error, regardless of the validity of the config.
        *   **GitOps Mitigation:** While GitOps tools typically retry on timeouts, relying on this behavior for a core feature is risky and can lead to unstable deployment loops.
    *   **Development Cost: High.** Requires implementing a new mechanism to inject GCP clients into the webhook, handling authentication/impersonation within the webhook context (which might differ from the controller's), and carefully managing timeouts and error propagation.
    *   **Testing Complexity: High.** End-to-end (E2E) testing is significantly more difficult. Since the webhook runs within the Kubernetes control plane, configuring it to communicate reliably with external APIs (especially mocks like MockGCP during tests) requires complex network setup. This introduces flakiness and makes it hard to verify behavior in standard CI environments compared to the controller approach.

#### Recommendation

The **Controller (Option 1)** is the recommended approach for the initial implementation and general usage. Its asynchronous nature aligns with Kubernetes best practices, ensuring high availability of the control plane and decoupling it from external API performance. The delayed feedback trade-off is acceptable in most GitOps workflows where status monitoring is already standard.

**Option 2 (Webhook)** should only be considered if there is a critical, specific requirement for synchronous feedback that outweighs the significant reliability risks and development costs—for example, in a specialized interactive tool or a strictly controlled environment where immediate rejection of invalid config is prioritized over system resilience.

## Comparison with Existing Features

| Feature | `force-acquire: "true"` | `actuationMode: Paused` | Standard Acquisition |
| :--- | :--- | :--- | :--- |
| **GCP Interaction** | **Reads** (Verify) | **None** (Stop) | **Reads & Writes** (Manage) |
| **Drift Detection** | **Yes** (Reports Error) | **No** | **Yes** (Auto-corrects) |
| **Actuation (Writes)** | **Blocked** | **Blocked** | **Allowed** |
| **Mutability Checks** | **Ignored** (Like new) | N/A | **Enforced** (after creation) |
| **Primary Use Case** | Safe Import, Audit | Maintenance, Emergency Stop | Normal Management |

## Alternatives Considered (for Bypassing Immutability)

### 1. Force Mutation Annotation (e.g., `force-mutate: "fieldA,fieldB"`)
*   **Description:** Introduce an annotation that lists specific fields where immutability checks should be bypassed, allowing KCC to attempt an update.
*   **Analysis:** While explicit, this is risky. If the underlying API does *not* actually support the update (despite the user's belief), it could lead to undefined behavior or complex errors. It also duplicates the logic of `force-acquire` but with dangerous write permissions. Validating which fields are truly mutable vs. destructive is difficult to do generically.
*   **Auditability:** Good. The intent to mutate specific fields is recorded directly on the object's metadata via the annotation and within GitOps history, providing a clear audit trail in both source control and Kubernetes audit logs.

### 2. Abandon and Re-acquire
*   **Description:** The user sets the `cnrm.cloud.google.com/deletion-policy: "abandon"` annotation, deletes the KCC resource (orphaning the GCP resource), updates the YAML manifest with the new desired state (which is allowed during creation), and re-applies it to "re-acquire" the resource.
*   **Analysis:** This uses existing primitives but is disruptive to GitOps workflows. It requires manual orchestration, breaks the resource's history in Kubernetes, and carries a risk of accidental deletion if the deletion policy is not set correctly. It also causes a temporary loss of visibility and management.
*   **Auditability:** Highly Challenging and Expensive. While GitOps tracks the delete and create commits, the cluster-level resource history is severed. The new resource has a new UID and no direct linkage to the old one within Kubernetes. Correlating the two requires extensive and complex log aggregation (e.g., querying historical audit logs to match the old resource's deletion with the new resource's creation based on the GCP resource ID across different systems), making it very difficult and costly to trace the full lifecycle or debug issues across the transition.

### 3. Power Tool / CLI
*   **Description:** Create a separate CLI tool that can patch the KCC resource status or directly call GCP APIs to force the change outside of the controller's reconciliation loop.
*   **Analysis:** This decouples dangerous operations from the controller but offers a poor developer experience. It fragments the tooling (requiring a separate binary and credentials), is not GitOps-friendly, and is hard to discover. It moves the complexity to the user rather than solving it in the platform.
*   **Auditability:** Highly Challenging and Expensive. If the tool is run ad-hoc by a developer, auditability is poor as it bypasses standard logs. To achieve good auditability, the tool must be wrapped in a CI/CD job with its configuration versioned in Git. However, this incurs a high setup and maintenance cost: it requires building a dedicated pipeline, managing credentials, and maintaining the wrapper scripts, effectively re-implementing a bespoke reconciliation loop just for this one operation.

**Why `force-acquire` is preferred:** It allows the KRM spec to be updated to the true desired state in a safe, declarative manner. The user can perform the mutation via other means (Console/gcloud) if needed, and then simply "verify and accept" the new state in KCC without risking accidental destruction or fighting with the controller. The use of the annotation and KCC's status updates provides a clear and auditable trail of the intent to verify and the outcome of that verification.

## Long-Term Vision

It is important to note that `force-acquire` is primarily a **stop-gap mechanism** to address the schema lag between KCC and the GCP API. It is not intended to be the default way to manage mutable fields.

The true long-term solution is to **minimize the gap** between KCC behavior and GCP API behavior. We aim to achieve this by:

1.  **Automated Pipeline Updates:** Improving our generation pipelines to rapidly detect and reflect API mutability changes in KCC schemas.
2.  **Prioritizing User Requests:** If users file issues requesting specific fields to be made mutable, prioritizing those schema updates is always preferred over directing them to use `force-acquire`. This ensures a cleaner, more standard Kubernetes experience for all users.
