*   Special shout-outs to @acpana, @anfernee, @anhdle-sso, @barney-s, @cdmello-g, @eugenenuke, @gemmahou, @katrielt, @ldanielmadariaga, @maqiuyujoyce, @sdowell, and @YpNo for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `AgentRegistryBinding`
    *   Manage [Agent Registry bindings](https://cloud.google.com/agent-registry/docs) to connect agent tools and authentication configurations.

*   `AIPlatformPersistentResource`
    *   Manage [Vertex AI Persistent Resources](https://cloud.google.com/vertex-ai/docs/training/persistent-resource-overview) for cluster computing and training jobs.

*   `AssuredWorkloadsWorkload`
    *   Manage [Assured Workloads](https://cloud.google.com/assured-workloads/docs) to enforce security and compliance controls for regulated workloads.

*   `BigtableSchemaBundle`
    *   Manage [Cloud Bigtable schema bundles](https://cloud.google.com/bigtable/docs) for storing protobuf schemas in Bigtable.

*   `CCInsightsConversation`
    *   Manage [Contact Center Insights conversations](https://cloud.google.com/contact-center/insights/docs) to store and analyze customer interaction data.

*   `CloudBuildConnection`
    *   Manage [Cloud Build connections](https://cloud.google.com/build/docs) to connect to source code repositories like GitHub, GitHub Enterprise, and GitLab.

*   `CloudNumberRegistryRegistryBook`
    *   Manage [Cloud Number Registry registry books](https://cloud.google.com/telecom-data-fabric/docs) for phone number allocation and inventory management.

*   `DataplexAspectType`
    *   Manage [Dataplex aspect types](https://cloud.google.com/dataplex/docs) to define metadata schemas for Dataplex entries.

*   `DataplexDataTaxonomy`
    *   Manage [Dataplex data taxonomies](https://cloud.google.com/dataplex/docs) for classifying and managing data governance policies.

*   `DiscoveryEngineSampleQuerySet`
    *   Manage [Discovery Engine sample query sets](https://cloud.google.com/generative-ai-app-builder/docs) to evaluate search engine performance.

*   `EventarcMessageBus`
    *   Manage [Eventarc message buses](https://cloud.google.com/eventarc/docs) to route events from publishers to subscribers.

*   `ModelArmorFloorSetting`
    *   Manage [Model Armor floor settings](https://cloud.google.com/model-armor/docs) at project, folder, or organization scope to enforce baseline safety and security filters for LLMs.

*   `NetworkSecurityAddressGroup`
    *   Manage [Network Security address groups](https://cloud.google.com/vpc/docs/firewall-policies-address-groups) to group IP addresses and CIDR ranges for firewall policies.

*   `NetworkSecurityInterceptDeployment`
    *   Manage [Network Security intercept deployments](https://cloud.google.com/secure-web-proxy/docs) to deploy and route traffic through third-party security appliances.

*   `VMMigrationGroup`
    *   Manage [Migrate to Virtual Machines groups](https://cloud.google.com/migrate/virtual-machines/docs) to organize and migrate batches of virtual machines.

## New Fields:

*   [`BackupDRBackupPlanAssociation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/backupdr/backupdrbackupplanassociation)
    *   Added `spec.sqlInstanceRef` field to support backing up Cloud SQL instances.

*   [`ComputeNetworkFirewallPolicyRule`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computenetworkfirewallpolicyrule)
    *   Added `spec.match.srcNetworkContext` and `spec.match.destNetworkContext` fields.

*   [`ComputeRegionNetworkEndpointGroup`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeregionnetworkendpointgroup)
    *   Added `INTERNET_IP_PORT`, `INTERNET_FQDN_PORT`, and `GCE_VM_IP_PORTMAP` values for `spec.networkEndpointType`.

*   [`RedisCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/redis/rediscluster)
    *   Added `spec.clusterEndpoints` and `status.clusterEndpoints` fields.
    *   Added `spec.gcsSource` field.
    *   Added `spec.managedBackupSource` field.

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
    *   Added `spec.settings.databaseReplicationEnabled` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   `FilestoreInstance`
    *   Support direct reconciliation (opt-in).

## New Features:

*   **Resource-Level Pause Support**: Support pausing and unpausing actuation at the individual resource level using the annotation `cnrm.cloud.google.com/actuation-mode`.

## Bug Fixes:

*   [`BigQueryDataset`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12383)
    *   Fixed reconciliation failure in projects whose project ID contains a colon (such as domain-scoped projects `example.com:my-project`), where creation previously failed with an invalid ID format error and `resourceID` / `selfLink` were not populated.

*   [`ComposerEnvironment`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12795)
    *   Fixed infinite reconciliation loops, false drift, and HTTP 400 errors when using empty map fields (`envVariables: {}`, `labels: {}`), partial `workloadsConfig` blocks, or disabling `scheduledSnapshotsConfig`.

*   [`ComputeRouterNAT`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/12813)
    *   Supported both short name and canonical long path in `spec.routerRef` in the direct controller.

*   **CRD Structural Schema Alignment**:
    *   Removed empty `status.observedState` from `v1alpha1` CRDs to align with Kubernetes structural schema requirements.
