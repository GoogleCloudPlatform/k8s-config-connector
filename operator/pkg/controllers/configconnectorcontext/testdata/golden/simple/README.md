# Simple ConfigConnectorContext Test

This test case verifies the reconciliation of a `ConfigConnectorContext` resource in a namespace when the cluster is in Namespaced mode.

## Configuration

- **Mode:** Namespaced mode (default when `spec.mode` is empty).
- **Namespace:** `testns1`.
- **Experiments:** None enabled.

## Expected Resources

When a `ConfigConnectorContext` is created in `testns1`, the operator is expected to spawn a namespaced controller instance in the `cnrm-system` namespace:

- **StatefulSet:** `cnrm-controller-manager-<namespace-id>`: A dedicated controller manager instance that only watches for resources in `testns1`.
- **ServiceAccount:** `cnrm-controller-manager-testns1`: The identity used by the namespaced manager.
- **RBAC:**
  - `ClusterRoleBinding`: Grants the namespaced manager permissions at the cluster level (e.g., to watch CRDs).
  - `RoleBinding` (in `testns1`): Grants the namespaced manager permissions to manage resources within the `testns1` namespace.
- **Service:** `cnrm-manager-<namespace-id>`: Endpoint for the namespaced manager's metrics and probes.
