*   Special shout-outs to @acpana, @anfernee, @anhdle-sso, @barney-s, @connorjclark, @eianiuk, @eugenenuke, @fkc1e100, @gemmahou, and @sdowell for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `AIPlatformContext`
    *   Manage [Vertex AI ML Metadata contexts](https://cloud.google.com/vertex-ai/docs/ml-metadata/introduction) to group lineage and metadata artifacts.

*   `APIHubAPI`
    *   Manage [API Hub APIs](https://cloud.google.com/api-hub/docs) to register and track APIs in API Hub.

*   `APIHubAttribute`
    *   Manage [API Hub attributes](https://cloud.google.com/api-hub/docs) to define custom metadata and taxonomy for APIs.

*   `APIHubCuration`
    *   Manage [API Hub curations](https://cloud.google.com/api-hub/docs) to curate API specifications and documentation.

*   `BigQueryReservationReservationGroup`
    *   Manage [BigQuery reservation groups](https://cloud.google.com/bigquery/docs/reservations-intro) to allocate dedicated workload capacity.

*   `CCInsightsPhraseMatcher`
    *   Manage [Contact Center Insights phrase matchers](https://docs.cloud.google.com/gemini-enterprise-cx/insights/phrase-matchers) to detect custom rules in conversations.

*   `CCInsightsQAScorecard`
    *   Manage [Contact Center Insights QA scorecards](https://docs.cloud.google.com/gemini-enterprise-cx/insights/qa-scorecards) to evaluate conversation quality.

*   `CloudBuildConnection`
    *   Manage [Cloud Build 2nd gen connections](https://cloud.google.com/build/docs/repositories) to connect source code repositories.

*   `ConfigDeliveryFleetPackage`
    *   Manage [Config Delivery fleet packages](https://cloud.google.com/kubernetes-engine/enterprise/config-sync/docs/concepts/fleet-packages) to package configurations for deployment across fleets.

*   `DatabaseMigrationConversionWorkspace`
    *   Manage [Database Migration Service conversion workspaces](https://cloud.google.com/database-migration/docs/overview) to convert schema and code during database migration.

*   `DataformFolder`
    *   Manage [Dataform folders](https://cloud.google.com/dataform/docs) to organize repository content.

*   `DataformReleaseConfig`
    *   Manage [Dataform release configurations](https://cloud.google.com/dataform/docs/release-configurations) to schedule and execute Dataform compilation runs.

*   `DataformTeamFolder`
    *   Manage [Dataform team folders](https://cloud.google.com/dataform/docs) to organize data assets across teams.

*   `DataformWorkflowConfig`
    *   Manage [Dataform workflow configurations](https://cloud.google.com/dataform/docs/workflow-configurations) to schedule and execute Dataform workflows.

*   `DataplexAsset`
    *   Manage [Dataplex assets](https://cloud.google.com/dataplex/docs) to attach storage or BigQuery datasets to Dataplex zones.

*   `DataplexDataAsset`
    *   Manage [Dataplex data assets](https://cloud.google.com/dataplex/docs) to manage metadata and governance for data assets.

*   `DataplexDataScan`
    *   Manage [Dataplex data scans](https://cloud.google.com/dataplex/docs/data-quality-tasks) to evaluate data quality and data profiles.

*   `DataplexGlossary`
    *   Manage [Dataplex glossaries](https://cloud.google.com/dataplex/docs/business-glossary) to define business terms and taxonomies.

*   `DevConnectAccountConnector`
    *   Manage [Developer Connect account connectors](https://cloud.google.com/developer-connect/docs) to configure integrations with developer tools.

*   `DeviceStreamingSession`
    *   Manage [Firebase device streaming sessions](https://firebase.google.com/docs/test-lab/android/device-streaming) to allocate cloud devices for interactive streaming.

*   `DialogflowConversationDataset`
    *   Manage [Dialogflow conversation datasets](https://cloud.google.com/dialogflow/docs) to manage datasets for training conversation models.

*   `DiscoveryEngineControl`
    *   Manage [Discovery Engine serving controls](https://cloud.google.com/generative-ai-app-builder/docs/serving-controls) to customize search and recommendation query serving.

*   `DiscoveryEngineSampleQuerySet`
    *   Manage [Discovery Engine sample query sets](https://cloud.google.com/generative-ai-app-builder/docs) to evaluate search quality.

*   `NetworkSecurityAddressGroup`
    *   Manage [Network Security address groups](https://cloud.google.com/firewall/docs/address-groups-overview) to manage IP address and CIDR collections.

*   `NetworkSecurityBackendAuthenticationConfig`
    *   Manage [Network Security backend authentication configurations](https://cloud.google.com/service-extensions/docs) to secure client-to-backend traffic.

*   `NetworkSecurityGatewaySecurityPolicy`
    *   Manage [Network Security gateway security policies](https://cloud.google.com/secure-web-proxy/docs/overview) to configure rules for Secure Web Proxy gateways.

*   `NetworkSecurityInterceptDeployment`
    *   Manage [Network Security intercept deployments](https://cloud.google.com/security/products/intercept) to intercept network traffic for security inspection.

*   `SaaSServiceMgmtUnitKind`
    *   Manage SaaS Service Management unit kinds to define resource schemas for deployment units.

*   `VertexAIFeatureOnlineStore`
    *   Manage [Vertex AI feature online stores](https://cloud.google.com/vertex-ai/docs/featurestore/latest/overview) to serve feature data for low-latency online prediction.

*   `WebSecurityScannerScanConfig`
    *   Manage [Web Security Scanner scan configurations](https://cloud.google.com/security-command-center/docs/concepts-web-security-scanner-overview) to detect security vulnerabilities in web applications.

## New Fields:

*   [`MemorystoreInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/memorystore/memorystoreinstance)
    *   Added support for `spec.pscAutoConnections[].ipVersion` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   `DialogflowCXAgent`
    *   Support direct reconciliation (opt-in).

The following resources now use the direct reconciler by default:

*   `AccessContextManagerGCPUserAccessBinding`
*   `ApigeeAddonsConfig`
*   `ApigeeNATAddress`
*   `ApigeeSyncAuthorization`
*   `AppEngineFirewallRule`
*   `AppEngineFlexibleAppVersion`
*   `AppEngineServiceSplitTraffic`
*   `AppEngineStandardAppVersion`
*   `BigQueryDatasetAccess`
*   `ComputeMachineImage`
*   `ComputeNetworkEndpoint`
*   `ComputeRegionPerInstanceConfig`
*   `DocumentAIProcessorDefaultVersion`
*   `FirebaseProject`
*   `HealthcareConsentStore`
*   `KMSCryptoKeyVersion`
*   `KMSKeyRingImportJob`
*   `KMSSecretCiphertext`
*   `NetworkServicesEdgeCacheService`

## New Features:

*   **Multi-Cluster Leader (MCL) Election Improvements**: Configured the kccmanager to fail closed at startup if no multi-cluster leader election configuration is found.
*   **Go Toolchain Update**: Config Connector images are now built with Go 1.26.11 (operator: Go 1.27.4) and upgraded dependencies.

## Bug Fixes:

*   [`ComputeForwardingRule`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeforwardingrule)
    *   Fixed Private Service Connect forwarding rules so `spec.ipAddress.addressRef` resolves to the `ComputeAddress` `selfLink` instead of the IP literal, and omitted `spec.loadBalancingScheme` defaults to `""` instead of `"EXTERNAL"`.

*   Fixed a proto comparison issue where reconcilers could be continuously invoked due to output-only or unmapped fields for resources including `OrgPolicyPolicy`, `APIQuotaPreference`, `GKEBackupBackupPlan`, `MetastoreService`, `DataCatalogTag`, and `DataformRepository`.

*   Fixed an issue where dynamic CRD schema updates were not propagated to running Terraform- and DCL-backed resource controllers until cnrm-controller-manager pods were restarted.
