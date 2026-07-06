# Journal Entry - ComputeTargetGRPCProxy

## Observations
- `ComputeTargetGRPCProxy` uses the proto message `TargetGrpcProxy` from the Compute API (`google.cloud.compute.v1.TargetGrpcProxy`).
- The KRM Kind has "GRPC" in uppercase, whereas the Proto name uses "Grpc" in camel case. Under `apis/compute/v1beta1/`, the lowercase file name must match the proto message, i.e., `targetgrpcproxy_types.go`.
- The resource reference field `urlMapRef` points to a `ComputeURLMap` resource, which is mapped via the shared `ComputeURLMapRef` defined in `urlmap_reference.go`.
- The `generate-mapper` tool commented out mapping of `urlMapRef` inside `mapper.generated.go` with a `// MISSING: URLMap` comment because it's a reference type mapped to a string field in the proto.
- We handcoded a simple mapper in `targetgrpcproxy_mapper.go` to handle mapping of the reference `urlMapRef.External` to the string field `UrlMap` in proto, which was automatically detected and skipped by the generator.
- We defined a fuzzer in `targetgrpcproxy_fuzzer.go` using the required type-safe helpers.
