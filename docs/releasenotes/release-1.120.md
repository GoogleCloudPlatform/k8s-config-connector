# v1.120.0

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* IAM configuration can now be applied to `PrivateCACAPool`, using our direct-actuation approach.

* Special shout-outs to ... for their
  contributions to this release.
TODO: list contributors with `git log v1.120.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`

## Resources promoted from alpha to beta:

*When resources are promoted from alpha to beta, we (generally) ensure they follow our best practices: use of refs on fields where appropriate,
output fields from GCP APIs are in `status.observedState.*`

* `CloudIDSEndpoint`
* `ComputeMangedSSLCertificate`

## New Resources:

* Added support for `PlaceholderKind` (v1beta1) resource.

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

  * Added `id` field to all widgets.
  * Added `prometheusQuery` and `outputFullDuration` to timeSeriesQuery.
  * Added `style` fields to text widgets.
  * Added `targetAxis` field to thresholds.

* `StorageBucket`
  * Added `spec.softDeletePolicy` field.
  * Added `status.observedState.softDeletePolicy` field.

* PlaceholderKind
  * Added `spec.placeholder` field.

