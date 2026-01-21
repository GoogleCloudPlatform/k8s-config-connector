# How to Create E2E Tests

This document describes how to create end-to-end (e2e) tests for Config Connector resources.

## Test Structure

E2E tests are located in `pkg/test/resourcefixture/testdata/basic`. The path structure generally follows:
`pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testName>/`

For example: `pkg/test/resourcefixture/testdata/basic/tags/v1alpha1/tagslocationtagbinding/tagslocationtagbinding_bucket/`

### Required Files

1.  **`create.yaml`**: This file contains the YAML configuration for the **primary** resource under test. It uses `${uniqueId}` for name uniqueness.

    ```yaml
    apiVersion: tags.cnrm.cloud.google.com/v1alpha1
    kind: TagsLocationTagBinding
    metadata:
      name: tagstagbinding-${uniqueId}
    spec:
      location: us-east4
      parentRef:
        kind: StorageBucket
        external:  "//storage.googleapis.com/projects/_/buckets/storagebucket-${uniqueId}"
      tagValueRef:
        name: tagstagvalue-${uniqueId}
    ```

2.  **`dependencies.yaml`** (Optional): This file contains any resources that must be created *before* the primary resource. You can define multiple resources here, separated by `---`.

    ```yaml
    apiVersion: storage.cnrm.cloud.google.com/v1beta1
    kind: StorageBucket
    metadata:
      name: storagebucket-${uniqueId}
    ...
    ```

### Optional Files

1.  **`update.yaml`**: If you want to test updating the resource, create this file with the modified configuration of the primary resource. This is highly recommended for test coverage. Note that some fields or resources might be immutable.

### Generated Golden Files

When you run the test, several files are generated. These are "golden" files and should be committed to the repository. They ensure no regressions in behavior.

*   **`_http.log`**: Captures the HTTP traffic (requests and responses) against the real GCP API or the mock.
*   **`_generated_object*.golden.yaml`**: Represents the final state of the object in Kubernetes after reconciliation.
*   **`_final_object_old_controller.golden.yaml`**: Similar to above, but for tests running against older controller implementations (if applicable).
*   **`_generated_export*.yaml`**: If the resource supports export, this file captures the exported YAML.

## Running Tests

We use helper scripts to run tests. Typically, we run against a **mock** GCP implementation because it is faster and doesn't require real credentials/projects for every developer.

### Running against Mock GCP

Use `hack/compare-mock` to run the test against the mock framework. You can specify the path to your specific test directory to save time.

```bash
hack/compare-mock pkg/test/resourcefixture/testdata/basic/tags/v1alpha1/tagslocationtagbinding/tagslocationtagbinding_bucket/
```

This command will:
1.  Spin up a local Kubernetes control plane (envtest).
2.  Start the Config Connector controller.
3.  Start the Mock GCP server.
4.  Apply `dependencies.yaml` and `create.yaml`.
5.  Wait for resources to be ready.
6.  Apply `update.yaml` (if present).
7.  Verify the state.
8.  Delete the resources.
9.  Generate/Update the golden files (e.g., `_http.log`).

### Running against Real GCP

To run against real GCP (requires credentials and a project), use `hack/record-gcp`.

```bash
hack/record-gcp pkg/test/resourcefixture/testdata/basic/tags/v1alpha1/tagslocationtagbinding/tagslocationtagbinding_bucket/
```

This is usually done when creating the test for the first time or when the mock needs to be updated with real API behavior.

## Tips

*   Use `${uniqueId}` in names to avoid collisions.
*   Use `${projectId}` if you need the GCP project ID.
*   Use `${TEST_ORG_ID}` if you need a test organization ID (mock environment usually provides this).
