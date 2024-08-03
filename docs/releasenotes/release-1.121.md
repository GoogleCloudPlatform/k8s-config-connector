# v1.121.0

* Special shout-outs to @600lyy, @acpana, @barney-s, @gemmahou, @haiyanmeng, @hankfreund, @jasonvigil, @jiefenghe, @jingyih, @justinsb, @maqiuyujoyce, @ostrain, @xiaoweim, @yuwenma, @ziyue-101 for their contributions to this release.

## Announcement 

* The `state-into-spec` is default to `Absent` in any *new* ConfigController clusters. 
We plan to apply this default to *all* the ConfigController clusters in v1.122 and then
to all ConfigConnector clusters in v1.123.   

## Direct Cloud Reconciler:

* `DataformRepository` (v1alpha1)

## Fixes:

* BigTable
  * When autoscaling is enabled (`spec.cluster[].autoscalingConfig.`), do not use `numNodes` (`spec.cluster[].numNodes=2`) as that applies only to manual scaling.

* BigQueryConnection
  * Added `status.observedState` field to store the output-only fields which are previously mistakenly defined in `spec`.