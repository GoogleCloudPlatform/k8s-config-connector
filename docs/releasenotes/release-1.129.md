# Follow `SAMPLE_XXX` format to write the content. 
# Delete `SAMPLE_` and its content before publishing the release note. 
# Delete the entire header if no updates.  

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* Special shout-outs to ... for their contributions to this release.
TODO: list contributors with `git log v1.128.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`.

## Announcement 

### SAMPLE_Simplified and More Reliable Resource Development 

* We launched a major improvement to the Config Connector resource development!  Our new approach significantly enhances reliability and provides a more native Kubernetes experience. Learn more in our [guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/docs/develop-resources)   

## New Beta Resources (Direct Reconciler):

* [`SAMPLE_BigQueryConnectionConnection`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryconnection/bigqueryconnectionconnection) (This is a sample, your actual release note should not contain `SAMPLE_`, otherwise it will be deleted)

     * Manage [connections](https://cloud.google.com/bigquery/docs/working-with-connections) to connect to Google services and external data sources

## New Alpha Resources (Direct Reconciler):

* `SAMPLE_KMSAutokeyConfig` (This is a sample, your actual release note should not contain `SAMPLE_`, otherwise it will be deleted)
 
    * Manage the [KMS auto key](https://cloud.google.com/kms/docs/autokey-overview) which simplifies the CMEKs provisioning and assignment. 
 
## New Fields:

* [`SAMPLE_AlloyDBInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/alloydb/alloydbinstance) (This is a sample, your actual release note should not contain `SAMPLE_`, otherwise it will be deleted)

  * Added `spec.networkConfig.enableOutboundPublicIp` field.
  * Added `status.outboundPublicIpAddresses` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in behaviour.  The API is unchanged. To use the direct reconciler, add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object.  The following resources now have direct reconciliation support (and we list some of the issues that this fixes): 

* [`SAMPLE_SQLInstance`](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlinstance) (This is a sample, your actual release note should not contain `SAMPLE_`, otherwise it will be deleted)

    * Fix the upgrade and downgrade issue between ENTERPRISE and ENTERPRISE_PLUS.
    * Supports "creating from clone" via `spec.cloneSource`


## New features:

* SAMPLE_Add cluster mode ...  (This is a sample, your actual release note should not contain `SAMPLE_`, otherwise it will be deleted)

    *  User can configure the ControllerReconciler object (Alpha) to set the rate-limit for all their cnrm manager controllers in the cluster. This example shows how to set up the configuration.

## Bug Fixes:

* [SAMPLE_Issue 3007](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/3007) ComputeBackendService cannot refer clientTLSPolicy due to invalid format (This is a sample, your actual release note should not contain `SAMPLE_`, otherwise it will be deleted)