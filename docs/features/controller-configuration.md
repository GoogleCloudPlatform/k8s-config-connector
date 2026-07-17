# Controller Implementation Overrides

Config Connector allows you to explicitly choose which underlying controller implementation (Direct, Terraform, or DCL) to use for a specific resource kind. This is primarily used when transitioning resources from legacy implementations to newer "Direct" controllers.

## Controller Types

Config Connector uses different underlying implementations (controller types) to manage Google Cloud resources:

*   **Direct Controllers:** Newer implementations that call Google Cloud APIs directly using `controller-runtime`. These are preferred for simplicity and performance.
*   **Terraform-based Controllers (TF):** Older implementations that wrap the Terraform Google provider.
*   **DCL-based Controllers:** Older implementations that wrap the Google Cloud Declarative Library (DCL).

## Overriding the Controller Type

You can override the default implementation for a specific resource kind within a namespace using the `experiments.controllerOverrides` field in `ConfigConnectorContext`.

**Note:** You can only override to a controller type that is explicitly supported for that resource. See [Finding Supported Controller Types](#finding-supported-controller-types).

### Example

The following `ConfigConnectorContext` configures Config Connector to use the `direct` controller for `BigQueryDataset` resources in the `my-namespace` namespace:

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

1.  **Legacy Resource Annotation (deprecated):** The `alpha.cnrm.cloud.google.com/reconciler` or `cnrm.cloud.google.com/reconciler` annotation on the individual resource.
2.  **ConfigConnectorContext Override:** The `experiments.controllerOverrides` entry in the namespace's `ConfigConnectorContext`.
3.  **Static Default:** The default implementation defined in the Config Connector version.

For more information on legacy behaviors and annotations, see [Legacy Behavior](legacy-behavior.md).

---

## Verifying the Controller Type

To determine which controller type is being used for a resource:

1.  **Check Overrides:** Inspect the `ConfigConnectorContext` in the resource's namespace.
2.  **Check Static Config:** Consult the [static_config.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/resourceconfig/static_config.go) file in the repository.
3.  **Inspect Logs:** View the logs of the `cnrm-controller-manager` pod.

### Finding Supported Controller Types

Consult the [static_config.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/resourceconfig/static_config.go) file in the repository. For each resource kind, the file defines:
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

## Further Reading

*   [Config Connector Operational Modes](config-connector-mode.md) for general configuration options (Identity, Billing, etc.).
*   [Legacy Behavior](legacy-behavior.md) for details on legacy annotations.
*   [Pausing Reconciliation](pause.md) for details on the `actuationMode` field.
