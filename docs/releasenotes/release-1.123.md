# v1.123.0

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* Special shout-outs to ... for their
  contributions to this release.
TODO: list contributors with `git log v1.122.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`

## Announcement

* Starting from this version, all the new CRs (CustomResources) will have the `cnrm.cloud.google.com/state-into-spec`
  annotation defaulted to `absent`. This means Config Connector will not populate any unspecified fields into the
  `spec` after a successful reconciliation of the resource. The behavior of existing CRs will not be impacted. More
  details about the Absent behavior can be found
  [here](https://cloud.google.com/config-connector/docs/concepts/ignore-unspecified-fields#absent).


## Direct Cloud Reconciler:

* `BigQueryDataTransferConfig` (v1alpha1)

## Resources promoted from alpha to beta:

*When resources are promoted from alpha to beta, we (generally) ensure they follow our best practices: use of refs on fields where appropriate,
output fields from GCP APIs are in `status.observedState.*`

* `DataformRepository` is now a v1beta1 resource.

## New Resources:

* Added support for `PlaceholderKind` (v1beta1) resource.

## New Fields:

* PlaceholderKind
  * Added `spec.placeholder` field.

## Bug Fixes:

* [fix: Enable TF-based reconciler for SQLInstances without clone source](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2731)
