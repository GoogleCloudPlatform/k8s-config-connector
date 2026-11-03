# Migration Journal: AccessContextManagerGCPUserAccessBinding

## Summary
Validation of direct migration takeover for `AccessContextManagerGCPUserAccessBinding` (`accesscontextmanager.cnrm.cloud.google.com/v1alpha1`) using fixture `pkg/test/resourcefixture/testdata/basic/accesscontextmanager/v1alpha1/accesscontextmanagergcpuseraccessbinding/accesscontextmanagergcpuseraccessbindingautogen`.

## Real GCP Recording Attempt
- **Command Run:**
  ```bash
  ./hack/record-gcp "TestMigrationToDirect/fixtures/accesscontextmanagergcpuseraccessbindingautogen$"
  ```
- **GCP Project:** `cnrm-barni-4`
- **Result:** Failed during Phase 1 creation due to missing organization-level permissions on live GCP.
- **Error Encountered:**
  ```text
  summary: Error creating GcpUserAccessBinding: googleapi: Error 403: The caller does not have permission
  ```
  and
  ```text
  summary: Error when reading or editing AccessContextManagerAccessLevel "accessPolicies/578359180191/accessLevels/...": googleapi: Error 403: The caller does not have permission
  ```
- **Root Cause:**
  `AccessContextManagerGCPUserAccessBinding` (`organizations/{org}/gcpUserAccessBindings/{bindingId}`) and `AccessContextManagerAccessLevel` are organization-scoped resources requiring organization-level Access Context Manager permissions (such as `roles/accesscontextmanager.policyAdmin` on `organizations/128653134652` and access policy `578359180191`). The service account credentials configured in the test environment (`overseer-kcc-tester@cnrm-barni-4.iam.gserviceaccount.com`) only possess project-level permissions within project `cnrm-barni-4` and lack organization-level Access Context Manager administrator permissions.

## Fallback Approach & MockGCP Validation
Because recording against live GCP is blocked by organization-level credential/permission constraints on the GCP organization, migration validation was executed hermetically against MockGCP:
- **Migration Test Command:**
  ```bash
  RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run "TestMigrationToDirect/fixtures/accesscontextmanagergcpuseraccessbindingautogen$"
  ```
- **Series E2E Test Command:**
  ```bash
  ./hack/compare-mock pkg/test/resourcefixture/testdata/basic/accesscontextmanager/v1alpha1/accesscontextmanagergcpuseraccessbinding/accesscontextmanagergcpuseraccessbindingautogen
  ```

### Verification Results
1. **Phase 1 (Legacy Create):** Successfully provisions the `AccessContextManagerAccessLevel` dependency and `AccessContextManagerGCPUserAccessBinding` resource using the legacy Terraform controller.
2. **Phase 2 (Legacy Re-reconcile):** Re-reconciles under the legacy controller with a no-op annotation (`re-reconcile-legacy-v1`) with 0 unexpected writes or modifications.
3. **Phase 3 (Direct Takeover):** Direct controller adopts the existing resource cleanly using read-only GET requests, achieving a perfect **0-write no-op reconciliation** with zero takeover diffs.
4. **Phase 4 (Direct Re-reconcile):** The Direct controller re-reconciles cleanly with a no-op annotation (`re-reconcile-direct-v1`), performing only read-only GET requests and 0 writes.
5. **Teardown & Deletion:** Both `AccessContextManagerGCPUserAccessBinding` and `AccessContextManagerAccessLevel` resources delete cleanly and completely.

## Ratcheting Exclusions
- Confirmed `AccessContextManagerGCPUserAccessBinding` is not present in `tests/e2e/ratcheting.go`.
