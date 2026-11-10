# KRMSyncer Architecture and Context

`syncer` (KRMSyncer) is a subproject of `kube-etl` that implements a Kubernetes controller for unidirectional resource synchronization between clusters.

## Goals

*   **Unidirectional Sync**: Copy KRM resources from a local cluster to a remote destination cluster.
*   **Dynamic configuration**: Driven by `KRMSyncer` Custom Resources (CRs).
*   **Status & Spec**: Support syncing both spec and status (optional).

## Architecture

This project is a standard Kubernetes controller built using `controller-runtime` and `kubebuilder` conventions.

### Directory Structure

*   `api/`: Contains the Custom Resource Definition (CRD) types (e.g., `KRMSyncer`).
*   `controllers/`: Contains the reconciliation logic (`KRMSyncerReconciler`).
*   `config/`: Kustomize configuration for deploying the controller and CRDs.
    *   `config/templates/`: YAML templates (sample `KRMSyncer` CR), filled in by `deploy.sh`.
*   `deploy.sh`: Deploys the controller to a destination GKE cluster and configures it to pull from a source GKE cluster.
*   `lib.sh`: Names and helpers shared by `deploy.sh` and `cleanup.sh` (e.g. the per-destination GSA name).
*   `cleanup.sh`: Revokes the IAM grants `deploy.sh` gave the destination cluster's GSA on source projects that are no longer in use.
*   `main.go`: The entry point for the controller manager.

## Key Components

*   **KRMSyncerReconciler**: The core logic that watches `KRMSyncer` objects and orchestrates the copying of resources defined in the rules.
*   **Dynamic Watching**: The controller likely needs to dynamically watch resources defined in the `KRMSyncer` rules (check `controllers/` implementation for details).

## Development

*   **Build**: Uses `Makefile` for building the binary and docker image.
*   **Dependencies**: Manages its own dependencies via `go.mod`.
