# v1.126.0

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* Special shout-outs to ... for their contributions to this release.
TODO: list contributors with `git log v1.125.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`

## Announcement 

* Promoted `ControllerReconciler` and `NamespacedControllerReconciler` CRDs from v1alpha1 to v1beta1.

## New Beta Resources (Direct Reconciler):

* `Placeholder`

## New Fields:

* `GkeHubFeatureMembership`
*   Added `spec.configmanagement.management` field to enable Config Sync Auto Upgrade.

## Modified Beta Reconciliation

We migrated the following reconciliation from the TF-based or DCL-based controller to the new Direct controller to enhance the reliability and performance. The existing API is unchanged. To use the direct controller, add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding CR object.

* **`SecretManagerSecret`**

  * #510 Enhanced `spec.rotation.nextRotationTime` to use a fixed datetime value to avoid relative `now()` friction. 
  * #1081 Fixed the `spec.replication.auto` immutable issue
  * #3051 Fixed the `spec.rotation.rotationPeriod` immutable issue 
  * (WIP Pending#3282)Added the in-use version aliases in `status.observedState.versionAliases`
  * Resolved update stalling issues. 

* **`SecretManagerSecretVersion`**
  
  * Resolved update stalling caused by `DependencyNotReady` errors.
  * Fixed the friction in `spec.enabled` that enabling/disabling a secret version does not take effect in GCP. 

  * **API Behavior Change** 
  `SecretManagerSecretVersion` uses [service generated ID](https://cloud.google.com/config-connector/docs/how-to/managing-resources-with-resource-ids). 
  The Terraform-based reconciler writes back the service generated ID into the `spec.resourceID` field. This conflicts with GitOps, and is a deviation from the Kubernetes API convetion for [`spec and status`](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status)
  
  With the direct controller, the ID is now stored in `status.version` and a new field `status.externalRef` is used to guardrail the identity. The direct controller does not edit the `spec`. 


* `GkeHubFeatureMembership` is now a direct resource

## New features:


## New Alpha Resources (Direct Reconciler):

* `Placeholder`

## Bug Fixes:
