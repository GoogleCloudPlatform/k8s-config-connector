Multi-Cluster Leader Election

## 1. Overview

This library provides a robust, decentralized, and Kubernetes-native leader election mechanism that allows a single replica of a controller to be elected as a leader from a pool of candidates running across multiple Kubernetes clusters.

## 2. Core Principles & Tenets

This design adheres to the following non-negotiable principles:

*   **Client-Side Cloud Agnosticism:** Client controllers (e.g., KCC) that consume this library should not need to compile in any cloud-provider-specific SDKs or dependencies. The integration library should be cloud-neutral.
*   **Decentralized Control Plane:** The election logic will be managed by a controller running in each participating cluster, eliminating a central control plane as a single point of failure.
*   **Seamless Integration:** The solution must be easily consumable by any controller built with `client-go` and `controller-runtime`, leveraging the standard `resourcelock.Interface` for a native-feeling integration.

## 3. High-Level Architecture

The system consists of three primary components:

*   **Client-Side Library (`MultiClusterLeaseLock`):** A lightweight, cloud-agnostic Go package that client controllers import. It implements the standard `resourcelock.Interface`.
*   **Decentralized Election Controller:** A controller that runs in each participating cluster. It is the only component that interacts with the global lock backend.
*   **Global Lock Backend:** An external, highly-available storage system that supports atomic compare-and-swap operations (e.g., a dedicated etcd cluster, GCS, DynamoDB).

The interaction flow is as follows:

1.  A Client Controller replica uses the Client-Side Library to create/update a `MultiClusterLease` CR in its local cluster. This serves as its candidacy declaration and liveness heartbeat.
2.  The Decentralized Election Controller in that same cluster observes this local CR.
3.  The Election Controller then contends for a lock on the Global Lock Backend on behalf of its local candidate.
4.  Based on the outcome of the global contention, the Election Controller updates the status of the local `MultiClusterLease` CR.
5.  The Client Controller learns it has become the leader by observing the change in the status of its local CR.

## 4. API Contract: MultiClusterLease CRD

The `MultiClusterLease` CRD is the central API contract. It cleanly separates the concerns between the client candidate and the election controller.

*   **`spec`** (Written by Client Controller): Represents the desired state of a candidate.
    *   `holderIdentity` (string): The unique ID of the candidate pod.
    *   `leaseDurationSeconds` (int): The duration the lease is considered valid.
    *   `renewTime` (`metav1.MicroTime`): The timestamp of the last heartbeat from the candidate. This is the primary liveness signal.

*   **`status`** (Written by Election Controller): Represents the observed state of the global election. This field is protected by the `/status` subresource, making it read-only for the client controller.
    *   `leader` (string): The `holderIdentity` of the confirmed global leader.
    *   `acquireTime` (`metav1.MicroTime`): Timestamp of when the global lock was acquired.
    *   `conditions` (`metav1.Condition`): Standard Kubernetes conditions for detailed, machine-readable status.

## 5 Architectural Comparison: CRD-based vs. Direct Global Resource Lock

Two primary architectural patterns were considered for implementing the custom `resourcelock.Interface` for multi-cluster leader election.

#### 1. CRD + Controller (Recommended Pattern)

This is a Kubernetes-native pattern where client controllers interact with a local `MultiClusterLease` Custom Resource (CR). A separate, decentralized election controller in each cluster observes this CR and reconciles its state with a true global lock (e.g., GCS, etcd).

*   **Cloud Neutrality:** Excellent. This pattern provides perfect decoupling. The client controller has zero knowledge of the global lock mechanism (GCS, GitHub Gist, etc.). Its only dependency is on the MultiClusterLease CRD definition and a small, cloud-agnostic library that knows how to interact with it. The complex, cloud-specific logic is entirely encapsulated within the separate, decentralized election controller. 
*   **Ease of Adoption:** High. Provides an idiomatic `controller-runtime` integration with no unwanted dependency pollution for the consumer.
*   **Operational Complexity:** Medium. Requires deploying and managing a new election controller in each cluster.

#### 2. Direct Global Resource Lock (Alternative Pattern)

This is a simpler pattern where the custom `resourcelock` implementation imported by the client controller interacts *directly* with the global lock (e.g., a GitHub Gist, GCS object).

*   **Cloud Neutrality:** Poor. The client controller is directly and tightly coupled to the global lock's SDK (e.g., GitHub, GCS). This violates the primary project constraint.
*   **Ease of Adoption:** Medium. It uses the correct `resourcelock.Interface` but forces consumers to compile in cloud-specific dependencies they wish to avoid.
*   **Operational Complexity:** Low. There are no new controller components to manage.

### Conclusion

The **CRD + Controller** pattern is the architecturally superior solution. While it introduces the operational requirement of a new controller, it uniquely satisfies the project's most critical constraint: keeping client controllers completely decoupled from any cloud-specific implementation details. It provides a clean, robust, and idiomatic solution that aligns with modern Kubernetes best practices.

## 6. Design Details and Considerations

Note that design details is subject to change.

### Detailed Workflow for the Custom `resourcelock` (Asynchronous Model)

The recommended approach is a fully asynchronous, non-blocking implementation of `client-go`'s `resourcelock.Interface` that uses a custom resource (CR) to bridge the gap between the client controller and the election controller. This design is efficient, resilient, and aligns with Kubernetes' core architectural principles.

*   **`Update()` Method (Heartbeating):**
    *   When a client controller replica (leader or follower) needs to renew its lease, the `client-go` `LeaderElector` calls the `Update()` method on our custom lock.
    *   This method's sole responsibility is to perform a fast, non-blocking update to the `spec` of its local CR, refreshing a `renewTime` timestamp. This action serves as a "heartbeat," signaling to the election controller that the candidate is still alive. The call returns immediately without waiting for any global confirmation.

*   **`Create()` Method (Initial Candidacy & Race Condition Prevention):**
    *   When the very first candidate starts, `client-go`'s `LeaderElector` finds no existing lock object and calls the `Create()` method.
    *   This method performs a fast, non-blocking `Create` of the local CR, which registers the pod as a candidate in the global election.
    *   **Crucially, to prevent a race condition where the first creator incorrectly assumes leadership**, the method immediately returns an `AlreadyExists` error. This deliberately tricks the `client-go` library into thinking the acquisition failed, forcing it into its standard retry loop instead of prematurely granting leadership.

*   **`Get()` Method (Verifying Leadership):**
    *   After every `Create` or `Update` attempt, the `client-go` `LeaderElector` calls the `Get()` method to learn the authoritative state of the lock.
    *   This method performs a fast, non-blocking read of the **`status`** subresource of the local CR. This `status` is written exclusively by the election controller after it has successfully contended for the global lock.
    *   The client controller learns it has won the election only when its identity appears in the `status.leader` field returned by this `Get()` call.

This asynchronous flow cleanly separates the client's "desired state" (declaring candidacy in the `spec`) from the "observed state" (learning the authoritative election result from the `status`).

### Alternative Synchronous Implementation (Anti-Pattern)

An alternative, synchronous implementation would involve making the `Create()` and `Update()` methods block until the election controller confirms global leadership.

*   **How it Would Work:** In this model, after creating or updating the local CR's `spec`, the `Create()`/`Update()` method would enter a polling or watch loop. It would repeatedly check the CR's `status` field, waiting for the election controller to perform the global election and write the result. The method would only return a successful response to the `client-go` library after this confirmation is received.

*   **Why This is an Anti-Pattern:**
    *   **Violates Kubernetes Principles:** It breaks the asynchronous, decoupled nature of Kubernetes controllers. It creates a tight, synchronous dependency, forcing the client controller to wait for the entire reconciliation loop of another controller to complete.
    *   **Creates a Brittle System:** This tight coupling makes the system fragile. If the election controller is slow, has crashed, or is partitioned from the global lock, the client controller's startup will hang indefinitely. The asynchronous model is more resilient, as the client can continue its retry loop independently.
    *   **Interferes with `client-go`:** The `client-go` `LeaderElector` has its own internal timeouts and logic. A long-blocking call can interfere with this logic, leading to unpredictable behavior and failures.

### Why Not Directly Reuse the "LeaseCandidate" Feature

Kubernetes' new "Coordinated Leader Election" feature introduces the `LeaseCandidate` resource, which seems conceptually similar to our goal of announcing candidacy. While we should absolutely adopt its **architectural pattern**, reusing the `LeaseCandidate` object itself is not recommended for three main reasons:

1.  **Semantic Mismatch and Scope:** The `LeaseCandidate` resource was designed specifically to solve **in-cluster version skew** during control plane upgrades. Its API fields, like `binaryVersion` and `strategy`, are tailored for this purpose and are irrelevant to our multi-cluster availability goal. Using it would be a misuse of the API.

2.  **Conflict with Built-in Controller:** If the `CoordinatedLeaderElection` feature is enabled on a cluster, a built-in Kubernetes controller within the kube-apiserver actively watches and reconciles `LeaseCandidate` objects. Our custom election controller would compete with this native Kubernetes logic, leading to race conditions and unpredictable behavior.

3.  **Lack of Extensibility:** `LeaseCandidate` is a built-in API type that we cannot modify. A robust multi-cluster solution may require additional fields in spec and status.
