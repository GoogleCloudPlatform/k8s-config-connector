# Migration Journal: NetworkServicesEdgeCacheService

## Summary
Validation of direct migration takeover for `NetworkServicesEdgeCacheService` (`networkservices.cnrm.cloud.google.com/v1alpha1`) using fixture `pkg/test/resourcefixture/testdata/basic/networkservices/v1alpha1/networkservicesedgecacheservice/networkservicesedgecacheservicebasic`.

## Real GCP Recording Attempt
- **Command Run:**
  ```bash
  ./hack/record-gcp "TestMigrationToDirect/fixtures/networkservicesedgecacheservicebasic$"
  ```
- **GCP Project:** `cnrm-barni-4`
- **Result:** Failed during Phase 1 dependency creation due to live GCP API allowlisting gating.
- **Error Encountered:**
  ```text
  Update call failed: error applying desired state: summary: Error creating EdgeCacheOrigin: googleapi: Error 404: Method not found.
  ```
- **Root Cause:**
  Google Cloud Media CDN (Edge Cache) resources (`edgeCacheServices`, `edgeCacheOrigins`, `edgeCacheKeysets`) are gated behind project allowlisting. In projects that have not been explicitly allowlisted for Media CDN, `networkservices.googleapis.com` returns `404 Method not found` for all edge cache endpoints. Additionally, running `gcloud iam list-testable-permissions //compute.googleapis.com/projects/cnrm-barni-4 | grep -i edgecache` returns no permissions, confirming the project is not allowlisted for Media CDN.

## Fallback Approach & MockGCP Validation
Because recording against live GCP is blocked due to the project allowlisting requirement for Media CDN, migration validation was executed hermetically against MockGCP:
- **Migration Test Command:**
  ```bash
  RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run "TestMigrationToDirect/fixtures/networkservicesedgecacheservicebasic$"
  ```
- **Series E2E Test Command:**
  ```bash
  RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run "TestAllInSeries/fixtures/networkservicesedgecacheservicebasic$"
  ```

### Verification Results
1. **Phase 1 (Legacy Create):** Successfully provisions the `NetworkServicesEdgeCacheOrigin` dependency and the `NetworkServicesEdgeCacheService` resource using the Terraform controller.
2. **Phase 2 (Legacy Re-reconcile):** Re-reconciles under the legacy controller with a no-op annotation without unexpected writes or modifications.
3. **Phase 3 (Direct Takeover):** The Direct controller adopts the existing resource cleanly with read-only GET requests, achieving a perfect **0-write no-op reconciliation** with zero takeover diffs (`"isNewObject": false` count is 0 in `_migration_diffs.json`).
4. **Phase 4 (Direct Re-reconcile):** The Direct controller re-reconciles cleanly with a no-op annotation, performing only read-only GET requests and 0 writes.
5. **Deletion:** Both `NetworkServicesEdgeCacheService` and `NetworkServicesEdgeCacheOrigin` delete cleanly.

## Ratcheting Exclusions
- Verified `NetworkServicesEdgeCacheService` is not present in `tests/e2e/ratcheting.go`.
