# ComputeRoute Fuzzer Journal

## Observations & Learnings

- **Fuzzer Implementation:** Implemented a new, round-trip KRM fuzzer for the `ComputeRoute` direct resource under `pkg/controller/direct/compute/computeroute_fuzzer.go`.
- **Status Mapper Mapping:** Added missing `ComputeRouteStatus_v1beta1_FromProto` and `ComputeRouteStatus_v1beta1_ToProto` mapper functions in `pkg/controller/direct/compute/computeroute_mapper.go` to support round-trip validation of status fields (`nextHopNetwork` and `selfLink`).
- **Unimplemented Field Detection:** Detected `.next_hop_interconnect_attachment` as an unimplemented field via the central fuzz testing framework and registered it using `f.Unimplemented_NotYetTriaged()`.
- **Validation:** Verified the fuzzer against the central fuzz testing framework using `FOCUS=ComputeRoute go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`, which passed successfully.
