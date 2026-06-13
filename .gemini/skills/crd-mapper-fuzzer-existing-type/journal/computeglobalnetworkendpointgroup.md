# Journal: ComputeGlobalNetworkEndpointGroup

Observations and learnings during the transition of `ComputeGlobalNetworkEndpointGroup` to direct KRM types.

## Type Mismatch Resolution
The `NetworkEndpointGroup` GCP protobuf type defines `default_port` as an `int32` pointer (`*int32`) and `network_endpoint_type` as a `string` pointer (`*string`), whereas the existing KRM/CRD baseline schema defines `defaultPort` as `*int64` and `networkEndpointType` as a plain `string`.
- This type mismatch results in compiler failures in automatically generated mapping assignments.
- We resolved this by creating a hand-coded mapper file (`pkg/controller/direct/compute/computeglobalnetworkendpointgroup_mapper.go`) containing:
  - `ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_FromProto`
  - `ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_ToProto`
- The generator automatically detects these pre-defined mapping functions and skips generating the conflicting assignments inside `mapper.generated.go`.

## Fuzzer Strategy for Spec-Only Mappers
Since this resource does not map any complex status fields (only `selfLink` and standard conditions), we did not define status mapping functions.
- To prevent generic type inference failures during fuzz testing, we used `fuzztesting.NewKRMTypedSpecFuzzer` to instantiate a spec-only fuzzer.
- We registered it using `fuzztesting.RegisterKRMSpecFuzzer` in `computeglobalnetworkendpointgroup_fuzzer.go` instead of `RegisterKRMFuzzer`. This avoids executing the status fuzzing logic, which would otherwise panic on `nil` mapping function pointers.
