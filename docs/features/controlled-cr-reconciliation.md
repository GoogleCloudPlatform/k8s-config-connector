# Controlled CR Reconciliation

Config Connector supports **Controlled CR Reconciliation**, allowing you to explicitly dictate which GCP resources Config Connector handles, both at the cluster (global) level and the namespace (local) level. 

By defining `ResourceSettings` within the `ConfigConnector` and `ConfigConnectorContext` configurations, platform administrators can opt-in or opt-out specific Resource groups/kinds to conserve cluster memory, avoid overlapping management, or restrict certain namespaces from provisioning specific resources.

## Precedence and Constraints

### Valid Modes
Resource lists can act in one of two modes:
* **Exclusive (Deny-list)**: `Enabled: false`. The system reconciles everything EXCEPT the resources explicitly listed. This is the default mode.
* **Inclusive (Allow-list)**: `Enabled: true`. The system reconciles ONLY the resources explicitly listed and ignores all others.

> [!WARNING]
> **Mode Conflicts Forbidden**: To prevent unpredictable behaviors, you cannot mix inclusive and exclusive modes across configurations. The `ConfigConnector` (global) and `ConfigConnectorContext` (namespace) **must** utilize the same `Enabled` value. If they clash, the `ConfigConnectorContext` will report an error (`Healthy: False`) and the Manager pod for that namespace will refuse to start tracking reconciliation until the conflict is resolved.

### Additive Layering
Because both global and local settings use the same mode, they combine **additively**:
* **In Exclusive Mode**: A resource is excluded (ignored) if it is listed in the `ConfigConnector` **OR** the `ConfigConnectorContext`.
* **In Inclusive Mode**: A resource is included (reconciled) if it is listed in the `ConfigConnector` **OR** the `ConfigConnectorContext`.

## Configuration Options

To configure this feature, manipulate the `ResourceSettings` block in your `ConfigConnector` and `ConfigConnectorContext` definitions.

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector   # OR ConfigConnectorContext
metadata:
  name: configconnector.core.cnrm.cloud.google.com # Or 'configconnectorcontext.core.cnrm.cloud.google.com'
spec:
  # ... other specs ...
  experiments:
    resourceSettings:
      enabled: false  # (Optional) Defaults to false. `false` = Exclusive, `true` = Inclusive
      resources:
        # Exclude/Include an entire Group
        - group: pubsub.cnrm.cloud.google.com
        # Exclude/Include a specific Kind within a Group
        - group: storage.cnrm.cloud.google.com
          kind: StorageBucket 
```

### Example: Exclusion (Deny-list)

In this example, the global `ConfigConnector` ignores all `PubSub` resources cluster-wide, and a specific `ConfigConnectorContext` further ignores `StorageBucket` resources for its namespace. 

```yaml
# Global Configurator (Excludes PubSub)
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  experiments:
    resourceSettings:
      enabled: false
      resources:
      - group: pubsub.cnrm.cloud.google.com
---
# Namespace Configurator (Further excludes StorageBucket)
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: test-namespace
spec:
  experiments:
    resourceSettings:
      enabled: false
      resources:
      - group: storage.cnrm.cloud.google.com
        kind: StorageBucket
```

With this setup:
* The `test-namespace` WILL NOT reconcile any `PubSub` resources (disabled globally).
* The `test-namespace` WILL NOT reconcile any `StorageBucket` resources (disabled locally).
* Other namespaces WILL reconcile `StorageBucket`s, but NOT `PubSub` resources.

### Example: Inclusion (Allow-list)

In this example, the system only manages explicitly authorized resources.

```yaml
# Global Configurator (Permits IAMServiceAccounts)
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  experiments:
    resourceSettings:
      enabled: true
      resources:
      - group: iam.cnrm.cloud.google.com
        kind: IAMServiceAccount
---
# Namespace Configurator (Further permits PubSubTopics)
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: data-namespace
spec:
  experiments:
    resourceSettings:
      enabled: true
      resources:
      - group: pubsub.cnrm.cloud.google.com
        kind: PubSubTopic
```

With this setup:
* The `data-namespace` WILL reconcile `IAMServiceAccount` and `PubSubTopic` ONLY. Rest will be ignored.
* Other namespaces WILL reconcile `IAMServiceAccount` ONLY. Rest will be ignored.
