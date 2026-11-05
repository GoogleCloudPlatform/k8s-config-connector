### [2026-06-12] LoggingLink Fuzzer Implementation
- **Context**: Implementing round-trip KRM fuzzer for LoggingLink under Issue #9789.
- **Problem**: The fuzzer file for LoggingLink was initially located at `logginglink_fuzzer.go` instead of the expected `link_fuzzer.go` (since the controller is named `link_controller.go`).
- **Solution**: Deleted `logginglink_fuzzer.go` and created a properly structured, fully documented KRM fuzzer at `pkg/controller/direct/logging/link_fuzzer.go`. Verified with central fuzz tests `go test -v -count=1 ./pkg/fuzztesting/fuzztests/...`.
- **Impact**: Ensures that any future schema changes to LoggingLink will be continuously validated through central round-trip fuzz testing and prevents file name clutter.

### [2026-06-15] LoggingLogMetric Direct Controller Verification
- **Context**: Verifying direct controller, KRM fuzzer, and E2E fixtures for `LoggingLogMetric`.
- **Problem**: Need to ensure that the pre-existing direct controller implementation and KRM fuzzer for `LoggingLogMetric` are fully compliant with direct controller design patterns, and that all 5 E2E fixtures are passing.
- **Solution**: Executed full E2E test runs against `mockgcp` for all 5 fixtures under `logginglogmetric` (explicitlogmetric, exponentiallogmetric, linearlogmetric, logbucketmetric, logginglogmetric-nometricdescriptor). All fixtures, schema validation tests, and global API field presence tests passed flawlessly. Additionally, executed complete fuzz round-trip tests to verify the `LoggingLogMetric` KRM fuzzer logic.
- **Impact**: Confirms that `LoggingLogMetric` has been completely migrated to a fully functional direct controller that is backward-compatible and structurally sound.

### [2026-10-23] LoggingLogScope Direct Controller and Fixtures Implementation
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `LoggingLogScope` under Issue #14312.
- **Problem**:
  1. The Google Cloud logging API defines LogScope under `ProjectsLocationsLogScopesService`, `FoldersLocationsLogScopesService`, and `OrganizationsLocationsLogScopesService` using Google API Client library `google.golang.org/api/logging/v2`.
  2. The PATCH API for log scopes requires an explicit `updateMask` parameter (e.g., `updateMask=description,resource_names`), returning HTTP 400 Bad Request if `update_mask` is empty or omitted.
  3. Re-reconciliation comparison with `diffs.GoogleAPI.Diff` returned field differences for `name` if `maskedActual.Name` was not aligned with `desired.Name`.
- **Solution**:
  1. Implemented `LoggingLogScopeAdapter` using the `google.golang.org/api/logging/v2` REST client, supporting Project, Folder, and Organization parents.
  2. In `Update`, mapped modified struct JSON field names (`description`, `resourceNames`) to `description,resource_names` in `updateMask`.
  3. Ensured `maskedActual.Name = desired.Name` in `compareScope`.
  4. Added `logginglogscope_fuzzer.go` using `NewKRMTypedFuzzer_NoProto`.
  5. Created minimal and maximal fixtures and recorded real GCP traffic (`RECORD_AUDIT_PROBE=1 ./hack/record-gcp`).
- **Impact**: Enables declarative management of GCP Cloud Logging LogScopes with KCC direct controller.

