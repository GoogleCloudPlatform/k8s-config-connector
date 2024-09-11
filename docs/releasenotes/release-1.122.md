# v1.122.0

* Special shout-outs to @600lyy, @acpana, @anhdle-sso, @barney-s, @CyberHippo, @gemmahou, @haiyanmeng, @hankfreund, @himanikh, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @marko7460, @xiaoweim, @yuwenma, @ziyue-101 for their contributions to this release.

## Direct Cloud Reconciler:

* `RedisCluster` (v1alpha1)
* `SQLInstance`

## New Resources:

* Added support for `RedisCluster` (v1alpha1) resource.

## New Fields:

* `ContainerCluster`
  * The `spec.nodeConfig.taint` can be updated.

* `ContainerNodePool`
  * The `spec.nodeConfig.taint` can be updated.

* `SQLInstance`
  * Add the `spec.cloneSource`.

* `RunJob`
  * Add the `spec.template.template.volumes[].cloudSqlInstance`


