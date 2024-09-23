# v1.123.1

* We do not recommend using 1.123.0 as it contains regression issues around the "state-into-spec" defaulting for resources `CloudIdentityGroup`, `CloudBuildTrigger` and `FirestoreIndex`.

* Special shout-outs to @600lyy ,@acpana ,@anhdle-sso ,@gemmahou ,@hankfreund ,@jasonvigil ,@jingyih ,@justinsb ,@maqiuyujoyce ,@nb-goog ,@xiaoweim ,@yuwenma ,@ziyue-101 for their contributions to this release.

## Announcement

* Starting from this version, all the new CRs (CustomResources) will have the `cnrm.cloud.google.com/state-into-spec`
  annotation defaulted to `absent`. This means Config Connector will not populate any unspecified fields into the
  `spec` after a successful reconciliation of the resource. The behavior of existing CRs will not be impacted. More
  details about the Absent behavior can be found
  [here](https://cloud.google.com/config-connector/docs/concepts/ignore-unspecified-fields#absent).

* You can use the `alpha.cnrm.cloud.google.com/reconciler: direct` label on `DataflowFlexTemplateJob` resource to opt-in
the Direct Cloud Reconciler, which provides an advanced status update solution for some timeout issues.


## Direct Cloud Reconciler:

* `BigQueryDataTransferConfig` (v1alpha1)
* `BigQueryConnectionConnection` (v1alpha1)

## Resources promoted from alpha to beta:

* `DataformRepository` is now a v1beta1 resource.

## New Resources:

* Added support for `FirestoreDatabase` (v1alpha1) resource.

## New Fields:

* BigQueryConnectionConnection
  * Added `spec.cloudSql`

## Bug Fixes:

* [fix: Enable TF-based reconciler for SQLInstances without clone source](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2731)
