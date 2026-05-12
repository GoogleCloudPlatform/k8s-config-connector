---
name: kcc-direct-controller-implementer
description: Implement the controller, mappers, and fuzzer for a direct KCC resource, ensuring package isolation and CI compliance. Use this when implementing the main reconciliation logic for a "direct" resource.
---

# KCC Direct Controller Implementer

This skill guides the implementation of the controller, mappers, and fuzzer for direct KCC resources, with a focus on package isolation to prevent symbol collisions and rigorous CI validation.

## Inputs
- `resource_kind`: The KCC Kind.
- `package_path`: The isolated package directory (e.g., `pkg/controller/direct/vertexai/examplestore/`).
- `proto_package`: The GCP proto package (e.g., `google.cloud.aiplatform.v1`).

## Workflow

1.  **Package Isolation**:
    You MUST implement all controller-related logic in the provided `package_path`. This prevents symbol collisions in `mapper.generated.go`.

2.  **Mappers**:
    Implement/verify `mapper.generated.go` and manual mappers. Ensure all references use the standard `Ref` pattern.

3.  **Fuzzer**:
    Implement `<resource_lower>_fuzzer.go` and register it with `fuzztesting.RegisterKRMFuzzer`.

4.  **Final Generation & Reporting**:
    Run `make ready-pr` from the repository root. This is a critical step that:
    - Regenerates all Go clients in `pkg/clients/generated`.
    - Updates global CRD reports (`docs/reports/crd_report.md` and `.csv`).
    - Runs `make fmt` and `make vet`.
    YOU MUST COMMIT ALL RESULTING CHANGES.

5.  **Validation & Last-Mile Tests**:
    Run the following tests to ensure CI compliance:
    - **Fuzzing**: `dev/ci/presubmits/fuzz-roundtrippers`
    - **E2E Scaffolding**: `dev/ci/presubmits/tests-e2e-fixtures-direct`
    - **Schema Integrity**: `go test ./pkg/crd/template/...`
    - **API Field Coverage**: `go test ./tests/apichecks/...`. 
      - If `TestCRDFieldPresenceInTestsForAlpha` fails, you may need to regenerate the exceptions file by running it with `WRITE_GOLDEN_OUTPUT=1`.

## Journaling
Append any controller/reconciliation quirks or API check hurdles to `.gemini/skills/kcc-direct-controller-implementer/journal.md`.
