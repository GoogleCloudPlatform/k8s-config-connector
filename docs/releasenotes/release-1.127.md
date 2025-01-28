# v1.127.0

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* Special shout-outs to ... for their contributions to this release.
TODO: list contributors with `git log v1.127.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`

## Announcement 

## New Beta Resources (Direct Reconciler):

* `WorkstationConfig`
* `Workstation`
* `FirestoreDatabase`

## New Fields:

* `GkeHubFeatureMembership`
  * Added `spec.configmanagement.management` field to enable Config Sync Auto Upgrade.

* `Placeholder`

## Modified Beta Reconciliation

We migrated the following reconciliation from the TF-based or DCL-based controller to the new Direct controller to enhance the reliability and performance. The resource CRD is unchanged.

* `GkeHubFeatureMembership` is now a direct resource.

* `Placeholder`

## New Alpha Resources (Direct Reconciler):

* `Placeholder`
