# Greenfield Resource Implementation Status

This file tracks the progress of new "Direct" resource implementations.

## Implementation Phases
- **Phase 1: Skeleton**: CRD, Types, IdentityV2, and isolated package structure.
- **Phase 2: Brain**: Controller logic (Adapter), manual mappers, and Real GCP E2E fixtures (Minimal/Maximal).
- **Phase 3: Proof**: MockGCP implementation and alignment with recorded golden logs.

## Resource Tracker

| Resource | Service | Phase 1 PR | Phase 2 PR | Phase 3 PR | Current Phase | State |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| VertexAICachedContent | vertexai | [#7997](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7997) | - | - | 1 | OPEN |
| VertexAICustomJob | vertexai | [#7996](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7996) | - | - | 1 | OPEN |
| VertexAIDataset | vertexai | [#7991](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7991) | - | - | 1 | OPEN |
| VertexAIEndpoint | vertexai | [#7994](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7994) | - | - | 1 | OPEN |
| VertexAIExampleStore | vertexai | [#7992](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7992) | - | - | 1 | OPEN |
| CloudIdentityDevice | cloudidentity | [#8077](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8077) | [#8077](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8077) | - | 2 | OPEN |
| ApigeeApiProduct | apigee | [#8078](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8078) | [#8078](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8078) | - | 2 | OPEN |
| NetworkConnectivityRegionalEndpoint | networkconnectivity | [#8071](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8071) | - | - | 1 | OPEN |
| KMSEKMConnection | kms | [#8076](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8076) | - | - | 1 | OPEN |
| NotebookRuntime | notebooks | [#8075](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8075) | [#8075](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8075) | - | 2 | OPEN |
| NetworkSecurityTLSInspectionPolicy | networksecurity | [#8474](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8474) | - | - | 1 | OPEN |
| BigLakeCatalog | biglake | [#8411](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8411) | - | - | 1 | OPEN |
| BeyondCorpClientGateway | beyondcorp | [#8419](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8419) | - | - | 1 | OPEN |
| BeyondCorpClientConnectorService | beyondcorp | [#8407](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8407) | - | - | 1 | OPEN |
| AutoMLDataset | automl | [#8417](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8417) | - | - | 1 | OPEN |
| AppHubServiceProjectAttachment | apphub | [#8418](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8418) | - | - | 1 | OPEN |
| ApiHubDeployment | apihub | [#8416](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8416) | - | - | 1 | OPEN |
| ApiHubDependency | apihub | [#8410](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8410) | - | - | 1 | OPEN |
| ApiHubCuration | apihub | [#8406](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8406) | - | - | 1 | OPEN |
| ApiHubAttribute | apihub | [#8412](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8412) | - | - | 1 | OPEN |
| ApiHubInstance | apihub | [#8424](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8424) | - | - | 1 | OPEN |
| ApigeeRegistryApi | apigeeregistry | [#8413](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8413) | - | - | 1 | OPEN |
| BigQueryAnalyticsHubDataExchange | analyticshub | [#8423](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8423) | - | - | 1 | OPEN |
| AIStreamsCluster | aistreams | [#8420](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8420) | - | - | 1 | OPEN |
| VertexAIModelDeploymentMonitoringJob | aiplatform | [#8425](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8425) | - | - | 1 | OPEN |
| VertexAIIndexEndpoint | aiplatform | [#8422](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8422) | - | - | 1 | OPEN |
| VertexAIIndex | aiplatform | [#8421](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8421) | - | - | 1 | OPEN |
| VertexAIHyperparameterTuningJob | aiplatform | [#8415](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8415) | - | - | 1 | OPEN |
| VertexAIFeatureOnlineStore | aiplatform | [#8408](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8408) | - | - | 1 | OPEN |
| VertexAIFeatureGroup | aiplatform | [#8414](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8414) | - | - | 1 | OPEN |
| SecretManagerRegionalSecret | secretmanager | [#8080](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8080) | - | - | 1 | OPEN |
| EventarcEnrollment | eventarc | [#8074](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8074) | [#8074](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8074) | - | 2 | OPEN |
| CertificateManagerTrustConfig | certificatemanager | [#8072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8072) | [#8072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8072) | - | 2 | OPEN |
| ArtifactRegistryVPCSCConfig | artifactregistry | [#8073](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8073) | - | - | 1 | OPEN |
| NetworkSecurityUrlList | networksecurity | [#7941](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7941) | - | - | 1 | OPEN |
| NetworkSecurityTLSInspectionPolicy | networksecurity | [#7943](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7943) | - | - | 1 | OPEN |
| NetworkSecurityGatewaySecurityPolicy | networksecurity | [#7942](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7942) | - | - | 1 | OPEN |
| ConnectorsConnection | connectors | - | - | - | 0 | PLANNED |
| ApigeeRegistryInstance | apigeeregistry | - | - | - | 0 | PLANNED |
| BlockchainNodeEngineNode | blockchainnodeengine | - | - | - | 0 | PLANNED |
| CloudSecurityFramework | cloudsecuritycompliance | - | - | - | 0 | PLANNED |
| CCInsightsQaScorecard | contactcenterinsights | - | - | - | 0 | PLANNED |
| ContentWarehouseSchema | contentwarehouse | - | - | - | 0 | PLANNED |
| DataLineageProcess | datalineage | - | - | - | 0 | PLANNED |
| DevConnectConnection | developerconnect | - | - | - | 0 | PLANNED |
| DeviceStreamingDeviceSession | devicestreaming | - | - | - | 0 | PLANNED |
| BigQueryMigrationWorkflow | bigquerymigration | - | - | - | 0 | PLANNED |
| CloudBatchResourceAllowance | batch | - | - | - | 0 | PLANNED |
| CloudAssetSavedQuery | cloudasset | - | - | - | 0 | PLANNED |
| CloudBuildConnection | cloudbuild | - | - | - | 0 | PLANNED |
| CloudTasksQueue | cloudtasks | - | - | - | 0 | PLANNED |
| InfraManagerDeploymentGroup | config | - | - | - | 0 | PLANNED |
| DataformFolder | dataform | - | - | - | 0 | PLANNED |
| DatabaseMigrationJob | datamigration | - | - | - | 0 | PLANNED |
| DataplexDataScan | dataplex | - | - | - | 0 | PLANNED |
| DataprocSession | dataproc | - | - | - | 0 | PLANNED |
| DialogflowConversationDataset | dialogflow | - | - | - | 0 | PLANNED |
| ModelArmorTemplate | modelarmor | - | - | - | 0 | PLANNED |
| DiscoveryEngineIdentityMappingStore | discoveryengine | - | - | - | 0 | PLANNED |
| DiscoveryEngineSampleQuerySet | discoveryengine | - | - | - | 0 | PLANNED |
| SecurityCenterBigQueryExport | securitycenter | - | - | - | 0 | PLANNED |
| SecurityCenterMuteConfig | securitycenter | - | - | - | 0 | PLANNED |
| SecurityCenterManagementEventThreatDetectionCustomModule | securitycentermanagement | - | - | - | 0 | PLANNED |
| SecurityCenterManagementSecurityHealthAnalyticsCustomModule | securitycentermanagement | - | - | - | 0 | PLANNED |
| CloudRunWorkerPool | run | - | - | - | 0 | PLANNED |
| NetworkSecurityAddressGroup | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityAuthzPolicy | networksecurity | - | - | - | 0 | PLANNED |
| CloudRunInstance | run | - | - | - | 0 | PLANNED |
| NetworkSecurityBackendAuthenticationConfig | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityDnsThreatDetector | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityFirewallEndpoint | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityFirewallEndpointAssociation | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityInterceptDeployment | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityInterceptDeploymentGroup | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityInterceptEndpointGroup | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityInterceptEndpointGroupAssociation | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityMirroringDeployment | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityMirroringDeploymentGroup | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityMirroringEndpointGroup | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityMirroringEndpointGroupAssociation | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityPartnerSSEGateway | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityPartnerSSERealm | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecuritySACRealm | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecuritySecurityProfile | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecuritySecurityProfileGroup | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityTlsInspectionPolicy | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityUrlList | networksecurity | - | - | - | 0 | PLANNED |
