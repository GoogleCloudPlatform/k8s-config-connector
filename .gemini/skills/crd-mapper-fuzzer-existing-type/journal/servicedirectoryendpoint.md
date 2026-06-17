# Journal: ServiceDirectoryEndpoint Transition to Direct KRM types

## Learnings & Observations

### 1. Handling of missing ObservedState in CRD schema
Just like `ServiceDirectoryService`, the baseline CRD's `status` for `ServiceDirectoryEndpoint` does not have an `observedState` field.
To ensure the fuzzer and generated mapper compile successfully while maintaining 100% strict CRD schema compatibility:
- We defined an empty `ServiceDirectoryEndpointObservedState` struct in `endpoint_types.go`.
- We did NOT add an `observedState` field to the `ServiceDirectoryEndpointStatus` struct.
This allows mapping functions (`ServiceDirectoryEndpointObservedState_FromProto` / `ToProto`) to be generated without changing the CRD schema.

### 2. Custom Reference Types mapping
`ServiceDirectoryEndpoint` references multiple other resources:
- `serviceRef`: points to `ServiceDirectoryService`
- `addressRef`: points to a `ComputeAddress` (imported from `computev1beta1`)
- `networkRef`: points to a `ComputeNetwork` (imported from `computev1beta1`)

The automatic mapper generated clean from/to proto translation helpers for `AddressRef` and `NetworkRef` leveraging `computev1beta1` reference structs, while `serviceRef` is mapped in the parent controller loop.

### 3. Strict Schema Compatibility
Running `dev/tasks/diff-crds` returned an empty output, confirming that our types match the baseline CRD schema exactly with zero changes.
