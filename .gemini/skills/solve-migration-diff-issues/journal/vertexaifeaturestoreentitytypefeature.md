# Migration Journal: VertexAIFeaturestoreEntityTypeFeature

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/vertexai/v1alpha1/vertexaifeaturestoreentitytypefeature/vertexaifeaturestoreentitytypefeature-basic`
- Resource: `VertexAIFeaturestoreEntityTypeFeature` (`vertexai.cnrm.cloud.google.com/v1alpha1`)
- Dependency resources: `VertexAIFeaturestore`, `VertexAIFeaturestoreEntityType`

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp "TestMigrationToDirect/fixtures/vertexaifeaturestoreentitytypefeature-basic$"
  ```
- **Execution Result:**
  - The E2E test `TestMigrationToDirect/fixtures/vertexaifeaturestoreentitytypefeature-basic` ran against live GCP and completed in ~449s with a PASS status.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully provisioned dependencies (`VertexAIFeaturestore`, `VertexAIFeaturestoreEntityType`) and the primary resource `VertexAIFeaturestoreEntityTypeFeature` using the legacy Terraform controller (`_http_migration_phase1_legacy_create.log`).
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller with a no-op annotation; verified clean 0-write re-reconciliation with read-only GET requests (`_http_migration_phase2_legacy_re-reconciliation.log`).
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero diffs (`"isNewObject": false`) were detected.
- **Phase 4 (Direct Re-reconcile):** Re-reconciled with a no-op annotation under the Direct controller with a single read-only GET request and a **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs were detected.
- The recorded `_migration_diffs.json` contains no takeover diffs (`"isNewObject": false` is absent), confirming full state compatibility between the legacy Terraform controller and the Direct controller.

### Ratcheting Exclusions
- Verified `VertexAIFeaturestoreEntityTypeFeature` is not present in the ratcheting exclusions in `tests/e2e/ratcheting.go`.

### Hermetic MockGCP Execution
- Executed `TestMigrationToDirect` against MockGCP:
  - `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run "TestMigrationToDirect/fixtures/vertexaifeaturestoreentitytypefeature-basic$"`
  - Passed cleanly with 0 unexpected writes or diffs (~23s).
- Executed `TestAllInSeries` against MockGCP:
  - `./hack/compare-mock pkg/test/resourcefixture/testdata/basic/vertexai/v1alpha1/vertexaifeaturestoreentitytypefeature/vertexaifeaturestoreentitytypefeature-basic`
  - Passed cleanly with 0 unexpected writes or diffs (~40s).
