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

## 6. Design Details: The Hybrid Synchronous `resourcelock`

The core of the client-side library is a custom implementation of `client-go`'s `resourcelock.Interface`. To guarantee correctness and prevent a "split-brain" race condition, this implementation uses a **Hybrid Synchronous Model**.

The methods (`Create` and `Update`) first send a fast, asynchronous signal to the election controller and then perform a short, synchronous wait to receive confirmation of the global election outcome. This "Signal then Poll" pattern provides the best of both worlds: it is loosely coupled and efficient, but also completely safe.

### Why a Purely Asynchronous Model Fails

A naive, purely asynchronous implementation of the `Update()` method is flawed in two ways:

1.  **Returning Success (`nil`) Causes Split-Brain:** If `Update()` simply patched the `spec` of the `MultiClusterLease` CR and immediately returned `nil`, the `client-go` `LeaderElector` would instantly assume it had acquired the lease. This would allow two different pods in two different clusters to both believe they are the leader simultaneously, causing a catastrophic split-brain.

2.  **Always Returning an Error Breaks Renewals:** An alternative asynchronous approach is to have `Update()` patch the `spec` but then always return a synthetic error to prevent the split-brain. While this works for *acquiring* the lease, it breaks the "happy path" for *renewing* it. The `LeaderElector` interprets any error from `Update()` as a failed renewal, causing a healthy, active leader to believe it has lost the lease and abdicate its position. This would lead to constant, incorrect leader churn.

The **Hybrid Synchronous Model** is the only design that solves both problems. It prevents the split-brain on acquisition while still allowing a healthy leader to successfully renew its lease.

### Detailed Workflow

*   **`Create()` and `Update()` Methods (Signal then Poll):**
    *   When a client calls `Create()` or `Update()`, the method first performs a fast, non-blocking `CREATE` or `PATCH` to the `.spec` of its local `MultiClusterLease` CR. This serves as the "Signal"—an immediate "I'm alive" heartbeat and candidacy declaration for the election controller to see.
    *   **Crucially, the method does not return yet.** It then enters the "Poll" phase: a short, blocking loop with a timeout (e.g., the `LeaseDuration`).
    *   Inside the loop, it repeatedly `GET`s the CR and reads the authoritative leader from the **`status.globalHolderIdentity`** field, also checking that the `status.observedGeneration` matches the `metadata.generation` of its earlier write.
    *   The method only returns a success (`nil`) if and when the identity in the `status` matches its own and the generation confirms the status is not stale. This guarantees the client pod only considers itself the leader after the backend election controller has authoritatively confirmed it.
    *   If the timeout is reached, an error is returned, correctly signaling to the `LeaderElector` that the acquisition/renewal attempt failed.

*   **`Get()` Method (Reading the Authoritative State):**
    *   This method remains a simple, non-blocking `GET` of the `MultiClusterLease` CR.
    *   It translates the fields from the CR's `spec` and `status` into the `LeaderElectionRecord` that the `client-go` `LeaderElector` expects. This is how the `LeaderElector` learns the authoritative state of the global lock.

### Why Not Directly Reuse the "LeaseCandidate" Feature

Kubernetes' new "Coordinated Leader Election" feature introduces the `LeaseCandidate` resource, which seems conceptually similar to our goal of announcing candidacy. While we should absolutely adopt its **architectural pattern**, reusing the `LeaseCandidate` object itself is not recommended for three main reasons:

1.  **Semantic Mismatch and Scope:** The `LeaseCandidate` resource was designed specifically to solve **in-cluster version skew** during control plane upgrades. Its API fields, like `binaryVersion` and `strategy`, are tailored for this purpose and are irrelevant to our multi-cluster availability goal. Using it would be a misuse of the API.

2.  **Conflict with Built-in Controller:** If the `CoordinatedLeaderElection` feature is enabled on a cluster, a built-in Kubernetes controller within the kube-apiserver actively watches and reconciles `LeaseCandidate` objects. Our custom election controller would compete with this native Kubernetes logic, leading to race conditions and unpredictable behavior.

3.  **Lack of Extensibility:** `LeaseCandidate` is a built-in API type that we cannot modify. A robust multi-cluster solution may require additional fields in spec and status.
