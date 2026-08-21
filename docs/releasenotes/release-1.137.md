# Release 1.137.0

*   Special shout-outs to 600lyy, acpana, anhdle-sso, gemmahou, georgecma, justinsb, maqiuyujoyce, shavonz, xiaoweim, yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   `DocumentAIProcessorVersion`
*   `EssentialContactsContact`
*   `BigQueryBigLakeTable`
*   `BackupDRBackupPlan`

## New Alpha Resources (Direct Reconciler):

*   `BigtableMaterializedView`

## New Fields:

*   `BigtableMaterializedView`: Added `spec.sourceTableRef` and `spec.definition`.
*   `BackupDRBackupPlan`: Added `spec.backupConfig.retentionPeriodDays` and `spec.backupConfig.backupWindow`.
*   `MemorystoreInstance`: Added support for `MEMCACHE` and `REDIS` instance types.

## Reconciliation Improvements:

*   Enabled opt-in for IAM partial policy management.
*   Enabled server-side apply for KMS resources.
*   Improved reconciliation for `BigtableLogicalView` by using deep reflection.
*   Improved reconciliation for `FirestoreDatabase` with identity pattern and export support.
*   Improved reconciliation for `RunJob` with export support.
*   Unified `ComputeTargetTCPProxy` direct API and controller.

## New features:

*   Added a preview command to see the changes that will be applied by Config Connector.
*   Added resource reconciliation tracking and early termination for the preview manager.

## Bug Fixes:

*   Fixed an issue where `ComputeBackendService` backends were not sorted.
*   Fixed an issue where `CloudFunctionsFunction` runtime was not a supported value.
*   Fixed an issue with labels for `BackupDRBackupPlan`.
*   Fixed an issue with labels for `RunJob`.
*   Fixed a fuzzing issue for `FirestoreField`.
*   Fixed an issue with `KMSCryptoKey` import.
*   Fixed a flakiness issue in the `MonitoringDashboard` fuzzer.
*   Fixed a flakiness issue in tests.
*   Fixed an issue with bad labels in tests.
*   Fixed an issue with etag in direct reconciliation.
