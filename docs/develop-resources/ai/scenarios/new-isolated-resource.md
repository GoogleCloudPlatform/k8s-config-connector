# Adding a New Resource (Isolated Mock Strategy)

This scenario applies when you need to implement a resource where the **GCP Go SDK does not yet have support** (e.g., alpha resources). You cannot use the standard Direct Controller pattern that relies on `google-cloud-go` SDK types. Instead, you must use the **Isolated Mock Strategy**.

## The Strategy

### The Problem
*   The GCP Service definitions exist in Protos (googleapis) but are not yet published in the Go SDK.
*   The KCC Direct Controller needs Go types to build the adapter.

### The Solution
Use the Go types generated *within* `mockgcp` as the source of truth for the controller.

## Recommended Workflow: The "Reordered" Stacked PRs

Unlike the standard [New Resource Workflow](./new-resource.md) (API -> Mock -> Controller), the Isolated Mock strategy requires you to **reorder** the PRs because the API generation depends on the Protos defined in the Mock.

### PR 1: MockGCP & Protos (The Foundation)
*   **Content**: `.proto` files, generated `pb.go` files, `mockgcp` service implementation, registration.
*   **Why**: `controllerbuilder` needs these protos to generate the API in the next step.
*   **Gemini Prompt**:
    > "I need to implement the `<Kind>` resource for the `<Service>` CIG.
    >
    > **Constraints:**
    > 1.  **Isolated Mock Strategy:** The GCP Go SDK does not support this resource yet. We must use `mockgcp` as the source of truth. Defines the proto in `mockgcp/apis/mockgcp/...` and generate the Go types there.
    > 2.  **Stacked PRs:** I want to implement this in 3 stacked PRs.
    >
    > **Goal:** Let's start with **PR 1 (MockGCP)**.
    > Please check if the proto exists in `mockgcp/apis`. If not, tell me where to copy it from, add it to the Makefile, and generate the mock types."

### PR 2: API & Mapper (The Contract)
*   **Content**: KRM API definitions (`types.go`), `controllerbuilder` config, generated mappers (`mapper.generated.go`), and manual mappers (`*_mappings.go`).
*   **Configuration**: You will configure `controllerbuilder` to read the protos from `mockgcp` (created in PR #1).
*   **Gemini Prompt** (after PR 1 is ready):
    > "Now for **PR 2 (API & Mapper)**:
    > 1.  Generate the KRM API for `<Kind>`.
    > 2.  Configure `controllerbuilder` to use the `mockgcp` proto we just added.
    > 3.  Generate the Mappers."

### PR 3: Controller & Tests (The Logic)
*   **Content**: `*_controller.go`, `create.yaml` (E2E tests), Release Notes.
*   **Same as Standard**: Implementation logic and verification.
*   **Gemini Prompt** (after PR 2 is ready):
    > "Now for **PR 3 (Controller & Tests)**:
    > 1.  Implement the Direct Controller using the `mockgcp` types.
    > 2.  Add E2E fixture tests verifying Create, Read, Update, Delete."

## Special Case: Adding a New Field (Isolated Mock)

**Prompt:**
> "I need to add the `<Field>` to `<Kind>` in KCC.
>
> **Context:**
> *   This resource uses the **Isolated Mock Strategy** (Direct Controller based on `mockgcp` types).
> *   The proto source is in `mockgcp/apis/mockgcp/...`.
>
> **Goal:**
> 1.  Update the `mockgcp` proto definition to include `<Field>`.
> 2.  Regenerate the mock types.
> 3.  Update the KRM API (`types.go`) and run `controllerbuilder` to regenerate mappers.
> 4.  Update the `create.yaml` test case to verify this new field."
