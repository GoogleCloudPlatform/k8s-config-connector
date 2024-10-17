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

Copy all existing test cases and add `-direct` suffix to the test names [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/61e85e1fc5f48de8c5c652cdb73aae48dd7dfecf/pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlinstance).

Enable direct controller for the the new test cases by setting the annotation `alpha.cnrm.cloud.google.com/reconciler: direct` in `create.yaml` and `update.yaml`. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/c0a77915723788e5068a819c93f22c4661e47b6b/pkg/test/resourcefixture/testdata/basic/dataflow/v1beta1/dataflowflextemplatejob/batchdataflowflextemplatejob-direct/create.yaml#L20)

This will override the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` annotations in the CRD and enable the direct controller for the new test cases. The previously-existing test cases will continue to use the TF/DCL-based controller.

Verify the new `-direct` test cases have equivalent behavior to the existing test cases (though not necessarily the exact same API interactions, if the direct controller achieves the same end result differently than the TF/DCL-based controller).

### Exit Criteria

* The PRs shall pass the MockGCP tests
* The roundtrip fuzz tests shall cover all the fields in `spec `and `status.observedState `fields [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92)
* There are equivalent Direct and TF/DCL-based mockgcp test cases, and all tests are passing.