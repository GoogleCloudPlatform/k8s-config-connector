# Journal: AppEngineDomainMapping MockGCP Implementation

This journal captures the details and findings of implementing the `AppEngineDomainMapping` mock in MockGCP (`mockgcp/mockappengine`) and attempting real GCP recording.

## Observations

1. **Service Proto and Generation**:
   - GCP Service: `appengine.googleapis.com` (`google.appengine.v1.DomainMappings`).
   - Go Client Types: `cloud.google.com/go/appengine/apiv1/appenginepb`.
   - The mock service was registered in `mockgcp/mockappengine/service.go` and implemented in `mockgcp/mockappengine/domainmappings.go`.

2. **Real GCP Recording Attempt & Blocker**:
   - **Command Run**: `hack/record-gcp "fixtures/^appenginedomainmappingautogen$"`
   - **GCP Project Used**: `cnrm-barni-4`
   - **Error Message**:
     ```
     Caller is not authorized to administer the domain <domain_name>. If you own <domain_name>, you can obtain authorization by verifying ownership of the domain, or any of its parent domains, via the Webmaster Central portal: https://www.google.com/webmasters/verification/verification?domain=...
     ```
   - **Cause**: Real GCP App Engine Admin API enforces domain ownership verification via Webmaster Central / Search Console for custom domain mappings. Arbitrary subdomains/domain names created in automated end-to-end tests cannot be verified in live GCP without manual Webmaster Central TXT/DNS verification.
   - **Fallback Approach**: As required by the guidelines, the fallback approach was used: golden logs (`_http_mock.log`) and golden object definitions (`_generated_object_*.golden.yaml`) were generated and verified hermetically through MockGCP using `hack/compare-mock`.

3. **Re-reconciliation Testing**:
   - The legacy Terraform provider for `google_app_engine_domain_mapping` defines `override_strategy` with `Default: "STRICT"`.
   - Because `override_strategy` is a write-only creation parameter not returned by GET `DomainMappings.GetDomainMapping`, Terraform compares the GET state against the schema default and detects a diff on `override_strategy`, triggering a `PATCH` request during re-reconciliation.
   - `AppEngineDomainMapping` is retained in `tests/e2e/ratcheting.go` with an explanatory comment for this legacy Terraform provider behavior.

4. **Validation**:
   - Ran `hack/compare-mock "fixtures/^appenginedomainmappingautogen$"`: passed (reconciles to `Ready`, deletion succeeds).
   - Ran `go test -v ./tests/apichecks/...`: passed with zero errors.
