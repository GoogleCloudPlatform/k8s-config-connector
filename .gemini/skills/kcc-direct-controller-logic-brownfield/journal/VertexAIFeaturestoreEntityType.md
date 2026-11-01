# VertexAIFeaturestoreEntityType Direct Migration Journal

## Overview
Successfully implemented the direct controller for `VertexAIFeaturestoreEntityType` (`vertexai.cnrm.cloud.google.com/v1alpha1`) and verified dual-controller compatibility, real GCP recording, and MockGCP golden log alignment.

## Key Implementation Details

1. **Client & Transport**:
   - Used GAPIC REST client (`aiplatform.NewFeaturestoreRESTClient`) with endpoint scoped to the resource location (`https://{location}-aiplatform.googleapis.com`).

2. **Reconciliation & Brownfield Alignment**:
   - **Create**: Calls `CreateEntityType` with `Parent` (`projects/{project}/locations/{location}/featurestores/{featurestore}`), `EntityTypeId`, and `EntityType`. Waits on the Long-Running Operation (`CreateEntityTypeOperation.Wait(ctx)`), then performs a follow-up `GetEntityType` call to retrieve the fully-populated resource before updating status.
   - **Update**: Uses `common.CompareBrownfieldSpecAndLabels` to compare desired KRM spec/labels against actual proto state. Server defaults (e.g. `monitoringInterval = "86400s"`, `offlineStorageTtlDays = 4000`) are merged into desired state to prevent spurious diffs. Updates are dispatched via `UpdateEntityType` with the calculated `UpdateMask`.
   - **Delete**: Calls `DeleteEntityType` and waits for completion. Gracefully handles 404 / NotFound for idempotency.
   - **Export**: Generates KRM object with canonical parent featurestore reference, resourceID, and labels.

3. **Validation & Verification**:
   - Recorded real GCP logs with `hack/record-gcp "fixtures/^vertexaifeaturestoreentitytype-basic$"` and in-place migration test with `hack/record-gcp "TestMigrationToDirect/fixtures/^vertexaifeaturestoreentitytype-basic$"` against GCP project `cnrm-barni-4` between Sat Oct 31 20:01:30 PDT 2026 and Sat Oct 31 20:23:51 PDT 2026.
   - Re-verified against MockGCP via `hack/compare-mock`.
   - Verified log alignment with `TestGoldenLogAlignment` passing with 0 errors.
   - Verified backward compatibility and schema checks with `go test ./tests/apichecks/...` and `go test ./pkg/crd/template/...`.
