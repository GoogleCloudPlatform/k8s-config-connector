# Controller Configuration

Config Connector can be configured through two primary resources to manage how it interacts with Google Cloud and how it reconciles your resources.

*   **ConfigConnector (CC):** A cluster-scoped resource that defines cluster-wide settings and the operational mode.
*   **ConfigConnectorContext (CCC):** A namespace-scoped resource that defines settings for a specific namespace, such as identity and billing.

## Critical User Journeys

This guide helps you achieve the following tasks:

*   **[Verify which controller type is managing my resource](#verifying-the-controller-type):** Learn how to check the active controller through logs or static configuration.
*   **[Identify supported controller types for a resource](#finding-supported-controller-types)**: Discover which controller implementations (Direct, TF, DCL) are available for a given resource kind.
*   **[Override the default controller implementation](#controller-implementation-overrides)**: Step-by-step example of using `ConfigConnectorContext` to switch a resource to a different controller type.
*   **[Configure the Google Service Account for a namespace](#identity-and-authentication)**: Set up the identity that Config Connector uses to act on your behalf.
*   **[Configure billing and quota project policies](#billing-and-quota)**: Control which Google Cloud project is charged for API requests.
*   **[Pause reconciliation of resources](#pausing-reconciliation)**: Temporarily stop Config Connector from making changes to Google Cloud resources.

---

## Configuration Resources

### ConfigConnector (Cluster-wide)

The `ConfigConnector` resource is used to configure the entire Config Connector installation.

| Field | Type | Description |
| :--- | :--- | :--- |
| `mode` | `enum` | **Required.** The operational mode: `namespaced` (default) or `cluster`. |
| `googleServiceAccount` | `string` | The Google Service Account (GSA) used in `cluster` mode with Workload Identity. |
| `credentialSecretName` | `string` | The Kubernetes secret containing a GSA key (used in `cluster` mode without Workload Identity). |
| `actuationMode` | `enum` | Cluster-wide default actuation mode (`Reconciling` or `Paused`). Overridden by `ConfigConnectorContext` in namespaced mode. |
| `stateIntoSpec` | `enum` | Cluster-wide default for the `state-into-spec` behavior (`Absent` or `Merge`). Overridden by `ConfigConnectorContext` in namespaced mode. |
| `experiments.multiClusterLease.leaseName` | `string` | The name of the `MultiClusterLease` object that Config Connector will create. |
| `experiments.multiClusterLease.namespace` | `string` | The namespace where the `MultiClusterLease` object will be created. |
| `experiments.multiClusterLease.clusterCandidateIdentity` | `string` | **Required.** Unique identity for the cluster in multi-cluster leader election. |

**Note:** `ConfigConnector` does not support controller implementation overrides at the cluster level. In cluster mode, overrides can only be applied at the individual resource level using legacy annotations.

### ConfigConnectorContext (Namespace-specific)

In `namespaced` mode, you must create a `ConfigConnectorContext` in each namespace where you want to use Config Connector.

| Category | Field | Description |
| :--- | :--- | :--- |
| **Identity** | `googleServiceAccount` | **Required.** The GSA used for this namespace. |
| **Billing** | `requestProjectPolicy` | Specifies which project to bill: `SERVICE_ACCOUNT_PROJECT` (default), `RESOURCE_PROJECT`, or `BILLING_PROJECT`. |
| | `billingProject` | The project ID to use when `requestProjectPolicy` is `BILLING_PROJECT`. |
| **Actuation** | `actuationMode` | Namespace-specific actuation mode (`Reconciling` or `Paused`). Takes precedence over `ConfigConnector` setting. |
| | `stateIntoSpec` | Namespace-specific `state-into-spec` behavior (`Absent` or `Merge`). Takes precedence over `ConfigConnector` setting. |
| **Advanced** | `experiments.controllerOverrides` | Map of `Kind.group` to controller type (`direct`, `tf`, `dcl`). |
| | `managerNamespace` | **Immutable.** Custom namespace for controller managers. |
| | `version` | Specifies the exact addon version to be deployed. |

---

## Controller Implementation Selection

Config Connector uses different underlying implementations (controller types) to manage Google Cloud resources:

*   **Direct Controllers:** Newer implementations that call Google Cloud APIs directly. These are generally preferred for simplicity and performance.
*   **Terraform-based Controllers (TF):** Older implementations that wrap the Terraform Google provider.
*   **DCL-based Controllers:** Older implementations that wrap the Google Cloud Declarative Library (DCL).

### Controller Implementation Overrides

You can override the default implementation for a specific resource kind within a namespace using the `experiments.controllerOverrides` field in `ConfigConnectorContext`.

**Note:** You can only override to a controller type that is explicitly supported for that resource. See [Finding Supported Controller Types](#finding-supported-controller-types).

#### Example

The following `ConfigConnectorContext` configures Config Connector to use the `direct` controller for `BigQueryDataset` resources:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: my-namespace
spec:
  googleServiceAccount: "kcc-sa@my-project.iam.gserviceaccount.com"
  experiments:
    controllerOverrides:
      BigQueryDataset.bigquery.cnrm.cloud.google.com: direct
```

### Precedence

Config Connector determines which controller type to use following this order of precedence:

1.  **Resource Annotation (Legacy):** The `alpha.cnrm.cloud.google.com/reconciler` annotation on the resource.
    *   *Note: This is legacy behavior. Use `ConfigConnectorContext` overrides for new configurations.*
2.  **ConfigConnectorContext Override:** The `experiments.controllerOverrides` entry in the namespace.
3.  **Static Default:** The default implementation defined in the Config Connector version.

---

## Verifying the Controller Type

To determine which controller type is being used for a resource:

1.  **Check Overrides:** Inspect the `ConfigConnectorContext` in the resource's namespace.
2.  **Check Static Config:** Consult [pkg/controller/resourceconfig/static_config.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/resourceconfig/static_config.go).
3.  **Inspect Logs:** View the logs of the `cnrm-controller-manager` pod.

### Finding Supported Controller Types

Consult the [static_config.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/resourceconfig/static_config.go) file in the repository.

For each resource kind, the file defines:
*   **`DefaultController`:** The controller type used if no override is specified.
*   **`SupportedControllers`:** The list of types that can be used as an override.

**Example entry:**
```go
{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}: {
    DefaultController: k8s.ReconcilerTypeTerraform, 
    SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerTypeDirect, k8s.ReconcilerTypeTerraform}
},
```

### Inspecting Logs

Look for the `"routing to controller"` message in the `cnrm-controller-manager` logs (requires `V(1)` verbosity).

```bash
kubectl -n cnrm-system logs pod/cnrm-controller-manager-0 | grep "routing to controller"
```

**Example Log Output:**
`Info  routing to controller {"resource": {"namespace": "my-namespace", "name": "my-resource"}, "type": "direct"}`

---

## Identity and Authentication

In `namespaced` mode, you must specify a `googleServiceAccount` in each `ConfigConnectorContext`. This GSA is used to authenticate all requests for resources in that namespace.

## Billing and Quota

Use `requestProjectPolicy` to control which project is billed for API calls:
*   `SERVICE_ACCOUNT_PROJECT`: Bill the project where the GSA resides.
*   `RESOURCE_PROJECT`: Bill the project where the resource is located.
*   `BILLING_PROJECT`: Bill a specific project defined in `billingProject`.

## Pausing Reconciliation

Set `actuationMode: Paused` in `ConfigConnector` (cluster-wide) or `ConfigConnectorContext` (namespace-specific) to stop Config Connector from making changes to Google Cloud. In this mode, Config Connector continues to monitor Kubernetes resources but does not actuate them to the cloud.
