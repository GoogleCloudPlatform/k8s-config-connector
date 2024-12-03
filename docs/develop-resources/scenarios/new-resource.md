# Add a new Config Connector resource

Adding a new direct resource is similar to migrating a Terraform/DCL based Alpha resource to the direct approach, with some additional challenges that you don't have the cheap option to verify each step until you have all the basic infra (API, Mapper, Controller) ready.

## A basic infra

Developing a pure direct resource requires you to have every step setup to validate your change. Basically you need the API and mapper to write the direct controller, then you can use the controller to write the MockGCP tests. This reverses the [deep-dives step 1](../deep-dives/1-add-mockgcp-tests.md) to the end compared to other scenarios. Thus, you first PR shall include

1. A basic API from the auto-generated code, including all required fields [deep-dives Step 2](../deep-dives/2-define-apis.md)
1. A basic direct controller [deep-dives Step 3](../deep-dives/3-add-mapper.md) and [Step 4](../deep-dives/4-add-controller.md).
1. A basic test suite to run against the real GCP (Step 1)

### PR reviews

* Code shall come from the auto-generate code.
* Only the basic test suite is required, full and other test suites can comment out as unimplemented fields.
* `_http.log` and  `_generated_object_<resource>.golden.yaml` shall verify the basic controller is functioning correctly.
* Commit the real GCP record and mockGCP record as contiguous git-commits, so the PR reviewers can verify.

## Add full test suite

Once the basic infra is done, you can add full field test coverage to prepare promoting the resource to Beta.  

### PR reviews

* The `_http.log` shall have minimum changes.
* The  `_generated_object_<resource>.golden.yaml`, `create.yaml` and `update.yaml` shall not change

## Add full API and mapping

* Add all API fields.
* Add full test suites.

### PR reviews 

For Beta resource, we require the following conditions to meet. 

* No `/*NOTYET .. */` in the api and mapping.
* We require the MockGCP test suits to cover all fields. 
* We require the roundtrip fuzz tests to cover all the fields in `spec` and `status.observedState` fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92)
* The version has to start from v1alpha1.