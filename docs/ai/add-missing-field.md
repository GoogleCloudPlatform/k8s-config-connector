# How to add a missing field

When the GCP service adds a field to their proto, we must also add the corresponding field to our CRDs.  We also need to make sure that our code passes the field through.

Begin by identifying the proto field, look for it in `./.build/third_party/googleapis/`.  For example, if adding support for DNS Endpoints to ContainerCluster,
you should look for the field `DNSEndpointConfig dns_endpoint_config = 1;` in `.build/third_party/googleapis/google/container/v1/cluster_service.proto`.  It is nested under ControlPlaneEndpointsConfig.

Find the corresponding KRM type, under `apis/<service>/<version>/<message>_types.go`.  Use the `// +kcc:spec:proto` and `// +kcc:proto` annotations to help you.

Then add the corresponding field to the go type.  Follow the example of other fields.  Please be sure to include a structured comment like `// +kcc:proto:field=google.cloud.apigateway.v1.Api.state` so we can more easily map fields to proto fields.

Next, find a test and add test coverage for the missing field to one of the test cases.  Ideally the test case will be `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/full<kind>`, but we are not consistent.
