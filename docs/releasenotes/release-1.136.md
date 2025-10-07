
** This version is not yet released; this document is gathering release notes for the future release **

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @gemmahou, @jingyih, @justinsb, @xiaoweim, @yuwenma for their contributions to this release.

## New Beta Resources

*   `AssetFeed`
*   `BigQueryReservationAssignment`
*   `CloudDeployDeliveryPipeline`
*   `ComposerEnvironment`

## New Fields

*   Added `spec.storageConfig` field to `ComposerEnvironment`.
*   Added `spec.config.workloadsConfig.dagProcessor` field to `ComposerEnvironment`.
*   Added `spec.config.workloadsConfig.triggerer` field to `ComposerEnvironment`.
*   Added `spec.config.softwareConfig.webServerPluginsMode` field to `ComposerEnvironment`.
*   Added `spec.config.softwareConfig.cloudDataLineageIntegration` field to `ComposerEnvironment`.

## Reconciliation Improvements

*   Introduce [Stateful Reconciliation for Direct Controllers](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/docs/designs/stateful-reconciliation-with-cookie.md). The controller will store a hash of the last successfully applied `.spec` in the resource's `.status`. This provides a lightweight, GitOps-safe historical record, allowing the controller to know with certainty when a user has modified their desired state. This enhanced reconciliation will be applied to the KRM-native controller gradually

## Bug Fixes

*   Added support for checking etag in spec for alpha resources.
*   Fixed an issue where `CloudIdentityMembership` roles comparison would fail.
*   Fixed a bug where the wrong GVK was reported in IAM controller.
*   Fixed a bug where errors were swallowed when reading a Secret.
*   Fixed an issue with LRO endTime in mockgcp.
*   Fixed a bug in the etag mapper where the `etag` field was missing from the `BackupDRBackupVaultObservedState` struct, causing issues with roundtrip conversions in tests.
*   Fixed a bug in the mapper generator to handle cases where a proto field is a slice and the corresponding KRM field is a single object (and vice-versa), preventing panics.
*   Fixed a bug in the mapper generator to correctly handle `OneOf` fields that are not of type `proto.Message` by generating helper functions.
*   Fixed an import for refs in the same package in controllerbuilder.
