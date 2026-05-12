### [2026-05-12] CloudIdentityDevice Missing Fields
- **Context**: Implementing CloudIdentityDevice resource via the direct approach.
- **Problem**: KCC's `generate-types` skips fields from `Device` protobuf because they are documented as "Output only." in the comments (which KCC ignores if they are missing from existing type files but then `generate-mapper` complains about them being MISSING). Additionally, the API documentation for `CreateDeviceRequest` specifies that `device_type` and `serial_number` are REQUIRED for creation, even though `device_type` is marked `Output only.` in the proto. Furthermore, `generate-mapper` will fail if trying to generate mappers for `cloudidentitypb` because KCC does not depend on the generated proto struct module for `cloudidentity`.
- **Solution**: We must manually map `Device` using `api.Device` (from `google.golang.org/api/cloudidentity/v1beta1`) instead of the generated `mapper.generated.go`. Also, `DeviceType` and `SerialNumber` must be put in the `Spec` since they are required for creation. All other read-only properties go to `ObservedState`. Finally, the `Customer` field is required by `CreateDeviceRequest` and we default it to `customers/my_customer`.
- **Impact**: Future agents working on CloudIdentity should know that some proto comments like `Output only` might be incorrect for `Create` requests. Also, manual mapping to the REST API is required since the GRPC proto is not available in KCC's module dependencies.

### [2026-05-12] WaitForCloudIdentityOp Polling Bug
- **Context**: Implementing LRO wait in CloudIdentityDevice controller.
- **Problem**: The existing WaitForCloudIdentityOp function in pkg/controller/direct/cloudidentity/lro.go does not actually fetch the updated operation from the API during its loop. It just checks the initial op.Done value repeatedly, which would lead to a timeout if the operation is not immediately complete.
- **Solution**: The function needs to be rewritten to actually call the GCP API gcpClient.Operations.Get(op.Name) to check the updated state. For now, it is used as-is following existing patterns.
- **Impact**: Any long-running operations in CloudIdentity that do not complete synchronously will fail with a timeout.

