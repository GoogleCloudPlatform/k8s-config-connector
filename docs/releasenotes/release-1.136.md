
** This version is not yet released; this document is gathering release notes for the future release **

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @gemmahou, @jingyih, @justinsb, @xiaoweim, @yuwenma for their contributions to this release.

## New Beta Resources

*   `AssetFeed`
*   `BigQueryReservationAssignment`
*   `CloudDeployDeliveryPipeline`
*   `ComposerEnvironment`
*   `NetworkSecurityAuthorizationPolicy`
*   `NetworkSecurityClientTlsPolicy`
*   `NetworkSecurityServerTlsPolicy`
*   `RunJob`

## New Fields

*   `ComposerEnvironment`:
    *   Added `spec.config.storageConfig` field.
    *   Added `spec.config.workloadsConfig.dagProcessor` field.
    *   Added `spec.config.workloadsConfig.triggerer` field.
    *   Added `spec.config.softwareConfig.webServerPluginsMode` field.
    *   Added `spec.config.softwareConfig.cloudDataLineageIntegration` field.

## Reconciliation Improvements

*   Introduce Stateful Reconciliation for Direct Controllers. The controller will store a hash of the last successfully applied `.spec` in the resource's `.status`. This provides a lightweight, GitOps-safe historical record, allowing the controller to know with certainty when a user has modified their desired state.

## New features

*   Added CEL annotation for immutable fields in IAM resources.
*   Added support for checking etag in spec.

## Bug Fixes

*   Fixed an issue where `CloudIdentityMembership` roles comparison would fail.
*   Fixed a bug where the wrong GVK was reported in IAM controller.
*   Fixed a bug where errors were swallowed when reading a Secret.
*   Fixed an issue with LRO endTime in mockgcp.
*   Fixed a bug in the etag mapper.
*   Fixed a bug in the mapper generator for slice and single object map.
*   Fixed a bug in the mapper generator for OneOf if the input is not proto.Message.
*   Fixed an import for refs in the same package in controllerbuilder.
