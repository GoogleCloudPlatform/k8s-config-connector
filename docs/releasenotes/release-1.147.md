# Release 1.147.1

** This version is not yet released; this document is gathering release notes for the future release **

* Special shout-outs to @anhdle-sso, @cheftako, @codebot-robot, and @xiaoweim for their contributions to this release.

## New features:

* Added a `--skip-name-validation` flag to Config Connector controllers to bypass duplicate controller name checks during registration. (#7075)
* Enhanced `config-connector preview` to support side-by-side comparison between default and alternative controllers. (#7083)

---

# Release 1.147.0

## This version has not yet been released, this document is to prepare release notes for the upcoming release

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @dhavalbhensdadiya-crest, @gemmahou, @katrielt, @maqiuyujoyce, @xiaoweim for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`CloudBuildTrigger`](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudbuild/cloudbuildtrigger)
    *   Direct reconciliation support added.

*   [`DataCatalogTaxonomy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/datacatalog/datacatalogtaxonomy)
    *   Direct reconciliation support added.

*   [`BigQueryRoutine`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigquery/bigqueryroutine)
    *   Promoted from Alpha to Stable (Beta status).

*   [`BigQueryAnalyticsHubDataExchange`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryanalyticshub/bigqueryanalyticshubdataexchange)
    *   Promoted from Alpha to Stable (Beta status).

*   [`ConfigControllerInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/configcontroller/configcontrollerinstance)
    *   Promoted from Alpha to Stable (Beta status).

*   [`TagsLocationTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagslocationtagbinding)
    *   Promoted from Alpha to Stable (Beta status).

*   [`TagsTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagstagbinding)
    *   Promoted from Alpha to Stable (Beta status).

*   [`TagsTagKey`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagstagkey)
    *   Promoted from Alpha to Stable (Beta status).

*   [`TagsTagValue`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagstagvalue)
    *   Promoted from Alpha to Stable (Beta status).

*   [`VertexAIMetadataStore`](https://cloud.google.com/config-connector/docs/reference/resource-docs/vertexai/vertexaimetadatastore)
    *   Promoted from Alpha to Stable (Beta status).

## New Alpha Resources (Direct Reconciler):

*   [`DataCatalogPolicyTag`](https://cloud.google.com/config-connector/docs/reference/resource-docs/datacatalog/datacatalogpolicytag)
    *   Direct reconciliation support added.

## Reconciliation Improvements

We have added support for structured diff reporting to the following direct controllers to improve logging and debugging:

*   [`MetastoreService`](https://cloud.google.com/config-connector/docs/reference/resource-docs/metastore/metastoreservice)
*   [`BigQueryTable`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigquery/bigquerytable)
*   [`BigQueryReservationAssignment`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigqueryreservation/bigqueryreservationassignment)
*   [`WorkflowsWorkflow`](https://cloud.google.com/config-connector/docs/reference/resource-docs/workflows/workflowsworkflow)
*   [`PrivilegedAccessManagerEntitlement`](https://cloud.google.com/config-connector/docs/reference/resource-docs/privilegedaccessmanager/privilegedaccessmanagerentitlement)
*   [`DataformRepository`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dataform/dataformrepository)
*   [`CloudIdentityGroup`](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudidentity/cloudidentitygroup)
*   [`BigQueryDataset`](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigquery/bigquerydataset)
*   [`CertificateManagerDNSAuthorization`](https://cloud.google.com/config-connector/docs/reference/resource-docs/certificatemanager/certificatemanagerdnsauthorization)
*   [`VMwareEngineExternalAccessRule`](https://cloud.google.com/config-connector/docs/reference/resource-docs/vmwareengine/vmwareengineexternalaccessrule)

## New features:

*   [`config-connector preview`](https://cloud.google.com/config-connector/docs/how-to/preview)
    *   Added support for reconciler overrides, allowing users to test direct reconciliation for supported resources.
    *   Added CRD filtering support to the preview recorder.

*   [`CloudDeployTarget`](https://cloud.google.com/config-connector/docs/reference/resource-docs/clouddeploy/clouddeploytarget)
    *   Renamed `multiTarget.targets` field to `targetRefs` to follow Config Connector naming conventions.

*   [`CloudDeployCustomTargetType`](https://cloud.google.com/config-connector/docs/reference/resource-docs/clouddeploy/clouddeploycustomtargettype)
    *   Renamed from `DeployCustomTargetType` to follow Google Cloud naming conventions.

## Bug Fixes:

*   [`StorageDefaultObjectAccessControl`](https://cloud.google.com/config-connector/docs/reference/resource-docs/storage/storagedefaultobjectaccesscontrol)
    *   Fixed typo in documentation.
