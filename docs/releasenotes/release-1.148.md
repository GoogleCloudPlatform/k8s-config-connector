# Release 1.148.0

*   Special shout-outs to acpana, anhdle-sso, barney-s, cheftako, dhavalbhensdadiya-crest, eugenenuke, gemmahou, gurusai-voleti, justinsb, katrielt, maqiuyujoyce, suwandim, and xiaoweim for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   [`ParameterManagerParameterVersion`](https://cloud.google.com/config-connector/docs/reference/resource-docs/parametermanager/parametermanagerparameterversion)

    *   Manage [Parameter Manager parameter versions](https://cloud.google.com/secret-manager/docs/parameter-manager) which allows you to manage regional parameters.

## Bug Fixes:

*   [SQLInstance](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7145) Fixed an issue where `settings.dataCacheConfig` was incorrectly detected as different when `dataCacheEnabled` was `false`.
*   [SQLInstance](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7200) Updated matching functions to treat nil values in KRM as equivalent to empty or default objects in GCP, preventing unnecessary re-reconciliation loops.
*   [TagKey/TagValue](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6943) Handle `ALREADY_EXISTS` error in TagKey and TagValue controllers by acquiring the existing resource.
*   [BigQueryAnalyticsHubDataExchange](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6774) Added structured reporting diff to improve visibility into resource changes and fixed reconciliation logic errors.
*   [CloudBuildTrigger](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7115) Restored missing descriptions in the CRD.
*   [RunService](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6693) Fixed a typo in environment variable values in samples and test fixtures.

## New features:

*   [MultiClusterLeaseSpec](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6919) now supports integration with a syncer for KRM objects. This will help KCC take ownership of resources with service generated IDs.
*   [kompanion](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7202) Added a Model Context Protocol (MCP) server to the `kompanion` tool to enable AI IDEs and assistants to interact with KCC resources.
*   [Config Connector controllers](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7075) Added a `--skip-name-validation` flag to bypass duplicate controller name checks during registration, facilitating integration tests and multi-manager scenarios.

## Documentation:

*   Added a [comprehensive guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/docs/features/controller-configuration.md) for controller configuration, detailing Direct, Terraform, and DCL controllers, including precedence rules and overrides.
*   Added [documentation](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/docs/features/containerresource.md) for enabling VerticalPodAutoscaler (VPA) for Config Connector pods using `ControllerResource` and `NamespacedControllerResource`.
*   Added [documentation](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/docs/cli/README.md) for the `config-connector` CLI and specifically for the `preview` command.
