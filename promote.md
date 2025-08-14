# KCC Resource Promotion Guide

This document outlines the process for promoting a KCC resource to a new API version (e.g., from `v1alpha1` to `v1beta1`). Promotion is a critical process that involves moving a resource to a more stable, "direct" controller implementation while ensuring backward compatibility for existing users.

## The Promotion Process

Promotion is a three-step process that should be performed in the following order. The `mcp-kcc` integration with the Gemini CLI provides tools to automate each step.

### 1. Promote the API

The first step is to create the new API version.

**Tool:** `promote_api`

**What it does:**
- Copies the source API package (e.g., `apis/<service>/v1alpha1`) to the new target version (e.g., `apis/<service>/v1beta1`).
- Updates all version strings in the new Go files from the source version to the target version.
- Adds crucial annotations to the primary type's `_types.go` file:
    - `// +kubebuilder:storageversion`: Marks the new version as the storage version for the CRD.
    - `// +kubebuilder:metadata:labels: "internal.cloud.google.com/additional-versions={source_version}"`: Ensures the CRD definition preserves the previous version, which is critical for backward compatibility.
- Runs `dev/tasks/generate-crds` to validate the changes and update the CRD files in the `crds/` directory.

**Example:**
To promote the `storage` API to `v1beta1`, you can instruct Gemini:
> promote the api in `apis/storage` to `v1beta1`

### 2. Promote the Controller

Next, promote the controller to use the new API version.

**Tool:** `promote_controller`

**What it does:**
- Updates the controller's Go files to import and use the new target API version. In complex cases where a controller directory contains files that reference types from both the old and new API versions, this tool intelligently refactors the code. It ensures that the new API version is aliased to `krm`, while the old version is given a version-specific alias (e.g., `krmv1alpha1`).

**Manual Intervention and Validation:**
While the `promote_controller` tool is powerful, complex scenarios may still result in compilation errors. If the tool fails, you can follow this manual process:

1.  **Use the `promote_controller_prompt`:** This tool provides detailed, context-specific instructions for fixing the compilation errors. It will guide you on how to correctly alias the imports and update type usages.
    
    **Example:**
    > give me a prompt to promote the controller in `pkg/controller/storage/storagebucket` for the api in `apis/storage` to `v1beta1`

2.  **Fix Imports Manually:** The core idea is to have two imports for the service's API. The new version should be aliased to `krm`, and the old version should get a version-specific alias.

    ```go
    import (
        krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
        krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1alpha1"
    )
    ```
    You must then update all type usages in the file to use the correct alias (`krm.` for promoted types, `krmv1alpha1.` for types remaining in the old version).

3.  **Validate Repeatedly:** After making manual changes, use the `promote_controller_validate` tool to check your work. You can run this tool repeatedly until all compilation errors are resolved.

    **Example:**
    > validate the controller promotion for `pkg/controller/storage/storagebucket`

**Example:**
To promote the `storagebucket` controller, you can instruct Gemini:
> promote the controller in `pkg/controller/storage/storagebucket` for the api in `apis/storage` to `v1beta1`

### 3. Promote the Tests

Finally, update the test fixtures to use the new API version.

**Tool:** `promote_tests`

**What it does:**
- Scans the test fixtures in `pkg/test/resourcefixture/testdata/basic/` for the specified Kind.
- Copies the relevant test fixtures to a new directory named after the target version.
- Updates the `apiVersion` in the test YAML files (`create.yaml`, `update.yaml`, etc.) to the new target version.

**Example:**
To promote the `StorageBucket` tests, you can instruct Gemini:
> promote the tests in `pkg/test/resourcefixture/testdata/basic/storage/v1alpha1/storagebucket` for kind `StorageBucket` to `v1beta1`

## Post-Promotion Verification

After promoting the API, controller, and tests, it is **critical** to verify that the new direct controller behaves identically to the old one.

1.  **Run the full test suite:** `make test`
2.  **Run the golden tests:**
    - Record new golden files against a real GCP project: `hack/record-gcp <test-name>`
    - Compare the new behavior against the mock GCP layer: `hack/compare-mock <test-name>`

This verification ensures that we do not introduce breaking changes for users.

## Key Takeaways

- **Promotion is a three-step process:** API, Controller, Tests. Always perform them in order.
- **Use Gemini and the integrated MCP tools** to automate the promotion steps.
- **Backward compatibility is paramount.** The `additional-versions` label is essential.
- **Verification is not optional.** Always run the golden tests to prevent regressions.
- **Manual intervention may be required,** especially for fixing compilation errors in the controller.
