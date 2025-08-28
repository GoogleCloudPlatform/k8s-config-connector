# Implementation Plan: Multi-Cluster Leader Election

This document outlines the steps to implement the asynchronous, multi-cluster leader election mechanism as described in `GEMINI.md`.

## Phase 1: Core API and Client Library

1.  **[x] Modify the CRD `spec`**
    -   [x] Add `holderIdentity` (string) to `MultiClusterLeaseSpec`.
    -   [x] Add `renewTime` (metav1.MicroTime) to `MultiClusterLeaseSpec`.
    -   [x] Run `make generate` to update the generated deepcopy functions.

2.  **[x] Implement the Client-Side `resourcelock.Interface`**
    -   [x] Create the file `pkg/client/resourcelock.go`.
    -   [x] Implement the `Create()` method with "Signal then Poll" logic.
    -   [x] Implement the `Update()` method with "Signal then Poll" logic.
    -   [x] Implement the `Get()` method.

## Phase 2: Controller Refactoring

3.  **[x] Refactor the Election Controller**
    -   [x] Update the controller to read the list of local candidates from the `MultiClusterLease` CRs.
    -   [x] The controller should use the `spec.holderIdentity` from the leading local candidate when contending for the global lock.
    -   [x] The controller must check the `spec.renewTime` to ensure a candidate is still live before renewing the global lock on its behalf.
    -   [x] The controller must write the confirmed global leader's identity into the `status.globalHolderIdentity` field.

## Phase 3: End-to-End Testing

4.  **[x] Create a `kind`-based E2E Test**
    -   [x] Create a new test file, e.g., `controllers/e2e_test.go`.
    -   [x] Write a Go test that:
        -   Connects to a Kubernetes cluster (expecting a `kind` cluster context).
        -   Creates a test namespace.
        -   Creates a `LeaderElector` instance configured with our custom `MultiClusterLeaseLock`.
        -   Runs the `LeaderElector` in a background goroutine.
        -   Asserts that the `onStartedLeading` callback is fired within a reasonable timeout.
        -   Asserts that the `MultiClusterLease` CR's `status.globalHolderIdentity` is correctly updated with the test's identity.
        -   Waits for a renewal period (e.g., `retryPeriod + 1s`).
        -   Asserts that the `status.globalRenewTime` on the CR has been updated, confirming the renewal "happy path".
    -   [x] Add a new `Makefile` target (e.g., `test-e2e`) that:
        -   Builds the controller Docker image.
        -   Creates a `kind` cluster.
        -   Loads the controller image into the `kind` cluster.
        -   Deploys the controller and its CRD.
        -   Runs the Go E2E test.
        -   Tears down the `kind` cluster.

## Phase 4: Refactor for Pluggable Storage Backend

5.  **[ ] Define the `Storage` Interface**
    -   [ ] Create a new directory: `pkg/storage`.
    -   [ ] Create a new file: `pkg/storage/interface.go`.
    -   [ ] In the new file, define the `LeaseData`, `LeaseObject`, and `Storage` interface with methods for `ReadLease`, `UpdateLease`, `CreateLease`, and `DeleteLease`.

6.  **[ ] Create the `GCSStorage` Implementation**
    -   [ ] Create a new file: `pkg/storage/gcs.go`.
    -   [ ] Add the `GCSStorage` struct and its methods to this file, implementing the `Storage` interface. This will contain all GCS-specific logic.

7.  **[ ] Refactor `LeaderElector` to Use the `Storage` Interface**
    -   [ ] In `pkg/leaderelection/leaderelection.go`, modify the `LeaderElector` struct to hold a `storage.Storage` interface instead of a `*storage.Client`.
    -   [ ] Update the `NewLeaderElector` constructor to accept a `storage.Storage` object.
    -   [ ] Rewrite the `AcquireOrRenew` method to use the `storage.Storage` interface. This method will now contain the core Read-Modify-Write and retry logic, making it the "brains" of the operation.
    -   [ ] Remove the low-level GCS helper methods (`readLease`, `updateLease`, etc.) from the `LeaderElector`, as that logic now lives in `GCSStorage`.

8.  **[ ] Wire Up the New Components**
    -   [ ] In `cmd/manager/main.go`, create an instance of `storage.GCSStorage`.
    -   [ ] In `controllers/multiclusterlease_controller.go`, update the `NewLeaderElector` constructor call to pass in the newly created `GCSStorage` instance.