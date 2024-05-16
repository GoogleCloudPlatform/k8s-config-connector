# v1.118.1

* This release introduces our new direct-reconciliation mechanism to reconcile KCC resources (without relying on terraform). Currently it only applies to `LoggingLogMetric`.

* Special shout-outs to @199201shubhamsahu, @acpana, @anhdle-sso, @barney-s, @cheftako, @gemmahou, @jingyih, @justinsb, @katrielt, @vmiglani, @xiaoweim and @yuwenma for their
  contributions to this release. 

## Direct Cloud Reconciler:

* `LoggingLogMetric`
  * This resource no longer depends on Terraform. Its reconciliation is moved to a KCC direct controller.
  * This is our very first KCC directly-reconciled resource. We will announce more in the upcoming releases. 🎉🎉🎉

## New Resource:

* `ComputeNetworkFirewallPolicyRule` (alpha)

## New Fields:

* `LoggingLogMetric`
  * Add the `spec.loggingLogBucketRef` field to support bucket reference.

## Fixes:

* `SQLInstance`
  * Fix the permanent diff bug in `spec.settings.edition` update.
