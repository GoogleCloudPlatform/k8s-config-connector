### [2026-07-22] Scaffolding types for LicenseManagerConfiguration
- **Context**: Implementing direct types, CRD, and IdentityV2 for `LicenseManagerConfiguration` as part of issue #10281.
- **Problem**: Greenfield resource generator outputs reachable/unreachable structs in different files, requiring manual copy-paste of types (such as `BillingInfo` and `BillingInfoObservedState`) to `licensemanagerconfiguration_types.go`, as well as setting proper kubebuilder validations, copyright year 2026, and stability-level labels.
- **Solution**: Manually incorporated and validated the `BillingInfo`, `UserCountBillingInfo`, and corresponding observed state definitions within `licensemanagerconfiguration_types.go`, adding `cnrm.cloud.google.com/stability-level=alpha`.
- **Impact**: Ensures that when the controller and mappers are subsequently generated/written, the KRM schemas align correctly with the GCP API contract.

### [2026-08-09] Implementing Direct Controller and E2E Fixtures for LicenseManagerConfiguration
- **Context**: Implementing the direct controller, mappers, fuzzers, and E2E test fixtures for `LicenseManagerConfiguration` under issue #12277.
- **Problem**: Encountered two critical GCP API constraints:
  1. The `display_name` field has a strict validation regex `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`. Standard test names with spaces or uppercase letters are rejected.
  2. The `nextBillingInfo` structure must have matching `userCount` with `currentBillingInfo` at creation and update time, or the API silently overrides or restricts them, causing unexpected diffs and infinite re-reconciliation loop.
- **Solution**: 
  1. Updated the minimal and maximal test fixtures' display names to use the hyphenated, all-lowercase format.
  2. Aligned `currentBillingInfo` and `nextBillingInfo` user count fields to match each other inside our test YAMLs.
  3. Added proper `startTime` field normalizations for both KRM objects and HTTP logs in `tests/e2e/normalize.go` and `tests/e2e/httplog.go`.
- **Impact**: All minimal and maximal E2E test fixtures record and run successfully against real GCP with clean, stable golden files.
