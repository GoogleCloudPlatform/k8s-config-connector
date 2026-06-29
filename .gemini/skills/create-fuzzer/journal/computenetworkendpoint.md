# Fuzzer Journal: ComputeNetworkEndpoint

## Resource Details
- **Kind**: `ComputeNetworkEndpoint`
- **Group**: `compute.cnrm.cloud.google.com`
- **Version**: `v1alpha1`
- **Proto**: `google.cloud.compute.v1.NetworkEndpoint`

## Observations & Findings
1. **Spec-only Mapping**: `ComputeNetworkEndpoint` is a lightweight resource that is managed inside a Network Endpoint Group (NEG). The KRM Spec maps to `NetworkEndpoint`, but KRM Status contains no GCP-managed fields representing observed state (only typical `Conditions` and `ObservedGeneration`). Therefore, we hand-coded trivial status mappers `FromProto` and `ToProto` returning zero-value structs/nil.
2. **IPAddress Case Near-miss**: The auto-generated mapper did not map `ipAddress` field automatically because of a near-miss: KRM uses `IpAddress` (lowercase p), whereas protobuf uses `IpAddress` (lowercase p) but the generator encountered an issue trying to match it (indicated by `// (near miss): "IPAddress" vs "IpAddress"` in `mapper.generated.go`). We hand-coded the mapping in `computenetworkendpoint_mapper.go` to handle it.
3. **ResourceID to Port Mapping**: The `ResourceID` (which represents the port number in string format) maps to the `Port` field (an `*int32`) in `pb.NetworkEndpoint`. We added explicit string-to-int32 conversions in our hand-coded spec mappers.
4. **Identity Field Mapping**: `pb.NetworkEndpoint` does not have a `.name` field in protobuf; its identity is composed of its `Instance`, `IpAddress`, and `Port` inside a NEG. Thus, we did not declare a `.name` field under `f.Unimplemented_Identity()`.
