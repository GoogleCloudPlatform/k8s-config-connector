*   Special shout-outs to @acpana, @anfernee, @anhdle-sso, @barney-s, @cdmello-g, @cheftako, @gemmahou, @GinnyJI, @himanigulati01, @ldanielmadariaga, @maqiuyujoyce, @sdowell, @suwandim, and @ziyue-101 for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `AIPlatformModel`
    *   Manage [Vertex AI models](https://cloud.google.com/vertex-ai/docs/model-registry/introduction)

*   `APIHubExternalAPI`
    *   Manage [API Hub external APIs](https://cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.externalApis)

*   `NotebookInstanceV2`
    *   Manage [Agent Platform Workbench instance](https://docs.cloud.google.com/gemini-enterprise-agent-platform/notebooks/workbench/reference/rest/v2/projects.locations.instances)

*   `RedisClusterEndpoint`
    *   Manage [Memorystore for Redis Cluster endpoints](https://cloud.google.com/memorystore/docs/cluster/reference/rest/v1/projects.locations.clusters#ClusterEndpoint)

## New Fields:

*   `NetworkServicesEdgeCacheService`
    *   Added `routeMethods` and `compressionMode` support.

## Bug Fixes:

*   [`NetworkConnectivityInternalRange`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12615)
    *   Fixed continuous reconciliation loop and 400 error on `NetworkConnectivityInternalRange` when auto-allocation fields are set.

## Other:

*   **BeyondCorpClientGateway**: Removed support for the deprecated BeyondCorp Enterprise client connector.
