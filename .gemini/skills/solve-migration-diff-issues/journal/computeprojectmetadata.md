# Migration Journal: ComputeProjectMetadata

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computeprojectmetadata`
- Resource: `ComputeProjectMetadata` (`compute.cnrm.cloud.google.com/v1beta1`)

### Real GCP Execution & Recording
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  E2E_TEST_TIMEOUT=10m TEST_DEPENDENT_ORG_PROJECT_ID=cnrm-barni-4 ./hack/record-gcp "TestMigrationToDirect/fixtures/computeprojectmetadata$"
  ```
- **Initial Issue & Resolution:**
  - Running `./hack/record-gcp` without setting `TEST_DEPENDENT_ORG_PROJECT_ID` resulted in empty string substitution for `${TEST_DEPENDENT_ORG_PROJECT_ID}` in `dependencies.yaml`, `create.yaml`, and `update.yaml`.
  - The `Project` dependency controller failed validation with:
    `could not parse resource ...: error validating 'spec.resourceID' field: the value of 'spec.resourceID' is invalid: '' (empty string)`.
  - The test environment credentials (`overseer-kcc-tester@cnrm-barni-4.iam.gserviceaccount.com`) lacked organization-level `resourcemanager.projects.create` permissions on organization `128653134652`.
  - Setting `TEST_DEPENDENT_ORG_PROJECT_ID=cnrm-barni-4` allowed the singleton metadata resource and dependencies to run against the authorized test project.

### Takeover & Re-reconciliation Verification
- **Phase 1 (Legacy Create):** Successfully created the `ComputeProjectMetadata` resource via the legacy Terraform controller.
- **Phase 2 (Legacy Re-reconcile):** Re-reconciled under the legacy controller; Terraform reported its standard metadata diffs as captured in `_migration_diffs.json`.
- **Phase 3 (Direct Takeover):** The Direct controller took over the resource cleanly with a single GET request and **0-write no-op reconciliation** (`_http_migration_phase3_direct_takeover.log`). No takeover diff was detected.
- **Phase 4 (Direct Re-reconcile):** The Direct controller re-reconciled cleanly with a single GET request and **0-write no-op reconciliation** (`_http_migration_phase4_direct_re-reconciliation.log`). No diff was detected.
- Both real GCP and MockGCP runs (`TestMigrationToDirect` and `TestAllInSeries`) passed cleanly with 0 direct diffs.
