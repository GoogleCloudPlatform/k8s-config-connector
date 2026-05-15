# Add a new field

_<span style="text-decoration:underline;">Note: If this is a Terraform-based or DCL-based resource, please migrate them to the direct approach([Alpha](./migrate-tf-resource-alpha.md), [Beta](./migrate-tf-resource-beta.md)) first before running the following steps.</span> _

## 1. Generate the field

Manually update the API definitions in the `/apis` directory to include the new field. You may also need to run the `generate-types-and-mappers` tool depending on your workflow.

    1. Run `dev/tasks/generate-types-and-mappers` to update the mapper files and generated code.
    2. Add the fields to `create.yaml `and `update.yaml `in corresponding test suites.
    3. Modify the MockGCP when necessary. The new fields should show up in `_http.log` and   `_generated_object_<resource>.golden.yaml` 

## 2. Resolve resource reference

If the newly added fields contain resource references, you should have a second PR to update the resource reference, following [4.2 resolve resource references](../deep-dives/4-add-controller.md#42-resolve-resource-references)

Add a new test suite with `dependencies.yaml` to cover the referenced fields.


## Add MockGCP coverage
 
Follow [Step 1](../deep-dives/1-add-mockgcp-tests.md)

### PR Reviews

* We require the PR to contain the real GCP record for `_generated_object_<resource>.golden.yaml` and `_http.log` 
* We require the PR reflects the new field in the `create.yaml`, `update.yaml`(if mutable),  `_generated_object_<resource>.golden.yaml` and `_http.log`
