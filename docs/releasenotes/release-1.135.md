# Release 1.135.0

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @fqbright, @gemmahou, @GoogleCloudPlatform, @himanikh, @jingyih, @justinsb, @kjasnoor0305, @xiaoweim, @yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`APIQuotaPreference`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apikeys/apiquotapreference)
    *   Manage API Quota Preferences to customize quota consumption for Google Cloud services.
*   `VMwareEngineExternalAddress`
    *   Manage external IP addresses for VMware Engine workloads.
*   [`PubSubSnapshot`](https://cloud.google.com/config-connector/docs/reference/resource-docs/pubsub/pubsubsnapshot)
    *   Manage Pub/Sub snapshots to retain messages for a specific duration.

## New Alpha Resources (Direct Reconciler):

*   `IAMDenyPolicy`
    *   Manage IAM Deny Policies to deny access to Google Cloud resources.
*   `MultiClusterLease`
    *   A new resource for multi-cluster leader election.

## New Fields:

*   [`AlloyDBCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/alloydb/alloydbcluster)
    *   Added `spec.databaseVersion` field to specify the database version of the AlloyDB cluster.

## Reconciliation Improvements

*   `OrgPolicy`: Optimize Update by comparing policies to avoid unnecessary updates.

## New features:

*   CRDs: The `kind` field is now optional for well-known Project references in CRDs.
*   Added support for direct reconciliation events, providing more visibility into the reconciliation process.

## Bug Fixes:

*   `AlloyDB`: Fixed a nil pointer dereference in the AlloyDB direct controller.
*   `PubSub`: Fixed a proto mismatch issue in the API, mapper, and fuzzer for Pub/Sub resources.
