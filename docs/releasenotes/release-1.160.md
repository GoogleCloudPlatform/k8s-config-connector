*   Special shout-outs to @acpana, @anhdle-sso, @barney-s, @eianiuk, @eugenenuke, @gemmahou, @ldanielmadariaga, @maqiuyujoyce, and @sdowell for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `CCInsightsAnalysisRule`
    *   Manage [Contact Center Insights analysis rules](https://cloud.google.com/contact-center/insights/docs) to define rules for analyzing conversational data.

*   `ConfigDeploymentGroup`
    *   Manage [Infrastructure Manager deployment groups](https://cloud.google.com/infrastructure-manager/docs) to orchestrate and manage deployments collectively.

*   `DataLineageProcess`
    *   Manage [Dataplex Data Lineage processes](https://cloud.google.com/dataplex/docs/data-lineage) to track data origin and movement.

*   `DataplexDataProduct`
    *   Manage [Dataplex data products](https://cloud.google.com/dataplex/docs) to package and share data assets across an enterprise.

*   `DialogflowKnowledgeBase`
    *   Manage [Dialogflow knowledge bases](https://cloud.google.com/dialogflow/docs) to provide FAQs and articles for agent answers.

*   `DiscoveryEngineSchema`
    *   Manage [Discovery Engine schemas](https://cloud.google.com/generative-ai-app-builder/docs) to define document structures and schemas for data stores.

*   `DiscoveryEngineSession`
    *   Manage [Discovery Engine sessions](https://cloud.google.com/generative-ai-app-builder/docs) to maintain conversation context and history for search and recommendations.

*   `DiscoveryEngineSitemap`
    *   Manage [Discovery Engine sitemaps](https://cloud.google.com/generative-ai-app-builder/docs) to crawl web pages for search data stores.

*   `ManagedKafkaConnectCluster`
    *   Manage [Apache Kafka Connect clusters](https://cloud.google.com/managed-kafka/docs) to run connectors that integrate Apache Kafka with Google Cloud and external systems.

*   `MigrationCenterPreferenceSet`
    *   Manage [Migration Center preference sets](https://cloud.google.com/migration-center/docs) to specify migration preferences and target cloud architecture assumptions.

*   `MigrationCenterSource`
    *   Manage [Migration Center sources](https://cloud.google.com/migration-center/docs) to configure data sources for importing asset data.

*   `NetAppActiveDirectory`
    *   Manage [NetApp Volumes Active Directory configurations](https://cloud.google.com/netapp/volumes/docs) to enable SMB and NFS domain-joined authentication.

*   `NetAppHostGroup`
    *   Manage [NetApp Volumes host groups](https://cloud.google.com/netapp/volumes/docs) to group client hosts for volume access.

*   `NetworkConnectivityTransport`
    *   Manage [Network Connectivity transports](https://cloud.google.com/network-connectivity/docs) for connecting network endpoints.

*   `NetworkSecurityMirroringEndpointGroupAssociation`
    *   Manage [Network Security mirroring endpoint group associations](https://cloud.google.com/vpc/docs/packet-mirroring) to associate packet mirroring endpoint groups with VPC networks.

*   `NetworkServicesLBEdgeExtension`
    *   Manage [Service Extensions load balancing edge extensions](https://cloud.google.com/service-extensions/docs) to customize Cloud CDN and edge proxy traffic processing.

*   `NetworkServicesLBTrafficExtension`
    *   Manage [Service Extensions load balancing traffic extensions](https://cloud.google.com/service-extensions/docs) to customize request and response processing in Application Load Balancers.

*   `NetworkServicesServiceLBPolicy`
    *   Manage [Service Extensions service load balancing policies](https://cloud.google.com/service-extensions/docs) to configure fine-grained traffic distribution across backend services.

*   `SaaSServiceMgmtUnit`
    *   Manage SaaS Service Management units to manage SaaS deployment units.

*   `StorageInsightsReportConfig`
    *   Manage [Cloud Storage Insights report configurations](https://cloud.google.com/storage/docs/insights/using-inventory-reports) to generate inventory and audit reports for Cloud Storage buckets.

*   `WorkloadManagerEvaluation`
    *   Manage [Workload Manager evaluations](https://cloud.google.com/workload-manager/docs) to scan and evaluate workloads against best practice rules.

## New Fields:

*   [`BigtableTable`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigtable/bigtabletable)
    *   Added support for `spec.columnFamily[].type`.

*   [`MemorystoreInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/memorystore/memorystoreinstance)
    *   Added support for `spec.gcsSource`.

*   [`MonitoringDashboard`](https://cloud.google.com/config-connector/docs/reference/resource-docs/monitoring/monitoringdashboard)
    *   Added support for `opsAnalyticsQuery` in `TimeSeriesQuery`.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   [`ComputeNetworkFirewallPolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computenetworkfirewallpolicy)
    *   Support direct reconciliation (opt-in).

*   [`ComputeProjectMetadata`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeprojectmetadata)
    *   Support direct reconciliation (opt-in).

*   [`FirebaseHostingSite`](https://cloud.google.com/config-connector/docs/reference/resource-docs/firebasehosting/firebasehostingsite)
    *   Support direct reconciliation (opt-in).

*   [`NetworkServicesEdgeCacheService`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networkservices/networkservicesedgecacheservice)
    *   Support direct reconciliation (opt-in).

*   [`OSLoginSSHPublicKey`](https://cloud.google.com/config-connector/docs/reference/resource-docs/oslogin/osloginsshpublickey)
    *   Support direct reconciliation (opt-in).

*   [`SQLSSLCert`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlsslcert)
    *   Support direct reconciliation (opt-in).

## New Features:

*   **Configurable Max Exponential Backoff Delay**: Support configurable max exponential backoff delay ceiling and retry halting for direct controllers via the `cnrm.cloud.google.com/backoff-max-delay-in-seconds` annotation.
*   **Halt Deletions in Paused Actuation Mode**: Halt delete operations when `cnrm.cloud.google.com/actuation-mode` is set to `Paused` on an individual resource.

## Bug Fixes:

*   [`AlloyDBCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/alloydb/alloydbcluster)
    *   Fixed false-positive diffs on `labels`, `initialUser`, and `network` during takeover and re-reconciliation in the direct controller.

*   [`ArtifactRegistryRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/artifactregistry/artifactregistryrepository)
    *   Fixed Terraform-to-Direct controller migration diffs on `cleanup_policies` and `cleanup_policy_dry_run`.

*   [`ComputeRouterNAT`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computerouternat)
    *   Aligned direct controller create defaults and delete behavior with the legacy controller.

*   [`ComputeSecurityPolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computesecuritypolicy)
    *   Fixed an issue where setting `rateLimitOptions.enforceOnKeyConfigs` failed with an API error unless `enforceOnKey: ""` was explicitly specified.

*   [`EssentialContactsContact`](https://cloud.google.com/config-connector/docs/reference/resource-docs/essentialcontacts/essentialcontactscontact)
    *   Supported automatic acquisition of existing contact by email on 409 ALREADY_EXISTS.

*   [`MemorystoreInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/memorystore/memorystoreinstance)
    *   Unmanaged `maintenanceVersion` if omitted from configuration YAML.
