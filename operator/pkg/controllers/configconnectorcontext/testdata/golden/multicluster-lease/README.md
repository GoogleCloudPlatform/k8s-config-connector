# Multi-Cluster Lease ConfigConnectorContext Test

This test case verifies the reconciliation of a `ConfigConnectorContext` when the global `multiClusterLease` experiment is enabled.

## Configuration

- **Mode:** Namespaced mode (default when `spec.mode` is empty).
- **Namespace:** `testns1`.
- **Experiments:** `multiClusterLease` is enabled in the global `ConfigConnector` resource.

## Expected Resources

The generated resources are similar to the [simple](../simple/README.md) case, but the namespaced manager StatefulSet is configured for multi-cluster leader election:

- **Leader Election:** The manager is started with `--leader-election-type=multicluster`.
- **Annotations:** The pod template includes lease-related annotations from the global `ConfigConnector` configuration:
  - `cnrm.cloud.google.com/lease-name`
  - `cnrm.cloud.google.com/lease-namespace`
  - `cnrm.cloud.google.com/lease-identity`
