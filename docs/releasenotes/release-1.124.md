# v1.124.0

* Special shout-outs to @600lyy, @acpana, @anhdle-sso, @benjamin-maynard, @cheftako, @gemmahou, @hankfreund, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @nancynh, @svetakvsundhar, @xiaoweim, @yuwenma for their contributions to this release.

## Announcement

### Simplified and More Reliable Resource Development 

* We launched a major improvement to the Config Connector resource development!  Our new approach significantly enhances reliability and provides a more native Kubernetes experience. Learn more in our [guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/docs/develop-resources)   


### Enhanced `ComputeForwardingRule` Reliability (Direct reconciliation)

* We improves the reliability of `ComputeForwardingRule` reconciliation and introduces advanced configuration options.  These enhancements include fixes for defaulting and label issues, and adding the Private Service Connect setup using the new `spec.target.googleApisBundle` field (accept value `all-apis` or `vpc-sc`). To enable these improvements, add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to your `ComputeForwardingRule` resources.

## New Beta Resources (Direct Reconciler):

* `RedisCluster`
* `CertificateManagerDNSAuthorization` (migrated)
* `ComputeForwardingRule` (migrated)

## New Fields:

* `CertificateManagerDNSAuthorization`

  * Added `spec.Location` field.

* `ComputeForwardingRule`

  * Added `spec.target.googleApisBundle` field. Note, when using this field, you are using the Direct reconciliation by default and you get all the benefits descript in "Enhanced `ComputeForwardingRule` Reliability" even if not adding the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation. 

## New Alpha Resources (Direct Reconciler):

* `PrivilegedAccessManagerEntitlement`
* `BigQueryAnalyticsHubDataExchange`
