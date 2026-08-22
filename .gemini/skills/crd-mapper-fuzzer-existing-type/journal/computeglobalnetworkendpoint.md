# ComputeGlobalNetworkEndpoint KRM Transition

## Observations and Learnings
- **Automatic Mapper Field Casing**: Renaming `Fqdn` to `FQDN` and `IpAddress` to `IPAddress` in the KRM Go struct matching the acronym casing expected by the controller-builder allows for automatic `FromProto` and `ToProto` generation of these fields, avoiding unnecessary manual mapping.
- **Spec-Only Fuzz Testing**: Since `ComputeGlobalNetworkEndpoint` does not have any status fields or custom status mapping, we should use `fuzztesting.RegisterKRMSpecFuzzer()` and `fuzztesting.NewKRMTypedSpecFuzzer[*pb.NetworkEndpoint, krm.ComputeGlobalNetworkEndpointSpec]` instead of a standard register which expects status mapper functions. This completely avoids `nil` dereference panics when executing status fuzz testing.
- **Port/ResourceID mismatch**: Handcoded `ComputeGlobalNetworkEndpointSpec_v1alpha1_FromProto` and `ComputeGlobalNetworkEndpointSpec_v1alpha1_ToProto` functions were written in `globalnetworkendpoint_mapper.go` to handle the `int32` (Port in Proto) vs `string` (ResourceID in KRM) type mismatch.
