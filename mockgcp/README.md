# MockGCP 

MockGCP contains stub implementations of GCP APIs, for use with testing Config Connector. The implementations
don't do anything (there is no "backend") - creating a mock VM does not launch a VM.  Instead, the
goal is to provide just-enough consistent CRUD operations to test the Config Connector providers.
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

To add a new service, the easiest way is to copy one of the existing PRs. Example PR: [mockGCP for cloudfunctions](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/869)

Broadly the steps are:

1. Add the proto to the Makefile and run `make gen-proto` (or just `make`).

   All google services are located in [googleapis GitHub repo](https://github.com/googleapis/googleapis/tree/master/google),
   refer to your resource's API documentation to identify the service name, for example [privateca](https://cloud.google.com/certificate-authority-service/docs/reference/rest#service:-privateca.googleapis.com).
   Once you identify the service, find the proper path to the proto files, for example:
   `cloud/security/privateca/v1/*.proto`. Then replace the prefix `googleapis/google/` to `./third_party/googleapis/mockgcp/`,
   and add into the Makefile.

   * Note: Ensure you pick the same version the controller uses to call the GCP API during GCP resource instantiation.

   * Note: If the Config Connector resource does not map to a GCP resource, please refer to the REST API of the relevant GCP resource and find the proto file that covers the resource's API endpoints.(You may not find a proto file that matches the name of that Config Connector resource)


   * (Optional) If you determine that the proto file is not up to date, or if it doesn't exist at all, refer to the [Generating Proto](#generating-proto) section
   
1. (Optional). If you're adding an API outside of `googleapis/google/cloud`,
   you may need to add commands to rename the API to mockgcp in [fixup-third-party.sh](fixup-third-party.sh). Example:
   ```
   ...
   mv google/logging/ mockgcp/
   ...
   find . -type f -print0 | xargs -0 sed -i -e "s@google/logging/@mockgcp/logging/@g"
   find . -type f -print0 | xargs -0 sed -i -e "s@google\.logging@mockgcp.logging@g"
   ...
   ```

1. Add a subdirectory called `mock<servicename>`.

   * Copying one of the existing ones. `mockprivateca` is a reasonable basic one. Note the files `names.go` and `service.go`,
   rename the third one to `<resourcename>.go`. We have started collapsing the contents of the `names.go` files into the relevant `<resourcename>.go` files.
1. Adjust your mock to use your protos (change the import paths), rename the types, etc.
1. Implement the core CRUD methods in `<resourcename>.go`.

   In the example file `capool.go`, we have GET, CREATE and DELETE implemented. More probably, you will need to implement
   some more CRUD methods.  In the log, you should see some errors indicating which method is not supported.
   For those methods, go into the service definition and implement a basic implementation - we have
   examples of most of the CRUD operations at this point.

   Note that until you started using the mock package you defined, VSCode may highlight a warning that the package is not used. See later steps for where to use the new mock package.
1. Register the mock service of your resource in the service.go file.
   [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/d10e4ac6241a454c995006ce2c83b5c4d20bb510/mockgcp/mockaiplatform/service.go#L58).
1. Add the service handler of your resource in the service.go file.
   [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/d10e4ac6241a454c995006ce2c83b5c4d20bb510/mockgcp/mockaiplatform/service.go#L62).
1. Add your service to mock_http_roundtrip.go, something like `services = append(services, mockcloudbilling.New(env, storage))`.
1. Add the kind(s) you are adding support for to `func MaybeSkip` in
   [harness.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/config/tests/samples/create/harness.go).
1. Run the tests.

   Example command: 
   ```
   E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures  2>&1 | tee log
   ```

   You can also specify a single test to run to save time, for example, `TestAllInSeries/fixtures/privatecacapool`, or
   reduce the timeout from 3600s to 600s or any other reasonable value depends on your test.

If you are lucky, everything will "just work" and you will see that your new tests are being invoked.

If something is not behaving as you would expect, you should be able to launch a debugger because it all runs in one process.
You can also use `ARTIFACTS=artifacts` to get detailed HTTP logs of the traffic, which is useful if you want to see the json requests & responses.
If you also use `E2E_GCP_TARGET=real` you can run against the real (non-mocked) GCP, and easily see what the actual behaviour should be.
Usually however, this is not necessary; the most common failure mode is that terraform or Config Connector expects a field to be automatically populated,
and it normally logs an error like "foo not set" (in this case, simply add that to your mock implementation.)

## Capture golden object and HTTP golden logs

1. Capture golden object and HTTP golden logs against real GCP.

   1. Run the following command to generate the golden object (`_generated_object_\[testname\].golden.yaml` file):
      `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=real GOLDEN_OBJECT_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/[testname]`.

   1. Run the following command to generate the golden request (`_http.log` file):
      `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=real GOLDEN_REQUEST_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/[testname]`.

   1. Ensure both `GOLDEN_OBJECT_CHECKS` and `GOLDEN_REQUEST_CHECKS` pass:
      `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=real GOLDEN_OBJECT_CHECKS=1 GOLDEN_REQUEST_CHECKS=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/[testname]`.

      1. If the test doesn't pass, check the output log and identify diffs. Normalize the values if needed in
         [tests/e2e/normalize.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/v1.120.1/tests/e2e/normalize.go#L66)
         for GOLDEN_OBJECT_CHECKS and in
         [tests/e2e/unified_test.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/v1.120.1/tests/e2e/unified_test.go#L523)
         for GOLDEN_REQUEST_CHECKS.

1. Ensure the test works when `E2E_GCP_TARGET=mock`:
   `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock GOLDEN_OBJECT_CHECKS=1 GOLDEN_REQUEST_CHECKS=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/[testname]`.

   1. If the test doesn't pass, update the CRUD methods in the mock to match the behavior needed by the golden files.

1. **In some edge cases**, we can consider getting HTTP logs against mockGCP after implementing the mockGCP service.
   For certain resources that are difficult to run against real GCP (for example, AttachedCluster requires an
   external cluster and EdgeContainer requires a physical machine, and etc), it's acceptable to capture the golden
   HTTP logs from the mock service.

   Example command: `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock GOLDEN_REQUEST_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/test-name`.

## Appendix

### Generating proto

If the proto file on [googleapis](https://github.com/googleapis/googleapis/commits/1e4137870560340a14700618a05e2d7162326af7/google/cloud/ids/v1/ids.proto) is out of date or nonexistent, you can generate a proto file from the [google generated api](https://github.com/googleapis/google-api-go-client/tree/b49e3b908a8ed562e068736f1c42e992538ba6e0) as such:

```shell
$ 	wget -O ids-api-v1.json https://raw.githubusercontent.com/googleapis/google-api-go-client/b49e3b908a8ed562e068736f1c42e992538ba6e0/ids/v1/ids-api.json
	mkdir -p apis/mockgcp/cloud/ids/v1/
	cd tools/gapic; go run . --proto-version=2 ../../ids-api-v1.json > ../../apis/mockgcp/cloud/ids/v1/service.proto
```

Modify the snippet above for your own service and add it to the Makefile target [`generate-proto-from-openapi`](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/bbdd7e244a8e9c1259ab939aa233c63fb38db1c2/mockgcp/Makefile#L73-L74).
