# Greenfield Resource Implementation Status

This file tracks the progress of new "Direct" resource implementations.

## Implementation Phases
- **Phase 1: Skeleton**: CRD, Types, IdentityV2, and isolated package structure.
- **Phase 2: Brain**: Controller logic (Adapter), manual mappers, and Real GCP E2E fixtures (Minimal/Maximal).
- **Phase 3: Proof**: MockGCP implementation and alignment with recorded golden logs.

## Resource Tracker

| Resource | Service | Phase 1 PR | Phase 2 PR | Phase 3 PR | Current Phase | State |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| VertexAICachedContent | vertexai | [#7997](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7997) | - | - | 1 | CLOSED |
| VertexAICustomJob | vertexai | [#7996](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7996) | - | - | 1 | OPEN |
| VertexAIDataset | vertexai | [#7991](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7991) | - | - | 1 | OPEN |
| VertexAIEndpoint | vertexai | [#7994](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7994) | - | - | 1 | OPEN |
| VertexAIExampleStore | vertexai | [#7992](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7992) | - | - | 1 | MERGED |
| CloudIdentityDevice | cloudidentity | [#8077](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8077) | [#8077](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8077) | - | 2 | OPEN |
| ApigeeApiProduct | apigee | [#8078](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8078) | [#8078](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8078) | - | 2 | OPEN |
| NetworkConnectivityRegionalEndpoint | networkconnectivity | [#8071](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8071) | - | - | 1 | MERGED |
| KMSEKMConnection | kms | [#8076](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8076) | - | - | 1 | CLOSED |
| NotebookRuntime | notebooks | [#8075](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8075) | [#8075](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8075) | - | 2 | OPEN |
| NetworkSecurityTLSInspectionPolicy | networksecurity | [#8474](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8474) | - | - | 1 | OPEN |
| BigLakeCatalog | biglake | [#8411](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8411) | - | - | 1 | CLOSED |
| BeyondCorpClientGateway | beyondcorp | [#8419](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8419) | - | - | 1 | OPEN |
| BeyondCorpClientConnectorService | beyondcorp | [#8407](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8407) | - | - | 1 | MERGED |
| AutoMLDataset | automl | [#8417](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8417) | - | - | 1 | MERGED |
| AppHubServiceProjectAttachment | apphub | [#8418](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8418) | - | - | 1 | MERGED |
| ApiHubDeployment | apihub | [#8416](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8416) | - | - | 1 | MERGED |
| ApiHubDependency | apihub | [#8410](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8410) | - | - | 1 | OPEN |
| ApiHubCuration | apihub | [#8406](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8406) | - | - | 1 | OPEN |
| ApiHubAttribute | apihub | [#8412](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8412) | - | - | 1 | OPEN |
| ApiHubInstance | apihub | [#8424](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8424) | - | - | 1 | OPEN |
| ApigeeRegistryApi | apigeeregistry | [#8413](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8413) | - | - | 1 | OPEN |
| BigQueryAnalyticsHubDataExchange | analyticshub | [#8423](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8423) | - | - | 1 | CLOSED |
| AIStreamsCluster | aistreams | [#8420](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8420) | - | - | 1 | CLOSED |
| VertexAIModelDeploymentMonitoringJob | aiplatform | [#8425](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8425) | - | - | 1 | OPEN |
| VertexAIIndexEndpoint | aiplatform | [#8422](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8422) | - | - | 1 | OPEN |
| VertexAIIndex | aiplatform | [#8421](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8421) | - | - | 1 | OPEN |
| VertexAIHyperparameterTuningJob | aiplatform | [#8415](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8415) | - | - | 1 | OPEN |
| VertexAIFeatureOnlineStore | aiplatform | [#8408](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8408) | - | - | 1 | OPEN |
| VertexAIFeatureGroup | aiplatform | [#8414](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8414) | - | - | 1 | OPEN |
| SecretManagerRegionalSecret | secretmanager | [#8080](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8080) | - | - | 1 | OPEN |
| EventarcEnrollment | eventarc | [#8074](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8074) | [#8074](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8074) | - | 2 | MERGED |
| CertificateManagerTrustConfig | certificatemanager | [#8072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8072) | [#8072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8072) | - | 2 | OPEN |
| ArtifactRegistryVPCSCConfig | artifactregistry | [#8073](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8073) | - | - | 1 | OPEN |
| NetworkSecurityUrlList | networksecurity | [#7941](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7941) | - | - | 1 | OPEN |
| NetworkSecurityTLSInspectionPolicy | networksecurity | [#7943](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7943) | - | - | 1 | OPEN |
| NetworkSecurityGatewaySecurityPolicy | networksecurity | [#7942](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7942) | - | - | 1 | OPEN |
| ConnectorsConnection | connectors | [#8687](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8687) | - | - | 1 | OPEN |
| ApigeeRegistryInstance | apigeeregistry | [#8684](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8684) | - | - | 1 | OPEN |
| BlockchainNodeEngineNode | blockchainnodeengine | [#8691](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8691) | - | - | 1 | OPEN |
| CloudSecurityFramework | cloudsecuritycompliance | [#8841](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8841) | - | - | 1 | OPEN |
| CCInsightsQaScorecard | contactcenterinsights | [#8701](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8701) | - | - | 1 | OPEN |
| ContentWarehouseSchema | contentwarehouse | [#8686](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8686) | - | - | 1 | OPEN |
| DataLineageProcess | datalineage | [#9110](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9110) | - | - | 1 | OPEN |
| DevConnectConnection | developerconnect | [#8685](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8685) | - | - | 1 | OPEN |
| DeviceStreamingSession | devicestreaming | [#8839](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8839) | - | - | 1 | OPEN |
| BigQueryMigrationWorkflow | bigquerymigration | [#8699](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8699) | - | - | 1 | OPEN |
| CloudBatchResourceAllowance | batch | - | - | - | 0 | PLANNED |
| CloudAssetSavedQuery | cloudasset | [#8696](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8696) | - | - | 1 | OPEN |
| CloudBuildConnection | cloudbuild | [#8700](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8700) | - | - | 1 | OPEN |
| CloudTasksQueue | cloudtasks | [#8677](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8677) | - | - | 1 | OPEN |
| InfraManagerDeploymentGroup | config | [#8688](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8688) | - | - | 1 | OPEN |
| DataformFolder | dataform | [#9051](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9051) | - | - | 1 | OPEN |
| DatabaseMigrationJob | datamigration | [#8690](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8690) | - | - | 1 | OPEN |
| DataplexDataScan | dataplex | [#8693](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8693) | - | - | 1 | OPEN |
| DataprocSession | dataproc | [#8890](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8890) | - | - | 1 | OPEN |
| DialogflowConversationDataset | dialogflow | [#8703](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8703) | - | - | 1 | OPEN |
| ModelArmorTemplate | modelarmor | [#8774](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8774) | - | - | 1 | OPEN |
| DiscoveryEngineIdentityMappingStore | discoveryengine | [#8889](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8889) | - | - | 1 | OPEN |
| DiscoveryEngineSampleQuerySet | discoveryengine | [#9055](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9055) | - | - | 1 | OPEN |
| SecurityCenterBigQueryExport | securitycenter | - | - | - | 0 | PLANNED |
| SecurityCenterMuteConfig | securitycenter | [#8831](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8831) | - | - | 1 | OPEN |
| SecurityCenterManagementEventThreatDetectionCustomModule | securitycentermanagement | [#8777](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8777) | - | - | 1 | OPEN |
| SecurityCenterManagementSecurityHealthAnalyticsCustomModule | securitycentermanagement | [#8765](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8765) | - | - | 1 | OPEN |
| CloudRunWorkerPool | run | [#8749](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8749) | - | - | 1 | OPEN |
| NetworkSecurityAddressGroup | networksecurity | [#8755](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8755) | - | - | 1 | OPEN |
| NetworkSecurityAuthzPolicy | networksecurity | [#8760](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8760) | - | - | 1 | OPEN |
| CloudRunInstance | run | [#8766](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8766) | - | - | 1 | OPEN |
| NetworkSecurityBackendAuthenticationConfig | networksecurity | [#8869](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8869) | - | - | 1 | OPEN |
| NetworkSecurityDnsThreatDetector | networksecurity | [#8758](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8758) | - | - | 1 | OPEN |
| NetworkSecurityFirewallEndpoint | networksecurity | [#8768](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8768) | - | - | 1 | OPEN |
| NetworkSecurityFirewallEndpointAssociation | networksecurity | [#8756](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8756) | - | - | 1 | OPEN |
| NetworkSecurityInterceptDeployment | networksecurity | [#8867](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8867) | - | - | 1 | OPEN |
| NetworkSecurityInterceptDeploymentGroup | networksecurity | [#8769](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8769) | - | - | 1 | OPEN |
| NetworkSecurityInterceptEndpointGroup | networksecurity | [#8835](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8835) | - | - | 1 | OPEN |
| NetworkSecurityInterceptEndpointGroupAssociation | networksecurity | [#8767](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8767) | - | - | 1 | OPEN |
| NetworkSecurityMirroringDeployment | networksecurity | [#9049](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9049) | - | - | 1 | OPEN |
| NetworkSecurityMirroringDeploymentGroup | networksecurity | [#8750](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8750) | - | - | 1 | OPEN |
| NetworkSecurityMirroringEndpointGroup | networksecurity | - | - | - | 0 | PLANNED |
| NetworkSecurityMirroringEndpointGroupAssociation | networksecurity | [#9053](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9053) | - | - | 1 | OPEN |
| NetworkSecurityPartnerSSEGateway | networksecurity | [#8770](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8770) | - | - | 1 | OPEN |
| NetworkSecurityPartnerSSERealm | networksecurity | [#8741](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8741) | - | - | 1 | OPEN |
| NetworkSecuritySACRealm | networksecurity | [#9054](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9054) | - | - | 1 | OPEN |
| NetworkSecuritySecurityProfile | networksecurity | [#9047](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9047) | - | - | 1 | OPEN |
| NetworkSecuritySecurityProfileGroup | networksecurity | [#8820](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8820) | - | - | 1 | OPEN |
| NetworkSecurityTlsInspectionPolicy | networksecurity | [#7943](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7943) | - | - | 1 | OPEN |
| NetworkSecurityUrlList | networksecurity | [#7941](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7941) | - | - | 1 | OPEN |
| AppOptimizeReport | appoptimize | [#9015](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9015) | - | - | 1 | OPEN |
| ContactCenterInsightsConversation | contactcenterinsights | [#9016](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9016) | - | - | 1 | OPEN |
| BlockchainNodeEngineBlockchainNode | blockchainnodeengine | [#9017](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9017) | - | - | 1 | OPEN |
| CESApp | ces | [#9018](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9018) | - | - | 1 | OPEN |
| ContentWarehouseDocument | contentwarehouse | [#9019](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9019) | - | - | 1 | OPEN |
| ConfigDeploymentGroup | config | [#9020](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9020) | - | - | 1 | OPEN |
| BigQueryDataTransferTransferConfig | bigquerydatatransfer | [#9021](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9021) | - | - | 1 | OPEN |
| BigQueryReservationReservationGroup | bigqueryreservation | [#9022](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9022) | - | - | 1 | OPEN |
| BigQueryMigrationMigrationWorkflow | bigquerymigration | [#9023](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9023) | - | - | 1 | OPEN |
| CloudSecurityComplianceCloudControl | cloudsecuritycompliance | [#9024](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9024) | - | - | 1 | CLOSED |
