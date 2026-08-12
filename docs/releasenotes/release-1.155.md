*   Special shout-outs to @MKand, @abhishekbin, @acpana, @ada-coder-bot, @anfernee, @anhdle-sso, @barney-s, @codebot-robot, @daedalus-agent-bot, @feynman-agent-bot, @gemmahou, @hopper-coder-bot, @ldanielmadariaga, @lntutor, @lovelace-coder-bot, @maqiuyujoyce, @neumann-coder-bot, @reviewbot-robot, @robertcodebot, @suwandim, and @walle-agent-bot for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `DiscoveryEngineLicenseConfig`
    *   Manage [Discovery Engine license configurations](https://cloud.google.com/generative-ai-app-builder/docs) to manage application licenses.

*   `StorageManagedFolder`
    *   Manage [GCS managed folders](https://cloud.google.com/storage/docs/managed-folders) to apply granular access control policies to subsets of storage objects.

*   `VertexAITensorboardExperiment`
    *   Manage [Vertex AI Tensorboard experiments](https://cloud.google.com/vertex-ai/docs/tensorboard) to organize and track runs.

## New Fields:

*   [`BigtableTable`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigtable/bigtabletable)
    *   Added `spec.automatedBackupPolicy` field.

*   [`CertificateManagerDNSAuthorization`](https://cloud.google.com/config-connector/docs/reference/resource-docs/certificatemanager/certificatemanagerdnsauthorization)
    *   Added `spec.type` field.

*   [`ComputeForwardingRule`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeforwardingrule)
    *   Added `spec.target.redisClusterServiceAttachment` field.

*   [`ComputeURLMap`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeurlmap)
    *   Added `spec.tests[].expectedOutputURL` field.
    *   Added `spec.tests[].expectedRedirectResponseCode` field.

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)
    *   Added `spec.nodeConfig.kubeletConfig.imageGcLowThresholdPercent` field.
    *   Added `spec.nodeConfig.kubeletConfig.imageGcHighThresholdPercent` field.
    *   Added `spec.nodeConfig.kubeletConfig.imageMinimumGcAge` field.
    *   Added `spec.nodeConfig.kubeletConfig.imageMaximumGcAge` field.
    *   Added `spec.nodeConfig.containerdConfig` field.
    *   Added `spec.inTransitEncryptionConfig` field.
    *   Added `spec.disableL4LbFirewallReconciliation` field.
    *   Added `spec.nodeConfig.resourceManagerTags` field.
    *   Added `spec.nodePoolAutoConfig.resourceManagerTags` field.

*   [`ContainerNodePool`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containernodepool)
    *   Added `spec.nodeConfig.kubeletConfig.imageGcLowThresholdPercent` field.
    *   Added `spec.nodeConfig.kubeletConfig.imageGcHighThresholdPercent` field.
    *   Added `spec.nodeConfig.kubeletConfig.imageMinimumGcAge` field.
    *   Added `spec.nodeConfig.kubeletConfig.imageMaximumGcAge` field.
    *   Added `spec.nodeConfig.containerdConfig` field.
    *   Added `spec.nodeConfig.resourceManagerTags` field.

*   [`StorageBucket`](https://cloud.google.com/config-connector/docs/reference/resource-docs/storage/storagebucket)
    *   Added `spec.autoclass.terminalStorageClass` field.
    *   Added `status.observedState.storageClass` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   [`NetworkServicesHTTPRoute`](https://cloud.google.com/config-connector/docs/reference/resource-docs/networkservices/networkserviceshttproute)
    *   Support direct reconciliation (opt-in).

## New Features:

*   **Configurable metrics server address**: Made the manager's built-in metrics server bind address configurable.
*   **Brownfield state comparison**: Added a generic helper function to compare desired and actual states in brownfield resources, improving reconciliation reliability.
*   **Irregular shortname pluralization**: Added support for irregular shortname pluralization of "corpus" to "corpora".

## Bug Fixes:

*   [`ComputeReservation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computereservation)
    *   Ignore diff for `specificReservation.inUseCount` to prevent infinite/unwanted reconciliations.

*   [`RedisInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/redis/redisinstance)
    *   Marked `MaintenanceSchedule` field as output only to align with GCP's behavior.

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
    *   Fixed legacy fuzzer roundtrip mismatch for `PscAutoConnectionPolicyEnabled`.

*   [`CloudFunctions2Function`](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudfunctions/cloudfunctions2function)
    *   Declared source fields mutable-but-unreadable to avoid spurious diffs.

*   [`Export` tool]
    *   Allowed exporting resource sets containing nested resource references.

*   [`setup-envtest` tool]
    *   Added check to verify that `KUBEBUILDER_ASSETS` directory exists.

*   [`GcsObject` resource]
    *   Removed `GCSObject` identity and reference.
