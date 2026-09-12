---
name: kcc-direct-mockgcp-implementer
description: Guides the implementation of Phase 3 (MockGCP and Alignment) for a direct KCC resource, verifying behavioral correctness against simulated GCP services. Use this when you need to implement or align mockgcp for a KCC resource.
---

# KCC Direct MockGCP Implementer

This skill guides you through implementing Phase 3 (MockGCP and Alignment) for a direct KCC resource to verify behavioral correctness against simulated GCP services.

## Inputs
- `ResourceKind`: The kind of the resource (e.g., `VertexAIDataset`).
- `service_name`: The short name of the GCP service (e.g., `aiplatform`).
- `api_version`: The KCC API version (e.g., `v1alpha1`, `v1beta1`).
- `group`: The API group (e.g., `vertexai`).
- `kind_lowercase`: The lowercase kind name (e.g., `vertexaidataset`).
- `testname`: The specific test folder name under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/`.

## Workflow

### 1. Locate E2E Fixtures
- The test fixtures are located under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/`.

### 2. Add or Enhance Mock Service
- If a mock service for `<service_name>` does not exist under `mockgcp/mock<service_name>/`, create one:
  - Follow the guide in `mockgcp/GEMINI.md` and `mockgcp/README.md`.
  - Add the relevant proto to the Makefile and run `make gen-proto` if needed.
  - Implement the mock service entrypoint in `mockgcp/mock<service_name>/service.go` and register it in `mockgcp/register.go`.
- If the mock service already exists, implement the necessary CRUD (Create, Read, Update, Delete) methods for `<ResourceKind>` in `mockgcp/mock<service_name>/<kind_lowercase>.go`.

### 2b. Remove from Ratcheting Exclusions (MANDATORY)
Before running the test cases against real or mock GCP, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which is a fundamental use case KCC resources must support.
1. Open `tests/e2e/ratcheting.go`.
2. Locate the function `ShouldTestRereconiliation`.
3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

### 3. Incremental Mock Alignment

  > [!WARNING]
  > **WHENEVER A TEST CASE IS UPDATED, WE MUST RECORD REAL GCP LOGS AGAIN.**
  > If you make any modifications to a test case configuration, manifest files (such as `create.yaml`, `update.yaml`, or `dependencies.yaml`), or the controller's runtime mapping configuration, you **MUST** run the test case against real GCP (`hack/record-gcp` or with `E2E_GCP_TARGET=real`) to regenerate the authentic `_http.log` baseline before comparing or committing any mock log changes. Do not attempt to manually edit the logs or bypass recording live traffic.

Use the `match-mockgcp-with-realgcp` skill at .gemini/skills/match-mockgcp-with-realgcp/SKILL.md) to align MockGCP behavior with real GCP golden logs for all fixtures under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/`.
  
Ensure all mock tests pass and both `_http.log` and `_http_mock.log` are present in all fixture directories with zero discrepancies.

### 4. Verify and Run Presubmits
- Run local validation: `scripts/validate-prereqs.sh`.
- Run the e2e fixtures presubmit: `./dev/ci/presubmits/tests-e2e-fixtures-<kind_lowercase>`.
- **Multi-Fixture Verification Requirement (MANDATORY)**:
  - You **MUST** verify and ensure that all test directories associated with your target resource are completely verified.
  1. Ensure that all discovered test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` have their golden files (`_generated_object_*.golden.yaml`) and mock HTTP logs (`_http_mock.log`) successfully generated. Note that a 0-diff in `git status` is perfectly valid if the resource doesn't contain any service-generated, randomized values in the responses or our normalizers have overwritten those values.
  2. Commit and push all regenerated and matched log/golden files for every discovered fixture in your Pull Request.
