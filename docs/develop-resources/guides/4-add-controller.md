# 4. Add the direct controller

Run the following command to generate a controller template 

```
cd dev/tools/controllerbuilder
go run main.go add --service <YOUR_SERVICE> --api-version <VERSION> --kind <YOUR_RESOURCE> --proto-resource <PROTO_RESOURCE>
```

Fix the generated code to make your SciFi running!

## 4.1 Implement the `model` interface

The controller template has implemented the model interface` find, create, update, delete `and` export`. You may need to update the code to fit your resource.


## 4.2 Resolve resource references

Most Config Connector resource need references like `spec.projectRef. `You should add those references in `AdapterForObject` using functions `Resolve<RefResource>`

if there is no previous reference method, You may need to add a new` Resolve<RefResource> `

Check  to make sure your validation is complete.


## 4.3 Register your controller

To wire your controller in the Config Connector operator, you need to register the controller [here](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/direct/register/register.go)


## 4.4 Verify your controller

To turn on the SciFi controller to reconcile resources:


### New Resource

```
hack/compare-mock fixtures/<your_resource_test>
```

### Existing DCL/TF based resource

```
KCC_USE_DIRECT_RECONCILERS=<YOUR KIND> hack/compare-mock fixtures/<your_resource_test>
```

### Turn on SciFi Controller

To fully turn on the SciFi controller, add your resource to [KCC_USE_DIRECT_RECONCILERS](https://github.com/xiaoweim/k8s-config-connector/blob/master/dev/tasks/run-e2e#L27) and make sure it passes all tests and there is a decent test coverage over customer use cases. Usually we do not turn it on for beta resources right away.
 

 * Note: some resources are used as dependency resources in other test scenarios, to only test the direct controller with your resource test, you can use this annotation `alpha.cnrm.cloud.google.com/reconciler: "direct"` in the test yamls instead of using the `KCC_USE_DIRECT_RECONCILERS` flag at CLI.

 * Note: Differences in http.logs such as `user-agent` and method url are expected. Please regenerate test logs before `compare-mock`.

### Exit Criteria

* The PRs shall pass the MockGCP tests
* The roundtrip fuzz tests shall cover all the fields in `spec `and `status.observedState `fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92)