# Legacy Behavior

This document tracks legacy behaviors and configurations in Config Connector that are still supported for backward compatibility but are no longer recommended for new installations.

## Per-Resource Reconciler Annotation

Before the introduction of `ConfigConnectorContext` overrides, users could specify which controller implementation to use for a specific resource using an annotation.

| Annotation | Values | Description |
| :--- | :--- | :--- |
| `alpha.cnrm.cloud.google.com/reconciler` | `direct`, `tf`, `dcl` | Specifies the underlying controller implementation for this specific resource instance. |

### Why it is Legacy

Using per-resource annotations makes it difficult to manage controller transitions at scale. We recommend using `ConfigConnectorContext` to manage overrides at the namespace level, which provides a more centralized and manageable way to configure your platform.

See [Controller Configuration](controller-configuration.md) for the recommended way to override controllers.

### Backward Compatibility

Config Connector still honors the `alpha.cnrm.cloud.google.com/reconciler` annotation. If present, it takes precedence over any `ConfigConnectorContext` level overrides.
