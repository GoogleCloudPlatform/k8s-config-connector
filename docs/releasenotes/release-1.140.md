** This version is not yet released; this document is gathering release notes for the future release **

*   Special shout-outs to @600lyy, @acpana, @gemmahou, @justinsb, @katrielt, @maqiuyujoyce, @shavonz, @xiaoweim, and @yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   `CertificateManagerCertificateIssuanceConfig`
    *   Manage [Certificate Manager certificate issuance configurations](https://cloud.google.com/certificate-manager/docs/reference/public-ca/rest/v1/projects.locations.certificateIssuanceConfigs) for automating certificate issuance.  


## New Alpha Resources (Direct Reconciler):

*   `AssuredWorkloadsWorkload`
    *   Manage [Assured Workloads workloads](https://cloud.google.com/assured-workloads/docs/creating-workloads-programmatically) to support compliance and security requirements.
*   `ConfigDeliveryResourceBundle`
    *   Manage [Config Delivery resource bundles](https://cloud.google.com/anthos-config-management/docs/config-delivery/managing-resources) for Anthos Config Management.

## New Fields:

*   `AlloyDBCluster`
    *   Added `spec.restoreContinuousBackupSource` and `spec.restoreBackupSource` fields to support restoring from a backup.
*   `BigQueryReservationAssignment`
    *   Added `spec.jobType` field.
*   `FirestoreDatabase`
    *   Added `spec.deleteProtectionState` field.
*   `FirestoreField`
    *   Added `spec.ttlConfig` field.
*   `RunJob`
    *   Added `spec.template.template.containers.dependsOn` field.

## Reconciliation Improvements

*   Integrated Multi-Cluster Leader Election for improved reliability in multi-cluster setups.
*   Added mock GCP support for `BigtableGCPolicy`, `SourceRepo`, `MonitoringDashboard`, and `NetworkServices` gateways to improve testing.

## Bug Fixes:

*   Fixed an issue where `BigQueryReservationAssignment` was not exposing `externalRef`.
*   Fixed an issue with `CertificateManagerDNSAuthorization` API, Fuzzer and Mapper.
*   Fixed an issue with `FirestoreDatabase` defaulting logic.
