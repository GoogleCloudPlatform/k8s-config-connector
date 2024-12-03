# Migrate Terraform/DCL-based Alpha to direct Alpha

## Add MockGCP tests
 
The MockGCP provides the test coverage and helps you develop the direct resource in the follow up steps. *The reason we needs the Terraform/DCL-based Alpha resource here is that it is a pre-exist cheap option to help you develop the Mock GCP faster, you don't need or relies on the Terraform/DCL based Alpha resource in the follow up steps.* 

Follow [deep-dives Step 1](../deep-dives/1-add-mockgcp-tests.md)

### PR Reviews

* The 1st git-commit shall contain `create.yaml` and `update.yaml` files with exactly the same config. 
* The 2nd git-commit shall have `update.yaml` changing all the mutable fields, reflected in the git-diff based on the first commit.
* The 3rd git-commit (or N-th git-commit) shall run against real GCP. The generated `_generated_object_<resource>.golden.yaml` and `_http.log` is the golden log reflecting the real GCP. The `_generated_object_<resource>.golden.yaml` matches the `spec` in `update.yaml`
* The 4rd git-commit (or N+1-th git-commit) shall run against the Mock GCP. The git-diff shows `_generated_object_<resource>.golden.yaml` is unchanged, and the `_http.log` is as much like the real GCP as possible. 

* If the resource has dependencies (`dependencies.yaml`), we suggest cover all the referenced fields. If the dependency resource does not have mockGCP, you need to implement those dependencies' MockGCP methods as well.

## Add API

Follow [deep-dives Step 2](../deep-dives/2-define-apis.md)

The PR shall contain the types and deepcopy codes. It shall follow the direct resource [recommended styles and conventions](../api-conventions). **It can change the existing fields since it is Alpha**.

* Make sure the generated code is placed under `./apis/<service>/v1alpha1`. If not, please check your flags.
* Add `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` to the API tag [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92), to continue using Terraform-based controllers. *This relies on the existing Terraform-based controller to validate your API.*

### PR Reviews

* If the config/crds is not changed after running `dev/tasks/generate-crds`, we require `the _generated_object_<resource>.golden.yaml` unchanged.
* If the API is updated, we require the `create.yaml` and `update.yaml` changes. Also, the `config/crds` and `_generated_object_<resource>.golden.yaml` shall match the corresponding changes after running.

## Add the mapper

Follow [deep-dives Step 3](../deep-dives/3-add-mapper.md)

This PR adds the direct mapper. You can do this together with the previous step or the next step if no additional manual changes are needed.

## Add the controller 

Follow [deep-dives Step 4](../deep-dives/4-add-controller.md).

* Remove the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` from the API tag [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92), and rerun `dev/tasks/generate-crds`.

Now, everything is on the direct reconciliation and the resource no longer depends on the Terraform or DCL based controllers. To verify the direct controller, the `the _generated_object_<resource>.golden.yaml` should be unchanged.

### PR Reviews

* We suggest to have the roundtrip fuzz tests to cover all the fields in `spec` and `status.observedState` fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92) (For mapper)
* We require removing the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` tags from the API and the presubmit test passes without `the _generated_object_<resource>.golden.yaml` change.
