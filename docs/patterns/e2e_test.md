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

### Optional Files

1.  **`dependencies.yaml`**: This file contains any resources that must be created *before* the primary resource. You can define multiple resources here, separated by `---`.

    ```yaml
    apiVersion: storage.cnrm.cloud.google.com/v1beta1
    kind: StorageBucket
    metadata:
      name: storagebucket-${uniqueId}
    ...
    ```

2.  **`update.yaml`**: If you want to test updating the resource, create this file with the modified configuration of the primary resource. This is highly recommended for test coverage. Note that some fields or resources might be immutable.

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

### Running a Single Fixture with `go test` and Envtest

You can also run a single fixture test directly using `go test`. First, configure the `KUBEBUILDER_ASSETS` environment variable for envtest:

```bash
source dev/tasks/setup-envtest
```

Then run the fixture against Mock GCP:

```bash
E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock go test -test.count=1 -timeout 3600s -v ./tests/e2e -run 'TestAllInSeries/fixtures/<testname>$'
```

Or against Real GCP (recording golden outputs):

```bash
E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=real GOLDEN_OBJECT_CHECKS=1 GOLDEN_REQUEST_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run 'TestAllInSeries/fixtures/<testname>$'
```

Alternatively, you can run a single fixture using `dev/tasks/run-e2e` by specifying `RUN_TESTS`:

```bash
RUN_TESTS='TestAllInSeries/fixtures/<testname>$' dev/tasks/run-e2e
```

### Re-reconciliation checks

After a fixture's create and update steps pass, the harness "touches" the primary object and checks what the re-reconciliation did.

*   **Write check (default).** `ShouldTestRereconiliation` in `tests/e2e/ratcheting.go` decides whether the check runs for a resource. A re-reconciliation that makes any write call to GCP fails the test. Read calls are allowed.
*   **No-diff check (opt-in).** `ShouldAssertNoDiffOnRereconcile` in `tests/e2e/ratcheting.go` matches test names by prefix. It currently matches only `computeinstance-293*`. It also fails a fixture when the Terraform controller still finds changes to apply after the touch, even if applying them makes no write calls (a perpetual diff).
*   **How to opt in.** Add the test name's prefix to `ShouldAssertNoDiffOnRereconcile`. The fixture must also pass `ShouldTestRereconiliation`, because the no-diff check only runs inside that branch of `tests/e2e/unified_test.go`.
*   **Terraform only.** The no-diff check is skipped for `ForceDirectController` runs, which keep a fixed 2-second pause instead.
*   **Waiting.** The check waits up to 30 seconds for the Terraform reconcile that starts after the touch to end (`diffRecorder.WaitForReconcileAfter`). If that times out, the test fails with `timed out after 30s waiting for a Terraform reconcile ... to end after the touch`.
*   **The reconcile must run through to diff evaluation.** If that reconcile ends with an error, is requeued, or is skipped (for example, with `cnrm.cloud.google.com/reconcile-interval-in-seconds: "0"`), the test fails with `did not run through to diff evaluation`. Without this rule, the check could pass without checking anything.
*   **Diffs fail the test.** Any diff recorded for the object fails the test with `re-reconciliation found changes to apply (re-applied on every reconcile)`.

For a worked example, see [ComputeInstance issue 293](../resource-behaviors/computeinstance/issue-293-networkip-and-image-family.md).

### Terminal errors

`isTerminalError` in `config/tests/samples/create/samples.go` treats an `UpdateFailed` condition whose message contains `cannot make changes to immutable field(s)` as a terminal error. This applies to all fixtures and samples. The wait for Ready then fails with `terminal error (Immutable field change)` as soon as the condition is observed. Without this rule, the fixture would keep waiting until the subtest timeout, which is 5 minutes for the mock target (`subtestTimeout` in `tests/e2e/unified_test.go`). An immutable-field error does not resolve on its own, so waiting only delays the failure.

## Tips

*   Use `${uniqueId}` in names to avoid collisions.
*   Use `${projectId}` if you need the GCP project ID.
*   Use `${TEST_ORG_ID}` if you need a test organization ID (mock environment usually provides this).
