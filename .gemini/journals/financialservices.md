### [2026-10-08] FinancialServicesInstance MockGCP and Alignment
- **Context**: Implementing Greenfield MockGCP and Alignment for `FinancialServicesInstance` (#13837)
- **Action**:
  1. Generated grpc-gateway bindings for `google.cloud.financialservices.v1` service proto under `mockgcp/generated/google/cloud/financialservices/v1/`.
  2. Implemented mock service `mockgcp/mockfinancialservices/` with CRUD and LRO operations for `Instance`.
  3. Registered `mockfinancialservices` in `mockgcp/register.go` and allowed `FinancialServicesInstance` in `config/tests/samples/create/harness.go`.
  4. Executed `hack/compare-mock` across both minimal and maximal fixtures to align mock logs and golden files.
  5. Verified all e2e tests and unit tests pass.

### [2026-10-02] FinancialServicesInstance Direct Controller, Fixtures, and Fuzzer
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `FinancialServicesInstance` (#13673)
- **Action**:
  1. Updated `apis/financialservices/generate.sh` to generate mappers and ran controller generation.
  2. Implemented `FinancialServicesInstance` direct controller in `pkg/controller/direct/financialservices/financialservicesinstance_controller.go` using GAPIC REST client (`NewAMLRESTClient`).
  3. Implemented KRM round-trip fuzzer in `pkg/controller/direct/financialservices/financialservicesinstance_fuzzer.go` and registered `financialservices` in `pkg/controller/direct/register/register.go`.
  4. Registered `FinancialServicesInstance` direct reconciler in `pkg/controller/resourceconfig/static_config.go`.
  5. Created minimal and maximal E2E fixtures covering all spec fields (`projectRef`, `location`, `resourceID`, `kmsKeyRef`, `labels`) under `pkg/test/resourcefixture/testdata/basic/financialservices/v1alpha1/financialservicesinstance/`.
  6. Successfully verified fuzzer round-trip tests (`TestFocusedMappers/FinancialServicesInstance`), CRD schema tests, and alpha API field coverage checks (`TestCRDFieldPresenceInTestsForAlpha`).
- **Observation**:
  - `financialservices.googleapis.com` (Financial Services / Anti-Money Laundering API) is an Early Access / allowlist-restricted Google Cloud service requiring `servicemanagement.services.bind` permissions for enablement.
