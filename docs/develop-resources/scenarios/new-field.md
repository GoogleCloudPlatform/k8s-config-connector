# Add a new field

_<span style="text-decoration:underline;">Note: If this is a Terraform-based or DCL-based resource, please migrate them to the direct approach([Alpha](./migrate-tf-resource-alpha.md), [Beta](./migrate-tf-resource-beta.md)) first before running the following steps.</span> _

## 1. Generate the field

Run the following command. This should add the new field and all of its dependency messages (if any) to the existing API files of the resource.

```bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT/dev/tools/controllerbuilder

go run . update-types insert \
    --parent "google.monitoring.dashboard.v1.Dashboard" \
    --field "row_layout" \
    --api-dir ${REPO_ROOT}/apis/monitoring/v1beta1
```

* `--parent`

Fully qualified name of the proto message holding the new field.

* `--field`

Name of the proto field to be inserted under the parent message.

* `--api-dir`

The `/apis` directory containing the existing API types for the resource. If a resource has multiple API versions, the `--api-dir` flag should specify the desired version. For example: `${REPO_ROOT}/apis/monitoring/v1beta1`.


    1. Run 3.1 Generate the API and proto mapper to update the mapper files.
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
