# Journal: Match real GCP behavior in MockGCP for VertexAIFeaturestoreEntityTypeFeature

## Overview
This journal documents the findings, design decisions, and learnings from matching MockGCP behavior with real GCP for `VertexAIFeaturestoreEntityTypeFeature` (and its parent `VertexAIFeaturestoreEntityType`).

## Observations & Findings

1. **Service and Proto Identification**:
   - The GCP proto for Featurestore is `google.cloud.aiplatform.v1beta1.FeaturestoreService`.
   - Featurestore service is already registered in `mockgcp/mockaiplatform/service.go`.
   - The proto methods are:
     - `CreateEntityType`, `GetEntityType`, `UpdateEntityType`, `DeleteEntityType`
     - `CreateFeature`, `GetFeature`, `UpdateFeature`, `DeleteFeature`
   - Create operations for both `EntityType` and `Feature` are LROs returning operations with `CreateEntityTypeOperationMetadata` and `CreateFeatureOperationMetadata` respectively.
   - Delete operations are LROs returning `DeleteOperationMetadata`.
   - Update operations (`UpdateEntityType`, `UpdateFeature`) are synchronous in the GCP API and return the updated resource directly.

2. **Ratcheting & Naming**:
   - `VertexAIFeaturestoreEntityTypeFeature` and `VertexAIFeaturestoreEntityType` were removed from `tests/e2e/ratcheting.go` to test SSA and re-reconciliation.
   - Vertex AI resource naming requires lowercase alphanumeric identifiers (dashes are not permitted in `featurestore`, `entityType`, or `feature` IDs).
   - In `dependencies.yaml`, `featurestore${uniqueId}` and `entitytype${uniqueId}` were used.
   - In `create.yaml` and `update.yaml`, `feature${uniqueId}` was used.

3. **Real GCP Recording**:
   - Recorded against live GCP project `cnrm-barni-4` using `./hack/record-gcp "fixtures/^vertexaifeaturestoreentitytypefeature-basic$"`.
   - Start timestamp: Thu Oct 22 17:55:50 PDT 2026 (00:55:50 UTC 2026).
   - Finish timestamp: Thu Oct 22 18:03:28 PDT 2026 (01:03:28 UTC 2026).
   - Successfully verified resource creation, update, deletion, and written golden output to `_http.log`.
   - Committed the recorded golden files in their own commit: `34570aad2a`.

4. **MockGCP Implementation & Alignment**:
   - Implemented `GetEntityType`, `CreateEntityType`, `UpdateEntityType`, and `DeleteEntityType` in `mockgcp/mockaiplatform/featurestore.go`.
   - Implemented `GetFeature`, `CreateFeature`, `UpdateFeature`, and `DeleteFeature` in `mockgcp/mockaiplatform/featurestore.go`.
   - Enabled `VertexAIFeaturestoreEntityType` and `VertexAIFeaturestoreEntityTypeFeature` in `config/tests/samples/create/harness.go`.
   - Executed `./hack/compare-mock "fixtures/^vertexaifeaturestoreentitytypefeature-basic$"`:
     - `_http_mock.log` was generated.
     - `TestGoldenLogAlignment` passed with 0 errors across all vertexai fixtures.
