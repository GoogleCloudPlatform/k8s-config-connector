# Pause KCC

> `Feature State`: `alpha` as of version v1.114

KCC can be configured to pause actuation of resouces on the cloud provider
(GCP). K8s objects continue to be reconciled with the api server but any
interaction with the cloud provider should be paused. This can be helpful for
debugging purposes or to have a hot standby.

## Pausing

We extended the API definitions for the Config Connector and Config ConnectorContext resources to
support a new field, `spec.actuationMode`. The field's current supported values are `Reconciling`
and `Paused` with `Reconciling` being the default for backwards compatiblity.
As such, KCC can be "paused" both globally and on a per-namespace level if running in namespace mode

### Pausing Globally

To pause KCC across namespaces it is sufficient to set the Config Connector's `actuationMode: Paused`. This will work when KCC runs in `Cluster` and `Namespaced` mode. To eventually resume actuation just set the field back to `Reconciling`.

### Pausing Per Namespace

When KCC is running in `Namespaced` mode (and only then), operators can set
`actuationMode: Paused` on the  `Config ConnectorContext` resource. To eventually
resume actuation for that namespace set the field back to `Reconciling`. The `actuationMode`
value on the `Config ConnectorContext` takes precedence over the value in `Config Connector`.

### Reconciling Per Namespace (only)

It can be handy to have KCC globally paused but reconciling on a per
namespace level. To do this make sure KCC is running in `Namespaced` mode
and the `actuationMode: Paused` on the `Config Connector` resource. Then
reconciling can be turned on for a namespace by changing the `actuationMode`
field for the`Config ConnectorContext` to `Reconciling` for that namespace.
NOTE: you can avoid any pausing in actuation by first changing the
`Config ConnectorContext` actuationMode.

## Caveats

### Eventual state transitions & Jitter

When you un-pause actuation, reconciliation will not happen immediately, but will instead start happening on the normal re-reconciliation interval (with jitter). We may enhance this behavior in future 