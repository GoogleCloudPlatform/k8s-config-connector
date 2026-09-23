*   Special shout-outs to @acpana, @anfernee, @anhdle-sso, @barney-s, @cdmello-g, @cheftako, @gemmahou, @GinnyJI, @himanigulati01, @ldanielmadariaga, @maqiuyujoyce, @sdowell, @suwandim, and @ziyue-101 for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `AIPlatformModel`
    *   Manage [Vertex AI models](https://cloud.google.com/vertex-ai/docs/model-registry/introduction)

*   `APIHubAPI`
    *   Manage [API Hub APIs](https://cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.apis)

*   `APIHubAttribute`
    *   Manage [API Hub attributes](https://cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.attributes)

*   `APIHubCuration`
    *   Manage [API Hub curation](https://cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.curations)

*   `APIHubDependency`
    *   Manage [API Hub dependencies](https://cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.dependencies)

*   `APIHubDeployment`
    *   Manage [API Hub deployments](https://cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.deployments)

*   `APIHubExternalAPI`
    *   Manage [API Hub external APIs](https://cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.externalApis)

*   `CCInsightsConversation`
    *   Manage [Contact Center AI Insights conversations](https://docs.cloud.google.com/gemini-enterprise-cx/insights/reference/rest/v1/projects.locations.conversations)

*   `CCInsightsPhraseMatcher`
    *   Manage [Contact Center AI Insights phrase matchers](https://docs.cloud.google.com/gemini-enterprise-cx/insights/reference/rest/v1/projects.locations.phraseMatchers)

*   `CCInsightsView`
    *   Manage [Contact Center AI Insights views](https://docs.cloud.google.com/gemini-enterprise-cx/insights/reference/rest/v1/projects.locations.views)

*   `ConfigDeliveryFleetPackage`
    *   Manage [Config Delivery fleet packages](https://docs.cloud.google.com/kubernetes-engine/config-sync/docs/reference/rest/v1/projects.locations.fleetPackages)

*   `DialogflowGenerator`
    *   Manage [Dialogflow generators](https://cloud.google.com/dialogflow/es/docs/reference/rest/v2/projects.locations.generators)

*   `LiveStreamAsset`
    *   Manage [Live Stream assets](https://cloud.google.com/livestream/docs/reference/rest/v1/projects.locations.assets)

*   `NetworkSecurityInterceptDeployment`
    *   Manage [Network Security intercept deployments](https://cloud.google.com/network-security-integration/docs/reference/rest/v1/projects.locations.interceptDeployments)

*   `NotebookInstanceV2`
    *   Manage [Agent Platform Workbench instance](https://docs.cloud.google.com/gemini-enterprise-agent-platform/notebooks/workbench/reference/rest/v2/projects.locations.instances)

*   `RedisClusterEndpoint`
    *   Manage [Memorystore for Redis Cluster endpoints](https://cloud.google.com/memorystore/docs/cluster/reference/rest/v1/projects.locations.clusters#ClusterEndpoint)

*   `SaasServiceMgmtRelease`
    *   Manage [SaaS Service Management releases](https://cloud.google.com/saas-runtime/docs/reference/rest/v1beta1/projects.locations.releases)

*   `VMMigrationGroup`
    *   Manage [VM Migration groups](https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.groups)

## New Fields:

*   `NetworkServicesEdgeCacheService`
    *   Added `routeMethods` and `compressionMode` support.

## Bug Fixes:

*   [`NetworkConnectivityInternalRange`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12615)
    *   Fixed continuous reconciliation loop and 400 error on `NetworkConnectivityInternalRange` when auto-allocation fields are set.

## Other:

*   **BeyondCorpClientGateway**: Removed support for the deprecated BeyondCorp Enterprise client connector.
