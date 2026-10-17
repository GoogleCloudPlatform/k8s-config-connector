# Simple ConfigConnector Test

This test case verifies the reconciliation of a `ConfigConnector` resource with default settings.

## Configuration
- **Mode:** Namespaced mode (default when `spec.mode` is empty).
- **Experiments:** None enabled.

## Expected Resources

The generated resources are similar to the [simple-cluster-mode](../simple-cluster-mode/README.md) case, without the `cnrm-controller-manager` StatefulSet:

- **StatefulSets:**
  - `cnrm-unmanaged-detector`: Indicates Config Connector is **not** properly enabled to manage resources. 
                               Because Config Connector is in namespaced-mode but the namespace does not have a ConfigConnectorContext object.
