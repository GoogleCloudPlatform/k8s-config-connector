# Multi-Cluster Lease ConfigConnector Test

This test case verifies the reconciliation of a `ConfigConnector` resource when the `multiClusterLease` experiment is enabled.

## Configuration

- **Mode:** Cluster mode (`spec.mode: cluster`).
- **Experiments:** `multiClusterLease` is enabled with specific lease details.

## Expected Resources

The generated resources are similar to the [simple-cluster-mode](../simple-cluster-mode/README.md) case, with the following modifications to the `cnrm-controller-manager` StatefulSet:

- **Leader Election:** The manager is started with `--leader-election-type=multicluster`.
- **Annotations:** The pod template includes lease-related annotations:
  - `cnrm.cloud.google.com/lease-name`
  - `cnrm.cloud.google.com/lease-namespace`
  - `cnrm.cloud.google.com/lease-identity`
