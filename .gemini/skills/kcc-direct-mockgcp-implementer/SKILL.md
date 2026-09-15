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

### 3. Align Mock with Real GCP
- **MANDATORY**: You **MUST strictly follow** the [`match-mockgcp-with-realgcp`](.gemini/skills/match-mockgcp-with-realgcp/SKILL.md) skill to remove ratcheting exclusions, establish real GCP baselines (if needed), execute `hack/compare-mock`, and resolve any discrepancies across all fixtures.

### 4. Run Presubmits
- Run local validation: `scripts/validate-prereqs.sh`.
- Run the e2e fixtures presubmit: `./dev/ci/presubmits/tests-e2e-fixtures-<kind_lowercase>`.
