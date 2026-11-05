# Pause KCC

KCC can be configured to pause actuation of resources on the cloud provider
(GCP). K8s objects continue to be reconciled with the api server but any
interaction with the cloud provider should be paused. This can be helpful for
debugging purposes or to have a hot standby.

## Pausing

Config Connector supports configuring actuation mode (`Reconciling` or `Paused`, with `Reconciling` as the default) at three levels:

1. **Global (Cluster) level**: via `spec.actuationMode` on the `ConfigConnector` resource.
2. **Namespace level**: via `spec.actuationMode` on the `ConfigConnectorContext` resource (when running in `namespaced` mode).
3. **Resource level (Direct controllers only)**: via the `cnrm.cloud.google.com/actuation-mode` annotation on an individual Config Connector resource.

### Precedence Order

When determining whether to actuate a resource on Google Cloud, Config Connector evaluates the configuration in the following order of precedence (highest to lowest):

1. **Resource annotation** (`cnrm.cloud.google.com/actuation-mode`, supported for Direct controllers only)
2. **`ConfigConnectorContext` `spec.actuationMode`** (only when `ConfigConnector` `spec.mode` is `namespaced`)
3. **`ConfigConnector` `spec.actuationMode`**
4. **Default**: `Reconciling`

### Pausing Globally

To pause KCC across namespaces it is sufficient to set the `ConfigConnector`'s `spec.actuationMode: Paused`. This will work when KCC runs in `cluster` and `namespaced` mode. To eventually resume actuation just set the field back to `Reconciling`.

### Pausing Per Namespace

When KCC is running in `namespaced` mode (and only then), operators can set
`spec.actuationMode: Paused` on the `ConfigConnectorContext` resource. To eventually
resume actuation for that namespace set the field back to `Reconciling`. The `actuationMode`
value on the `ConfigConnectorContext` takes precedence over the value in `ConfigConnector`.

### Reconciling Per Namespace (only)

It can be handy to have KCC globally paused but reconciling on a per
namespace level. To do this make sure KCC is running in `namespaced` mode
and the `spec.actuationMode: Paused` on the `ConfigConnector` resource. Then
reconciling can be turned on for a namespace by changing the `spec.actuationMode`
field for the `ConfigConnectorContext` to `Reconciling` for that namespace.
NOTE: you can avoid any pausing in actuation by first changing the
`ConfigConnectorContext` `actuationMode`.

### Pausing or Reconciling Individual Resources (Direct Controllers Only)

> [!NOTE]
> **Scope**: The `cnrm.cloud.google.com/actuation-mode` annotation is currently supported **exclusively for resources managed by Direct controllers**.

You can pause or resume actuation for an individual resource by setting the `cnrm.cloud.google.com/actuation-mode` annotation on the resource's metadata:

| Annotation Value | Behavior |
| :--- | :--- |
| **`"Paused"`** | Skips all actuation against Google Cloud for this resource, including **Create, Update, periodic Drift correction, and Delete** operations. |
| **`"Reconciling"`** | Enables actuation for this resource, overriding `spec.actuationMode: Paused` on the `ConfigConnectorContext` or `ConfigConnector`. |
| **Unset or `""`** | Falls back to the `ConfigConnectorContext` or `ConfigConnector` `spec.actuationMode` (defaulting to `Reconciling`). |
| **Any other value** | Treated as an invalid configuration; reconciliation fails and records `Ready: False` (`UpdateFailed` or `DeleteFailed`) in `.status.conditions`. |

#### Example 1: Pause a Specific Resource

```yaml
apiVersion: logging.cnrm.cloud.google.com/v1beta1
kind: LoggingLogMetric
metadata:
  name: logginglogmetric-sample
  annotations:
    cnrm.cloud.google.com/actuation-mode: "Paused"
spec:
  filter: "resource.type=gae_app AND severity>=ERROR"
  projectRef:
    external: "projects/my-project-id"
```

When `cnrm.cloud.google.com/actuation-mode: "Paused"` is set:
* **Create / Update / Drift**: Config Connector ensures the controller finalizer (`cnrm.cloud.google.com/finalizer`) is present on the Kubernetes object and skips making any create, update, or drift-correction requests to Google Cloud.
* **Delete**: If the Kubernetes object is deleted (`kubectl delete`), Config Connector **halts the deletion operation** against Google Cloud and retains the finalizer. The Kubernetes object remains in the `Terminating` state until you remove the annotation or change it to `"Reconciling"`.

#### Example 2: Reconcile a Single Resource While the Namespace or Cluster is Paused

If your namespace (`ConfigConnectorContext`) or cluster (`ConfigConnector`) has `spec.actuationMode: Paused`, you can selectively allow a single direct-managed resource to reconcile by setting `cnrm.cloud.google.com/actuation-mode: "Reconciling"`:

```yaml
apiVersion: logging.cnrm.cloud.google.com/v1beta1
kind: LoggingLogMetric
metadata:
  name: logginglogmetric-active
  annotations:
    cnrm.cloud.google.com/actuation-mode: "Reconciling"
spec:
  filter: "resource.type=gae_app AND severity>=ERROR"
  projectRef:
    external: "projects/my-project-id"
```

## Caveats

### Eventual state transitions & Jitter

* **Cluster and Namespace level (`ConfigConnector` / `ConfigConnectorContext`)**: Changing `spec.actuationMode` on `ConfigConnector` or `ConfigConnectorContext` does not immediately enqueue all resources in the cluster or namespace; resources observe the new mode on their next scheduled reconciliation (or when modified).
* **Resource level (`cnrm.cloud.google.com/actuation-mode`)**: Changing or removing the `cnrm.cloud.google.com/actuation-mode` annotation on an individual resource **immediately** triggers reconciliation for that resource.

---

## Further Reading

*   [Config Connector Operational Modes](config-connector-mode.md) for general configuration options (Identity, Billing, etc.).
*   [Controller Implementation Overrides](controller-configuration.md) for details on choosing between Direct, Terraform, and DCL controllers.
