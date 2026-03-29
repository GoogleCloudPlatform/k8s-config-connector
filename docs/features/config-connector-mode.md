# Config Connector Operational Modes

Config Connector supports two primary modes of operation: **Cluster Mode** and **Namespace Mode**. The operational mode determines how controllers are deployed and how identity is managed across your Kubernetes cluster.

## Cluster Mode

In **Cluster Mode**, a single instance of the Config Connector controller manages resources across all namespaces in your cluster.

*   **Unified Management:** One controller process watches for Config Connector resources in all namespaces.
*   **Centralized Identity:** It typically runs using a single Google Service Account (GSA) for the entire cluster. This GSA must have the necessary IAM permissions to manage all Google Cloud resources specified across all namespaces.
*   **Setup:** Configured by setting `spec.mode: cluster` in the cluster-scoped `ConfigConnector` resource.

## Namespace Mode

In **Namespace Mode**, Config Connector allows for granular control and isolation by running separate controller instances for each enabled namespace. This is the recommended mode for multi-tenant clusters.

*   **Isolation:** Each namespace can be configured independently, and controller processes are isolated.
*   **Flexible Identity:** Each namespace can have its own Google Service Account, allowing you to follow the principle of least privilege for different teams or projects.
*   **Enablement:** 
    1. Set `spec.mode: namespaced` in the cluster-scoped `ConfigConnector` resource.
    2. Create a `ConfigConnectorContext` resource in each namespace where you want to use Config Connector. This "enables" the namespace and defines its specific configuration (Identity, Billing, etc.).

---

# Configuration Reference

Config Connector is configured using two primary Kubernetes resources.

## ConfigConnector (Cluster-wide)

The `ConfigConnector` resource defines the cluster-wide operational mode and default settings.

| Field | Type | Description |
| :--- | :--- | :--- |
| `mode` | `enum` | The operational mode: `namespaced` (default) or `cluster`. |
| `googleServiceAccount` | `string` | The Google Service Account (GSA) used in `cluster` mode with Workload Identity. |
| `credentialSecretName` | `string` | The Kubernetes secret containing a GSA key (used in `cluster` mode without Workload Identity). |
| `actuationMode` | `enum` | Cluster-wide default actuation mode (`Reconciling` or `Paused`). Overridden by `ConfigConnectorContext` in namespaced mode. See [Pausing Reconciliation](pause.md). |
| `stateIntoSpec` | `enum` | Cluster-wide default for the `state-into-spec` behavior (`Absent` or `Merge`). Overridden by `ConfigConnectorContext` in namespaced mode. |
| `experiments.multiClusterLease.leaseName` | `string` | The name of the `MultiClusterLease` object that Config Connector will create. |
| `experiments.multiClusterLease.namespace` | `string` | The namespace where the `MultiClusterLease` object will be created. |
| `experiments.multiClusterLease.clusterCandidateIdentity` | `string` | **Required.** Unique identity for the cluster in multi-cluster leader election. |

## ConfigConnectorContext (Namespace-specific)

In `namespaced` mode, the `ConfigConnectorContext` (CCC) resource defines settings for a specific namespace.

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

---

# Identity and Billing

## Identity and Authentication

In `namespaced` mode, you must specify a `googleServiceAccount` in each `ConfigConnectorContext`. This GSA is used to authenticate all requests for resources in that namespace. Ensure that the GSA has the necessary IAM permissions to manage the GCP resources defined in that namespace.

## Billing and Quota

Use `requestProjectPolicy` to control which project is billed for API calls and whose quota is consumed:
*   `SERVICE_ACCOUNT_PROJECT`: Bill the project where the GSA resides.
*   `RESOURCE_PROJECT`: Bill the project where the Google Cloud resource is located.
*   `BILLING_PROJECT`: Bill a specific project defined in `billingProject`.
