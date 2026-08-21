# Release 1.139.0

*   Special shout-outs to @anhdle-sso, @gemmahou, @justinsb, @maqiuyujoyce, @shavonz, @xiaoweim, @yuwenma for their contributions to this release.

## New Alpha Resources (Direct Reconciler):
*   `FirestoreField`
*   `FirestoreDocument`

## Reconciliation Improvements
*   `IAM partial policy management`: Now supports direct reconciliation.

## New features:
*   The controller type is now reported at the start and end of reconciliation.
*   Mockgcp now supports `iap oauth brands` and `bigtable materializedview`.

## Bug Fixes:
*   Reduces the memory footprint of the recorder.
*   `SQLInstance`: Fixes an issue where empty `maintenanceVersion` patches were sent. The `settings` and `maintenanceVersion` fields are now unmanaged.
*   `FirestoreDatabase`: Fixes boolean value exports.