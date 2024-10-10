# 5. Release

## 5.1 Make Direct controller the default (for migration from TF/DCL)

* Remove the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` go tag from the CRD struct, and run `dev/tasks/generate-crds` to use Direct as the permanent controller.

* Remove all of the duplicate `-direct` mockgcp test cases, and re-record mockgcp interactions (now using direct controller instead of legacy TF/DCL).

## 5.2 Bump from v1alpha1 to v1beta1

* Copy-paste `./apis/service>/v1alpha1 `to` ./apis/service>/v1beta1 `using `git mv` (this is for PR review) and make sure the `./apis/service>/v1alpha1 `is still there.

* Update the new folder apiVersion to `v1beta1`

* Update all imports and update the MockGCP test suite to use `v1beta1`

* Add `kubebuilder:storageversion` tag to `v1beta1 `API. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/1b19153411653329177f4ba0991c982f36970707/apis/cloudbuild/v1beta1/workerpool_types.go#L155)

## 5.3 Turn on doc auto-generation (Direct resource only)

* Add the direct resource to` config/servicemappings`, only specify the `name`, `kind` and <code>direct(true) </code>[example](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2182/files#diff-e463d47dab0190c35c12d64604451db84e0e7b6316ce33524a2a4eb29e0f2e47)