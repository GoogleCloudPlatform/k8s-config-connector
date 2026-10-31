# Journal: Match Real GCP Behavior in MockGCP for ApigeeAddonsConfig

This journal captures the details and findings of verifying and aligning the `ApigeeAddonsConfig` mock in MockGCP (`mockgcp/mockapigee`) with authentic real GCP behavior.

## Observations

1. **Existing Mock Implementation**:
   - The Apigee mock service (`mockgcp/mockapigee`) already implemented `SetAddonsOrganization` and `GetOrganization`.
   - However, in test environments, the parent `Organization` resource (`organizations/${projectId}`) does not exist by default unless created or pre-seeded. In real GCP test projects, Apigee organizations take 24 hours to delete due to retention policies, so test projects use a pre-existing organization named after the GCP project (`organizations/${projectId}`).
   - `mockgcp/mockapigee/organization.go` was enhanced with `populateDefaultOrganization` so that `GetOrganization` and `SetAddonsOrganization` automatically initialize the pre-existing organization for the project with default fields matching authentic GCP responses (`state: ACTIVE`, `billingType: PAYG`, `runtimeType: CLOUD`, `apigeeProjectId`, `caCertificate`, `caCertificates`, and default properties). Deletions are tracked via `deletedOrgs` to ensure that deleted organizations properly return 404 `NOT_FOUND` on subsequent `GetOrganization` or `SetAddonsOrganization` calls instead of being re-populated.

2. **Real GCP Addons Behavior & PAYG Billing Type**:
   - In real GCP projects with PAYG billing, calling `:setAddons` with add-ons such as `advancedApiOpsConfig`, `integrationConfig`, `connectorsPlatformConfig`, or `monetizationConfig` enabled fails with `400 FAILED_PRECONDITION: add-ons are not supported for PAYG billing type`.
   - `apiSecurityConfig: { enabled: true }` is supported on PAYG organizations and successfully reconciles in real GCP.
   - When disabling/deleting an add-on, GCP returns `{}` for `addonsConfig` on subsequent `GET organizations/${projectId}` calls.
   - Handled `req.GetOrganization() == nil` and empty request bodies gracefully in `SetAddonsOrganization` to support resource deletion (`DELETE` in the Terraform provider calls `:setAddons` with an empty body `{}`).

3. **Ratcheting & Re-reconciliation**:
   - Removed `ApigeeAddonsConfig` from the ratcheting exclusion list in `tests/e2e/ratcheting.go`.
   - When configured with `apiSecurityConfig: { enabled: true }`, re-reconciliation passes cleanly without any extraneous diffs or write operations against either real GCP or MockGCP.

4. **Real GCP Recording**:
   - Ran `hack/record-gcp "fixtures/^apigeeaddonsconfig-basic$"` against real GCP project `cnrm-barni-4`.
   - Started at: Wed Oct 28 20:37:01 PDT 2026.
   - Finished at: Wed Oct 28 20:44:08 PDT 2026.
   - Baseline golden files (`_http.log` and `_generated_object_apigeeaddonsconfig-basic.golden.yaml`) recorded authentically and committed.

5. **Mock Alignment Verification**:
   - Verified that `hack/compare-mock "fixtures/^apigeeaddonsconfig-basic$"` passes cleanly with 0 diffs.
   - Verified that `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/apigee"` passes with 0 errors.
   - Verified `TestRealHTTPLogsDoNotContainMockGCP` and all `tests/apichecks` tests pass.
