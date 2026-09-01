# Pause Actuation at Resource Level

## 1. Overview

In Kubernetes Config Connector (KCC), controllers continuously reconcile the desired state defined in Kubernetes (`spec`) against the live state in Google Cloud Platform (GCP). Under default `controller-runtime` semantics, any error returned during reconciliation causes the resource to be requeued with exponential backoff and retried indefinitely.

While continuous retries are essential for transient failures (e.g., temporary network blips, eventual consistency propagation, server-side 5xx errors), they become infinite reconciliation loops when encountering permanent client errors (such as `InvalidArgument` or `Unimplemented`) or KCC diffing bugs (such as false-positive diffs due to missing API domain knowledge).

These tight loops exhaust the underlying GCP API quota across customer projects, flood operator logs, and frequently escalate into critical on-call operational burdens. To stop infinite reconciliation loops and provide targeted incident mitigation, we decided that introducing **resource-level pause** capabilities is the most effective approach. This enables KCC to automatically halt retries on non-retryable errors and provides users with fine-grained actuation control to mitigate operational impact.

## 2. Use Case Deep Dive

Resource-level actuation pause supports two essential use cases:

1. **Automatic Pause on Non-Retryable Errors (KCC-Initiated):**
   * **When** reconciliation fails due to a permanent/non-retryable client error (such as gRPC `InvalidArgument` / `Unimplemented` or HTTP `400` with an invalid argument error message / HTTP `501`; see [Appendix 8.1](#81-detailed-categorization-of-non-retryable-vs-retryable-errors) for full categorization), **KCC should** update `status.conditions` to reflect the error, mark actuation as paused on that resource, and halt workqueue retries to stop API quota burn and log spam. **When** the user updates the `spec`, **KCC should** automatically clear the pause and resume reconciliation.

2. **Manual Pause for Mitigation & Operational Control (User-Initiated):**
   * **When** a resource is stuck in an update loop (e.g., false-positive diffs), failing during an incident, or undergoing maintenance, **users can** manually pause actuation on that specific resource to immediately halt GCP actuation without affecting other resources in the namespace or cluster. **When** ready, **users can** manually resume actuation.

---

## 3. Evaluation of Existing Options & Gaps

| Existing Mechanism | Behavior | Identified Gaps & Limitations |
| :--- | :--- | :--- |
| **`spec.actuationMode: Paused`** in `ConfigConnector` / `ConfigConnectorContext` | Halts all reconciliation across the entire cluster or namespace. | **Granularity Gap:** Coarse-grained, all-or-nothing blast radius. Cannot pause a single problematic resource without stopping all reconciliation in the cluster or namespace. |
| **`reconcile-interval-in-seconds: "0"`** | Disables periodic drift correction when the resource is `UpToDate`. | **Failure Mode Gap:** Only takes effect when `status.conditions[Ready].status == "True"`. If a resource fails or enters an error state (`Ready == False`), the `controller-runtime` workqueue continues exponential backoff retries indefinitely. |
| **`deletion-policy: "abandon"`** | Allows deleting the KRM resource from Kubernetes without deleting the underlying GCP cloud resource. | **GitOps & Operational UX Gap:** Requires an imperative multi-step sequence (annotate with `abandon`, wait for the annotation to sync to the cluster, then prune the YAML from Git). If the file is deleted from Git before the `abandon` annotation syncs to the cluster, GitOps tools will delete the live resource without the abandon policy, risking accidental cloud resource deletion. |

---

## 4. Proposed Solution: Resource-Level Actuation Mode

To extend the actuation concept down to individual resources, we introduce the annotation:
`cnrm.cloud.google.com/actuation-mode: "Reconciling" | "Paused" | "AutoPaused"`

When specified on an individual resource, this annotation overrides the cluster-level (`ConfigConnector.spec.actuationMode`) and namespace-level (`ConfigConnectorContext.spec.actuationMode`) configurations, following [KCC's standard configuration hierarchy](#82-kccs-standard-configuration-hierarchy) to control custom resource behaviors while cleanly distinguishing between **user-initiated manual pause** and **KCC-initiated automated pause**.

### 4.1. Supported Actuation Modes & Semantics

| `cnrm.cloud.google.com/actuation-mode` | Initiator | Behavior on Reconciliation | Behavior on `.spec` Update (`generation > observedGeneration`) |
| :--- | :--- | :--- | :--- |
| **`"Paused"`** | User | Freezes all GCP actuation (`Create`, `Update`, and `Delete`) and periodic drift correction. Deletions are halted to protect cloud resources. | **Remains Paused.** Does not auto-resume; requires explicit user action to unpause. |
| **`"AutoPaused"`** | KCC | Set automatically upon encountering a non-retryable error. Skips GCP `Create`/`Update` actuation and halts workqueue retries. Deletion is always processed. | **Auto-Resumes.** KCC detects the spec modification, automatically clears the annotation, and reconciles the fix. |
| **`"Reconciling"`** | User | Actively reconciles the resource against GCP (`Create`, `Update`, and `Delete`). | Standard active reconciliation. Can be used to override a parent namespace or cluster-level `"Paused"` mode for a specific resource. |
| *(Unset)* | Default | Inherits actuation mode from the namespace (`ConfigConnectorContext`) or cluster (`ConfigConnector`), defaulting to `"Reconciling"`. | Standard inheritance. |

---

### 4.2. Mode 1: Automatic Pause on Non-Retryable Errors (`AutoPaused`)

When reconciliation fails due to a non-retryable (terminal) client error—such as `codes.InvalidArgument`, `codes.Unimplemented`, HTTP `400` with invalid arguments, or HTTP `501` (see [Appendix 8.1](#81-detailed-categorization-of-non-retryable-vs-retryable-errors) for the complete taxonomy)—retrying with the same configuration against the same API state is guaranteed to fail. KCC automatically halts workqueue retries and places the resource into `AutoPaused` mode.

#### 4.2.1. Error Inspection Logic (Example Implementation)
To accurately classify errors across client libraries (gRPC, Google REST clients, standard wrapped Go errors), a centralized error evaluator can be implemented in `pkg/controller/lifecyclehandler`. The following is an illustrative example implementation:

```go
func IsNonRetryableError(err error) bool {
    if err == nil {
        return false
    }

    // 1. Check gRPC status code
    if s, ok := status.FromError(err); ok {
        switch s.Code() {
        case codes.InvalidArgument, codes.Unimplemented:
            return true
        }
    }

    // 2. Check Google API Error (REST/gax-go)
    var apiErr *apierror.APIError
    if errors.As(err, &apiErr) {
        if apiErr.HTTPCode() == http.StatusBadRequest || apiErr.HTTPCode() == http.StatusNotImplemented {
            return true
        }
    }

    var googleAPIErr *googleapi.Error
    if errors.As(err, &googleAPIErr) {
        if googleAPIErr.Code == http.StatusBadRequest || googleAPIErr.Code == http.StatusNotImplemented {
            return true
        }
    }

    return false
}
```

#### 4.2.2. Auto-Pause State Machine & Auto-Resume Flow

1. **Failure Handling (Entering `AutoPaused`):**
   * When `adapter.Create()` or `adapter.Update()` returns an error, the controller evaluates `lifecyclehandler.IsNonRetryableError(err)`.
   * The controller calls `handleUpdateFailed(ctx, u, err)` to record `status.conditions[Ready]=False` and update `status.observedGeneration = metadata.generation`.
   * If non-retryable:
     1. The controller annotates the resource (using `retry.RetryOnConflict` or `client.Patch` to handle metadata conflicts safely):
        ```yaml
        metadata:
          annotations:
            cnrm.cloud.google.com/actuation-mode: "AutoPaused"
        ```
     2. The controller ensures status clearly communicates the auto-pause state alongside the root-cause failure (e.g. `reason: UpdateFailedNonRetryable`, with message: `Update failed: <GCP Error Message>. Actuation is automatically paused to halt retry loops; modify .spec to resume reconciliation.`).
     3. The controller halts further workqueue retries by returning `reconcile.Result{}, nil` to `controller-runtime`.
2. **Reconciliation Check & Auto-Resume:**
   * We extend KCC's existing actuation mode evaluation logic (in `pkg/controller/resourceactuation` and reconcilers like `directbase`, `tf`, `dcl`):
     * When `Reconcile()` is triggered for an object in `AutoPaused`:
       * Check if `metadata.deletionTimestamp` is set. If the user deletes a resource that failed validation, **proceed immediately with deletion and finalizer cleanup**.
       * Compare `metadata.generation` with `status.observedGeneration`:
         * **If `metadata.generation > status.observedGeneration`:** The user modified the `.spec` (e.g. fixed the typo in Git). The controller automatically clears the `cnrm.cloud.google.com/actuation-mode` annotation, logs an informational message, and proceeds with full reconciliation.
         * **If `metadata.generation == status.observedGeneration`:** The resource remains in the terminal error state. The controller exits with `reconcile.Result{}, nil`.

---

### 4.3. Mode 2: Manual Per-Resource Actuation Pause (`Paused`)

To empower users during incident mitigation, maintenance, or troubleshooting, users can explicitly pause actuation on any individual resource at any time:

```yaml
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeInstance
metadata:
  name: my-instance
  annotations:
    cnrm.cloud.google.com/actuation-mode: "Paused"
```

#### 4.3.1. Behavior When Manually Paused
* **All Actuation Halted:** Consistent with `ConfigConnector` and `ConfigConnectorContext` `spec.actuationMode: "Paused"`, setting `actuation-mode: "Paused"` freezes **all actuation against GCP (`Create`, `Update`, and `Delete`)** and stops periodic drift correction.
* **Deletion Safety Protection:** If a resource is deleted (`kubectl delete`) while `actuation-mode: "Paused"`, KCC ensures finalizers remain in place and halts deletion against GCP, preserving the underlying cloud resource (see [Appendix 8.3](#83-event-dropping--unpause-latency-analysis) for detailed lifecycle analysis).
* **Spec Modification Isolation:** Unlike `AutoPaused`, modifying the `.spec` will **not** clear a manual `"Paused"` annotation. The resource remains paused until explicitly unpaused.
* **Status Untouched (Retains Prior Data):** The controller does not fetch live state from GCP nor write to `.status` while paused.

#### 4.3.2. Resuming Actuation
Users can resume reconciliation at any time by:
1. Removing the `cnrm.cloud.google.com/actuation-mode` annotation, or
2. Setting `cnrm.cloud.google.com/actuation-mode: "Reconciling"`.

#### 4.3.3. Visibility & Warning on Manual Pause
Because setting a resource to `"Paused"` halts all actuation indefinitely (including `Delete`), KCC can return a non-blocking **Validating Admission Webhook Warning** (`admissionv1.AdmissionResponse.Warnings`) to provide immediate feedback:
* **On `kubectl apply` or `kubectl delete` with `actuation-mode: "Paused"`:**
  ```text
  Warning: Resource 'my-instance' has cnrm.cloud.google.com/actuation-mode: "Paused". All actuation against GCP (including Create, Update, and Delete) is halted until unpaused.
  ```

---

## 5. Options & Alternatives Considered

### 5.1. Option 1: Single Boolean Annotation (`cnrm.cloud.google.com/actuation-paused: "true"`)

Use a single boolean annotation for both automatic error pausing and manual user pausing.

* **Pros:**
  * **Seamless GitOps Auto-Resume:** Any user edit to `.spec` automatically increments `.metadata.generation` via the Kubernetes API server. When `metadata.generation > status.observedGeneration`, the controller automatically unpauses and applies the fix without requiring the user to know about or manually remove internal annotations.
  * **Durable & Introspectable on the KRM Object:** Storing the pause state directly on the resource (`metadata.annotations` and `status`) ensures it persists across controller pod restarts and leader failovers (preventing sudden retry storms), while remaining fully observable and queryable via `kubectl describe` or `kubectl get -o yaml`.
  * **No CRD Schema Changes:** Works out-of-box using standard Kubernetes metadata annotations and existing status conventions across all KCC CRDs.

* **Cons & Trade-offs:**
  * **Unable to Pause Actuation on Spec Changes (Even with Manual Annotation):** Because the reconciler automatically clears `actuation-paused` whenever `metadata.generation > status.observedGeneration` (to support automatic resume after fixing errors), this design cannot pause actuation if a user modifies the `.spec` while intending for actuation to remain paused. Any `.spec` change increments `metadata.generation` and forces the controller to clear the pause annotation and actuate against GCP.
  * **Potential GitOps Sync Fighting:** If a GitOps engine (e.g., ArgoCD, Flux) is configured with aggressive drift correction on annotations, it might attempt to prune the dynamically added `actuation-paused` annotation from the cluster.
  * **Manual Intervention for Out-of-Band Fixes:** If an error is resolved by an out-of-band change in GCP (e.g. enabling a GCP API service) rather than editing the KRM `spec`, `metadata.generation` will not change. The user must manually remove the annotation to force a retry.
  * **Outcome:** Rejected in favor of the multi-state `cnrm.cloud.google.com/actuation-mode: "Paused" | "AutoPaused" | "Reconciling"` design, which cleanly separates manual pause from automatic error pause.

### 5.2. Option 2: Workqueue Max Retry Limit (Dropping Key After Max Retries)

Configure `controller-runtime` to stop requeueing a key after reaching a maximum retry count (e.g. 5 attempts).

* **Cons & Trade-offs:**
  * Violates the level-triggered, eventual consistency model expected of Kubernetes controllers.
  * Does not persist failure state across controller pod restarts. If the controller pod restarts, the retry count resets to zero, causing an immediate burst of API calls.
  * Lacks visibility: Users cannot see why retries stopped without searching controller logs.

### 5.3. Option 3: Admission Webhook Validation

Catch all `InvalidArgument` errors at admission time using validating webhooks.

* **Cons & Trade-offs:**
  * Admission webhooks cannot validate GCP-side dynamic constraints, project quotas, or asynchronous backend invariants that only GCP APIs evaluate at runtime.
  * Webhooks do not help when GCP deprecates or alters API behavior out-of-band.

---

## 6. Open Questions

### 6.1. Dynamic Annotation Management vs. Status-Driven Error Pause in GitOps

* **The Problem:** When KCC dynamically sets `cnrm.cloud.google.com/actuation-mode: "AutoPaused"` on a live resource upon encountering a non-retryable error, GitOps engines (such as Config Sync, ArgoCD, or Flux) that enforce Git as the sole source of truth may treat this uncommitted annotation as drift and prune it.
* **The Thrashing Loop:**
  1. KCC sets `actuation-mode: "AutoPaused"` after a non-retryable error from GCP.
  2. GitOps detects drift on metadata annotations and removes `actuation-mode: "AutoPaused"`.
  3. The annotation deletion fires a Kubernetes Watch event, triggering KCC to reconcile again.
  4. GCP API rejects the unchanged spec with the same non-retryable error.
  5. KCC re-adds `actuation-mode: "AutoPaused"`.
  6. GitOps removes it again $\rightarrow$ creating an infinite cyclic fighting loop that burns CPU, API quota, and floods operator logs.
* **Alternative Under Consideration: Status-Driven Error Pause:**
  * Restrict `cnrm.cloud.google.com/actuation-mode: "Paused" | "Reconciling"` strictly to user-managed declarative annotations for manual pause/unpause.
  * For automated non-retryable error pauses, track the pause state purely within `status.conditions[Ready]` (with `reason: "UpdateFailedNonRetryable"` and `observedGeneration == metadata.generation`) without mutating `metadata.annotations`.
  * Because GitOps engines never manage `.status`, this eliminates annotation thrashing completely while preserving seamless auto-resume when `.spec` is updated in Git (`generation > observedGeneration`).
  * **Trade-off & Operational Implication:**
    * **Requires `.spec` Update to Retrigger:** Because the pause is evaluated against `metadata.generation == status.observedGeneration`, modifying metadata (like labels or non-actuation annotations) will not trigger a retry.
    * If an issue is resolved out-of-band in GCP without changing `.spec` (e.g., enabling a GCP service API or updating external IAM roles), users are forced to either modify `.spec` (e.g., a dummy field or whitespace edit in tools that increment generation) or explicitly apply `cnrm.cloud.google.com/actuation-mode: "Reconciling"` to force a retry.

### 6.2. Behavior on Non-Retryable Error When Actuation Mode is Explicitly `"Reconciling"`

* **The Dilemma:** When an operator explicitly applies `cnrm.cloud.google.com/actuation-mode: "Reconciling"` (e.g. to override a parent namespace pause or force reconciliation), they express an explicit declarative intent: *"I want this resource to actively reconcile against GCP."* If KCC encounters a permanent client error (e.g., `InvalidArgument` / `400`), automatically switching the annotation to `"AutoPaused"` violates the user's explicit declarative intent. Conversely, ignoring the error and continuing tight retries burns GCP API quota and spams logs.

* **Current Proposed Solution (Section 4):**
  * When a non-retryable error is encountered, KCC unilaterally overwrites `metadata.annotations["cnrm.cloud.google.com/actuation-mode"]` to `"AutoPaused"` (even if it was previously set to `"Reconciling"`) and returns `(reconcile.Result{}, nil)`.
  * The resource remains paused until the user modifies `.spec` (`generation > observedGeneration`), which automatically clears `"AutoPaused"` and resumes reconciliation.
  * **Limitation:** Mutating the user's explicit `"Reconciling"` annotation directly overrides declarative user intent and can confuse operators who expect user-applied annotations to remain untouched.

* **Alternative Options Under Consideration:**
  * **Option A: Strict User Intent (Continuous Exponential Backoff)**
    * KCC respects the explicit `"Reconciling"` annotation and does **not** auto-pause. KCC returns `(Result{}, err)` to `controller-runtime`, continuing exponential backoff retries.
    * *Pros:* Preserves strict declarative intent. If an out-of-band GCP change resolves the error, KCC automatically succeeds without requiring user intervention.
    * *Cons:* Continues infinite retry loops and API quota exhaustion for genuinely broken configurations until the user removes the `"Reconciling"` annotation.
  * **Option B: Degraded Requeue / Low-Frequency Polling (Slow Retry)**
    * When a non-retryable error occurs on an explicitly `"Reconciling"` resource, KCC reports `Ready: False` (`reason: UpdateFailedNonRetryable`), but instead of tight exponential retries or complete halting, it requeues with a long interval (e.g. `reconcile.Result{RequeueAfter: 30 * time.Minute}`).
    * *Pros:* Respects user intent to keep reconciling while protecting GCP API quota and eliminating log spam.
    * *Cons:* Introduces an arbitrary requeue interval and continues making low-frequency failing API calls indefinitely.
  * **Option C: Status-Driven Error Pause (Preserve User Annotation, Halt Current Generation)**
    * KCC **never** modifies the user's `actuation-mode: "Reconciling"` annotation.
    * KCC records the non-retryable failure in `status.conditions[Ready]` with `observedGeneration == generation` and halts retries for the *current* generation (`reconcile.Result{}, nil`).
    * *Re-triggering:* Modifying `.spec` automatically resumes reconciliation (`generation > observedGeneration`). If resolved out-of-band, re-applying or updating the annotation acts as a one-shot trigger to attempt one new reconciliation for the current generation.
    * *Pros:* Decouples user intent (annotations) from controller status (conditions), prevents annotation mutation, and stops quota burn.
    * *Cons:*
      * **Opaque & Misleading System State:** An operator inspecting `metadata.annotations` sees `actuation-mode: "Reconciling"` and naturally assumes the controller is actively retrying/reconciling against GCP, whereas KCC is actually silently halted.
      * **Violates User Intent Without Explicit Feedback:** It fails to honor the user's explicit command to reconcile, yet hides the fact that it stopped. This creates a confusing mental model where the resource's declared metadata contradicts the controller's runtime behavior.

---

## 7. Future Improvements

### 7.1. Metrics & Observability
We can introduce dedicated Prometheus metrics to track actuation pause behavior across the cluster:
* `kcc_reconciler_auto_paused_total{group, version, kind, namespace, reason}`: Counter incremented whenever a resource is automatically paused due to a non-retryable error.
* `kcc_reconciler_paused_resources{group, version, kind, namespace}`: Gauge tracking currently paused resources.

### 7.2. Automated Fan-Out Enqueue on Cluster/Namespace Unpause
To improve the UX of `ConfigConnectorContext` / `ConfigConnector` unpause, controllers could attach an event handler to `ConfigConnectorContext` that maps `PAUSED -> RECONCILING` transitions to automatically enqueue all managed resources in the namespace (`handler.EnqueueRequestsFromMapFunc`), eliminating the reliance on periodic re-reconciliation.

---

## 8. Appendix

### 8.1. Detailed Categorization of Non-Retryable vs. Retryable Errors

Classifying errors as non-retryable must balance **preventing infinite loops** against **avoiding premature halts on transient failures**. The rationale for each error category is detailed below:

#### A. Non-Retryable Error Categories
* **`codes.InvalidArgument` (gRPC 3) / HTTP 400 (with Invalid Argument Message):**
  * **Rationale:** Indicates client-side errors such as malformed syntax, invalid field combinations, unparseable parameters, or attempts to modify immutable fields. Retrying an identical request against the same API state is deterministically guaranteed to fail. Halting retries protects quota, while updating the `spec` will automatically resume reconciliation by incrementing `metadata.generation`.
* **`codes.Unimplemented` (gRPC 12) / HTTP 501 (Not Implemented):**
  * **Rationale:** Indicates that the requested API method or capability is unsupported by the backend service. Retrying cannot succeed until the server implements the feature or the user removes the unsupported configuration.

#### B. Ambiguous or Conditionally Retryable Categories (Kept as Retryable)
* **`codes.FailedPrecondition` (gRPC 9) / HTTP 412 / HTTP 400 (Precondition):**
  * **Rationale:** Often indicates that the system is not in a state required for the operation (e.g., a parent resource is still creating/deleting, or an asynchronous operation is in flight). Many precondition errors resolve automatically once external state converges. Therefore, treating them as retryable is safer to allow eventual consistency.
* **`codes.PermissionDenied` (gRPC 7) / HTTP 403 & `codes.Unauthenticated` (gRPC 16) / HTTP 401:**
  * **Rationale:** While permission errors look like client configuration issues, IAM bindings and service account roles are often granted out-of-band in GCP *without* updating the KRM `spec`. Continuing retries allows the controller to automatically converge once IAM permissions propagate.
* **`codes.NotFound` (gRPC 5) / HTTP 404:**
  * **Rationale:** Often caused by eventual consistency where a newly created dependency (e.g., a VPC network or project) has not yet propagated across GCP backend regions. Retrying allows eventual resolution.
* **`codes.ResourceExhausted` (gRPC 8) / HTTP 429 (Too Many Requests):**
  * **Rationale:** Rate limit and quota exhaustion are transient. The `controller-runtime` exponential backoff allows quota replenishment.
* **`codes.Unavailable` (gRPC 14) / `codes.DeadlineExceeded` (gRPC 4) / `codes.Internal` (gRPC 13) / HTTP 5xx:**
  * **Rationale:** Standard transient infrastructure or server-side issues that must continue retrying.

### 8.2. KCC's Standard Configuration Hierarchy

Config Connector uses a standardized multi-tier configuration hierarchy to allow platform administrators and application teams to configure custom resource (CR) behaviors with fine-grained scoping. 

#### A. The Multi-Tier Precedence Pattern
When determining runtime behavior for any given resource, KCC resolves settings in order of decreasing specificity:

```
┌─────────────────────────────────────────────────────────────────┐
│ 1. Resource-Level Annotation (Per-CR Override)                  │
│    Highest precedence. Explicitly configured on the CR.         │
└────────────────────────────────┬────────────────────────────────┘
                                 │ (if unset)
┌────────────────────────────────▼────────────────────────────────┐
│ 2. Namespace-Level Configuration (ConfigConnectorContext)       │
│    Scoped to an individual namespace (in namespaced mode).      │
└────────────────────────────────┬────────────────────────────────┘
                                 │ (if unset / in cluster mode)
┌────────────────────────────────▼────────────────────────────────┐
│ 3. Cluster-Level Configuration (ConfigConnector)                │
│    Global fallback across all namespaces in the cluster.        │
└────────────────────────────────┬────────────────────────────────┘
                                 │ (if unset)
┌────────────────────────────────▼────────────────────────────────┐
│ 4. Built-in System Default                                      │
└─────────────────────────────────────────────────────────────────┘
```

#### B. Application to Resource-Level Actuation Controls
The proposed `cnrm.cloud.google.com/actuation-mode` annotation follows this exact pattern:
* **Resource Level:** `cnrm.cloud.google.com/actuation-mode: "Paused" | "AutoPaused" | "Reconciling"` (allows per-resource pause or per-resource unpause override).
* **Namespace Level:** `ConfigConnectorContext.spec.actuationMode: "Paused" | "Reconciling"`.
* **Cluster Level:** `ConfigConnector.spec.actuationMode: "Paused" | "Reconciling"`.
* **Default:** `"Reconciling"`.

#### C. Case Study: `state-into-spec`
The `state-into-spec` option is an example of this hierarchy:

1. **Resource Level:** An individual resource can declare `cnrm.cloud.google.com/state-into-spec: "merge"` (or `"absent"`) in `metadata.annotations`. This always takes highest precedence for that specific resource.
2. **Namespace Level:** `ConfigConnectorContext.spec.stateIntoSpec: "Merge"` configures the default behavior for all resources within that namespace.
3. **Cluster Level:** `ConfigConnector.spec.stateIntoSpec: "Merge"` configures the default behavior across the entire cluster.
4. **Resolution Logic:**
   * If the resource annotation exists, use the annotation value.
   * Else if running in namespaced mode and `ConfigConnectorContext.spec.stateIntoSpec` is set, use `CCC.spec.stateIntoSpec`.
   * Else if `ConfigConnector.spec.stateIntoSpec` is set, use `CC.spec.stateIntoSpec`.
   * Else, fallback to default (`"absent"`).

#### D. Case Study: `project-id` (Legacy Pre-Operator Pattern)
The `cnrm.cloud.google.com/project-id` annotation demonstrates an early resolution hierarchy that predates the KCC Operator and the introduction of `ConfigConnector` (CC) and `ConfigConnectorContext` (CCC) CRDs. 

**Note:** This pattern is no longer recommended for new features unless a use case is specifically evaluated and justified. The modern pattern using `ConfigConnector` and `ConfigConnectorContext` CRDs is preferred:

1. **Resource Level:** An individual resource can declare `cnrm.cloud.google.com/project-id: "my-custom-project"` (or specify `.spec.projectRef`), explicitly targeting a GCP project.
2. **Namespace Level:** If unset on the resource, KCC inspects the standard Kubernetes `Namespace` object for a `cnrm.cloud.google.com/project-id` annotation.
3. **Default:** If unset on both the resource and the namespace, KCC defaults the GCP project ID to the Kubernetes namespace name (`namespace.metadata.name`).
4. **Legacy Pattern (`project-id`):** Relies on annotating native Kubernetes `Namespace` objects directly, which lacks OpenAPI schema validation, lacks CRD status reporting, and requires broad cluster-level permissions for app teams to configure namespace-level defaults.

### 8.3. Event Dropping & Unpause Latency Analysis

An important operational UX consideration in pause and unpause lifecycles is how Kubernetes events are handled when transitioning between `PAUSED` and `RECONCILING`.

#### A. Event Types: Apply vs. Delete
In Kubernetes controllers, reconciliation is triggered by two primary categories of Watch events:
1. **Apply Events (`Create` / `Update`):** A user creates a resource or updates its `.spec` in Kubernetes (or via GitOps).
2. **Delete Events (`Delete`):** A user initiates deletion (`kubectl delete` or namespace deletion), causing the Kubernetes API server to set `metadata.deletionTimestamp`.

#### B. Event Behavior in Cluster / Namespace Pause (`ConfigConnector` / `ConfigConnectorContext`)
When an entire cluster or namespace is paused via `spec.actuationMode: "Paused"` on `ConfigConnector` (CC) or `ConfigConnectorContext` (CCC):

1. **Apply Events while Paused:**
   * When a user creates or updates a resource, the API server emits an Apply Watch event to the reconciler.
   * The controller checks the actuation mode, sees `"Paused"`, and returns `reconcile.Result{}, nil` immediately without queueing further work. The event is consumed and dropped.

2. **Delete Events while Paused:**
   * When a user deletes a resource (`kubectl delete`), the API server sets `metadata.deletionTimestamp` and emits a Delete Watch event.
   * The controller observes `"Paused"`, ensures finalizers remain in place to protect the cloud provider resource, and returns `reconcile.Result{}, nil` without calling GCP Delete. The deletion event is dropped.
   * **Resulting State:** The Kubernetes object remains stuck in `Terminating` state, while the underlying GCP resource is fully preserved and continues running.

3. **UX After Unpaused:**
   * When an administrator transitions CC/CCC `spec.actuationMode` from `"Paused"` to `"Reconciling"`, no watch events are emitted for individual custom resources because the CRs themselves were not modified.
   * **Latency UX Gap:**
     * **Pending `.spec` updates** made during the pause window remain unactuated.
     * **Pending deletion requests** (objects in `Terminating` status) remain stuck without finalizing in GCP.
   * Neither operation will execute until either:
     * A user manually touches each CR (e.g. running `kubectl annotate` on the CR to trigger a watch event), or
     * Periodic re-reconciliation happens (which can take 10+ minutes, or never if `reconcile-interval-in-seconds: "0"`).

#### C. Event Behavior in Resource-Level Design (`cnrm.cloud.google.com/actuation-mode`)
The resource-level actuation mode design addresses event lifecycles on individual resources with immediate, deterministic feedback:

1. **`AutoPaused` (KCC-Initiated on Terminal Error):**
   * **1.1. Apply Events while AutoPaused:**
     * **When `.spec` is modified (`generation > observedGeneration`):** Updating the `.spec` (e.g. fixing a typo or invalid field) mutates the Custom Resource directly in Kubernetes. The API server emits an immediate Watch event. The controller detects that the spec was updated, automatically clears `actuation-mode: "AutoPaused"`, and reconciles the fix against GCP immediately without waiting for periodic re-reconciliation.
     * **When `.spec` is unchanged (`generation == observedGeneration`):** Metadata updates (or background watch triggers) that do not change `.spec` are evaluated and skipped (`reconcile.Result{}, nil`), halting retry loops.
   * **1.2. Delete Events while AutoPaused:**
     * When a user deletes a resource (`kubectl delete`) in `AutoPaused` mode, the controller observes `metadata.deletionTimestamp != nil`.
     * Because `AutoPaused` was triggered by a configuration error during Create/Update, the controller does not block deletion: it immediately proceeds with deletion and finalizer cleanup, allowing the failed resource to be deleted cleanly.

2. **`Paused` (User-Initiated Manual Pause):**
   * **2.1. Apply Events while Paused:**
     * **While Paused:** `.spec` updates trigger watch events that are ignored/dropped by the controller, supporting the same function but isolated to one single resource.
     * **On Unpause:** When the user unpauses the resource (setting `actuation-mode: "Reconciling"` or removing the annotation), that metadata update is written directly to the CR. The API server emits an immediate Watch event, prompting the controller to reconcile the pending `.spec` against GCP immediately without waiting for periodic re-reconciliation.
   * **2.2. Delete Events while Paused:**
     * **While Paused:** If a deletion is requested (`kubectl delete`), the deletion event is dropped. KCC leaves finalizers intact and halts deletion against GCP, preserving the cloud resource. The Kubernetes object remains in `Terminating` state.
     * **Resolving a Paused Deletion:**
       * *To actuate deletion in GCP:* The user unpauses the resource (setting `actuation-mode: "Reconciling"` or removing the annotation), triggering an immediate Watch event on the CR that actuates GCP deletion and strips finalizers.
       * *To abandon without touching GCP:* The user sets `cnrm.cloud.google.com/deletion-policy: "abandon"` and unpauses the resource, allowing KCC to remove the Kubernetes object without calling GCP Delete APIs.
