## New Fields:

*   [`NetworkServicesEdgeCacheService`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networkservices/networkservicesedgecacheservice)
    *   Added `spec.routing.pathMatchers.routeRules.routeAction.compressionMode` field.
    *   Added `spec.routing.pathMatchers.routeRules.routeMethods` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in behaviour. The API is unchanged. To use the direct reconciler, add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object. The following resources now have direct reconciliation support (and we list some of the issues that this fixes):

*   [`NetworkServicesEdgeCacheService`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networkservices/networkservicesedgecacheservice)
    *   Migrated to Direct Controller architecture.
