# v1.125.0

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* Special shout-outs to ... for their
  contributions to this release.
TODO: list contributors with `git log v1.124.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`

## Resources promoted from alpha to beta:

* `BigQueryAnlayticsHubDataExchange` is now a v1beta1 resource.
* `PrivilegedAccessManagerEntitlement` is now a v1beta1 resource.
* `RedisCluster` is now a v1beta1 resource.

## Modified Beta Reconciliation

We migrated the following reconciliation from the TF-based or DCL-based controller to the new Direct controller to enhance the reliability and performance. The resource CRD is unchanged.

* `ComputeFirewallPolicyRule`

  * You can use the alpha.cnrm.cloud.google.com/reconciler: direct annotation on ComputeFirewallPolicyRule resource to opt-in
    the Direct Cloud Reconciler, which fixes the issue when updating `targetResources`.

## New Resources:

* Added support for `PlaceholderKind` (v1beta1) resource.

## New Fields:

* `AlloyDBInstance`
  * Added `spec.networkConfig.enableOutboundPublicIp` field.
  * Added `status.outboundPublicIpAddresses` field.

## New features:

* Allow more customization of resource reconciliation in cluster mode
  * Added a new `ControllerReconciler` CRD (v1alpha1). See [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/operator/config/samples/controller_reconciler_customization_sample.yaml)
  * This feature allows users to customize the client-side kube-apiserver request rate limit when Config Connector is runnning in cluster mode.

## Bug Fixes:

* [Incorrect format of clientTLSPolicy when referenced from ComputeBackendService](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/3007)
