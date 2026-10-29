# Migration Journal: KMSSecretCiphertext

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/kms/v1alpha1/kmssecretciphertext/kmssecretciphertext`
- Resource: `KMSSecretCiphertext` (`kms.cnrm.cloud.google.com/v1alpha1`)

### Real GCP Execution
- **Command**: `./hack/record-gcp TestMigrationToDirect/fixtures/kmssecretciphertext$`
- **GCP Project**: `cnrm-barni-4`
- **Results across phases**:
  - **Phase 1 (Legacy Create)**: Successfully provisioned dependencies (`KMSKeyRing`, `KMSCryptoKey`) and the primary resource `KMSSecretCiphertext` using the legacy Terraform controller.
  - **Phase 2 (Legacy Re-reconcile)**: Verified clean 0-write re-reconciliation with a single read-only `GET` request to `CryptoKey`.
  - **Phase 3 (Direct Takeover)**: Direct controller took over cleanly with 0-write no-op reconciliation (single read-only `GET` request to `CryptoKey`).
  - **Phase 4 (Direct Re-reconcile)**: Direct controller performed clean 0-write re-reconciliation with a single read-only `GET` request to `CryptoKey`.
- **Diff Analysis**:
  - `_migration_diffs.json` contains zero `"isNewObject": false` diff blocks during direct takeover (Phase 3) or direct re-reconciliation (Phase 4).
  - Only initial creation diffs (`"isNewObject": true`) from Terraform are recorded.
  - Full state compatibility is confirmed between the legacy Terraform controller and the direct controller.

### Ratcheting Exclusions
- Verified `KMSSecretCiphertext` is not present in `tests/e2e/ratcheting.go`.

### Hermetic MockGCP Execution
- Ran `TestMigrationToDirect` and `TestAllInSeries` against MockGCP:
  - `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestMigrationToDirect/fixtures/kmssecretciphertext$`
  - `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestAllInSeries/fixtures/kmssecretciphertext$`
- Both tests passed cleanly with 0 unexpected writes or diffs.
