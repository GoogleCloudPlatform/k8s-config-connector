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
