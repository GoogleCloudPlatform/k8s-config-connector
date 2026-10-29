# Migration Journal: KMSCryptoKeyVersion

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/kms/v1alpha1/kmscryptokeyversion/kmscryptokeyversion`
- Resource: `KMSCryptoKeyVersion` (`kms.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution
- **Command**: `./hack/record-gcp TestMigrationToDirect/fixtures/kmscryptokeyversion$`
- **GCP Project**: `cnrm-barni-4`
- **Results across phases**:
  - **Phase 1 (Legacy Create)**: Successfully created the resource using the legacy Terraform controller.
  - **Phase 2 (Legacy Re-reconcile)**: Verified clean 0-write re-reconciliation with a single read-only `GET` request.
  - **Phase 3 (Direct Takeover)**: Direct controller took over cleanly with 0-write no-op reconciliation (single read-only `GET` request).
  - **Phase 4 (Direct Re-reconcile)**: Direct controller performed clean 0-write re-reconciliation with a single read-only `GET` request.
- **Diff Analysis**:
  - `_migration_diffs.json` contains zero `"isNewObject": false` diff blocks for `KMSCryptoKeyVersion`.
  - Full state compatibility is confirmed between the legacy Terraform controller and the direct controller.

### Ratcheting Exclusions
- Verified `KMSCryptoKeyVersion` is not in `tests/e2e/ratcheting.go`.

### Hermetic MockGCP Execution
- Ran `TestMigrationToDirect` and `TestAllInSeries` against MockGCP:
  - `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestMigrationToDirect/fixtures/kmscryptokeyversion$`
  - `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestAllInSeries/fixtures/kmscryptokeyversion$`
- Both tests passed cleanly with 0 unexpected writes or diffs.
