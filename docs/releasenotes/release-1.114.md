# v1.114.1

(Version 1.114.0 contained a minor regression discovered after tagging, and was not published; we recommend 1.114.1 instead)

* Stop merging sensitive fields in SQLInstance and ComputeBackendService

* Fix resource deletion of `AlloyDBInstance` and `EdgeContainerNodePool` when their "parent objects" no longer exist.

* Initial support (alpha stability) for pausing actuation of resources onto Google Cloud. Operators
  can set Config Connector's or Config ConnectorContext's spec.actuationMode to `Paused`, depending
  on whether to pause on the whole cluster or just a namespace.
  See the [pause feature docs](./../features/pause.md) for more information.

* Initial support (alpha stability) for defaulting state-into-spec to absent (the recommended setting),
  by setting `spec.stateIntoSpec: Absent` in the Config ConnectorContext.

* Experimental "powertools" area of the CLI, containing experimental/dangerous functionality that should not be
  part of normal operation, but can sometimes nonetheless be useful.

* Special shout-outs to Hamzawy63@, hkundag@, katrielt@ for their
  contributions to this release.

## Resources promoted from alpha to beta:

*When resources are promoted from alpha to beta, we (generally) ensure they follow our best practices: use of refs on fields where appropriate,
output fields from GCP APIs are in `status.observedState`.*

* `AccessContextManagerServicePerimeterResource`

## New Resources:

* Added support for `ComputeNetworkFirewallPolicyAssociation` (v1beta1) resource.

* Added support for `APIKeysKey` (v1alpha1) resource.

## New Fields:

* BigQueryDataSet
  * Added `access[].iamMember` field.

* ComputeAddress
  * Added `status.observedState.address` field.

* ComputeTargetHttpsProxy
  * Added `spec.certificateManagerCertificates` field.

* DNSRecordSet
  * Added `spec.routingPolicy` field.

* GKEHubFeatureMembership
  * Added `spec.policycontroller` field.
