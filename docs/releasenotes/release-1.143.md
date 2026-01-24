# Release Notes 1.143.0

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @dhavalbhensdadiya-crest, @gemmahou, @himanikh, @justinsb, @kjasnoor0305, @maqiuyujoyce, @nancynh, @xiaoweim, @yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler)

*   [`ArtifactRegistryRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/artifactregistry/artifactregistryrepository)
*   [`LoggingLink`](https://cloud.google.com/config-connector/docs/reference/resource-docs/logging/logginglink)
*   [`MemorystoreInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/memorystore/memorystoreinstance)
*   [`PrivateCACAPool`](https://cloud.google.com/config-connector/docs/reference/resource-docs/privateca/privatecacapool)

## New Alpha Resources (Direct Reconciler)

*   `ParameterManagerParameter`

## New Fields

*   [`AlloyDBInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/alloydb/alloydbinstance)
    *   Added `spec.connectionPoolConfig` field.
    *   Added `status.connectionPoolConfig` field.

## Reconciliation Improvements

*   [`TagsTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagstagbinding)
    *   Added support for `organizations` in `parentRef`.
    *   Added support for multiple targets in `parentRef`.

*   Resource References (`refs.Ref`) support added for the following resources to improve reference resolution:
    *   `BigQueryTable`
    *   `BigQueryDataset`
    *   `CloudRunService`
    *   `CloudRunJob`
    *   `ArtifactRegistryRepository`
    *   `StorageBucket`

*   Added structured diff reporting for the following Direct Reconciler resources:
    *   `Workstation`
    *   `NotebookInstance`
    *   `BackupPolicy`
    *   `ManagedKafkaTopic`
    *   `ManagedKafkaCluster`
    *   `DataprocBatch`
    *   `ComposerEnvironment`
    *   `CloudBuildWorkerPool`
    *   `BigtableMaterializedView`
    *   `AppProfileIdentity`

## New features

*   Set `GOMEMLIMIT` for KCC workloads to improve memory management and stability.

## Bug Fixes

*   Fixed a race condition in `kccmanager`.
*   [Issue 6221](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6221): `ComputeBackendService` can now correctly refer to `clientTLSPolicy`.
*   [Issue 6156](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6156): `BigQueryTable` now ignores `int64` to `int32` schema changes when configured.
*   [Issue 6026](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6026): Fixed identity parsing for `TagsTagValue`.
