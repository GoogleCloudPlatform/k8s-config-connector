# Journal: KRM Fuzzer for FilestoreInstance

## Background
The `FilestoreInstance` direct controller mapper contains several helper functions to convert KRM types to/from GCP protos.
The Status type for `FilestoreInstance` does not have an `ObservedState` sub-struct, but rather maps directly to the `Status` fields (`CreateTime`, `Etag`, `State`, `StatusMessage`).

## Implementation Details
1. **Added Status Mappers**: We handwrote `FilestoreInstanceStatus_FromProto` and `FilestoreInstanceStatus_ToProto` in `pkg/controller/direct/filestore/mapper.go` to cleanly map direct status fields like `CreateTime`, `Etag`, `State`, and `StatusMessage` to/from protobuf format.
2. **Created Fuzzer**: We implemented `pkg/controller/direct/filestore/filestoreinstance_fuzzer.go` declaring the KRM fuzzer matching the `Instance` proto.
3. **Categorized Fields**:
   - Spec fields: `.description`, `.file_shares`, `.networks`, `.tier`.
   - Status fields: `.create_time`, `.etag`, `.state`, `.status_message`.
   - Identity field: `.name`.
   - Unimplemented or unsupported fields (including `.labels` and `.networks[].connect_mode`) are cleanly marked as unimplemented using standard helpers.
4. **Anonymous Registration**: We registered the direct controller package anonymous import in `pkg/controller/direct/register/register.go` to ensure its automatic inclusion in central fuzz testing.
