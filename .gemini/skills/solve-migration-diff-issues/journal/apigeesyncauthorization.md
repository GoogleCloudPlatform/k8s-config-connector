# Migration Journal: ApigeeSyncAuthorization

## Summary
Validation of the direct migration takeover for `ApigeeSyncAuthorization` (`apigee.cnrm.cloud.google.com/v1alpha1`) based on fixture `pkg/test/resourcefixture/testdata/basic/apigee/v1alpha1/apigeesyncauthorization/apigeesyncauthorization-basic`.

## Real GCP Recording Attempt & Limitation
1. **Command Run:**
   ```bash
   ./hack/record-gcp TestMigrationToDirect/fixtures/apigeesyncauthorization-basic$
   ```
2. **Project Used:** `cnrm-barni-4`
3. **Full Error Encountered:**
   ```
   googleapi: Error 400: GetSyncAuthorization is not allowed for organization with runtimeType CLOUD
   Details:
   [
     {
       "@type": "type.googleapis.com/google.rpc.RequestInfo",
       "requestId": "..."
     }
   ]
   ```
4. **Root Cause:**
   The GCP Apigee APIs (`GetSyncAuthorization` and `SetSyncAuthorization`) are only supported for Apigee Organizations configured with `runtimeType: HYBRID`. Organizations with `runtimeType: CLOUD` explicitly reject `GetSyncAuthorization` with HTTP 400 `FAILED_PRECONDITION`. Creating a `HYBRID` Apigee organization requires Apigee Hybrid entitlements/licensing not enabled on standard PAYG GCP projects (`runtimeType HYBRID is not supported for PAYG_V2, use CLOUD instead`).
5. **Fallback Approach:**
   Following skill instructions for allowlist-gated/restricted GCP services, validated migration hermetically against MockGCP. As documented in the skill guidelines, mock-generated migration logs are not committed as real GCP golden logs.

## Migration Test Framework Findings & Fix
- In `tests/e2e/migration_test.go`, Phase 3 (direct takeover) and Phase 4 (direct re-reconciliation) verify that the controller performs a 0-write no-op reconciliation by checking that incoming HTTP requests are read-only.
- `ApigeeSyncAuthorization` uses a POST HTTP method for its read endpoint (`POST /v1/organizations/{org}:getSyncAuthorization`). This is also documented in `tests/e2e/ratcheting.go`.
- `tests/e2e/migration_test.go` was updated to classify `POST` requests matching `:getSyncAuthorization` as read-only, matching the API semantics.

## MockGCP Validation
- Verified `TestMigrationToDirect/fixtures/apigeesyncauthorization-basic$` passes cleanly with 0 unexpected writes and 0 diffs on MockGCP.
- Verified `TestAllInSeries/fixtures/apigeesyncauthorization-basic$` passes cleanly on MockGCP.
