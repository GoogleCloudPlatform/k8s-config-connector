# ConfigConnector and ConfigConnectorContext Reference

Config Connector is configured using two primary Kubernetes resources: `ConfigConnector` and `ConfigConnectorContext`.

*   **ConfigConnector (CC):** A cluster-scoped resource that defines cluster-wide settings and the operational mode.
*   **ConfigConnectorContext (CCC):** A namespace-scoped resource that defines settings for a specific namespace, such as identity and billing.

## ConfigConnector (Cluster-wide)

The `ConfigConnector` resource configures the entire Config Connector installation.

| Field | Type | Description |
| :--- | :--- | :--- |
| `mode` | `enum` | The operational mode: `namespaced` (default) or `cluster`. |
| `googleServiceAccount` | `string` | The Google Service Account (GSA) used in `cluster` mode with Workload Identity. Cannot be used with `credentialSecretName`. |
| `credentialSecretName` | `string` | The Kubernetes secret containing a GSA key (used in `cluster` mode without Workload Identity). Cannot be used with `googleServiceAccount`. |
| `actuationMode` | `enum` | Cluster-wide default actuation mode (`Reconciling` or `Paused`). Overridden by `ConfigConnectorContext` in namespaced mode. See [Pausing Reconciliation](pause.md). |
| `stateIntoSpec` | `enum` | Cluster-wide default for the `state-into-spec` behavior (`Absent` or `Merge`). Overridden by `ConfigConnectorContext` in namespaced mode. |
| `experiments.multiClusterLease.leaseName` | `string` | The name of the `MultiClusterLease` object that Config Connector will create. |
| `experiments.multiClusterLease.namespace` | `string` | The namespace where the `MultiClusterLease` object will be created. |
| `experiments.multiClusterLease.clusterCandidateIdentity` | `string` | **Required.** Unique identity for the cluster in multi-cluster leader election. |

## ConfigConnectorContext (Namespace-specific)

In `namespaced` mode, you must create a `ConfigConnectorContext` in each namespace where you want to use Config Connector.

| Category | Field | Description |
| :--- | :--- | :--- |
| **Identity** | `googleServiceAccount` | **Required.** The GSA used for this namespace. |
| **Billing** | `requestProjectPolicy` | Specifies which project to bill: `SERVICE_ACCOUNT_PROJECT` (default), `RESOURCE_PROJECT`, or `BILLING_PROJECT`. |
| | `billingProject` | The project ID to use when `requestProjectPolicy` is `BILLING_PROJECT`. |
| **Actuation** | `actuationMode` | Namespace-specific actuation mode (`Reconciling` or `Paused`). Takes precedence over `ConfigConnector` setting. See [Pausing Reconciliation](pause.md). |
| | `stateIntoSpec` | Namespace-specific `state-into-spec` behavior (`Absent` or `Merge`). Takes precedence over `ConfigConnector` setting. |
| **Advanced** | `experiments.controllerOverrides` | Map of `Kind.group` to controller type (`direct`, `tf`, `dcl`). See [Controller Implementation Overrides](controller-configuration.md). |
| | `managerNamespace` | **Immutable.** Custom namespace for controller managers. |
| | `version` | Specifies the exact addon version to be deployed. |

## Identity and Authentication

In `namespaced` mode, you must specify a `googleServiceAccount` in each `ConfigConnectorContext`. This GSA is used to authenticate all requests for resources in that namespace. Ensure that the GSA has the necessary IAM permissions to manage the GCP resources you define in that namespace.

## Billing and Quota

Use `requestProjectPolicy` to control which project is billed for API calls and whose quota is consumed:
*   `SERVICE_ACCOUNT_PROJECT`: Bill the project where the GSA resides.
*   `RESOURCE_PROJECT`: Bill the project where the resource is located.
*   `BILLING_PROJECT`: Bill a specific project defined in `billingProject`.
