# Journal: Match real GCP behavior in MockGCP for VertexAIFeaturestoreEntityType

## Overview
This journal documents the findings, design decisions, and learnings from matching MockGCP behavior with real GCP for `VertexAIFeaturestoreEntityType`.

## Observations & Findings

1. **Service and Proto Identification**:
   - The GCP proto for Featurestore is `google.cloud.aiplatform.v1beta1.FeaturestoreService`.
   - Featurestore service is registered in `mockgcp/mockaiplatform/service.go`.
   - The entityType proto methods are: `CreateEntityType`, `GetEntityType`, `UpdateEntityType`, `DeleteEntityType`.
   - Create and Delete are long-running operations (LROs) returning `CreateEntityTypeOperationMetadata` and `DeleteOperationMetadata`.
   - `UpdateEntityType` is synchronous and returns `*pb.EntityType` directly.

2. **Real GCP Recording**:
   - Recorded against live GCP project `cnrm-barni-4` using `./hack/record-gcp "fixtures/^vertexaifeaturestoreentitytype-basic$"`.
   - Successfully verified resource creation, update, re-reconciliation, and deletion, writing golden output to `_http.log`.
   - Golden files and test fixtures were committed in their own commit: `ce64366788`.

3. **GCP API Field Behaviors**:
   - `monitoringInterval`: When `snapshotAnalysis.monitoringIntervalDays` is provided (e.g., 1 day), real GCP automatically populates `snapshotAnalysis.monitoringInterval = "86400s"` (derived as `days * 86400s`). MockGCP now computes this default in `populateDefaultsForEntityType`.
   - `offlineStorageTtlDays`: Real GCP defaults `offlineStorageTtlDays` to `4000` when unset.
   - `monitoringConfig` on Update: In real GCP, `monitoringConfig` is ignored during `UpdateEntityType` when updating an EntityType via the Terraform provider / REST API. Attempting to change `monitoringConfig` during update causes continuous re-reconciliation drift if the mock allows updates. MockGCP filters out `monitoring_config` / `monitoringConfig` from updateMask paths to match real GCP behavior.

4. **MockGCP Implementation & Alignment**:
   - Implemented `populateDefaultsForEntityType` in `mockgcp/mockaiplatform/featurestore.go` to compute `monitoringInterval` from `monitoringIntervalDays` and default `offlineStorageTtlDays` to 4000.
   - Updated `UpdateEntityType` to ignore `monitoring_config` / `monitoringConfig` field mask paths, matching real GCP behavior.
   - Executed `./hack/compare-mock "fixtures/^vertexaifeaturestoreentitytype-basic$"`:
     - All tests passed cleanly including re-reconciliation.
     - `TestGoldenLogAlignment` passed with 0 errors across all vertexai fixtures.
