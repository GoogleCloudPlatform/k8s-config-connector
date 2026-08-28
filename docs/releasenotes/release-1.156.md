*   Special shout-outs to @NickASC4, @acpana, @ada-coder-bot, @anfernee, @anhdle-sso, @barney-s, @daedalus-agent-bot, @feynman-agent-bot, @gemmahou, @hopper-coder-bot, @ldanielmadariaga, @lovelace-coder-bot, @maqiuyujoyce, @neumann-coder-bot, @reviewbot-robot, @sdowell, and @walle-agent-bot for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `CCInsightsQAScorecard`
    *   Manage [Contact Center Insights QA scorecards](https://cloud.google.com/contact-center/insights/docs) to manage and evaluate agent performance.

*   `ContentWarehouseSynonymSet`
    *   Manage [Document AI Warehouse Synonym Sets](https://cloud.google.com/document-ai-warehouse/docs) to manage custom synonym groups for searches.

*   `DevConnectAccountConnector`
    *   Manage [Developer Connect Account Connectors](https://cloud.google.com/developer-connect/docs) to connect GKE clusters to developer systems.

*   `DiscoveryEngineEngine`
    *   Manage [Discovery Engine search engines](https://cloud.google.com/generative-ai-app-builder/docs) to power search experiences.

*   `DiscoveryEngineServingConfig`
    *   Manage [Discovery Engine serving configurations](https://cloud.google.com/generative-ai-app-builder/docs) to control search, recommendation, and listing features.

*   `GKEHubFleet`
    *   Manage [GKE Hub Fleets](https://cloud.google.com/anthos/multicluster-management/connect/docs) to logically group and manage clusters.

*   `ModelArmorTemplate`
    *   Manage [Model Armor templates](https://cloud.google.com/model-armor/docs) to define safety and security policies for large language models.

*   `NetworkSecurityAuthzPolicy`
    *   Manage [Network Security Authorization Policies](https://cloud.google.com/traffic-director/docs/security-authz-policy) to authorize traffic.

*   `RapidMigrationAssessmentCollector`
    *   Manage [Rapid Migration Assessment collectors](https://cloud.google.com/migration-center/docs/rma) to gather environment discovery data for cloud migration.

*   `SecurityCenterManagementEventThreatDetectionCustomModule`
    *   Manage [Security Command Center event threat detection custom modules](https://cloud.google.com/security-command-center/docs) to customize threat detection.

*   `StorageInsightsDatasetConfig`
    *   Manage [Cloud Storage Insights dataset configurations](https://cloud.google.com/storage/docs/insights) to automatically inventory and analyze storage datasets.

*   `VectorSearchCollection`
    *   Manage [Vertex AI Vector Search collections](https://cloud.google.com/vertex-ai/docs/vector-search) to manage similarity search indexes.

## New Fields:

*   [`BigtableTable`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigtable/bigtabletable)
    *   Added `spec.automatedBackupPolicy.locations` field.

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)
    *   Added `spec.nodeConfig.swapConfig` field.

*   [`ContainerNodePool`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containernodepool)
    *   Added `spec.nodeConfig.swapConfig` field.

*   [`DataprocCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dataproc/dataproccluster)
    *   Added `spec.secondaryWorkerConfig.instanceFlexibilityPolicy` field.

*   [`NetworkSecurityFirewallEndpoint`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networksecurity/networksecurityfirewallendpoint)
    *   Added `spec.organizationRef` field to support organizing resources at the organization level, and made `spec.projectRef` optional.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   `VertexAITensorboard`
    *   Support direct reconciliation (opt-in).

## New Features:

*   **Optional NAT IP Allocate Option**: Made `natIpAllocateOption` an optional field in `ComputeRouterNat` to support dynamic allocation.
*   **Preview Summary CLI Improvements**: Added namespace and current status to the preview summary report.

## Bug Fixes:

*   [`NetworkSecurityFirewallEndpoint`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networksecurity/networksecurityfirewallendpoint)
    *   Fixed reconciliation logic and verified behavior against real GCP.

*   [`NetworkSecurityFirewallEndpointAssociation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networksecurity/networksecurityfirewallendpointassociation)
    *   Fixed reconciliation logic and verified behavior against real GCP.

*   [`NetworkSecurityTLSInspectionPolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networksecurity/networksecuritytlsinspectionpolicy)
    *   Fixed reconciliation logic and verified behavior against real GCP.
