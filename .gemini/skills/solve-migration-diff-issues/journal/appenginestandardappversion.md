# Migration Journal: AppEngineStandardAppVersion

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/appengine/v1alpha1/appenginestandardappversion/appenginestandardappversionbasic`
- Resource: `AppEngineStandardAppVersion` (`appengine.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution & Investigation
- Ran `./hack/record-gcp TestMigrationToDirect/fixtures/appenginestandardappversionbasic$` against live GCP project `cnrm-barni-4`.
- **Baseline Observation**:
  - In Phase 1 (Legacy Create), the legacy Terraform controller successfully creates the `AppEngineStandardAppVersion` resource on GCP.
  - In Phase 2 (Legacy Re-reconcile), the legacy Terraform provider detects diffs on `deployment.0.files`, `entrypoint.0.shell`, and `env_variables` because `resourceAppEngineStandardAppVersionRead` does not read back or persist these write-only/unflattened fields into Terraform state. Consequently, Terraform triggers an update via `resourceAppEngineStandardAppVersionUpdate` by issuing a `POST https://appengine.googleapis.com/v1/apps/${projectId}/services/default/versions?alt=json` (CreateVersion) request and polls an App Engine Long Running Operation (LRO).
  - Because App Engine version creation takes 40-50 seconds on GCP, and `waitForReconciliationAfterPatch` in `script_test.go` has a 5-second polling timeout, the test advances to Phase 3 while Terraform's Phase 2 operation is still in-flight.
  - The in-flight Phase 2 `POST` request completes during Phase 3, leading to `migration_test.go:336: FAIL: unexpected write request during migration reconciliation: POST https://appengine.googleapis.com/v1/apps/cnrm-barni-4/services/default/versions?alt=json`.
  - The baseline diffs and HTTP cassettes were recorded and committed in the diagnosis commit (`42883a86d5`).

### Direct Controller Takeover Verification (MockGCP)
- Ran `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestMigrationToDirect/fixtures/appenginestandardappversionbasic$`.
- In MockGCP, where LRO operations complete immediately without the 40-50s GCP propagation delay:
  - Phase 1 (Legacy Create) successfully provisions the resource.
  - Phase 2 (Legacy Re-reconciliation) completes within milliseconds.
  - Phase 3 (Direct Takeover) verifies that the direct controller adopts the resource cleanly with a 0-write no-op reconciliation (only issuing `GET` verification requests).
  - Phase 4 (Direct Re-reconciliation) verifies that the direct controller re-reconciles cleanly with a 0-write no-op reconciliation (only issuing `GET` verification requests).
  - `_migration_diffs.json` confirms that the Direct controller detects 0 takeover diffs (`"isNewObject": false` is absent for the direct controller).
  - The direct controller's existing normalization in `normalizeStandardAppVersionForDiff` (clearing write-only `Entrypoint` and server-computed `Deployment.Files[].Sha1Sum`) ensures clean state comparison.
- Ran dual-controller test `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestAllInSeries/fixtures/appenginestandardappversionbasic$` and `./hack/compare-mock fixtures/appenginestandardappversionbasic$`, which both passed cleanly.

### Ratcheting Exclusions
- Verified `tests/e2e/ratcheting.go`: `AppEngineStandardAppVersion` is currently excluded from `ShouldTestRereconiliation` due to the legacy Terraform provider's re-reconciliation update behavior on `deployment.0.files`. Attempting to suppress diffs in the legacy Terraform provider causes Terraform's update method to omit `deployment`, failing real GCP updates with `400 Bad Request: version.deployment is required`. Therefore, `AppEngineStandardAppVersion` properly remains in `tests/e2e/ratcheting.go` until the default controller is promoted to Direct.
