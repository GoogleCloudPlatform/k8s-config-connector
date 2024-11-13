# 2. Define API

Config Connector builds the API using the Google Cloud Client proto.   

## 2.1 Build the Google Cloud Proto (one-off step)

This step generates the Google Cloud Client proto in a single file under ./dev/tools/proto-to-mapper/build. This should be a one-off command only when you need to update the proto or the first time you develop the direct resource.

Make sure the` generate-pb` rule in [proto-to-mapper/Makefile](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/dev/tools/proto-to-mapper/Makefile#L2) contains your proto. If not, add one using the file path in [https://github.com/googleapis/googleapis/tree/master](https://github.com/googleapis/googleapis/tree/master). [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/6ce31faf38dfaf6f44dd964802f43f9228d5a869/dev/tools/proto-to-mapper/Makefile#L16)

Run the following command to generate the Google proto.


```
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT/dev/tools/proto-to-mapper
make generate-pb
```

## 2.2 Generate the Config Connector Types (repeat-run safe) 

This step generates the Config Connector API types you need from the generated proto in 2.1. The generated API is placed under `./apis` together with some other files that the [Kubernetes controller-gen](https://book.kubebuilder.io/reference/controller-gen) can leverage to build the Config Connector CRD and the Controller `runtime.Object`.

Run the following command

```
cd $REPO_ROOT/dev/tools/controllerbuilder

go run main.go generate-types \
     --service google.storage.v1  \
     --proto-source-path ../proto-to-mapper/build/googleapis.pb \
     --output-api $REPO_ROOT/apis \
     --kind StorageNotification  \ 
     --proto-resource Notification \
     --api-version "storage.cnrm.cloud.google.com/v1beta1"
```

* `--service`

The proto name of the GCP service, you can find them in [https://github.com/googleapis/googleapis.git](https://github.com/googleapis/googleapis.git). For example, the SQL service is [https://github.com/googleapis/googleapis/tree/master/google/cloud/sql/v1beta4](https://github.com/googleapis/googleapis/tree/master/google/cloud/sql/v1beta4). The `–service` should be `google.cloud.sql.v1beta4`


* `--proto-source-path`

The path to the one-off file we generated in 2.1

* `--output-api`

The apis directory to where to write the result to. Shall always be   $REPO_ROOT/apis

* `--kind`

The Config Connector resource kind, camel case. Normally it should contain the service name for example `SpannerInstance`, `SQLInstance`.

* `--proto-resource`

The proto name of the resource, you can find them in [https://github.com/googleapis/googleapis.git](https://github.com/googleapis/googleapis.git). For example, the SQLInstance is named `instance` under [https://github.com/googleapis/googleapis/tree/master/google/cloud/sql/v1beta4](https://github.com/googleapis/googleapis/tree/master/google/cloud/sql/v1beta4). The proto-source should be `instance` instead of `SQLInstance`

* `--api-version`

The Config Connector apiVersion value, shall be <code><service>.[cnrm.cloud.google.com/](http://cnrm.cloud.google.com/)<version></code>, where the generated file will be placed under<code>$REPO_ROOT/apis/<service>/<version></code> (if the dir does not exist, the command will create one).  


## 2.3 Generate CRD (repeat-run safe)

You can run this command repeatedly. This step uses `controller-gen` to create or update the `zz_generated.deepcopy.go` and `config/crds` from the generated API types.


```
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT

./dev/tasks/generate-crds
```

## 2.4 Edit the Spec and Status

The Config Connector API needs to convert the Google Cloud Client Proto API to a Kubernetes-native declarative API, where that the user configurable fields are placed in the object `spec` field, to desrcibe the user's desired status. And the Config Connector will update the object `status` field to reflect the actual status of the resource. 

This step helps you define the `spec` and `status` to be declarative friendly, following Config Connector's best practice.

**There are 3 scenarios**

1. A new direct resource should follow the best practice as described below.

1. Migrating an *existing* Terraform-based or DCL-based Alpha resource should 
follow the best practice, and change the existing APIs when necessary.   

1. Migrating an *existing* Terraform-based or DCL-based **Beta** resource should keep the existing API and its behavior for backward compatibility reasons, even if it is not following the best practice. To be more specific, 
    * The API field name must **not** change. You need to manually modify the field if the existing field name mismatches the proto name.
    * The go comment of the field can be changed. Changing the go comment will update the Google Reference Doc "Fields" description once the resource is released.

According to the above principles, you can decide how to process the following.


* Add the parent field and mark as required

    * See detailed [requirements and example](../api-conventions/validations.md#rule-3-parent).

* If a field refers to a GCP object, replace it from `<resource or any custom name>` to  `<resource>Ref`. If the `<resource>` kind does not exist under `$REPO_ROOT/apis/v1beta1/refs.go`, write the `Resolve<resource>Ref` function and put it under your own resource api `*_reference.go` file.

    * See [direct resource reference validation rules](../api-conventions/resource-reference.md). This validation shall consider the actual resource field’ usage.

* Add **non** output-only fields to `spec`, excluding imperative fields.

* Add other CR validations according to the [CustomResource field validations guide](../api-conventions/validations.md)

* (Only for Terraform/DCL Beta) Existing `spec` and `status` fields should still be there, except the [output-only spec](https://paste.googleplex.com/4694303066030080) should be removed from `spec` and add to `status.observedState`.

### Exit Criteria

* The API PRs shall pass the MockGCP tests. This requires a dirct controller or a Terraform/DCL based controller.
* For Beta resource, all fields shall be covered and properly handled (no `/*NOTYET*/` comments).
