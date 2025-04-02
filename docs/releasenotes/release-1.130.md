* Special shout-outs to acpana@, anhdle-sso@, barney-s@, cheftako@, ericpang777@, gemmahou@, jasonvigil@, jingyih@, justinsb@, maqiuyujoyce@, renovate-bot@, xiaoweim@, yufan-su@, yuwenma@ for their contributions to this release.

## Announcement    

## New Beta Resources (Direct Reconciler):

* [`ManagedKafkaTopic`](https://cloud.google.com/config-connector/docs/reference/resource-docs/managedkafka/managedkafkatopic)
* [`SecureSourceManagerInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/securesourcemanager/securesourcemanagerinstance.md)
* [`SecureSourceManagerRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/securesourcemanager/securesourcemanagerrepository.md)

## New Alpha Resources (Direct Reconciler):

* `DataplexLake`
* `DocumentAIProcessor`
* `CodeDeployDeliveryPipeline`
* `NetAppBackupPolicy`
* `BigLakeTable`
* `NotebooksEnvironment`
* `SpannerInstanceConfig`
* `VertexAIFeaturestore`
* `WorkflowExecution`
* `ApphubApplication`
* `BatchJob`
* `TaskQueue`
 
## New Fields:

* [`GKEHubFeatureMembership`](https://cloud.google.com/config-connector/docs/reference/resource-docs/gkehub/gkehubfeaturemembership)

  * Added `spec.configmanagement.configSync.stopSyncing` field since 1.129.

* [SpannerInstance](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance). For opt-in direct controller,
  * Added `spec.labels` field.
  * Added `spec.defaultBackupScheduleType` field.
* [SecretManagerSecret](https://cloud.google.com/config-connector/docs/reference/resource-docs/secretmanager/secretmanagersecret). For opt-in direct controller,
  * Added `spec.labels` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in behaviour.  The API is unchanged. To use the direct reconciler, add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object.  The following resources now have direct reconciliation support (and we list some of the issues that this fixes):

## Bug Fixes:

* FirewallPolicyRule cannot refer to FirewallPolicy with resourceID "firewallPolicies/firewallPolicyID" after the direct controller migration. The now accepted formats are "firewallPolicies/firewallPolicyID" and "firewallPolicyID".
