# ComputeRoute KRM Types Transition Journal

## Learnings & Observations

1. **Custom Reference Types Without `kind` Field**:
   In the baseline CRD for `ComputeRoute`, references like `networkRef`, `nextHopILBRef`, `nextHopInstanceRef`, and `nextHopVPNTunnelRef` lack a `kind` field. Using `k8sv1alpha1.ResourceRef` introduces `kind=string` in the generated CRD schema, resulting in schema mismatch errors from `diff-crds`.
   - **Solution**: Hand-coded custom reference structs (`ComputeRouteNetworkRef`, `ComputeRouteNextHopILBRef`, etc.) matching the baseline reference schema perfectly (containing only `external`, `name`, and `namespace`). This preserved strict schema compatibility with the baseline CRD.

2. **Pointer required for auto-mapper**:
   Go pointer vs value types affect automatic mapping code generation. In particular, KRM reference fields should be declared as pointer types (`*ComputeRouteNetworkRef`) to avoid compilation issues in the generated `mapper.generated.go` file. Required validations can still be enforced at the spec-level in OpenAPI schema independently of the Go struct type.

3. **Type Mismatch Hand-coded Mapper**:
   When GCP protobuf fields have types like `*string` (optional string) or `*uint32` (optional uint32) but the CRD schema requires a plain string (`string`) or uses signed types (`*int32`), the automatic mapper will fail to compile.
   - **Solution**: Handcoded `ComputeRouteSpec_v1beta1_FromProto` and `ComputeRouteSpec_v1beta1_ToProto` in `pkg/controller/direct/compute/computeroute_mapper.go`. The generator detects these methods and correctly skips auto-generating conflicting versions.

4. **Fuzzing Casing & Path Constraints**:
   Proto paths in `fuzztesting` helper methods can match the proto field casing directly (e.g. `.next_hop_ip`). Standard fuzzer feedback indicates precisely which fields need to be marked as `Unimplemented_NotYetTriaged` to round-trip flawlessly.
