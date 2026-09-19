# Release 1.157.0

** This version is not yet released; this document is gathering release notes
for the future release **

*   Special shout-outs to GinnyJI, acpana, anfernee, anhdle-sso, barney-s, cdmello-g, cheftako, gemmahou, himanigulati01, ldanielmadariaga, maqiuyujoyce, sdowell, suwandim, and ziyue-101 for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   [`AIPlatformModel`](https://cloud.google.com/config-connector/docs/reference/resource-docs/aiplatform/aiplatformmodel)
    *   Manage [Vertex AI models](https://cloud.google.com/vertex-ai/docs/model-registry/introduction)

*   [`APIHubAttribute`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apihub/apihubattribute)
    *   Manage [API Hub attributes](https://cloud.google.com/apihub/docs/reference/rest/v1/projects.locations.attributes)

*   [`APIHubCuration`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apihub/apihubcuration)
    *   Manage [API Hub curation](https://cloud.google.com/apihub/docs/reference/rest/v1/projects.locations.apis.versions.specs)

*   [`APIHubDependency`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apihub/apihubdependency)
    *   Manage [API Hub dependencies](https://cloud.google.com/apihub/docs/reference/rest/v1/projects.locations.dependencies)

*   [`APIHubExternalAPI`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apihub/apihubexternalapi)
    *   Manage [API Hub external APIs](https://cloud.google.com/apihub/docs/reference/rest/v1/projects.locations.externalApis)

*   [`CCInsightsConversation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/contactcenterinsights/ccinsightsconversation)
    *   Manage [Contact Center AI Insights conversations](https://cloud.google.com/contact-center/insights/docs/reference/rest/v1/projects.locations.conversations)

*   [`DialogflowGenerator`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dialogflow/dialogflowgenerator)
    *   Manage [Dialogflow generators](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.generators)

*   [`LiveStreamAsset`](https://cloud.google.com/config-connector/docs/reference/resource-docs/livestream/livestreamasset)
    *   Manage [Live Stream assets](https://cloud.google.com/video-stitcher/docs/reference/rest/v1/projects.locations.assets)

*   [`NetworkSecurityInterceptDeployment`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networksecurity/networksecurityinterceptdeployment)
    *   Manage [Network Security intercept deployments](https://cloud.google.com/network-security/docs/reference/rest/v1/projects.locations.interceptDeployments)

*   [`RedisClusterEndpoint`](https://cloud.google.com/config-connector/docs/reference/resource-docs/redis/redisclusterendpoint)
    *   Manage [Memorystore for Redis Cluster endpoints](https://cloud.google.com/memorystore/docs/cluster/reference/rest/v1/projects.locations.clusters.endpoints)

*   [`VMMigrationGroup`](https://cloud.google.com/config-connector/docs/reference/resource-docs/vmmigration/vmmigrationgroup)
    *   Manage [VM Migration groups](https://cloud.google.com/migrate/vmmigration/docs/reference/rest/v1/projects.locations.groups)

## New Fields:

*   [`NetworkServicesEdgeCacheService`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networkservices/networkservicesedgecacheservice)
    *   Added `routeMethods` and `compressionMode` support.

## Bug Fixes:

*   [`NetworkConnectivityInternalRange`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12615)
    *   Fixed continuous reconciliation loop and 400 error on `NetworkConnectivityInternalRange` when auto-allocation fields are set.

*   [`NotebookInstanceV2`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12314)
    *   Fixes and improvements for `NotebookInstanceV2` resource.

*   [`ComputeRouterNAT`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12716)
    *   Fixed issue where `ComputeRouterNATRef.Normalize` fallback function constructed external identity before checking resource readiness.

## Other:

*   **BeyondCorpClientGateway**: Removed support for the deprecated BeyondCorp Enterprise client connector.
