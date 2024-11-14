# v1.125.0

* Special shout-outs to @600lyy, @acpana, @anhdle-sso, @cheftako, @ericpang777, @gemmahou, @hankfreund, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @nb-goog, @svetakvsundhar, @xiaoweim, @yuwenma, @zicongmei, @ziyue-101 for their contributions to this release.

## New Beta Resources (Direct Reconciler):

* [`BigQueryConnectionConnection`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryconnection/bigqueryconnectionconnection)

     * Manage [connections](https://cloud.google.com/bigquery/docs/working-with-connections) to connect to Google services and external data sources

* [`BigQueryAnalyticsHubDataExchange`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryanalyticshub/bigqueryanalyticshubdataexchange.md)

  * Manage [data exchange](https://cloud.google.com/bigquery/docs/analytics-hub-introduction#data_exchanges) to enable self-service data sharing

* [`PrivilegedAccessManagerEntitlement`](https://cloud.google.com/config-connector/docs/reference/resource-docs/privilegedaccessmanager/privilegedaccessmanagerentitlement) 

     * Manage [entitlements](https://cloud.google.com/iam/docs/pam-create-entitlements) to grant for projects, folders, and organizations
 
* [`WorkstationCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/workstations/workstationcluster.md)

  * Manage [workstation cluster](https://cloud.google.com/workstations/docs/overview) to define a group of workstations in a particular region and the VPC network they're attached to.

## New Alpha Resources (Direct Reconciler):

* `KMSAutokeyConfig`
 
    * Manage the [KMS auto key](https://cloud.google.com/kms/docs/autokey-overview) which simplifies the CMEKs provisioning and assignment. 
 
## New Fields:

* [`AlloyDBInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/alloydb/alloydbinstance) (Beta)
  * Added `spec.networkConfig.enableOutboundPublicIp` field.
  * Added `status.outboundPublicIpAddresses` field.

## Reconciliation Improvements

We migrated the following resources from the Terraform-based or DCL-based controller to the new Direct Controller to enhance the reliability and performance. The resource CRD is unchanged.

* [`SQLInstance`](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlinstance)  

  * You can use the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on the SQLInstance CR object to opt-in the direct controller.
  * The direct reconciler contains 2 fix and improvement: 
    * Fix the upgrade and downgrade issue between ENTERPRISE and ENTERPRISE_PLUS.
    * Supports creating from clone functionality via `spec. cloneSource`
  * We migrated the `SQLInstance` from the Terraform-based or DCL-based controller to the new Direct Controller to enhance the reliability and performance. The CRD is unchanged.

* [`ComputeFirewallPolicyRule`](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computefirewallpolicyrule)

  * You can use the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on the `ComputeFirewallPolicyRule` CR object to opt-in the direct controller, which fixes the `targetResources` error *required value "priority" could not be found*.

## New features:

* Add *cluster mode* to manage the rate-limit for the Config Connector requests

  * In v1.119, we added [rate-limit control in namespace mode](https://cloud.google.com/config-connector/docs/how-to/customize-controller-manager-rate-limit). Users can configure `NamespacedControllerReconciler` object (Alpha) to set the rate-limit for the reconciling requests to the kube-apiserver for their Config Connector resources. 
  * In this release, we add this feature for cluster mode. User can configure `ControllerReconciler` object (Alpha) to set the rate-limit for all their cnrm manager controllers in the cluster. This [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/operator/config/samples/controller_reconciler_customization_sample.yaml) shows how to set up the configuration.

## Bug Fixes:

* [Issue 3007](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/3007) ComputeBackendService cannot refer clientTLSPolicy due to invalid format
* [Issue 2973](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2973) kubelet_config has `insecure_kubelet_readonly_port_enabled: true` set even if not configured in the `ContainerNodePool` object. 
* [Issue 3140](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/3007) BigQueryConnectionConnection requires UUID to acquire the resource. 
