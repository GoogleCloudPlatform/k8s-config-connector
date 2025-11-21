# Add a New Config Connector Resource (AI-Assisted Workflow)

This guide outlines the process for adding a new, direct-reconciliation resource using the Gemini CLI.

## 1. Instruct Gemini

Use the following prompt to instruct Gemini to add the new resource. Be sure to replace the placeholder values with the specifics of your resource.

**Prompt:**
```
Add a new direct resource for `<Kind>` for the service `<google.service.version>`. The resource is defined in the `<path/to/proto_file.proto>` file. Start with the `v1alpha1` version.

Example:
Add a new direct resource for `StorageBucketLock` for the service `google.storage.v1`. The resource is defined in the `google/storage/v1/storage.proto` file. Start with the `v1alpha1` version.
```

---

## 2. Understand the Process

Gemini will perform all the necessary implementation steps to create the foundational infrastructure for your new resource.

### What Gemini Will Do for You

1.  **Generate a Basic API:** Gemini will generate the initial Kubernetes API types for your resource, including all the required fields, based on the GCP service's proto definition. This is covered in [Deep Dive: Define API](../deep-dives/2-define-apis.md).
2.  **Generate the Mapper:** It will create the mapping logic that translates data between the Kubernetes resource (KRM) and the GCP API's proto message format. See [Deep Dive: Add Mapper](../deep-dives/3-add-mapper.md).
3.  **Generate a Basic Controller:** It will generate the boilerplate for a new direct controller. See [Deep Dive: Add Controller](../deep-dives/4-add-controller.md).
4.  **Create a Basic Test Suite:** Gemini will create a `create.yaml` and `update.yaml` for a "basic" test case that covers the required fields.
5.  **Implement the MockGCP Server:** It will add the necessary logic to our MockGCP framework to simulate the real GCP API for your new resource.
6.  **Record Golden Logs:** Gemini will run the test against the real GCP to record the authoritative `_http.log` and `_generated_object_<resource>.golden.yaml` files. It will then run the test against the MockGCP server and ensure the results are consistent. This process is detailed in [Deep Dive: Add MockGCP Tests](../deep-dives/1-add-mockgcp-tests.md).

After this initial PR, you can then instruct Gemini in a follow-up prompt to add the remaining optional fields to achieve full test coverage.

---

## 3. Verify Gemini's Work

Your role is to review the generated artifacts to ensure they are correct.

### PR Review Checklist

-   **Golden Logs are Key:** The most important part of your review is to inspect the generated golden log files.
    -   Does the `_http.log` show the expected sequence of HTTP requests (GET 404 -> CREATE -> GET 200)?
    -   Does the `_generated_object_<resource>.golden.yaml` file show the correct `spec` and `status` for the resource after creation?
-   **API and Controller Sanity Check:** Briefly look over the generated files in `apis/` and `pkg/controller/direct/` to ensure they look reasonable.
-   **Contiguous Commits:** Ensure that the PR contains two contiguous commits for the log recording: one for the real GCP run and one for the mock comparison. This allows for easy verification of the mock's accuracy.
-   **Initial Version:** The resource's version should start at `v1alpha1`.