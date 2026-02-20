# Controller Configuration

Config Connector uses different underlying implementations (controller types) to manage Google Cloud resources. These include:

*   **Direct Controllers:** Newer implementations that call Google Cloud APIs directly. These are generally preferred for their simplicity and better integration.
*   **Terraform-based Controllers (Legacy):** Older implementations that wrap the Terraform Google provider.
*   **DCL-based Controllers (Legacy):** Older implementations that wrap the Google Cloud Declarative Library (DCL).

While Config Connector selects a default controller for each resource, you can override this behavior at the namespace level (using `ConfigConnectorContext`) for supported resources.

## Critical User Journeys

This guide helps you achieve the following tasks:

*   **[Verify which controller is managing my resource](#verifying-the-controller):** Learn how to check the active controller through logs or static configuration.
*   **[Identify supported controller types for a resource](#finding-supported-controller-types):** Discover which controller implementations (Direct, TF, DCL) are available for a given resource kind.
*   **[Override the default controller](#example):** Step-by-step example of using `ConfigConnectorContext` to switch a resource to a different controller type.

## Configuration Resources

There are two primary resources for configuring Config Connector:

1.  **ConfigConnector (CC):** A cluster-scoped resource that controls cluster-wide options, such as the operational mode (cluster vs. namespaced).
    *   **Note:** `ConfigConnector` does not currently support controller type overrides at the cluster level. In cluster mode, overrides can only be applied at the individual resource level using legacy annotations.
2.  **ConfigConnectorContext (CCC):** A namespace-scoped resource that controls options for a specific namespace.
    *   **Note:** Controller overrides via `ConfigConnectorContext` are only supported when Config Connector is running in **namespaced mode**.

### ConfigConnectorContext Reference

The `ConfigConnectorContext` resource allows you to configure identity, billing, and experimental features for all Config Connector resources within a specific namespace.

#### Identity and Billing
| Field | Type | Description |
| :--- | :--- | :--- |
| `googleServiceAccount` | `string` | **Required.** The Google Service Account to be used by Config Connector to authenticate with Google Cloud APIs in the associated namespace. |
| `requestProjectPolicy` | `enum` | Specifies which project to use for preconditions, quota, and billing. Supported values: `SERVICE_ACCOUNT_PROJECT` (default), `RESOURCE_PROJECT`, or `BILLING_PROJECT`. |
| `billingProject` | `string` | The project ID to use for billing when `requestProjectPolicy` is set to `BILLING_PROJECT`. |

#### Reconciliation and Actuation
| Field | Type | Description |
| :--- | :--- | :--- |
| `actuationMode` | `enum` | Controls how resources are actuated. Supported values: `Reconciling` (default), `Paused`. |
| `stateIntoSpec` | `enum` | Overrides the default `cnrm.cloud.google.com/state-into-spec` annotation. Supported values: `Absent`, `Merge`. |

#### Experimental and Advanced
| Field | Type | Description |
| :--- | :--- | :--- |
| `experiments.controllerOverrides` | `map[string]string` | Allows overriding the controller for specific resource types within the namespace. |
| `managerNamespace` | `string` | **Immutable.** Instructs Config Connector to deploy controller managers in a specific namespace instead of the standard `cnrm-system`. |
| `version` | `string` | Specifies the exact addon version to be deployed. |

## Controller Overrides

The `experiments.controllerOverrides` field in `ConfigConnectorContext` allows you to specify which controller to use for a given resource kind within the namespace, overriding the system default.

**Note:** Not all resources support all controller types. You can only override the controller with a type that is explicitly supported for that resource. See [Verifying the Controller](#verifying-the-controller) for how to check supported types.

The key for each entry must follow the format `Kind.group`, and the value must be one of the supported controller types: `direct`, `tf`, or `dcl`.

### Finding Resource Kind and Group

To find the Kind and Group for a resource, you can inspect its YAML definition or its CRD.

*   **Kind:** Found in the `kind` field of the resource (e.g., `BigQueryDataset`).
*   **Group:** Found in the `apiVersion` before the slash (e.g., `bigquery.cnrm.cloud.google.com`).

For example, for a `BigQueryDataset`, the override key would be `BigQueryDataset.bigquery.cnrm.cloud.google.com`.

### Example

The following `ConfigConnectorContext` configures Config Connector to use the experimental `direct` controller for `BigQueryDataset` resources:

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

## Precedence

Config Connector determines which controller to use following this order of precedence:

1.  **Resource Annotation (Legacy):** If the `alpha.cnrm.cloud.google.com/reconciler` annotation is present on the resource. Config Connector still honors this legacy behavior for backward compatibility, but it is moving away from it and you should use `ConfigConnectorContext` to override controllers.
2.  **ConfigConnectorContext Override:** If an entry for the resource's `Kind.group` exists in the namespace's `ConfigConnectorContext`.
3.  **Static Default:** The default controller implementation defined within the Config Connector version you are running.

## Verifying the Controller

To determine which controller is being used for a resource, you should check the following in order:

1.  **ConfigConnectorContext Overrides:** Check the `ConfigConnectorContext` in the resource's namespace for any `experiments.controllerOverrides`.
2.  **Static Configuration:** If no override is present, Config Connector uses the default defined in [pkg/controller/resourceconfig/static_config.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/resourceconfig/static_config.go).

### Finding Supported Controller Types

To find which controller types are available for a specific resource Kind, consult the [pkg/controller/resourceconfig/static_config.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/resourceconfig/static_config.go) file in the repository.

For each resource, the file defines:
*   **`DefaultController`:** The controller type used if no override is specified.
*   **`SupportedControllers`:** The list of all controller types that can be successfully used as an override for that resource.

**Example entry in `static_config.go`:**
```go
{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}: {
    DefaultController: k8s.ReconcilerTypeTerraform, 
    SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerTypeDirect, k8s.ReconcilerTypeTerraform}
},
```
In this example, `BigQueryDataset` defaults to `tf` but can be overridden to `direct`.

### Inspecting Logs

You can also verify the routing decision by inspecting the logs of the `cnrm-controller-manager` pod. Look for the message `"routing to controller"` which includes the selected controller type.

To view these logs (ensure you have appropriate verbosity enabled, typically `V(1)`):

```bash
kubectl -n cnrm-system logs pod/cnrm-controller-manager-0 | grep "routing to controller"
```

The log entry will look similar to this:
`Info  routing to controller {"resource": {"namespace": "my-namespace", "name": "my-resource"}, "type": "direct"}`
