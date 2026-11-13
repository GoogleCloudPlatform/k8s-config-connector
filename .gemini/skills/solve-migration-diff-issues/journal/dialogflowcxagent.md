# Migration Journal: DialogflowCXAgent

## Observations

### Test Fixtures
- Resource: `DialogflowCXAgent` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- Tested Fixtures:
  - `pkg/test/resourcefixture/testdata/basic/dialogflowcx/v1alpha1/dialogflowcxagent/dialogflowcxagent-minimal`
  - `pkg/test/resourcefixture/testdata/basic/dialogflowcx/v1alpha1/dialogflowcxagent/dialogflowcxagent-maximal`

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Commands:**
  ```bash
  ./hack/record-gcp TestMigrationToDirect/fixtures/dialogflowcxagent-minimal$
  ./hack/record-gcp TestMigrationToDirect/fixtures/dialogflowcxagent-maximal$
  ```
- **Execution Result:**
  - Both E2E migration tests ran against live GCP and passed cleanly.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully provisioned `DialogflowCXAgent` using the legacy Terraform controller (`_http_migration_phase1_legacy_create.log`).
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller with a no-op annotation; only read requests (`GET https://global-dialogflow.googleapis.com/v3/projects/${projectId}/locations/global/agents/...`) were made (`_http_migration_phase2_legacy_re-reconciliation.log`).
- **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with a single read-only `GET` request and a **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). Zero takeover diffs (`"isNewObject": false`) were detected.
- **Phase 4 (Direct Re-reconcile):** Re-reconciled with a no-op annotation under the Direct controller with a single read-only `GET` request and a **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). Zero diffs were detected.
- The recorded `_migration_diffs.json` files confirm that there are zero `"isNewObject": false` diff blocks for `DialogflowCXAgent`. Full state compatibility is maintained between the legacy Terraform controller and the direct controller.

### Ratcheting Exclusions
- `DialogflowCXAgent` is confirmed not present in the ratcheting exclusions list in `tests/e2e/ratcheting.go`.
- Hermetic MockGCP migration validation was also verified (`TestMigrationToDirect/fixtures/dialogflowcxagent-minimal` and `TestMigrationToDirect/fixtures/dialogflowcxagent-maximal`) and passed cleanly.
- `TestAllInSeries` fixtures were also verified against MockGCP via `./hack/compare-mock` and passed cleanly.
