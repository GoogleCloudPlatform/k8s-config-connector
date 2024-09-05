# 5. Release

## 5.1 Turn on your Direct controller (TF/DCL Beta Only)

### For TF-based Beta resource

* Remove the `cnrm.cloud.google.com/tf2crd: "true"` label from the CRD will turn on SciFi controller. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/196a4b9a28b59b17936a443d5b36bb65f3c42fd9/apis/apikeys/v1alpha1/apikey_type.go#L44)

### For DCL-based Beta resource

* Remove the `cnrm.cloud.google.com/dcl2crd: "true"` label from the CRD will turn on SciFi controller.

## 5.2 Bump from v1alpha1 to v1beta1

* Copy-paste `./apis/service>/v1alpha1 `to` ./apis/service>/v1beta1 `using `git mv` (this is for PR review) and make sure the `./apis/service>/v1alpha1 `is still there.

* Update the new folder apiVersion to `v1beta1`

* Update all imports and update the MockGCP test suite to use `v1beta1`

* Add `kubebuilder:storageversion` tag to `v1beta1 `API. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/1b19153411653329177f4ba0991c982f36970707/apis/cloudbuild/v1beta1/workerpool_types.go#L155)

## 5.3 Turn on doc auto-generation (Direct resource only)

* Add the direct resource to` config/servicemappings`, only specify the `name`, `kind` and <code>direct(true) </code>[example](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2182/files#diff-e463d47dab0190c35c12d64604451db84e0e7b6316ce33524a2a4eb29e0f2e47)