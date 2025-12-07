# AI-Assisted Workflow for KCC Resource Development

This guide introduces the Gemini-powered workflow for developing Config Connector resources. This approach accelerates development by automating the implementation details, allowing you to focus on providing the correct inputs and verifying the final results.

## Target Audience

This guide is for developers who: 
*   Have a good understanding of Kubernetes concepts. Here're some basic ideas: 
    * [Kuberetes Setup](https://kubernetes.io/docs/setup/)
    * [CRD and Controller](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
    * [API convention](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
*   Have an in-depth understanding about declarative friendly API. See https://google.aip.dev/news/2020-10.
*   Have an in-depth understanding about Kubernetes reconciliation mechanism.
*   Have a clear understanding of [what Config Connector is](https://docs.cloud.google.com/config-connector/docs).
*   Know the differences between Config Connector, Terraform, and DCL. 



## Development Environment

To effectively use the Gemini-powered workflow, ensure your environment is set up for efficiency:

### Tools
*   **VS Code** or **Antigravity** (antigravity.google): Recommended IDE. Install the Go extension.
*   **Gemini CLI**: Ensure you have the latest version of the internal `gemini-cli` tool.
    *   **Recommendation**: Use `gemini-3.0-pro` model for best results with code generation and reasoning.
*   **GitHub CLI (`gh`)**: Essential for managing Stacked PRs.
    *   Install: `sudo apt install gh` (or equivalent).
    *   Auth: `gh auth login`.

### Best Practices
*   **Terminal**: Keep multiple terminals open:
    1.  **Gemini Session**: For prompt interaction.
    2.  **Build/Test**: For running `make` and `go test`.
    3.  **Git/GH**: For branch management and PR creation.

## The Philosophy: You Prompt, Gemini Executes, You Verify

The core of this workflow is a partnership between you and the Gemini CLI.
-   **Your Role (The Supervisor):** Your primary responsibility is to initiate the process with a clear, scenario-based prompt and then to act as a reviewer. You will verify that the code and the behavior (as captured in the golden test logs) are correct and adhere to Config Connector's standards.
-   **Gemini's Role (The Implementer):** Gemini's job is to perform all the hands-on-keyboard work. It will generate APIs, mappers, fuzzers, controllers, mockgcp and test files, and it will run the necessary commands to record test data.

## How to Use This Guide

1.  **Find Your Scenario:** Navigate to the `scenarios` directory and find the document that matches your development task (e.g. adding a new resource, adding a field).
2.  **Use the Provided Prompt:** Each scenario document now begins with a canonical prompt. Copy this prompt and provide it to the Gemini CLI.
3.  **Understand the Process:** The scenario document outlines the steps that Gemini will take. Use this to follow along with its progress.
4.  **Verify the Results:** After Gemini completes the task, use the "How to Verify" section in the scenario document. This section will guide you on which files to review and will link to our `deep-dives` and `api-conventions` documents, which serve as the technical reference for what "good" looks like.

## Advanced Guidance

### Manual Mappers (When Autogen isn't Enough)

The `controllerbuilder` is powerful but has limits, especially with:
*   Complex nested structs.
*   "Lazy" pointers vs values mismatches.
*   Custom logic needed during conversion.

**Pattern**:
*   The generator looks for `*_mappings.go`.
*   If you implement `Function_FromProto` or `Function_ToProto` manually in that file, the generator *skips* generating it and uses yours.
*   **Tip**: Let the generator run once. If it produces bad code (or misses fields), copy the generated function to `*_mappings.go` and fix it manually. The generator will then respect your manual override.

### Testing Tips

*   **Mock Verification**:
    *   Use `E2E_GCP_TARGET=mock` to run tests against your in-process mock server.
    *   Fast feedback loop: `go test -v ./tests/e2e -run TestAllInSeries/../yourresource`
*   **Harness Whitelist**:
    *   If your test skips with "not supported", check `config/tests/samples/create/harness.go` and add your Kind to the allowlist.

## Scenarios Index

*   [**New Standard Resource**](./scenarios/new-resource.md): Standard flow for resources supported by `google-cloud-go` SDK.
*   [**New Isolated Mock Resource**](./scenarios/new-isolated-resource.md): **Start here** if the resource is NOT in the Go SDK yet (e.g., Alpha resources).
*   [**New Field**](./scenarios/new-field.md): Adding a field to an existing resource.
*   [**Promote Alpha to Beta**](./scenarios/alpha-to-beta.md): Promoting a resource version.

## Quick Start Prompts

Use these prompts to get started immediately. For full details, see the [Scenarios Index](#scenarios-index).

### New Isolated Mock Resource
**Goal**: Implement a resource not yet in the Go SDK (Alpha/Private).
**Prompt**:
> "I need to implement the `<Kind>` resource for the `<Service>` CIG.
>
> **Constraints:**
> 1.  **Isolated Mock Strategy:** The GCP Go SDK does not support this resource yet. We must use `mockgcp` as the source of truth.
> 2.  **Stacked PRs:** I want to implement this in 3 stacked PRs.
>
> **Goal:** Let's start with **PR 1 (MockGCP)**.
> Please check if the proto exists in `mockgcp/apis`. If not, download it from googleapis/master, add it to the Makefile, and generate the mock types."

### Fixing Build & Linter Issues
**Goal**: Resolve common CI failures (missing DeepCopy, license headers, stale protos).
**Prompt**:
> "I need to fix build and linter issues in the current PRs.
> 1. Check for missing `DeepCopy` methods (run `make generate` to fix).
> 2. Check for missing proto files in `mockgcp` (fetch from `googleapis` or generate).
> 3. Ensure `mockgcp` service registration is complete (uncomment `NewHTTPMux`).
> 4. Run `make fmt` to fix license headers and formatting.
> 5. Verify everything compiles with `go build ./mockgcp/...`."

## Essential Manual Steps (The "Gotchas")
While Gemini handles most tasks, some steps might require manual intervention or specific instructions:

1.  **Missing Protos**: If `googleapis` repo is pinned to an older version, you might need to `wget` the `.proto` file from `googleapis/master` manually.
2.  **Missing HTTP Annotations**: If the proto lacks `google.api.http` options, the gRPC Gateway (REST) code won't generate. You must patch the `.proto` file to add them.
    *   *Prompt*: "The proto keys `google.api.http` annotations. Please patch it to add standard CRUD mappings."
3.  **Stale Generated Files**: If `go build` fails with conflicting types, try deleting the `generated` directory or running `make clean` (if available) before regenerating.