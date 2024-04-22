# MockGCP

MockGCP contains stub implementations of GCP APIs, for use with testing KCC.  The implementations
don't do anything (there is no "backend") - creating a mock VM does not launch a VM.  Instead, the
goal is to provide just-enough consistent CRUD operations to test the KCC providers.
If the tests pass, we consider that to be a "good enough" mock for GCP.  When
we discover an issue that the mocks failed to catch, we should update mockgcp
to more accurately simulate GCP, ideally triggering the bug and then showing
that the fix addresses the issue.

At their core, most GCP APIs are CRUD-based, so implementing a basic implementation is fairly straightforward.
We are establishing some patterns to make the implementation as simple as possible, such as a simple
storage driver that works with the existing protos.

In general, we use the protos published by GCP at [https://github.com/googleapis/googleapis.git](https://github.com/googleapis/googleapis.git).  These protos make it relatively easy to implement the CRUD operations needed.

In addition, the protos have the annotations needed for GRPC <-> HTTP interop.
We use [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to create HTTP servers from GRPC servers.

It is then fairly straightforward to inject our own HTTP handlers directly into the serving paths
for tests, so we don't even need to start a real webserver (i.e. we don't even need to listen on a port)

## Adding a new service

To add a new service, the easiest way is to copy one of the existing CLs. Example CL: [mockGCP for cloudfunctions](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/869)

Broadly the steps are:

1. Add the proto to the Makefile and run `make gen-proto` (or just `make`, as it's the default target).

   All google services are located in [googleapis Github repo](https://github.com/googleapis/googleapis/tree/master/google),
   refer to your resource's API documentation to identify the service name, for example [privateca](https://cloud.google.com/certificate-authority-service/docs/reference/rest#service:-privateca.googleapis.com).
   Once you identify the service, find the proper path to the proto files, for example:
   `cloud/security/privateca/v1/*.proto`. Then replace the prefix `googleapis/google/` to `./third_party/googleapis/mockgcp/`,
   and add into the Makefile.

1. (Optional). If you're adding an API outside of googleapis/google/cloud,
   you may need to add commands to rename the API o mockgcp in [fixup-third-party.sh](fixup-third-party.sh). Example:
   ```
   ...
   mv google/logging/ mockgcp/
   ...
   find . -type f -print0 | xargs -0 sed -i -e "s@google/logging/@mockgcp/logging/@g"
   find . -type f -print0 | xargs -0 sed -i -e "s@google\.logging@mockgcp.logging@g"
   ...
   ```

1. Add a subdirectory called `mock<servicename>`.

   Copying one of the existing ones. `mockprivateca` is a reasonable basic one. Keep the files names.go and service.go,
   rename the third one to `<resourcename>.go`.
1. Adjust your mock to use your protos (change the import paths), rename the types, etc.
1. Implement the core CRUD methods in `<resourcename>.go`.

   In the example file `capool.go`, we have GET, CREATE and DELETE implemented. More probably, you will need to implement
   some more CRUD methods.  In the log, you should see some errors indicating which method is not supported.
   For those methods, go into the service definition and implement a basic implementation - we have
   examples of most of the CRUD operations at this point.
1. Register the mock service of your resource in the service.go file.
   [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/d10e4ac6241a454c995006ce2c83b5c4d20bb510/mockgcp/mockaiplatform/service.go#L58).
1. Add the service handler of your resource in the service.go file.
   [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/d10e4ac6241a454c995006ce2c83b5c4d20bb510/mockgcp/mockaiplatform/service.go#L62).
1. Add your service to mock_http_roundtrip.go, something like `services = append(services, mockcloudbilling.New(env, storage))`.
1. Add the kind(s) you are adding support for to `func MaybeSkip` in
   [harness.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/config/tests/samples/create/harness.go).
1. Run test and compare mock result.

   1. Run the test against the real GCP, and record the HTTP requests and responses. The log can be use to identify the fields that needs to be mocked.

      ```shell
      ./hack/record-gcp fixtures/[TEST_NAME]
      ```

   1. Commit the `_http.log` within the test folder but do not push it. Run the test against the mock target to compare if any fields changed. It normally logs an error like "foo not set" (in this case, simply add that to your mock implementation.)

      ```shell
      ./hack/compare-mock fixtures/[TEST_NAME]
      ```

   1. If there are fields keeps changing that keeps changing everytime you run the test such as `createdTime` or `id`. You need to adjust normalize logic in [unified_test.go](tests/e2e/unified_test.go) to produce reliable test result.

   1. If there are fields containing sensitive/confidential information in the log. The fields needs to be removed. (Example)[https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/tests/e2e/unified_test.go#L456-L462].

## Capture HTTP golden logs

Capture HTTP logs against real GCP:

Example command: `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=real GOLDEN_REQUEST_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/test-name`.

Capture HTTP logs against mock GCP:

We can consider getting HTTP logs against mockGCP after implementing the mockGCP service.
For certain resources that are difficult to run against real GCP (for example, AttachedCluster requires an external cluster and EdgeContainer requires a physical machine, and etc),
it's acceptable to capture the golden HTTP logs from the mock service.

Example command: `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock GOLDEN_REQUEST_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/test-name`.

A new file named "_http.log" will be created in the test folder during the first time the command is executed, that's our golden HTTP log.
We will compare new HTTP logs with the golden HTTP log as a part of the mockGCP test if GOLDEN_REQUEST_CHECKS is specified.

To prevent any problems when comparing with the golden logs, it is necessary to replace all generated values in the HTTP log with identical values.

See [here](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/tests/e2e/unified_test.go#L167-L329) to get some ideas.