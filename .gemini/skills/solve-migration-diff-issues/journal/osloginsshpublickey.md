# Migration Journal: OSLoginSSHPublicKey

## Summary
Validation of the direct migration takeover for `OSLoginSSHPublicKey` (`oslogin.cnrm.cloud.google.com/v1alpha1`) based on fixture `pkg/test/resourcefixture/testdata/basic/oslogin/v1alpha1/osloginsshpublickey/osloginsshpublickey`.

## Real GCP Recording Attempt
- **Command Run:** `./hack/record-gcp TestMigrationToDirect/fixtures/osloginsshpublickey$`
- **GCP Project:** `cnrm-barni-4`
- **Result:** Failed during Phase 1 creation due to API authentication/credential constraint on live GCP.
- **Error Encountered:**
  ```text
  googleapi: Error 403: End user credentials must match the user specified in the request. Request for user [test-user-...@example.com] does not match the credential for [overseer-kcc-tester@cnrm-barni-4.iam.gserviceaccount.com].
  ```
- **Root Cause:**
  Google Cloud OS Login API (`oslogin.googleapis.com/v1/users/{user}:importSshPublicKey`) enforces that the authenticated caller must be an end user credential (e.g. human user Google account or Workspace profile) matching the `{user}` path parameter. Automated test runners running with service account credentials (`overseer-kcc-tester@cnrm-barni-4.iam.gserviceaccount.com`) cannot manage OS Login profiles for arbitrary end users on live GCP.

## Fallback Approach & MockGCP Validation
Because recording against live GCP is blocked by the OS Login API's end-user credential constraint, migration validation was executed hermetically against MockGCP:
- **Migration Test Command:** `E2E_GCP_TARGET=mock RUN_TESTS="TestMigrationToDirect/fixtures/osloginsshpublickey$" dev/tasks/run-e2e`
- **Fixture Verification Command:** `./hack/compare-mock fixtures/osloginsshpublickey$`

### Verification Results
1. **Phase 1 (Legacy Create):** Successfully provisions the `OSLoginSSHPublicKey` resource using the legacy Terraform controller.
2. **Phase 2 (Legacy Re-reconciliation):** Re-reconciles cleanly with no-op updates and 0 unexpected writes.
3. **Phase 3 (Direct Takeover):** The Direct controller successfully identifies the resource using `status.fingerprint` resolved via `GetIdentity`. The `compareOSLoginSSHPublicKey` comparison logic finds no discrepancies between the desired spec and actual GCP state. The resource is adopted cleanly with a **0-write no-op reconciliation** (only read requests are made).
4. **Phase 4 (Direct Re-reconciliation):** Direct controller re-reconciles cleanly with 0 writes.
5. **Takeover Diffs:** Inspection of `_migration_diffs.json` confirms no `isNewObject: false` blocks exist, confirming zero diffs during direct takeover and direct re-reconciliation.
