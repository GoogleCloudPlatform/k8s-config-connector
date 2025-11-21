# Add a New Field to a Direct Resource (AI-Assisted Workflow)

This guide explains how to add a new field to an existing direct resource using the Gemini CLI.

> **Note:** If the resource is still using the legacy Terraform or DCL controller, you must first migrate it to the direct approach. Please follow the [migration guide](./migrate-tf-resource-beta.md) before proceeding.

## 1. Instruct Gemini

Use the following prompt to instruct Gemini to add the new field. Be sure to replace the placeholder values with the specifics of your field and resource.

**Prompt:**
```
Add the `<fieldName>` field from the `<google.service.version.Message>` proto to the `<Kind>` resource in the `<group>/<version>` API.

Example:
Add the `softDeletePolicy` field from the `google.storage.v1.Bucket` proto to the `StorageBucket` resource in the `storage.cnrm.cloud.google.com/v1beta1` API.
```

---

## 2. Understand the Process

Gemini will handle the necessary code generation and test updates to integrate the new field.

### What Gemini Will Do for You

1.  **Update API Types:** Gemini will run the `update-types` tool to add the new field and any of its nested message types to the resource's API definition in the `/apis` directory.
2.  **Update Mapper:** It will regenerate the KRM-to-Proto mapper to include functions for the new field.
3.  **Update Test Suites:** It will add the new field to the `create.yaml` and `update.yaml` files in the relevant MockGCP test suites to ensure it is covered.
4.  **Update MockGCP Server:** If necessary, it will update the MockGCP server logic to handle the new field in its simulated API responses.
5.  **Re-record Golden Logs:** Gemini will re-run the `record-gcp` and `compare-mock` steps to update the `_http.log` and `_generated_object.golden.yaml` files with the new field's behavior.
6.  **Resolve Resource References:** If the new field is a reference to another resource, Gemini will also add the necessary logic to the controller to resolve this reference, as detailed in [Deep Dive: Add Controller](../deep-dives/4-add-controller.md#42-resolve-resource-references).

---

## 3. Verify Gemini's Work

Your role is to review the changes to ensure the new field is implemented correctly.

### PR Review Checklist

-   **Verify Golden Logs:** This is the most important step.
    -   Inspect the `_http.log` diff. Does it show the new field being sent in the `CREATE` (POST/PUT) and `UPDATE` (PATCH) requests?
    -   Inspect the `_generated_object.golden.yaml` diff. Does it show the new field populated correctly in both the `spec` and `status.observedState`?
-   **Review Test Files:** Check the changes to `create.yaml` and `update.yaml`. Do they correctly specify the new field to provide adequate test coverage?
-   **API and Mapper Sanity Check:** Briefly review the generated changes in the `apis/` and `pkg/controller/direct/` directories to ensure they are sensible.
