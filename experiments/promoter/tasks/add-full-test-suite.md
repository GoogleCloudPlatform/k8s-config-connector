**Goal:** Add a comprehensive test suite for the KCC resource `<KIND>` in the `<SERVICE>` service.

This process involves creating `create.yaml` and `update.yaml` test fixtures and ensuring they cover all possible fields in the Custom Resource Definition (CRD).

**Step 1: Create the `create.yaml` file**

1.  **Create the test directory:**
    If it doesn't already exist, create the following directory:
    `pkg/test/resourcefixture/testdata/basic/<SERVICE>/v1beta1/<KIND_LOWERCASE>/<KIND_LOWERCASE>-full/`
    (Replace `<SERVICE>` with the service name and `<KIND_LOWERCASE>` with the lowercase version of the Kind.)

2.  **Create `create.yaml`:**
    Inside the new directory, create a file named `create.yaml`. This file will define the initial state of the resource for testing.

3.  **Populate `create.yaml`:**
    -   The `create.yaml` file is a Kubernetes Custom Resource (CR) based on the resource's Custom Resource Definition (CRD).
    -   You can find the CRD for your resource at: `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<PLURAL_KIND_LOWERCASE>.<SERVICE>.cnrm.cloud.google.com.yaml`. Note that `<PLURAL_KIND_LOWERCASE>` might be different from just adding an 's' to the lowercase kind.
    -   In the `create.yaml` file:
        -   Set `.metadata.name` to `<KIND_LOWERCASE>-${uniqueId}`.
        -   If `.spec.projectRef` exists, set it to `spec.projectRef.external: ${projectId}`.
        -   Fill in valid values for as many fields in the `.spec` as possible. Try to understand the purpose of each field to provide a meaningful value.

**Step 2: Verify Field Coverage**

To ensure your `create.yaml` covers all possible fields, run the following test:

```bash
TARGET_KIND=<KIND> go test ./tests/apichecks/... -run TestCRDFieldPresenceInTests
```

-   Replace `<KIND>` with the resource Kind (e.g., `StorageBucket`).
-   If this test fails, the output will tell you exactly which fields from the CRD are missing in your `create.yaml`. Add the missing fields to your `create.yaml` and re-run the test until it passes.

**Step 3: Record the GCP Traffic for `create.yaml`**

Once the field coverage test passes, record the live GCP API calls for your `create.yaml`:

```bash
E2E_TEST_TIMEOUT=20s hack/record-gcp fixtures/<KIND_LOWERCASE>-full 
```

-   This command will create a `_http.log` file in your test directory.
-   If the command fails, examine the error messages. You may need to adjust the values in your `create.yaml` or fix issues in the resource's controller located in `pkg/controller/direct/<SERVICE>/`. You can retry this command up to 10 times.

**Step 4: Create the `update.yaml` file**

1.  **Create `update.yaml`:**
    In the same directory, create a file named `update.yaml`.

2.  **Populate `update.yaml`:**
    -   Copy the content from `create.yaml`.
    -   Modify the values of all the **mutable** fields in the `.spec`. Choose new, valid values for these fields. Try to modify as many mutable fields as possible.

**Step 5: Record the GCP Traffic for `update.yaml`**

Now, record the live GCP API calls for your `update.yaml`:

```bash
E2E_TEST_TIMEOUT=20s hack/record-gcp fixtures/<KIND_LOWERCASE>-full
```

-   This will update the `_http.log` with the API calls for the update operation.
-   If the command fails, you might have an issue with the values in `update.yaml` or with the `Update` method in the controller (`pkg/controller/direct/<SERVICE>/*_controller.go`).

**Step 6: Verify the Mock GCP Implementation**

Finally, verify that the mock GCP implementation behaves the same as the real GCP API.

1.  **Stage the recorded traffic:**
    ```bash
    git add pkg/test/resourcefixture/testdata/basic/<SERVICE>/v1beta1/<KIND_LOWERCASE>/<KIND_LOWERCASE>-full/
    ```

2.  **Compare with the mock:**
    ```bash
    E2E_TEST_TIMEOUT=10s hack/compare-mock fixtures/<KIND_LOWERCASE>-full
    ```
    -   If this command fails, it means there's a difference between the real and mock GCP interactions.
    -   Examine the diff to see the differences:
        ```bash
        git diff pkg/test/resourcefixture/testdata/basic/<SERVICE>/v1beta1/<KIND_LOWERCASE>/<KIND_LOWERCASE>-full/_http.log
        ```
    -   Modify the mock implementation in `mockgcp/mock<SERVICE>/` to match the real behavior. For example, the `Insert`, `Update`, and `Delete` functions in `mockgcp/mocksql/sqlinstance.go` should mirror the behavior of the `a.sqlInstancesClient.Insert`, `sqlInstancesClient.Update`, and `sqlInstancesClient.Delete` calls in `pkg/controller/direct/sql/sqlinstance_controller.go`.
