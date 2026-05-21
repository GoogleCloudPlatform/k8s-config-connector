# Greenfield Resource Implementation Status

This file tracks the progress of new "Direct" resource implementations.

## Implementation Phases
- **Phase 1: Skeleton**: CRD, Types, IdentityV2, and isolated package structure.
- **Phase 2: Brain**: Controller logic (Adapter), manual mappers, and Real GCP E2E fixtures (Minimal/Maximal).
- **Phase 3: Proof**: MockGCP implementation and alignment with recorded golden logs.

## Resource Tracker

| Resource | Service | Phase 1 PR | Phase 2 PR | Phase 3 PR | Current Phase | State |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| VertexAIDataLabelingJob | vertexai | [#7998](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7998) | - | - | 1 | MERGED |
| VertexAICustomJob | vertexai | [#7996](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7996) | - | - | 1 | OPEN |
| VertexAIDataset | vertexai | [#7991](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7991) | - | - | 1 | OPEN |
| VertexAIExampleStore | vertexai | [#7992](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7992) | - | - | 1 | OPEN |
| AIPlatformBatchPredictionJob | vertexai | [#7993](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7993) | - | - | 1 | OPEN |
| VertexAIEndpoint | vertexai | [#7994](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7994) | - | - | 1 | OPEN |
| VertexAIDeploymentResourcePool | vertexai | [#7995](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7995) | - | - | 1 | OPEN |
| VertexAICachedContent | vertexai | [#7997](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7997) | - | - | 1 | OPEN |
| ActionsPreview | actions | [#7989](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7989) | - | - | 1 | CLOSED |
| ActionsVersion | actions | [#7990](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7990) | - | - | 1 | CLOSED |
| NetworkConnectivityRegionalEndpoint | networkconnectivity | [#8071](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8071) | - | - | 1 | OPEN |
| CertificateManagerTrustConfig | certificatemanager | [#8072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8072) | - | - | 1 | OPEN |
| ArtifactRegistryVPCSCConfig | artifactregistry | [#8073](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8073) | - | - | 1 | OPEN |
| EventarcEnrollment | eventarc | [#8074](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8074) | - | - | 1 | OPEN |
| NotebookRuntime | notebooks | [#8075](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8075) | - | - | 1 | OPEN |
| KMSEKMConnection | kms | [#8076](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8076) | - | - | 1 | OPEN |
| CloudIdentityDevice | cloudidentity | [#8077](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8077) | - | - | 1 | OPEN |
| ApigeeApiProduct | apigee | [#8078](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8078) | - | - | 1 | OPEN |
| SecretManagerRegionalSecret | secretmanager | [#8080](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8080) | - | - | 1 | OPEN |
| BinaryAuthorizationPlatformPolicy | binaryauthorization | [#8081](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8081) | - | - | 1 | MERGED |
| VertexAIFeatureGroup | aiplatform | [#8386](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8386) | - | - | 1 | ISSUE_CREATED |
| VertexAIFeatureOnlineStore | aiplatform | [#8387](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8387) | - | - | 1 | ISSUE_CREATED |
| VertexAIHyperparameterTuningJob | aiplatform | [#8388](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8388) | - | - | 1 | ISSUE_CREATED |
| VertexAIIndex | aiplatform | [#8389](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8389) | - | - | 1 | ISSUE_CREATED |
| VertexAIIndexEndpoint | aiplatform | [#8390](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8390) | - | - | 1 | ISSUE_CREATED |
| VertexAIModelDeploymentMonitoringJob | aiplatform | [#8391](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8391) | - | - | 1 | ISSUE_CREATED |
| AIStreamsCluster | aistreams | [#8392](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8392) | - | - | 1 | ISSUE_CREATED |
| BigQueryAnalyticsHubDataExchange | analyticshub | [#8393](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8393) | - | - | 1 | ISSUE_CREATED |
| ApigeeRegistryApi | apigeeregistry | [#8394](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8394) | - | - | 1 | ISSUE_CREATED |
| ApiHubInstance | apihub | [#8395](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8395) | - | - | 1 | ISSUE_CREATED |
| ApiHubAttribute | apihub | [#8396](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8396) | - | - | 1 | ISSUE_CREATED |
| ApiHubCuration | apihub | [#8397](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8397) | - | - | 1 | ISSUE_CREATED |
| ApiHubDependency | apihub | [#8398](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8398) | - | - | 1 | ISSUE_CREATED |
| ApiHubDeployment | apihub | [#8399](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8399) | - | - | 1 | ISSUE_CREATED |
| AppHubServiceProjectAttachment | apphub | [#8400](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8400) | - | - | 1 | ISSUE_CREATED |
| AutoMLDataset | automl | [#8401](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8401) | - | - | 1 | ISSUE_CREATED |
| BatchResourceAllowance | batch | [#8402](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8402) | - | - | 1 | ISSUE_CREATED |
| BeyondCorpClientConnectorService | beyondcorp | [#8403](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8403) | - | - | 1 | ISSUE_CREATED |
| BeyondCorpClientGateway | beyondcorp | [#8404](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8404) | - | - | 1 | ISSUE_CREATED |
| BigLakeCatalog | biglake | [#8405](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8405) | - | - | 1 | ISSUE_CREATED |
