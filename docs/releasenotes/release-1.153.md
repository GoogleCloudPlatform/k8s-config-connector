*   Special shout-outs to @acpana, @anfernee, @anhdle-sso, @barney-s, @codebot-robot, @fqbright, @gemmahou, @himanigulati01, @iamkonohamaru, @justinsb, @katrielt, @ldanielmadariaga, @maqiuyujoyce, @yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`OrgPolicyPolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/orgpolicy/orgpolicypolicy)
    *   Manage [Organization Policies](https://cloud.google.com/resource-manager/docs/organization-policy/overview) to configure constraints across your Google Cloud resources.

## New Alpha Resources (Direct Reconciler):

*   `AIStreamsCluster`
    *   Manage [AI Streams](https://cloud.google.com/ai-streams/docs) clusters to ingest and process real-time video streams.

*   `AutoMLDataset`
    *   Manage [AutoML datasets](https://cloud.google.com/vertex-ai/docs/datasets/overview) in Vertex AI.

*   `CloudBatchResourceAllowance`
    *   Manage [Batch resource allowances](https://cloud.google.com/batch/docs) to control resource usage and limits for batch jobs.

*   `BeyondCorpClientGateway`
    *   Manage [BeyondCorp client gateways](https://cloud.google.com/beyondcorp/docs) to secure access to private applications.

*   `BigLakeDatabase`
    *   Manage [BigLake databases](https://cloud.google.com/bigquery/docs/biglake-intro) for unified analytics over data lakes.

*   `BigtableAuthorizedView`
    *   Manage [Bigtable authorized views](https://cloud.google.com/bigtable/docs/authorized-views) to control access to specific subsets of data in a table.

*   `BigtableBackup`
    *   Manage [Bigtable backups](https://cloud.google.com/bigtable/docs/backups) to preserve table data.

*   `BillingAccount`
    *   Manage [Billing accounts](https://cloud.google.com/billing/docs) (read-only/reference representation).

*   `CertificateManagerCertificateIssuanceConfig`
    *   Manage [Certificate Manager certificate issuance configurations](https://cloud.google.com/certificate-manager/docs/issuance-configs) for automated certificate provisioning.

*   `CloudDeployDeployPolicy`
    *   Manage [Cloud Deploy deploy policies](https://cloud.google.com/deploy/docs) to define deployment constraints and approvals.

*   `FirestoreBackupSchedule`
    *   Manage [Firestore backup schedules](https://cloud.google.com/firestore/docs/backups) for automated database backups.

*   `IAMDenyPolicy`
    *   Manage [IAM deny policies](https://cloud.google.com/iam/docs/deny-overview) to explicitly deny permissions.

*   `NetworkSecurityMirroringDeployment`
    *   Manage [Network Security mirroring deployments](https://cloud.google.com/secure-web-proxy/docs) to mirror network traffic for inspection.

*   `NetworkSecurityMirroringEndpointGroup`
    *   Manage [Network Security mirroring endpoint groups](https://cloud.google.com/secure-web-proxy/docs) to associate mirroring endpoints.

*   `SecurityCenterBigQueryExport`
    *   Manage [Security Command Center BigQuery exports](https://cloud.google.com/security-command-center/docs/how-to-analyze-findings-in-bigquery) to automatically export findings to BigQuery.

*   `SecurityCenterMuteConfig`
    *   Manage [Security Command Center mute configs](https://cloud.google.com/security-command-center/docs/how-to-mute-findings) to control finding visibility.

*   `VertexAIExampleStore`
    *   Manage [Vertex AI example stores](https://cloud.google.com/vertex-ai/docs) for managing example data.

## New Fields:

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)
    *   Added `spec.confidentialNodes.confidentialInstanceType` and `spec.nodeConfig.confidentialNodes.confidentialInstanceType` fields.

*   [`ContainerNodePool`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containernodepool)
    *   Added `spec.nodeConfig.confidentialNodes.confidentialInstanceType` field.
    *   Added `spec.nodeConfig.windowsNodeConfig` field.

*   [`RedisCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/redis/rediscluster)
    *   Added `spec.automatedBackupConfig` field.
    *   Added `spec.crossClusterReplicationConfig` field.
    *   Added `spec.kmsKeyRef` field.
    *   Added `spec.maintenancePolicy` and `status.observedState.maintenancePolicy` fields.
    *   Added `status.observedState.pscServiceAttachments` field.

*   [`MemorystoreInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/memorystore/memorystoreinstance)
    *   Added `spec.maintenancePolicy` and `status.observedState.maintenancePolicy` fields.
    *   Added `spec.kmsKeyRef` field.
    *   Added `status.observedState.encryptionInfo` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in behaviour. The API is unchanged. To use the direct reconciler, add the `cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object.

*   [`ArtifactRegistryRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/artifactregistry/artifactregistryrepository)
*   [`CertificateManagerCertificate`](https://cloud.google.com/config-connector/docs/reference/resource-docs/certificatemanager/certificatemanagercertificate)
*   [`DNSManagedZone`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dns/dnsmanagedzone)
*   [`DNSPolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dns/dnspolicies)
*   [`LoggingLogBucket`](https://cloud.google.com/config-connector/docs/reference/resource-docs/logging/logginglogbucket)
*   [`LoggingLogSink`](https://cloud.google.com/config-connector/docs/reference/resource-docs/logging/logginglogsink)
*   [`LoggingLogView`](https://cloud.google.com/config-connector/docs/reference/resource-docs/logging/logginglogview)
*   [`PubSubSchema`](https://cloud.google.com/config-connector/docs/reference/resource-docs/pubsub/pubsubschema)
*   [`PubSubTopic`](https://cloud.google.com/config-connector/docs/reference/resource-docs/pubsub/pubsubtopic)
*   [`ServiceDirectoryService`](https://cloud.google.com/config-connector/docs/reference/resource-docs/servicedirectory/servicedirectoryservice)

## Bug Fixes:

*   [ComposerEnvironment](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9114): Support resource reference resolution and fix drift on private environment config.
*   [ComputeBackendService](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/10414): Prevent the automatic injection of `subsetting: { policy: "NONE" }` for regional backend services with `INTERNAL_SELF_MANAGED` load balancing scheme.
*   [`DataprocCluster`](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7222): userServiceAccountMapping field in DataprocCluster resources is now mutable.
*   [DNSPolicy](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9594): Fixes to DNSPolicy for real GCP integration.
*   [MemorystoreInstance](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/10417): Fix crossInstanceReplicationConfig field, allowing replication roles to be synced and updated correctly.
*   [NetworkSecurity](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8828): Fix UpdateAuthorizationPolicy panic with empty updateMask.
*   [PrivilegedAccessManager](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7567): Fix mock drift in PrivilegedAccessManager by correcting Delete LRO and array element retention.
*   [RedisCluster](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9602): Fix permanent diff caused by automated backup configuration default value.
