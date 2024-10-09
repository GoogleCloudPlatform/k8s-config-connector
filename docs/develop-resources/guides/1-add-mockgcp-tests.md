# 1. Add MockGCP tests

## General rules
MockGCP test is required for the Direct resource development. 

1. This is required no matter if the feature request is small or big. For example, both adding a single new field or adding an entire new resource require the MockGCP test coverage. If the MockGCP does not exist yet when adding a single field, developers should add the corresponding MockGCP coverage for the entire resource.  
1. Each Config Connector `spec` field should be covered in both create and update(if applicable) cases. 
1. MockGCP check is auto-enabled in Presubmit check and shall not be skipped. Code changes to a Config Connector resource must not merge if not covered by MockGCP check. 
1. New PRs should give steady golden log output before they can merge in. Reasonable golden object and http log changes are okay and sometimes expected.
1. MockGCP should be up-to-date when migrating resource from TF/DCL based controller to the Direct Controller
1. Real GCP record should be uploaded to reflect and validate the change in scenarios like adding a new resource, migrating from TF or DCL based controller to the direct controller, adding a new field, etc. 

### Suggestion

**Choose a TF-based resource to get started**. If this is your first time to add a MockGCP test, we suggest you start from a TF-based resource rather than picking a new Direct resource (note: all direct resources already have MockGCP coverage). Because you need an existing controller to help *validate* your MockGCP is functioning. Once you are familiar with writing a MockGCP test, you can work on the advanced pure Direct resource development.

## 1.1 Write the MockGCP server

Follow [this guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/README.md) to write a mock gcp server.

## 1.2 Add the test suite

Create a directory for your resource [pkg/test/resourcefixture/testdata/basic](pkg/test/resourcefixture/testdata/basic), following the naming convention as other directories. The bottommost directory name is the test suite name. You can create as many directory as your test needs. 

Just to add the `create.yaml`,  `update.yaml` and `dependencies.yaml` (if applicable) to each bottommost directory.
These three files are all Kubernetes config files, excep that `create.yaml` and `update.yaml` only holds the target CR object, and the `dependencies.yaml` holds all the dependent objects. Your test suite is done. 

## 1.3 Verify the MockGCP server

To write a new MockGCP server or to validate an existing one for the Direct resource, you should record the real GCP log and compare it with your MockGCP log, and make sure they are as likely as each other.

#### Record real GCP log 

```
hack/record-gcp fixtures/<your_resource_test>
```

`your_resource_test` is the folder directory name where you put the create.yaml and update.yaml files in. You can also use keywords to run multiple test suites. i.e. `hack/record-gcp fixtures/sql` matches all test suites that contains `sql`


#### Record mock GCP log. 

Note: This will override the output!
(remove the` WRITE_GOLDEN_OUTPUT=1` in hack/compare-mock if you don’t want to override)

```
hack/compare-mock fixtures/<your_resource_test>
```

## Exit Criteria

You shall finish the following code to claim a MockGCP task is done.  


### MockGCP Server

* Implement the mocked methods for CREATE, UPDATE, GET, and DELETE.
* If the Cloud API client returns long running operation (LRO), use `startRLO ` [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/8a350a029803a322e2889fd693cabf9780828c47/mockgcp/mockcloudbuild/workerpool.go#L56)
* CREATE method shall assign the “output-only” fields. [Example 1](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/5e08dbffb5fa3922dd43c451f35fdec45882205a/mockgcp/mockresourcemanager/tagkeys.go#L99C23-L99C37) [Example 2](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/611abaff651af81bed4517f62f915318f1b239bd/mockgcp/mocksql/sqlinstance.go#L68-L180)
* UPDATE method shall be able to update all mutable fields and support field masks if the real server does. [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/8a350a029803a322e2889fd693cabf9780828c47/mockgcp/mockcloudbuild/workerpool.go#L100)
* Make sure the log gives a stable outcome. You can override the value in<code> normalize.go </code>[example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/ba513862c2fb6ec3e54a05f6483c76b0337d6cbd/tests/e2e/normalize.go#L100)

### A "basic" test suite

* `create.yaml` that only specifies the required fields
* `update.yaml` that changes all the mutable fields     
* `_http.log `that reflects the workflow of GET resource (return 404) → CREATE resource (return 200) → GET resource (return 200)  → UPDATE resource (return 200) →  GET resource (return 200) → DELETE resource (return 204/200)[^1]` `


* `_generated_object_<resource>.golden.yaml `that contains `.status` that reflects the updated fields
* `_generated_export_<resource>.golden.yaml` (export is optional for now)

### A "full" test suite

* `create.yaml` that specifies all the configurable fields
* `update.yaml` that changes all the mutable fields
* `_http.log `
* `_generated_object_<resource>.golden.yaml `that contains `.status` that reflects the updated fields 
* `_generated_export_<resource>.golden.yaml` (export is optional for now)

### More dependencies test suite

* `create.yaml` and `update.yaml` that cover necessary fields
* `dependencies.yaml `to cover resource references 
    * No additional test suite if the resource ref is projectRef,
    * No separate test suite if the resource ref is a required field. You shall put this in the “basic” test suite. 
* `_http.log `contains not only the CRUD for your resource, but the CRUD for the dependent resources. 
* `_generated_object_<resource>.golden.yaml `that contains `.status` that reflects the updated fields 
* `_generated_export_<resource>.golden.yaml` (export is optional for now)
