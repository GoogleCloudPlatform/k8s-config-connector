# Certificate Manager Journals

### 2026-06-08 Direct Controller Implementation for CertificateManagerCertificate
- **Context**: Implementing direct controller and recording/verifying E2E fixtures for `CertificateManagerCertificate` under issue #9555.
- **Problem**: When migrating brownfield resources from legacy Terraform controllers to direct controllers, the HTTP logs (golden logs) often fail verification due to missing or mismatched headers like `User-Agent` (`kcc/${kccVersion}`).
- **Solution**: Set `forceDirect = true` in `tests/e2e/unified_test.go` and regenerate golden logs by running `WRITE_GOLDEN_OUTPUT=1 dev/ci/presubmits/tests-e2e-fixtures-certificatemanager`. This captures the new direct controller traffic correctly.
- **Impact**: Ensures that when other brownfield controllers are migrated, they can correctly update and align mockgcp and real GCP E2E HTTP verification logs.

### 2026-09-15 Greenfield Direct Controller for CertificateManagerTrustConfig
- **Context**: Implementing Greenfield direct controller and recording/verifying E2E fixtures for `CertificateManagerTrustConfig` under issue #11793.
- **Problem**: YAML parsing of multiline PEM certificate blocks using stripping operators like `|-` drops trailing newlines. When GCP stores and returns PEM certificates, it normalizes them with a trailing newline. This slight discrepancy in the PEM string payload (with vs without a trailing newline) causes `common.DiffForTopLevelFields` to flag a false diff during re-reconciliation, triggering constant update loops.
- **Solution**: Added a `normalizeTrustConfig` helper in the controller that trims all leading/trailing whitespace/newlines from PEM certificate strings and ensures a standard single trailing newline is appended. This is applied to both desired and maskedActual configs before performing top-level comparison.
- **Impact**: Clean, 100% correct, and stable re-reconciliation of `CertificateManagerTrustConfig` resources without any false diffs or endless API calls.
