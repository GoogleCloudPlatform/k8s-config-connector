# Follow `SAMPLE_XXX` format to write the content.Add commentMore actions

# Delete `SAMPLE_` and its content before publishing the release note.

# Delete the entire header if no updates.

# Run [mdformat](go/mdformat) before publishing this release notes.

** This version is not yet released; this document is gathering release notes
for the future release **

*   Special shout-outs to ... for their contributions to this release. TODO:
    list contributors with `git log v1.128.0... | grep Merge | grep from | awk
    '{print $6}' | cut -d '/' -f 1 | sort | uniq`.

## Announcement

### SAMPLE_Simplified and More Reliable Resource Development

*   We launched a major improvement to the Config Connector resource
    development! Our new approach significantly enhances reliability and
    provides a more native Kubernetes experience. Learn more in our
    [guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/docs/develop-resources)

## New Beta Resources (Direct Reconciler):

*   [`APIGatewayAPI`](https://cloud.google.com/config-connector/docs/reference/docs/reference/resource-docs/apigateway/apigatewayapi.md)

*   [`AppHubApplication`](https://cloud.google.com/config-connector/docs/reference/docs/reference/resource-docs/apphub/apphubapplication.md)

*   `StorageAnywhereCache`

*   [`SAMPLE_BigQueryConnectionConnection`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryconnection/bigqueryconnectionconnection)
    (This is a sample, your actual release note should not contain `SAMPLE_`,
    otherwise it will be deleted)

    *   Manage
        [connections](https://cloud.google.com/bigquery/docs/working-with-connections)
        to connect to Google services and external data sources

## New Alpha Resources (Direct Reconciler):

*   `BigtableLogicalView`

## New Fields:

*   [`SAMPLE_AlloyDBInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/alloydb/alloydbinstance)
    (This is a sample, your actual release note should not contain `SAMPLE_`,
    otherwise it will be deleted)

    *   Added `spec.networkConfig.enableOutboundPublicIp` field.
    *   Added `status.outboundPublicIpAddresses` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   [`SAMPLE_SQLInstance`](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlinstance)
    (This is a sample, your actual release note should not contain `SAMPLE_`,
    otherwise it will be deleted)

    *   Fix the upgrade and downgrade issue between ENTERPRISE and
        ENTERPRISE_PLUS.
    *   Supports "creating from clone" via `spec.cloneSource`

## Modified Beta Reconciliation

We migrated the following resources from the Terraform-based or DCL-based controller to the new Direct Controller. The resource CRD is unchanged.

*  `BigQueryTable`
  * You can use the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation on the `BigQueryTable` CR object to opt-in the direct controller.
  * The direct controller also supports adding BigQueryDataPolicies directly to BigQueryTable columns within `spec.schema`.

## New features:

*   SAMPLE_Add cluster mode ... (This is a sample, your actual release note
    should not contain `SAMPLE_`, otherwise it will be deleted)

    *   User can configure the ControllerReconciler object (Alpha) to set the
        rate-limit for all their cnrm manager controllers in the cluster. This
        example shows how to set up the configuration.

## Bug Fixes:

*   [PR#4808](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/4808)
    filtered out Kubernetes labels that are invalid for GCP in the
    ComputeForwardingRule direct controller, ensuring backward compatibility
    after migrating to the direct controller.
