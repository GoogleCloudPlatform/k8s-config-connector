* Special shout-outs to @acpana, @anhdle-sso, @cheftako, @ericpang777, @gemmahou, @haiyanmeng, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @nb-goog, @xiaoweim, @yuwenma, @ziyue-101 for their contributions to this release.

## Announcement 


## New Beta Resources (Direct Reconciler):

* [`BigQueryAnalyticsHubListing`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryanalyticshub/bigqueryanalyticshublisting)

* [`FirestoreDatabase`](https://cloud.google.com/config-connector/docs/reference/resource-docs/firestore/firestoredatabase)

* [`WorkstationConfig`](https://cloud.google.com/config-connector/docs/reference/resource-docs/workstations/workstationconfig)

* [`Workstation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/workstations/workstation)

## New Fields:

* `BigQueryDataTransferConfig`

  * Added `spec.scheduleOptionsV2` to customize the different types of data transfer schedule.
  * Added `status.observedState.error` with detailed information about reason of the latest config failure.

* `GKEHubFeatureMembership`

  * Added `spec.configmanagement.management` to enable Config Sync Auto Upgrade.


## Modified Beta Reconciliation

We added the direct controller support for the following 3 resources to enhance the reliability and performance. Add `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on the object to use the direct controller. CRD is backward compatible.

* `GKEHubFeatureMembership`

  * Added Config Sync Auto-Upgrade support.

* `SecretManagerSecret`

  * [#510](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/510) Enhanced `spec.rotation.nextRotationTime` to use a fixed datetime value to avoid relative `now()` friction. 
  * [#1081](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/1081) Fixed the `spec.replication.auto` immutable issue
  * [#3051](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/3051) Fixed the `spec.rotation.rotationPeriod` immutable issue 
  * Added the in-use version aliases in `status.observedState.versionAliases`
  * Resolved update stalling issues. 
  * Clarify the TTL use. See the problems and share your use in [#3395](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/3395) 

* `SecretManagerSecretVersion`
  
  * Resolved update stalling caused by `DependencyNotReady` errors.
  * Fixed the friction in `spec.enabled` that enabling/disabling a secret version does not always take effect in GCP.
  * **API Behavior Change**  The service generated ID is changed from `spec.resourceID` to `status.version` with `status.externalRef` (new field) to guardrail the identity. See the rational behind and share your feedback in [#3445](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/3445)
