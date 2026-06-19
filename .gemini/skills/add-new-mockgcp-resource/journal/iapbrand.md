# IAPBrand MockGCP Match Journal

## 2026-06-19

*   Enabled `iap.googleapis.com` service in MockGCP's `mockserviceusage/knownservices.go`. Without this, the test framework's serviceusage client fails to enable `iap.googleapis.com` with a 403 error.
*   Registered `iap.cnrm.cloud.google.com/IAPBrand` under `config/tests/samples/create/harness.go` to mark it as supported by MockGCP.
*   Generated and verified the E2E mock golden files (`_http.log`, `_generated_object_iapbrand.golden.yaml`, and `_identities.yaml`) for the `iapbrand` test fixture, completing the match with MockGCP.
