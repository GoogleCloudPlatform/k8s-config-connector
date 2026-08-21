# Cloud Tasks Journal

### 2026-06-01 MockGCP CreateQueue Custom Values Support
- **Context**: Implementing the direct controller E2E fixtures for `TasksQueue` (Issue #8911).
- **Problem**: When testing `tasksqueue-maximal` against MockGCP, the mock implementation of `CreateQueue` was completely overwriting any incoming `RateLimits` and `RetryConfig` with hardcoded default values, which caused immediate drift and reconciliation loops.
- **Solution**: Refactored `mockgcp/mockcloudtasks/queue.go`'s `CreateQueue` to check for the presence of individual subfields of `RateLimits` and `RetryConfig` (such as `MaxBurstSize`, `MaxAttempts`, etc.) and only apply the default values if those subfields are unconfigured/omitted (set to zero or nil).
- **Impact**: E2E test fixtures can now reliably configure custom values on creation of TasksQueue without incurring artificial drift against MockGCP, matches real GCP behavior where supplied parameters are respected.
