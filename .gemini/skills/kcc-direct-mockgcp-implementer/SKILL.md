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
- If the mock service already exists, implement the necessary CRUD (Create, Read, Update, Delete) methods for `<ResourceKind>` in `mockgcp/mock<service_name>/<kind_lowercase>.go` (and List, if the controller uses it).
- **Parents**: Take the parent from the request (`projects/{project}/locations/{location}` or the parent resource's name). If real GCP rejects a missing parent (e.g. with a 404 when the parent doesn't exist), the mock must too.
- **LROs**: Return long-running operations wherever real GCP does.

### 2b. Remove from Ratcheting Exclusions (MANDATORY)
Before running the test cases against real or mock GCP, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which is a fundamental use case KCC resources must support.
1. Open `tests/e2e/ratcheting.go`.
2. Locate the function `ShouldTestRereconiliation`.
3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

### 3. Incremental Mock Alignment

  > [!WARNING]
  > **WHENEVER A TEST CASE IS UPDATED, WE MUST RECORD REAL GCP LOGS AGAIN.**
  > If you make any modifications to a test case configuration or manifest files (such as `create.yaml`, `update.yaml`, or `dependencies.yaml`), or modify the controller's GCP request structures, you **MUST** run the test case against real GCP (`hack/record-gcp` or with `E2E_GCP_TARGET=real`) to regenerate the authentic `_http.log` baseline before comparing or committing any mock log changes. Do not attempt to manually edit the logs or bypass recording live traffic.
  >
  > **Do NOT re-run real GCP during MockGCP alignment**: When updating MockGCP implementations or normalizers, do NOT re-run `hack/record-gcp`. Mock alignment is performed exclusively using `hack/compare-mock`.
  >
  > **Real GCP logs are the baseline**: Every fixture directory must already have an `_http.log` recorded against real GCP. If one is missing, stop and report it instead of aligning against nothing. Never hand-edit `_http.log`, `_http_mock.log` or golden files. If you had to re-record a fixture, say which one and why in the PR.

Use the `match-mockgcp-with-realgcp` skill (at `.gemini/skills/match-mockgcp-with-realgcp/SKILL.md`) to align MockGCP behavior with real GCP golden logs for all fixtures under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/`.
  
Ensure all mock tests pass and both `_http.log` and `_http_mock.log` are present in all fixture directories with zero discrepancies.

### 4. Verify and Run Presubmits
Run these commands (the presubmit is what CI runs for the group, against MockGCP):
```bash
hack/compare-mock "fixtures/<kind_lowercase>"
go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/<kind_lowercase>"
./dev/ci/presubmits/tests-e2e-fixtures-suite <group>
scripts/validate-prereqs.sh
```
- `hack/compare-mock` must have no failures, and `TestGoldenLogAlignment` must pass with zero errors for every fixture.
- **Multi-Fixture Verification Requirement (MANDATORY)**:
  - You **MUST** verify and ensure that all test directories associated with your target resource are completely verified.
  1. Ensure that all discovered test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` have their golden files (`_generated_object_*.golden.yaml`) and mock HTTP logs (`_http_mock.log`) successfully generated. Note that a 0-diff in `git status` is perfectly valid if the resource doesn't contain any service-generated, randomized values in the responses or our normalizers have overwritten those values.
  2. Commit and push all regenerated and matched log/golden files for every discovered fixture in your Pull Request.

## Review checklist

Use this when reviewing a PR produced with this skill. Check every item:

1. **Mock service**: `<ResourceKind>` is implemented in `mockgcp/mock<service_name>/` (a new service is registered in `mockgcp/register.go`). CRUD handles the parent like real GCP, including rejecting a missing parent.
2. **Real-GCP baseline**: No `_http.log` changed, unless the PR explains a fixture or request change and the new log comes from `hack/record-gcp`. No log or golden file was edited by hand.
3. **Alignment and CI**: Every fixture has `_http_mock.log` and goldens. `TestGoldenLogAlignment/<kind_lowercase>` passes with zero errors, and `./dev/ci/presubmits/tests-e2e-fixtures-suite <group>` passes.
4. **Normalizers**: `Previsit` is scoped to `<service_name>.googleapis.com` URLs. API-version and URL-prefix replacements are exact strings for this service only, with no broad wildcards.
5. **Mock behavior and scope**: Defaults are set on both write and read paths. Proto3 zero-value gaps have `TODO`s. No legacy (Terraform/DCL) controller code changed, no new `missingfields.txt` lines, and `<ResourceKind>` isn't excluded in `ShouldTestRereconiliation`.
