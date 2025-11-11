# Release 1.137.0

*   Special shout-outs to acpana, anhdle-sso, gemmahou, justinsb, maqiuyujoyce, xiaoweim, yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   `BackupDRBackupVault`
*   `OrgPolicyCustomConstraint`

## New Alpha Resources (Direct Reconciler):

*   `FirestoreBackupSchedule`
*   `FirestoreDocument`

## New Fields:

*   `OrgPolicyPolicy`: Added `spec.dryRunSpec`.

## Reconciliation Improvements:

*   Improved Normalization logic for `OrgPolicy`, `RunJob`, `TagsTagBinding`, and `VertexAIIndex` resources.

## New features:

*   The export logic was updated to handle `status.name` in export URIs.

## Bug Fixes:

*   TODO
