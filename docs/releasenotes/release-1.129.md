
* Special shout-outs to acpana@, anhdle-sso@, barney-s@, cheftako@, ericpang777@, gemmahou@, jasonvigil@, jingyih@, justinsb@, maqiuyujoyce@, xiaoweim@, yuwenma@ for their contributions to this release.

## New Beta Resources (Direct Reconciler):

* [`ManagedKafkaCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/managedkafka/managedkafkacluster)

* [`ApigeeInstance`](https://cloud.google.com/apigee/docs/api-platform/system-administration/instances)

## New Alpha Resources (Direct Reconciler):

* `ManagedKafkaTopic`
* `ApigeeInstanceAttachment`
* `ApigeeEnvgroupAttachment`
* `ApigeeEndpointAttachment`
 

## Reconciliation Improvements

* [`SQLInstance`](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlinstance)

    * All SQLInstance types are now reconciled using the new direct controller (instead of the legacy Terraform-based controller). The previous "opt-in" annotation (documented [here](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/docs/features/optin.md)) no longer applies. Users no longer need to apply the "opt-in" annotation to SQLInstance resources to enable the direct controller. Regardless of the presence (or absence) of an opt-in annotation on SQLInstance resources, the direct reconciler will be used.
    * This change enables all SQLInstance resources to switch from edition ENTERPRISE -> ENTERPRISE_PLUS and fixes [this bug](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/2907).
