# Certificate Manager Journals

### 2026-06-08 Direct Controller Implementation for CertificateManagerCertificate
- **Context**: Implementing direct controller and recording/verifying E2E fixtures for `CertificateManagerCertificate` under issue #9555.
- **Problem**: When migrating brownfield resources from legacy Terraform controllers to direct controllers, the HTTP logs (golden logs) often fail verification due to missing or mismatched headers like `User-Agent` (`kcc/${kccVersion}`).
- **Solution**: Set `forceDirect = true` in `tests/e2e/unified_test.go` and regenerate golden logs by running `WRITE_GOLDEN_OUTPUT=1 dev/ci/presubmits/tests-e2e-fixtures-certificatemanager`. This captures the new direct controller traffic correctly.
- **Impact**: Ensures that when other brownfield controllers are migrated, they can correctly update and align mockgcp and real GCP E2E HTTP verification logs.

### 2026-06-08 Verified CertificateManagerDNSAuthorization Direct Controller and E2E Fixtures
- **Context**: Verification of direct controller and E2E fixtures for `CertificateManagerDNSAuthorization` ([Issue #9544](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9544)).
- **Verification**: 
  - Ran E2E tests for the `certificatemanagerdnsauthorization-minimal`, `certificatemanagerdnsauthorization-full`, and `certificatemanagerdnsauthorization` fixtures against MockGCP.
  - All test scenarios passed successfully with no golden file diffs, verifying that the implementation of the controller, mappers, external resource mapping, fuzzer, and mock API behavior is fully correct, integrated, and functioning as expected.
- **Result**: The direct controller and E2E fixtures are 100% complete and verified.
