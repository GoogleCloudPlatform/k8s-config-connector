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

### 3. Incremental Mock Alignment
- Run `hack/compare-mock "fixtures/^<testname>$"` to execute the tests against the mock implementation.
- Use the `fix-diffs-mockgcp` skill (`mockgcp/.gemini/skills/fix-diffs-mockgcp/SKILL.md`) to align the mock logs with the real GCP output:
  - **Output-Only Fields/IDs**: If real GCP produces dynamic values that mockgcp lacks, implement a `populate<ResourceKind>Defaults` function in `mockgcp/mock<service_name>/<kind_lowercase>.go` called on `Insert` and `Get` to match the required format.
  - **Volatile/Random Values**: For values like timestamps or etags that are functionally identical but structurally unpredictable, update `normalize.go` for the service.
  - **Critical Rule**: Always scope the `Previsit` normalization in `normalize.go` to ensure it only applies to your service URL (e.g. `strings.Contains(event.URL(), "<service_name>.googleapis.com")`) to prevent log corruption in unrelated services.
- Iterate on running `hack/compare-mock "fixtures/^<testname>$"` and making incremental code updates until the HTTP logs match real GCP perfectly with clean, minimal diffs.
- Once the logs align, run the comparison with `WRITE_GOLDEN_OUTPUT=1` against mock GCP to generate the `_http_mock.log` file:
  ```bash
  WRITE_GOLDEN_OUTPUT=1 RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run "TestAllInSeries/fixtures/<testname>"
  ```
  Make sure both `_http.log` and `_http_mock.log` are present in the fixture directory.

### 4. Verify and Run Presubmits
- Run local validation: `scripts/validate-prereqs.sh`.
- Run the e2e fixtures presubmit: `./dev/ci/presubmits/tests-e2e-fixtures-<kind_lowercase>`.
- Make sure to stage and commit both `_http.log` and `_http_mock.log` files in your Pull Request.
