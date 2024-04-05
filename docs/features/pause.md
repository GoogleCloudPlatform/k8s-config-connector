# Pause KCC

> `Feature State`: KCC version v1.114+ (alpha)

KCC can be configured to pause actuation of resouces on the cloud provider
(GCP). K8s objects continue to be reconciled with the api server but any
interaction with the cloud provider should be paused. This can be helpful for
debugging purposes or to have a hot standby.

## Pausing

We extended the API definitions for the ConfigConnector and ConfigConnectorContext resources to
support a new field, `spec.actuationMode`. The field's current supported values are `Reconciling`
and `Paused` with `Reconciling` being the default for backwards compatiblity.
As such, KCC can be "paused" both globablly and on a per namespace level if running in namesapce mode

### Pausing Globally

To pause KCC across namespaces it is suffcient to set the ConfigConnector's `actuationMode: Paused`. This will work when KCC runs in `Cluster` and `Namesapced` mode. To eventually resume actuation just set the field back to `Reconciling`.

### Pausing Per Namespace

When KCC is running in `Namesapced` mode (and only then), operators can set
`actuationMode: Paused` on the  `ConfigConnectorContext` resource. To eventually
resume actuation for that namespace set the field back to `Reconciling`.

### Reconciling Per Namespace (only)

It can be handy to have KCC globally paused but reconciling on a per
namespace level. To do this make sure KCC is running in `Namespaced` mode
and the `actuationMode: Paused` on the `ConfigConnector` resource. Then
reconciling can be turned on for a namespace by changing the `actuationMode`
field for the`ConfigConnectorContext` to `Reconciling` for that namespace.
NOTE: you can avoid any pausing in actuation by first changing the
`ConfigConnectorContext` actuationMode first.

## Caveats

### Eventual state transitions & Jitter

When changing actuation modes for a cluster or namespace, a KCC resource should eventually actuate
(or not depending on the intent). Note that, in order to protect the k8s api server, KCC resources
are re-enqueued with a jitter amount.  