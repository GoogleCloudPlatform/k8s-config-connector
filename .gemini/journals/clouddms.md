### [2026-05-24] CloudDMSConversionWorkspace controller implementation
- **Context**: Implemented direct controller and MockGCP support for `CloudDMSConversionWorkspace` (Issue #8638 / PR #8639).
- **Problem**: The resource returns `latestCommitTime` in MockGCP responses which is dynamic and causes golden E2E file comparison failures.
- **Solution**: Added normalization for `.latestCommitTime` and `.status.observedState.latestCommitTime` in `mockgcp/mockclouddms/normalize.go` and `tests/e2e/normalize.go` to standardise volatile fields to `mockgcpregistry.PlaceholderTimestamp`.
- **Impact**: E2E tests against MockGCP run hermetically, stably, and pass 100% cleanly.
