*   Special shout-outs to @acpana, @anfernee, @anhdle-sso, @barney-s, @eianiuk, @fkc1e100, @gemmahou, @ldanielmadariaga, @maqiuyujoyce, @mattiefu, and @sdowell for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `AIPlatformReasoningEngine`
    *   Manage [Vertex AI Reasoning Engines](https://cloud.google.com/vertex-ai/docs/reasoning-engine/overview) to deploy and run customized reasoning applications.

*   `CCInsightsAnalysisRule`
    *   Manage [Contact Center Insights analysis rules](https://cloud.google.com/contact-center/insights/docs) to configure conversation analysis criteria.

*   `CloudDeployCustomTargetType`
    *   Manage [Cloud Deploy custom target types](https://cloud.google.com/deploy/docs) to define custom deployment targets and rendering actions.

*   `ConfigDeliveryResourceBundle`
    *   Manage [Config Delivery resource bundles](https://cloud.google.com/kubernetes-engine/docs/concepts/fleet-deployment) for deploying declarative fleet packages.

*   `DataformFolder`
    *   Manage [Dataform folders](https://cloud.google.com/dataform/docs) in Dataform repositories to organize development assets.

*   `DataLineageProcess`
    *   Manage [Dataplex Data Lineage processes](https://cloud.google.com/dataplex/docs/data-lineage) to track data origin and movement.

*   `DataplexDataProduct`
    *   Manage [Dataplex data products](https://cloud.google.com/dataplex/docs) to govern and share data domain assets.

*   `DialogflowKnowledgeBase`
    *   Manage [Dialogflow knowledge bases](https://cloud.google.com/dialogflow/cx/docs) for conversational question-answering.

*   `DiscoveryEngineControl`
    *   Manage [Discovery Engine controls](https://cloud.google.com/generative-ai-app-builder/docs) to boost or filter search results.

*   `EventarcPipeline`
    *   Manage [Eventarc pipelines](https://cloud.google.com/eventarc/docs) to transform and route events across Google Cloud services.

*   `LiveStreamChannel`
    *   Manage [Live Stream channels](https://cloud.google.com/livestream/docs) to ingest, encode, and distribute live video streams.

*   `LiveStreamInput`
    *   Manage [Live Stream inputs](https://cloud.google.com/livestream/docs) to configure video ingestion endpoints.

*   `ManagedKafkaConnectCluster`
    *   Manage [Managed Service for Apache Kafka Connect clusters](https://cloud.google.com/managed-kafka/docs) for Apache Kafka Connect integrations.

*   `MapManagementMapConfig`
    *   Manage [Maps Management map configurations](https://developers.google.com/maps/documentation) to customize Google Maps features.

*   `MapManagementStyleConfig`
    *   Manage [Maps Management style configurations](https://developers.google.com/maps/documentation) to define custom map styles.

*   `NetworkConnectivityMulticloudDataTransferConfig`
    *   Manage [Network Connectivity multicloud data transfer configs](https://cloud.google.com/network-connectivity/docs) for data transfer between cloud providers.

*   `NetworkSecurityDNSThreatDetector`
    *   Manage [Network Security DNS threat detectors](https://cloud.google.com/secure-web-proxy/docs) to monitor and mitigate DNS-based threats.

*   `NetworkSecurityMirroringEndpointGroupAssociation`
    *   Manage [Network Security mirroring endpoint group associations](https://cloud.google.com/firewall/docs/packet-mirroring) to associate mirroring endpoint groups with VPC networks.

## New Fields:

*   [`ArtifactRegistryRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/artifactregistry/artifactregistryrepository)
    *   Added support for `spec.remoteRepositoryConfig.commonRepository`.

*   [`ComposerEnvironment`](https://cloud.google.com/config-connector/docs/reference/resource-docs/composer/composerenvironment)
    *   Added mutable update masks for Composer 3 workloads configuration (`spec.config.workloadsConfig.dagProcessor`).

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)
    *   Added support for `spec.privilegedAdmissionConfig` (Autopilot Privileged Admission).

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
    *   Added support for disaster recovery observed state fields under `status.observedState` (`failoverDRReplicaName`, `drReplica`, `psaWriteEndpoint`, `masterInstanceName`).

## New Features:

*   **Cloud SQL Advanced Disaster Recovery**: Added support for Cloud SQL Enterprise Plus Disaster Recovery (DR) pair setup and drift suppression for `SQLInstance` via the `cnrm.cloud.google.com/sqlinstance-advanced-dr: "enabled"` annotation.
*   **Operator ResourceSettings Pod Restart**: Automatically restart manager pods when `resourceSettings` change in operator configuration.
*   **GitOps Management Guide**: Added documentation and guide for managing Config Connector using GitOps tools (such as ArgoCD).

## Bug Fixes:

*   [`BigQueryDataset`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13352)
    *   Fixed false-positive diffs during reconciliation, normalized access entry diffing and cloning, and avoided redundant `PATCH` requests when resources are already up to date.

*   [`ComputeRouterNAT`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13472)
    *   Fixed false drift and steady-state reconciliation issues during Direct controller takeover.

*   [`MonitoringDashboard`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13574)
    *   Migrated direct controller to the Google REST client for improved enum compatibility and reliable reconciliation.
