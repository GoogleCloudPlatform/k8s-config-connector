# v1.124.0

* Special shout-outs to @600lyy, @acpana, @anhdle-sso, @benjamin-maynard, @cheftako, @gemmahou, @hankfreund, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @nancynh, @svetakvsundhar, @xiaoweim, @yuwenma for their contributions to this release.

## Announcement

### Simplified and More Reliable Resource Development 

* We launched a major improvement to the Config Connector resource development!  Our new approach significantly enhances reliability and provides a more native Kubernetes experience. Learn more in our [guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/docs/develop-resources)   

## New Beta Resources (Direct Reconciler):

* `RedisCluster`

## New Fields:

* `CertificateManagerDNSAuthorization`

  * Added `spec.Location` field.

* `ComputeForwardingRule`

  * Added `spec.target.googleApisBundle` field (accept value `all-apis` or `vpc-sc`). Note, when configured this field, you are using the new Direct reconciliation.  

## Modified Beta Reconciliation

We migrated the following reconciliation from the TF-based or DCL-based controller to the new Direct controller to enhance the reliability and performance. The resource CRD is unchanged.

* `CertificateManagerDNSAuthorization`

## New Alpha Resources (Direct Reconciler):

* `PrivilegedAccessManagerEntitlement`
* `BigQueryAnalyticsHubDataExchange`
