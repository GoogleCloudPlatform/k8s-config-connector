# v1.124.0

* Special shout-outs to @600lyy, @acpana, @anhdle-sso, @benjamin-maynard, @cheftako, @gemmahou, @hankfreund, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @nancynh, @svetakvsundhar, @xiaoweim, @yuwenma for their contributions to this release.

## Announcement

* We launched a major improvement to the Config Connector resource development!  Our new approach significantly enhances reliability and provides a more native Kubernetes experience. Learn more in our [guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/docs/develop-resources)   

* You can use the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on `ComputeForwardingRule` resource to opt-in
the Direct Cloud Reconciler, which provides sustainable reconciliation support for different types of ForwardingRule uses, including setting labels for ServiceAttachments and Private Service Connect.

## Resources promoted from alpha to beta:

* `RedisCluster` (v1beta1)

## Direct Cloud Reconciler:

* `BigQueryAnalyticsHubDataExchange` (v1alpha1).
* `CertificateManagerDNSAuthorization` (v1beta1).
* `ComputeForwardingRule` (v1beta1).

## New Resources:

* Added support for `PrivilegedAccessManagerEntitlement` (v1alpha1) resource.

## New Fields:

* CertificateManagerDNSAuthorization

  * Added `spec.Location` field.

* ComputeForwardingRule

  * Added `spec.target.googleApisBundle` field
