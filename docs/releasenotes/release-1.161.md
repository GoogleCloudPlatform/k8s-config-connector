*   Special shout-outs to @KengoA, @acpana, @anfernee, @anhdle-sso, @barney-s, @eianiuk, @fkc1e100, @gemmahou, @himanigulati01, @ldanielmadariaga, @maqiuyujoyce, @mattiefu, @mutianf, @sdowell, and @suonenjoki for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `AIPlatformCustomJob`
    *   Manage [Vertex AI custom jobs](https://cloud.google.com/vertex-ai/docs/training/create-custom-jobs) to run custom machine learning training workloads.

*   `AIPlatformFeatureView`
    *   Manage [Vertex AI Feature Store feature views](https://cloud.google.com/vertex-ai/docs/featurestore/latest/overview) to define a read-only projection of features.

*   `BigQueryAnalyticsHubQueryTemplate`
    *   Manage [BigQuery Analytics Hub query templates](https://cloud.google.com/bigquery/docs/analytics-hub-introduction) to share reusable SQL queries across data exchanges.

*   `ComposerUserWorkloadsConfigMap`
    *   Manage [Cloud Composer user workloads ConfigMaps](https://cloud.google.com/composer/docs/composer-2/manage-environments) to supply configuration data to Airflow tasks.

*   `ComposerUserWorkloadsSecret`
    *   Manage [Cloud Composer user workloads Secrets](https://cloud.google.com/composer/docs/composer-2/manage-environments) to pass sensitive data to Airflow tasks.

*   `ConfigDeliveryRelease`
    *   Manage [Config Delivery releases](https://cloud.google.com/config-delivery/docs) to represent a specific version of configuration for deployment.

*   `DataplexEntry`
    *   Manage [Dataplex entries](https://cloud.google.com/dataplex/docs) to represent cataloged metadata items in Dataplex Universal Catalog.

*   `LoggingLogScope`
    *   Manage [Cloud Logging log scopes](https://cloud.google.com/logging/docs) to define resource boundaries for log routing and analytics.

*   `ManagedKafkaACL`
    *   Manage [Managed Service for Apache Kafka ACLs](https://cloud.google.com/managed-kafka/docs) to control access permissions for Apache Kafka topics and consumer groups.

*   `SecureSourceManagerBranchRule`
    *   Manage [Secure Source Manager branch rules](https://cloud.google.com/secure-source-manager/docs) to enforce branch protection policies.

*   `VectorsearchIndex`
    *   Manage [Vertex AI Vector Search indexes](https://cloud.google.com/vertex-ai/docs/vector-search/overview) to serve large-scale vector similarity search.

## New Fields:

*   [`BigQueryReservationReservation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryreservation/bigqueryreservationreservation)
    *   Added support for `spec.schedulingPolicy` field.
    *   Added support for `spec.reservationGroupRef` field.

*   [`DiscoveryEngineDataStore`](https://cloud.google.com/config-connector/docs/reference/resource-docs/discoveryengine/discoveryenginedatastore)
    *   Added support for `spec.skipDefaultSchemaCreation` field.

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
    *   Added `spec.settings.ipConfiguration.serverCaMode`, `serverCaPoolRef`, and `customSubjectAlternativeNames` fields.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   [`ContainerNodePool`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containernodepool)
    *   Support direct reconciliation (opt-in).

The following resources now use the direct reconciler by default:

*   [`APIKeysKey`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apikeys/apikeyskey)
*   [`BigQueryReservationCapacityCommitment`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryreservation/bigqueryreservationcapacitycommitment)
*   [`ComputeAutoscaler`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeautoscaler)
*   [`ComputeBackendServiceSignedURLKey`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computebackendservicesignedurlkey)
*   [`ComputeDiskResourcePolicyAttachment`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computediskresourcepolicyattachment)
*   [`DNSResponsePolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dns/dnsresponsepolicy)
*   `FirebaseHostingSite`
*   `OSLoginSSHPublicKey`

## New Features:

*   **Multi-Architecture Container Builds**: Support multi-architecture (AMD64 and ARM64) container builds.
*   **Go Toolchain Update**: Config Connector images are now built with Go 1.26.10 (operator: Go 1.27.3), which distrusts certificate chains through compromised eMudhra/emSign intermediate CAs ([golang/go#82345](https://github.com/golang/go/issues/82345)).
*   **Per-Resource Max Concurrent Reconciles**: Added experimental `spec.experiments.maxConcurrentReconcilesOverrides` field to `ConfigConnector` and `ConfigConnectorContext` to configure per-resource `MaxConcurrentReconciles`.
*   **IAM Policy Member Support for LoggingLogView**: Support `IAMPolicyMember` for `LoggingLogView`.

## Bug Fixes:

*   [`ComputeManagedSSLCertificate`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computemanagedsslcertificate)
    *   Support `status.observedState.managed.status` and `status.observedState.managed.domainStatus` fields in direct controller.

*   [`ComputeReservation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computereservation)
    *   Fixed an issue where the `ComputeReservation` direct controller failed reconciliation when optional immutable fields (such as `description` or `minCpuPlatform`) were set in GCP but omitted from the CR `spec`.

*   [`ComputeURLMap`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeurlmap)
    *   Allow switching between mutually exclusive fields (`defaultService`, `defaultUrlRedirect`, `defaultRouteAction`) on `ComputeURLMap`.

*   [`StorageAnywhereCache`](https://cloud.google.com/config-connector/docs/reference/resource-docs/storage/storageanywherecache)
    *   `admissionPolicy: admit-on-second-miss` is now treated as `admit-on-first-miss` to match the Cloud Storage API, and no longer causes repeated updates. `desiredState: paused` is deprecated; the resource now reports a clear error instead of calling the removed pause API.

*   Fixed an issue where `BigQueryAnalyticsHubDataExchange`, `VertexAIDataset`, and `VertexAIEndpoint` CRDs were excluded from release bundles.
