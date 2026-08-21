### [2026-05-12] Missing protobufs for PlatformPolicy
- **Context**: Implementing `BinaryAuthorizationPlatformPolicy` direct resource types.
- **Problem**: `generate-types` failed because the `PlatformPolicy` protobuf message is missing from the public `googleapis` repository, even though it exists in REST (discovery). `generate-controller` also had a bug where the `--service` flag was required but not bound. Finally, `cloud.google.com/go/binaryauthorization/apiv1` did not support `PlatformPolicy`, forcing the use of the REST client `google.golang.org/api/binaryauthorization/v1`.
- **Solution**: 
  1. Manually constructed `mockgcp/apis/google/cloud/binaryauthorization/v1/platform_policy.proto` using the REST Discovery Document schema.
  2. Modified `generate-proto.sh` to compile it.
  3. Ran `generate-types` to scaffold types, then manually fixed the type references in the Go code.
  4. Rewrote `platformpolicy_controller.go` and `mapper.generated.go` to use the `google.golang.org/api/binaryauthorization/v1` REST client instead of the gRPC client since the gRPC client doesn't have the methods.
- **Impact**: Any agent implementing a newer GCP API that relies on Discovery docs rather than public protobufs will need to manually scaffold the protobufs into `mockgcp/apis/` or entirely skip `generate-types` and write the KRM and Mappers manually.
