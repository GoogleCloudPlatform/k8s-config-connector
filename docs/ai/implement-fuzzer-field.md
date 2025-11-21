# How to add an unimplemented field in a fuzzer

In a fuzzer, we mark some fields as unimplemented.  To fix one of them, we must add the field to the CRD, ensure it is mapped by the mapper, and then add an example to our tests.


1. Begin by identifying the proto field, look for it in `./.build/third_party/googleapis/`.  For example, if adding support for DNS Endpoints to ContainerCluster,
you should look for the field `DNSEndpointConfig dns_endpoint_config = 1;` in `.build/third_party/googleapis/google/container/v1/cluster_service.proto`.  It is nested under ControlPlaneEndpointsConfig.

1. Determine whether this is a spec or a status field.  Input fields are Spec fields.  Output fields are Status fields.

1. Find the corresponding KRM type, under `apis/<service>/<version>/<message>_types.go`.  Use the `// +kcc:spec:proto` and `// +kcc:proto` annotations to help you.

1. Then add the corresponding field to the go type.  Follow the example of other fields.  Please be sure to include a structured comment like `// +kcc:proto:field=google.cloud.apigateway.v1.Api.state` so we can more easily map fields to proto fields.

1. Next, find a test and add test coverage for the missing field to one of the test cases.  Ideally the test case will be `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<kind>-full`, but we are not consistent.

1. In the fuzzer, replace the call that marks the method as unimplemented with either a call to `f.Spec("<fieldName>")` or `f.Status("<fieldName>")`.  The fuzzer should be `pkg/controller/direct/<service>/<message>_fuzzer.go`
