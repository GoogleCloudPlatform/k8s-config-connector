# 1. Add MockGCP tests

## Goals

The MockGCP provides presubmit coverage to your Config Connector resource:
- It validates your resources' fields behavior REQUIRED, OPTIONAL, (IM)MUTABLE.
- It guardrails your resources' when used by other resources as a reference.
- It mocks the real GCP to give a more stable and predictable outcome
- It improves the developing experience and reduces the long running operation waiting time from real GCP to mocked GCP in a factor of X * 10 times.
- It sets the test coverage for the follow-up direct controller developments. You don't need to write other tests after this step.

## General rules

MockGCP test is required for the direct resource development.

1. This is required no matter if the feature request is small or big. For example, both adding a single new field or adding an entire new resource require the MockGCP test coverage. If the MockGCP does not exist yet when adding a single field, developers should add the corresponding MockGCP coverage for the entire resource.  
1. Each Config Connector `spec` field should be covered in both create and update(if applicable) cases. 
1. MockGCP check is auto-enabled in Presubmit check and shall not be skipped. Code changes to a Config Connector resource must not merge if not covered by MockGCP check. 
1. New PRs should give steady golden log output before they can merge in. Reasonable golden object and http log changes are okay and sometimes expected.
1. MockGCP should be up-to-date when migrating resource from TF/DCL based controller to the direct controller
1. Real GCP record should be uploaded to reflect and validate the change in scenarios like adding a new resource, migrating from TF or DCL based controller to the direct controller, adding a new field, etc. *Please commit your real GCP record and mockGCP record as contiguous git-commits, so the PR reviewers can verify the MockGCP.*

### Suggestion

## 1.1 Add the test suite

Create a directory for your resource [pkg/test/resourcefixture/testdata/basic](pkg/test/resourcefixture/testdata/basic), following the naming convention as other directories. The bottommost directory name is the test suite name. You can create as many directory as your test needs.

Just to add the `create.yaml`,  `update.yaml` and `dependencies.yaml` (if applicable) to each bottommost directory.
These three files are all Kubernetes config files, except that `create.yaml` and `update.yaml` only holds the your Config Connector CR object, and the `dependencies.yaml` holds all the dependent objects.

Note: Do not add your GCP Project (the output of `gcloud config get project`) as a dependency in `dependencies.yaml` and run `hack/record-gcp` (even with the `cnrm.cloud.google.com/deletion-policy: abandon` annotation). If your GCP project is listed as a dependency, Config Connector will add the label `cnrm-test: true` to your project during test runs (because it labels all test resources with `cnrm-test: true`). This is likely not desirable, because the label indicates that your GCP project is a test resource. Instead, you can refer to your development project in the `create.yaml` and `update.yaml` resources via an external project reference, with the `${projectId}` template variable. For example:
```
  projectRef:
    external: ${projectId}
```

## 1.2 Write the MockGCP server

Follow [this guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/README.md) to write a mock gcp server.


## 1.3 Verify the MockGCP server

To write a new MockGCP server or to validate an existing one for the Direct resource, you should record the real GCP log and compare it with your MockGCP log, and make sure they are as likely as each other.

*Please commit your real GCP record and mockGCP record as contiguous git-commits, so the PR reviewers can verify the MockGCP.*

#### Record real GCP log

```
hack/record-gcp fixtures/<your_resource_test>
```

- `your_resource_test` is the folder directory name where you put the create.yaml and update.yaml files in. You can also use keywords to run multiple test suites. i.e. `hack/record-gcp fixtures/sql` matches all test suites that contains `sql`

```bash
git add <golden_file> <log_file> && git commit -m "log: real gcp"
```

- `golden_file` is the `_generated_object_<your_resource_test>.golden.yaml` file based on *real GCP reconciliation*.
- `log_file` is the `_http.log` file recording the *real GCP log*.


#### Record mock GCP log.

```
hack/compare-mock fixtures/<your_resource_test>
```

- `your_resource_test` is the folder directory name where you put the create.yaml and update.yaml files in. You can also use keywords to run multiple test suites. i.e. `hack/record-gcp fixtures/sql` matches all test suites that contains `sql`

```bash
git add <golden_file> <log_file> && git commit -m "log: real gcp"
```

- `golden_file` is the `_generated_object_<your_resource_test>.golden.yaml` file based on *mock GCP reconciliation*.
- `log_file` is the `_http.log` file recording the *mock GCP log*.

## Exit Criteria

For Config Connector **Beta** resources, your MockGCP should meet the following requirements.

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

### MockGCP Server

* Implement the mocked methods for CREATE, UPDATE, GET, and DELETE.
* If the Cloud API client returns long running operation (LRO), use `startRLO ` [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/8a350a029803a322e2889fd693cabf9780828c47/mockgcp/mockcloudbuild/workerpool.go#L56)
* CREATE method shall assign the “output-only” fields. [Example 1](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/5e08dbffb5fa3922dd43c451f35fdec45882205a/mockgcp/mockresourcemanager/tagkeys.go#L99C23-L99C37) [Example 2](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/611abaff651af81bed4517f62f915318f1b239bd/mockgcp/mocksql/sqlinstance.go#L68-L180)
* UPDATE method shall be able to update all mutable fields and support field masks if the real server does. [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/8a350a029803a322e2889fd693cabf9780828c47/mockgcp/mockcloudbuild/workerpool.go#L100)
* Make sure the log gives a stable outcome. You can override the value in<code> normalize.go </code>[example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/ba513862c2fb6ec3e54a05f6483c76b0337d6cbd/tests/e2e/normalize.go#L100)
