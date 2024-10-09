# 2. Define API

Config Connector builds the API using the Google Cloud Client proto.   

## 2.1 Build the Google Cloud Proto (one-off step)

This should be a one-off command only when you need to generate or update the proto.

Make sure the` generate-pb` rule in [proto-to-mapper/Makefile](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/dev/tools/proto-to-mapper/Makefile#L2) contains your proto. If not, add one using the file path in [https://github.com/googleapis/googleapis/tree/master](https://github.com/googleapis/googleapis/tree/master). [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/6ce31faf38dfaf6f44dd964802f43f9228d5a869/dev/tools/proto-to-mapper/Makefile#L16)

Run the following command to generate the Google proto.


```
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT/dev/tools/proto-to-mapper
make generate-pb
```

## 2.2 Generate the Config Connector Types (repeat-run safe) 

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

You can run this command repeatedly. This will create or update the `zz_generated.deepcopy.go` and `config/crds` 


```
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT

./dev/tasks/generate-crds
```

## 2.4 Edit the Spec and Status

**There are 3 scenarios**

1. Add a new SciFi resource (start from Alpha)
2. Migrate a TF-based or DCL-based Alpha resource 
3. Migrate a TF-based or DCL-based Beta resource

For scenarios #1 and #2, we accept breaking changes. So it is fine if the CRD outcome in 2.3 is different from the already released CRDs.

For scenarios #3, Config Connector has to be backward compatible. So we have to keep the CRD **existing** spec and status the same.

* The API field name must **not** change. You need to manually modify the field if it’s changed. 
* The field go comment is the CRD’s field description, this can be modified.
* Beta resources must **not** use  `excluded_resources` ([link](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/eca4722eac14047ed5e0879cdd89f313bdbc9d44/dev/tasks/generate-crds#L58)) flag to bypass the presubmit check about the CRD changes. But rather, it should rely on the outcome of the` ./dev/tasks/generate-crds `as the self-development guide to make sure the generated Direct API gives the same CRD output as the existing beta CRD. See [migrate from TF-beta to Direct resource](https://docs.google.com/document/d/1Az2yOr9dGHrj-IGvYEhGxjpajGmcd65mQ4OjsAQSXyE/edit?resourcekey=0-Mf91G-_QDqq6XEVzf_CnQQ&tab=t.0#heading=h.8mqswg27ngyc) for more detailed guide.

According to the above principles, you shall decide how to process the following. Please check out the [Migrate a TF/DCL Beta Resource to SciFi Beta](https://docs.google.com/document/d/1Az2yOr9dGHrj-IGvYEhGxjpajGmcd65mQ4OjsAQSXyE/edit?resourcekey=0-Mf91G-_QDqq6XEVzf_CnQQ&tab=t.0#heading=h.8mqswg27ngyc) and [Bump a TF/DCL Alpha Resource to SciFi Beta](https://docs.google.com/document/d/1Az2yOr9dGHrj-IGvYEhGxjpajGmcd65mQ4OjsAQSXyE/edit?resourcekey=0-Mf91G-_QDqq6XEVzf_CnQQ&tab=t.0#heading=h.t9v0v643cc5f)** **to see the PR review requirements and detailed steps.


* Add the parent field and mark as required

    * See detailed [requirements and example](../api-conventions/validations.md#rule-3-parent).

* Replace any field that is a resource reference to `<resource>Ref `and add the resource to `./apis/refs` if not exist.` `

    * See [direct resource reference validation rules](../api-conventions/resource-reference.md). This validation shall consider the actual resource field’ usage and can be done in follow-up PRs 

* Add **non** output-only fields to `spec`, excluding imperative fields.

* Add other CR validations according to the [Direct Resource API Validation Guide](../api-conventions/validations.md)

* (Only for TF/DCL Beta) Existing `spec` and `status` fields should still be there, except the [output-only spec](https://paste.googleplex.com/4694303066030080) should be removed. 


* (Only for TF/DCL Beta) Add the [output-only spec](https://paste.googleplex.com/4694303066030080) fields to `status.observedState`.

### Exit Criteria

* The API PRs shall pass the MockGCP tests. This requires a Dirct controller or a TF/DCL based controller.   
* All fields shall be covered and properly handled (no `/*NOTYET*/` comments). 
