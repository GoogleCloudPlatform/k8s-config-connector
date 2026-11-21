### [2026-11-06] FinancialServicesBacktestResult MockGCP and Alignment
- **Context**: Implementing Greenfield MockGCP and Alignment for `FinancialServicesBacktestResult` (#15092).
- **Action**:
  1. Implemented MockGCP CRUD operations and LRO handling for `BacktestResult` in `mockgcp/mockfinancialservices/financialservicesbacktestresult.go`.
  2. Added resource name parsing for `backtestResultName` and updated instance name parsing in `mockgcp/mockfinancialservices/names.go`.
  3. Scoped `Previsit` normalization in `mockgcp/mockfinancialservices/normalize.go` to `financialservices.googleapis.com`.
  4. Registered `FinancialServicesBacktestResult` in `config/tests/samples/create/harness.go`.
  5. Implemented unit tests in `mockgcp/mockfinancialservices/financialservicesbacktestresult_test.go` covering full lifecycle and parent validation.
  6. Generated mock HTTP logs and golden object files for minimal and maximal fixtures via `hack/compare-mock`.
  7. Verified fixtures using `./dev/ci/presubmits/tests-e2e-fixtures-suite financialservices`.

### [2026-11-06] FinancialServicesDataset Direct Controller, Fixtures, and Fuzzer
- **Context**: Greenfield direct controller, E2E fixtures, and fuzzer for `FinancialServicesDataset` (#15091).
- **Action**:
  1. Verified `FinancialServicesDataset` direct controller (`pkg/controller/direct/financialservices/financialservicesdataset_controller.go`) and round-trip fuzzer (`pkg/controller/direct/financialservices/financialservicesdataset_fuzzer.generated.go`).
  2. Verified `pkg/controller/resourceconfig/static_config.go` with `generate_static_config.py`.
  3. Created minimal and maximal E2E fixtures under `pkg/test/resourcefixture/testdata/basic/financialservices/v1alpha1/financialservicesdataset/` referencing `parentRef`, `resourceID`, `dateRange`, `tableSpecs`, and `timeZone`.
  4. Ran `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha` to update alpha missing fields exceptions.
  5. Verified 200-seed fuzzer round-trip (`TestFocusedMappers/FinancialServicesDataset`), CRD schema tests, and direct controller compilation.
- **Observation**:
  - `financialservices.googleapis.com` (Financial Services / Anti-Money Laundering API) is an Early Access / allowlist-restricted Google Cloud service requiring `servicemanagement.services.bind` permissions for enablement; recording golden logs against real GCP is blocked in non-allowlisted projects.

### [2026-11-06] FinancialServicesModel Direct Controller, Fixtures, and Fuzzer
- **Context**: Greenfield direct controller, E2E fixtures, and fuzzer for `FinancialServicesModel` (#15055).
- **Action**:
  1. Verified `FinancialServicesModel` direct controller (`pkg/controller/direct/financialservices/financialservicesmodel_controller.go`) and round-trip fuzzer (`pkg/controller/direct/financialservices/financialservicesmodel_fuzzer.generated.go`).
  2. Verified `pkg/controller/resourceconfig/static_config.go` mapping (`FinancialServicesModel` mapped to Direct controller).
  3. Created minimal and maximal E2E fixtures under `pkg/test/resourcefixture/testdata/basic/financialservices/v1alpha1/financialservicesmodel/` referencing `parentRef`, `engineConfigRef` (external), `primaryDatasetRef` (external), and `endTime`.
  4. Ran `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha` to update alpha missing fields exceptions.
  5. Verified 200-seed fuzzer round-trip (`TestFocusedMappers/FinancialServicesModel`), CRD schema tests, and direct controller compilation.

### [2026-11-05] FinancialServicesPredictionResult MockGCP and Alignment
- **Context**: Implementing Greenfield MockGCP and Alignment for `FinancialServicesPredictionResult` (#15017).
- **Action**:
  1. Implemented MockGCP CRUD handlers (`GetPredictionResult`, `CreatePredictionResult`, `UpdatePredictionResult`, `DeletePredictionResult`) for `pb.PredictionResult` with LRO operations in `mockgcp/mockfinancialservices/financialservicespredictionresult.go`.
  2. Implemented `predictionResultName` and `parsePredictionResultName` in `mockgcp/mockfinancialservices/names.go`.
  3. Registered `FinancialServicesPredictionResult` in `config/tests/samples/create/harness.go` under MockGCP supported GroupKinds.
  4. Executed `WRITE_GOLDEN_OUTPUT=1 hack/compare-mock` across `financialservicespredictionresult-minimal` and `financialservicespredictionresult-maximal` to generate `_http_mock.log` and align object goldens.
  5. Verified `./dev/ci/presubmits/tests-e2e-fixtures-suite financialservices` passed for all 4 tests with 0 errors.
- **Observation**:
  - `financialservices.googleapis.com` is an allowlist-restricted service requiring `servicemanagement.services.bind` on GCP projects, so testing is conducted hermetically against MockGCP.

### [2026-11-05] FinancialServicesEngineConfig Direct Controller, Fixtures, and Fuzzer
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `FinancialServicesEngineConfig` (#14967).
- **Action**:
  1. Verified existing direct controller `financialservicesengineconfig_controller.go` and fuzzer `financialservicesengineconfig_fuzzer.generated.go`.
  2. Updated `financialservicesengineconfig_controller.go` to ensure `common.NormalizeGreenfieldReferences` is called prior to `GetIdentity` in `AdapterForObject`.
  3. Created minimal and maximal E2E fixtures under `pkg/test/resourcefixture/testdata/basic/financialservices/v1alpha1/financialservicesengineconfig/` covering all `FinancialServicesEngineConfigSpec` fields (`parentRef`, `resourceID`, `engineVersion`, `tuning`, `performanceTarget`, `hyperparameterSourceType`, `hyperparameterSource`).
  4. Verified fuzzer tests (`FOCUS=FinancialServicesEngineConfig go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers`), apichecks tests (`TestCRDFieldPresenceInTestsForAlpha`), and updated `alpha-missingfields.txt`.
- **Observation**:
  - `financialservices.googleapis.com` remains an Early Access / allowlist-restricted Google Cloud service requiring `servicemanagement.services.bind` permissions on real GCP projects. Fixtures and golden files were modeled according to GAPIC client protocol and REST endpoints.

### [2026-11-05] FinancialServicesDataset Greenfield Types and Identity Finalization
- **Context**: Triaging judgement queue and finalizing KRM types and identity for `FinancialServicesDataset` (#14961)
- **Action**:
  1. Updated `FinancialServicesDatasetSpec` to model `ParentRef *FinancialServicesInstanceRef` as required and immutable, replacing `ProjectRef` and `Location`.
  2. Removed top-level `Labels` from `FinancialServicesDatasetSpec` (reconciled via `metadata.labels`).
  3. Ran `apis/financialservices/generate.sh` to update types, `FinancialServicesDatasetIdentity`, mappers, and CRDs.
  4. Resolved all `FinancialServicesDataset` entries in `apis/financialservices/judgement_queue.yaml` (`untriaged-bulk-generation`, `identity-root-unknown`, `identity-server-generated-id`, `.spec.instanceRef`, `.spec.location`, `.spec.tableSpecs`).
  5. Regenerated controller and fuzzer with one-shot `controllerbuilder` commands.
  6. Verified fuzzer round-trip (`TestFocusedMappers/FinancialServicesDataset`), `apichecks` tests, `gcpurls`, and package unit tests.

### [2026-11-05] FinancialServicesModel Greenfield Types Judgement and Identity Finalization
- **Context**: Triaging judgement queue and finalizing direct KRM types and identity for `FinancialServicesModel` (#14928)
- **Action**:
  1. Replaced `ProjectRef`, `Location`, and guessed `InstanceRef` on `FinancialServicesModelSpec` with `ParentRef *FinancialServicesInstanceRef` (`spec.parentRef`).
  2. Converted `EngineConfig` and `PrimaryDataset` to `EngineConfigRef *FinancialServicesEngineConfigRef` and `PrimaryDatasetRef *FinancialServicesDatasetRef`.
  3. Removed top-level `Labels` from `FinancialServicesModelSpec`.
  4. Resolved all open entries for `FinancialServicesModel` in `apis/financialservices/judgement_queue.yaml`.
  5. Regenerated code via `./apis/financialservices/generate.sh` and updated controller/fuzzer via `controllerbuilder`.

### [2026-11-05] FinancialServicesPredictionResult Direct Controller, Fixtures, and Fuzzer
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `FinancialServicesPredictionResult` (#15012).
- **Action**:
  1. Verified direct controller in `pkg/controller/direct/financialservices/financialservicespredictionresult_controller.go` adheres to canonical greenfield invariants (REST AMLClient, reference normalization, CompareGreenfieldSpec, ReportDiff, LRO polling, status updates).
  2. Verified round-trip KRM fuzzer tests in `pkg/controller/direct/financialservices/financialservicespredictionresult_fuzzer.generated.go` pass via `FOCUS=FinancialServicesPredictionResult go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers`.
  3. Created minimal and maximal E2E test fixtures under `pkg/test/resourcefixture/testdata/basic/financialservices/v1alpha1/financialservicespredictionresult/` covering all `FinancialServicesPredictionResultSpec` fields (`parentRef`, `resourceID`, `datasetRef`, `modelRef`, `endTime`, `predictionPeriods`, `outputs.predictionDestination`, `outputs.explainabilityDestination`).
  4. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` to remove all missing field entries for `FinancialServicesPredictionResult`.
  5. Verified `financialservices.cnrm.cloud.google.com/FinancialServicesPredictionResult` is registered as Direct in `static_config.go` and remains in `BLOCKED_KINDS`.
- **Observation**:
  - `financialservices.googleapis.com` (Financial Services / Anti-Money Laundering API) is an Early Access / allowlist-restricted Google Cloud service requiring `servicemanagement.services.bind` permissions for enablement.

### [2026-11-05] FinancialServicesPredictionResult Types Triage
- **Context**: Triaging judgement queue and finalizing KRM types and identity for `FinancialServicesPredictionResult` (#14960).
- **Action**:
  1. Resolved `judgement_queue.yaml` entries for `FinancialServicesPredictionResult`.
  2. Modeled `FinancialServicesPredictionResult` as a child resource of `FinancialServicesInstance` using immutable `spec.parentRef` (`*FinancialServicesInstanceRef`), removing `ProjectRef`, `Location`, and guessed `InstanceRef`.
  3. Converted `spec.dataset` to `DatasetRef` (`*FinancialServicesDatasetRef`) and `spec.model` to `ModelRef` (`*FinancialServicesModelRef`).
  4. Retained `tableURI` as a string in `PredictionResult_Outputs` BigQueryDestination since `bq://` URIs are not KCC resource references.
  5. Removed top-level `spec.labels` as labels are managed via `metadata.labels`.
  6. Regenerated types, identity, mappers, CRDs, direct controller, and fuzzer. Verified with tests.

### [2026-11-05] FinancialServicesBacktestResult Direct Controller, Fixtures, and Fuzzer
- **Context**: Greenfield direct controller, E2E fixtures, and fuzzer for `FinancialServicesBacktestResult` (#14968).
- **Action**:
  1. Verified `FinancialServicesBacktestResult` direct controller (`pkg/controller/direct/financialservices/financialservicesbacktestresult_controller.go`) and round-trip fuzzer (`pkg/controller/direct/financialservices/financialservicesbacktestresult_fuzzer.generated.go`).
  2. Generated `pkg/controller/resourceconfig/static_config.go` with `generate_static_config.py`.
  3. Created minimal and maximal E2E fixtures under `pkg/test/resourcefixture/testdata/basic/financialservices/v1alpha1/financialservicesbacktestresult/` referencing `parentRef`, `datasetRef` (external), `modelRef` (external), `endTime`, `backtestPeriods`, and `performanceTarget`.
  4. Ran `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha` to update alpha missing fields exceptions.
  5. Verified 200-seed fuzzer round-trip (`TestFocusedMappers/FinancialServicesBacktestResult`), CRD schema tests, and direct controller compilation.

### [2026-11-04] FinancialServicesBacktestResult Types Triage
- **Context**: Triaging judgement queue and finalizing KRM types and identity for `FinancialServicesBacktestResult` (#14921).
- **Action**:
  1. Resolved `judgement_queue.yaml` entries for `FinancialServicesBacktestResult`.
  2. Modeled `FinancialServicesBacktestResult` as a child resource of `FinancialServicesInstance` using immutable `spec.parentRef` (`*FinancialServicesInstanceRef`), removing `ProjectRef`, `Location`, and guessed `InstanceRef`.
  3. Converted `spec.dataset` to `DatasetRef` (`*FinancialServicesDatasetRef`) and `spec.model` to `ModelRef` (`*FinancialServicesModelRef`).
  4. Removed top-level `spec.labels` as labels are managed via `metadata.labels`.
  5. Regenerated types, identity, mappers, CRDs, direct controller, and fuzzer. Verified with `validate-prereqs.sh` and tests.

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
