*   Special shout-outs to @acpana, @anhdle-sso, @barney-s, @cheftako, @dhavalbhensdadiya-crest, @fkc1e100, @gemmahou, @justinsb, @katrielt, @maqiuyujoyce, and @xiaoweim for their contributions to this release.

## New Alpha Resources (Direct Reconciler)

*   `CloudDeployCustomTargetType`
    *   The `DeployCustomTargetType` (v1alpha1) resource is no longer supported and has been replaced by the new `CloudDeployCustomTargetType` (v1alpha1) resource. Please remove any instances of the `DeployCustomTargetType` resource.

## New Fields

*   [`CloudDeployTarget`](https://cloud.google.com/config-connector/docs/reference/resource-docs/clouddeploy/clouddeploytarget)
    *   Updated the `targetIds` field to be a resource reference.

## Reconciliation Improvements

*   Improved structured reporting diffs to provide better visibility into what changed during reconciliation for the following resources:
    *   `BigQueryDataset`
    *   `BigQueryReservationAssignment`
    *   `BigQueryTable`
    *   `CertificateManagerDNSAuthorization`
    *   `CloudIdentityGroup`
    *   `DataformRepository`
    *   `MetastoreService`
    *   `PrivilegedAccessManagerEntitlement`
    *   `WorkflowsWorkflow`

## New features

*   Enhanced `config-connector preview` to support side-by-side comparison between default and alternative controllers.
*   Added a `--skip-name-validation` flag to Config Connector controllers to bypass duplicate controller name checks during registration.

## Bug Fixes

*   Added CRD filtering for the preview recorder to skip non-CNRM objects.
*   Fixed a typo in `StorageDefaultObjectAccessControl` documentation.
*   Added missing reference documentation and samples for various resources.
