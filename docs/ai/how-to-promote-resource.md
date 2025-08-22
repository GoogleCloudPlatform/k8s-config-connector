# How to Promote a KCC Resource from v1alpha1 to v1beta1

This guide outlines the process for promoting a Config Connector resource from an alpha API version (`v1alpha1`) to a beta API version (`v1beta1`). The promotion is a three-step process that involves promoting the API definition, the controller, and the test fixtures.

## Overview

The promotion process requires careful, sequential updates to the API definitions, controller logic, and test data. While some automation is available, this guide presents the manual process as the primary workflow to ensure a thorough understanding and to handle the common complexities that arise from cross-version dependencies.

The three main steps are:
1.  **Promote the API:** Create the new `v1beta1` API definition from the `v1alpha1` source.
2.  **Promote the Controller:** Update the resource's controller to use the new `v1beta1` API.
3.  **Promote the Test Fixtures:** Update the test data and golden files for the new version.

---

## Step 1: Promote the API

The first step is to create the new `v1beta1` API files based on the existing `v1alpha1` files.

### Action

1.  **Prepare the `v1beta1` directory:**
    *   Create the new directory (e.g., `apis/myservice/v1beta1`).
    *   Copy `generate.sh` from the `v1alpha1` directory to the `v1beta1` directory.
    *   If `generate.sh` doesn't exist, create a new one inside `v1beta1` by using `apis/documentai/generate.sh` as a template.
    *   In the new `v1beta1/generate.sh`, update the `--api-version` flag to point to `v1beta1` and ensure the correct resources are listed.

2.  **Copy and Modify Core API Files:**
    *   Copy `_types.go`, `_identity.go`, and `_reference.go` from the `v1alpha1` directory to the `v1beta1` directory.
    *   In each of these copied Go files, change the package name from `v1alpha1` to `v1beta1`.
    *   Update any import paths within these files that might be pointing to the old version.

3.  **Generate and Validate Code:**
    *   Run the `./generate.sh` script from within the `v1beta1` directory. This will create or update the `types.generated.go` file.
    *   After the `generate.sh` script succeeds, run `make generate-crds` from the root of the repository. This validates the changes and updates the CRD definition.

4.  **Finalize Promotion and Cleanup:**
    *   Once `make generate-crds` succeeds, edit the `v1beta1/..._types.go` file. Add the following label to the `+kubebuilder:metadata:labels` annotation on the main resource struct: `"internal.cloud.google.com/additional-versions=v1alpha1"`.
    *   Finally, delete the `_types.go`, `_identity.go`, and `_reference.go` files from the original `v1alpha1` directory.

### How to Handle Failures

If `make generate-crds` fails, it is almost always due to Go compilation errors caused by cross-version dependencies.

**To fix these errors:**

1.  **Analyze the compilation error:** The error message will tell you which file and which type is causing the problem.
2.  **Identify the dependency type:**
    *   **Scenario A: The new `v1beta1` resource depends on another resource that is still in `v1alpha1`.**
        *   **Fix:** Edit the new `v1beta1/..._types.go` file. Add a new, aliased import for the `v1alpha1` package and update the struct field to use the aliased type.
            ```go
            // Add this import
            import krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/some-service/v1alpha1"

            // Update the field's type
            // From: MyField *SomeOtherTypeRef `json:"myField"`
            // To:   MyField *krmv1alpha1.SomeOtherTypeRef `json:"myField"`
            ```
    *   **Scenario B: An existing `v1alpha1` resource depends on the resource you are promoting.**
        *   **Fix:** Edit the `v1alpha1/..._types.go` file of the *other* resource. Add an aliased import for the new `v1beta1` package and update its struct field to point to the new type.
3.  **Re-validate:** Run `make generate-crds` again. If it succeeds, the API promotion is complete.

### Alternative (Advanced): Using `promote_api`

For simple cases without complex dependencies, you can use the `promote_api` tool to automate the steps above.

**Example:**

```
promote_api(apiPath='apis/storage/v1alpha1/storagebucket_types.go', targetVersion='v1beta1')
```

If this tool fails, you will need to follow the manual process and the "How to Handle Failures" guide to resolve the issues.

---

## Step 2: Promote the Controller

Next, update the controller to make it aware of the new `v1beta1` API version. To minimize code changes, the convention is to keep `krm` as the import alias for `v1alpha1` and introduce `krmv1beta1` as the alias for `v1beta1`. These changes may apply to the `*_controller.go` file, the resource mapper, and any fuzz test files.

### Action

1.  **Analyze for compilation errors:** The message will point to an undefined type in a specific file (the controller, mapper, or fuzzer).

2.  **Manually Refactor Imports:** Open the failing `.go` file.
    *   If the file only needs to refer to the newly promoted `v1beta1` types, simply change the import path to the `v1beta1` package while keeping the `krm` alias.
    *   If the file needs to reference types from both API versions, ensure there are two imports for the API, aliased correctly:
        ```go
        import (
            krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1alpha1"
            krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
        )
        ```

3.  **Update Type Usages:**
    *   Search for all usages of types from the API within the file.
    *   Prefix types that were promoted to the new version with `krmv1beta1.` (e.g., `krmv1beta1.StorageBucketSpec`).
    *   Prefix types that remain in the old version with `krm.` (e.g., `krm.SomeOtherSpec`).

4.  **Re-validate:** Run `go build ./...` from the controller directory. Repeat the process for any other files with compilation errors until all are resolved.

### Alternative (Advanced): Using `promote_controller`

For simple cases, you can use the `promote_controller` tool to attempt to automatically refactor the import aliases and type usages. If it fails, follow the manual process above to resolve any compilation errors.

**Example:**

```
promote_controller(controllerPath='pkg/controller/direct/storage/storagebucket_controller.go', apiPath='apis/storage/v1alpha1/storagebucket_types.go', targetVersion='v1beta1')
```

---

## Step 3: Promote the Test Fixtures

Finally, move the test fixtures from `v1alpha1` to `v1beta1` and update any other tests that depend on the promoted resource.

### Action

1.  **Identify and Move Test Directories:**
    *   Navigate to `pkg/test/resourcefixture/testdata/basic/<service>/v1alpha1/`.
    *   Identify all subdirectories containing tests for the resource being promoted (by checking the `kind` in the `create.yaml` file).
    *   Move these directories to the corresponding `v1beta1` path (e.g., `.../<service>/v1beta1/`).

2.  **Update Moved Test Fixtures:**
    *   In each of the newly moved test directories, edit the `create.yaml` and `update.yaml` files.
    *   Change the `apiVersion` field in these two files to use the `v1beta1` version (e.g., `myservice.cnrm.cloud.google.com/v1beta1`).
    *   **Important:** In this step, do *not* change the `apiVersion` for any resources listed in the `dependencies.yaml` file within these moved directories.

3.  **Update Dependent Fixtures Across the Project:**
    *   Search across the entire `pkg/test/resourcefixture/testdata/basic/` directory for any `dependencies.yaml` files that use the resource you just promoted.
    *   For each `dependencies.yaml` file that you find, update the entry for the promoted resource to use its new `v1beta1` `apiVersion`.

4.  **Validate All Changes:**
    *   Run the test validation script (e.g., `hack/compare-mock`) to ensure the mock GCP layer behaves as expected. You may need to run it for the promoted resource's fixtures as well as any fixtures whose dependencies you updated.

### Alternative (Advanced): Using `promote_tests`

For simple cases, the `promote_tests` tool can automate copying the test directories and updating the `apiVersion` in the YAML files. You may still need to manually check for dependencies in other test fixtures as described in step 3.

**Example:**

```
promote_tests(testFixturePath='pkg/test/resourcefixture/testdata/basic/storage/v1alpha1/storagebucket', kind='StorageBucket', targetVersion='v1beta1')
```
