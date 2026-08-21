# v1.115.0

* Better support for AlloyDB secondary clusters and instances.

* Special shout-out to @199201shubhamsahu for their contributions to this release.

## New Fields:

* AlloyDBCluster
  * Added `spec.clusterType` field.
  * Added `spec.deletionPolicy` field.
  * Added `spec.secondaryConfig` field.

* AlloyDBInstance
  * Added `spec.instanceTypeRef` field.
