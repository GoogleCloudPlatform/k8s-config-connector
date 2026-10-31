# Simple ConfigConnector Test

This test case verifies the reconciliation of a `ConfigConnector` resource with default settings.

## Configuration
- **Mode:** Cluster mode (`spec.mode: cluster`).
- **Experiments:** None enabled.

## Expected Resources

When reconciled, the operator is expected to create/manage the following cluster-wide components in the `cnrm-system` namespace:

- **StatefulSets:**
  - `cnrm-controller-manager`: The core controller responsible for managing GCP resources.
  - `cnrm-deletiondefender`: Prevents accidental deletion of GCP resources.
- **Deployments:**
  - `cnrm-resource-stats-recorder`: Collects and records resource statistics.
  - `cnrm-webhook-manager`: Manages admission webhooks for Config Connector resources.
- **RBAC:**
  - `ServiceAccount`: Identity for each component.
  - `ClusterRole` & `ClusterRoleBinding`: Cluster-wide permissions for managing CRDs, Namespaces, and GCP resources.
- **Scaling:**
  - `HorizontalPodAutoscaler`: Automatically scales the webhook manager.
- **Networking:**
  - `Service`: Provides stable endpoints for the manager, webhook, and recorder.
- **CRDs:**
  - All supported Config Connector Custom Resource Definitions.
