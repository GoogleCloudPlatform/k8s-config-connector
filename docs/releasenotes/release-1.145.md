# Release 1.145.0

## This version has not yet been released, this document is to prepare release notes for the upcoming release

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @dhavalbhensdadiya-crest, @eugenenuke, @gemmahou, @iamkonohamaru, @justinsb, @katrielt, @maqiuyujoyce, @xiaoweim, @yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`ComputeSecurityPolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computesecuritypolicy)
    *   Manage Google Cloud Armor security policies.

## New Fields:

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)
    *   Added `spec.clusterAutoscaling.defaultComputeClassConfig` field.

*   [`RunJob`](https://cloud.google.com/config-connector/docs/reference/resource-docs/run/runjob)
    *   Added `spec.template.template.volumes.nfs` field to support NFS backed Volumes.
    *   Added `spec.template.template.volumes.gcs` field to support GCS backed Volumes.

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
    *   Added `spec.settings.failoverDrReplicaRef` field to support designating CloudSQL Enterprise Plus DR Replicas.

## Reconciliation Improvements

We have added support for structured diff reporting to the following direct controllers to improve logging and debugging:

*   [`ApigeeEnvgroup`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apigee/apigeeenvgroup)
*   [`ApigeeInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apigee/apigeeinstance)
*   [`AssetFeed`](https://cloud.google.com/config-connector/docs/reference/resource-docs/asset/assetfeed)
*   [`AssetSavedQuery`](https://cloud.google.com/config-connector/docs/reference/resource-docs/asset/assetsavedquery)
*   [`BackupVaultIdentity`](https://cloud.google.com/config-connector/docs/reference/resource-docs/backupdr/backupvaultidentity)
*   [`BigtableLogicalView`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigtable/bigtablelogicalview)
*   [`CloudDeployDeliveryPipeline`](https://cloud.google.com/config-connector/docs/reference/resource-docs/clouddeploy/clouddeploydeliverypipeline)
*   [`ColabRuntime`](https://cloud.google.com/config-connector/docs/reference/resource-docs/colab/colabruntime)
*   [`RedisCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/redis/rediscluster)
*   [`SpannerBackupSchedule`](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerbackupschedule)
*   [`SpannerInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance)
*   [`SpannerInstanceConfig`](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstanceconfig)
*   [`TaskQueue`](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudtasks/taskqueue)
*   [`WorkstationConfig`](https://cloud.google.com/config-connector/docs/reference/resource-docs/workstations/workstationconfig)

## New features:

*   Added the `preview` command to the `config-connector` CLI. The `preview` command has been removed from the experimental `kompanion` tool.

## New Beta Resources (Direct Reconciler):

* MemorystoreInstance to manage [Memorystore for Valkey Instances](https://cloud.google.com//memorystore/docs/valkey/reference/rest/v1/projects.locations.instances)

## Bug Fixes:

*   [`DataformRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dataform/dataformrepository)
    *   Fixed a bug where the `serviceAccountRef` field could not be updated.

*   [`SpannerBackupSchedule`](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerbackupschedule)
    *   Fixed an issue with invalid update masks by handling output-only fields.