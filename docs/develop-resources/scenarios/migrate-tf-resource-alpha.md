# Migrate TF/DCL-based Alpha to Direct Alpha

## Add MockGCP tests
 
Follow [Step 1](https://github.com/yuwenma/k8s-config-connector/blob/scifi-guide/docs/develop-resources/guides/1-add-mockgcp-tests.md)

1. The 1st PR should set the `create.yaml `and  `update.yaml `fields the same value for both test suites, with `_http.log `telling the matching HTTP request/response, and `_generated_object_<resource>.golden.yaml` telling the output-only fields. It shall record against real GCP 

2. The 2nd PR should have `update.yaml `fields update all mutable fields for both basic and full test suites, with  `_http.log` and` _generated_object_<resource>.golden.yaml `telling the corresponding git diff. 

### PR Reviews

* We require the PR to contain the real GCP record for `_generated_object_<resource>.golden.yaml` and `_http.log` 
* We require the 2nd PR git diff can show the mutable fileds in `update.yaml`.
* We require the `_generated_object_<resource>.golden.yaml` reflecting the mutable fields are successfully updated.
* We require the `dependencies.yaml` to cover all referenced fields, and the `_http.log` showing the Cloud requests. You need to implement those dependencies' MockGCP methods as well.

## Add API

Follow [Step 2](https://github.com/yuwenma/k8s-config-connector/blob/scifi-guide/docs/develop-resources/guides/2-define-apis.md)

The PR shall contain the types and deepcopy codes. It shall follow the Direct resource recommended styles and conventions. (TODO add the link) **It can change the existing fields since this is a Alpha resource**.

* We always start from Alpha resources. So the migration shall first be placed under `./apis/service>/v1alpha1`. Once the alpha resource is released, we then bump it to Beta.
* Add `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` to the API tag [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92), to continue using TF-based controllers. 

### PR Reviews

* If the config/crds is not changed after running dev/tasks/generate-crds, we require the _generated_object_<resource>.golden.yaml unchanged
* If the API is updated, we require the create.yaml and update.yaml changes. Also, the config/crds and _generated_object_<resource>.golden.yaml shall match the corresponding changes.

## Add the mapper

Follow [Step 3](https://github.com/yuwenma/k8s-config-connector/blob/scifi-guide/docs/develop-resources/guides/3-add-mapper.md)

This PR adds the Direct mapper. You can do this together with the previous step or the next step if no additional manual changes are needed.

## Add the controller 

Follow [Step 4](https://github.com/yuwenma/k8s-config-connector/blob/scifi-guide/docs/develop-resources/guides/4-add-controller.md).

* Use the `KCC_USE_DIRECT_RECONCILERS` flag [exampe](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/dev/tasks/run-e2e#L27). This will override the tf2crd and dcl2crd label to force using the Direct controller. 


### PR Reviews

* We require the roundtrip fuzz tests to cover all the fields in `spec` and `status.observedState` fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92) (For mapper)
* We require removing the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` tags from the API and the presubmit test passes.
