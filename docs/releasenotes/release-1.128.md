* Special shout-outs to @600lyy, @acpana, @anhdle-sso, @barney-s, @Camila-B, @cheftako, @ericpang777, @gemmahou, @haiyanmeng, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @nb-goog, @tarynlucas, @xiaoweim, @yuwenma, @ziyue-101 for their contributions to this release. 

## Announcement

* `ComputeFirewallPolicyRule` is switched to the direct approach by default. Previously this direct approach is introduced as a opt-in since release 1.125. 

## New Beta Resources (Direct Reconciler):

* [`ApigeeEnvgroup`](https://cloud.google.com/config-connector/docs/reference/resource-docs/apigee/apigeeenvgroup)

     * Define [environment groups](https://cloud.google.com/apigee/docs/api-platform/fundamentals/environments/envgroups) to specify the hostnames for routing traffic to Apigee environments.

* [`KMSAutokeyConfig`](https://cloud.google.com/config-connector/docs/reference/resource-docs/kms/kmsautokeyconfig)

    * Manage the [KMS auto key](https://cloud.google.com/kms/docs/autokey-overview) which simplifies the CMEKs provisioning and assignment.

## New Alpha Resources (Direct Reconciler):

* `IAPSettings`
    * Customize the [Identity-Aware Proxy (IAP)](https://cloud.google.com/iap/docs/customizing) settings for applications and services running on Google Cloud Platform.
    
* `SecureSourceManangerInstance`

* `SecureSourceManangerRepository`

## New Fields:

* [`SpannerInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance)

    * You need to use the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on `SpannerInstance` resource to opt-in these features.

        * `spec.autoscalingConfig`
        * `spec.edition`

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in behaviour.  The API is unchanged. To use the direct reconciler, add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object.  The following resources now have direct reconciliation support (and we list some of the issues that this fixes):

* [`AlloyDBInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/alloydb/alloydbinstance)

* [`SpannerInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance)

## Bug Fixes:

* [Fixed](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/3521) the incorrect format validation for the following fields in resource [`DataformRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dataform/dataformrepository).

    * `spec.gitRemoteSettings.authenticationTokenSecretVersionRef`
    * `spec.gitRemoteSettings.sshAuthenticationConfig.userPrivateKeySecretVersionRef`
    * `spec.npmrcEnvironmentVariablesSecretVersionRef`
