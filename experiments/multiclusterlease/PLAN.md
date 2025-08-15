# Implementation Plan: Multi-Cluster Leader Election

This document outlines the steps to implement the asynchronous, multi-cluster leader election mechanism as described in `GEMINI.md`.

## Phase 1: Core API and Client Library

1.  **[ ] Modify the CRD `spec`**
    -   [ ] Add `holderIdentity` (string) to `MultiClusterLeaseSpec`.
    -   [ ] Add `renewTime` (metav1.MicroTime) to `MultiClusterLeaseSpec`.
    -   [ ] Run `make generate` to update the generated deepcopy functions.

2.  **[ ] Implement the Client-Side `resourcelock.Interface`**
    -   [ ] Create the file `pkg/client/resourcelock.go`.
    -   [ ] Implement the `Create()` method:
        -   It should create a `MultiClusterLease` CR with the client's `holderIdentity` and `renewTime` in the `spec`.
        -   It MUST immediately return an `AlreadyExists` error to prevent premature leadership claims.
    -   [ ] Implement the `Update()` method:
        -   It should perform a non-blocking patch to update the `.spec.renewTime` field, acting as a heartbeat.
    -   [ ] Implement the `Get()` method:
        -   It should read the `.status.leader` field to get the authoritative leader.

## Phase 2: Controller Refactoring

3.  **[ ] Refactor the Election Controller**
    -   [ ] Update the controller to read the list of local candidates from the `MultiClusterLease` CRs.
    -   [ ] The controller should use the `spec.holderIdentity` from the leading local candidate when contending for the global lock.
    -   [ ] The controller must check the `spec.renewTime` to ensure a candidate is still live before renewing the global lock on its behalf.
    -   [ ] The controller must write the confirmed global leader's identity into a new `status.leader` field.
