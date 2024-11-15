# v1.125.0

* Special shout-outs to @600lyy, @acpana, @anhdle-sso, @cheftako, @ericpang777, @gemmahou, @hankfreund, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @nb-goog, @svetakvsundhar, @xiaoweim, @yuwenma, @zicongmei, @ziyue-101 for their contributions to this release.

## Announcement 

## New Beta Resources (Direct Reconciler):

* `BigQueryAnlayticsHubDataExchange`
* `PrivilegedAccessManagerEntitlement`
* `RedisCluster`
* `WorkstationCluster`

## New Fields:

* `AlloyDBInstance`
  * Added `spec.networkConfig.enableOutboundPublicIp` field.
  * Added `status.outboundPublicIpAddresses` field.

## Modified Beta Reconciliation

We migrated the following resources from the Terraform-based or DCL-based controller to the new Direct Controller to enhance the reliability and performance. The resource CRD is unchanged.

* `SQLInstance`

  * You can use the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on the `SQLInstance` CR object to opt-in the direct controller, which fixes issues with updating from ENTERPRISE -> ENTERPRISE_PLUS edition and allows "create from clone" functionality.

* `ComputeFirewallPolicyRule`

  * You can use the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on the `ComputeFirewallPolicyRule` CR object to opt-in the direct controller, which fixes the `targetResources` error *required value "priority" could not be found*.

## New features:

* Add *cluster mode* to manage the rate-limit for the Config Connector requests

  * Previously we added [rate-limit control in namespace mode](https://cloud.google.com/config-connector/docs/how-to/customize-controller-manager-rate-limit). Users can configure the `NamespacedControllerReconciler` object to set the rate-limit for the reconciling requests to the kube-apiserver for their Config Connector resources. 
  * This release adds this feature in the cluster mode in `ControllerReconciler` object. This [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/operator/config/samples/controller_reconciler_customization_sample.yaml) shows how to set up the configuration.

## Bug Fixes:

* [Issue 3007](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/3007) ComputeBackendService cannot refer clientTLSPolicy due to invalid format
* [Issue 2973](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2973) kubelet_config has `insecure_kubelet_readonly_port_enabled: true` set even if not configured in the `ContainerNodePool` object. 
* [Issue 3140](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/3007) BigQueryConnectionConnection requires UUID to acquire the resource. 