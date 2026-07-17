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
    Run the tests against real GCP to record the traffic and object state. Ensure you use a sufficient timeout (e.g., 30-60 minutes) as GCP resource creation can be slow:
    ```bash
    # Run from the repository root
    RUN_E2E=1 \
    E2E_GCP_TARGET=real \
    E2E_KUBE_TARGET=envtest \
    GOLDEN_REQUEST_CHECKS=1 \
    GOLDEN_OBJECT_CHECKS=1 \
    WRITE_GOLDEN_OUTPUT=1 \
    go test -v ./tests/e2e \
      -timeout 60m \
      -run TestAllInSeries/fixtures/<resource_lower>-minimal
    ```
    Repeat for the `-maximal` fixture. Commit the resulting `_http.log` and `_generated_object_*.golden.yaml` files.

5.  **Verify Field Coverage**:
    Run the API check tests:
    - For alpha: `WRITE_GOLDEN_OUTPUT=1 go test -v ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha`
    - Verify that your "Maximal" test reduces the number of missing fields in the exceptions file.


## Journaling
Append any reconciliation hurdles, GCP SDK quirks, or other controller issues to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
