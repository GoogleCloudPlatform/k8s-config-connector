# Release Notes 1.150.0

** This version is not yet released; this document is gathering release notes
for the future release **

*   Special shout-outs to acpana, anhdle-sso, barney-s, cheftako, fedebongio, fkc1e100, gemmahou, justinsb, katrielt, ldanielmadariaga, maqiuyujoyce, suwandim, and xiaoweim for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`CloudDeployTarget`](https://cloud.google.com/config-connector/docs/reference/resource-docs/clouddeploy/clouddeploytarget)

    *   Manage [Cloud Deploy targets](https://cloud.google.com/deploy/docs/targets) which define where your application is deployed.

## New Alpha Resources (Direct Reconciler):

*   `GKEHubScope`

    *   Manage [GKE Hub Scopes](https://cloud.google.com/anthos/fleet-management/docs/concepts#fleet-level-features) which allow you to group fleet resources for fine-grained management.

## New Fields:

*   [`CertificateManagerCertificate`](https://cloud.google.com/config-connector/docs/reference/resource-docs/certificatemanager/certificatemanagercertificate)

    *   Added `status.observedState` field, which includes `managed.state`, `managed.authorizationAttemptInfo`, and `managed.provisioningIssue`. This allows tracking the current state of the certificate even when `state-into-spec: absent` is used.

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)

    *   Added `spec.ipAllocationPolicy.additionalIpRangesConfigs` field.

*   [`ContainerNodePool`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containernodepool)

    *   Added `spec.networkConfig.subnetworkRef` field.

## New Features:

*   **License Agent**

    *   Introduced a new license agent to help manage and automate license header checks and updates across the codebase.

## General Improvements:

*   **Enhanced Resource Creation Logging**

    *   Improved the logging and event reporting during resource creation for both Direct and DCL-based controllers by including structured diffs.

*   **Preview Tool Enhancements**

    *   The `preview` tool now supports read-only POST requests, improving its compatibility with certain GCP API methods.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support:

*   [`BigqueryTable`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigquery/bigquerytable)

    *   Fixed a permanent difference in the `policyTag` field when using the direct controller, ensuring safer upgrades.

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)

    *   Added detailed diff reporting for the `userLabels` field.

*   [`DataplexLake`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dataplex/dataplexlake)

    *   Added structured diff reporting to improve visibility into resource changes.

## Bug Fixes:

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)

    *   Fixed a permanent difference in the `databaseEncryption.state` field and added support for the `ALL_OBJECTS_ENCRYPTION_ENABLED` value.

*   [`MemorystoreInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/memorystore/memorystoreinstance)

    *   Updated the controller to use change cookies, improving reconciliation stability and correctness.
