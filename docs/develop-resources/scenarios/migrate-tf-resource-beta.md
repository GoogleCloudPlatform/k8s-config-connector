# Migrate TF/DCL-based Beta to Direct Beta

## Add MockGCP tests
 
Follow [Step 1](../guides/1-add-mockgcp-tests.md)

1. The 1st PR should set the `create.yaml `and  `update.yaml `fields the same value for both test suites, with `_http.log `telling the matching HTTP request/response, and `_generated_object_<resource>.golden.yaml` telling the output-only fields.

2. The 2nd PR should have `update.yaml `fields update all mutable fields for both basic and full test suites, with  `_http.log` and` _generated_object_<resource>.golden.yaml `telling the corresponding git diff. 

### PR Reviews

* We require the 1st PR to show git diff between the real GCP record and the mock GCP record for `_generated_object_<resource>.golden.yaml` and `_http.log` 
* We require the 2nd PR git diff can show the mutable fileds in `update.yaml`.
* We require the `_generated_object_<resource>.golden.yaml` reflecting the mutable fields are successfully updated.
* We require the `dependencies.yaml` to cover all referenced fields, and the `_http.log` showing the Cloud requests. You need to implement those dependencies' MockGCP methods as well.

## Add the backward compatible API

Follow [Step 2](../guides/2-define-apis.md)

The PR shall contain the types and deepcopy codes. It shall make modifications to make sure the CRDs are the same, because Beta resource has to be backward compatible. 

* You may need to modify the auto-generated `types.go` to keep the existing fields the same (even if it is not following the recommended styles and conventions). You can run `dev/tasks/generate-crds` (repeatedly) to make sure the CRD are the same (comment changes are acceptable).

* Add `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` to the API tag [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92), to continue using DCL-based or TF-based controllers. 

* You may see some new fields added to the CRD. These are expected since the TF/DCL based resources could be out of date (and users are looking forward to these new fields!). You **shall comment out** those new fields using `/*NOTYET ..*/` in this PR if they are not supported in the TF-based controller yet (we will add them later).

Note: *Do not use* `excluded_resources`, we want the presubmit to validate the existing APIs are unchanged and the behaviors are the same.

### PR Reviews

* We require the `_http.log`, `_generated_object_<resource>.golden.yaml`, `create.yaml` and `update.yaml` unchanged. 

## Add the backward compatible mapper

Follow [Step 3](../guides/3-add-mapper.md)

The PR adds the Direct mapper. You can do this together with the previous step or the next step if no additional manual changes are needed. Using  `/*NOTYET .. */` to comment out new functions, same as the last step.

## Add the direct controller 

Follow [Step 4](../guides/4-add-controller.md)

* Use the `KCC_USE_DIRECT_RECONCILERS` flag [exampe](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/dev/tasks/run-e2e#L27). 

*Tips* The `KCC_USE_DIRECT_RECONCILERS` will override the `tf2crd` and `dcl2crd` label to force using the Direct controller, but it will not affect the releases which will still use the TF/DCL based controllers until the Direct controller is ready. This allows developing the API and controller separately. 

### PR Reviews

* We require the roundtrip fuzz tests to cover all the fields in `spec` and `status.observedState` fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92) (For mapper)
* We require the MockGCP pass without any change to the `create.yaml`, `update.yaml` and `_generated_object_<resource>.golden.yaml` files. 

## Add new API and mapper

The PR shall add back those fields that commented out in Step 3. 
You need to update both `types.go` and `mapping.go`, and may need to adjust the direct controller as well to support those new fields.
 
### PR Reviews

* The MockGCP test shall reflect the new fields in `create.yaml`, `update.yaml` (if mutable),  `_generated_object_<resource>.golden.yaml`, and `_http.log`.
* We require the roundtrip fuzz tests to cover the new fields in `spec` or/and `status.observedState` fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92) (For mapper)
* We require no /*NOTYET .. */ in the API and Mapper code.

## Release

* Remove the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` go tag, and run `dev/tasks/generate-crds` to use Direct as the permanent controller (step 5)
* Remove your resource from the `KCC_USE_DIRECT_RECONCILERS` flag since it is no longer needed.

### PR Reviews

* We require the MockGCP unchange and pass. 