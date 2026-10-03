# Pause KCC

> `Feature State`: `alpha` as of version v1.114

KCC can be configured to pause actuation of resources on the cloud provider
(GCP). K8s objects continue to be reconciled with the api server but any
interaction with the cloud provider should be paused. This can be helpful for
debugging purposes or to have a hot standby.

## Pausing

KCC can be paused at three levels:
1. Per-resource (via annotation)
2. Per-namespace (via `ConfigConnectorContext`)
3. Globally (via `ConfigConnector`)

The `cnrm.cloud.google.com/actuation-mode` annotation on an individual resource takes highest precedence, followed by the `spec.actuationMode` on `ConfigConnectorContext`, and finally `spec.actuationMode` on `ConfigConnector`.

### Pausing an Individual Resource

To pause actuation for an individual resource, set the annotation `cnrm.cloud.google.com/actuation-mode: "Paused"` on the resource:

```yaml
apiVersion: dataflow.cnrm.cloud.google.com/v1beta1
kind: DataflowJob
metadata:
  name: sample-job
  annotations:
    cnrm.cloud.google.com/actuation-mode: "Paused"
spec:
  # ...
```

To resume actuation for that resource, set the annotation to `Reconciling` or remove the annotation (in which case it falls back to the namespace/global setting).

#### Paused at Creation & Job Launch Gating (e.g. Kueue Integration)

You can create a resource with `cnrm.cloud.google.com/actuation-mode: "Paused"` to defer its initial creation and launch on Google Cloud.

When an object is created with `actuation-mode: "Paused"`:
* Finalizers are ensured on the Kubernetes object.
* The object's `Ready` condition is set to `status: "False"`, `reason: "Paused"`.
* No create, update, or launch API calls are made to GCP.

When an external orchestrator or queueing system (such as [Kueue](https://kueue.sigs.k8s.io/)) admits the workload, it can switch the annotation to `cnrm.cloud.google.com/actuation-mode: "Reconciling"` (or remove the annotation), which prompts KCC to launch/actuate the remote GCP resource.

#### Status Condition while Paused

While a resource is paused, KCC sets its `Ready` status condition to `status: "False"` with reason `Paused` and message `The resource is paused and will not be actuated onto the cloud provider`. This allows controllers and users to distinguish between an object that is paused awaiting launch and an object still initializing or encountering errors.

#### Deletion of Paused Resources

When a resource has `cnrm.cloud.google.com/actuation-mode: "Paused"`, KCC also skips deletion on GCP. If you delete the Kubernetes object while it is paused, KCC does not call the GCP delete API. The finalizers stay on the object, and it remains in the `Terminating` state.

To let the deletion proceed, set the annotation to `Reconciling` or remove it. KCC then deletes the resource on GCP, removes the finalizers, and the Kubernetes object goes away. If you remove the annotation, the namespace or global actuation mode applies, so the deletion only proceeds if that mode is `Reconciling`.

### Pausing Globally

To pause KCC across namespaces it is sufficient to set the Config Connector's `actuationMode: Paused`. This will work when KCC runs in `Cluster` and `Namespaced` mode. To eventually resume actuation just set the field back to `Reconciling`.

### Pausing Per Namespace

When KCC is running in `Namespaced` mode (and only then), operators can set
`actuationMode: Paused` on the `ConfigConnectorContext` resource. To eventually
resume actuation for that namespace set the field back to `Reconciling`. The `actuationMode`
value on the `ConfigConnectorContext` takes precedence over the value in `ConfigConnector`.

### Reconciling Per Namespace (only)

It can be handy to have KCC globally paused but reconciling on a per
namespace level. To do this make sure KCC is running in `Namespaced` mode
and the `actuationMode: Paused` on the `ConfigConnector` resource. Then
reconciling can be turned on for a namespace by changing the `actuationMode`
field for the `ConfigConnectorContext` to `Reconciling` for that namespace.
NOTE: you can avoid any pausing in actuation by first changing the
`ConfigConnectorContext` actuationMode.

## Caveats

### Eventual state transitions & Jitter


---

## Further Reading

*   [Config Connector Operational Modes](config-connector-mode.md) for general configuration options (Identity, Billing, etc.).
*   [Controller Implementation Overrides](controller-configuration.md) for details on choosing between Direct, Terraform, and DCL controllers.
