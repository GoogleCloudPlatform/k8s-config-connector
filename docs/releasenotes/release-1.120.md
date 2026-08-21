# v1.120.1


* IAM configuration can now be applied to `PrivateCACAPool`, using our direct-actuation approach.

* You can configure the Config Connector operator to roll back to install the v1.119.0 CRDs by specifying `spec.version: 1.119.0` in the `Config ConnectorContext` CR (namespaced mode). 

* Special shout-outs to 600lyy,acpana,barney-s,coperni,gemmahou,hankfreund,jasonvigil,justinsb,maqiuyujoyce,nancynh,xiaoweim,yuwenma,zicongmei,ziyue-101 for their contributions to this release.

## Direct Cloud Reconciler:

* `CloudBuildWorkerPool`
* `MonitoringDashboard`

## Resources promoted from alpha to beta:

* `CloudBuildWorkerPool`
* `CloudIDSEndpoint`
* `ComputeMangedSSLCertificate`

## New Fields:

* `MonitoringAlertPolicy`
  * Added `spec.severity` field.

* `MonitoringDashboard`

  * Added `dashboardFilters` support.
  * Added `alertChart` widgets.
  * Added `collapsibleGroup` widgets.
  * Added `pieChart` widgets.
  * Added `sectionHeader` widgets.
  * Added `singleViewGroup` widgets.
  * Added `timeSeriesTable` widgets.

  * Added `blankView` to `scorecard` widgets.
  * Added `dataSets.targetAxis` and `y2Axis` fields to `xyChart` widgets.
  * Added `id` field to all widgets.
  * Added `prometheusQuery` and `outputFullDuration` to timeSeriesQuery.
  * Added `style` fields to text widgets.
  * Added `targetAxis` field to thresholds.

* `StorageBucket`
  * Added `spec.softDeletePolicy` field.
  * Added `status.observedState.softDeletePolicy` field.
