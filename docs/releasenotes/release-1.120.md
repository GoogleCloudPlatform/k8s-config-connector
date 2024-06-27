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

* `PlaceholderKind`

## New Resources:

* Added support for `PlaceholderKind` (v1beta1) resource.

## New Fields:

* `MonitoringAlertPolicy`
  * Added `spec.severity` field.

* `MonitoringDashboard`
  * Added `alertChart` widgets.
  * Added `collapsibleGroup` widgets.
  * Added `style` fields to text widgets.
  * Added `sectionHeader` widgets.
  * Added `pieChart` widgets.

* `StorageBucket`
  * Added `spec.softDeletePolicy` field.
  * Added `status.observedState.softDeletePolicy` field.

* PlaceholderKind
  * Added `spec.placeholder` field.

