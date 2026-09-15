---
name: kcc-direct-controller-logic-implementer
description: Implement the core reconciliation logic (Adapter) and E2E fixtures for a direct KCC resource. Use this after Step 1 is complete.
---

# KCC Direct Controller Logic Implementer

This skill guides the implementation of the `Adapter` interface and the creation of "Minimal" and "Maximal" E2E fixtures to verify the resource against real GCP.

## Inputs
- `resource_kind`: The KCC Kind.
- `service_name`: The GCP service name (short, e.g., `apigee`).
- `api_version`: The KCC API version.

## Workflow

1.  **Implement Adapter Logic**:
    Update `pkg/controller/direct/<service>/<resource_lower>_controller.go`.
    - Implement `Find`, `Create`, `Update`, and `Delete`.
    - Use the generated mappers and manual mappers as needed.
    - Ensure correct error handling (e.g., handling 404s in `Find`).

2.  **Create Minimal Fixture**:
    Create directory `pkg/test/resourcefixture/testdata/basic/<service_name>/<api_version>/<resource_lower>/<resource_lower>-minimal/`.
    - Add `create.yaml`: Use the bare minimum **Required** fields.
    - Use `${uniqueId}` for resource names.

3.  **Create Maximal Fixture**:
    Create directory `pkg/test/resourcefixture/testdata/basic/<service_name>/<api_version>/<resource_lower>/<resource_lower>-maximal/`.
    - Add `create.yaml`: Include **every supported field** in the Spec.
    - Add `update.yaml`: Update all **mutable** fields.
    - Add `dependencies.yaml` if the resource requires other KCC resources to exist first.

4.  **Record Golden Files (Real GCP)**:
    You **MUST strictly follow** the [`record-real-gcp`](file:///usr/local/google/home/maqiuyujoyce/4-k8s-config-connector/.gemini/skills/record-real-gcp/SKILL.md) skill to discover affected test fixtures and record authentic GCP golden logs.

5.  **Verify Field Coverage**:
    Run the API check tests:
    - For alpha: `WRITE_GOLDEN_OUTPUT=1 go test -v ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha`
    - Verify that your "Maximal" test reduces the number of missing fields in the exceptions file.


## Journaling
Append any reconciliation hurdles, GCP SDK quirks, or other controller issues to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
