# v1.116.0

* This release includes enhanced support for DNSRecordSet,
  enabling advanced configurations such as geo-routing, primary/backup,
  and weighted round-robin load-balancing.

## New Fields:

* ContainerCluster
  * Added `spec.nodeConfig.linuxNodeConfig.cgroupMode` field.

* ContainerNodePool
  * Added `spec.nodeConfig.linuxNodeConfig.cgroupMode` field.

* DNSRecordSet
  * Added `spec.routingPolicy.geo.healthCheckedTargets` field.
  * Added `spec.routingPolicy.primaryBackup` field.
  * Added `spec.routingPolicy.wrr` field.

* EventArcTrigger
  * Added `spec.destination.httpEndpoint` field.
  * Added `spec.destination.networkConfig` field.

* LoggingLogBucket
  * Added `spec.enableAnalytics` field.
