# v1.121.0

* Special shout-outs to @600lyy, @acpana, @barney-s, @gemmahou, @haiyanmeng, @hankfreund, @jasonvigil, @jiefenghe, @jingyih, @justinsb, @maqiuyujoyce, @ostrain, @xiaoweim, @yuwenma, @ziyue-101 for their contributions to this release.

## Announcement 

* We plan to apply the `state-into-spec` default value  `Absent` to  *all the Config Connector clusters* in the  v1.123 (next to the next release).

## Direct Cloud Reconciler:

* `DataformRepository` (v1alpha1)

## Fixes:

* BigTable
  * When autoscaling is enabled (`spec.cluster[].autoscalingConfig.`), do not use `numNodes` (`spec.cluster[].numNodes=2`) as that applies only to manual scaling.

* BigQueryConnection
  * Added `status.observedState` field to store the output-only fields which are previously mistakenly defined in `spec`.
