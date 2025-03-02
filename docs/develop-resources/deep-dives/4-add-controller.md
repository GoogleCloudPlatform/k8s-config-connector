# 4. Add the direct controller

Run the following command to generate a controller template 

```
cd dev/tools/controllerbuilder
go run main.go generate-controller --service <YOUR_SERVICE> --api-version <VERSION> --resource <YOUR_RESOURCE>:<PROTO_RESOURCE>
```

Fix the generated code to make your SciFi running!

## 4.1 Implement the `model` interface

The controller template has implemented the model interface` find, create, update, delete `and` export`. You may need to update the code to fit your resource.


## 4.2 Resolve resource references

Most Config Connector resource need references like `spec.projectRef. `You should add those references in `AdapterForObject` using functions `Resolve<RefResource>`

if there is no previous reference method, You may need to add a new` Resolve<RefResource> `

Check  to make sure your validation is complete.

## 4.3 Verify your controller

To turn on the SciFi controller to reconcile resources:


### New Resource

```
hack/compare-mock fixtures/<your_resource_test>
```

### Existing DCL/TF based resource

```
KCC_USE_DIRECT_RECONCILERS=<YOUR KIND> hack/compare-mock fixtures/<your_resource_test>
```

 * Note: Differences in http.logs such as `user-agent` and method url are expected. Please regenerate test logs before `compare-mock`.

### Exit Criteria

* The PRs shall pass the MockGCP tests
* For Beta resource, the roundtrip fuzz tests shall cover all the fields in `spec `and `status.observedState `fields. [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/f313b00c52f09c4a52a2eb5fe2c15fa4b30a05fd/pkg/controller/direct/discoveryengine/fuzzers.go#L26-L47)
