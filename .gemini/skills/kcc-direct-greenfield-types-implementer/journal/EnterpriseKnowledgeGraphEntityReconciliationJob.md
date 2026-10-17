# EnterpriseKnowledgeGraphEntityReconciliationJob Journal

## Context
Implementing Greenfield direct KRM types, identity, reference, and generate.sh support for `EnterpriseKnowledgeGraphEntityReconciliationJob` under `enterpriseknowledgegraph.cnrm.cloud.google.com/v1alpha1`.

## Resource Details
- Proto message: `google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob`
- Service: `google.cloud.enterpriseknowledgegraph.v1`
- Canonical URL pattern: `projects/{project}/locations/{location}/entityReconciliationJobs/{entity_reconciliation_job}`
- Identity type: `identity.ServerGeneratedIdentity` because `CreateEntityReconciliationJob` does not accept a user-specified entity reconciliation job ID in its RPC request; the resource ID is GCP-generated.

## Schema & Reference Decisions
- Spec fields:
  - `projectRef` (`*refsv1beta1.ProjectRef`): Project reference for the resource.
  - `location` (`*string`): Location string pointer.
  - `resourceID` (`*string`): Optional resource ID.
  - `inputConfig` (`*InputConfig`):
    - `bigqueryInputConfigs`: array of `BigQueryInputConfig`, each with `bigqueryTableRef` (`*bigqueryv1beta1.BigQueryTableRef`) and `gcsURI` (`*string`).
    - `entityType`: `*string` with enum validation `ENTITY_TYPE_UNSPECIFIED;PEOPLE;ESTABLISHMENT;PROPERTY;PRODUCT;ORGANIZATION;LOCAL_BUSINESS;PERSON`.
    - `previousResultBigqueryTableRef`: `*bigqueryv1beta1.BigQueryTableRef`.
  - `outputConfig` (`*OutputConfig`):
    - `bigqueryDatasetRef`: `*bigqueryv1beta1.DatasetRef`.
  - `reconConfig` (`*ReconConfig`):
    - `connectedComponentsConfig.weightThreshold`: `*float32`.
    - `affinityClusteringConfig.compressionRoundCount`: `*int64`.
    - `options.enableGeocodingSeparation`: `*bool`.
    - `modelConfig.modelName`: `*string`.
    - `modelConfig.versionTag`: `*string`.
- Status & ObservedState:
  - `observedGeneration`: `*int64`.
  - `externalRef`: `*string`.
  - `observedState`:
    - `state`: `*string` with job state enum validation.
    - `error`: `*common.Status`.
    - `createTime`, `endTime`, `updateTime`: `*string`.

## Observations & Notes
- There is currently no Go client SDK module for enterpriseknowledgegraph published on google-cloud-go (`cloud.google.com/go/enterpriseknowledgegraph`), so `generate-mapper` is omitted in `generate.sh` as this phase only covers types, CRD, identity, and reference.
- CI presubmit script `dev/ci/presubmits/tests-e2e-fixtures-enterpriseknowledgegraphentityreconciliationjob` was added to run fixtures for the `enterpriseknowledgegraph.cnrm.cloud.google.com` APIGroup.
