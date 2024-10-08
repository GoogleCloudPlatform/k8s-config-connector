# Add a new Config Connector resource

## A basic infra 

Different than other scenarios, developing a pure Direct resource requires you to have every step setup first to validate your change. Basically you need the API and mapper to write the Direct controller, then you can use the controller to write the MockGCP tests. This reverses the step 1 to the end compared to other scenarios. Thus, you first PR shall include

1. A basic API from the auto-generated code, including all required fields (Step 2)
2. A basic Direct controller (Step 3 and 4).
3. Define `create.yaml` and `update.yaml` to run against the real GCP (Step 1 [record real gcp](../guides/1-add-mockgcp-tests.md#record-real-gcp-log))

### PR reviews

* Code shall come from the auto-generate code.
* Only the basic test suite is required, full and other test suites can comment out unimplemented fields.
* `_http.log` and  `_generated_object_<resource>.golden.yaml` shall verify the basic controller is functioning.
* We require the PR to contain the real GCP record for `_generated_object_<resource>.golden.yaml` and `_http.log` 

## Add MockGCP 

Add the mockGCP (Step 1) 

### PR reviews

* The `_http.log` shall have minimum changes.
* The  `_generated_object_<resource>.golden.yaml`, `create.yaml` and `update.yaml` shall not change

## Add full api and mapping

* Add all API fields.
* Add full test suites.

### PR reviews 

* No `/*NOTYET .. */` in the api and mapping.
* We require the MockGCP test suits to cover all fields. 
* We require the roundtrip fuzz tests to cover all the fields in `spec` and `status.observedState` fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92)
* The version has to start from v1alpha1. 