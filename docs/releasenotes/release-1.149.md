# Release 1.149.1

*   Special shout-outs to acpana, barney-s, cheftako, dhavalbhensdadiya-crest, gemmahou, justinsb, katrielt, ldanielmadariaga, maqiuyujoyce, suwandim, and xiaoweim for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   [`NetworkServicesLBRouteExtension`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networkservices/networkserviceslbrouteextension)

    *   Manage [LB route extensions](https://cloud.google.com/service-extensions/docs/optimize-proxies-lb-route-extensions) which allow you to inject custom logic into the load balancing path.

*   [`ParameterManagerParameterVersion`](https://cloud.google.com/config-connector/docs/reference/resource-docs/parametermanager/parametermanagerparameterversion)

    *   Manage [Parameter Manager parameter versions](https://cloud.google.com/secret-manager/docs/parameter-manager) which allows you to manage regional parameters.

## New Fields:

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)

    *   Added `spec.nodeConfig.ephemeralStorageLocalSsdConfig.dataCacheCount` field to support GKE Data Cache.

*   [`ContainerNodePool`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containernodepool)

    *   Added `spec.nodeConfig.ephemeralStorageLocalSsdConfig.dataCacheCount` field to support GKE Data Cache.

## New Features:

*   [Controlled CR reconciliation](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6243) Added support for unmanaging specific resources via `resourceSettings` in `ConfigConnector` (global) and `ConfigConnectorContext` (per-namespace). This allows users to selectively disable reconciliation for specific Group/Kinds to save memory or manage resources differently.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in behaviour. The API is unchanged. To use the direct reconciler, add the `cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object.

*   [`BigQueryDatasetAccess`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigquery/bigquerydatasetaccess)

## Bug Fixes:

*   [Preview Tool](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7642) Fixed a connection error in the KCC preview tool and enforced read-only access to the cluster for improved security.
