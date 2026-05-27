# Fix Flaky Test

This skill collects patterns and approaches to fixing flaky tests in Config Connector.

## mockgcp Flakiness
- **Infinite Reconciles:** If an e2e test flakes heavily on GitHub Actions but not locally, and uses `mockgcp`, check the `mockgcp` API implementation for the resource. Ensure `GetInstance` does not overwrite fields or populate `UpdateTime` with `time.Now()` on every read. This causes constant state drift and triggers infinite reconciliations, which can overload `envtest` and cause context cancellations.
- **Storage Updates:** `storage.Update` in mockgcp completely replaces the object. Update operations must always `Get`, modify, then `Update`.
